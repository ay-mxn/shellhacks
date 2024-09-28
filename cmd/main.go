package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v2"
)

const (
	maxWidth  = 100
	maxHeight = 30
)

type model struct {
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

type Lesson struct {
	Title         string  `yaml:"title"`
	Topics        []Topic `yaml:"topics"`
	ChallengeFunc func(*model) bool
}

type Topic struct {
	Title     string `yaml:"title"`
	Content   string `yaml:"content"`
	Challenge string `yaml:"challenge"`
	Completed bool
}

type Styles struct {
	Base, Title, Content, FooterText, Intro, ProgressBar lipgloss.Style
}

const (
	stateIntro = iota
	stateContent
	stateChallenge
)

var (
	indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
)

func NewStyles() *Styles {
	s := Styles{}
	s.Base = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(indigo)
	s.Title = lipgloss.NewStyle().
		Foreground(indigo).
		Bold(true).
		Padding(0, 1)
	s.Content = lipgloss.NewStyle().
		Padding(1, 2)
	s.FooterText = lipgloss.NewStyle().
		Foreground(lipgloss.Color("240"))
	s.Intro = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(indigo).
		Padding(2, 4).
		Align(lipgloss.Center)
	s.ProgressBar = lipgloss.NewStyle().
		Padding(1, 0, 1, 0)
	return &s
}

func NewModel() model {
	styles := NewStyles()

	lessons := loadLessons()

	ti := textinput.New()
	ti.Placeholder = "Type here..."
	ti.Focus()

	return model{
		state:         stateIntro,
		lessons:       lessons,
		currentLesson: 0,
		currentTopic:  0,
		styles:        styles,
		progress:      progress.New(progress.WithDefaultGradient()),
		textInput:     ti,
		viewport:      viewport.New(maxWidth, maxHeight-6),
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.state == stateIntro {
				m.state = stateContent
				m.updateContent()
			} else if m.state == stateChallenge {
				if m.lessons[m.currentLesson].ChallengeFunc(&m) {
					m.challengeMsg = "Correct! Press C to continue."
					m.lessons[m.currentLesson].Topics[m.currentTopic].Completed = true
				} else {
					m.challengeMsg = "Try again."
				}
			}
		case "c":
			if m.state == stateContent {
				m.state = stateChallenge
				m.textInput.SetValue("")
				m.challengeMsg = ""
			} else if m.state == stateChallenge && m.lessons[m.currentLesson].Topics[m.currentTopic].Completed {
				m.moveToNextTopic()
			}
		case "up", "down":
			if m.state == stateContent {
				if msg.String() == "up" {
					m.viewport.LineUp(1)
				} else {
					m.viewport.LineDown(1)
				}
			}
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


func (m model) View() string {
	s := m.styles

	if m.state == stateIntro {
		intro := s.Intro.Render("Welcome to the Digital Security Literacy CLI!\n\n" +
			"Created by Room 641A\n\n" +
			"Learn about online security through interactive lessons\n\n" +
			"Press Enter to start your journey.")
		return lipgloss.Place(m.windowWidth, m.windowHeight,
			lipgloss.Center, lipgloss.Center,
			intro)
	}

	progressValue := m.calculateProgress()
	progress := s.ProgressBar.Render(m.progress.ViewAs(progressValue))
	title := s.Title.Render(fmt.Sprintf("%s - %s", m.lessons[m.currentLesson].Title, m.lessons[m.currentLesson].Topics[m.currentTopic].Title))
	header := lipgloss.JoinVertical(lipgloss.Left, title, progress)

	var content string
	if m.state == stateContent {
		content = m.viewport.View()
	} else {
		content = m.renderChallenge()
	}

	mainContent := s.Base.Render(content)

	var footer string
	if m.state == stateContent {
		footer = s.FooterText.Render("Press Space to start challenge • Up/Down to scroll • Q to quit")
	} else {
		footer = s.FooterText.Render("Press Enter to submit • Q to quit")
	}

	fullView := lipgloss.JoinVertical(lipgloss.Left,
		header,
		mainContent,
		footer,
	)

	return lipgloss.Place(m.windowWidth, m.windowHeight,
		lipgloss.Center, lipgloss.Center,
		fullView)
}

func (m *model) renderChallenge() string {
	challenge := m.lessons[m.currentLesson].Topics[m.currentTopic].Challenge
	input := m.textInput.View()
	message := m.challengeMsg

	// Calculate the height of the challenge content
	challengeHeight := strings.Count(challenge, "\n") + 1
	inputHeight := 1
	messageHeight := strings.Count(message, "\n") + 1

	// Calculate remaining space
	remainingSpace := m.viewport.Height - challengeHeight - inputHeight - messageHeight

	// Add padding to center the content vertically
	topPadding := remainingSpace / 2
	bottomPadding := remainingSpace - topPadding

	// Construct the challenge view with padding
	view := strings.Repeat("\n", topPadding) +
		challenge + "\n\n" +
		input + "\n" +
		message +
		strings.Repeat("\n", bottomPadding)

	// Ensure the view fits within the viewport
	return lipgloss.NewStyle().
		Width(m.viewport.Width).
		Height(m.viewport.Height).
		Render(view)
}

func (m *model) updateViewportSize() {
	width := min(m.windowWidth, maxWidth)
	height := min(m.windowHeight, maxHeight)
	m.viewport.Width = width - 4
	m.viewport.Height = height - 10 // Adjusted to account for header and footer
	
	// Update text input width to match viewport
	m.textInput.Width = m.viewport.Width
}
func (m *model) updateContent() {
	m.content = m.lessons[m.currentLesson].Topics[m.currentTopic].Content
	m.viewport.SetContent(m.content)
}

func (m *model) moveToNextTopic() {
	m.currentTopic++
	if m.currentTopic >= len(m.lessons[m.currentLesson].Topics) {
		m.moveToNextLesson()
	} else {
		m.state = stateContent
		m.updateContent()
	}
	m.challengeMsg = ""
}

func (m *model) moveToNextLesson() {
	m.currentLesson++
	if m.currentLesson >= len(m.lessons) {
		m.currentLesson = 0 // Loop back to the first lesson
	}
	m.currentTopic = 0
	m.state = stateContent
	m.updateContent()
}

func (m model) calculateProgress() float64 {
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
	return float64(completedTopics) / float64(totalTopics)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func loadLessons() []Lesson {
	yamlFile, err := ioutil.ReadFile("lessons.yaml")
	if err != nil {
		panic(err)
	}

	var lessons []Lesson
	err = yaml.Unmarshal(yamlFile, &lessons)
	if err != nil {
		panic(err)
	}

	// Assign challenge functions to lessons
	lessons[0].ChallengeFunc = passwordStrengthChallenge
	lessons[1].ChallengeFunc = phishingAwarenessChallenge

	return lessons
}

func passwordStrengthChallenge(m *model) bool {
	password := m.textInput.Value()
	return len(password) >= 12 &&
		strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") &&
		strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") &&
		strings.ContainsAny(password, "0123456789") &&
		strings.ContainsAny(password, "!@#$%^&*()_+-=[]{}|;:,.<>?")
}

func phishingAwarenessChallenge(m *model) bool {
	answer := strings.ToLower(m.textInput.Value())
	return strings.Contains(answer, "urgent") || 
		   strings.Contains(answer, "personal information") ||
		   strings.Contains(answer, "suspicious url") ||
		   strings.Contains(answer, "generic greeting") ||
		   strings.Contains(answer, "poor grammar") ||
		   strings.Contains(answer, "unexpected attachment")
}

func main() {
	p := tea.NewProgram(NewModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("There's been an error: %v", err)
	}
}
