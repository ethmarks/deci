package main

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

const (
	headerTextRight = "by @ethmarks"
	headerPadding   = 2
)

var (
	headerTextLeft = fmt.Sprintf("deci %v", version)

	headerLeftStyle  = lipgloss.NewStyle().PaddingLeft(headerPadding)
	headerRightStyle = lipgloss.NewStyle().PaddingRight(headerPadding)
	headerStyle      = lipgloss.NewStyle().
				Foreground(lipgloss.Black).
				Background(lipgloss.White)
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
