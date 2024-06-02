package service

import (
	productproto "shopito/pkg/protobuf/products"
	categoryrepository "shopito/services/products/internal/repository/category"
	productrepository "shopito/services/products/internal/repository/product"
	"strconv"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	CreateProductService(product *productproto.Product) (int, error)
	DeleteProductService(id string) error
	GetProductService(id string) (*productproto.Product, error)
	UpdateProductService(id string, product *productproto.Product) error
	CreateCategoryService(category *productproto.Category) (int, error)
	DeleteCategoryService(id string) error
	GetCategoryService(id string) (*productproto.Category, error)
	UpdateCategoryService(id string, category *productproto.Category) error
	ListCategoriesService() ([]*productproto.Category, error)
}

type ProductService struct {
	productRepo  productrepository.Repository
	categoryRepo categoryrepository.Repository
}

func New(productRepo *productrepository.ProductRepository, categoryRepo *categoryrepository.CategoryRepository) *ProductService {
	return &ProductService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *ProductService) CreateProductService(product *productproto.Product) (int, error) {
	id, err := s.productRepo.Insert(product)
	if err != nil {
		logrus.WithError(err).Error("Internal Server Error")
		return -1, status.Errorf(codes.Internal, "Internal Server error")
	}
	return id, nil
}

func (s *ProductService) DeleteProductService(id string) error {
	productId, err := strconv.Atoi(id)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Couldn't process the id")
	}
	if !s.productRepo.Exist(productId) {
		return status.Errorf(codes.NotFound, "Product not found")
	}
	if err := s.productRepo.Delete(productId); err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error")
	}
	return nil
}

func (s *ProductService) GetProductService(id string) (*productproto.Product, error) {
	productId, err := strconv.Atoi(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Couldn't process the id")
	}
	if !s.productRepo.Exist(productId) {
		return nil, status.Errorf(codes.NotFound, "Product not found")
	}
	product, err := s.productRepo.Get(productId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	return product, nil
}

func (s *ProductService) UpdateProductService(id string, product *productproto.Product) error {
	productId, err := strconv.Atoi(id)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Couldn't process the id")
	}
	if err := s.productRepo.Update(productId, product); err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error")
	}
	return nil
}

func (s *ProductService) CreateCategoryService(category *productproto.Category) (int, error) {
	id, err := s.categoryRepo.Insert(category)
	if err != nil {
		logrus.WithError(err).Error("Internal Server Error")
		return -1, status.Errorf(codes.Internal, "Internal Server error")
	}
	return id, nil
}

func (s *ProductService) DeleteCategoryService(id string) error {
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Couldn't process the id")
	}
	if !s.categoryRepo.Exist(categoryId) {
		return status.Errorf(codes.NotFound, "Category not found")
	}
	if err := s.categoryRepo.Delete(categoryId); err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error")
	}
	return nil
}

func (s *ProductService) GetCategoryService(id string) (*productproto.Category, error) {
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Couldn't process the id")
	}
	if !s.categoryRepo.Exist(categoryId) {
		return nil, status.Errorf(codes.NotFound, "Category not found")
	}
	category, err := s.categoryRepo.Get(categoryId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	return category, nil
}

func (s *ProductService) UpdateCategoryService(id string, category *productproto.Category) error {
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Couldn't process the id")
	}
	if err := s.categoryRepo.Update(categoryId, category); err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error")
	}
	return nil
}

func (s *ProductService) ListCategoriesService() ([]*productproto.Category, error) {
	categories, err := s.categoryRepo.List()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	return categories, nil
}
