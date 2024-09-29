package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type textInputWrapper struct {
	model textinput.Model
}

const (
	textinputPlaceholder string = "go"
)

func newTextInput() *textInputWrapper {
	return &textInputWrapper{textinput.New()}

}

func (ti *textInputWrapper) Default() *textInputWrapper {
	ti.model.Placeholder = textinputPlaceholder
	return ti
}

func (ti *textInputWrapper) Focus() *textInputWrapper {
	ti.model.Focus()
	return ti
}

func (ti *textInputWrapper) Update(msg tea.Msg) (textinput.Model, tea.Cmd) {
	return ti.model.Update(msg)
}
