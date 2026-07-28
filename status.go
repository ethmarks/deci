package main

import (
	"charm.land/lipgloss/v2"
)

const (
	statusTextWelcome = "Welcome to deci!"
)

var (
	statusStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Black).
		Background(lipgloss.White).
		Align(lipgloss.Center).
		Padding(0, 2)
)

func getStatusBar(status string, termWidth int) string {
	if status == "" {
		return ""
	}

	if lipgloss.Width(status) >= termWidth {
		return status
	}

	content := statusStyle.Render(status)

	placed := lipgloss.PlaceHorizontal(termWidth, lipgloss.Center, content)

	return placed
}
