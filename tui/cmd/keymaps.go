package main

import "github.com/charmbracelet/bubbles/key"

type ListKeyMap struct {
	Query           key.Binding
	BookmarkToggle  key.Binding
	ReadToggle      key.Binding
	ReadLaterToggle key.Binding
	IgnoreToggle    key.Binding
	PreviewPost     key.Binding
	GoToFeeds       key.Binding
	GoToPosts       key.Binding
}

func NewListKeyMap() *ListKeyMap {
	return &ListKeyMap{
		GoToFeeds: key.NewBinding(
			key.WithKeys("f"),
			key.WithHelp("f", "go to feeds"),
		),
		GoToPosts: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "go to posts"),
		),
		PreviewPost: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "preview post"),
		),
		Query: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "query posts"),
		),
		BookmarkToggle: key.NewBinding(
			key.WithKeys("b"),
			key.WithHelp("b", "toggle bookmarked"),
		),
		ReadToggle: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "toggle mark as read"),
		),

		IgnoreToggle: key.NewBinding(
			key.WithKeys("x", "d"),
			key.WithHelp("d/x", "ignore post"),
		),
		ReadLaterToggle: key.NewBinding(
			key.WithKeys("m"),
			key.WithHelp("m", "add to read later"),
		),
	}
}
