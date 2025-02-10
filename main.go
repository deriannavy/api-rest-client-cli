package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	app "github.com/deriannavy/api-rest-client-cli/application"
	"github.com/deriannavy/api-rest-client-cli/components"
	"github.com/deriannavy/api-rest-client-cli/panel"
)

var (
	ListConfig app.ListConfiguration

	AppStyle = lipgloss.NewStyle().Padding(1, 2)
)

type model struct {
	list  list.Model
	panel panel.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// indx := m.list.Index()
	// selectedConfig := ListConfig.getConfigByIndex(indx)

	// selectedConfigView := fmt.Sprintf(
	// 	"%s \n%s %s",
	// 	selectedConfig.Name,
	// 	selectedConfig.Request.Method,
	// 	urlStyle.Render(selectedConfig.getUri()),
	// )
	// CurrentConfig = fmt.Sprintf("%v", selectedConfigView)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := AppStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var (
		cmds []tea.Cmd
		listCmd tea.Cmd
		panelCmd tea.Cmd
	)
	
	m.list, listCmd = m.list.Update(msg)
	cmds = append(cmds, listCmd)

	globalIndex := m.list.Index()
	m.panel.SetCurrentConfig(globalIndex)

	m.panel, panelCmd = m.panel.Update(msg)
	cmds = append(cmds, panelCmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {

	return AppStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.list.View(),
			m.panel.View(),
		),
	)
}

func main() {

	app.LoadConfig(&ListConfig)

	items := ListConfig.GetItemList()
	itemsConfig := ListConfig.Configurations

	m := model{
		list:  components.NewList(items),
		panel: components.NewPanel(itemsConfig),
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
