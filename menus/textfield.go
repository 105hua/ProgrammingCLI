package menus

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errorMsg error
)

type Model struct {
	textInput textinput.Model
	err       error
}

func InitialModel() Model {
	textInput := textinput.New()
	textInput.Placeholder = "Ask anything..."
	textInput.Focus()
	textInput.CharLimit = 156
	textInput.Width = 100

	return Model{
		textInput: textInput,
		err:       nil,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case errorMsg:
		m.err = msg
		return m, nil
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return fmt.Sprintf(
		"%s %s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

// GetValue returns the current value of the text input
func (m Model) GetValue() string {
	return m.textInput.Value()
}
