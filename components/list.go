package components

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

func NewList(items []list.Item) list.Model {

	ls := list.NewDefaultDelegate()

	c := lipgloss.Color("#ff00ff")
	style := ls.Styles.SelectedTitle.Foreground(c).BorderLeftForeground(c)
	ls.Styles.SelectedTitle = style
	ls.Styles.SelectedDesc = style

	nlist := list.New(items, ls, 0, 0)

	nlist.SetShowTitle(false)
	nlist.SetShowStatusBar(false)
	nlist.InfiniteScrolling = true

	return nlist
}
