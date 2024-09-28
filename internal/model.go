package internal

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	maxWidth  = 100
	maxHeight = 30
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
}

const (
	stateIntro = iota
	stateContent
	stateChallenge
)

func NewModel() Model {
	styles := NewStyles()
	lessons := loadLessons()

	ti := textinput.New()
	ti.Placeholder = "Type here..."
	ti.Focus()

	progressBar := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(80),
		progress.WithoutPercentage(),
	)

	m := Model{
		state:         stateIntro,
		lessons:       lessons,
		styles:        styles,
		progress:      progressBar,
		textInput:     ti,
		viewport:      viewport.New(maxWidth, maxHeight-6),
	}

	m.updateProgress()
	if len(m.lessons) > 0 && len(m.lessons[0].Topics) > 0 {
		m.updateContent()
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
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
}

func (m *Model) setPhishingEmailContent() {
	// Set the viewport content for the phishing challenge
	m.viewport.SetContent(phishingEmail)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
