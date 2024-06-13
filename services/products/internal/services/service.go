package service

import (
	productproto "shopito/pkg/protobuf/products"
	"shopito/services/products/internal/repository"
)

type Service interface {
	CreateProductService(product *productproto.Product) (int64, error)
	DeleteProductService(id int64) error
	GetProductService(id int64) (*productproto.Product, error)
	UpdateProductService(id int64, product *productproto.Product) error
	ListProductService(limit, offset int32, filter string) ([]*productproto.Product, error)
	ListProductByCategoryService(category string, limit, offset int32) ([]*productproto.Product, error)

	CreateCategoryService(category *productproto.Category) (int64, error)
	DeleteCategoryService(id int64) error
	GetCategoryService(id int64) (*productproto.Category, error)
	UpdateCategoryService(id int64, category *productproto.Category) error
	ListCategoriesService() ([]*productproto.Category, error)
}

type ProductService struct {
	repo *repository.Queries
}

func New(repo *repository.Queries) *ProductService {
	return &ProductService{
		repo: repo,
	}
}
