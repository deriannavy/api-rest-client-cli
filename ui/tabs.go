package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/deriannavy/api-rest-client-cli/handler"
	"github.com/deriannavy/api-rest-client-cli/styles"
)

type Tab struct {
	Name  string
	Badge int
}

type Tabs struct {
	// Styles & Keymaps
	TabType string
	Styles  styles.TabsStyle
	KeyMap  handler.KeyMap
	// Tabs Index & Name
	index     int
	Sections  []Tab
	Separator string
	// Window Size
	Size handler.SizeSpec
}

func NewTabComponent(TabType string, sections []string, width, height int) Tabs {

	var tabSection []Tab

	for i, s := range sections {
		tabSection = append(tabSection, Tab{s, i})
	}

	return Tabs{
		// Styles & Keymaps
		TabType: TabType,
		Styles:  styles.DefaultTabsStyle(),
		KeyMap:  handler.DefaultKeyMap(),
		// Tabs Index & Name
		index:     0,
		Sections:  tabSection,
		Separator: "  ",
		// Window Size
		Size: handler.NewSizeSpec(width, height),
	}
}

func (t Tabs) SectionFormat(tab Tab, isSelected bool) string {
	var (
		cursor = t.Styles.SelectedCursor.Render(" ")
		title  = t.Styles.NormalTitle.Render(tab.Name)
	)
	if isSelected {
		cursor = t.Styles.SelectedCursor.Render(styles.TabIndicator)
		title = t.Styles.SelectedTitle.Render(tab.Name)
	}
	return cursor + title
}

func (t Tabs) SectionBorderFormat(tab Tab, isSelected bool, i int) string {
	var (
		leftBorder  = handler.Ternary(i == 0, " ", t.Styles.NormalBorderTitle.Render("│ "))
		badgeNumber = strconv.Itoa(tab.Badge)
		style       = t.Styles.NormalBorderTitle
	)
	if isSelected {
		// badgeNumber = " "
		style = t.Styles.SelectedBorderTitle
	}
	return leftBorder + style.Render(tab.Name) + " " + t.Styles.BadgeStyle.Render(badgeNumber)
}

func (t Tabs) Update(msg tea.Msg) (Tabs, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, t.KeyMap.NextTab):
			cmds = append(cmds, t.NextTab())

		case key.Matches(msg, t.KeyMap.PrevTab):
			cmds = append(cmds, t.PrevTab())
		}
	}

	return t, tea.Batch(cmds...)
}

func (t *Tabs) PrevTab() tea.Cmd {
	if t.index == 0 {
		t.index = (len(t.Sections) - 1)
	} else {
		t.index--
	}
	return func() tea.Msg {
		return handler.NewTabMoveMsg(t.index)
	}
}

func (t *Tabs) NextTab() tea.Cmd {
	if t.index == (len(t.Sections) - 1) {
		t.index = 0
	} else {
		t.index++
	}
	return func() tea.Msg {
		return handler.NewTabMoveMsg(t.index)
	}
}

func (t Tabs) View() string {

	var b strings.Builder

	for i, tab := range t.Sections {
		if t.TabType == "Horizontal" {
			fmt.Fprintf(&b, "%s%s", t.SectionFormat(tab, i == t.index), t.Separator)
		} else if t.TabType == "Vertical" {
			fmt.Fprintf(&b, "%s", t.SectionBorderFormat(tab, i == t.index, i))
		}
		// if (i+1) == len(t.Sections) && t.TabType == "Vertical" {
		// 	fmt.Fprintf(&b, "%s", " │")
		// }
	}

	fmt.Fprintf(&b, "\n")

	return b.String()
}
