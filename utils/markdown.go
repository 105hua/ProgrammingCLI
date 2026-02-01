package utils

import (
	"os"

	markdown "github.com/MichaelMure/go-term-markdown"
	"golang.org/x/term"
)

func RenderMarkdown(content string) (string, error) {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80 // Default width fallback
	}

	// Render the markdown
	// source, width, leftPadding
	result := markdown.Render(content, width, 0)

	return string(result), nil
}
