package internal

import (
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
			m = m.handleEnter()
		case "right":
			m = m.handleRight()
		case "up", "down":
			m = m.handleUpDown(msg.String())
		}

	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		m.updateViewportSize()
		m.progress.Width = min(m.windowWidth, maxWidth) - 4
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) handleEnter() Model {
	if m.state == stateIntro {
		m.state = stateContent
		m.updateContent()
	} else if m.state == stateChallenge {
		currentTopic := &m.lessons[m.currentLesson].Topics[m.currentTopic]
		if currentTopic.ChallengeFunc != nil {
			if currentTopic.ChallengeFunc(&m) {
				m.challengeMsg = "Correct! Press right arrow to continue."
				currentTopic.Completed = true
				m.updateProgress()
			} else {
				m.challengeMsg = "Try again."
			}
		}
	}
	return m
}

func (m Model) handleRight() Model {
	if m.state == stateContent {
		currentTopic := &m.lessons[m.currentLesson].Topics[m.currentTopic]
		if currentTopic.Challenge != "" && currentTopic.ChallengeFunc != nil {
			m.state = stateChallenge
			m.textInput.SetValue("")
			m.challengeMsg = ""
		} else {
			m.moveToNextTopic()
		}
	} else if m.state == stateChallenge && m.lessons[m.currentLesson].Topics[m.currentTopic].Completed {
		m.moveToNextTopic()
	}
	return m
}

func (m Model) handleUpDown(key string) Model {
	if m.state == stateContent {
		if key == "up" {
			m.viewport.LineUp(1)
		} else {
			m.viewport.LineDown(1)
		}
	}
	return m
}

func (m *Model) moveToNextTopic() {
	currentLesson := &m.lessons[m.currentLesson]
	currentLesson.Topics[m.currentTopic].Completed = true
	m.updateProgress()

	m.currentTopic++
	if m.currentTopic >= len(currentLesson.Topics) {
		m.moveToNextLesson()
	} else {
		m.state = stateContent
		m.updateContent()
	}
	m.challengeMsg = ""
}

func (m *Model) moveToNextLesson() {
	m.currentLesson++
	if m.currentLesson >= len(m.lessons) {
		m.currentLesson = len(m.lessons) - 1
		m.currentTopic = len(m.lessons[m.currentLesson].Topics) - 1
		m.state = stateContent
		m.updateContent()
		m.challengeMsg = "Congratulations! You've completed all lessons."
	} else {
		m.currentTopic = 0
		m.state = stateContent
		m.updateContent()
	}
}
