package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type screen int

const (
	Editor screen = iota
	Help
	Preview
)

type model struct {
	lines       *[]string
	editorLines []string
	helpLines   []string

	filename string

	status string
	err    error
	screen screen

	showNums   bool
	rawContent bool
	pagerMode  bool

	cursorY     int
	cursorX     int             // column of the displayed caret
	cursorPrefX int             // preferred column
	cursorShape tea.CursorShape // just an iota

	termWidth  int
	termHeight int

	reservedFromTop    int
	reservedFromBottom int
	reservedFromLeft   int
	reservedFromRight  int

	paneOffsetX int
	paneOffsetY int

	previewStyle glamourStyle
}

func initialModel(editorLines []string, filename string, created bool) model {
	status := statusTextWelcome

	if created {
		status = fmt.Sprintf("Created %v", filename)
	}

	return model{
		lines:       &editorLines,
		editorLines: editorLines,
		helpLines:   getHelpLines(),

		filename: filename,

		status: status,
		screen: Editor,

		showNums:   true,
		rawContent: false,
		pagerMode:  false,

		reservedFromTop:    1, // for the header
		reservedFromBottom: 2, // for the keybinds and status bar
		reservedFromLeft:   0, // for the line nums
		reservedFromRight:  0, // unused
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
