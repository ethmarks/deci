package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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

type errMsg struct{ err error }
type statusMsg struct{ status string }

const (
	headerTextLeft  = "  deci 0.0.1"
	headerTextRight = "by @ethmarks  "
)

var (
	inverseStyle = lipgloss.NewStyle().Reverse(true)
)

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

func (m model) getClampedCursorX() int {
	cursorLine := m.lines[m.cursorY]
	clampedX := min(m.cursorPrefX, len(cursorLine))
	return clampedX
}

func (m model) getHeader() string {
	if len(headerTextLeft)+len(headerTextRight) >= m.termWidth {
		return headerTextLeft + "" + headerTextRight
	}
	return fmt.Sprintf("%s%*s", headerTextLeft, m.termWidth-len(headerTextLeft), headerTextRight)
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case errMsg:
		m.err = msg.err
		return m, nil

	case statusMsg:
		m.status = msg.status
		return m, nil

	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height

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
			if m.cursorY > 0 {
				m.cursorY -= 1
				m.cursorX = m.getClampedCursorX()
			}
			return m, nil
		case "down":
			if m.cursorY < len(m.lines)-1 {
				m.cursorY += 1
				m.cursorX = m.getClampedCursorX()
			}
			return m, nil
		case "left":
			if m.cursorPrefX > 0 {
				m.cursorX -= 1
				m.cursorPrefX = m.cursorX
			}
			return m, nil
		case "right":
			if m.cursorPrefX < len(m.lines[m.cursorY]) {
				m.cursorPrefX += 1
				m.cursorX = m.getClampedCursorX()
			}
			return m, nil

		case "backspace":
			if m.cursorX > 0 {
				updatedLine := backspaceAt(m.lines[m.cursorY], m.cursorX)
				m.lines[m.cursorY] = updatedLine
				m.cursorX -= 1
			}

			m.status = ""
			return m, nil

		case "enter":
			line := m.lines[m.cursorY]

			before := line[:m.cursorX]
			after := string(line[m.cursorX]) + line[m.cursorX+1:]

			m.lines[m.cursorY] = before

			m.lines = slices.Insert(m.lines, m.cursorY+1, after)

			m.cursorPrefX = 0
			m.cursorX = 0
			m.cursorY += 1

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

			updatedLine := insertAt(m.lines[m.cursorY], char, m.cursorX)
			m.lines[m.cursorY] = updatedLine

			m.cursorX += 1
			m.cursorPrefX = m.cursorX

			m.status = ""
			return m, nil
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func makeCharGrid(width, height int) [][]string {
	grid := make([][]string, height)

	for y, _ := range grid {
		row := make([]string, width)
		for x, _ := range row {
			row[x] = " "
		}
		grid[y] = row
	}

	return grid
}

func (m model) View() tea.View {
	if m.err != nil {
		return tea.NewView(fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err))
	}

	linesToDisplay := m.termHeight - m.reservedFromTop - m.reservedFromBottom
	colsToDisplay := m.termWidth - m.reservedFromLeft - m.reservedFromRight

	if linesToDisplay < 1 || colsToDisplay < 1 {
		return tea.NewView("")
	}

	grid := makeCharGrid(m.termWidth, m.termHeight)

	// content
	for lineIndex := range linesToDisplay {
		y := lineIndex + m.reservedFromTop

		if y >= m.termHeight || lineIndex >= len(m.lines) {
			break
		}

		line := m.lines[lineIndex]

		for charIndex := range colsToDisplay {
			x := charIndex + m.reservedFromLeft

			if x >= m.termWidth || charIndex >= len(line) {
				break
			}

			char := line[charIndex]

			grid[y][x] = string(char)
		}
	}

	// header
	headerY := 0
	header := m.getHeader()
	for x, char := range header {
		grid[headerY][x] = inverseStyle.Render(string(char))
	}

	// cursor
	absCursorY := m.cursorY + m.reservedFromTop
	absCursorX := m.cursorX + m.reservedFromLeft
	grid[absCursorY][absCursorX] = inverseStyle.Render(grid[absCursorY][absCursorX])

	// Send the UI for rendering
	outLines := make([]string, len(grid))
	for y, line := range grid {
		outLines[y] = strings.Join(line, "")
	}

	v := tea.NewView(strings.Join(outLines, "\n"))

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
