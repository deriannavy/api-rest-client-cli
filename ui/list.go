package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deriannavy/api-rest-client-cli/handler"
)

type List struct {
	// Styles & Keymaps
	Styles ListStyle
	KeyMap handler.KeyMap
	// Components
	itemComplement ItemComplement
	// Window Size
	Size handler.SizeSpec
	// width  int
	// height int
	// Items & index
	index       int
	itemsLength int
	items       []Item
}

func NewList(items []Item, width, height int) List {
	return List{
		// Styles & Keymaps
		Styles: DefaultListStyle(),
		KeyMap: handler.DefaultKeyMap(),
		// Components
		itemComplement: NewComplement(width, 1),
		// Window Size
		Size: handler.NewSizeSpec(width, height),
		// Items & index
		index:       0,
		items:       items,
		itemsLength: len(items),
	}
}

// Get the page Size based on the list height / single item height
func (l List) PageSize() int {
	return max(1, l.Size.Height()/l.itemComplement.Size.Height())
}

// Get the number current page based in the index
func (l List) CurrentNumberPage() int {
	return l.index / l.PageSize()
}

func (l List) CurrentPageBounds() (int, int) {
	start := (l.PageSize() * l.CurrentNumberPage())
	end := min(l.itemsLength, start+l.PageSize())
	return start, end
}

func (l List) CurrentPageItems() []Item {
	start, end := l.CurrentPageBounds()
	return l.items[start:end]
}

func (l List) Update(msg tea.Msg) (List, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, l.KeyMap.CursorUp):
			l.CursorUp()

		case key.Matches(msg, l.KeyMap.CursorDown):
			l.CursorDown()
		}
	}

	return l, tea.Batch(cmds...)
}

func (l *List) CursorUp() {
	if l.index == 0 {
		l.index = (l.itemsLength - 1)
	} else {
		l.index--
	}
}

func (l *List) CursorDown() {
	if l.index == (l.itemsLength - 1) {
		l.index = 0
	} else {
		l.index++
	}
}

func (l List) View() string {

	var b strings.Builder

	// Empty states
	if l.itemsLength <= 0 {
		return l.Styles.NoItems.Render("No Items.")
	}

	for _, item := range l.CurrentPageItems() {
		// item := l.itemComplement.Render(item, l.index)
		item := item.View(l.itemComplement, l.index == item.Index)
		fmt.Fprintf(&b, "%s\n", item)
	}

	return b.String()
}
