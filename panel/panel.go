package panel

type Model struct {
	Styles Styles
}

func New() Model {

	styles := DefaultStyles()

	return Model{
		Styles: styles,
	}
}

func (m Model) View() string {
	return "info"
}
