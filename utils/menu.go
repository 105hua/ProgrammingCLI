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
	// Open first text box.
	p := tea.NewProgram(menus.InitialModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
