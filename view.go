package main

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"fmt"
	"strings"
)

var (
	inverseStyle = lipgloss.NewStyle().Reverse(true)
	errStyle     = lipgloss.NewStyle().Foreground(lipgloss.Red).Reverse(true)
)

func makeCharGrid(width, height int) [][]string {
	grid := make([][]string, height)

	for y, _ := range grid {
		row := make([]string, width)
		for x, _ := range row {
			row[x] = " "
		}
		grid[y] = row
	}

	return grid
}

func (m model) View() tea.View {
	if m.err != nil {
		return tea.NewView(fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err))
	}

	linesToDisplay := m.termHeight - m.reservedFromTop - m.reservedFromBottom
	colsToDisplay := m.termWidth - m.reservedFromLeft - m.reservedFromRight

	if linesToDisplay < 1 || colsToDisplay < 1 {
		return tea.NewView("")
	}

	grid := makeCharGrid(m.termWidth, m.termHeight)

	// content
	for lineIndex := range linesToDisplay {
		y := lineIndex + m.reservedFromTop

		if y >= m.termHeight || lineIndex >= len(m.lines) {
			break
		}

		line := m.lines[lineIndex]

		for charIndex := range colsToDisplay {
			x := charIndex + m.reservedFromLeft

			if x >= m.termWidth || charIndex >= len(line) {
				break
			}

			char := line[charIndex]

			grid[y][x] = string(char)
		}
	}

	// header
	header := getHeader(m.termWidth)

	// Send the UI for rendering
	outLines := make([]string, len(grid))
	for y, line := range grid {
		outLines[y] = strings.Join(line, "")
	}
	out := strings.Join(outLines, "\n")

	v := tea.NewView(header + out)

	// cursor
	absCursorY := m.cursorY + m.reservedFromTop
	absCursorX := m.cursorX + m.reservedFromLeft
	v.Cursor = &tea.Cursor{
		Position: tea.Position{
			X: absCursorX,
			Y: absCursorY,
		},
		Shape: tea.CursorBar,
		Blink: true,
	}

	v.AltScreen = true

	return v
}
