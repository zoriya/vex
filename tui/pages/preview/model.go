package preview

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/zoryia/vex/tui/models"

	"errors"
)

type Model struct {
	Entry    models.Entry
	Viewport viewport.Model
}

// RenderMarkdown renders the markdown content with glamour.
func RenderMarkdown(width int, content string) (string, error) {
	background := "light"

	if lipgloss.HasDarkBackground() {
		background = "dark"
	}

	r, _ := glamour.NewTermRenderer(
		glamour.WithWordWrap(width),
		glamour.WithStandardStyle(background),
	)

	out, err := r.Render(content)
	if err != nil {
		return "", errors.Unwrap(err)
	}

	return out, nil
}

func renderMarkdownCmd(width int, entry models.Entry) tea.Cmd {
	return func() tea.Msg {
		_, err := RenderMarkdown(width, entry.Content)
		if err != nil {
			//return errorMsg(err)
		}
		return nil
		// return renderMarkdownMsg(markdownContent)
	}
}

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
)

// TODO: render as markdown with glamour
func (m Model) View() string {
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.Viewport.View(), m.footerView())
}

func (m Model) VerticalMarginHeight() int {
	headerHeight := lipgloss.Height(m.headerView())
	footerHeight := lipgloss.Height(m.footerView())
	return headerHeight + footerHeight
}

func (m Model) headerView() string {
	title := titleStyle.Render(m.Entry.ArticleTitle)
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m Model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}
