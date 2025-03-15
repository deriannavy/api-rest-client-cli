package styles

import "github.com/charmbracelet/lipgloss"

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
