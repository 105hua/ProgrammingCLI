package utils

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/styles"
)

func CreateRenderer() (*glamour.TermRenderer, error) {
	style := styles.DarkStyleConfig
	bg := "#3a3a3a"
	style.CodeBlock.BackgroundColor = &bg

	renderer, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithStyles(style),
	)
	if err != nil {
		return nil, err
	}

	return renderer, nil
}
