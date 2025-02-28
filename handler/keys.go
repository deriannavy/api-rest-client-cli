package handler

import "github.com/charmbracelet/bubbles/key"

// KeyMap defines general keybindings
type KeyMap struct {
	// Keybindings used when browsing the list.
	CursorUp   key.Binding
	CursorDown key.Binding

	// The quit keybinding. This won't be caught when filtering.
	Quit key.Binding

	// The quit-no-matter-what keybinding. This will be caught when filtering.
	ForceQuit key.Binding
}

// DefaultKeyMap returns a default set of keybindings.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		// Browsing list.
		CursorUp: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		CursorDown: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),

		// Quitting.
		Quit: key.NewBinding(
			key.WithKeys("q", "esc"),
			key.WithHelp("q", "quit"),
		),
		ForceQuit: key.NewBinding(key.WithKeys("ctrl+c")),
	}
}
