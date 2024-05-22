package userusecase

import (
	"context"
	"database/sql"
	"fmt"
	userdomain "shopito/api/internal/domain/user"
	userrepository "shopito/api/internal/repository/user"
	"shopito/api/pkg/types/errors"
	"time"

	"github.com/sirupsen/logrus"
)

type UserUseCase interface {
	GetUsers(ctx context.Context) (*[]userdomain.User, *errors.Error)
	GetUser(ctx context.Context, id int) (*userdomain.User, *errors.Error)
	UpdateUser(ctx context.Context, id int, user *userdomain.User) *errors.Error
	DeleteUser(ctx context.Context, id int) *errors.Error
}

type UserUseCaseImpl struct {
	userrepo userrepository.UserRepository
}

func New(db *sql.DB) *UserUseCaseImpl {
	return &UserUseCaseImpl{
		userrepo: userrepository.New(db),
	}
}

func (u *UserUseCaseImpl) GetUsers(ctx context.Context) (*[]userdomain.User, *errors.Error) {
	users, err := u.userrepo.GetAll()
	if err != nil {
		logrus.WithField("err", err.Error()).Error("db err")
		return nil, errors.ErrInternal.SetMessage("db err")
	}
	return users, nil
}

func (u *UserUseCaseImpl) GetUser(ctx context.Context, id int) (*userdomain.User, *errors.Error) {
	if !u.userrepo.ExistById(id) {
		return nil, errors.ErrNotFound.SetMessage(fmt.Sprintf("user with id=%v not found", id))
	}
	user, err := u.userrepo.GetById(id)
	if err != nil {
		logrus.WithField("err", err.Error()).Error("db err")
		return nil, errors.ErrInternal.SetMessage("Internal Server Error")
	}
	return user, nil
}

func (u *UserUseCaseImpl) UpdateUser(ctx context.Context, id int, user *userdomain.User) *errors.Error {
	if !u.userrepo.ExistById(id) {
		return errors.ErrNotFound.SetMessage(fmt.Sprintf("user with id=%v not found", id))
	}
	user.UpdatedAt = time.Now()
	if err := u.userrepo.Update(id, user); err != nil {
		logrus.WithField("err", err.Error()).Error("db err")
		return errors.ErrInternal.SetMessage("Internal Server Error")
	}
	return nil
}

func (u *UserUseCaseImpl) DeleteUser(ctx context.Context, id int) *errors.Error {
	if !u.userrepo.ExistById(id) {
		return errors.ErrNotFound.SetMessage(fmt.Sprintf("user with id=%v not found", id))
	}
	if err := u.userrepo.Delete(id); err != nil {
		logrus.WithField("err", err.Error()).Error("db err")
		return errors.ErrInternal.SetMessage("Internal Server Error")
	}
	return nil
}
