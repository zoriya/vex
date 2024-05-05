package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	. "github.com/zoryia/vex/tui/pages"
)

func (m Model) AuthView() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.auth.LoginForm.View(),
		m.auth.RegisterForm.View(),
	)
}
func (m Model) EntriesView() string {
	return lipgloss.JoinVertical(lipgloss.Left, m.textInput.View(), m.list.View())
}
func (m Model) FeedsView() string {
	return fmt.Sprintf("%s ", *m.auth.Jwt)
}
func (m Model) TagsView() string {
	return ""
}
func (m Model) View() string {
	switch m.page {
	case LOGIN:
		return m.AuthView()
	case REGISTER:
		return m.AuthView()
	case ENTRIES:
		return m.EntriesView()
	case FEEDS:
		return m.FeedsView()
	case TAGS:
		return m.TagsView()
	case PREVIEW:
		return m.Preview.View()
	}
	return "Really unexpected state, get help"
}
