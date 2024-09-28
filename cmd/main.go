package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
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
    Title         string   `yaml:"title"`
    Topics        []Topic  `yaml:"topics"`
    ChallengeFunc func(*model) bool `yaml:"-"`
    ChallengeType string   // Add this line
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


func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

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
            if len(m.lessons) > 0 {
                if m.lessons[m.currentLesson].ChallengeFunc != nil {
                    result := m.lessons[m.currentLesson].ChallengeFunc(&m)
                    log.Printf("Challenge function result: %v", result)
                    if result {
                        m.challengeMsg = "Correct! Press right arrow to continue."
                        m.lessons[m.currentLesson].Topics[m.currentTopic].Completed = true
                    } else {
                        m.challengeMsg = "Try again."
                    }
                } else {
                    log.Printf("No challenge function for lesson: %s", m.lessons[m.currentLesson].Title)
                    m.challengeMsg = "No challenge function defined for this lesson."
                }
            } else {
                log.Printf("No lessons available")
                m.challengeMsg = "No lessons available."
            }
        }
		case "right":
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

func loadLessons() ([]Lesson, error) {
    lessonsDir := "assets/lessons"
    files, err := ioutil.ReadDir(lessonsDir)
    if err != nil {
        return nil, fmt.Errorf("failed to read lessons directory: %v", err)
    }

    var allLessons []Lesson
    for _, file := range files {
        if filepath.Ext(file.Name()) != ".yaml" && filepath.Ext(file.Name()) != ".yml" {
            continue
        }

        filePath := filepath.Join(lessonsDir, file.Name())
        yamlFile, err := ioutil.ReadFile(filePath)
        if err != nil {
            return nil, fmt.Errorf("failed to read lesson file %s: %v", file.Name(), err)
        }

        var lessons []Lesson
        err = yaml.Unmarshal(yamlFile, &lessons)
        if err != nil {
            return nil, fmt.Errorf("failed to unmarshal lessons from %s: %v", file.Name(), err)
        }

        log.Printf("Loaded %d lessons from %s", len(lessons), file.Name())

        // Set the lesson title from the filename if not specified in the YAML
        for i := range lessons {
            if lessons[i].Title == "" {
                lessons[i].Title = strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
            }
            log.Printf("Lesson title: %s", lessons[i].Title)

            // Assign challenge functions based on lesson title
            // Assign challenge functions based on lesson title
						switch strings.ToLower(lessons[i].Title) {
						case "passwords":
								lessons[i].ChallengeFunc = passwordStrengthChallenge
								lessons[i].ChallengeType = "passwordStrength"
								log.Printf("Assigned passwordStrengthChallenge to lesson: %s", lessons[i].Title)
						case "phishing":
								lessons[i].ChallengeFunc = phishingAwarenessChallenge
								lessons[i].ChallengeType = "phishingAwareness"
								log.Printf("Assigned phishingAwarenessChallenge to lesson: %s", lessons[i].Title)
						default:
								lessons[i].ChallengeFunc = defaultChallenge
								lessons[i].ChallengeType = "default"
								log.Printf("Assigned defaultChallenge to lesson: %s", lessons[i].Title)
						}
        }

        allLessons = append(allLessons, lessons...)
    }

    if len(allLessons) == 0 {
        return nil, fmt.Errorf("no lessons found in %s", lessonsDir)
    }

    log.Printf("Total lessons loaded: %d", len(allLessons))
    return allLessons, nil
}

func NewModel() model {
	styles := NewStyles()

	lessons, err := loadLessons()
	if err != nil {
		log.Printf("Failed to load lessons: %v", err)
		lessons = []Lesson{}
	}

	ti := textinput.New()
	ti.Placeholder = "Type here..."
	ti.Focus()

	m := model{
		state:         stateIntro,
		lessons:       lessons,
		currentLesson: 0,
		currentTopic:  0,
		styles:        styles,
		progress:      progress.New(progress.WithDefaultGradient()),
		textInput:     ti,
		viewport:      viewport.New(maxWidth, maxHeight-6),
	}

	if len(m.lessons) > 0 && len(m.lessons[0].Topics) > 0 {
		m.updateContent()
	}

	return m
}

func (m *model) updateContent() {
	if len(m.lessons) == 0 || len(m.lessons[m.currentLesson].Topics) == 0 {
		m.content = "No lessons or topics available."
		return
	}
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

// Modify the challenge functions to include logging
func passwordStrengthChallenge(m *model) bool {
    password := m.textInput.Value()
    log.Printf("Password challenge input: %s", password)
    result := len(password) >= 12 &&
        strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") &&
        strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") &&
        strings.ContainsAny(password, "0123456789") &&
        strings.ContainsAny(password, "!@#$%^&*()_+-=[]{}|;:,.<>?")
    log.Printf("Password challenge result: %v", result)
    return result
}

func phishingAwarenessChallenge(m *model) bool {
    answer := strings.ToLower(m.textInput.Value())
    log.Printf("Phishing challenge input: %s", answer)
    result := strings.Contains(answer, "urgent") || 
           strings.Contains(answer, "personal information") ||
           strings.Contains(answer, "suspicious url") ||
           strings.Contains(answer, "generic greeting") ||
           strings.Contains(answer, "poor grammar") ||
           strings.Contains(answer, "unexpected attachment")
    log.Printf("Phishing challenge result: %v", result)
    return result
}

func defaultChallenge(m *model) bool {
    log.Printf("Default challenge called with input: %s", m.textInput.Value())
    return false
}



func main() {
	p := tea.NewProgram(NewModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Printf("There's been an error: %v", err)
	}
}
