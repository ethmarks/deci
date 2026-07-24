package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	filename string
	lines    []string
	status   string
	err      error

	cursorY     int
	cursorX     int // column of the displayed caret
	cursorPrefX int // preferred column

	termWidth  int
	termHeight int

	reservedFromTop    int
	reservedFromBottom int
	reservedFromLeft   int
	reservedFromRight  int
}

func initialModel(lines []string, filename string, created bool) model {
	var status string

	if created {
		status = fmt.Sprintf("Created %v", filename)
	}

	return model{
		filename: filename,
		lines:    lines,
		status:   status,

		reservedFromTop:    1, // for the header
		reservedFromBottom: 2, // for the keybinds and status bar
		reservedFromLeft:   0, // will be updated for the line nums
		reservedFromRight:  0, // unused
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
