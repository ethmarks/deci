package main

import (
	"charm.land/lipgloss/v2"
	"strings"
)

type hyperlink struct {
	text string
	url  string
}

var (
	// https://patorjk.com/software/taag/#p=display&f=Colossal&t=deci&x=none&v=4&h=4&w=80&we=false
	artText = `
       888                   d8b
       888                   Y8P
       888
   .d88888  .d88b.   .d8888b 888
  d88" 888 d8P  Y8b d88P"    888
  888  888 88888888 888      888
  Y88b 888 Y8b.     Y88b.    888
   "Y88888  "Y8888   "Y8888P 888`

	subtitleText = `"your second-to-last next editor"`

	bodyText = `
deci is a text editor like nano.
`

	hyperlinks = []hyperlink{
		hyperlink{
			text: "source",
			url:  "https://github.com/ethmarks/deci/",
		},
		hyperlink{
			text: "site",
			url:  "https://ethmarks.github.io/deci/",
		},
	}
)

func getHelpLines() []string {
	art := lipgloss.NewStyle().Foreground(lipgloss.Cyan).Render(artText)

	subtitle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Foreground(lipgloss.Blue).
		Render(subtitleText)

	var links string
	for _, link := range hyperlinks {
		text := lipgloss.NewStyle().
			Render(link.text)

		url := lipgloss.NewStyle().
			Foreground(lipgloss.Green).
			Hyperlink(link.url).
			Render(link.url)
		links += text + ": " + url + "\n"
	}

	s := art + "\n\n" + subtitle + "\n" + bodyText + "\n" + links

	return strings.Split(s, "\n")[1:]
}
