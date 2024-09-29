package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	listTitle string = "Templates"
)

type listWrapper struct {
	model list.Model
}

type item struct {
	title string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return i.title }

func NewList(items []list.Item, itemDelegate list.ItemDelegate, width, height int) *listWrapper {
	li := list.New(items, itemDelegate, width, height)
	return &listWrapper{li}
}

func NewDefaultList() *listWrapper {
	items := []list.Item{
		item{"go"},
		item{"gorilla"},
		item{"rust"},
		item{"zig"},
		item{"dart"},
		item{"flutter"},
	}

	return NewList(items, list.NewDefaultDelegate(), 30, 10).
		Default()
}

func (l *listWrapper) Default() *listWrapper {
	l.model.Title = listTitle
	return l
}

func (l *listWrapper) Update(msg tea.Msg) (list.Model, tea.Cmd) {
	return l.model.Update(msg)
}
