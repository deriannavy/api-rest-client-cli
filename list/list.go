package list

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/deriannavy/api-rest-client-cli/item"
	"github.com/deriannavy/api-rest-client-cli/paginator"
)

type Model struct {
	// Styles
	Styles Styles

	// Titles & text
	itemNameSingular string
	itemNamePlural   string

	// Components
	Paginator      paginator.Model
	itemComplement item.ItemComplement

	// Window Size
	width  int
	height int

	// Items & indexs
	cursor int
	items  []item.Item
}

func New(items []item.Item, width, height int) Model {

	s := DefaultStyles()

	ic := item.NewComplement()

	p := paginator.New()
	p.ActiveDot = s.ActivePaginationDot.String()
	p.InactiveDot = s.InactivePaginationDot.String()

	m := Model{
		Styles:           DefaultStyles(),
		itemNameSingular: "item",
		itemNamePlural:   "items",
		itemComplement:   ic,
		Paginator:        p,
		width:            width,
		height:           height,

		// > Lists
		items: items,
	}

	m.updatePagination()

	return m

}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {

	items := m.items

	var b strings.Builder

	// Empty states
	if len(items) <= 0 {
		return m.Styles.NoItems.Render("No " + m.itemNamePlural + ".")
	}

	if len(items) > 0 {
		start, end := m.Paginator.GetSliceBounds(len(items))
		docs := items[start:end]

		for i, item := range docs {

			isSelected := i == m.Index()

			m.itemComplement.Render(&b, m.width, isSelected, i+start, item)
			
			if i != len(docs)-1 {
				fmt.Fprint(
					&b, 
					strings.Repeat(
						"\n", 
						m.itemComplement.Spacing() + 1,
					),
				)
			}
		}
	}

	return b.String()
}

// Index returns the index of the currently selected item as it is stored in the
// filtered list of items.
// Using this value with SetItem() might be incorrect, consider using
// GlobalIndex() instead.
func (m Model) Index() int {
	return m.Paginator.Page*m.Paginator.PerPage + m.cursor
}

// SetSize sets the width and height of this component.
func (m *Model) SetSize(width, height int) {
	m.setSize(width, height)
}

// > Set size private function
func (m *Model) setSize(width, height int) {
	m.width = width
	m.height = height
	m.updatePagination()
}

// Update pagination according to the amount of items for the current state.
func (m *Model) updatePagination() {
	index := m.Index()
	availHeight := m.height

	m.Paginator.PerPage = max(1, availHeight/m.itemComplement.TotalHeight())

	if pages := len(m.items); pages < 1 {
		m.Paginator.SetTotalPages(1)
	} else {
		m.Paginator.SetTotalPages(pages)
	}

	// Restore index
	m.Paginator.Page = index / m.Paginator.PerPage
	m.cursor = index % m.Paginator.PerPage

	// Make sure the page stays in bounds
	if m.Paginator.Page >= m.Paginator.TotalPages-1 {
		m.Paginator.Page = max(0, m.Paginator.TotalPages-1)
	}
}
