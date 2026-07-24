package main

import (
	"fmt"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	// first arg is always the path of the program
	args := os.Args

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
