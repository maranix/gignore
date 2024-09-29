package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type state struct {
	textinput *textInputWrapper
	list      *listWrapper
	err       error
}

func newState() *state {
	ti := newTextInput().
		Default().
		Focus()

	li := NewDefaultList()

	return &state{
		textinput: ti,
		list:      li,
		err:       nil,
	}
}

func (s *state) Init() tea.Cmd {
	return textinput.Blink
}

func (s *state) View() string {
	return fmt.Sprintf(
		"Select a .gitignore template to fetch\n\n%s\n\n%s\n\n%s",
		s.textinput.model.View(),
		lipgloss.NewStyle().Render(s.list.model.View()),
		"(esc or ctrl-c to quit)",
	) + "\n"
}

func (s *state) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

	var textinputCmd tea.Cmd
	var listCmd tea.Cmd

	s.textinput.model, textinputCmd = s.textinput.Update(msg)
	s.list.model, listCmd = s.list.Update(msg)

	return s, tea.Batch(textinputCmd, listCmd)
}
