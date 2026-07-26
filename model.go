package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type Mode int

const (
	Editor Mode = iota
	Help
)

type model struct {
	lines       *[]string
	editorLines []string

	filename string

	status string
	err    error
	mode   Mode

	showNums bool

	cursorY     int
	cursorX     int // column of the displayed caret
	cursorPrefX int // preferred column

	termWidth  int
	termHeight int

	reservedFromTop    int
	reservedFromBottom int
	reservedFromLeft   int
	reservedFromRight  int

	paneOffsetX int
	paneOffsetY int
}

func initialModel(editorLines []string, filename string, created bool) model {
	status := statusTextWelcome

	if created {
		status = fmt.Sprintf("Created %v", filename)
	}

	return model{
		editorLines: editorLines,
		lines:       &editorLines,

		filename: filename,

		status: status,
		mode:   Editor,

		showNums: true,

		reservedFromTop:    1, // for the header
		reservedFromBottom: 2, // for the keybinds and status bar
		reservedFromLeft:   0, // for the line nums
		reservedFromRight:  0, // unused
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
