package paginator

// Model is the Bubble Tea model for this user interface.
type Model struct {
	// Page is the current page number.
	Page int
	// PerPage is the number of items per page.
	PerPage int
	// TotalPages is the total number of pages.
	TotalPages int
	// ActiveDot is used to mark the current page under the Dots display type.
	ActiveDot string
	// InactiveDot is used to mark inactive pages under the Dots display type.
	InactiveDot string

	// KeyMap encodes the keybindings recognized by the widget.
	KeyMap KeyMap
}

func New() Model {
	return Model{
		Page:        0,
		PerPage:     1,
		TotalPages:  1,
		KeyMap:      DefaultKeyMap,
		ActiveDot:   "•",
		InactiveDot: "○",
	}
}

// PrevPage is a helper function for navigating one page backward. It will not
// page beyond the first page (i.e. page 0).
func (m *Model) PrevPage() {
	if m.Page > 0 {
		m.Page--
	}
}

// NextPage is a helper function for navigating one page forward. It will not
// page beyond the last page (i.e. totalPages - 1).
func (m *Model) NextPage() {
	if !m.OnLastPage() {
		m.Page++
	}
}

// OnLastPage returns whether or not we're on the last page.
func (m Model) OnLastPage() bool {
	return m.Page == m.TotalPages-1
}

// OnFirstPage returns whether or not we're on the first page.
func (m Model) OnFirstPage() bool {
	return m.Page == 0
}

// ItemsOnPage is a helper function for returning the number of items on the
// current page given the total number of items passed as an argument.
func (m Model) ItemsOnPage(totalItems int) int {
	if totalItems < 1 {
		return 0
	}
	start, end := m.GetSliceBounds(totalItems)
	return end - start
}

// GetSliceBounds is a helper function for paginating slices. Pass the length
// of the slice you're rendering and you'll receive the start and end bounds
// corresponding to the pagination. For example:
//
//	bunchOfStuff := []stuff{...}
//	start, end := model.GetSliceBounds(len(bunchOfStuff))
//	sliceToRender := bunchOfStuff[start:end]
func (m *Model) GetSliceBounds(length int) (start int, end int) {
	start = m.Page * m.PerPage
	end = min(m.Page*m.PerPage+m.PerPage, length)
	return start, end
}

// SetTotalPages is a helper function for calculating the total number of pages
// from a given number of items. Its use is optional since this pager can be
// used for other things beyond navigating sets. Note that it both returns the
// number of total pages and alters the model.
func (m *Model) SetTotalPages(items int) int {
	if items < 1 {
		return m.TotalPages
	}
	n := items / m.PerPage
	if items%m.PerPage > 0 {
		n++
	}
	m.TotalPages = n
	return n
}

// > Get the min integer
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
