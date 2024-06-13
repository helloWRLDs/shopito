package service

import (
	"context"
	"database/sql"
	productproto "shopito/pkg/protobuf/products"
	"shopito/services/products/internal/repository"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ProductService) CreateProductService(product *productproto.Product) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	categoryId := sql.NullInt64{Int64: product.CategoryId, Valid: true}

	id, err := s.repo.CreateProduct(ctx, repository.CreateProductParams{
		Name:       product.GetName(),
		Price:      product.GetPrice(),
		Stock:      product.GetStock(),
		CategoryID: categoryId,
		ImgUrl:     sql.NullString{String: product.GetImgUrl(), Valid: true},
	})
	if err != nil {
		logrus.WithError(err).Error("Internal Server Error")
		return -1, status.Errorf(codes.Internal, "Internal Server error")
	}
	return id, nil
}

func (s *ProductService) DeleteProductService(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	if id == 0 {
		return status.Errorf(codes.InvalidArgument, "Couldn't process the id")
	}
	exist, err := s.repo.IsProductExistByID(ctx, id)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error")
	} else if !exist {
		return status.Errorf(codes.NotFound, "Product not found")
	}
	if err := s.repo.DeleteProduct(ctx, id); err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error")
	}
	return nil
}

func (s *ProductService) GetProductService(id int64) (*productproto.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Couldn't process the id")
	}
	exist, err := s.repo.IsProductExistByID(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	} else if !exist {
		return nil, status.Errorf(codes.NotFound, "Product not found")
	}
	product, err := s.repo.GetProductById(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	return &productproto.Product{
		Id:         product.ID,
		Name:       product.Name,
		ImgUrl:     product.ImgUrl.String,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryId: product.CategoryID.Int64,
	}, nil
}

func (s *ProductService) UpdateProductService(id int64, product *productproto.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	if id == 0 {
		return status.Errorf(codes.InvalidArgument, "Couldn't process the id")
	}
	exist, err := s.repo.IsProductExistByID(ctx, id)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error")
	} else if !exist {
		return status.Errorf(codes.NotFound, "Product not found")
	}
	params := repository.UpdateProductParams{
		Name:       product.GetName(),
		ImgUrl:     sql.NullString{String: product.GetImgUrl(), Valid: true},
		Price:      product.GetPrice(),
		Stock:      product.GetStock(),
		CategoryID: sql.NullInt64{Int64: product.GetCategoryId(), Valid: true},
		ID:         id,
	}
	if err := s.repo.UpdateProduct(ctx, params); err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error")
	}
	return nil
}

func (s *ProductService) ListProductService(limit, offset int32, filter string) ([]*productproto.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	params := repository.ListProductsParams{
		Name:    "%" + filter + "%",
		Column2: "id",
		Limit:   limit,
		Offset:  offset,
	}
	result, err := s.repo.ListProducts(ctx, params)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	products := make([]*productproto.Product, len(result))
	for i, p := range result {
		product := productproto.Product{
			Id:         p.ID,
			Name:       p.Name,
			ImgUrl:     p.ImgUrl.String,
			Price:      p.Price,
			Stock:      p.Stock,
			CategoryId: p.CategoryID.Int64,
		}

		products[i] = &product
	}
	return products, nil
}

func (s *ProductService) ListProductByCategoryService(category string, limit, offset int32) ([]*productproto.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	params := repository.ListProductsByCategoryNameParams{
		Name:   "%" + category + "%",
		Offset: offset,
		Limit:  limit,
	}
	result, err := s.repo.ListProductsByCategoryName(ctx, params)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	products := make([]*productproto.Product, len(result))
	for i, p := range result {
		product := productproto.Product{
			Id:         p.ID,
			Name:       p.Name,
			ImgUrl:     p.ImgUrl.String,
			Price:      p.Price,
			Stock:      p.Stock,
			CategoryId: p.CategoryID.Int64,
		}

		products[i] = &product
	}
	return products, nil
}
