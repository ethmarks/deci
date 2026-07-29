package main

import (
	"unicode"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
)

func (m model) handleCursorMove(key string) model {
	linesHeight := len(*m.lines) - 1

	switch key {
	case "up":
		if m.cursorY > 0 {
			m.cursorY -= 1
			m.cursorX = m.getClampedCursorX()
		} else {
			m.cursorX = 0
		}
	case "down":
		if m.cursorY < linesHeight {
			m.cursorY += 1
			m.cursorX = m.getClampedCursorX()
		} else {
			m.cursorX = lipgloss.Width((*m.lines)[m.cursorY])
		}
	case "left":
		if m.cursorX > 0 {
			m.cursorX -= 1
			m.cursorPrefX = m.cursorX
		} else if m.cursorY > 0 {
			m.cursorY -= 1
			m.cursorPrefX = lipgloss.Width((*m.lines)[m.cursorY])
			m.cursorX = m.cursorPrefX
		}
	case "right":
		if m.cursorX < lipgloss.Width((*m.lines)[m.cursorY]) {
			m.cursorX += 1
			m.cursorPrefX = m.cursorX
		} else if m.cursorY < linesHeight {
			m.cursorY += 1
			m.cursorX = 0
			m.cursorPrefX = m.cursorX
		}
	}
	return m
}

func (m model) handleCtrlCursorMove(key string) model {
	line := []rune((*m.lines)[m.cursorY])
	lineLen := ansi.StringWidth((*m.lines)[m.cursorY])

	switch key {
	case "ctrl+left":
		if m.cursorX == 0 {
			return m.handleCursorMove("left")
		}

		for m.cursorX > 0 && isDelimiter(line[m.cursorX-1]) {
			m = m.handleCursorMove("left")
		}

		for m.cursorX > 0 && !isDelimiter(line[m.cursorX-1]) {
			m = m.handleCursorMove("left")
		}
	case "ctrl+right":
		if m.cursorX >= lineLen {
			return m.handleCursorMove("right")
		}

		for m.cursorX < lineLen && !isDelimiter(line[m.cursorX]) {
			m = m.handleCursorMove("right")
		}

		for m.cursorX < lineLen && isDelimiter(line[m.cursorX]) {
			m = m.handleCursorMove("right")
		}
	}

	return m
}

func (m model) handlePager(key string) model {
	switch key {
	case "up":
		if m.paneOffsetY > 0 {
			m.paneOffsetY -= 1
		}
		if m.status == "reached end of file" {
			m.status = ""
		}
	case "down":
		linesHeight := len(*m.lines) - 1

		if m.paneOffsetY < linesHeight-1 {
			m.paneOffsetY += 1
		} else {
			m.status = "reached end of file"
		}
	}
	return m
}

func (m model) updateOffsets() model {
	contentRows := m.termHeight - m.reservedFromTop - m.reservedFromBottom
	contentCols := m.termWidth - m.reservedFromLeft - m.reservedFromRight

	if contentRows < 1 || contentCols < 1 {
		return m
	}

	if m.cursorY < m.paneOffsetY {
		m.paneOffsetY = m.cursorY
	} else if m.cursorY >= m.paneOffsetY+contentRows {
		m.paneOffsetY = m.cursorY - contentRows + 1
	}

	if m.cursorX < m.paneOffsetX {
		m.paneOffsetX = m.cursorX
	} else if m.cursorX >= m.paneOffsetX+contentCols {
		m.paneOffsetX = m.cursorX - contentCols + 1
	}

	return m
}

func (m model) getClampedCursorX() int {
	clampedX := min(m.cursorPrefX, lipgloss.Width((*m.lines)[m.cursorY]))
	return clampedX
}

func cursorShapeToString(shape tea.CursorShape) string {
	switch shape {
	case 0:
		return "block"
	case 1:
		return "underline"
	case 2:
		return "bar"
	default:
		return "block"
	}
}

func isDelimiter(r rune) bool {
	return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r)
}
