package adminusecase

import (
	"context"
	"database/sql"
	"fmt"
	"shopito/api/config"
	userrepository "shopito/api/internal/repository/user"
	"shopito/api/pkg/types/errors"
	"shopito/api/pkg/types/response"
	emailutil "shopito/api/pkg/util/email"

	"github.com/sirupsen/logrus"
)

type AdminUseCase interface {
	PromoteUser(ctx context.Context, id int) *errors.Error
	NotifyUsers(ctx context.Context, message response.EmailMessage) *errors.Error
	DemoteUser(ctx context.Context, id int) *errors.Error
}

type AdminUseCaseImpl struct {
	userrepo userrepository.UserRepository
}

func New(db *sql.DB) *AdminUseCaseImpl {
	return &AdminUseCaseImpl{
		userrepo: userrepository.New(db),
	}
}

func (u *AdminUseCaseImpl) PromoteUser(ctx context.Context, id int) *errors.Error {
	if !u.userrepo.ExistById(id) {
		return errors.ErrNotFound.SetMessage(fmt.Sprintf("users with id=%v not found", id))
	}
	user, err := u.userrepo.GetById(id)
	if err != nil {
		return errors.ErrInternal.SetMessage("Internal Server Error")
	}
	user.IsAdmin = true
	if err := u.userrepo.Update(id, user); err != nil {
		return errors.ErrInternal.SetMessage("Internal Server Error")
	}
	return nil
}

func (u *AdminUseCaseImpl) DemoteUser(ctx context.Context, id int) *errors.Error {
	if !u.userrepo.ExistById(id) {
		return errors.ErrNotFound.SetMessage(fmt.Sprintf("users with id=%v not found", id))
	}
	user, err := u.userrepo.GetById(id)
	if err != nil {
		return errors.ErrInternal.SetMessage("Internal Server Error")
	}
	user.IsAdmin = false
	if err := u.userrepo.Update(id, user); err != nil {
		return errors.ErrInternal.SetMessage("Internal Server Error")
	}
	return nil
}

func (u *AdminUseCaseImpl) NotifyUsers(ctx context.Context, message response.EmailMessage) *errors.Error {
	users, err := u.userrepo.GetAll()
	if err != nil {
		return errors.ErrInternal.SetMessage("couldn't load the users")
	}
	credentials := emailutil.New(
		config.EMAIL.USERNAME,
		config.EMAIL.PASSWORD,
		config.EMAIL.HOST,
		config.EMAIL.PORT,
		config.EMAIL.FROM,
	)
	for _, user := range *users {
		err := credentials.SendEmail(user.Email, message.Subject, message.Body)
		if err != nil {
			logrus.WithError(err).Error("couldn't send email")
		}
	}
	return nil
}
