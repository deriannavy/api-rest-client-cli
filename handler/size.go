package handler

type SizeSpec struct {
	// Attributes
	height int
	width  int
	// Available
	availableHeight int
	availableWidth  int
}

func NewSizeSpec(width, height int) SizeSpec {
	return SizeSpec{
		height: height,
		width:  width,
	}
}

func (ss *SizeSpec) SetSize(width, height int) {
	ss.width = width
	ss.availableWidth = width
	ss.height = height
	ss.availableHeight = height
}

func (ss *SizeSpec) SetWidth(width int) {
	ss.width = width
	ss.availableWidth = width
}

func (ss *SizeSpec) Width() int {
	return ss.width
}

func (ss *SizeSpec) Height() int {
	return ss.height
}

func (ss *SizeSpec) SubstractAvailableHeight(clear bool, substract int) {
	if clear {
		ss.availableHeight = ss.Height()
	}
	ss.availableHeight -= ss.height
}
