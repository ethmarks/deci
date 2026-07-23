package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type pos struct {
	x int
	y int
}

type model struct {
	lines  []string
	cursor pos
	err    error
}

type errMsg struct{ err error }
type readFileMsg struct{ lines []string }

func initialModel() model {
	defaultLines := make([]string, 0)

	for i := range defaultLines {
		defaultLines[i] = "ahoy, world!"
	}

	return model{
		lines:  defaultLines,
		cursor: pos{},
		err:    nil,
	}
}

func readFileCmd(filename string) tea.Cmd {
	return func() tea.Msg {
		data, err := os.ReadFile(filename)
		if err != nil {
			return errMsg{err}
		}

		lines := strings.Split(string(data), "\n")

		return readFileMsg{lines}
	}
}

func insertAt(line string, char string, index int) string {
	before := line[:index]
	after := line[index:]

	return before + char + after
}

func (m model) Init() tea.Cmd {
	return readFileCmd("test.txt")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case readFileMsg:
		m.lines = msg.lines
		return m, nil

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

	case errMsg:
		m.err = msg.err
		return m, tea.Quit
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() tea.View {
	if m.err != nil {
		return tea.NewView(fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err))
	}

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
