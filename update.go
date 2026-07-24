package main

import (
	tea "charm.land/bubbletea/v2"
	"slices"
)

type errMsg struct{ err error }
type statusMsg struct{ status string }

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

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "ctrl+x":
			return m, tea.Quit

		// These keys should write the lines
		case "ctrl+s", "ctrl+o":
			return m, writeFileCmd(m.filename, m.lines)

		case "up":
			if m.cursorY > 0 {
				m.cursorY -= 1
				m.cursorX = m.getClampedCursorX()
			}
			return m, nil
		case "down":
			if m.cursorY < len(m.lines)-1 {
				m.cursorY += 1
				m.cursorX = m.getClampedCursorX()
			}
			return m, nil
		case "left":
			if m.cursorPrefX > 0 {
				m.cursorX -= 1
				m.cursorPrefX = m.cursorX
			}
			return m, nil
		case "right":
			if m.cursorPrefX < len(m.lines[m.cursorY]) {
				m.cursorX += 1
				m.cursorPrefX = m.cursorX
			}
			return m, nil

		case "backspace":
			if m.cursorX > 0 {
				updatedLine := backspaceAt(m.lines[m.cursorY], m.cursorX-1)
				m.lines[m.cursorY] = updatedLine
				m.cursorX -= 1
			} else if m.cursorY > 0 {
				m.cursorX = len(m.lines[m.cursorY-1])

				m.lines[m.cursorY-1] = m.lines[m.cursorY-1] + m.lines[m.cursorY]
				m.lines = slices.Delete(m.lines, m.cursorY, m.cursorY+1)

				m.cursorY -= 1
			}

			m.cursorPrefX = m.cursorX
			m.status = ""

			return m, nil

		case "enter":
			line := m.lines[m.cursorY]

			before := line[:m.cursorX]
			after := ""

			if m.cursorX < len(line)-1 {
				after = string(line[m.cursorX]) + line[m.cursorX+1:]
			}

			m.lines[m.cursorY] = before

			if m.cursorY == len(m.lines)-1 {
				m.lines = append(m.lines, after)
			} else {
				m.lines = slices.Insert(m.lines, m.cursorY+1, after)
			}

			m.cursorY += 1
			m.cursorX = 0
			m.cursorPrefX = m.cursorX

			return m, nil

		// All other keys
		default:
			char := msg.String()

			if char == "space" {
				char = " "
			}

			// This rejects both modifiers and non-ASCII chars
			if len(char) != 1 {
				return m, nil
			}

			updatedLine := insertAt(m.lines[m.cursorY], char, m.cursorX)
			m.lines[m.cursorY] = updatedLine

			m.cursorX += 1
			m.cursorPrefX = m.cursorX

			m.status = ""
			return m, nil
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func insertAt(line string, char string, index int) string {
	before := line[:index]
	after := line[index:]

	return before + char + after
}

func overwriteAt(line string, char string, index int) string {
	if index < 0 || index >= len(line) {
		return line
	}

	before := line[:index]
	after := line[index+1:]

	return before + char + after
}

func backspaceAt(line string, index int) string {
	if index < 0 || index >= len(line) {
		return line
	}

	before := line[:index]
	after := line[index+1:]

	return before + after
}

func (m model) getClampedCursorX() int {
	cursorLine := m.lines[m.cursorY]
	clampedX := min(m.cursorPrefX, len(cursorLine))
	return clampedX
}
