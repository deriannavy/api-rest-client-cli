package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	bullet   = "•"
	ellipsis = "…"
	cursor   = "→"
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
		ActivePaginationDot:   lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#847A85", Dark: "#979797"}).SetString(bullet),
		InactivePaginationDot: lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#DDDADA", Dark: "#3C3C3C"}).SetString(bullet),
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
		NormalTitle:   lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#aaaaaa"}).Padding(0, 0, 0, 1), //nolint:mnd
		SelectedTitle: lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#ffffff"}).Padding(0, 0, 0, 1),

		// Cursor Style
		SelectedCursor: lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Padding(0, 0, 0, 1),

		// Methods Style
		GetMethod:     lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#68D696"}).Align(lipgloss.Right).Padding(0, 0, 0, 1),
		PostMethod:    lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#F5DB7B"}).Align(lipgloss.Right).Padding(0, 0, 0, 1),
		PutMethod:     lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#486385"}).Align(lipgloss.Right).Padding(0, 0, 0, 1),
		PatchMethod:   lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#BEA6DE"}).Align(lipgloss.Right).Padding(0, 0, 0, 1),
		DeleteMethod:  lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#F2958D"}).Align(lipgloss.Right).Padding(0, 0, 0, 1),
		OptionsMethod: lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#D959AC"}).Align(lipgloss.Right).Padding(0, 0, 0, 1),
		HeadMethod:    lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#68D696"}).Align(lipgloss.Right).Padding(0, 0, 0, 1),
		UnknowMethod:  lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#999999"}).Align(lipgloss.Right).Padding(0, 0, 0, 1),

		// Url Style
		UrlStyle: lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#aaaaaa"}).Padding(0, 0, 0, 1), //nolint:mnd
	}
}

type PanelStyle struct {
}

func DefaultPanelStyle() PanelStyle {
	return PanelStyle{}
}
