package main

import (
	"strings"

	"charm.land/glamour/v2"
)

type glamourStyle int

const (
	glamourDracula glamourStyle = iota
	glamourDark
	glamourAscii
	glamourTokyoNight
	glamourLight
	glamourNotty
	glamourPink
)

func (style glamourStyle) String() string {
	switch style {
	case glamourDracula:
		return "dracula"
	case glamourDark:
		return "dark"
	case glamourAscii:
		return "ascii"
	case glamourTokyoNight:
		return "tokyo-night"
	case glamourLight:
		return "light"
	case glamourNotty:
		return "notty"
	case glamourPink:
		return "pink"
	default:
		return "dracula"
	}
}

func previewMd(editorLines []string, style glamourStyle, termWidth int) ([]string, error) {
	in := strings.Join(editorLines, "\n")

	r, _ := glamour.NewTermRenderer(
		glamour.WithWordWrap(min(termWidth, 80)),
		glamour.WithStylePath(style.String()),
	)

	out, err := r.Render(in)

	if err != nil {
		return nil, err
	}

	outLines := strings.Split(out, "\n")

	return outLines, nil
}
