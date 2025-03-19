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

type Tabs struct {
	// Item
	item Item
	// Styles & Keymaps
	Styles styles.TabsStyle
	KeyMap handler.KeyMap
	// Tabs Index & Name
	index     int
	Sections  []Tab
	Separator string
	// Window Size
	Size handler.SizeSpec
}

func NewTabComponent(item Item, width, height int) Tabs {

	return Tabs{
		// Item
		item: item,
		// Styles & Keymaps
		Styles: styles.DefaultTabsStyle(),
		KeyMap: handler.DefaultKeyMap(),
		// Tabs Index & Name
		index:     0,
		Separator: "  ",
		// Window Size
		Size: handler.NewSizeSpec(width, height),
	}
}

func (t *Tabs) SetItem(item Item) {
	t.item = item
	t.SetBadges()
}

func (t *Tabs) SetBadges() {
	for i := range t.Sections {
		badge := ""
		switch t.Sections[i].Name {
		case "Headers":
			n := len(t.item.Request.Header)
			badge = strconv.Itoa(n)
		case "Parameters":
			n := len(t.item.Request.Url.Query)
			badge = strconv.Itoa(n)
		case "Body":
			badge = handler.Ternary(t.item.Request.Body.Mode != "", styles.Bullet, "-")
		}
		t.Sections[i].SetBadge(badge)
	}
}

func (t *Tabs) AddTab(tab Tab) {
	t.Sections = append(t.Sections, tab)
}

func (t *Tabs) AddDefaultTabs(sections ...string) {
	for _, s := range sections {
		t.AddTab(Tab{Name: s})
	}
	t.SetBadges()
}

func (t Tabs) CurrentTab() Tab {
	return t.Sections[t.index]
}

func (t Tabs) SectionFormat(tab Tab, isSelected bool, i int) string {
	var (
		leftBorder = handler.Ternary(i == 0, " ", t.Styles.NormalBorderTitle.Render("│ "))
		style      = t.Styles.NormalBorderTitle
		bstyle     = t.Styles.BadgeStyle
	)
	if isSelected {
		style = t.Styles.SelectedBorderTitle
		bstyle = t.Styles.BadgeSelectedStyle
	}
	return leftBorder + style.Render(tab.Name) + bstyle.Render(tab.Badge)
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
	t.index = handler.TernaryNumber(t.index == 0, (len(t.Sections) - 1), t.index-1)
	return func() tea.Msg {
		return handler.NewTabMoveMsg(t.index)
	}
}

func (t *Tabs) NextTab() tea.Cmd {
	t.index = handler.TernaryNumber(t.index == (len(t.Sections)-1), 0, t.index+1)
	return func() tea.Msg {
		return handler.NewTabMoveMsg(t.index)
	}
}

func (t Tabs) View() string {

	var b strings.Builder

	for i, tab := range t.Sections {
		fmt.Fprintf(&b, "%s", t.SectionFormat(tab, i == t.index, i))
	}
	fmt.Fprintf(&b, "\n")

	currenTab := t.CurrentTab()
	fmt.Fprintf(&b, "%s\n", currenTab.Render(t.item))

	fmt.Fprintf(&b, "\n")

	return b.String()
}
