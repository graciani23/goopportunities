package main

import (
	"github.com/graciani23/goopportunities/config"
	"github.com/graciani23/goopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errf("Error initializing configs: %v", err)
		return
	}
	// initialize the gin router
	router.Init()
}
