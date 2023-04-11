package main

import (
	"PTOBuilder/config"
	"PTOBuilder/internal/api"
	"PTOBuilder/pkg/logging"
)

func main() {
	// Get logger for to display and save all api behavior
	log := logging.GetLogger()

	// Read all configurations from the yaml config file using the viper module
	config.GetConfigs()

	// Create and initialize api structure with all dependencies
	API := api.NewAPI(&log)
	API.Init()

	// Start our REST API
	API.Start()
}
