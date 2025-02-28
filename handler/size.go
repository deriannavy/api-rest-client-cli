package handler

type SizeSpec struct {
	height int
	width  int
}

func NewSizeSpec(width, height int) SizeSpec {
	return SizeSpec{
		height: height,
		width:  width,
	}
}

func (w *SizeSpec) SetSize(width, height int) {
	w.setSize(width, height)
}

func (ss *SizeSpec) setSize(width, height int) {
	ss.width = width
	ss.height = height
}

func (ss *SizeSpec) Height() int {
	return ss.height
}

func (ss *SizeSpec) Width() int {
	return ss.width
}
