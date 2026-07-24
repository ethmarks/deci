package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
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

	absCursorY := m.cursorY + m.reservedFromTop
	absCursorX := m.cursorX + m.reservedFromLeft

	if linesToDisplay < 1 || colsToDisplay < 1 {
		return tea.NewView("")
	}

	grid := makeCharGrid(colsToDisplay, linesToDisplay)

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

	// status bar
	statusBar := getStatusBar(m.status, m.termWidth)

	// Send the UI for rendering
	outLines := make([]string, len(grid))
	for y, chars := range grid {
		line := strings.Join(chars, "")
		if absCursorY == y {
			line = cursorLineStyle.Render(line)
		}
		outLines[y] = baseStyle.Render(line)
	}
	out := strings.Join(outLines, "\n")

	v := tea.NewView(header + out + "\n" + statusBar)

	// cursor
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
