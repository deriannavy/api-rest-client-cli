package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deriannavy/api-rest-client-cli/handler"
)

type Panel struct {
	// Styles & Keymaps
	Styles PanelStyle
	KeyMap handler.KeyMap
	// Window Size
	Size handler.SizeSpec
	// Components
	ItemComplement ItemComplement
	Tabs           Tabs
	RequestTabs    Tabs
	ResponseTabs   Tabs
	// Item
	Item Item
}

func NewPanel(item Item, width, height int) Panel {
	return Panel{
		// Styles & Keymaps
		Styles: DefaultPanelStyle(),
		KeyMap: handler.DefaultKeyMap(),
		// Window Size
		Size: handler.NewSizeSpec(width, height),
		// Components
		ItemComplement: NewComplement(width, 1),
		Tabs:           NewTabComponent("Horizontal", []string{"Request", "Response"}, width, 1),
		RequestTabs:    NewTabComponent("Vertical", []string{"Params", "Headers", "Body"}, width, 1),
		ResponseTabs:   NewTabComponent("Vertical", []string{"Headers", "Body"}, width, 1),
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

	// switch msg := msg.(type) {

	// 	case tea.KeyMsg:
	// 	switch {
	// 	case key.Matches(msg, l.KeyMap.CursorUp):
	// 		cmds = append(cmds, l.CursorUp())

	// 	case key.Matches(msg, l.KeyMap.CursorDown):
	// 		cmds = append(cmds, l.CursorDown())
	// 	}

	// }

	p.Tabs, cmdTabs = p.Tabs.Update(msg)
	cmds = append(cmds, cmdTabs)

	return p, tea.Batch(cmds...)
}

func (p Panel) Render() string {
	var b strings.Builder

	fmt.Fprintf(&b, "%s\n", p.Item.TitleFormat(p.ItemComplement, true))
	p.Size.AddUsedHeight(false, 2)
	fmt.Fprintf(&b, "%s%s\n", p.Item.MethodFormat(p.ItemComplement, "right"), p.Item.UrlFormat(p.ItemComplement))
	p.Size.AddUsedHeight(false, 1)

	return b.String()
}

func (p Panel) View() string {

	tabs := p.Tabs.View()
	p.Size.AddUsedHeight(true, lipgloss.Height(tabs))

	RequestTabs := p.RequestTabs.View()
	p.Size.AddUsedHeight(false, lipgloss.Height(RequestTabs))
	// strings.Repeat("\n", p.Size.AvailableHeight())
	return p.Styles.BorderLeftStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			tabs,
			p.Render(),
			RequestTabs,
		),
	)

}
