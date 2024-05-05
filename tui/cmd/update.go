package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	huh "github.com/charmbracelet/huh"
	. "github.com/zoryia/vex/tui/models"
	. "github.com/zoryia/vex/tui/pages"
	"github.com/zoryia/vex/tui/pages/feeds"
)

func (m Model) LoginUpdate(msg tea.Msg) (tea.Model, []tea.Cmd) {
	var cmds []tea.Cmd
	if m.page != LOGIN {
		return m, nil
	}
	return m, cmds
}

func (m Model) GlobalUpdate(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.initList(msg.Width, msg.Height-2)
		m.queryInput.Width = msg.Width - 5
		m.Preview.Viewport.Width = msg.Width
		m.Preview.Viewport.Height = msg.Height - m.Preview.VerticalMarginHeight()
		m.Feeds.List.SetWidth(msg.Width)
		m.Feeds.List.SetHeight(msg.Height)

	case getEntriesSuccessMsg:
		var entries []list.Item
		for _, e := range msg {
			entries = append(entries, e)
		}
		m.list.SetItems(entries)

	case invalidJwtMsg:
		m.Auth.Jwt = new(string)
		m.page = LOGIN
		return m, nil

	case loginSuccessMsg:
		*m.Auth.Jwt = msg.string
		m.page = ENTRIES
		return m, m.getEverything()

	case registerSuccessMsg:
		*m.Auth.Jwt = msg.string
		m.page = ENTRIES
		return m, m.getEverything()

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}
	return m, nil
}

func loginUpdate(m Model, msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlT:
			m.page = REGISTER
		}
	}
	form, cmd := m.Auth.LoginForm.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.Auth.LoginForm = f
		cmds = append(cmds, cmd)
	}

	if m.Auth.LoginForm.State == huh.StateCompleted {
		username := m.Auth.LoginForm.GetString("email")
		password := m.Auth.LoginForm.GetString("password")
		cmds = append(cmds, login(username, password))
	}

	return m, tea.Batch(cmds...)
}
func registerUpdate(m Model, msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlT:
			m.page = LOGIN
		}
	}

	// Process the form
	// LOGIN
	registerForm, cmd := m.Auth.RegisterForm.Update(msg)
	if f, ok := registerForm.(*huh.Form); ok {
		m.Auth.RegisterForm = f
		cmds = append(cmds, cmd)
	}

	if m.Auth.RegisterForm.State == huh.StateCompleted {
		username := m.Auth.RegisterForm.GetString("username")
		password := m.Auth.RegisterForm.GetString("password")
		email := m.Auth.RegisterForm.GetString("email")
		cmds = append(cmds, register(username, password, email))
	}

	return m, tea.Batch(cmds...)
}
func entriesUpdate(m Model, msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	var blurredNow = false
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.queryInput.Focused() {
				// Get entries with query
				m.queryInput.Blur()
				blurredNow = true
			}
		}
		switch {
		case key.Matches(msg, m.queryInput.KeyMap.DeleteCharacterBackward) && m.queryInput.Focused():
			words := strings.Split(m.queryInput.Value(), " ")
			if len(words) > 0 && (strings.HasPrefix(words[len(words)-1], "tag:") || strings.HasPrefix(words[len(words)-1], "feed:")) {
				m.deleteWordBackward()
			}
		}
		if m.queryInput.Focused() == false {
			switch {
			case key.Matches(msg, m.keys.GoToFeeds):
				m.page = FEEDS

			case key.Matches(msg, m.keys.IgnoreToggle):
				var e = m.list.SelectedItem().(Entry)
				cmds = append(cmds, ignorePost(m.Auth.Jwt, e))

			case key.Matches(msg, m.keys.ReadToggle):
				var e = m.list.SelectedItem().(Entry)
				cmds = append(cmds, toggleRead(m.Auth.Jwt, e))

			case key.Matches(msg, m.keys.ReadLaterToggle):
				var e = m.list.SelectedItem().(Entry)
				cmds = append(cmds, toggleReadLater(m.Auth.Jwt, e))

			case key.Matches(msg, m.keys.BookmarkToggle):
				var e = m.list.SelectedItem().(Entry)
				cmds = append(cmds, toggleBookmark(m.Auth.Jwt, e))

			case key.Matches(msg, m.keys.Query):
				m.queryInput.Focus()
				m.queryInput.SetValue("")

			case key.Matches(msg, m.keys.PreviewPost) && m.queryInput.Focused() == false && blurredNow == false:
				var e = m.list.SelectedItem().(Entry)
				m.Preview.Entry = e
				m.Preview.Viewport.SetContent(e.Content)
				m.page = PREVIEW

			}
		}
		m, cmd = m.handleSearchCompletion()
		cmds = append(cmds, cmd)
	}
	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)
	m.queryInput, cmd = m.queryInput.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
func feedsUpdate(m Model, msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, feeds.FeedsKeyMaps().AddFeed):
			m.Feeds.AddFeed.Run()
		}
	}
	registerForm, cmd := m.Auth.RegisterForm.Update(msg)
	if f, ok := registerForm.(*huh.Form); ok {
		m.Auth.RegisterForm = f
		cmds = append(cmds, cmd)
	}

	if m.Auth.RegisterForm.State == huh.StateCompleted {
		url := m.Auth.RegisterForm.GetString("url")
		tags := m.Auth.RegisterForm.Get("tags").([]string)
		cmds = append(cmds, addFeed(m.Auth.Jwt, url, tags))
	}
	m.Feeds.List, cmd = m.Feeds.List.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
func tagsUpdate(m Model, msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}
func previewUpdate(m Model, msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	m.Preview.Viewport, cmd = m.Preview.Viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
func ignoredUpdate(m Model, msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

type updateFunc func(Model, tea.Msg) (Model, tea.Cmd)

func getUpdateMap() map[VexPage]updateFunc {

	updateMap := make(map[VexPage]updateFunc)
	updateMap[LOGIN] = loginUpdate
	updateMap[REGISTER] = registerUpdate
	updateMap[ENTRIES] = entriesUpdate
	updateMap[FEEDS] = feedsUpdate
	updateMap[TAGS] = tagsUpdate
	updateMap[IGNORED] = ignoredUpdate
	updateMap[PREVIEW] = previewUpdate
	return updateMap
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	updateMap := getUpdateMap()
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m, cmd = m.GlobalUpdate(msg)
	cmds = append(cmds, cmd)
	m, cmd = updateMap[m.page](m, msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
