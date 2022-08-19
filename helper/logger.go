package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func Log(level logrus.Level, message ...interface{}) {
	logger := logrus.New()
	logrus.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile(generateLogFile(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Log(level, message...)
}

func generateLogFile() (filename string) {
	currentTime := time.Now()

	folder := fmt.Sprintf("log/%d/%d", currentTime.Year(), currentTime.Month())
	filename = fmt.Sprintf("application-%d.log", currentTime.Day())

	_, err := pathExists(folder)
	if err != nil {
		os.MkdirAll(folder, 0666)
	}

	filename = fmt.Sprintf("%s/%s", folder, filename)
	return
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}
