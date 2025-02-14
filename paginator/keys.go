package paginator

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	// > Key binding to prev page
	PrevPage key.Binding
	// > Key binding to next page
	NextPage key.Binding
}

// DefaultKeyMap is the default set of key bindings for navigating and acting
// upon the paginator.
var DefaultKeyMap = KeyMap{
	// > Key binding to prev page
	PrevPage: key.NewBinding(key.WithKeys("pgup", "left", "h")),
	// > Key binding to prev page
	NextPage: key.NewBinding(key.WithKeys("pgdown", "right", "l")),
}
