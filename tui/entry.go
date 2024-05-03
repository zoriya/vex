package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
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
	list list.Model
	err  error
}

func New() *Model {
	return &Model{}
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
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.list.View()
}

func main() {
	m := New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
