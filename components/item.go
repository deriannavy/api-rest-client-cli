package components

type Item struct {
	id          string
	title       string
	description string
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.description }
func (i Item) FilterValue() string { return i.title + i.description }

func NewItem(id, title, description string) Item {
	return Item{id, title, description}
}
