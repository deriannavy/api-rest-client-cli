package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	ListConfig ListConfiguration

	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Background(lipgloss.Color("#25A065")).Padding(0, 1)

	urlStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("204")).Background(lipgloss.Color("235")).Padding(1, 2)

	CurrentConfig string
)

type item struct {
	id          string
	title       string
	description string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title + i.description }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	indx := m.list.Index()
	selectedConfig := ListConfig.getConfigByIndex(indx)
	// getUri
	selectedConfigView := fmt.Sprintf(
		"%s \n%s",
		selectedConfig.getName(),
		urlStyle.Render(selectedConfig.getUri()),
	)
	CurrentConfig = fmt.Sprintf("%v", selectedConfigView)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return appStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top, m.list.View(), CurrentConfig))
}

func main() {

	loadConfig(&ListConfig)

	var items []list.Item

	for _, lci := range ListConfig.Configurations {
		items = append(items, lci.toItem())
	}

	ls := list.NewDefaultDelegate()

	c := lipgloss.Color("#ff00ff")
	style := ls.Styles.SelectedTitle.Foreground(c).BorderLeftForeground(c)
	ls.Styles.SelectedTitle = style
	ls.Styles.SelectedDesc = style

	m := model{list: list.New(items, ls, 0, 0)}
	m.list.Title = "Requests"
	m.list.Styles.Title = titleStyle

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
