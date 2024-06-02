package productservice

import (
	"context"
	"shopito/services/api-gw/config"
	"shopito/services/api-gw/protobuf"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service interface {
	DeleteProductService(id string) error
	CreateProductService(product *protobuf.Product) (string, error)
	GetProductService(id string) (*protobuf.Product, error)
	ListProductService() (*protobuf.ListProductsResponse, error)
	Close()
}

type ProductService struct {
	clientGRPC protobuf.ProductServiceClient
	conn       *grpc.ClientConn
}

func New() *ProductService {
	conn, err := grpc.NewClient(config.PRODUCTS_ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	client := protobuf.NewProductServiceClient(conn)
	return &ProductService{
		clientGRPC: client,
		conn:       conn,
	}
}

func (s *ProductService) DeleteProductService(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.clientGRPC.DeleteProduct(ctx, &protobuf.DeleteProductRequest{Id: id})
	return err
}

func (s *ProductService) CreateProductService(product *protobuf.Product) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := s.clientGRPC.CreateProduct(ctx, &protobuf.CreateProductRequest{
		Name:   product.GetName(),
		ImgUrl: product.GetImgUrl(),
		Price:  product.GetPrice(),
		Stock:  product.GetStock(),
	})
	return r.GetId(), err
}

func (s *ProductService) GetProductService(id string) (*protobuf.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := s.clientGRPC.GetProduct(ctx, &protobuf.GetProductRequest{Id: id})
	return r, err
}

func (s *ProductService) ListProductService() (*protobuf.ListProductsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := s.clientGRPC.ListProducts(ctx, &protobuf.ListProductsRequest{})
	return r, err
}

func (s *ProductService) Close() {
	err := s.conn.Close()
	if err != nil {
		logrus.WithError(err).Error("Couldn't close products service grpc connection")
	} else {
		logrus.Info("products service grpc conn closed")
	}
}
