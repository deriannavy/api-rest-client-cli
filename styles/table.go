package styles

import "github.com/charmbracelet/lipgloss"

type TableStyle struct {
	// Headers Style
	HeaderStyle lipgloss.Style
}

func DefaultTableStyle() TableStyle {
	return TableStyle{
		// Border Style
		HeaderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#444444")),
	}
}
