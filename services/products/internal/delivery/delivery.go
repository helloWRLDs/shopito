package delivery

import (
	"context"
	productservice "shopito/services/products/internal/services"
	"shopito/services/products/protobuf"
	"strconv"
)

type Delivery struct {
	protobuf.UnimplementedProductServiceServer
	productServ productservice.Service
}

func New(product *productservice.ProductService) *Delivery {
	return &Delivery{
		productServ: product,
	}
}

func (d *Delivery) CreateCategory(ctx context.Context, request *protobuf.CreateCategoryRequest) (*protobuf.CreateCategoryResponse, error) {
	response := protobuf.CreateCategoryResponse{
		Id:      "-1",
		Success: false,
	}
	id, err := d.productServ.CreateCategoryService(&protobuf.Category{
		Name: request.GetName(),
	})
	if err != nil {
		return &response, err
	}
	response.Success = true
	response.Id = strconv.Itoa(id)
	return &response, nil
}

func (d *Delivery) CreateProduct(ctx context.Context, request *protobuf.CreateProductRequest) (*protobuf.CreateProductResponse, error) {
	response := protobuf.CreateProductResponse{}
	product := protobuf.Product{
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

func (d *Delivery) DeleteCategory(ctx context.Context, request *protobuf.DeleteCategoryRequest) (*protobuf.DeleteCategoryResponse, error) {
	response := protobuf.DeleteCategoryResponse{
		Success: false,
	}
	if err := d.productServ.DeleteCategoryService(request.GetId()); err != nil {
		return &response, err
	}
	response.Success = true
	return &response, nil
}

func (d *Delivery) DeleteProduct(ctx context.Context, request *protobuf.DeleteProductRequest) (*protobuf.DeleteProductResponse, error) {
	response := &protobuf.DeleteProductResponse{Success: false}
	if err := d.productServ.DeleteProductService(request.GetId()); err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}

func (d *Delivery) GetCategory(ctx context.Context, request *protobuf.GetCategoryRequest) (*protobuf.GetCategoryResponse, error) {
	category, err := d.productServ.GetCategoryService(request.GetId())
	if err != nil {
		return nil, err
	}
	response := protobuf.GetCategoryResponse{
		Category: category,
	}
	return &response, nil
}

func (d *Delivery) GetProduct(ctx context.Context, request *protobuf.GetProductRequest) (*protobuf.GetProductResponse, error) {
	product, err := d.productServ.GetProductService(request.GetId())
	if err != nil {
		return nil, err
	}
	response := protobuf.GetProductResponse{
		Product: product,
	}
	return &response, nil
}

func (d *Delivery) ListCategories(ctx context.Context, request *protobuf.ListCategoriesRequest) (*protobuf.ListCategoriesResponse, error) {
	categories, err := d.productServ.ListCategoriesService()
	if err != nil {
		return nil, err
	}
	response := protobuf.ListCategoriesResponse{
		Categories: categories,
	}
	return &response, nil
}

func (d *Delivery) ListProducts(ctx context.Context, request *protobuf.ListProductsRequest) (*protobuf.ListProductsResponse, error) {
	return nil, nil
}

func (d *Delivery) ListProductsByCategory(ctx context.Context, request *protobuf.ListProductsByCategoryRequest) (*protobuf.ListProductsByCategoryResponse, error) {
	return nil, nil
}

func (d *Delivery) UpdateCategory(ctx context.Context, request *protobuf.UpdateCategoryRequest) (*protobuf.UpdateCategoryResponse, error) {
	response := protobuf.UpdateCategoryResponse{
		Success: false,
	}
	err := d.productServ.UpdateCategoryService(request.GetId(), request.GetCategory())
	if err != nil {
		return &response, err
	}
	response.Success = true
	return &response, nil
}

func (d *Delivery) UpdateProduct(ctx context.Context, request *protobuf.UpdateProductRequest) (*protobuf.UpdateProductResponse, error) {
	response := protobuf.UpdateProductResponse{
		Success: false,
	}
	if err := d.productServ.UpdateProductService(request.GetId(), request.GetProduct()); err != nil {
		return &response, err
	}
	response.Success = true
	return &response, nil
}
