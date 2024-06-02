package main

import (
	"net"
	"shopito/pkg/datastore/postgres"
	"shopito/pkg/log"
	"shopito/services/users/config"
	"shopito/services/users/internal/delivery"
	"shopito/services/users/internal/repository"
	"shopito/services/users/internal/service"

	"shopito/services/users/protobuf"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	log.Init("users")
	db, err := postgres.Open(config.DB.HOST, config.DB.PORT, config.DB.USER, config.DB.PASSWORD, config.DB.NAME)
	if err != nil {
		logrus.WithError(err).Panic("db conn failure")
	}

	lis, err := net.Listen("tcp", config.ADDR)
	if err != nil {
		logrus.WithError(err).Error("failed to listen")
	}
	srv := grpc.NewServer()

	repo := repository.New(db)
	service := service.New(repo)
	delivery := delivery.New(service)

	protobuf.RegisterUserServiceServer(srv, delivery)

	logrus.WithField("addr", config.ADDR).Info("server started")
	if err := srv.Serve(lis); err != nil {
		logrus.WithError(err).Error("filed to serve")
	}
}
