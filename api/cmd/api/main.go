package main

import (
	"fmt"
	"net/http"
	"os"
	"shopito/api/config"
	v1 "shopito/api/internal/delivery/http"
	"shopito/api/pkg/datastore/postgres"
	"shopito/api/pkg/log"
	"time"

	logrus "github.com/sirupsen/logrus"
)

func main() {
	log.Init()
	db, err := postgres.Open(config.DB.USER, config.DB.PASSWORD, config.DB.NAME)
	if err != nil {
		logrus.WithError(err).Error("db conn failure")
		os.Exit(1)
	}
	logrus.Info("db conn established")

	router := v1.Router(db)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.ADDR),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logrus.WithField("addr", config.ADDR).Info("server started")
	if err := srv.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}
