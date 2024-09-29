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

		if m.state == stateAllCompleted {
			return m.renderAllCompleted()
		}

    progress := s.ProgressBar.Render(m.progress.View())

    var title string
    if len(m.lessons) > m.currentLesson && len(m.lessons[m.currentLesson].Topics) > m.currentTopic {
        title = s.Title.Render(fmt.Sprintf("%s - %s",
            m.lessons[m.currentLesson].Title,
            m.lessons[m.currentLesson].Topics[m.currentTopic].Title))
    } else {
        title = s.Title.Render("No lesson or topic available")
    }

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
	intro := m.styles.Intro.Render("ShellHacked 2024 - a Digital Literacy Lesson\n\n" +
		"Developed by Team 0x641a\n\n" +
		"Learn about online security through interactive lessons!\n\n" +
		"Press Enter to start :)")
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
	// bottomPadding := remainingSpace - topPadding

	view := strings.Repeat("\n", topPadding) +
		challenge + "\n\n" +
		input +
		message

	return lipgloss.NewStyle().
		Width(m.viewport.Width).
		Height(m.viewport.Height).
		Render(view)
}

func (m Model) renderFooter() string {
	var footerText string
	switch m.state {
	case stateIntro:
		footerText = "Press Enter to start • Q to quit"
	case stateContent:
		footerText = "← → to navigate between lessons • ↑↓ to scroll • Q to quit"
	case stateChallenge:
		footerText = "Enter to submit • ← to go back • Q to quit"
	case stateAllCompleted:
		footerText = "Press Enter to exit"
	}
	return m.styles.FooterText.Render(footerText)
}

func (m Model) renderAllCompleted() string {
	message := m.styles.Intro.Render("Congratulations! You've completed all the lessons.\n\n" +
		"And the biggest lesson is: don't run any binaries like this!\n\n" +
		"Press Enter to exit.")
	return lipgloss.Place(m.windowWidth, m.windowHeight,
		lipgloss.Center, lipgloss.Center,
		message)
}
