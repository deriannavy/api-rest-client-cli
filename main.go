package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	app "github.com/deriannavy/api-rest-client-cli/application"
	"github.com/deriannavy/api-rest-client-cli/handler"
	"github.com/deriannavy/api-rest-client-cli/ui"
)

var (
	Configuration app.Configuration

	keyMap = handler.DefaultKeyMap()

	AppStyle = lipgloss.NewStyle().Padding(1, 0, 0, 0)
)

type model struct {
	keyMap handler.KeyMap
	list   ui.List
	// panel panel.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keyMap.Quit) {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := AppStyle.GetFrameSize()
		m.list.Size.SetSize(msg.Width-h, msg.Height-v)
	}

	var (
		cmds    []tea.Cmd
		cmdList tea.Cmd
	)

	m.list, cmdList = m.list.Update(msg)
	cmds = append(cmds, cmdList)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {

	return AppStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.list.View(),
			// m.panel.View(),
		),
	)
}

func main() {

	app.LoadConfiguration(&Configuration)

	m := model{
		keyMap: keyMap,
		list:   ui.NewList(Configuration.Items, 1, 1),
		// panel: components.NewPanel(itemsConfig),
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
