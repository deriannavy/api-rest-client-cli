package styles

import "github.com/charmbracelet/lipgloss"

type TabsStyle struct {
	// The selected and normal state for horizontal tabs.
	NormalTitle   lipgloss.Style
	SelectedTitle lipgloss.Style
	// The selected and normal state for vertical tabs.
	NormalBorderTitle   lipgloss.Style
	SelectedBorderTitle lipgloss.Style
	// Cursor Style
	SelectedCursor lipgloss.Style
	// Badge Style
	BadgeStyle lipgloss.Style
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
		// Badge Style
		BadgeStyle: lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#750404"}),
	}
}
