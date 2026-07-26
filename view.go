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

	grid = renderLines(grid, m.lines, linesToDisplay, colsToDisplay, m.paneOffsetY, m.paneOffsetX)

	header := getHeader(m.termWidth)

	statusBar := getStatusBar(m.status, m.termWidth)

	out := getOutString(
		grid,
		m.cursorY,
		m.paneOffsetY,
		m.reservedFromLeft,
		len(m.lines),
		m.showNums,
	)

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

func getOutString(
	grid [][]string,
	cursorY, paneOffsetY, reservedFromLeft, lineCount int,
	showNums bool,
) string {
	outLines := make([]string, len(grid))

	for gridY, chars := range grid {
		absY := gridY + paneOffsetY
		line := strings.Join(chars, "")

		lineStyle := baseStyle
		numStyle := lineNumStyle.Width(reservedFromLeft)

		if cursorY == absY {
			lineStyle = cursorLineStyle
			numStyle = numStyle.Foreground(lipgloss.White)
		}

		numStyle = numStyle.Inherit(lineStyle)

		lineNum := ""
		if showNums && absY < lineCount {
			lineNum = strconv.Itoa(absY + 1)
			lineNum = numStyle.Render(lineNum)
		}

		outLines[gridY] = lineNum + lineStyle.Render(line)
	}

	return strings.Join(outLines, "\n")
}

func renderLines(
	grid [][]string,
	lines []string,
	linesToDisplay, colsToDisplay, paneOffsetY, paneOffsetX int,
) [][]string {
	for gridY := range linesToDisplay {
		absY := gridY + paneOffsetY

		if absY >= len(lines) {
			break
		}

		line := lines[absY]

		for gridX := range colsToDisplay {
			absX := gridX + paneOffsetX

			if absX >= len(line) {
				break
			}

			grid[gridY][gridX] = string(line[absX])
		}
	}

	return grid
}
