package handler

type CursorMoveMsg struct {
	Index int
}

func NewCursorMoveMsg(index int) CursorMoveMsg {
	return CursorMoveMsg{Index: index}
}
