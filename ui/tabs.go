package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/deriannavy/api-rest-client-cli/handler"
)

type Tabs struct {
	// Styles & Keymaps
	TabType string
	Styles  TabsStyle
	KeyMap  handler.KeyMap
	// Tabs Index & Name
	index     int
	Sections  []string
	Separator string
	// Window Size
	Size handler.SizeSpec
}

func NewTabComponent(TabType string, sections []string, width, height int) Tabs {
	return Tabs{
		// Styles & Keymaps
		TabType: TabType,
		Styles:  DefaultTabsStyle(),
		KeyMap:  handler.DefaultKeyMap(),
		// Tabs Index & Name
		index:     0,
		Sections:  sections,
		Separator: "  ",
		// Window Size
		Size: handler.NewSizeSpec(width, height),
	}
}

func (t Tabs) SectionFormat(section string, isSelected bool) string {
	var (
		cursor = t.Styles.SelectedCursor.Render(" ")
		title  = t.Styles.NormalTitle.Render(section)
	)
	if isSelected {
		cursor = t.Styles.SelectedCursor.Render(tabIndicator)
		title = t.Styles.SelectedTitle.Render(section)
	}
	return cursor + title
}

func (t Tabs) SectionBorderFormat(section string, isSelected bool) string {

	s := t.Styles.NormalBorderTitle
	if isSelected {
		s = t.Styles.SelectedBorderTitle
	}
	return t.Styles.NormalBorderTitle.Render("| ") + s.Render(section) + " "
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

	for i, s := range t.Sections {
		if t.TabType == "Horizontal" {
			fmt.Fprintf(&b, "%s%s", t.SectionFormat(s, i == t.index), t.Separator)
		} else if t.TabType == "Vertical" {
			fmt.Fprintf(&b, "%s", t.SectionBorderFormat(s, i == t.index))
		}
		if (i+1) == len(t.Sections) && t.TabType == "Vertical" {
			fmt.Fprintf(&b, "%s", " |")
		}
	}

	fmt.Fprintf(&b, "\n")

	return b.String()
}
