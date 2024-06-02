package main

import (
	"net"
	"shopito/pkg/log"
	notifierproto "shopito/pkg/protobuf/notifier"
	"shopito/services/notifier/config"
	"shopito/services/notifier/internal/adapter"
	"shopito/services/notifier/internal/delivery"
	"shopito/services/notifier/internal/service"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	log.Init("notifier")

	lis, err := net.Listen("tcp", config.ADDR)
	if err != nil {
		logrus.WithError(err).Error("failed to listen")
	}
	srv := grpc.NewServer()

	adapter := adapter.New(config.EMAIL.USERNAME, config.EMAIL.PASSWORD, config.EMAIL.HOST, config.EMAIL.PORT, config.EMAIL.FROM)
	service := service.New(adapter)
	delivery := delivery.New(service)

	notifierproto.RegisterNotifierServiceServer(srv, delivery)

	logrus.WithField("addr", config.ADDR).Info("server started")
	if err := srv.Serve(lis); err != nil {
		logrus.WithError(err).Error("filed to serve")
	}
}
