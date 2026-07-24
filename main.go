package main

import (
	"fmt"
	"log"
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
type statusMsg struct{ status string }

var cursor string = lipgloss.NewStyle().Reverse(true).Render(" ")

func initialModel(lines []string, filename string, created bool) model {
	var status string

	if created {
		status = fmt.Sprintf("Created %v", filename)
	}

	return model{
		filename: "test.txt",
		lines:    lines,
		cursor:   pos{},
		status:   status,
	}
}

func readFile(filename string) ([]string, error) {
	lines := make([]string, 1)

	data, err := os.ReadFile(filename)

	if err == nil {
		lines = strings.Split(string(data), "\n")
	}

	return lines, err
}

func createFile(filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	file.Close()
	return nil
}

func readOrCreateFile(filename string) ([]string, error, bool) {
	lines, err := readFile(filename)
	created := false

	if err != nil {
		// if we got an error reading it the first time, create it and try
		// reading the file one more time.
		err = createFile(filename)

		if err == nil {
			created = true
			lines, err = readFile(filename)
		}
	}

	return lines, err, created
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

func overwriteAt(line string, char string, index int) string {
	if index < 0 || index >= len(line) {
		return line
	}

	before := line[:index]
	after := line[index+1:]

	return before + char + after
}

func backspaceAt(line string, index int) string {
	if index < 0 || index >= len(line) {
		return line
	}

	before := line[:index]
	after := line[index+1:]

	return before + after
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

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
			if m.cursor.y < len(m.lines)-1 {
				m.cursor.y += 1
			}
			return m, nil
		case "left":
			if m.cursor.x > 0 {
				m.cursor.x -= 1
			}
			return m, nil
		case "right":
			if m.cursor.x < len(m.lines[m.cursor.y]) {
				m.cursor.x += 1
			}
			return m, nil

		case "backspace":
			if m.cursor.x > 0 {
				updatedLine := backspaceAt(m.lines[m.cursor.y], m.cursor.x)
				m.lines[m.cursor.y] = updatedLine
				m.cursor.x -= 1
			}

			m.status = ""
			return m, nil

		// All other keys
		default:
			char := msg.String()

			if char == "space" {
				char = " "
			}

			// This rejects both modifiers and non-ASCII chars
			if len(char) != 1 {
				return m, nil
			}

			updatedLine := insertAt(m.lines[m.cursor.y], char, m.cursor.x)
			m.lines[m.cursor.y] = updatedLine
			m.cursor.x += 1

			m.status = ""
			return m, nil
		}

	case statusMsg:
		m.status = msg.status
		return m, nil

	case errMsg:
		m.err = msg.err
		return m, nil
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
		output[m.cursor.y] = overwriteAt(output[m.cursor.y], cursor, m.cursor.x)
	}

	// Send the UI for rendering
	v := tea.NewView(strings.Join(output, "\n"))

	v.AltScreen = true

	return v
}

func main() {
	// first arg is always the path of the program
	args := os.Args

	errStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Red).
		Reverse(true)

	if len(args) == 1 {
		log.Fatal(errStyle.Render("You must provide a file as an argument"))
	} else if len(args) > 2 {
		log.Fatal(errStyle.Render("Too many arguments!"))
	}

	filepath := args[1]

	lines, err, created := readOrCreateFile(filepath)

	if err != nil {
		log.Fatal(errStyle.Render(err.Error()))
	}

	p := tea.NewProgram(initialModel(lines, filepath, created))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
