package main

import (
	"ProgrammingCLI/utils"
)

func main() {
	// Load configuration from config.json
	config := utils.LoadConfig() // Load configuration from file.
	_ = config                   // Initialize application with loaded config.

	// Start a new conversation immediately.
	utils.NewConversationScreen()
}
