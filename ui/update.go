package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

const maxWidth = 100

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
		m.progress.Width = m.width - 4
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit // Handle quitting
		case "enter":
			// If the form is completed, move to the next topic
			if m.form.State == huh.StateCompleted {
				return m.moveToNextTopic(), nil
			}
		}
	}

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f

		// Check if form is completed
		if m.form.State == huh.StateCompleted {
			// Handle restart logic if the quiz is completed
			if m.quizCompleted {
				if m.form.GetString("confirm") == "yes" {
					return m.restart(), nil
				}
				return m, tea.Quit
			}
			// Move to the next topic
			return m.moveToNextTopic(), nil
		}
	}

	return m, cmd
}
