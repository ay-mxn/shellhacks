package internal

import (
	"os/exec"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

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
	case tickMsg:
		if m.progress.Percent() < m.targetPercent {
			cmd := m.progress.IncrPercent(0.01)
			return m, tea.Batch(tickCmd(), cmd)
		}
		return m, tickCmd()
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	}

	if m.state == stateChallenge {
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		cmds = append(cmds, cmd)
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

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
		m.textInput.SetValue("")
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
			return m, m.focusTextInput
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

func (m Model) focusTextInput() tea.Msg {
	m.textInput.Focus()
	return nil
}

func generateASCIIArt() string {
	cmd := exec.Command("bash", "-c", "figlet -f roman -t -c Digital Security | lolcat")
	asciiArt, err := cmd.Output()
	if err != nil {
		return "Digital Security Literacy CLI\n"
	}
	return string(asciiArt)
}
