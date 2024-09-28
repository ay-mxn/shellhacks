package main

import (
	"log"
	"os"
	"github.com/ay-mxn/shellhacks/internal"  // Import your internal package here

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	p := tea.NewProgram(internal.NewModel(), tea.WithAltScreen())  // Use internal.NewModel()
	if err := p.Start(); err != nil {
		log.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
