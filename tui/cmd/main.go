package main

import (
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	huh "github.com/charmbracelet/huh"
	"github.com/google/uuid"
	. "github.com/zoryia/vex/tui/models"
	. "github.com/zoryia/vex/tui/pages"
	"github.com/zoryia/vex/tui/pages/auth"
	"github.com/zoryia/vex/tui/pages/preview"
)

type Model struct {
	list      list.Model
	textInput textinput.Model
	err       error
	auth      auth.Model
	page      VexPage
	query     string
	feeds     []Feed
	entries   []Entry
	tags      []string
	keys      *ListKeyMap
	Preview   preview.Model
}

func New() *Model {
	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.CharLimit = 156
	ti.Width = 56
	return &Model{textInput: ti, auth: auth.New(), page: ENTRIES, keys: NewListKeyMap(), Preview: preview.Model{Viewport: viewport.New(0, 0)}}
}

func (m Model) getEverything() tea.Cmd {
	return func() tea.Msg {
		return tea.Batch(getEntries(m.auth.Jwt)) // getTags, getFeeds)
	}
}

func (m *Model) initList(width int, height int) {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.list.Title = "Posts"
	m.list.SetFilteringEnabled(false)
	var f = Feed{Id: uuid.UUID{}, Tags: []string{"Devops", "Kubernetes"}, Name: "zwindler", Url: "zwindler.blog", FaviconUrl: "zwindler.blog.favicon"}
	m.list.SetItems([]list.Item{
		Entry{Id: uuid.UUID{}, ArticleTitle: "yay", Content: "ouin ouin ouinouin ouin ouinouin ouin ouinouin ouin ouinouin ouin ouinouin ouin ouinouin ouin ouinouin ouin ouinouin ouin ouinouin ouin ouin", Link: "awd", Date: time.Now(), IsRead: false, IsIgnored: false, IsReadLater: false, IsBookmarked: false, Feed: f},
		Entry{Id: uuid.UUID{}, ArticleTitle: "grrrrr", Content: "ouin ouin ouin", Link: "awd", Date: time.Now(), IsRead: false, IsIgnored: false, IsReadLater: false, IsBookmarked: false, Feed: f},
		Entry{Id: uuid.UUID{}, ArticleTitle: "my life is pain", Content: "ouin ouin ouin", Link: "awd", Date: time.Now(), IsRead: false, IsIgnored: false, IsReadLater: false, IsBookmarked: false, Feed: f},
	})
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.auth.LoginForm.Init(), m.auth.RegisterForm.Init(), checkJwt(m.auth.Jwt))

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

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd
	var blurredNow = false
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.initList(msg.Width, msg.Height)
		m.Preview.Viewport.Width = msg.Width
		m.Preview.Viewport.Height = msg.Height - m.Preview.VerticalMarginHeight()

	case invalidJwtMsg:
		m.auth.Jwt = new(string)
		m.page = LOGIN
		return m, nil
	case loginSuccessMsg:
		*m.auth.Jwt = msg.string
		m.page = FEEDS
		return m, m.getEverything()

	case registerSuccessMsg:
		*m.auth.Jwt = msg.string
		m.page = FEEDS
		return m, m.getEverything()

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyCtrlT:
			if m.page == LOGIN {
				m.page = REGISTER
			} else if m.page == REGISTER {
				m.page = LOGIN
			}
		case tea.KeyEnter:
			if m.textInput.Focused() {
				// Get entries with query
				m.textInput.Blur()
				blurredNow = true
			}
		}
		switch {
		case key.Matches(msg, m.textInput.KeyMap.DeleteCharacterBackward): //TODO: add only when query
			words := strings.Split(m.textInput.Value(), " ")
			if len(words) > 0 && (strings.HasPrefix(words[len(words)-1], "tag:") || strings.HasPrefix(words[len(words)-1], "feed:")) {
				m.deleteWordBackward()
			}
		case key.Matches(msg, m.keys.IgnoreToggle) && m.page == "FEEDS":
			var e = m.list.SelectedItem().(Entry)
			cmds = append(cmds, ignorePost(m.auth.Jwt, e))

		case key.Matches(msg, m.keys.ReadToggle):
			var e = m.list.SelectedItem().(Entry)
			cmds = append(cmds, toggleRead(m.auth.Jwt, e))

		case key.Matches(msg, m.keys.ReadLaterToggle):
			var e = m.list.SelectedItem().(Entry)
			cmds = append(cmds, toggleReadLater(m.auth.Jwt, e))

		case key.Matches(msg, m.keys.BookmarkToggle):
			var e = m.list.SelectedItem().(Entry)
			cmds = append(cmds, toggleBookmark(m.auth.Jwt, e))

		case key.Matches(msg, m.keys.Query):
			m.textInput.Focus()
			m.textInput.SetValue("")

		case key.Matches(msg, m.keys.PreviewPost) && m.textInput.Focused() == false && blurredNow == false:
			var e = m.list.SelectedItem().(Entry)
			m.Preview.Entry = e
			m.Preview.Viewport.SetContent(e.Content)
			m.page = PREVIEW
		}

	}

	// Process the form
	// LOGIN
	if m.page == LOGIN {
		form, cmd := m.auth.LoginForm.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.auth.LoginForm = f
			cmds = append(cmds, cmd)
		}

		if m.auth.LoginForm.State == huh.StateCompleted {
			username := m.auth.LoginForm.GetString("email")
			password := m.auth.LoginForm.GetString("password")
			cmds = append(cmds, login(username, password))
		}
	}

	if m.page == REGISTER {

		// Process the form
		// LOGIN
		registerForm, cmd := m.auth.RegisterForm.Update(msg)
		if f, ok := registerForm.(*huh.Form); ok {
			m.auth.RegisterForm = f
			cmds = append(cmds, cmd)
		}

		if m.auth.RegisterForm.State == huh.StateCompleted {
			username := m.auth.RegisterForm.GetString("username")
			password := m.auth.RegisterForm.GetString("password")
			email := m.auth.RegisterForm.GetString("email")
			cmds = append(cmds, register(username, password, email))
		}

	}

	var cmd tea.Cmd
	m.Preview.Viewport, cmd = m.Preview.Viewport.Update(msg)
	cmds = append(cmds, cmd)
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
