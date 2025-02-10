package panel

import (
	"github.com/deriannavy/api-rest-client-cli/application"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	ItemsConfig []application.Config
	CurrentConfig application.Config
	Styles        Styles
}

func New(itemsConfig []application.Config) Model {

	styles := DefaultStyles()

	return Model{
		Styles:        styles,
		ItemsConfig: itemsConfig,
		CurrentConfig: itemsConfig[0],
	}
}


func (m Model) SetCurrentConfig(index int) {
	// centralStyle.Render(CurrentConfig)
	m.CurrentConfig = m.ItemsConfig[index]
}

func (m Model) View() string {
	// centralStyle.Render(CurrentConfig)
	return m.CurrentConfig.GetUri()
}




func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd

	// m.CurrentConfig = ListConfig.GetConfigByIndex(0)
	

	return m, tea.Batch(cmds...)
}