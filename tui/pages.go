package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	LOGIN    = "LOGIN"
	REGISTER = "REGISTER"
	ENTRIES  = "ENTRIES"
	FEEDS    = "FEEDS"
	TAGS     = "TAGS"
)

type VexPage string

func (m Model) LoginView() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.auth.loginForm.View(),
		m.auth.registerForm.View(),
	)
}
func (m Model) EntriesView() string {
	return ""
}
func (m Model) FeedsView() string {
	return fmt.Sprintf("%s ", *m.auth.jwt)
}
func (m Model) TagsView() string {
	return ""
}
func (m Model) View() string {
	switch m.page {
	case LOGIN:
		return m.LoginView()
	case REGISTER:
		return m.LoginView()
	case ENTRIES:
		return m.EntriesView()
	case FEEDS:
		return m.FeedsView()
	case TAGS:
		return m.TagsView()
	}
	return m.textInput.View() + m.list.View()
}
