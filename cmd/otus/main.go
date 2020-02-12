package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Hello World")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		logrus.Fatal("Port is not set")
	}
}
