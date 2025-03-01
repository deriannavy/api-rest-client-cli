package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/deriannavy/api-rest-client-cli/handler"
)

type Panel struct {
	// Styles & Keymaps
	Styles ListStyle
	KeyMap handler.KeyMap
	// Window Size
	Size handler.SizeSpec
	// Components
	ItemComplement ItemComplement
	// Item
	Item Item
}

func NewPanel(item Item, width, height int) Panel {
	return Panel{
		// Styles & Keymaps
		Styles: DefaultListStyle(),
		KeyMap: handler.DefaultKeyMap(),
		// Window Size
		Size: handler.NewSizeSpec(width, height),
		// Components
		ItemComplement: NewComplement(width, 1),
		// Item
		Item: item,
	}
}

func (p *Panel) SetItem(item Item) {
	p.Item = item
}

func (p Panel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	var cmds []tea.Cmd

	// switch msg := msg.(type) {

	// 	case tea.KeyMsg:
	// 	switch {
	// 	case key.Matches(msg, l.KeyMap.CursorUp):
	// 		cmds = append(cmds, l.CursorUp())

	// 	case key.Matches(msg, l.KeyMap.CursorDown):
	// 		cmds = append(cmds, l.CursorDown())
	// 	}

	// }

	return p, tea.Batch(cmds...)
}

func (p Panel) View() string {

	var b strings.Builder

	fmt.Fprintf(&b, "%s\n", p.Item.TitleFormat(p.ItemComplement, true))

	fmt.Fprintf(&b, "%s%s\n", p.Item.MethodFormat(p.ItemComplement, "right"), p.Item.UrlFormat(p.ItemComplement))

	return b.String()
}
