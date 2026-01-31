package utils

import (
	"fmt"
	"strings"
)

func DisplayTitle() {
	fmt.Println(strings.Repeat("─", 100))
	// fmt.Println("")
	fmt.Println("     █")
	fmt.Println("   █████")
	fmt.Println("   █ █ █      Programming CLI | v0.0.1")
	fmt.Println("   █████")
	fmt.Println("   █ █ █")
	// fmt.Println("")
	fmt.Println(strings.Repeat("─", 100))
}
