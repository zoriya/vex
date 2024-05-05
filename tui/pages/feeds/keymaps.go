package feeds

import "github.com/charmbracelet/bubbles/key"

type ListKeyMap struct {
	GoToPosts key.Binding
	AddFeed   key.Binding
}

func FeedsKeyMaps() *ListKeyMap {
	return &ListKeyMap{
		GoToPosts: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "Go to posts"),
		),
		AddFeed: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "Add feed"),
		),
	}
}
