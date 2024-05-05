package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	huh "github.com/charmbracelet/huh"
)

func (m Model) handleSearchCompletion() (Model, tea.Cmd) {
	var cmd tea.Cmd
	queryWords := strings.Split(m.queryInput.Value(), " ")
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
		m.queryInput.SetValue(m.queryInput.Value() + feed)
		m.queryInput.CursorEnd()
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
		m.queryInput.SetValue(m.queryInput.Value() + feed)
		m.queryInput.CursorEnd()
	}
	return m, cmd
}

func (m *Model) deleteWordBackward() {
	if m.queryInput.Position() == 0 || len(m.queryInput.Value()) == 0 {
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
	oldPos := m.queryInput.Position() //nolint:ifshort

	m.queryInput.SetCursor(oldPos - 1)
	// ECHO character?
	for m.queryInput.Value()[m.queryInput.Position()] == ' ' {
		if m.queryInput.Position() <= 0 {
			break
		}
		// ignore series of whitespace before cursor
		m.queryInput.SetCursor(m.queryInput.Position() - 1)
	}

	for m.queryInput.Position() > 0 {
		if m.queryInput.Value()[m.queryInput.Position()] != ' ' {
			m.queryInput.SetCursor(m.queryInput.Position() - 1)
		} else {
			if m.queryInput.Position() > 0 {
				// keep the previous space
				m.queryInput.SetCursor(m.queryInput.Position() + 1)
			}
			break
		}
	}

	if oldPos > len(m.queryInput.Value()) {
		m.queryInput.SetValue(m.queryInput.Value()[:m.queryInput.Position()])
	} else {
		m.queryInput.SetValue(m.queryInput.Value()[:m.queryInput.Position()] + m.queryInput.Value()[oldPos:])
	}
}
