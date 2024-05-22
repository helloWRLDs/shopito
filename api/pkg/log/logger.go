package log

import (
	"io"
	"log"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var once sync.Once

func Init() {
	once.Do(func() {
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableTimestamp: false,
			DisableColors:    false,
			TimestampFormat:  "2006-01-02 15:04:05",
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "@time",
				logrus.FieldKeyFunc:  "from",
				logrus.FieldKeyLevel: "lvl",
			},
		})
		logFile, err := os.OpenFile("api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		logrus.SetOutput(io.MultiWriter(os.Stdout, logFile))
	})
}
