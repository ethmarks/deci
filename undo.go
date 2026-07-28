package main

import "fmt"

type historyState struct {
	editorLines []string

	status string

	cursorY     int
	cursorX     int
	cursorPrefX int
}

func (m model) saveSnapshot() model {
	if m.historyIndex < len(m.history)-1 {
		m.history = m.history[:m.historyIndex+1]
	}

	snapshot := historyState{
		editorLines: deepCopyLines(m.editorLines),
		status:      m.status,
		cursorY:     m.cursorY,
		cursorX:     m.cursorX,
		cursorPrefX: m.cursorPrefX,
	}

	m.history = append(m.history, snapshot)
	m.historyIndex = len(m.history) - 1

	return m
}

func (m model) undo() model {
	if m.historyIndex <= 0 {
		return m
	}

	m.historyIndex--
	state := m.history[m.historyIndex]

	m.editorLines = deepCopyLines(state.editorLines)
	m.cursorX = state.cursorX
	m.cursorY = state.cursorY
	m.cursorPrefX = state.cursorPrefX

	m.status = fmt.Sprintf("undid '%v'", m.history[m.historyIndex+1].status)
	m = m.updateOffsets()

	return m
}

func (m model) redo() model {
	if m.historyIndex >= len(m.history)-1 {
		return m
	}

	m.historyIndex++
	state := m.history[m.historyIndex]

	m.editorLines = deepCopyLines(state.editorLines)
	m.cursorX = state.cursorX
	m.cursorY = state.cursorY
	m.cursorPrefX = state.cursorPrefX

	m.status = fmt.Sprintf("redid '%v'", state.status)
	m = m.updateOffsets()

	return m
}

func deepCopyLines(in []string) []string {
	cpy := make([]string, len(in))
	copy(cpy, in)
	return cpy
}
