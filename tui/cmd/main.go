package main

import (
	"os"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	. "github.com/zoryia/vex/tui/models"
	. "github.com/zoryia/vex/tui/pages"
	"github.com/zoryia/vex/tui/pages/auth"
	"github.com/zoryia/vex/tui/pages/feeds"
	"github.com/zoryia/vex/tui/pages/preview"
)

type Model struct {
	// entries
	list       list.Model
	queryInput textinput.Model

	err  error
	page VexPage
	tags []string
	keys *ListKeyMap

	Preview preview.Model
	Feeds   feeds.Model
	Auth    auth.Model
}

func queryInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Search Query"
	ti.CharLimit = 156
	ti.Width = 56
	return ti

}

func New() *Model {
	return &Model{
		queryInput: queryInput(),
		page:       LOGIN,
		keys:       NewListKeyMap(),

		Preview: preview.Model{Viewport: viewport.New(0, 0)},
		Feeds:   feeds.New(),
		Auth:    auth.New(),
	}
}

func (m Model) getEverything() tea.Cmd {
	return tea.Batch(getEntries(m.Auth.Jwt)) // getTags, getFeeds)
}

func (m *Model) initList(width int, height int) {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.list.Title = "Posts"
	m.list.SetFilteringEnabled(false)
	m.list.DisableQuitKeybindings()
	m.list.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			m.keys.Query,
			m.keys.PreviewPost,
		}
	}
	m.list.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			m.keys.GoToFeeds,
			m.keys.ReadToggle,
			m.keys.ReadLaterToggle,
			m.keys.BookmarkToggle,
			m.keys.IgnoreToggle,
			m.keys.Query,
			m.keys.PreviewPost,
		}
	}
	var f = Feed{Id: uuid.UUID{}, Tags: []string{"Devops", "Kubernetes"}, Name: "zwindler", Url: "zwindler.blog", FaviconUrl: "zwindler.blog.favicon"}
	m.list.SetItems([]list.Item{
		Entry{Id: uuid.UUID{}, ArticleTitle: "yay", Content: "ouin ouin ouinouin ouin ouinouin ouin ouinouin ouin ouinouin ouin ouinouin ouin ouinouin ouin ouinouin ouin ouinouin ouin ouinouin ouin ouin", Link: "awd", Date: time.Now(), IsRead: false, IsIgnored: false, IsReadLater: false, IsBookmarked: false, Feed: f},
		Entry{Id: uuid.UUID{}, ArticleTitle: "grrrrr", Content: "ouin ouin ouin", Link: "awd", Date: time.Now(), IsRead: false, IsIgnored: false, IsReadLater: false, IsBookmarked: false, Feed: f},
		Entry{Id: uuid.UUID{}, ArticleTitle: "my life is pain", Content: "ouin ouin ouin", Link: "awd", Date: time.Now(), IsRead: false, IsIgnored: false, IsReadLater: false, IsBookmarked: false, Feed: f},
	})
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.Auth.LoginForm.Init(),
		m.Auth.RegisterForm.Init(),
		checkJwt(m.Auth.Jwt),
	)
}

func main() {
	tea.LogToFile("vex.log", "")
	m := New()
	p := tea.NewProgram(m,
		tea.WithAltScreen(), // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(),
	)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
