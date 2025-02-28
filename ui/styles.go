package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	bullet   = "•"
	ellipsis = "…"
)

// Styles contains style definitions for the item component
type ListStyle struct {
	// List style
	NoItems lipgloss.Style

	// Styled characters.
	ActivePaginationDot   lipgloss.Style
	InactivePaginationDot lipgloss.Style
}

type ItemStyle struct {
	// The Normal state.
	NormalTitle lipgloss.Style

	// The selected item state.
	SelectedTitle lipgloss.Style
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

// Default list styles
func DefaultItemStyle() ItemStyle {
	return ItemStyle{
		// The Normal state.
		NormalTitle: lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).Padding(0, 0, 0, 1), //nolint:mnd

		// The selected item state.
		SelectedTitle: lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, false, true).BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#990000"}).Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#990000"}).Padding(0, 0, 0, 0),
	}
}
