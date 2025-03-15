package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deriannavy/api-rest-client-cli/handler"
	"github.com/deriannavy/api-rest-client-cli/styles"
)

type Panel struct {
	// Styles & Keymaps
	Styles styles.PanelStyle
	KeyMap handler.KeyMap
	// Window Size
	Size handler.SizeSpec
	// Components
	ItemComplement ItemComplement
	Tabs           Tabs
	// Item
	Item Item
}

func NewPanel(item Item, width, height int) Panel {

	return Panel{
		// Styles & Keymaps
		Styles: styles.DefaultPanelStyle(),
		KeyMap: handler.DefaultKeyMap(),
		// Window Size
		Size: handler.NewSizeSpec(width, height),
		// Components
		ItemComplement: NewComplement(width, 1),
		Tabs:           NewTabComponent(item, width, 1),
		// Item
		Item: item,
	}
}

func (p *Panel) SetItem(item Item) {
	p.Item = item
}

func (p Panel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	var (
		cmds    []tea.Cmd
		cmdTabs tea.Cmd
	)

	switch msg.(type) {
	case handler.CursorMoveMsg:
		p.Tabs.SetItem(p.Item)
	}

	p.Tabs, cmdTabs = p.Tabs.Update(msg)
	cmds = append(cmds, cmdTabs)

	return p, tea.Batch(cmds...)
}

func (p Panel) Render() string {
	var b strings.Builder

	fmt.Fprintf(&b, "%s\n", p.Item.TitleFormat(p.ItemComplement, true))
	p.Size.AddUsedHeight(false, 2)
	fmt.Fprintf(&b, "%s%s\n", p.Item.MethodFormatStyle(p.ItemComplement, "left", false), p.Item.UrlFormat(p.ItemComplement))
	p.Size.AddUsedHeight(false, 1)

	return b.String()
}

func (p Panel) View() string {

	Tabs := p.Tabs.View()
	p.Size.AddUsedHeight(true, lipgloss.Height(Tabs))

	// t := NewTable()

	// t.AddHeaders("Key", "Value", "Description")
	// t.AddRow("offset", "5", "About what digit start")
	// t.AddRow("limit", "10", "Limit set to the list")
	// t.AddRow("page_size", "20", "Page size")
	// t.AddRow("asss", "20", "Page size")

	// strings.Repeat("\n", p.Size.AvailableHeight())
	return p.Styles.BorderLeftStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			p.Render(),
			Tabs,
		),
	)

}
