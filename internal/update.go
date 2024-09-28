package internal

import (
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

// Update handles all the application logic and state updates
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			return m.handleEnter()
		case "right":
			return m.handleRight()
		case "up", "down":
			return m.handleUpDown(msg.String())
		}
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		m.updateViewportSize()
		m.progress.Width = min(m.windowWidth, maxWidth) - 4
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds := append([]tea.Cmd{cmd}, m.updateTextInput(msg)...)

	return m, tea.Batch(cmds...)
}

func (m Model) handleEnter() (tea.Model, tea.Cmd) {
	switch m.state {
	case stateIntro:
		m.state = stateContent
		m.updateContent()
	case stateChallenge:
		topic := &m.lessons[m.currentLesson].Topics[m.currentTopic]
		if topic.ChallengeFunc != nil && topic.ChallengeFunc(&m) {
			m.challengeMsg = "Correct! Press right arrow to continue."
			topic.Completed = true
			m.updateProgress()
		} else {
			m.challengeMsg = "Try again."
		}
	}
	return m, nil
}

func (m Model) handleRight() (tea.Model, tea.Cmd) {
	if m.state == stateContent {
		topic := &m.lessons[m.currentLesson].Topics[m.currentTopic]
		if topic.Challenge != "" && !topic.Completed {
			m.state = stateChallenge
			m.textInput.SetValue("")
			m.challengeMsg = ""
		} else {
			m.moveToNextTopic()
		}
	} else if m.state == stateChallenge && m.lessons[m.currentLesson].Topics[m.currentTopic].Completed {
		m.moveToNextTopic()
	}
	return m, nil
}

func (m Model) handleUpDown(key string) (tea.Model, tea.Cmd) {
	if m.state == stateContent {
		if key == "up" {
			m.viewport.LineUp(1)
		} else {
			m.viewport.LineDown(1)
		}
	}
	return m, nil
}

func (m *Model) moveToNextTopic() {
	m.currentTopic++
	if m.currentTopic >= len(m.lessons[m.currentLesson].Topics) {
		m.currentLesson++
		m.currentTopic = 0
		if m.currentLesson >= len(m.lessons) {
			m.currentLesson = 0
			m.state = stateIntro
		} else {
			m.state = stateContent
		}
	} else {
		m.state = stateContent
	}
	m.updateContent()
	m.updateProgress()
}

func (m Model) updateTextInput(msg tea.Msg) []tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)
	return cmds
}

// Additional helper function to generate ASCII art
func generateASCIIArt() string {
	cmd := exec.Command("bash", "-c", "figlet -f roman -t -c Digital Security | lolcat")
	asciiArt, err := cmd.Output()
	if err != nil {
		return "Digital Security Literacy CLI\n"
	}
	return string(asciiArt)
}
