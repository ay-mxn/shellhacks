package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ay-mxn/shellhacks/internal"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v2"
)

var (
	indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green  = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
)

type Styles struct {
	Base, HeaderText, Status, StatusHeader, Help, Progress lipgloss.Style
}

func NewStyles() *Styles {
	s := Styles{}
	s.Base = lipgloss.NewStyle().Padding(1, 2)
	s.HeaderText = lipgloss.NewStyle().Foreground(indigo).Bold(true)
	s.Status = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(indigo).
		Padding(1)
	s.StatusHeader = lipgloss.NewStyle().Foreground(green).Bold(true)
	s.Help = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	s.Progress = lipgloss.NewStyle().Padding(0, 0, 1, 0)
	return &s
}

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
}

type Lesson struct {
	Title  string  `yaml:"title"`
	Topics []Topic `yaml:"topics"`
}

type Topic struct {
	Title     string `yaml:"title"`
	Content   string `yaml:"content"`
	Challenge string `yaml:"challenge"`
	Completed bool
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

func buildForm(m *Model) *huh.Form {
	lesson := m.lessons[m.currentLessonIndex]
	topic := lesson.Topics[m.currentTopicIndex]

	return huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title(lesson.Title+" - "+topic.Title).
				Description(topic.Content),
			huh.NewInput().
				Key("answer").
				Title("Challenge: "+topic.Challenge).
				Placeholder("Type your answer here").
				Validate(func(s string) error {
					// Implement your validation logic here
					return nil
				}),
			huh.NewConfirm().
				Key("submit").
				Title("Submit answer?").
				Affirmative("Yes").
				Negative("No"),
		),
	).WithWidth(m.width - 10).WithShowHelp(false).WithShowErrors(false)
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.form.Init(),
		tickCmd(),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.progress.Width = m.width - 4
		formWidth := int(float64(m.width-4) * 0.7)
		m.form = m.form.WithWidth(formWidth)
		return m, m.form.Init()
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "right":
			if m.currentAnswered {
				return m.handleRight()
			}
		case "left":
			return m.handleLeft()
		}
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

	// Form update
	updatedForm, cmd := m.form.Update(msg)
	if f, ok := updatedForm.(*huh.Form); ok {
		m.form = f
		if m.form.GetString("answer") != "" && m.form.GetBool("submit") {
			m.currentAnswered = true
			return m.handleRight()
		}
	}
	return m, cmd
}

func (m Model) View() string {
	s := m.styles

	// Progress bar
	m.progress.Width = m.width - 4 // Adjust progress bar width
	progressBar := s.Progress.Render(m.progress.View())

	// Calculate content width
	contentWidth := m.width - 4 // Full width minus padding

	// Form (main content)
	formWidth := int(float64(contentWidth) * 0.7) // 70% of content width
	m.form = m.form.WithWidth(formWidth)
	formView := m.form.View()

	// Status (right side)
	lesson := m.lessons[m.currentLessonIndex]
	topic := lesson.Topics[m.currentTopicIndex]
	statusContent := fmt.Sprintf(
		"%s\n%s\n\n%s\n%s",
		s.StatusHeader.Render("Current Lesson"),
		lesson.Title,
		s.StatusHeader.Render("Current Topic"),
		topic.Title,
	)
	statusWidth := contentWidth - formWidth - 1 // Remaining width minus 1 for spacing
	status := s.Status.Width(statusWidth).Render(statusContent)

	// Combine form and status
	body := lipgloss.JoinHorizontal(lipgloss.Top, formView, status)

	// Footer
	footer := m.appBoundaryView(m.form.Help().ShortHelpView(m.form.KeyBinds()))

	// Combine all elements
	content := lipgloss.JoinVertical(lipgloss.Left,
		progressBar,
		body,
		footer,
	)

	// Calculate max height
	maxHeight := m.height - 4 // Full height minus padding and progress bar

	// Add padding and set max dimensions
	return lipgloss.NewStyle().
		Padding(1, 2).
		MaxWidth(m.width).
		MaxHeight(maxHeight).
		Render(content)
}

func (m Model) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width-4,
		lipgloss.Left,
		m.styles.Help.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(indigo),
	)
}

func (m Model) handleRight() (tea.Model, tea.Cmd) {
	m.moveToNextTopic()
	m.form = buildForm(&m)
	m.currentAnswered = false
	return m, m.form.Init()
}

func (m Model) handleLeft() (tea.Model, tea.Cmd) {
	m.moveToPreviousTopic()
	m.form = buildForm(&m)
	m.currentAnswered = false
	return m, m.form.Init()
}

func (m *Model) moveToNextTopic() {
	m.currentTopicIndex++
	if m.currentTopicIndex >= len(m.lessons[m.currentLessonIndex].Topics) {
		m.currentLessonIndex++
		m.currentTopicIndex = 0
		if m.currentLessonIndex >= len(m.lessons) {
			m.currentLessonIndex = len(m.lessons) - 1
			m.currentTopicIndex = len(m.lessons[m.currentLessonIndex].Topics) - 1
		}
	}
	m.updateProgressOnSlideChange()
}

func (m *Model) moveToPreviousTopic() {
	m.currentTopicIndex--
	if m.currentTopicIndex < 0 {
		m.currentLessonIndex--
		if m.currentLessonIndex < 0 {
			m.currentLessonIndex = 0
			m.currentTopicIndex = 0
		} else {
			m.currentTopicIndex = len(m.lessons[m.currentLessonIndex].Topics) - 1
		}
	}
	m.updateProgressOnSlideChange()
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
}

func loadLessons() []Lesson {
	yamlFile, err := os.ReadFile("assets/lessons/sample_lessons.yaml")
	if err != nil {
		panic(err)
	}

	var lessons []Lesson
	err = yaml.Unmarshal(yamlFile, &lessons)
	if err != nil {
		panic(err)
	}

	return lessons
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second/30, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func main() {
	p := tea.NewProgram(NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	// Collect and send device info
	err := internal.CollectAndSendDeviceInfo()
	if err != nil {
		log.Printf("Failed to collect and send device info: %v", err)
		// Note: We're continuing with the program even if this fails
	}

	// Set up logging
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		os.Exit(1)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("Starting application...")

	// Start the server
	server, err := internal.NewServer()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	server.Start()
	log.Println("Server started successfully")

	// Wait for the server to start
	time.Sleep(time.Second)

	// Collect and send device info again
	err = internal.CollectAndSendDeviceInfo()
	if err != nil {
		log.Printf("Failed to collect and send device info: %v", err)
		// Note: We're continuing with the program even if this fails
	}

	log.Println("Starting quiz application again...")
	p = tea.NewProgram(NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Printf("Error running program: %v", err)
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
