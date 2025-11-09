package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/shawon-crosen/dashboard-go/server/pkg/server"
)

func main() {
	path, err := filepath.Abs("../config.yaml")
	config, err := os.ReadFile(path)

	if err != nil {
		log.Println(err)
	}

	server.Start(config)
}
