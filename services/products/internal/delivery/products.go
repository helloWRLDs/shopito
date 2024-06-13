package delivery

import (
	"context"
	productproto "shopito/pkg/protobuf/products"
)

func (d *Delivery) CreateProduct(ctx context.Context, request *productproto.CreateProductRequest) (*productproto.CreateProductResponse, error) {
	response := productproto.CreateProductResponse{Success: false, Id: -1}
	product := productproto.Product{
		Id:         request.GetId(),
		Name:       request.GetName(),
		ImgUrl:     request.GetImgUrl(),
		Price:      request.GetPrice(),
		Stock:      request.GetStock(),
		CategoryId: request.GetCategoryId(),
	}
	id, err := d.service.CreateProductService(&product)
	if err != nil {
		response.Success = false
		return &response, err
	}
	response.Id = id
	response.Success = true
	return &response, nil
}

func (d *Delivery) DeleteProduct(ctx context.Context, request *productproto.DeleteProductRequest) (*productproto.DeleteProductResponse, error) {
	response := &productproto.DeleteProductResponse{Success: false}
	if err := d.service.DeleteProductService(request.GetId()); err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}

func (d *Delivery) GetProduct(ctx context.Context, request *productproto.GetProductRequest) (*productproto.GetProductResponse, error) {
	product, err := d.service.GetProductService(request.GetId())
	if err != nil {
		return nil, err
	}
	response := productproto.GetProductResponse{
		Product: product,
	}
	return &response, nil
}

func (d *Delivery) ListProducts(ctx context.Context, request *productproto.ListProductsRequest) (*productproto.ListProductsResponse, error) {
	products, err := d.service.ListProductService(request.GetLimit(), request.GetOffset(), request.GetFilter())
	if err != nil {
		return nil, err
	}
	return &productproto.ListProductsResponse{Products: products}, nil
}

func (d *Delivery) UpdateProduct(ctx context.Context, request *productproto.UpdateProductRequest) (*productproto.UpdateProductResponse, error) {
	response := productproto.UpdateProductResponse{
		Success: false,
	}
	if err := d.service.UpdateProductService(request.GetId(), request.GetProduct()); err != nil {
		return &response, err
	}
	response.Success = true
	return &response, nil
}

func (d *Delivery) ListProductsByCategory(ctx context.Context, request *productproto.ListProductsByCategoryRequest) (*productproto.ListProductsByCategoryResponse, error) {
	products, err := d.service.ListProductByCategoryService(request.GetCategoryName(), request.GetLimit(), request.GetOffset())
	if err != nil {
		return nil, err
	}
	return &productproto.ListProductsByCategoryResponse{Products: products}, nil
}
