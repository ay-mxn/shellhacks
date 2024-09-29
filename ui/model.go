package ui

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	styles             *Styles
	form               *huh.Form
	width              int
	height             int
	lessons            []Lesson
	currentLessonIndex int
	currentTopicIndex  int
	progress           progress.Model
	targetPercent      float64
	currentAnswered    bool
	quizCompleted      bool
}

func NewModel() Model {
	m := Model{
		styles:  NewStyles(),
		lessons: loadLessons(),
		progress: progress.New(
			progress.WithDefaultGradient(),
			progress.WithWidth(78),
		),
	}
	m.form = buildForm(&m)
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Restart the quiz
func (m Model) restart() tea.Model {
	newModel := NewModel() // Restart with a fresh model
	newModel.width = m.width
	return newModel
}

// Move to the next topic
func (m *Model) moveToNextTopic() tea.Model {
	m.currentTopicIndex++
	if m.currentTopicIndex >= len(m.lessons[m.currentLessonIndex].Topics) {
		m.currentLessonIndex++
		m.currentTopicIndex = 0
		if m.currentLessonIndex >= len(m.lessons) {
			m.quizCompleted = true
			m.form = buildCompletionForm(m)
			return m
		}
	}
	m.form = buildForm(m)
	return m
}

func (m *Model) updateProgressOnSlideChange() {
	totalSlides := 0
	currentSlide := 0
	for i, lesson := range m.lessons {
		for j := range lesson.Topics {
			totalSlides++
			if i < m.currentLessonIndex || (i == m.currentLessonIndex && j <= m.currentTopicIndex) {
				currentSlide++
			}
		}
	}
	m.targetPercent = float64(currentSlide) / float64(totalSlides)
	m.progress.SetPercent(m.targetPercent) // Update the progress value
}
