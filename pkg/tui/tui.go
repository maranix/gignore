package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Tui struct {
	app *tea.Program
}

type model struct {
	textinput textinput.Model
	err       error
}

func initModel() model {
	ti := textinput.New()
	ti.Placeholder = "go"
	ti.Focus()

	return model{
		textinput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) View() string {
	return fmt.Sprintf(
		"Select a .gitignore template to fetch\n\n%s\n\n%s",
		m.textinput.View(),
		"(esc or ctrl-c to quit)",
	) + "\n"
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEscape:
			return m, tea.Quit
		}

	case error:
		m.err = msg
		return m, nil
	}

	m.textinput, cmd = m.textinput.Update(msg)
	return m, cmd
}

func NewApp() *Tui {
	p := tea.NewProgram(initModel())

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
