package main

import (
	"github.com/zfabricio/gopportunities/config"
	"github.com/zfabricio/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	// Initializer Configs
	err := config.Init()
	if err != nil {
		logger.ErrF("Config initization error: %v", err)
		return
	}
	// Initialize Router
	router.Initialize()
}
