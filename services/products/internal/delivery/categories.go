package delivery

import (
	"context"
	productproto "shopito/pkg/protobuf/products"
)

func (d *Delivery) CreateCategory(ctx context.Context, request *productproto.CreateCategoryRequest) (*productproto.CreateCategoryResponse, error) {
	response := productproto.CreateCategoryResponse{Id: -1, Success: false}
	id, err := d.service.CreateCategoryService(&productproto.Category{
		Name: request.GetName(),
	})
	if err != nil {
		return &response, err
	}
	response.Success = true
	response.Id = id
	return &response, nil
}

func (d *Delivery) DeleteCategory(ctx context.Context, request *productproto.DeleteCategoryRequest) (*productproto.DeleteCategoryResponse, error) {
	response := productproto.DeleteCategoryResponse{Success: false}
	if err := d.service.DeleteCategoryService(request.GetId()); err != nil {
		return &response, err
	}
	response.Success = true
	return &response, nil
}

func (d *Delivery) GetCategory(ctx context.Context, request *productproto.GetCategoryRequest) (*productproto.GetCategoryResponse, error) {
	category, err := d.service.GetCategoryService(request.GetId())
	if err != nil {
		return nil, err
	}
	response := productproto.GetCategoryResponse{
		Category: category,
	}
	return &response, nil
}

func (d *Delivery) ListCategories(ctx context.Context, request *productproto.ListCategoriesRequest) (*productproto.ListCategoriesResponse, error) {
	categories, err := d.service.ListCategoriesService()
	if err != nil {
		return nil, err
	}
	response := productproto.ListCategoriesResponse{
		Categories: categories,
	}
	return &response, nil
}

func (d *Delivery) UpdateCategory(ctx context.Context, request *productproto.UpdateCategoryRequest) (*productproto.UpdateCategoryResponse, error) {
	response := productproto.UpdateCategoryResponse{
		Success: false,
	}
	err := d.service.UpdateCategoryService(request.GetId(), request.GetCategory())
	if err != nil {
		return &response, err
	}
	response.Success = true
	return &response, nil
}
