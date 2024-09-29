package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
		m.progress.Width = m.width - 4
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f

		if m.form.State == huh.StateCompleted {
			if m.quizCompleted {
				if m.form.GetConfirmation("restart") {
					return m.restart(), nil
				}
				return m, tea.Quit
			}
			return m.moveToNextTopic(), nil
		}
	}

	return m, cmd
}

func (m Model) moveToNextTopic() (tea.Model, tea.Cmd) {
	m.currentTopicIndex++
	if m.currentTopicIndex >= len(m.lessons[m.currentLessonIndex].Topics) {
		m.currentLessonIndex++
		m.currentTopicIndex = 0
		if m.currentLessonIndex >= len(m.lessons) {
			m.quizCompleted = true
			m.form = buildCompletionForm(&m)
			return m, m.form.Init()
		}
	}
	m.updateProgress()
	m.form = buildForm(&m)
	return m, m.form.Init()
}

func (m *Model) updateProgress() {
	totalTopics := 0
	completedTopics := 0
	for _, lesson := range m.lessons {
		totalTopics += len(lesson.Topics)
		for _, topic := range lesson.Topics {
			if topic.Completed {
				completedTopics++
			}
		}
	}
	m.progress.SetPercent(float64(completedTopics) / float64(totalTopics))
}

func (m Model) restart() tea.Model {
	newModel := NewModel()
	newModel.width = m.width
	return newModel
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
