package authusecase

import (
	"context"
	"database/sql"
	"fmt"
	"shopito/api/config"
	userdomain "shopito/api/internal/domain/user"
	userrepository "shopito/api/internal/repository/user"
	verifyrepository "shopito/api/internal/repository/verification"
	"shopito/api/pkg/types/errors"
	emailutil "shopito/api/pkg/util/email"
	jwtutil "shopito/api/pkg/util/jwt"

	"github.com/sirupsen/logrus"
)

type AuthUseCase interface {
	RegisterUser(ctx context.Context, user *userdomain.User) (int, *errors.Error)
	VerifyUser(ctx context.Context, id int, token string) *errors.Error
	RetryVerification(ctx context.Context, id int) *errors.Error
	LoginUser(ctx context.Context, user *userdomain.User) (*string, *errors.Error)
}

type AuthUseCaseImpl struct {
	userrepo   userrepository.UserRepository
	verifyrepo verifyrepository.VerifyRepository
}

func New(db *sql.DB) *AuthUseCaseImpl {
	return &AuthUseCaseImpl{
		userrepo:   userrepository.New(db),
		verifyrepo: verifyrepository.New(db),
	}
}

func (u *AuthUseCaseImpl) LoginUser(ctx context.Context, user *userdomain.User) (*string, *errors.Error) {
	p := user.Password
	user, err := u.userrepo.GetByEmail(user.Email)
	if err != nil {
		return nil, errors.ErrNotFound.SetMessage(fmt.Sprintf("user not found with email=%v", user.Email))
	}
	if !user.AuthenticatePassword(p) {
		return nil, errors.ErrNotAuthorized.SetMessage("bad credentials: email or password is incorrect")
	}
	token, err := jwtutil.GenerateToken(user.ID, user.IsAdmin, user.IsVerified)
	if err != nil {
		logrus.WithField("err", err.Error()).Error("token gen err")
		return nil, errors.ErrInternal.SetMessage("token generation error")
	}
	return token, nil
}

func (u *AuthUseCaseImpl) RegisterUser(ctx context.Context, user *userdomain.User) (int, *errors.Error) {
	if err := user.IsValid(); err != nil {
		return -1, errors.ErrBadRequest.SetMessage(err.Error())
	}
	if err := user.HashPassword(); err != nil {
		return -1, errors.ErrBadRequest.SetMessage(err.Error())
	}
	if u.userrepo.ExistByEmail(user.Email) {
		return -1, errors.ErrConflict.SetMessage("such email already exists")
	}
	id, err := u.userrepo.Insert(user)
	if err != nil {
		return -1, errors.ErrInternal.SetMessage("couldn't perform an operation")
	}

	code := emailutil.GenerateCode()
	subject := "Welcome, please verify your email address"
	body := fmt.Sprintf("Your confirmation code: %v", code)
	error := emailutil.New(
		config.EMAIL.USERNAME,
		config.EMAIL.PASSWORD,
		config.EMAIL.HOST,
		config.EMAIL.PORT,
		config.EMAIL.FROM,
	).SendEmail(user.Email, subject, body)
	if error != nil {
		u.userrepo.Delete(id)
		return -1, errors.ErrBadRequest.SetMessage("email is not valid")
	}
	u.verifyrepo.Insert(id, code)
	return id, nil
}

func (u *AuthUseCaseImpl) RetryVerification(ctx context.Context, id int) *errors.Error {
	code := emailutil.GenerateCode()
	user, err := u.userrepo.GetById(id)
	if err != nil {
		return errors.ErrNotFound.SetMessage(fmt.Sprintf("user not found with id=%v", id))
	}

	subject := "Welcome, please verify your email address"
	body := fmt.Sprintf("Your confirmation code: %v", code)

	error := emailutil.New(
		config.EMAIL.USERNAME,
		config.EMAIL.PASSWORD,
		config.EMAIL.HOST,
		config.EMAIL.PORT,
		config.EMAIL.FROM,
	).SendEmail(user.Email, subject, body)
	if error != nil {
		u.userrepo.Delete(id)
		return errors.ErrBadRequest.SetMessage("email is not valid")
	}
	if err := u.verifyrepo.Update(id, code); err != nil {
		return errors.ErrBadRequest.SetMessage("email is not valid")
	}
	return nil
}

func (u *AuthUseCaseImpl) VerifyUser(ctx context.Context, id int, token string) *errors.Error {
	if !u.userrepo.ExistById(id) {
		return errors.ErrNotFound.SetMessage(fmt.Sprintf("user with id=%v doesn't exist", id))
	}
	codeSource, err := u.verifyrepo.GetCodeByUser(id)
	if err != nil {
		return errors.ErrInternal.SetMessage("couldn't access the verify token")
	}
	if codeSource != token {
		return errors.ErrNotAuthorized.SetMessage("wrong verification token")
	}
	u.verifyrepo.DeleteByUser(id)
	user, err := u.userrepo.GetById(id)
	if err != nil {
		return errors.ErrNotFound.SetMessage(fmt.Sprintf("user with id=%v doesn't exist", id))
	}
	user.IsVerified = true
	if err := u.userrepo.Update(id, user); err != nil {
		return errors.ErrInternal.SetMessage("couldn't verify the user")
	}
	return nil
}
