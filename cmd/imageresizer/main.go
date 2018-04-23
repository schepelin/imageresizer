package main

import (
	"log"
	"os"

	"github.com/schepelin/imageresizer/pkg/http"
)

func main() {
	const address string = "0.0.0.0:8081"

	const dbConnect string = "postgres://schepelin:topsecret@localhost/image_resizer?sslmode=disable"

	logger := log.New(os.Stdout, "", log.LstdFlags)

	startServer(address, logger)
}

func startServer(addr string, logger *log.Logger) {
	server := http.NewServer(addr, logger)
	server.Start()
}
