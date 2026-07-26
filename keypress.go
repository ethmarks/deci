package main

import (
	"fmt"
	"slices"
	"strings"
)

func (m model) handleEditorKeypress(key string) model {
	// For status messages.
	// Display is one-indexed, but the cursor pos is zero-indexed.
	lineNum := m.cursorY + 1
	colNum := m.cursorX + 1

	switch key {

	case "up", "down", "left", "right":
		return m.handleCursorMove(key)

	case "backspace":
		return m.handleBackspace()

	case "delete":
		return m.handleDelete()

	case "enter":
		return m.handleEnter()

	case "space":
		m.lines[m.cursorY] = insertAt(m.lines[m.cursorY], " ", m.cursorX)
		m.cursorX += 1
		m.cursorPrefX = m.cursorX
		m.status = fmt.Sprintf("inserted space at %v:%v", lineNum, colNum)

		return m

	case "tab":
		for range spacesPerTab {
			m.lines[m.cursorY] = insertAt(m.lines[m.cursorY], " ", m.cursorX)
			m.cursorX += 1
		}
		m.cursorPrefX = m.cursorX
		m.status = fmt.Sprintf("inserted tab at %v:%v", lineNum, colNum)

		return m

	// All other keys
	default:
		// This rejects both modifiers and non-ASCII chars
		if len(key) != 1 {
			return m
		}

		m.status = fmt.Sprintf("inserted '%v' at %v:%v", key, lineNum, colNum)

		m.lines[m.cursorY] = insertAt(m.lines[m.cursorY], key, m.cursorX)

		m.cursorX += 1
		m.cursorPrefX = m.cursorX

		return m
	}
}

func (m model) handleBackspace() model {
	// For status messages.
	// Display is one-indexed, but the cursor pos is zero-indexed.
	lineNum := m.cursorY + 1
	colNum := m.cursorX + 1

	if m.cursorX > 0 {
		m.status = fmt.Sprintf("removed %v:%v", lineNum, colNum-1)
		updatedLine := backspaceAt(m.lines[m.cursorY], m.cursorX)
		m.lines[m.cursorY] = updatedLine
		m.cursorX -= 1
	} else if m.cursorY > 0 {
		if strings.TrimSpace(m.lines[m.cursorY]) == "" {
			m.status = fmt.Sprintf("removed line %v", lineNum)
		} else {
			m.status = fmt.Sprintf("merged line %v with %v", lineNum, lineNum-1)
		}

		m.cursorX = len(m.lines[m.cursorY-1])

		m.lines[m.cursorY-1] = m.lines[m.cursorY-1] + m.lines[m.cursorY]
		m.lines = slices.Delete(m.lines, m.cursorY, m.cursorY+1)

		m.cursorY -= 1
	}

	m.cursorPrefX = m.cursorX

	return m
}

func (m model) handleDelete() model {
	// For status messages.
	// Display is one-indexed, but the cursor pos is zero-indexed.
	lineNum := m.cursorY + 1
	colNum := m.cursorX + 1

	if m.cursorX < len(m.lines[m.cursorY]) {
		m.status = fmt.Sprintf("removed %v:%v", lineNum, colNum)
		updatedLine := deleteAt(m.lines[m.cursorY], m.cursorX)
		m.lines[m.cursorY] = updatedLine
	} else if m.cursorY < len(m.lines)-1 {
		if strings.TrimSpace(m.lines[m.cursorY]) == "" {
			m.status = fmt.Sprintf("removed line %v", lineNum)
		} else {
			m.status = fmt.Sprintf("merged line %v with %v", lineNum+1, lineNum)
		}

		m.lines[m.cursorY] = m.lines[m.cursorY] + m.lines[m.cursorY+1]
		m.lines = slices.Delete(m.lines, m.cursorY+1, m.cursorY+2)
	}

	m.cursorPrefX = m.cursorX

	return m
}

func (m model) handleEnter() model {
	// For status messages.
	// Display is one-indexed, but the cursor pos is zero-indexed.
	lineNum := m.cursorY + 1

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

	m.reservedFromLeft = m.getLeftReserve()

	m.cursorY += 1
	m.cursorX = 0
	m.cursorPrefX = m.cursorX

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
