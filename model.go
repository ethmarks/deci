package main

import (
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

	history      []historyState
	historyIndex int
}

func initialModel(editorLines []string, filename string) model {
	m := model{
		lines:       &editorLines,
		editorLines: editorLines,
		helpLines:   getHelpLines(),

		filename: filename,

		status: statusTextWelcome,
		screen: Editor,

		showNums:   true,
		rawContent: false,
		pagerMode:  false,

		reservedFromTop:    1, // for the header
		reservedFromBottom: 2, // for the keybinds and status bar
		reservedFromLeft:   0, // for the line nums
		reservedFromRight:  0, // unused
	}

	m = m.saveSnapshot()

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) resetViewport() model {
	m.paneOffsetY = 0
	m.paneOffsetX = 0
	m.cursorY = 0
	m.cursorX = 0
	m.cursorPrefX = m.cursorX

	return m
}
