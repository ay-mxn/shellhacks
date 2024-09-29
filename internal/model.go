package internal

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	maxWidth  = 100
	maxHeight = 30
	padding   = 2
)

type Model struct {
	state         int
	lessons       []Lesson
	currentLesson int
	currentTopic  int
	styles        *Styles
	progress      progress.Model
	textInput     textinput.Model
	challengeMsg  string
	viewport      viewport.Model
	content       string
	windowWidth   int
	windowHeight  int
	targetPercent float64
	allCompleted  bool  // New field to track completion of all lessons
}

const (
	stateIntro = iota
	stateContent
	stateChallenge
	stateAllCompleted  // New state for when all lessons are completed
)

func NewModel() Model {
	styles := NewStyles()
	lessons := loadLessons()

	ti := textinput.New()
	ti.Placeholder = "Type here..."
	ti.Focus()

	prog := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(maxWidth-padding*2),
	)

	m := Model{
		state:         stateIntro,
		lessons:       lessons,
		styles:        styles,
		progress:      prog,
		textInput:     ti,
		viewport:      viewport.New(maxWidth, maxHeight-6),
		targetPercent: 0,
	}

	m.updateProgress()
	if len(m.lessons) > 0 && len(m.lessons[0].Topics) > 0 {
		m.updateContent()
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		textinput.Blink,
		tickCmd(),
	)
}

func (m Model) handleUpDown(key string) (Model, tea.Cmd) {
	if m.state == stateContent {
		if key == "up" {
			m.viewport.LineUp(1)
		} else {
			m.viewport.LineDown(1)
		}
	}
	return m, nil
}

func (m *Model) updateContent() {
	if len(m.lessons) == 0 || len(m.lessons[m.currentLesson].Topics) == 0 {
		m.content = "No lessons or topics available."
		return
	}
	m.content = m.lessons[m.currentLesson].Topics[m.currentTopic].Content
	m.viewport.SetContent(m.content)
}

func (m *Model) updateViewportSize() {
	width := min(m.windowWidth, maxWidth)
	height := min(m.windowHeight, maxHeight)
	m.viewport.Width = width - 4
	m.viewport.Height = height - 10
	m.textInput.Width = m.viewport.Width
	m.progress.Width = width - padding*2
}

func (m *Model) setPhishingEmailContent() {
	m.viewport.SetContent(phishingEmail)
}

func (m *Model) updateProgress() {
	m.targetPercent = m.calculateProgress()
}

func (m *Model) calculateProgress() float64 {
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
	if totalTopics == 0 {
		return 0
	}
	return float64(completedTopics) / float64(totalTopics)
}

func (m *Model) canMoveToNextTopic() bool {
	currentTopic := m.lessons[m.currentLesson].Topics[m.currentTopic]
	return currentTopic.Challenge == "" || currentTopic.Completed
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second/30, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
