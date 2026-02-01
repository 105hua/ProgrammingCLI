package utils

import (
	"ProgrammingCLI/menus"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
)

// Commands - user input commands
const (
	CmdExit = "/exit"
	CmdNew  = "/new"
)

// ANSI escape sequences for terminal control
const (
	AnsiClearScreen = "\033[H\033[2J"
)

// Display constants
const (
	TitleLineWidth = 100
	AppVersion     = "v0.0.1"
)

// Pre-configured color functions for consistent styling
var (
	robotColor = color.New(color.FgHiYellow).SprintFunc()
	lineColor  = color.New(color.FgYellow).SprintFunc()
	boldText   = color.New(color.Bold).SprintFunc()
	aiColor    = color.New(color.FgGreen, color.Bold).SprintFunc()
	errorColor = color.New(color.FgRed, color.Bold).SprintFunc()
)

// ClearScreen clears the terminal screen.
func ClearScreen() {
	fmt.Print(AnsiClearScreen)
}

// DisplayTitle prints the application header with ASCII art.
func DisplayTitle() {
	fmt.Println(lineColor(strings.Repeat("─", TitleLineWidth)))
	fmt.Println(robotColor("     ▗▄▓▓▄▖"))
	fmt.Println(robotColor("    ▗▓▀  ▀▓▖"))
	fmt.Println(robotColor("    ▐▌ ██ ▐▌      ") + boldText("Programming CLI | "+AppVersion))
	fmt.Println(robotColor("    ▝▓▄  ▄▓▘"))
	fmt.Println(robotColor("     ▝▀▓▓▀▘"))
	fmt.Println(lineColor(strings.Repeat("─", TitleLineWidth)))
}

// CommandResult represents the outcome of processing a user command.
type CommandResult int

const (
	CommandContinue CommandResult = iota // Continue the conversation loop
	CommandExit                          // Exit the conversation
	CommandNew                           // Start a new conversation
	CommandMessage                       // Process as a message to AI
)

// handleCommand processes user input and determines the appropriate action.
func handleCommand(input string) CommandResult {
	switch input {
	case "", CmdExit:
		return CommandExit
	case CmdNew:
		return CommandNew
	default:
		return CommandMessage
	}
}

// displayAIResponse sends a message to the AI and displays the response.
// Returns an error if the API call or rendering fails.
func displayAIResponse(userInput string, config Config) error {
	fmt.Printf("%s ", aiColor("\nAI:\n"))

	response := GetCompletion(userInput, nil, config.ApiKey)
	rendered, err := RenderMarkdown(response.Content)
	if err != nil {
		return fmt.Errorf("failed to render markdown: %w", err)
	}

	fmt.Printf("%s\n\n", rendered)
	return nil
}

// displayError shows a user-friendly error message.
func displayError(message string) {
	fmt.Printf("\n%s %s\n\n", errorColor("Error:"), message)
}

// runInputLoop handles the main conversation input loop.
// Returns true if a new conversation should be started, false to exit.
func runInputLoop(config Config) bool {
	for {
		p := tea.NewProgram(menus.InitialModel())
		finalModel, err := p.Run()
		if err != nil {
			displayError(fmt.Sprintf("Input error: %v", err))
			return false
		}

		m, ok := finalModel.(menus.Model)
		if !ok {
			displayError("Failed to get user input")
			return false
		}

		userInput := m.GetValue()

		switch handleCommand(userInput) {
		case CommandExit:
			return false
		case CommandNew:
			return true
		case CommandMessage:
			if err := displayAIResponse(userInput, config); err != nil {
				displayError(err.Error())
				// Continue the loop instead of crashing
			}
		}
	}
}

// NewConversation starts and manages a new CLI conversation session.
func NewConversation() {
	config := LoadConfig()

	for {
		// Reset conversation state for new conversation
		ResetConversation()

		ClearScreen()
		DisplayTitle()

		if !runInputLoop(config) {
			return
		}
	}
}
