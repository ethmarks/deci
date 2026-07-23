package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type pos struct {
	x int
	y int
}

type model struct {
	filename string
	lines    []string
	cursor   pos
	status   string
	err      error
}

type errMsg struct{ err error }
type readFileMsg struct{ lines []string }
type statusMsg struct{ status string }

var cursor string = lipgloss.NewStyle().Reverse(true).Render(" ")

func initialModel() model {
	defaultLines := make([]string, 1)

	return model{
		filename: "test.txt",
		lines:    defaultLines,
		cursor:   pos{},
		status:   "",
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

func writeFileCmd(filename string, lines []string) tea.Cmd {
	return func() tea.Msg {
		data := []byte(strings.Join(lines, "\n"))

		// no idea what it means but its the one they use in the docs
		// https://pkg.go.dev/os#WriteFile
		perm := os.FileMode(0666)

		err := os.WriteFile(filename, data, perm)

		if err != nil {
			return errMsg{err}
		}

		return statusMsg{fmt.Sprintf("Wrote %v lines", len(lines))}
	}
}

func insertAt(line string, char string, index int) string {
	before := line[:index]
	after := line[index:]

	return before + char + after
}

func replaceAt(line string, char string, index int) string {
	if index < 0 || index >= len(line) {
		return line
	}

	before := line[:index]
	after := line[index+1:]

	return before + char + after
}

func (m model) Init() tea.Cmd {
	return readFileCmd(m.filename)
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
		case "ctrl+c", "ctrl+x":
			return m, tea.Quit

		// These keys should write the lines
		case "ctrl+s", "ctrl+o":
			return m, writeFileCmd(m.filename, m.lines)

		case "up":
			if m.cursor.y > 0 {
				m.cursor.y -= 1
			}
			return m, nil
		case "down":
			m.cursor.y += 1
			return m, nil
		case "left":
			if m.cursor.x > 0 {
				m.cursor.x -= 1
			}
			return m, nil
		case "right":
			m.cursor.x += 1
			return m, nil

		// All other keys
		default:
			m.status = ""
			return m, nil
		}

	case statusMsg:
		m.status = msg.status
		return m, nil

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

	// copy of m.lines
	output := append([]string(nil), m.lines...)

	if m.status != "" {
		if len(output) == 0 {
			return tea.NewView(m.status)
		}
		output[len(output)-1] = m.status
	}

	if len(output) > m.cursor.y {
		output[m.cursor.y] = replaceAt(output[m.cursor.y], cursor, m.cursor.x)
	}

	// Send the UI for rendering
	v := tea.NewView(strings.Join(output, "\n"))

	v.AltScreen = true

	return v
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
