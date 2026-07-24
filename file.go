package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
)

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
