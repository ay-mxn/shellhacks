package ui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	s := m.styles

	// Render the progress bar at the top
	progressBar := m.renderProgressBar()

	header := m.appBoundaryView("Security Quiz")

	formView := m.form.View()

	status := s.Status.Render(
		s.StatusHeader.Render("Current Lesson") + "\n" +
			m.lessons[m.currentLessonIndex].Title + "\n\n" +
			s.StatusHeader.Render("Current Topic") + "\n" +
			m.lessons[m.currentLessonIndex].Topics[m.currentTopicIndex].Title,
	)

	body := lipgloss.JoinHorizontal(lipgloss.Top, formView, status)

	footer := m.appBoundaryView(m.form.Help().ShortHelpView(m.form.KeyBinds()))

	// Add margins to avoid missing elements
	return s.Base.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			progressBar, // Progress bar now at the top
			header,      // Header below the progress bar
			body,        // Main content with form and status
			footer,      // Footer for navigation help
		),
	)
}

func (m Model) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(indigo),
	)
}

func (m Model) renderProgressBar() string {
	return m.styles.ProgressBar.Render(m.progress.View())
}
