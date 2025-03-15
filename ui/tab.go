package ui

type Tab struct {
	Name  string
	Badge string
}

func (t *Tab) SetBadge(badge string) {
	t.Badge = badge
}
