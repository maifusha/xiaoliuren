package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"xiaoliuren/internal/config"
	"xiaoliuren/internal/util/file"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	filepath := path.Join(os.Getenv("GOPATH"), config.Conf.Logfile.Runtime)

	logger = logrus.New()
	logger.SetOutput(file.MustOpen(filepath))
	logger.SetFormatter(&logrus.JSONFormatter{})
}

func Printf(format string, args ...interface{}) {
	logger.WithFields(CallerFileFields()).Printf(format, args...)
}

func Println(args ...interface{}) {
	logger.WithFields(CallerFileFields()).Println(args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.WithFields(CallerFileFields()).Fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
	logger.WithFields(CallerFileFields()).Fatalln(args...)
}

func CallerFileFields() logrus.Fields {
	_, f, l, _ := runtime.Caller(2)

	return logrus.Fields{"file": fmt.Sprintf("%s:%d", f, l)}
}
