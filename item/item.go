package item

import (
	"fmt"
	"io"
)

type Item struct {
	// > Json data
	Name    string  `json:"name"`
	Request Request `json:"request"`
}

type ItemComplement struct {
	// > App data
	Styles  DefaultItemStyles
	spacing int
	height  int
}

// NewDefaultDelegate creates a new delegate with default styles.
func NewComplement() ItemComplement {
	return ItemComplement{
		Styles:  NewDefaultItemStyles(),
		height:  1,
		spacing: 0,
	}
}

func (i Item) View() string {
	return i.Name
}

// Height returns the delegate's preferred height.
// This has effect only if ShowDescription is true,
// otherwise height is always 1.
func (ic ItemComplement) Height() int {
	return ic.height
}

// Spacing returns the delegate's spacing.
func (ic ItemComplement) Spacing() int {
	return ic.spacing
}

func (ic ItemComplement) TotalHeight() int {
	return ic.Height() + ic.Spacing()
}

func (ic ItemComplement) Render(w io.Writer, isSelected bool, item Item) {

	var (
		title string
		s     = &ic.Styles
	)

	title = item.Request.Method + " " + item.Name

	// textwidth := width - s.NormalTitle.GetPaddingLeft() - s.NormalTitle.GetPaddingRight()
	// title = ansi.Truncate(title, textwidth, ellipsis)

	if isSelected {
		title = s.SelectedTitle.Render(title)
	} else {
		title = s.NormalTitle.Render(title)
	}

	fmt.Fprintf(w, "%s\n", title) //nolint: errcheck
}
