package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"lora-ingress-http-go/cmd/main/handlers"
	"net/http"
	"os"
)

var port = ":8080"

func init() {
	if os.Getenv("PORT") != "" {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Infof("listening on port %s", port)
}

//main entrypoint
func main() {
	handlers.Initialize()
	logrus.Fatal(http.ListenAndServe(port, nil))
}
