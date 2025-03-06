package styles

import "github.com/charmbracelet/lipgloss"

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
