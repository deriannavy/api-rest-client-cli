package panel

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	RequestName lipgloss.Style
	RequestUri  lipgloss.Style
}

func DefaultStyles() (s Styles) {

	s.RequestName = lipgloss.NewStyle()
	s.RequestUri = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("204")).Background(lipgloss.Color("235"))

	return s
}
