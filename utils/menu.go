package utils

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func DisplayTitle() {
	robotColor := color.New(color.FgHiYellow).SprintFunc()
	lineColor := color.New(color.FgYellow).SprintFunc()
	boldText := color.New(color.Bold).SprintFunc()

	fmt.Println(lineColor(strings.Repeat("─", 100)))
	fmt.Println(robotColor("     █"))
	fmt.Println(robotColor("   █████"))
	fmt.Println(robotColor("   █ █ █      ") + boldText("Programming CLI | v0.0.1"))
	fmt.Println(robotColor("   █████"))
	fmt.Println(robotColor("   █ █ █"))
	fmt.Println(lineColor(strings.Repeat("─", 100)))
}
