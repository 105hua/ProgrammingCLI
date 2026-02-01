package utils

import (
	"ProgrammingCLI/menus"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
)

func DisplayTitle() {
	robotColor := color.New(color.FgHiYellow).SprintFunc()
	lineColor := color.New(color.FgYellow).SprintFunc()
	boldText := color.New(color.Bold).SprintFunc()

	fmt.Println(lineColor(strings.Repeat("─", 100)))
	fmt.Println(robotColor("     ▗▄▓▓▄▖"))
	fmt.Println(robotColor("    ▗▓▀  ▀▓▖"))
	fmt.Println(robotColor("    ▐▌ ██ ▐▌      ") + boldText("Programming CLI | v0.0.1"))
	fmt.Println(robotColor("    ▝▓▄  ▄▓▘"))
	fmt.Println(robotColor("     ▝▀▓▓▀▘"))
	fmt.Println(lineColor(strings.Repeat("─", 100)))
}

func NewConversationScreen() {
	// Clear screen.
	fmt.Print("\033[H\033[2J")
	// Display title screen.
	DisplayTitle()

	// Load config once
	config := LoadConfig()

	// Open first text box.
	p := tea.NewProgram(menus.InitialModel())
	finalModel, err := p.Run()
	if err != nil {
		panic(err)
	}

	// Get the user input from the final model state
	if m, ok := finalModel.(menus.Model); ok {
		userInput := m.GetValue()
		if userInput != "" {
			// Send to OpenAI API and get response
			aiColor := color.New(color.FgGreen, color.Bold).SprintFunc()
			fmt.Printf("%s ", aiColor("\nAI:\n"))

			// Get response from OpenRouter API and render it.
			response := GetCompletion(userInput, nil, config.ApiKey)
			renderer, err := CreateRenderer()
			if err != nil {
				panic(err)
			}
			rendered, err := renderer.Render(response.Content)
			if err != nil {
				panic(err)
			}

			// Print rendered content.
			fmt.Printf("%s\n\n", rendered)
		}
	}
}
