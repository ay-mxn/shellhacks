package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	s := m.styles

	header := m.appBoundaryView("Security Quiz")

	// Form (left side)
	formView := m.form.View()

	// Status (right side)
	status := s.Status.Render(
		s.StatusHeader.Render("Current Lesson") + "\n" +
			m.lessons[m.currentLessonIndex].Title + "\n\n" +
			s.StatusHeader.Render("Current Topic") + "\n" +
			m.lessons[m.currentLessonIndex].Topics[m.currentTopicIndex].Title,
	)

	body := lipgloss.JoinHorizontal(lipgloss.Top, formView, status)

	progressBar := m.renderProgressBar()

	footer := s.Help.Render("q: quit • enter: submit")

	return s.Base.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			header,
			body,
			footer,
			progressBar,
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
