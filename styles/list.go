package styles

import "github.com/charmbracelet/lipgloss"

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
