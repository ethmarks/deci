package main

import (
	"fmt"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
)

const (
	version = "0.4.2"
)

func main() {
	// first arg is always the path of the program
	args := os.Args

	if len(args) > 2 {
		log.Fatal(errStyle.Render("Too many arguments!"))
	}

	filepath := "deci_out.txt"

	if len(args) == 2 {
		filepath = args[1]
	}

	lines, err := readFile(filepath)

	if err != nil {
		lines = make([]string, 1)
	}

	p := tea.NewProgram(initialModel(lines, filepath))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
