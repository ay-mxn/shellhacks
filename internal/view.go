package internal

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	s := m.styles

	if m.state == stateIntro {
		return m.renderIntro()
	}

	progress := s.ProgressBar.Render(m.progress.View())
	title := s.Title.Render(fmt.Sprintf("%s - %s (Progress: %.0f%%)", 
		m.lessons[m.currentLesson].Title, 
		m.lessons[m.currentLesson].Topics[m.currentTopic].Title,
		m.progress.Percent()*100))
	header := lipgloss.JoinVertical(lipgloss.Left, title, progress)

	var content string
	if m.state == stateContent {
		content = m.viewport.View()
	} else {
		content = m.renderChallenge()
	}

	mainContent := s.Base.Render(content)

	footer := m.renderFooter()

	fullView := lipgloss.JoinVertical(lipgloss.Left,
		header,
		mainContent,
		footer,
	)

	return lipgloss.Place(m.windowWidth, m.windowHeight,
		lipgloss.Center, lipgloss.Center,
		fullView)
}

func (m Model) renderIntro() string {
	intro := m.styles.Intro.Render("Welcome to the Digital Security Literacy CLI!\n\n" +
		"Created by Room 641A\n\n" +
		"Learn about online security through interactive lessons\n\n" +
		"Press Enter to start your journey.")
	return lipgloss.Place(m.windowWidth, m.windowHeight,
		lipgloss.Center, lipgloss.Center,
		intro)
}

func (m Model) renderChallenge() string {
	challenge := m.lessons[m.currentLesson].Topics[m.currentTopic].Challenge
	input := m.textInput.View()
	message := m.challengeMsg

	challengeHeight := strings.Count(challenge, "\n") + 1
	inputHeight := 1
	messageHeight := strings.Count(message, "\n") + 1

	remainingSpace := m.viewport.Height - challengeHeight - inputHeight - messageHeight

	topPadding := remainingSpace / 2
	bottomPadding := remainingSpace - topPadding

	view := strings.Repeat("\n", topPadding) +
		challenge + "\n\n" +
		input + "\n" +
		message +
		strings.Repeat("\n", bottomPadding)

	return lipgloss.NewStyle().
		Width(m.viewport.Width).
		Height(m.viewport.Height).
		Render(view)
}

func (m Model) renderFooter() string {
	if m.state == stateContent {
		return m.styles.FooterText.Render("Press right arrow to continue • Up/Down to scroll • Q to quit")
	}
	return m.styles.FooterText.Render("Press Enter to submit • Q to quit")
}
