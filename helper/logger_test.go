package helper

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestLog(t *testing.T) {
	logger := logrus.New()
	currentTime := time.Now()

	folder := fmt.Sprintf("log/%d/%d", currentTime.Year(), currentTime.Month())
	filename := fmt.Sprintf("application-%d.log", currentTime.Day())

	_, err := pathExists(folder)
	if err != nil {
		fmt.Println(err)
		os.MkdirAll(folder, 0666)
	}

	filename = fmt.Sprintf("%s/%s", folder, filename)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	logger.SetOutput(file)

	logger.Info("Logg info")
}
