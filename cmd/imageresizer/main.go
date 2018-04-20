package main

import (
	"log"
	"os"

	"github.com/schepelin/imageresizer/pkg/http"
)

func main() {
	const address string = "0.0.0.0:8081"

	logger := log.New(os.Stdout, "", log.LstdFlags)

	startServer(address, logger)
}

func startServer(addr string, logger *log.Logger) {
	server := http.NewServer(addr, logger)
	server.Start()
}
