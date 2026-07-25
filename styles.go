package main

import (
	"charm.land/lipgloss/v2"
)

var (
	baseStyle    = lipgloss.NewStyle()
	inverseStyle = lipgloss.NewStyle().
			Reverse(true)
	errStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Red).
			Reverse(true)

	cursorLineStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("0"))
	lineNumStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8")).
			Align(lipgloss.Right).
			PaddingRight(lineNumPadding)
)
