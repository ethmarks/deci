package main

import (
	tea "charm.land/bubbletea/v2"
	"fmt"
	"slices"
	"strings"
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

	// Is it a key press?
	case tea.KeyPressMsg:
		return m.handleKeypress(msg)
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) handleKeypress(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	// For status messages.
	// Display is one-indexed, but the cursor pos is zero-indexed.
	lineNum := m.cursorY + 1
	colNum := m.cursorX + 1

	switch key := msg.String(); key {

	// These keys should exit the program.
	case "ctrl+c", "ctrl+x":
		return m, tea.Quit

	// These keys should write out.
	case "ctrl+s", "ctrl+o":
		return m, writeFileCmd(m.filename, m.lines)

	// These keys should move the cursor.
	case "up", "down", "left", "right":
		return m.handleCursorMove(key), nil

	case "backspace":
		if m.cursorX > 0 {
			m.status = fmt.Sprintf("removed %v:%v", lineNum, colNum-1)
			updatedLine := backspaceAt(m.lines[m.cursorY], m.cursorX)
			m.lines[m.cursorY] = updatedLine
			m.cursorX -= 1
		} else if m.cursorY > 0 {
			m.status = fmt.Sprintf("merged line %v with %v", lineNum, lineNum-1)
			m.cursorX = len(m.lines[m.cursorY-1])

			m.lines[m.cursorY-1] = m.lines[m.cursorY-1] + m.lines[m.cursorY]
			m.lines = slices.Delete(m.lines, m.cursorY, m.cursorY+1)

			m.cursorY -= 1
		}

		m.cursorPrefX = m.cursorX

		return m, nil

	case "delete":
		if m.cursorX < len(m.lines[m.cursorY]) {
			m.status = fmt.Sprintf("removed %v:%v", lineNum, colNum)
			updatedLine := deleteAt(m.lines[m.cursorY], m.cursorX)
			m.lines[m.cursorY] = updatedLine
		} else if m.cursorY < len(m.lines)-1 {
			m.status = fmt.Sprintf("merged line %v with %v", lineNum+1, lineNum)
			m.lines[m.cursorY] = m.lines[m.cursorY] + m.lines[m.cursorY+1]
			m.lines = slices.Delete(m.lines, m.cursorY+1, m.cursorY+2)
		}

		m.cursorPrefX = m.cursorX

		return m, nil

	case "enter":
		line := m.lines[m.cursorY]

		before := line[:m.cursorX]
		after := ""

		if m.cursorX < len(line)-1 {
			after = string(line[m.cursorX]) + line[m.cursorX+1:]
		}

		if strings.TrimSpace(before) == "" {
			m.status = fmt.Sprintf("created new line at %v", lineNum)
		} else if strings.TrimSpace(after) == "" {
			m.status = fmt.Sprintf("created new line at %v", lineNum+1)
		} else {
			m.status = fmt.Sprintf("split line %v to line %v", lineNum, lineNum+1)
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

	case "space":
		m.lines[m.cursorY] = insertAt(m.lines[m.cursorY], " ", m.cursorX)
		m.cursorX += 1
		m.cursorPrefX = m.cursorX
		m.status = fmt.Sprintf("inserted space at %v:%v", lineNum, colNum)

		return m, nil

	case "tab":
		for range spacesPerTab {
			m.lines[m.cursorY] = insertAt(m.lines[m.cursorY], " ", m.cursorX)
			m.cursorX += 1
		}
		m.cursorPrefX = m.cursorX
		m.status = fmt.Sprintf("inserted tab at %v:%v", lineNum, colNum)

		return m, nil

	// All other keys
	default:
		// This rejects both modifiers and non-ASCII chars
		if len(key) != 1 {
			return m, nil
		}

		m.status = fmt.Sprintf("inserted '%v' at %v:%v", key, lineNum, colNum)

		m.lines[m.cursorY] = insertAt(m.lines[m.cursorY], key, m.cursorX)

		m.cursorX += 1
		m.cursorPrefX = m.cursorX

		return m, nil
	}
}

func (m model) handleCursorMove(key string) model {
	switch key {
	case "up":
		if m.cursorY > 0 {
			m.cursorY -= 1
			m.cursorX = m.getClampedCursorX()
		}
	case "down":
		if m.cursorY < len(m.lines)-1 {
			m.cursorY += 1
			m.cursorX = m.getClampedCursorX()
		}
	case "left":
		if m.cursorPrefX > 0 {
			m.cursorX -= 1
			m.cursorPrefX = m.cursorX
		} else if m.cursorY > 0 {
			m.cursorY -= 1
			m.cursorPrefX = len(m.lines[m.cursorY])
			m.cursorX = m.cursorPrefX
		}
	case "right":
		if m.cursorPrefX < len(m.lines[m.cursorY]) {
			m.cursorX += 1
			m.cursorPrefX = m.cursorX
		} else if m.cursorY < len(m.lines)-1 {
			m.cursorY += 1
			m.cursorX = 0
			m.cursorPrefX = m.cursorX
		}
	}
	return m
}

func insertAt(line string, char string, index int) string {
	if index < 0 || index > len(line) {
		return line
	}
	return line[:index] + char + line[index:]
}

func overwriteAt(line string, char string, index int) string {
	if index < 0 || index >= len(line) {
		return line
	}
	return line[:index] + char + line[index+1:]
}

func backspaceAt(line string, index int) string {
	if index > len(line) {
		return line
	}
	return line[:index-1] + line[index:]
}

func deleteAt(line string, index int) string {
	if index >= len(line) {
		return line
	}

	return line[:index] + line[index+1:]
}

func (m model) getClampedCursorX() int {
	cursorLine := m.lines[m.cursorY]
	clampedX := min(m.cursorPrefX, len(cursorLine))
	return clampedX
}
