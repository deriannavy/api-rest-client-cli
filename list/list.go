package list

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deriannavy/api-rest-client-cli/item"
	"github.com/deriannavy/api-rest-client-cli/paginator"
)

type Model struct {
	// Styles & Keymaps
	Styles Styles
	KeyMap KeyMap

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
		Styles:           s,
		KeyMap:           DefaultKeyMap(),
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

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.CursorUp):
			m.CursorUp()

		case key.Matches(msg, m.KeyMap.CursorDown):
			m.CursorDown()
		}
	}

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

			isSelected := i == m.Cursor()

			m.itemComplement.Render(&b, m.width, isSelected, i+start, item)

			if i != len(docs)-1 {
				fmt.Fprint(
					&b,
					strings.Repeat(
						"\n",
						m.itemComplement.Spacing()+1,
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

// Cursor returns the index of the cursor on the current page.
func (m Model) Cursor() int {
	return m.cursor
}

// CursorUp moves the cursor up. This can also move the state to the previous
// page.
func (m *Model) CursorUp() {
	m.cursor--

	// If we're at the start, stop
	if m.cursor < 0 && m.Paginator.Page == 0 {
		// if infinite scrolling is enabled, go to the last item
		m.Paginator.Page = m.Paginator.TotalPages - 1
		m.cursor = m.Paginator.ItemsOnPage(len(m.items)) - 1
		return
	}

	// Move the cursor as normal
	if m.cursor >= 0 {
		return
	}

	// Go to the previous page
	m.Paginator.PrevPage()
	m.cursor = m.Paginator.ItemsOnPage(len(m.items)) - 1
}

// CursorDown moves the cursor down. This can also advance the state to the
// next page.
func (m *Model) CursorDown() {
	itemsOnPage := m.Paginator.ItemsOnPage(len(m.items))

	m.cursor++

	// If we're at the end, stop
	if m.cursor < itemsOnPage {
		return
	}

	// Go to the next page
	if !m.Paginator.OnLastPage() {
		m.Paginator.NextPage()
		m.cursor = 0
		return
	}

	// During filtering the cursor position can exceed the number of
	// itemsOnPage. It's more intuitive to start the cursor at the
	// topmost position when moving it down in this scenario.
	if m.cursor > itemsOnPage {
		m.cursor = 0
		return
	}

	m.cursor = itemsOnPage - 1

	// if infinite scrolling is enabled, go to the first item
	m.Paginator.Page = 0
	m.cursor = 0
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
