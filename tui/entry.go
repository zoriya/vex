package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	huh "github.com/charmbracelet/huh"
)

type Feed struct {
	id         string
	name       string
	url        string
	faviconUrl string
	tags       []string
}

type Entry struct {
	id      string
	title   string
	content string
	link    string
	date    time.Time

	author       *string // author not always specified
	isRead       bool
	isBookmarked bool
	isIgnored    bool
	isReadLater  bool
	feed         Feed
}

func (e Entry) FilterValue() string {
	return e.title
}

func (e Entry) Title() string {
	return e.title
}

func (e Entry) Description() string {
	return fmt.Sprintf("%s", "my desc") // TODO: real description
}

type Model struct {
	list      list.Model
	textInput textinput.Model
	err       error
}

func New() *Model {
	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	return &Model{textInput: ti}
}

func (m *Model) initList(width int, height int) {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.list.Title = "Posts"
	m.list.SetFilteringEnabled(false)
	var f = Feed{id: "1", tags: []string{"Devops", "Kubernetes"}, name: "zwindler", url: "zwindler.blog", faviconUrl: "zwindler.blog.favicon"}
	m.list.SetItems([]list.Item{
		Entry{id: "1", title: "yay", content: "ouin ouin ouin", link: "awd", date: time.Now(), isRead: false, isIgnored: false, isReadLater: false, isBookmarked: false, feed: f},
		Entry{id: "2", title: "grrrrr", content: "ouin ouin ouin", link: "awd", date: time.Now(), isRead: false, isIgnored: false, isReadLater: false, isBookmarked: false, feed: f},
		Entry{id: "3", title: "my life is pain", content: "ouin ouin ouin", link: "awd", date: time.Now(), isRead: false, isIgnored: false, isReadLater: false, isBookmarked: false, feed: f},
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.initList(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	m.textInput, cmd = m.textInput.Update(msg)
	if strings.HasSuffix(m.textInput.Value(), "feed:") {
		var feed string
		huh.NewSelect[string]().
			Title("Pick a country.").
			Options(
				huh.NewOption("United States", "US"),
				huh.NewOption("Germany", "DE"),
				huh.NewOption("Brazil", "BR"),
				huh.NewOption("Canada", "CA"),
			).
			Value(&feed).Run()
		m.textInput.SetValue(m.textInput.Value() + feed)
		// var suggestions = []string{"cncf.io/rss", "zwindler.blog/index.xml"}
		// s := make([]string, len(suggestions))
		// for i := range suggestions {
		// 	s[i] = fmt.Sprintf("%s%s", m.textInput.Value(), suggestions[i])
		// }
		// m.textInput.SetSuggestions(s)
		// m.textInput.ShowSuggestions = true
	}
	return m, cmd
}

func (m Model) View() string {
	return m.textInput.View() // + m.list.View()
}

func main() {
	m := New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
