package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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
	RequestTab     Tabs
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
		RequestTab:     NewTabComponent(item, []string{"Params", "Headers", "Body"}, width, 1),
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

	p.RequestTab, cmdTabs = p.RequestTab.Update(msg)
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

	RequestTabs := p.RequestTab.View()
	p.Size.AddUsedHeight(true, lipgloss.Height(RequestTabs))

	rows := [][]string{
		{"offset", "5", "About what digit start"},
		{"limit", "10", "Limit set to the list"},
		{"page_size", "20", "Page size"},
		{"——————", "20", "Page size"},
	}

	tt := NewTable(
		[]string{"Key", "Value", "Description"},
		[][]string{[]string{}},
	)

	t := table.New().
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#444444"))).
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false).
		BorderBottom(true).
		// BorderColumn(false).
		BorderRow(true).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Bold(true).Align(lipgloss.Center)
			case row%2 == 0:
				return lipgloss.NewStyle().Padding(0, 1).Width(14).Foreground(lipgloss.Color("241"))
			default:
				return lipgloss.NewStyle().Padding(0, 1).Width(14).Foreground(lipgloss.Color("245"))
			}
		}).
		Headers("key", "Value", "Description").
		Rows(rows...)

	// strings.Repeat("\n", p.Size.AvailableHeight())
	return p.Styles.BorderLeftStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			p.Render(),
			RequestTabs,
			tt.View(),
			t.String(),
		),
	)

}
