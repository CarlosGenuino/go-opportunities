package main

import (
	"github.com/CarlosGenuino/go-opportunities/config"
	"github.com/CarlosGenuino/go-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initilizer Config
	err := config.Init()

	if err != nil {

		logger.Errorf("config initilization error: %v", err)
		return
	}

	//Initilize Router
	router.Initialize()
}
