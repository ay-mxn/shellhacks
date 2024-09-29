package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ay-mxn/shellhacks/internal"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Setup logging
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Printf("Failed to create log file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	// Create and start the Bubble Tea program
	p := tea.NewProgram(
		internal.NewModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(), // Enable mouse support for better interactivity
	)

	if _, err := p.Run(); err != nil {
		log.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}
