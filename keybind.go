package main

import (
	"charm.land/lipgloss/v2"
)

type keybind struct {
	bind string
	text string
}

var keybinds = []keybind{
	keybind{
		bind: "ctrl+h",
		text: "Toggle Help",
	},
	keybind{
		bind: "ctrl+c",
		text: "Exit",
	},
	keybind{
		bind: "ctrl+o",
		text: "Write Out",
	},
}

var (
	bindStyle     = lipgloss.NewStyle().Reverse(true)
	bindTextStyle = lipgloss.NewStyle().PaddingLeft(1).PaddingRight(3)
)

func getKeybindBar(termWidth int) string {
	bar := ""

	for _, kb := range keybinds {
		bind := bindStyle.Render(kb.bind)
		text := bindTextStyle.Render(kb.text)

		s := lipgloss.JoinHorizontal(lipgloss.Left, bind, text)

		new := lipgloss.JoinHorizontal(lipgloss.Left, bar, s)

		if lipgloss.Width(new) < termWidth {
			bar = new
		} else {
			break
		}
	}

	return bar
}
