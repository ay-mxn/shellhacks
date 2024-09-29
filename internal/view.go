package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

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
type SummaryResponse struct {
	TotalHosts     int      `json:"total_hosts"`
	TotalMemoryGB  int      `json:"total_memory_gb"`
	UniqueIPCount  int      `json:"unique_ip_count"`
	UniqueOSCount  int      `json:"unique_os_count"`
	TotalCPUCores  int      `json:"total_cpu_cores"`
	AccessTypes    []string `json:"access_types"`
	OldestTimestamp string  `json:"oldest_timestamp"`
	NewestTimestamp string  `json:"newest_timestamp"`
}

func (m Model) renderError(err error) string {
	errorMessage := fmt.Sprintf("An error occurred: %v\n\nPress Enter to exit.", err)
	return lipgloss.Place(m.windowWidth, m.windowHeight,
		lipgloss.Center, lipgloss.Center,
		m.styles.Base.Render(errorMessage))
}


func (m Model) renderAllCompleted() string {
	summary, err := getSummaryOnce()
	if err != nil {
		return m.renderError(fmt.Errorf("failed to get summary: %v", err))
	}

	if summary == nil {
		return m.renderError(fmt.Errorf("summary is nil"))
	}

	message := m.styles.Intro.Render("Congratulations! You've completed all the lessons.\n\n" +
		"And the biggest lesson is:\ndon't run any binaries like this!\n\nRunning binaries you don't know the origin from\ncan give an attacker complete control over your computer.\n\n" +
		fmt.Sprintf("Here's a summary of the data collected:\n%d computers, %d GB of memory, "+
			"%d CPU cores, %d unique IPs.\n\n",
			summary.TotalHosts, summary.TotalMemoryGB, summary.TotalCPUCores,
			summary.UniqueIPCount) +
		"Be safe online, and keep hacking.\n\n- Team 0x641a <3 \n\nPress Enter to exit.")

	return lipgloss.Place(m.windowWidth, m.windowHeight,
		lipgloss.Center, lipgloss.Center,
		message)
}
var (
	summaryOnce  sync.Once
	summaryCached *SummaryResponse
	summaryErr   error
)

func getSummaryOnce() (*SummaryResponse, error) {
	summaryOnce.Do(func() {
		summaryCached, summaryErr = fetchSummary()
	})
	return summaryCached, summaryErr
}

func fetchSummary() (*SummaryResponse, error) {
	resp, err := http.Get("https://shellhacked.share.zrok.io/summary")
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var summary SummaryResponse
	err = json.Unmarshal(body, &summary)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return &summary, nil
}
