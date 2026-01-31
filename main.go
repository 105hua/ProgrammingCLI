package main

import (
	"ProgrammingCLI/menus"
	"ProgrammingCLI/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Load configuration from config.json
	config := utils.LoadConfig()

	// Initialize the application with the loaded config
	// The config now contains ApiKey and Model for API calls
	_ = config

	// Display the title message.
	utils.DisplayTitle()

	// Make text box with tea.
	p := tea.NewProgram(menus.InitialModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
