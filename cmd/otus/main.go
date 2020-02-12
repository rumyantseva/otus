package main

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Hello World")

	port := os.Getenv("PORT")
	if port == "" {
		logrus.Fatal("Port is not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":"+port, nil)

}
