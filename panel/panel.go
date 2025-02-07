package panel

import "github.com/deriannavy/api-rest-client-cli/application"

type Model struct {
	CurrentConfig application.Config
	Styles        Styles
}

func New(config application.Config) Model {

	styles := DefaultStyles()

	return Model{
		Styles:        styles,
		CurrentConfig: config,
	}
}

func (m Model) View() string {
	// centralStyle.Render(CurrentConfig)
	return "info"
}
