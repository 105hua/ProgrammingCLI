package main

import (
	"ProgrammingCLI/utils"
)

func main() {
	utils.DisplayTitle()

	// Load configuration from config.json
	config := utils.LoadConfig()

	// Initialize the application with the loaded config
	// The config now contains ApiKey and Model for API calls
	_ = config
}
