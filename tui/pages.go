package main

const (
	LOGIN   = "LOGIN"
	ENTRIES = "ENTRIES"
	FEEDS   = "FEEDS"
	TAGS    = "TAGS"
)

type VexPage string

func (m Model) LoginView() string {
	return m.auth.form.View()
}
func (m Model) EntriesView() string {
	return ""
}
func (m Model) FeedsView() string {
	return "feeds"
}
func (m Model) TagsView() string {
	return ""
}
func (m Model) View() string {
	switch m.page {
	case LOGIN:
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
