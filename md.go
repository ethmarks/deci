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

func previewMd(editorLines []string, style glamourStyle) ([]string, error) {
	in := strings.Join(editorLines, "\n")

	out, err := glamour.Render(in, style.String())

	if err != nil {
		return nil, err
	}

	outLines := strings.Split(out, "\n")

	return outLines, nil
}
