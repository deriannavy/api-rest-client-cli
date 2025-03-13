package styles

import "github.com/charmbracelet/lipgloss"

type TableStyle struct {
	// Header Style
	HeaderStyle lipgloss.Style
	// Rows Style
	RowOddStyle  lipgloss.Style
	RowEvenStyle lipgloss.Style
}

func DefaultTableStyle() TableStyle {
	return TableStyle{
		// Header Style
		HeaderStyle: lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, true, false).BorderForeground(lipgloss.Color("#444444")).Padding(0, 1).Bold(true).Foreground(lipgloss.Color("240")),
		// Border Style
		RowOddStyle:  lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, true, false).BorderForeground(lipgloss.Color("#444444")).Padding(0, 1).Foreground(lipgloss.Color("245")),
		RowEvenStyle: lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, true, false).BorderForeground(lipgloss.Color("#444444")).Padding(0, 1).Foreground(lipgloss.Color("240")),
	}
}
