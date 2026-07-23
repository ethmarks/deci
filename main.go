package main

import (
	tea "charm.land/bubbletea/v2"
	"fmt"
	"os"
	"strings"
)

type pos struct {
	x int
	y int
}

type model struct {
	lines  []string
	cursor pos
}

func initialModel() model {
	defaultLines := make([]string, 5)

	for i := range defaultLines {
		defaultLines[i] = "ahoy, world!"
	}

	return model{
		lines:  defaultLines,
		cursor: pos{},
	}
}

func insertAt(line string, char string, index int) string {
	before := line[:index]
	after := line[index:]

	return before + char + after
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

			/**
			// The "up" and "k" keys move the cursor up
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}

			// The "down" and "j" keys move the cursor down
			case "down", "j":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}

			// The "enter" key and the space bar toggle the selected state
			// for the item that the cursor is pointing at.
			case "enter", "space":
				_, ok := m.selected[m.cursor]
				if ok {
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}
			*/
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() tea.View {
	s := strings.Join(m.lines, "\n")

	// Send the UI for rendering
	return tea.NewView(s)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
