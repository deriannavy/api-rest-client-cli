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

	AppStyle = lipgloss.NewStyle().Padding(0, 0, 1, 0)
)

type model struct {
	keyMap   handler.KeyMap
	list     ui.List
	panel    ui.Panel
	response string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keyMap.MakeRequest):
			item := m.panel.GetItem()
			data := MakeRequest(item)
			m.response = data
		}
	case tea.WindowSizeMsg:
		h, v := AppStyle.GetFrameSize()
		listMaxWidth := 25 // Refac

		m.list.Size.SetSize(listMaxWidth-h, msg.Height-(v+3)) // remover el +3
		m.list.ItemComplement.Size.SetWidth(listMaxWidth)

		m.panel.Size.SetSize(msg.Width, msg.Height-v)
		m.panel.ItemComplement.Size.SetWidth(msg.Width)

	case handler.CursorMoveMsg:
		currentItem := Configuration.Items[msg.Index]
		m.panel.SetItem(currentItem)
	}

	var (
		cmds     []tea.Cmd
		cmdList  tea.Cmd
		cmdPanel tea.Cmd
	)

	m.list, cmdList = m.list.Update(msg)
	cmds = append(cmds, cmdList)

	m.panel, cmdPanel = m.panel.Update(msg)
	cmds = append(cmds, cmdPanel)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {

	return AppStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.list.View(),
			m.panel.View(),
			m.response,
		),
	)
}

func main() {

	app.LoadConfiguration(&Configuration)

	panel := ui.NewPanel(Configuration.Items[0], 1, 1)
	panel.Tabs.AddDefaultTabs("Parameters", "Headers", "Body")

	m := model{
		keyMap: keyMap,
		list:   ui.NewList(Configuration.Items, 1, 1),
		panel:  panel,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
