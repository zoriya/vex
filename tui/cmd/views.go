package main

import (
	"github.com/charmbracelet/lipgloss"
	. "github.com/zoryia/vex/tui/pages"
)

func (m Model) AuthView() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.Auth.LoginForm.View(),
		m.Auth.RegisterForm.View(),
	)
}
func (m Model) EntriesView() string {
	return lipgloss.JoinVertical(lipgloss.Left, m.queryInput.View(), m.list.View())
}
func (m Model) FeedsView() string {
	return m.Feeds.View()
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
