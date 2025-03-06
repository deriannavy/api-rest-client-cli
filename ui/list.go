package ui

import (
	"fmt"
	"strconv"
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
	ItemComplement ItemComplement
	// Window Size
	Size handler.SizeSpec
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
		ItemComplement: NewComplement(width, 1),
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
	return max(1, l.Size.AvailableHeight()/l.ItemComplement.Size.Height())
}

func (l List) ShowPageDot(index int) string {
	if index == l.CurrentNumberPage() {
		return l.Styles.ActivePaginationDot.Render(bullet)
	}
	return l.Styles.InactivePaginationDot.Render(bullet)
}

// Generate pagination
func (l List) GeneratePagination() string {
	pagination := "   "
	for i := range l.TotalPages() {
		pagination += l.ShowPageDot(i)
	}
	return pagination
}

// Get the Total Pages based on the page size and total Items
// If the total page is equal to 1 return 0
func (l List) TotalPages() int {
	return (l.itemsLength / l.PageSize()) + 1
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
			cmds = append(cmds, l.CursorUp())

		case key.Matches(msg, l.KeyMap.CursorDown):
			cmds = append(cmds, l.CursorDown())
		}
	}

	return l, tea.Batch(cmds...)
}

func (l *List) CursorUp() tea.Cmd {
	n := handler.TernaryNumber(l.index == 0, (l.itemsLength - 1), l.index-1)
	l.index = n
	return func() tea.Msg {
		return handler.NewCursorMoveMsg(l.index)
	}
}

func (l *List) CursorDown() tea.Cmd {
	n := handler.TernaryNumber(l.index == (l.itemsLength-1), 0, l.index+1)
	l.index = n
	return func() tea.Msg {
		return handler.NewCursorMoveMsg(l.index)
	}
}

func (l List) View() string {

	// Empty states
	if l.itemsLength <= 0 {
		return l.Styles.NoItems.Render("No Items.")
	}

	var b strings.Builder

	// T I T L E | Show title and item count
	l.Size.AddUsedHeight(false, 2)
	fmt.Fprintf(&b, "    %s\n\n", "Request • "+strconv.Itoa(l.itemsLength)+" items")

	// P A G I N A T O R | Save space for pagination if this is greater than 1
	l.Size.AddUsedHeight(false, 2)
	pagination := l.GeneratePagination()

	// L I S T | Print list lines
	for _, item := range l.CurrentPageItems() {
		l.Size.AddUsedHeight(false, 1)
		fmt.Fprintf(&b, "%s\n", item.View(l.ItemComplement, l.index == item.Index))
	}

	// L I S T | Set lines for every available line to push pagination at the end
	fmt.Fprint(&b, strings.Repeat("\n", l.Size.AvailableHeight()))

	// P A G I N A T O R | Print pagination
	fmt.Fprintf(&b, "\n%s", pagination)

	return b.String()
}
