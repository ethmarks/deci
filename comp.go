package main

import (
	"strings"

	"charm.land/lipgloss/v2"
)

const (
	headerTextLeft  = "deci 0.0.1"
	headerTextRight = "by @ethmarks"
	headerPadding   = 2

	statusTextWelcome = "Welcome to deci!"
)

var (
	headerLeftStyle  = lipgloss.NewStyle().PaddingLeft(headerPadding)
	headerRightStyle = lipgloss.NewStyle().PaddingRight(headerPadding)
	headerStyle      = lipgloss.NewStyle().
				Foreground(lipgloss.Black).
				Background(lipgloss.White)

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Black).
			Background(lipgloss.White).
			Align(lipgloss.Center).
			Padding(0, 2)
)

func getHeader(termWidth int) string {
	left := headerLeftStyle.Render(headerTextLeft)
	right := headerRightStyle.Render(headerTextRight)

	leftWidth := lipgloss.Width(left)
	rightWidth := lipgloss.Width(right)

	spacerWidth := termWidth - leftWidth - rightWidth
	if spacerWidth < 0 {
		spacerWidth = 0
	}

	spacer := strings.Repeat(" ", spacerWidth)

	raw := lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		spacer,
		right,
	)

	return headerStyle.Width(termWidth).MaxWidth(termWidth).Render(raw)
}

func getStatusBar(status string, termWidth int) string {
	if lipgloss.Width(status) >= termWidth {
		return status
	}

	content := statusStyle.Render(status)

	placed := lipgloss.PlaceHorizontal(termWidth, lipgloss.Center, content)

	return placed
}
