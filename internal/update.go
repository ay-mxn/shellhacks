package internal

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mbndr/figlet4go"
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
			if m.state == stateAllCompleted {
				return m, tea.Quit
			}

			return m.handleEnter()
		case "right":
			return m.handleRight()
		case "left":
			return m.handleLeft()
		case "up", "down":
			updatedModel, cmd := m.handleUpDown(msg.String())
			return updatedModel, cmd
		}
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		m.updateViewportSize()
	case tickMsg:
		if m.progress.Percent() < m.targetPercent {
			cmd := m.progress.IncrPercent(0.01)
			return m, tea.Batch(tickCmd(), cmd)
		} else if m.progress.Percent() > m.targetPercent {
			cmd := m.progress.SetPercent(m.targetPercent)
			return m, tea.Batch(tickCmd(), cmd)
		}
		return m, tickCmd()
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	}

	if m.targetPercent >= 1.0 && !m.allCompleted {
		m.allCompleted = true
		m.state = stateAllCompleted
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
		m.updateProgressOnSlideChange()
	case stateChallenge:
		topic := &m.lessons[m.currentLesson].Topics[m.currentTopic]
		if topic.ChallengeFunc != nil && topic.ChallengeFunc(&m) {
			topic.Completed = true
			m.moveToNextTopic()
			m.updateProgressOnSlideChange()
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
		} else if m.canMoveToNextTopic() {
			m.moveToNextTopic()
			m.updateProgressOnSlideChange()
		}
	} else if m.state == stateChallenge && m.lessons[m.currentLesson].Topics[m.currentTopic].Completed {
		m.moveToNextTopic()
		m.updateProgressOnSlideChange()
	}
	return m, nil
}

func (m Model) handleLeft() (tea.Model, tea.Cmd) {
	if m.state == stateContent || m.state == stateChallenge {
		m.moveToPreviousTopic()
		m.updateProgressOnSlideChange()
	}
	return m, nil
}

func (m *Model) moveToNextTopic() {
	m.currentTopic++
	if m.currentTopic >= len(m.lessons[m.currentLesson].Topics) {
		if m.currentLesson < len(m.lessons)-1 {
			m.currentLesson++
			m.currentTopic = 0
		} else {
			m.currentTopic = len(m.lessons[m.currentLesson].Topics) - 1
		}
	}
	m.state = stateContent
	m.updateContent()
}

func (m *Model) moveToPreviousTopic() {
	m.currentTopic--
	if m.currentTopic < 0 {
		if m.currentLesson > 0 {
			m.currentLesson--
			m.currentTopic = len(m.lessons[m.currentLesson].Topics) - 1
		} else {
			m.currentTopic = 0
		}
	}
	m.state = stateContent
	m.updateContent()
}

func (m *Model) updateProgressOnSlideChange() {
	totalSlides := 0
	currentSlide := 0
	for i, lesson := range m.lessons {
		for j := range lesson.Topics {
			totalSlides++
			if i < m.currentLesson || (i == m.currentLesson && j <= m.currentTopic) {
				currentSlide++
			}
		}
	}
	m.targetPercent = float64(currentSlide) / float64(totalSlides)
}

func (m Model) focusTextInput() tea.Msg {
	m.textInput.Focus()
	return nil
}

func generateASCIIArt() string {
	ascii := figlet4go.NewAsciiRender()

	// Set custom options
	options := figlet4go.NewRenderOptions()
	options.FontName = "roman"
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorRed,
		figlet4go.ColorYellow,
		figlet4go.ColorGreen,
		figlet4go.ColorCyan,
		figlet4go.ColorBlue,
		figlet4go.ColorMagenta,
	}

	renderStr, err := ascii.RenderOpts("Digital Security", options)
	if err != nil {
		return "Digital Security Literacy CLI\n"
	}

	return renderStr
}
