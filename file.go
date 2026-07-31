package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/anyascii/go"
)

func readFile(filename string) ([]string, error) {
	lines := make([]string, 1)

	data, err := os.ReadFile(filename)

	if err == nil {
		s := string(data)

		s = strings.ReplaceAll(s, "\t", strings.Repeat(" ", spacesPerTab))

		s = anyascii.Transliterate(s)

		lines = strings.Split(s, "\n")
	}

	return lines, err
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
