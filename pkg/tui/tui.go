package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Tui struct {
	app *tea.Program
}

func NewApp() *Tui {
	p := tea.NewProgram(newState())

	return &Tui{
		app: p,
	}
}

func (t *Tui) Run() error {
	if _, err := t.app.Run(); err != nil {
		return err
	}

	return nil
}
