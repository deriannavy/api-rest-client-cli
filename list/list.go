package list

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/deriannavy/api-rest-client-cli/item"
)

type Model struct {
	Styles Styles

	itemNameSingular string
	itemNamePlural   string

	width  int
	height int

	items []item.Item
}

func New(items []item.Item, width, height int) Model {

	return Model{
		Styles:           DefaultStyles(),
		itemNameSingular: "item",
		itemNamePlural:   "items",

		width:  width,
		height: height,
		items:  items,
	}

}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {

	items := m.items

	var b strings.Builder

	// Empty states
	if len(items) == 0 {
		return m.Styles.NoItems.Render("No " + m.itemNamePlural + ".")
	}

	// if len(items) > 0 {
	// 	start, end := m.Paginator.GetSliceBounds(len(items))
	// 	docs := items[start:end]

	// 	for i, item := range docs {
	// 		m.delegate.Render(&b, m, i+start, item)
	// 		if i != len(docs)-1 {
	// 			fmt.Fprint(&b, strings.Repeat("\n", m.delegate.Spacing()+1))
	// 		}
	// 	}
	// }
	fmt.Fprint(&b, strconv.Itoa(m.width))
	fmt.Fprint(&b, strconv.Itoa(m.height))

	return b.String()
}

// SetSize sets the width and height of this component.
func (m *Model) SetSize(width, height int) {
	m.setSize(width, height)
}

func (m *Model) setSize(width, height int) {
	m.width = width
	m.height = height
}
