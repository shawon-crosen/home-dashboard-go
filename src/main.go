package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/shawon-crosen/dashboard-go/pkg/server"
)

func main() {
	path, err := filepath.Abs("src/config.yaml")
	config, err := ioutil.ReadFile(path)

	if err != nil {
		log.Println(err)
	}

	server.Start(config)
}
