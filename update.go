package main

import (
	"fmt"
	"math"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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

	case tea.PasteMsg:
		if m.screen == Editor {
			for _, key := range msg.Content {
				m = m.handleEditorKeypress(string(key))
			}

			fits := lipgloss.Width(msg.Content) <= m.termWidth/2 && strings.Contains(msg.Content, "\n")

			if fits {
				m.status = fmt.Sprintf("pasted '%v'", msg.Content)
			} else {
				m.status = fmt.Sprintf("pasted %v characters", lipgloss.Width(msg.Content))
			}

			m = m.updateOffsets()
		} else {
			m.status = "this screen is read-only"
		}

	case tea.KeyPressMsg:
		switch key := msg.String(); key {

		// These keys should exit the program.
		case "ctrl+c", "ctrl+x":
			return m, tea.Quit

		// These keys should write out.
		case "ctrl+s", "ctrl+o":
			if m.screen == Editor {
				return m, writeFileCmd(m.filename, m.editorLines)
			} else {
				return m, nil
			}

		case "up", "down", "left", "right":
			m = m.handleCursorMove(key)
			m = m.updateOffsets()
			return m, nil

		case "ctrl+h":
			m.cursorY = 0
			m.cursorX = 0
			m.cursorPrefX = m.cursorX

			if m.screen != Help {
				m.screen = Help
				m.rawContent = true
				m.showNums = false
				m.status = "switched to help"
			} else {
				m.screen = Editor
				m.rawContent = false
				m.showNums = true
				m.status = "switched to editor"
			}
			m.reservedFromLeft = m.getLeftReserve()

		case "alt+c":
			m.cursorShape = (m.cursorShape + 2) % 3
			return m, nil

		case "ctrl+p":
			m.cursorY = 0
			m.cursorX = 0
			m.cursorPrefX = m.cursorX

			if m.screen != Preview {
				m.screen = Preview
				m.rawContent = true
				m.showNums = false
				m.status = "previewing as markdown"
			} else {
				m.screen = Editor
				m.rawContent = false
				m.showNums = true
				m.status = "switched to editor"
			}
			m.reservedFromLeft = m.getLeftReserve()

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
	case Preview:
		md, err := previewMd(*m.lines)

		if err != nil {
			m.err = err
			return m, nil
		}

		m.lines = &md

		return m, nil
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
