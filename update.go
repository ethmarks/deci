package main

import (
	"math"

	tea "charm.land/bubbletea/v2"
)

type errMsg struct{ err error }
type statusMsg struct{ status string }

const (
	spacesPerTab = 4
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case errMsg:
		m.err = msg.err
		return m, nil

	case statusMsg:
		m.status = msg.status
		return m, nil

	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height
		m.reservedFromLeft = m.getLeftReserve()
		m = m.updateOffsets()

	case tea.KeyPressMsg:
		switch key := msg.String(); key {

		// These keys should exit the program.
		case "ctrl+c", "ctrl+x":
			return m, tea.Quit

		// These keys should write out.
		case "ctrl+s", "ctrl+o":
			return m, writeFileCmd(m.filename, m.editorLines)

		case "up", "down", "left", "right":
			return m.handleCursorMove(key), nil

		case "ctrl+h":
			m.cursorY = 0
			m.cursorX = 0
			m.cursorPrefX = m.cursorX

			if m.screen == Help {
				m.screen = Editor
				m.status = "switched to editor"
			} else {
				m.screen = Help
				m.status = "switched to help"
			}

		case "alt+c":
			m.cursorShape = (m.cursorShape + 2) % 3
			return m, nil

		default:
			if m.screen == Editor {
				m = m.handleEditorKeypress(key)

				m = m.updateOffsets()
			} else {
				m.status = "this screen is read-only"
			}
		}
	}

	switch m.screen {
	case Editor:
		m.lines = &m.editorLines
	case Help:
		m.lines = &m.helpLines
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) getLeftReserve() int {
	if !m.showNums {
		return 0
	}

	lineCount := len(*m.lines)

	magn := max(math.Log10(float64(lineCount)), minLineNumMagn)

	return int(math.Ceil(magn)) + lineNumPadding
}
