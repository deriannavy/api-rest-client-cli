package item

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
