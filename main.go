package main

import (
	"github.com/pires-guilherme/go-jobs-opportunities/config"
	"github.com/pires-guilherme/go-jobs-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("config init error : %v", err)
	}

	router.Initialize()
}
