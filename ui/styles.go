package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	bullet       = "•"
	ellipsis     = "…"
	cursor       = "→"
	tabIndicator = "•"
)

// Styles contains style definitions for the item component
type ListStyle struct {
	// List style
	NoItems lipgloss.Style
	// Styled characters.
	ActivePaginationDot   lipgloss.Style
	InactivePaginationDot lipgloss.Style
}

// Default list styles
func DefaultListStyle() ListStyle {
	return ListStyle{
		// List style
		NoItems: lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#909090", Dark: "#626262"}),

		// Styled characters.
		ActivePaginationDot:   lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#847A85", Dark: "#FFFFFF"}),
		InactivePaginationDot: lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#DDDADA", Dark: "#3C3C3C"}),
	}
}

type ItemStyle struct {
	// The selected and normal state.
	NormalTitle   lipgloss.Style
	SelectedTitle lipgloss.Style

	// Cursor Style
	SelectedCursor lipgloss.Style
	DisabledCursor lipgloss.Style

	// Methods Style
	GetMethod     lipgloss.Style
	PostMethod    lipgloss.Style
	PutMethod     lipgloss.Style
	PatchMethod   lipgloss.Style
	DeleteMethod  lipgloss.Style
	OptionsMethod lipgloss.Style
	HeadMethod    lipgloss.Style
	UnknowMethod  lipgloss.Style

	// Url Style
	UrlStyle lipgloss.Style
}

// Default list styles
func DefaultItemStyle() ItemStyle {
	return ItemStyle{
		// The selected and normal state.
		NormalTitle:   lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#aaaaaa"}).Padding(0, 1, 0, 1), //nolint:mnd
		SelectedTitle: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#ffffff"}).Padding(0, 1, 0, 1),

		// Cursor Style
		SelectedCursor: lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#750404"}).Padding(0, 0, 0, 1),

		// Methods Style
		GetMethod:     lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#68D696"}).Padding(0, 0, 0, 1),
		PostMethod:    lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#F5DB7B"}).Padding(0, 0, 0, 1),
		PutMethod:     lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#486385"}).Padding(0, 0, 0, 1),
		PatchMethod:   lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#BEA6DE"}).Padding(0, 0, 0, 1),
		DeleteMethod:  lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#BC4B47"}).Padding(0, 0, 0, 1),
		OptionsMethod: lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#D959AC"}).Padding(0, 0, 0, 1),
		HeadMethod:    lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#68D696"}).Padding(0, 0, 0, 1),
		UnknowMethod:  lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#999999"}).Padding(0, 0, 0, 1),

		// Url Style
		UrlStyle: lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#aaaaaa"}).Padding(0, 0, 0, 1), //nolint:mnd
	}
}

type PanelStyle struct {
	// Border Style
	BorderLeftStyle lipgloss.Style
}

func DefaultPanelStyle() PanelStyle {
	return PanelStyle{
		// Border Style
		BorderLeftStyle: lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, false, true).BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#444444"}),
	}
}

type TabsStyle struct {
	// The selected and normal state for horizontal tabs.
	NormalTitle   lipgloss.Style
	SelectedTitle lipgloss.Style
	// The selected and normal state for vertical tabs.
	NormalBorderTitle   lipgloss.Style
	SelectedBorderTitle lipgloss.Style
	// Cursor Style
	SelectedCursor lipgloss.Style
}

func DefaultTabsStyle() TabsStyle {
	return TabsStyle{
		// The selected and normal state.
		NormalTitle:   lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#aaaaaa"}).Padding(0, 0, 0, 1), //nolint:mnd
		SelectedTitle: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#ffffff"}).Padding(0, 0, 0, 1),
		// The selected and normal state.
		NormalBorderTitle:   lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#aaaaaa"}).Padding(0, 0, 0, 1), //nolint:mnd
		SelectedBorderTitle: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#ffffff"}).Padding(0, 0, 0, 1),
		// Cursor Style
		SelectedCursor: lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#750404"}).Padding(0, 0, 0, 1),
	}
}
