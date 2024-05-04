package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	huh "github.com/charmbracelet/huh"
)

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
	auth      Auth
	page      VexPage
	query     string
	feeds     []Feed
	entries   []Entry
	tags      []string
}

func New() *Model {
	loginForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Username").
				Key("username"),
			// Validating fields is easy. The form will mark erroneous fields
			// and display error messages accordingly.
			// TODO:
			// Validate(func(str string) error {
			// 	if str == "octopus773" {
			// 		return errors.New("Not you")
			// 	}
			// 	return nil
			// })))
			huh.NewInput().
				Title("Password").
				Key("password").
				Password(true),
		))
	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 56
	return &Model{textInput: ti, auth: Auth{form: loginForm, jwt: new(string)}, page: "LOGIN"}
}

func (m Model) getEverything() tea.Cmd {
	return func() tea.Msg {
		return tea.Batch(getEntries(m.auth.jwt)) // getTags, getFeeds)
	}
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
	return tea.Batch(checkServer, m.auth.form.Init())

}

func (m Model) handleSearchCompletion() (Model, tea.Cmd) {
	var cmd tea.Cmd
	queryWords := strings.Split(m.textInput.Value(), " ")
	if len(queryWords) == 0 {
		return m, cmd
	}
	lastWord := queryWords[len(queryWords)-1]
	if lastWord == "tag:" {
		var feed string
		huh.NewSelect[string]().
			Title("Pick a feed.").
			Options(
				huh.NewOption("United States", "US"),
				huh.NewOption("Germany", "DE"),
				huh.NewOption("Brazil", "BR"),
				huh.NewOption("Canada", "CA"),
			).
			Value(&feed).Run()
		m.textInput.SetValue(m.textInput.Value() + feed)
		m.textInput.CursorEnd()
	}

	if lastWord == "feed:" {
		var feed string
		var feeds = []string{"Devops", "System", "Angular"}
		huh.NewSelect[string]().
			Title("Pick a feed.").
			Options(
				huh.NewOptions(feeds...)...,
			).
			Value(&feed).Run()
		m.textInput.SetValue(m.textInput.Value() + feed)
		m.textInput.CursorEnd()
	}
	return m, cmd
}

func (m *Model) deleteWordBackward() {
	if m.textInput.Position() == 0 || len(m.textInput.Value()) == 0 {
		return
	}

	// TODO: wtf are other echo modes, dont care
	//if m.textInput.EchoMode != textinput.EchoNormal {
	//	m.deleteBeforeCursor()
	//	return
	//}

	// Linter note: it's critical that we acquire the initial cursor position
	// here prior to altering it via SetCursor() below. As such, moving this
	// call into the corresponding if clause does not apply here.
	oldPos := m.textInput.Position() //nolint:ifshort

	m.textInput.SetCursor(oldPos - 1)
	// ECHO character?
	for m.textInput.Value()[m.textInput.Position()] == ' ' {
		if m.textInput.Position() <= 0 {
			break
		}
		// ignore series of whitespace before cursor
		m.textInput.SetCursor(m.textInput.Position() - 1)
	}

	for m.textInput.Position() > 0 {
		if m.textInput.Value()[m.textInput.Position()] != ' ' {
			m.textInput.SetCursor(m.textInput.Position() - 1)
		} else {
			if m.textInput.Position() > 0 {
				// keep the previous space
				m.textInput.SetCursor(m.textInput.Position() + 1)
			}
			break
		}
	}

	if oldPos > len(m.textInput.Value()) {
		m.textInput.SetValue(m.textInput.Value()[:m.textInput.Position()])
	} else {
		m.textInput.SetValue(m.textInput.Value()[:m.textInput.Position()] + m.textInput.Value()[oldPos:])
	}
}

const url = "localhost:3000"

func checkServer() tea.Msg {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := c.Get(url)
	if err != nil {
		return errMsg{err}
	}
	defer res.Body.Close() // nolint:errcheck

	return statusMsg(res.StatusCode)

}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.initList(msg.Width, msg.Height)
	case loginSuccessMsg:
		*m.auth.jwt = msg.string
		m.page = FEEDS

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
		switch {
		case key.Matches(msg, m.textInput.KeyMap.DeleteCharacterBackward):
			words := strings.Split(m.textInput.Value(), " ")
			if len(words) > 0 && (strings.HasPrefix(words[len(words)-1], "tag:") || strings.HasPrefix(words[len(words)-1], "feed:")) {
				m.deleteWordBackward()
			}
		}
		// if writing
	}
	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.auth.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.auth.form = f
		cmds = append(cmds, cmd)
	}

	if m.auth.form.State == huh.StateCompleted {
		// Quit when the form is done.
		username := m.auth.form.GetString("username")
		password := m.auth.form.GetString("password")
		cmds = append(cmds, login(username, password))
	}

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)
	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		_ = msg
		m, cmd = m.handleSearchCompletion()
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func main() {
	tea.LogToFile("yay.log", "")
	m := New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}