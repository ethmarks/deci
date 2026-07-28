package main

import (
	"fmt"
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
)

const (
	lineNumPadding = 1
	minLineNumMagn = 3
)

func (m model) View() tea.View {
	if m.err != nil {
		return tea.NewView(fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err))
	}

	contentRows := m.termHeight - m.reservedFromTop - m.reservedFromBottom
	contentCols := m.termWidth - m.reservedFromLeft - m.reservedFromRight

	absCursorY := (m.cursorY - m.paneOffsetY) + m.reservedFromTop
	absCursorX := (m.cursorX - m.paneOffsetX) + m.reservedFromLeft

	if contentRows < 1 || contentCols < 1 {
		return tea.NewView("")
	}

	content := getContentString(
		*m.lines,
		contentRows, contentCols,
		m.paneOffsetY, m.paneOffsetX,

		m.cursorY, m.reservedFromLeft,

		m.showNums, m.rawContent, m.pagerMode,
	)

	header := getHeader(m.termWidth)
	statusBar := getStatusBar(m.status, m.termWidth)
	keybindBar := m.getKeybindBar()

	v := tea.NewView(header + "\n" + content + "\n" + statusBar + "\n" + keybindBar)

	// cursor
	if !m.pagerMode {
		v.Cursor = &tea.Cursor{
			Position: tea.Position{
				X: absCursorX,
				Y: absCursorY,
			},
			Shape: m.cursorShape,
			Blink: true,
		}
	}

	v.AltScreen = true

	return v
}

func getContentString(
	fullLines []string,
	contentRows, contentCols,
	paneOffsetY, paneOffsetX int,

	cursorY, reservedFromLeft int,

	showNums, rawContent, pagerMode bool,
) string {
	outLines := make([]string, contentRows)

	for row := range outLines {
		absY := row + paneOffsetY

		contentStyle := baseStyle.Width(contentCols)
		numStyle := lineNumStyle.Width(reservedFromLeft)

		if cursorY == absY && !pagerMode {
			contentStyle = contentStyle.
				Background(cursorLineBackground)
			numStyle = numStyle.
				Background(cursorLineBackground).
				Foreground(lipgloss.White)
		}

		var lineNum string
		if showNums && absY < len(fullLines) {
			lineNum = strconv.Itoa(absY + 1)
		}

		var content string
		if absY < len(fullLines) {
			content = ansi.Cut(fullLines[absY], paneOffsetX, contentCols+paneOffsetX)
		}

		if !rawContent {
			content = contentStyle.Render(content)
		} else {
			numStyle = numStyle.Padding(0)
		}

		outLines[row] = numStyle.Render(lineNum) + content
	}

	return strings.Join(outLines, "\n")
}
