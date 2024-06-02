package delivery

import (
	"context"
	productproto "shopito/pkg/protobuf/products"
	productservice "shopito/services/products/internal/services"
	"strconv"
)

type Delivery struct {
	productproto.UnimplementedProductServiceServer
	productServ productservice.Service
}

func New(product *productservice.ProductService) *Delivery {
	return &Delivery{
		productServ: product,
	}
}

func (d *Delivery) CreateCategory(ctx context.Context, request *productproto.CreateCategoryRequest) (*productproto.CreateCategoryResponse, error) {
	response := productproto.CreateCategoryResponse{
		Id:      "-1",
		Success: false,
	}
	id, err := d.productServ.CreateCategoryService(&productproto.Category{
		Name: request.GetName(),
	})
	if err != nil {
		return &response, err
	}
	response.Success = true
	response.Id = strconv.Itoa(id)
	return &response, nil
}

func (d *Delivery) CreateProduct(ctx context.Context, request *productproto.CreateProductRequest) (*productproto.CreateProductResponse, error) {
	response := productproto.CreateProductResponse{}
	product := productproto.Product{
		Id:         request.GetId(),
		Name:       request.GetName(),
		ImgUrl:     request.GetImgUrl(),
		Price:      request.GetPrice(),
		Stock:      request.GetStock(),
		CategoryId: request.GetCategoryId(),
	}
	id, err := d.productServ.CreateProductService(&product)
	if err != nil {
		response.Success = false
		return &response, err
	}
	response.Id = strconv.Itoa(id)
	response.Success = true
	return &response, nil
}

func (d *Delivery) DeleteCategory(ctx context.Context, request *productproto.DeleteCategoryRequest) (*productproto.DeleteCategoryResponse, error) {
	response := productproto.DeleteCategoryResponse{
		Success: false,
	}
	if err := d.productServ.DeleteCategoryService(request.GetId()); err != nil {
		return &response, err
	}
	response.Success = true
	return &response, nil
}

func (d *Delivery) DeleteProduct(ctx context.Context, request *productproto.DeleteProductRequest) (*productproto.DeleteProductResponse, error) {
	response := &productproto.DeleteProductResponse{Success: false}
	if err := d.productServ.DeleteProductService(request.GetId()); err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}

func (d *Delivery) GetCategory(ctx context.Context, request *productproto.GetCategoryRequest) (*productproto.GetCategoryResponse, error) {
	category, err := d.productServ.GetCategoryService(request.GetId())
	if err != nil {
		return nil, err
	}
	response := productproto.GetCategoryResponse{
		Category: category,
	}
	return &response, nil
}

func (d *Delivery) GetProduct(ctx context.Context, request *productproto.GetProductRequest) (*productproto.GetProductResponse, error) {
	product, err := d.productServ.GetProductService(request.GetId())
	if err != nil {
		return nil, err
	}
	response := productproto.GetProductResponse{
		Product: product,
	}
	return &response, nil
}

func (d *Delivery) ListCategories(ctx context.Context, request *productproto.ListCategoriesRequest) (*productproto.ListCategoriesResponse, error) {
	categories, err := d.productServ.ListCategoriesService()
	if err != nil {
		return nil, err
	}
	response := productproto.ListCategoriesResponse{
		Categories: categories,
	}
	return &response, nil
}

func (d *Delivery) ListProducts(ctx context.Context, request *productproto.ListProductsRequest) (*productproto.ListProductsResponse, error) {
	return nil, nil
}

func (d *Delivery) ListProductsByCategory(ctx context.Context, request *productproto.ListProductsByCategoryRequest) (*productproto.ListProductsByCategoryResponse, error) {
	return nil, nil
}

func (d *Delivery) UpdateCategory(ctx context.Context, request *productproto.UpdateCategoryRequest) (*productproto.UpdateCategoryResponse, error) {
	response := productproto.UpdateCategoryResponse{
		Success: false,
	}
	err := d.productServ.UpdateCategoryService(request.GetId(), request.GetCategory())
	if err != nil {
		return &response, err
	}
	response.Success = true
	return &response, nil
}

func (d *Delivery) UpdateProduct(ctx context.Context, request *productproto.UpdateProductRequest) (*productproto.UpdateProductResponse, error) {
	response := productproto.UpdateProductResponse{
		Success: false,
	}
	if err := d.productServ.UpdateProductService(request.GetId(), request.GetProduct()); err != nil {
		return &response, err
	}
	response.Success = true
	return &response, nil
}
