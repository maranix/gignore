package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type state struct {
	textinput textInputWrapper
	err       error
}

func newState() state {
	ti := *newTextInput().
		Default().
		Focus()

	return state{
		textinput: ti,
		err:       nil,
	}
}

func (s state) Init() tea.Cmd {
	return textinput.Blink
}

func (s state) View() string {
	return fmt.Sprintf(
		"Select a .gitignore template to fetch\n\n%s\n\n%s",
		s.textinput.model.View(),
		"(esc or ctrl-c to quit)",
	) + "\n"
}

func (s state) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEscape:
			return s, tea.Quit
		}

	case error:
		s.err = msg
		return s, nil
	}

	s.textinput.model, cmd = s.textinput.Update(msg)
	return s, cmd
}