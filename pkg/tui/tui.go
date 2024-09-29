package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Tui struct {
	app   *tea.Program
	state *state
}

func NewApp() *Tui {
	s := newState()
	p := tea.NewProgram(s)

	return &Tui{
		app:   p,
		state: s,
	}
}

func (t *Tui) Run() error {
	if _, err := t.app.Run(); err != nil {
		return err
	}

	return nil
}
