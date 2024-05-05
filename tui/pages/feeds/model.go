package feeds

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/huh"
	"github.com/google/uuid"

	"github.com/zoryia/vex/tui/models"
)

type Model struct {
	List    list.Model
	AddFeed *huh.Form
}

func (m Model) View() string {
	return m.List.View()
}

func FeedsKeys() []key.Binding {
	return []key.Binding{
		key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "Add a new feed"),
		),
		key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "Go to posts"),
		),
		key.NewBinding(
			key.WithKeys("x", "d"),
			key.WithHelp("x/d", "Ignore feed"),
		),
	}
}

func initFeedsList() list.Model {
	feeds := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	feeds.Title = "Feeds"
	feeds.SetFilteringEnabled(false)
	feeds.DisableQuitKeybindings()
	feeds.AdditionalShortHelpKeys = FeedsKeys
	feeds.AdditionalFullHelpKeys = FeedsKeys
	feeds.SetItems([]list.Item{
		models.Feed{Id: uuid.UUID{}, Tags: []string{"Devops", "Kubernetes"}, Name: "zwindler", Url: "zwindler.blog", FaviconUrl: "zwindler.blog.favicon"},
		models.Feed{Id: uuid.UUID{}, Tags: []string{"Devops", "Kubernetes"}, Name: "zwindler", Url: "zwindler.blog", FaviconUrl: "zwindler.blog.favicon"},
		models.Feed{Id: uuid.UUID{}, Tags: []string{"Devops", "Kubernetes"}, Name: "zwindler", Url: "zwindler.blog", FaviconUrl: "zwindler.blog.favicon"},
	})
	return feeds
}

func AddFeedForm(tags []string) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Feed Url").
				Key("url"),
			huh.NewMultiSelect[string]().
				Title("Tags").
				Key("tags").
				Options(

					huh.NewOptions(tags...)...,
				),
		)).WithWidth(40)
}

func New() Model {
	return Model{List: initFeedsList(), AddFeed: AddFeedForm([]string{})}

}
