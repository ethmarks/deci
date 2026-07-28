package main

import (
	"charm.land/lipgloss/v2"
)

type keybind struct {
	bind      string
	text      string
	condition func(m model) bool
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
		bind:      "ctrl+o",
		text:      "Write Out",
		condition: func(m model) bool { return m.screen == Editor },
	},
	keybind{
		bind: "ctrl+p",
		text: "Toggle Preview",
		condition: func(m model) bool {
			return m.screen == Editor || m.screen == Preview
		},
	},
	keybind{
		bind:      "ctrl+z",
		text:      "Undo",
		condition: func(m model) bool { return m.screen == Editor },
	},
	keybind{
		bind:      "ctrl+y",
		text:      "Redo",
		condition: func(m model) bool { return m.screen == Editor },
	},
	keybind{
		bind:      "alt+c",
		text:      "Change Cursor Shape",
		condition: func(m model) bool { return !m.pagerMode },
	},
	keybind{
		bind:      "alt+p",
		text:      "Change Preview Style",
		condition: func(m model) bool { return m.screen == Preview },
	},
}

var (
	bindStyle     = lipgloss.NewStyle().Reverse(true)
	bindTextStyle = lipgloss.NewStyle().PaddingLeft(1).PaddingRight(3)
)

func (m model) getKeybindBar() string {
	bar := ""

	for _, kb := range keybinds {
		if kb.condition != nil && !kb.condition(m) {
			continue
		}

		bind := bindStyle.Render(kb.bind)
		text := bindTextStyle.Render(kb.text)

		s := lipgloss.JoinHorizontal(lipgloss.Left, bind, text)

		new := lipgloss.JoinHorizontal(lipgloss.Left, bar, s)

		if lipgloss.Width(new) < m.termWidth {
			bar = new
		} else {
			break
		}
	}

	return bar
}
