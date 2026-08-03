package main

import (
	"fmt"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
)

const (
	version         = "1.0.0"
	defaultFilepath = "deci_out.txt"
)

func main() {
	// first arg is always the path of the program
	args := os.Args

	if len(args) > 2 {
		log.Fatal(errStyle.Render("Too many arguments!"))
	}

	filepath := defaultFilepath

	if len(args) == 2 {
		arg := args[1]

		if arg == "--version" || arg == "-v" {
			fmt.Printf("deci %v\n", version)
			os.Exit(0)
		}

		if arg == "--help" || arg == "-h" {
			fmt.Print(helpCliMsg)
			os.Exit(0)
		}

		filepath = arg
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
