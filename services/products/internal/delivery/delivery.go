package delivery

import (
	productproto "shopito/pkg/protobuf/products"
	productservice "shopito/services/products/internal/services"
)

type Delivery struct {
	productproto.UnimplementedProductServiceServer
	service productservice.Service
}

func New(product *productservice.ProductService) *Delivery {
	return &Delivery{
		service: product,
	}
}
