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
	headerStyle      = lipgloss.NewStyle().Reverse(true)
)

func getCenterText(file string) string {
	if file == defaultFilepath {
		return ""
	}

	return file
}

func getHeader(centerText string, termWidth int) string {
	left := headerLeftStyle.Render(headerTextLeft)
	right := headerRightStyle.Render(headerTextRight)
	center := centerText

	leftWidth := lipgloss.Width(left)
	rightWidth := lipgloss.Width(right)
	centerWidth := lipgloss.Width(center)

	spacerWidth := termWidth - leftWidth - rightWidth - centerWidth
	if spacerWidth < 0 {
		spacerWidth = 0
	}

	halfSpacer := strings.Repeat(" ", spacerWidth/2)

	raw := lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		halfSpacer,
		center,
		halfSpacer,
		right,
	)

	return headerStyle.Width(termWidth).MaxWidth(termWidth).Render(raw)
}
