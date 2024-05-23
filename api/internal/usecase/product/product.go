package productusecase

import (
	"context"
	"database/sql"
	"fmt"
	productdomain "shopito/api/internal/domain/product"
	productrepository "shopito/api/internal/repository/product"
	"shopito/api/pkg/types/errors"
)

type ProductUseCase interface {
	InsertProduct(ctx context.Context, product *productdomain.Product) (int, *errors.Error)
	GetProduct(ctx context.Context, id int) (*productdomain.Product, *errors.Error)
	DeleteProduct(ctx context.Context, id int) *errors.Error
	UpdateProduct(ctx context.Context, id int, product *productdomain.Product) *errors.Error
}

type ProductUseCaseImpl struct {
	productrepo productrepository.ProductRepository
}

func New(db *sql.DB) *ProductUseCaseImpl {
	return &ProductUseCaseImpl{
		productrepo: productrepository.New(db),
	}
}

func (u *ProductUseCaseImpl) InsertProduct(ctx context.Context, product *productdomain.Product) (int, *errors.Error) {
	id, err := u.productrepo.Insert(product)
	if err != nil {
		return -1, errors.ErrInternal.SetMessage("Internal Server Error")
	}
	return id, nil
}

func (u *ProductUseCaseImpl) GetProduct(ctx context.Context, id int) (*productdomain.Product, *errors.Error) {
	if !u.productrepo.Exist(id) {
		return nil, errors.ErrNotFound.SetMessage(fmt.Sprintf("product not found with id=%v", id))
	}
	product, err := u.productrepo.GetById(id)
	if err != nil {
		return nil, errors.ErrInternal.SetMessage("Internal Server Error")
	}
	return product, nil
}

func (u *ProductUseCaseImpl) DeleteProduct(ctx context.Context, id int) *errors.Error {
	if !u.productrepo.Exist(id) {
		return errors.ErrNotFound.SetMessage(fmt.Sprintf("product not found with id=%v", id))
	}
	if err := u.productrepo.Delete(id); err != nil {
		return errors.ErrInternal.SetMessage("Internal Server Error")
	}
	return nil
}

func (u *ProductUseCaseImpl) UpdateProduct(ctx context.Context, id int, product *productdomain.Product) *errors.Error {
	if !u.productrepo.Exist(id) {
		return errors.ErrNotFound.SetMessage(fmt.Sprintf("product not found with id=%v", id))
	}
	if err := u.productrepo.Update(id, product); err != nil {
		return errors.ErrInternal.SetMessage("Internal Server Error")
	}
	return nil
}
