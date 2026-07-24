package main

import (
	"charm.land/lipgloss/v2"
)

var (
	inverseStyle = lipgloss.NewStyle().Reverse(true)
	errStyle     = lipgloss.NewStyle().Foreground(lipgloss.Red).Reverse(true)

	cursorLineStyle = lipgloss.NewStyle().Background(lipgloss.Color("0"))
)
