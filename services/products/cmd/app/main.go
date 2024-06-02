package main

import (
	"net"
	"shopito/pkg/datastore/postgres"
	"shopito/pkg/log"
	"shopito/services/products/config"
	"shopito/services/products/internal/delivery"
	categoryrepository "shopito/services/products/internal/repository/category"
	productrepository "shopito/services/products/internal/repository/product"
	productservice "shopito/services/products/internal/services"
	"shopito/pkg/protobuf/products"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	log.Init("products")
	db, err := postgres.Open(config.DB.HOST, config.DB.PORT, config.DB.USER, config.DB.PASSWORD, config.DB.NAME)
	if err != nil {
		logrus.WithError(err).Panic("db conn failure")
	}

	lis, err := net.Listen("tcp", config.ADDR)
	if err != nil {
		logrus.WithError(err).Error("failed to listen")
	}
	srv := grpc.NewServer()

	service := productservice.New(
		productrepository.New(db),
		categoryrepository.New(db),
	)
	delivery := delivery.New(service)

	productproto.RegisterProductServiceServer(srv, delivery)

	logrus.WithField("addr", config.ADDR).Info("server started")
	if err := srv.Serve(lis); err != nil {
		logrus.WithError(err).Error("filed to serve")
	}
}
