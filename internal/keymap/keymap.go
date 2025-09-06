package keymap

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

var DefaultKeyMap = huh.KeyMap{
	Text: huh.TextKeyMap{
		Submit: key.NewBinding(
			key.WithKeys("alt+enter"),
			key.WithHelp("alt+enter", "submit"),
		),
		NewLine: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "newline"),
		),
	},
}
