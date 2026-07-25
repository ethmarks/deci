package main

import (
	"fmt"
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const (
	lineNumPadding = 1
	minLineNumMagn = 3
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

	absCursorY := (m.cursorY - m.paneOffsetY) + m.reservedFromTop
	absCursorX := (m.cursorX - m.paneOffsetX) + m.reservedFromLeft

	if linesToDisplay < 1 || colsToDisplay < 1 {
		return tea.NewView("")
	}

	grid := makeCharGrid(colsToDisplay, linesToDisplay)

	// content
	for gridY := range linesToDisplay {
		absY := gridY + m.paneOffsetY

		if absY >= len(m.lines) {
			break
		}

		line := m.lines[absY]

		for gridX := range colsToDisplay {
			absX := gridX + m.paneOffsetX

			if absX >= len(line) {
				break
			}

			grid[gridY][gridX] = string(line[absX])
		}
	}

	// header
	header := getHeader(m.termWidth)

	// status bar
	statusBar := getStatusBar(m.status, m.termWidth)

	// Send the UI for rendering
	outLines := make([]string, len(grid))
	for gridY, chars := range grid {
		absY := gridY + m.paneOffsetY
		line := strings.Join(chars, "")
		lineNum := ""

		if absY < len(m.lines) {
			lineNum = strconv.Itoa(absY + 1)
		}

		lineStyle := baseStyle
		numStyle := lineNumStyle.Width(m.reservedFromLeft)

		if m.cursorY == absY {
			lineStyle = cursorLineStyle
			numStyle = numStyle.Foreground(lipgloss.White)
		}

		numStyle = numStyle.Inherit(lineStyle)
		outLines[gridY] = numStyle.Render(lineNum) + lineStyle.Render(line)
	}
	out := strings.Join(outLines, "\n")

	v := tea.NewView(header + "\n" + out + "\n" + statusBar)

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
