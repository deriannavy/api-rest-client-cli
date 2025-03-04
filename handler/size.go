package handler

type SizeSpec struct {
	// Attributes
	height int
	width  int
	// Used
	usedHeight int
	usedWidth  int
}

func NewSizeSpec(width, height int) SizeSpec {
	return SizeSpec{
		height:     height,
		width:      width,
		usedHeight: 0,
		usedWidth:  0,
	}
}

func (ss *SizeSpec) SetSize(width, height int) {
	ss.width = width
	ss.usedWidth = 0
	ss.height = height
	ss.usedHeight = 0
}

func (ss *SizeSpec) SetWidth(width int) {
	ss.width = width
	ss.usedWidth = width
}

func (ss *SizeSpec) Width() int {
	return ss.width
}

func (ss *SizeSpec) Height() int {
	return ss.height
}

func (ss *SizeSpec) AddUsedHeight(clear bool, add int) {
	if clear {
		ss.usedHeight = 0
	}
	ss.usedHeight += add
}

func (ss SizeSpec) UsedHeight() int {
	return ss.usedHeight
}

func (ss SizeSpec) AvailableHeight() int {
	availableHeight := ss.height - ss.usedHeight

	if availableHeight < 0 {
		availableHeight = 0
	}

	return availableHeight
}
