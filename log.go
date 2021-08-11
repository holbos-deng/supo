package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.WarnLevel)
	Log = logrus.StandardLogger()
}
