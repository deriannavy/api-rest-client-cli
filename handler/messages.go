package handler

type CursorMoveMsg struct {
	Index int
}

type TabMoveMsg struct {
	Index int
}

func NewCursorMoveMsg(index int) CursorMoveMsg {
	return CursorMoveMsg{Index: index}
}

func NewTabMoveMsg(index int) TabMoveMsg {
	return TabMoveMsg{Index: index}
}
