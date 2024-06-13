package service

import (
	"context"
	productproto "shopito/pkg/protobuf/products"
	"shopito/services/products/internal/repository"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ProductService) InsertCategoryService(name string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	exist, err := s.repo.IsCategoryExistByName(ctx, name)

	if err != nil {
		logrus.WithError(err).Error("Internal Server Error")
		return -1, status.Errorf(codes.Internal, "Internal Server Error")
	} else if exist {
		return -1, status.Errorf(codes.AlreadyExists, "Category with such name already exists")
	}

	id, err := s.repo.CreateCategory(ctx, name)
	if err != nil {
		logrus.WithError(err).Error("Internal Server Error")
		return -1, status.Errorf(codes.Internal, "Internal Error")
	}
	return id, nil
}

func (s *ProductService) CreateCategoryService(category *productproto.Category) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	exist, err := s.repo.IsCategoryExistByName(ctx, category.GetName())
	if err != nil {
		return -1, status.Errorf(codes.Internal, "Internal Server Error")
	} else if exist {
		return -1, status.Errorf(codes.NotFound, "Product already exist")
	}
	id, err := s.repo.CreateCategory(ctx, category.GetName())
	if err != nil {
		logrus.WithError(err).Error("Internal Server Error")
		return -1, status.Errorf(codes.Internal, "Internal Server error")
	}
	return id, nil
}

func (s *ProductService) DeleteCategoryService(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	if id == 0 {
		return status.Errorf(codes.InvalidArgument, "Couldn't process the id")
	}
	exist, err := s.repo.IsCategoryExistByID(ctx, id)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error")
	} else if !exist {
		return status.Errorf(codes.NotFound, "Category not found")
	}

	if err := s.repo.DeleteCategory(ctx, id); err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error")
	}
	return nil
}

func (s *ProductService) GetCategoryService(id int64) (*productproto.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Couldn't process the id")
	}
	exist, err := s.repo.IsCategoryExistByID(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	} else if !exist {
		return nil, status.Errorf(codes.NotFound, "Category not found")
	}
	category, err := s.repo.GetCategoryById(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	protoCategory := productproto.Category{
		Id:   id,
		Name: category.Name,
	}
	return &protoCategory, nil
}

func (s *ProductService) UpdateCategoryService(id int64, category *productproto.Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	if id == 0 {
		return status.Errorf(codes.InvalidArgument, "Couldn't process the id")
	}
	exist, err := s.repo.IsCategoryExistByID(ctx, id)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error")
	} else if !exist {
		return status.Errorf(codes.NotFound, "Category not found")
	}
	params := repository.UpdateCategoryParams{
		Name: category.GetName(),
		ID:   id,
	}
	if err := s.repo.UpdateCategory(ctx, params); err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error")
	}
	return nil
}

func (s *ProductService) ListCategoriesService() ([]*productproto.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	categories, err := s.repo.ListCategories(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	var protoCategories = make([]*productproto.Category, len(categories))
	for i, category := range categories {
		c := productproto.Category{
			Id:   category.ID,
			Name: category.Name,
		}
		protoCategories[i] = &c
	}

	return protoCategories, nil
}
