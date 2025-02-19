package item

import (
	"fmt"
	"io"
	"strings"
	"github.com/charmbracelet/x/ansi"
)

type Item struct {
	// > Json data
	Name    string  `json:"name"`
	Request Request `json:"request"`
}


type ItemComplement struct {
	// > App data
	ShowDescription bool
	Styles          DefaultItemStyles
	spacing         int
	height          int
}


// NewDefaultDelegate creates a new delegate with default styles.
func NewComplement() ItemComplement {
	const defaultHeight = 2
	const defaultSpacing = 1
	return ItemComplement{
		ShowDescription: true,
		Styles:          NewDefaultItemStyles(),
		height:          defaultHeight,
		spacing:         defaultSpacing,
	}
}

// Height returns the delegate's preferred height.
// This has effect only if ShowDescription is true,
// otherwise height is always 1.
func (ic ItemComplement) Height() int {
	if ic.ShowDescription {
		return ic.height
	}
	return 1
}

// Spacing returns the delegate's spacing.
func (ic ItemComplement) Spacing() int {
	return ic.spacing
}

func (ic ItemComplement) TotalHeight() int {
	return ic.Height() + ic.Spacing()
}

 
func (ic ItemComplement) Render(w io.Writer, width int, isSelected bool, index int, item Item) {

	var (
		title, desc  string
		s            = &ic.Styles
	)

	title = item.Name
	desc = item.Name
	

	if width <= 0 {
		// short-circuit
		return
	}

	textwidth := width - s.NormalTitle.GetPaddingLeft() - s.NormalTitle.GetPaddingRight()
	title = ansi.Truncate(title, textwidth, ellipsis)
	if ic.ShowDescription {
		var lines []string
		for i, line := range strings.Split(desc, "\n") {
			if i >= ic.height-1 {
				break
			}
			lines = append(lines, ansi.Truncate(line, textwidth, ellipsis))
		}
		desc = strings.Join(lines, "\n")
	}

	if isSelected {
		title = "» " + title
		title = s.SelectedTitle.Render(title)
		desc = s.SelectedDesc.Render(desc)
	} else {
		title = s.NormalTitle.Render(title)
		desc = s.NormalDesc.Render(desc)
	}


	if ic.ShowDescription {
		fmt.Fprintf(w, "%s\n%s", title, desc) //nolint: errcheck
		return
	}
	fmt.Fprintf(w, "%s", title) //nolint: errcheck

}