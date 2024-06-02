package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

var once sync.Once

func Init(fileName string) {
	logPath := fmt.Sprintf("./logs/%v.log", fileName)
	err := os.MkdirAll("./logs", 0755)
	if err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}
	once.Do(func() {
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableTimestamp:       false,
			DisableColors:          false,
			TimestampFormat:        "2006-01-02 15:04:05",
			FullTimestamp:          true,
			DisableLevelTruncation: true,
			CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
				return "", fmt.Sprintf("%v:%v", formatFilePath(f.File), f.Line)
			},
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "time",
				logrus.FieldKeyFunc:  "from",
				logrus.FieldKeyLevel: "lvl",
			},
		})
		logrus.SetReportCaller(true)
		logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		logrus.SetOutput(io.MultiWriter(os.Stdout, logFile))
	})
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
