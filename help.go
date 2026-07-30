package main

import (
	"strconv"
	"strings"

	"charm.land/lipgloss/v2"
)

type hyperlink struct {
	text string
	url  string
}

const (
	dAscii = `
     888
     888
     888
 .d88888
d88" 888
888  888
Y88b 888
 "Y88888`

	eAscii = `
 .d88b.
d8P  Y8b
88888888
Y8b.
 "Y8888`
	cAscii = `
 .d8888b
d88P"
888
Y88b.
 "Y8888P`
	iAscii = `
d8b
Y8P

888
888
888
888
888`
)

var (
	subtitleText = `"your second-to-last next editor"`

	descText = "deci is a terminal text editor (like nano)."

	separator = lipgloss.NewStyle().
			SetString(strings.Repeat("─", 40)). // from lipgloss.RoundedBorder()
			Foreground(lipgloss.Color("240")).
			Bold(true).
			Padding(1, 0).
			String()

	linkText = "links:"

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

	bodyText = `
special keybinds:

ctrl+h: toggles this help screen.
ctrl+c: exits deci and returns to the shell.
ctrl+o: writes editor content out to the file.
ctrl+p: toggles the preview screen. It works best with Markdown content, but if
        you try to preview a Go file or something, it won't crash.
alt+c:  changes the cursor shape. The default is a solid block, but you can
        change it to vertical bar or an underline if you want.
alt+p:  changes the preview style. The styles, in order, are: dracula, dark,
        ascii, tokyo-night, light, notty, and pink.
`
)

func getHelpLines() []string {
	asciiLetters := []string{dAscii, eAscii, cAscii, iAscii}

	for i, letter := range asciiLetters {
		asciiLetters[i] = lipgloss.NewStyle().
			Foreground(lipgloss.Color(strconv.Itoa(i + 3))).
			PaddingLeft(1).
			Render(letter)
	}

	artText := lipgloss.JoinHorizontal(lipgloss.Bottom, asciiLetters...)

	art := lipgloss.NewStyle().
		Foreground(lipgloss.Cyan).
		Render(artText)

	subtitle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Foreground(lipgloss.Blue).
		Render(subtitleText)

	var links strings.Builder
	for _, link := range hyperlinks {
		text := lipgloss.NewStyle().
			Render(link.text)

		url := lipgloss.NewStyle().
			Foreground(lipgloss.Green).
			Hyperlink(link.url).
			Render(link.url)

		links.WriteString(text)
		links.WriteString(": ")
		links.WriteString(url)
		links.WriteString("\n")
	}

	s := art + "\n\n" +
		subtitle + "\n\n" +
		descText + "\n" + separator + "\n" +
		linkText + "\n\n" + links.String() + separator +
		bodyText + separator

	return strings.Split(s, "\n")[1:]
}
