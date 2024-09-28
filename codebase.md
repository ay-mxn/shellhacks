# lessons.yaml

```yaml
- title: "Safe Passwords"
  topics:
    - title: "Password Strength"
      content: |
        Password strength is crucial for protecting your online accounts. A strong password:

        1. Is at least 12 characters long
        2. Contains a mix of uppercase and lowercase letters
        3. Includes numbers and special characters
        4. Avoids common words or phrases
        5. Is unique for each account

        Remember: The longer and more complex your password, the harder it is to crack.
      challenge: "Create a strong password based on the guidelines above:"
    - title: "Password Managers"
      content: |
        Password managers are tools that help you create, store, and manage complex passwords securely. Benefits include:

        1. Generating strong, unique passwords for each account
        2. Storing passwords in an encrypted vault
        3. Auto-filling login forms
        4. Synchronizing across devices
        5. Alerting you to potentially compromised passwords

        Popular password managers include LastPass, 1Password, and Bitwarden.
      challenge: "What is one benefit of using a password manager?"

- title: "Phishing Awareness"
  topics:
    - title: "Recognizing Phishing Attempts"
      content: |
        Phishing is a technique used by cybercriminals to trick you into revealing sensitive information. Common signs of phishing include:

        1. Urgent or threatening language
        2. Requests for personal information
        3. Suspicious or mismatched URLs
        4. Generic greetings
        5. Poor spelling and grammar
        6. Unexpected attachments

        Always verify the sender's identity and be cautious of unsolicited messages.
      challenge: "Name one sign of a phishing attempt:"
    - title: "Safe Email Practices"
      content: |
        Protecting yourself from email-based threats involves following these best practices:

        1. Don't click on links or download attachments from unknown senders
        2. Use spam filters and keep them updated
        3. Enable two-factor authentication on your email account
        4. Be cautious of emails asking for personal information
        5. Verify the sender's email address
        6. Use a secure email service that offers encryption

        When in doubt, contact the supposed sender through a known, trusted channel to verify the email's legitimacy.
      challenge: "What should you do if you're unsure about an email's legitimacy?"

```

# go.mod

```mod
module github.com/ay-mxn/shellhacks

go 1.22.3

require (
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/charmbracelet/bubbles v0.20.0 // indirect
	github.com/charmbracelet/bubbletea v1.1.1 // indirect
	github.com/charmbracelet/harmonica v0.2.0 // indirect
	github.com/charmbracelet/lipgloss v0.13.0 // indirect
	github.com/charmbracelet/x/ansi v0.2.3 // indirect
	github.com/charmbracelet/x/term v0.2.0 // indirect
	github.com/erikgeiser/coninput v0.0.0-20211004153227-1c3628e74d0f // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-localereader v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6 // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/sahilm/fuzzy v0.1.1 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.3.8 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

```

# debug.log

```log

```

# SECURITY.md

```md
# Detailed Digital Security Literacy CLI Hackathon Plan (36 Hours)

## Hour-by-Hour Breakdown

### Hour 0-1: Project Setup and Planning
- All: Quick team meeting to align on goals and approach
- Backend: Set up Git repository, create project structure, initialize Go module
- Frontend: Set up development environment, install Charm libraries
- Curriculum: Start outlining lesson structure and topics

### Hours 1-4: Core Framework and First Lesson
- Backend (3h):
  1. Create main.go with basic Bubble Tea app structure (30m)
  2. Implement basic app state management (1h)
  3. Create content loading function (reads from string constants initially) (1h)
  4. Set up basic navigation logic between screens (30m)
- Frontend (3h):
  1. Create main menu UI (1h)
  2. Implement basic lesson display UI (1h)
  3. Add simple styling with Lipgloss (1h)
- Curriculum (3h):
  1. Write content for "Safe Passwords" lesson (2h)
  2. Create 10 quiz questions for "Safe Passwords" (1h)

### Hours 4-8: Quiz Implementation and Second Lesson
- Backend (4h):
  1. Implement quiz logic (presenting questions, checking answers) (2h)
  2. Add score calculation and storage (1h)
  3. Integrate quiz flow with main app (1h)
- Frontend (4h):
  1. Create quiz interface UI (2h)
  2. Implement progress indicator for lessons/quizzes (1h)
  3. Add transitions between screens (1h)
- Curriculum (4h):
  1. Write content for "Phishing Awareness" lesson (2h)
  2. Create 10 quiz questions for "Phishing Awareness" (1h)
  3. Start on "Safe Browsing Habits" lesson (1h)

### Hours 8-12: Enhance User Experience and Third Lesson
- Backend (4h):
  1. Implement user progress tracking (2h)
  2. Add error handling and input validation (1h)
  3. Create a simple file-based storage system for persistence (1h)
- Frontend (4h):
  1. Add animations for text display in lessons (2h)
  2. Implement interactive elements in lessons (e.g., clickable examples) (2h)
- Curriculum (4h):
  1. Finish "Safe Browsing Habits" lesson (2h)
  2. Create 10 quiz questions for "Safe Browsing Habits" (1h)
  3. Start on "Data Privacy" lesson (1h)

### Hours 12-16: Fourth Lesson and Advanced Features
- Backend (4h):
  1. Implement a simple achievement system (2h)
  2. Add timed quizzes option (2h)
- Frontend (4h):
  1. Create UI for achievements (1h)
  2. Implement timed quiz UI elements (1h)
  3. Add more advanced styling and layouts (2h)
- Curriculum (4h):
  1. Finish "Data Privacy" lesson (2h)
  2. Create 10 quiz questions for "Data Privacy" (1h)
  3. Start on "Device Security" lesson (1h)

### Hours 16-20: Fifth Lesson and Polish
- Backend (4h):
  1. Implement difficulty levels for quizzes (2h)
  2. Add a review mode for missed questions (2h)
- Frontend (4h):
  1. Create UI for difficulty selection (1h)
  2. Implement review mode interface (2h)
  3. Add final polish to all UI elements (1h)
- Curriculum (4h):
  1. Finish "Device Security" lesson (2h)
  2. Create 10 quiz questions for "Device Security" (1h)
  3. Write app introduction and conclusion texts (1h)

### Hours 20-24: Integration and Testing
- All (4h each):
  1. Integrate all components
  2. Conduct thorough testing
  3. Fix bugs and address issues
  4. Optimize performance

### Hours 24-28: Extra Features and Content
- Backend (4h):
  1. Implement a hint system for quizzes (2h)
  2. Add a "quick quiz" mode that pulls random questions (2h)
- Frontend (4h):
  1. Create animations for achievements and completion (2h)
  2. Implement a help/about screen with app information (2h)
- Curriculum (4h):
  1. Create additional quiz questions for all topics (2h)
  2. Write "cybersecurity tips of the day" feature (2h)

### Hours 28-32: Final Polish and Preparation
- All (4h each):
  1. Final round of bug fixes and optimizations
  2. Prepare demonstration script and slides
  3. Create a brief user guide
  4. Plan future feature ideas to discuss during presentation

### Hours 32-36: Rehearsal and Last-Minute Improvements
- All (4h each):
  1. Rehearse demonstration
  2. Address any last-minute issues
  3. Prepare answers for potential questions
  4. Final team sync and motivation boost!

## Key Milestones:
1. Hour 4: Basic app with one lesson and main menu
2. Hour 8: Quiz functionality and two lessons
3. Hour 16: Four lessons, basic achievements, and polished UI
4. Hour 24: All five lessons, advanced features, fully integrated
5. Hour 32: Extra features, polished product, ready for final testing

Remember to take short breaks, stay hydrated, and support each other throughout the hackathon. Good luck!

```

# README.md

```md
# Digital Security Literacy CLI

A command-line tool using Charm libraries to educate users about online security through interactive lessons and quizzes.

## Core Features

1. Interactive Lessons
   - Text-based lessons with animated elements using Charm libraries
   - User interaction to reinforce learning

2. Quizzes
   - Multiple-choice and true/false questions
   - Immediate feedback and explanations

3. Progress Tracking
   - Save user progress locally
   - Display completion status for each lesson

4. Customizable Content
   - Easy to add or modify lessons and quizzes

5. Offline Functionality
   - All content available offline after initial download

## Lesson Topics

1. Safe Passwords
   - Password strength
   - Password managers
   - Multi-factor authentication

2. Phishing Awareness
   - Recognizing phishing attempts
   - Safe email practices

3. Safe Browsing Habits
   - HTTPS importance
   - Recognizing fake websites
   - Browser security settings

4. Data Privacy
   - Understanding data collection
   - Privacy settings on social media
   - VPNs and their uses

5. Device Security
   - Device encryption
   - Software updates
   - Secure Wi-Fi usage

## Technical Implementation

- Use Bubble Tea for the main application loop
- Implement custom UI components with Lipgloss
- Store progress and quiz results locally using a simple file-based system
- Utilize Bubbles for interactive elements like text input and selection

## User Experience

- Clear, step-by-step lessons with animated text
- Interactive quizzes with immediate feedback
- Progress bar showing overall course completion
- Command-line arguments for quick access to specific lessons or quizzes

## Ethical Considerations

- Be transparent about all functionality
- No collection of personal data
- Provide accurate, up-to-date security information
- Include resources for further learning

```

# CODE_OF_CONDUCT.md

```md
# Digital Security Literacy CLI Architecture

## Tech Stack
- Language: Go
- CLI Framework: [Bubble Tea](https://github.com/charmbracelet/bubbletea) (for TUI)
- Styling: [Lipgloss](https://github.com/charmbracelet/lipgloss)
- Input Handling: [Bubbles](https://github.com/charmbracelet/bubbles)
- Storage: Local file system for progress and settings
- Content Management: YAML files for lesson content and quizzes

## Project Structure
\`\`\`
security-literacy-cli/
├── cmd/
│   └── securitycli/
│       └── main.go
├── internal/
│   ├── tui/
│   │   ├── app.go
│   │   ├── menu.go
│   │   ├── lesson.go
│   │   ├── quiz.go
│   │   └── progress.go
│   ├── content/
│   │   ├── loader.go
│   │   └── models.go
│   ├── storage/
│   │   └── progress.go
│   └── utils/
│       └── helpers.go
├── pkg/
│   ├── lesson/
│   │   └── renderer.go
│   └── quiz/
│       └── engine.go
├── assets/
│   ├── lessons/
│   │   ├── passwords.yaml
│   │   ├── phishing.yaml
│   │   └── ...
│   └── quizzes/
│       ├── passwords_quiz.yaml
│       ├── phishing_quiz.yaml
│       └── ...
├── go.mod
├── go.sum
└── README.md
\`\`\`

## Component Overview

1. **Main Application (cmd/securitycli/main.go)**
   - Entry point for the application
   - Initializes the Bubble Tea application
   - Handles command-line arguments

2. **TUI Components (internal/tui/)**
   - `app.go`: Main TUI application structure
   - `menu.go`: Main menu for navigating lessons and quizzes
   - `lesson.go`: UI for displaying lesson content
   - `quiz.go`: UI for quiz interactions
   - `progress.go`: UI for displaying user progress

3. **Content Management (internal/content/)**
   - `loader.go`: Functions to load and parse YAML content
   - `models.go`: Structures for lesson and quiz data

4. **Storage (internal/storage/)**
   - `progress.go`: Functions to save and load user progress

5. **Utility Functions (internal/utils/)**
   - `helpers.go`: Common utility functions

6. **Lesson Rendering (pkg/lesson/)**
   - `renderer.go`: Functions to render lesson content with animations

7. **Quiz Engine (pkg/quiz/)**
   - `engine.go`: Quiz logic, scoring, and feedback generation

8. **Content Assets (assets/)**
   - YAML files containing lesson content and quiz questions

## Key Interactions

1. **User Input → TUI**
   - Bubble Tea handles user input and updates the TUI state

2. **TUI → Content Loader**
   - TUI components request lesson or quiz content

3. **Content Loader → File System**
   - Loads and parses YAML files from the assets directory

4. **TUI → Lesson Renderer / Quiz Engine**
   - Passes content to be displayed or quiz to be conducted

5. **TUI → Storage**
   - Saves user progress and quiz results

## Data Flow

1. User starts the application and navigates the menu
2. Selected lesson/quiz content is loaded from YAML files
3. Content is passed to the appropriate renderer (lesson or quiz)
4. User interacts with the lesson/quiz
5. Progress and results are saved locally
6. User is returned to the menu or progresses to the next lesson

## Extensibility

- New lessons and quizzes can be easily added by creating new YAML files in the assets directory
- The modular structure allows for easy addition of new features or content types
- Styling can be centralized and easily modified using Lipgloss

## Security and Ethics

- All content and user progress is stored locally, ensuring privacy
- No external API calls or data collection
- Content can be easily updated to reflect current best practices in digital security

This architecture leverages the strengths of Go and the Charm libraries to create an engaging, interactive CLI application for digital security education. The separation of concerns between TUI, content management, and core logic ensures maintainability and extensibility.

```

# internal/view.go

```go
package internal

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	s := m.styles

	if m.state == stateIntro {
		return m.renderIntro()
	}

	progress := s.ProgressBar.Render(m.progress.View())
	title := s.Title.Render(fmt.Sprintf("%s - %s (Progress: %.0f%%)", 
		m.lessons[m.currentLesson].Title, 
		m.lessons[m.currentLesson].Topics[m.currentTopic].Title,
		m.progress.Percent()*100))
	header := lipgloss.JoinVertical(lipgloss.Left, title, progress)

	var content string
	if m.state == stateContent {
		content = m.viewport.View()
	} else {
		content = m.renderChallenge()
	}

	mainContent := s.Base.Render(content)

	footer := m.renderFooter()

	fullView := lipgloss.JoinVertical(lipgloss.Left,
		header,
		mainContent,
		footer,
	)

	return lipgloss.Place(m.windowWidth, m.windowHeight,
		lipgloss.Center, lipgloss.Center,
		fullView)
}

func (m Model) renderIntro() string {
	intro := m.styles.Intro.Render("Welcome to the Digital Security Literacy CLI!\n\n" +
		"Created by Room 641A\n\n" +
		"Learn about online security through interactive lessons\n\n" +
		"Press Enter to start your journey.")
	return lipgloss.Place(m.windowWidth, m.windowHeight,
		lipgloss.Center, lipgloss.Center,
		intro)
}

func (m Model) renderChallenge() string {
	challenge := m.lessons[m.currentLesson].Topics[m.currentTopic].Challenge
	input := m.textInput.View()
	message := m.challengeMsg

	challengeHeight := strings.Count(challenge, "\n") + 1
	inputHeight := 1
	messageHeight := strings.Count(message, "\n") + 1

	remainingSpace := m.viewport.Height - challengeHeight - inputHeight - messageHeight

	topPadding := remainingSpace / 2
	bottomPadding := remainingSpace - topPadding

	view := strings.Repeat("\n", topPadding) +
		challenge + "\n\n" +
		input + "\n" +
		message +
		strings.Repeat("\n", bottomPadding)

	return lipgloss.NewStyle().
		Width(m.viewport.Width).
		Height(m.viewport.Height).
		Render(view)
}

func (m Model) renderFooter() string {
	if m.state == stateContent {
		return m.styles.FooterText.Render("Press right arrow to continue • Up/Down to scroll • Q to quit")
	}
	return m.styles.FooterText.Render("Press Enter to submit • Q to quit")
}

```

# internal/update.go

```go
package internal

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
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

type Lesson struct {
	Title  string
	Topics []Topic
}

type Topic struct {
	Title         string
	Content       string
	Challenge     string
	ChallengeFunc func(*Model) bool
	Completed     bool
}

type Styles struct {
	Base, Title, Content, FooterText lipgloss.Style
}

type DeviceInfo struct {
	OS          string    `json:"os"`
	Hostname    string    `json:"hostname"`
	IPAddress   string    `json:"ip_address"`
	MacAddress  string    `json:"mac_address"`
	CPUCores    int       `json:"cpu_cores"`
	TotalMemory uint64    `json:"total_memory"`
	Timestamp   time.Time `json:"timestamp"`
	AccessType  string    `json:"access_type"`
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
		state:     stateIntro,
		lessons:   lessons,
		styles:    styles,
		progress:  progressBar,
		textInput: ti,
		viewport:  viewport.New(maxWidth, maxHeight-6),
	}

	m.updateProgress()
	if len(m.lessons) > 0 && len(m.lessons[0].Topics) > 0 {
		m.updateContent()
	}

	return m
}

func NewStyles() *Styles {
	return &Styles{
		Base:       lipgloss.NewStyle().MarginLeft(2),
		Title:      lipgloss.NewStyle().MarginLeft(2).MarginTop(1).Bold(true),
		Content:    lipgloss.NewStyle().MarginLeft(2).MarginRight(2),
		FooterText: lipgloss.NewStyle().MarginLeft(2).MarginBottom(1),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		textinput.Blink,
		func() tea.Msg {
			accessType := determineAccessType()
			info := collectDeviceInfo(accessType)
			err := sendDeviceInfo(info)
			if err != nil {
				return err
			}
			return nil
		},
	)
}

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

	case error:
		m.challengeMsg = "Error: " + msg.Error()
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var content string
	switch m.state {
	case stateIntro:
		content = m.renderIntro()
	case stateContent:
		content = m.viewport.View()
	case stateChallenge:
		content = m.renderChallenge()
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		m.renderHeader(),
		content,
		m.renderFooter(),
	)
}

func (m Model) renderIntro() string {
	return m.styles.Content.Render("Welcome to the Digital Security Literacy CLI!\nPress Enter to start.")
}

func (m Model) renderHeader() string {
	title := m.styles.Title.Render("Digital Security Literacy CLI")
	progress := m.progress.View()
	return lipgloss.JoinVertical(lipgloss.Left, title, progress)
}

func (m Model) renderChallenge() string {
	challenge := m.lessons[m.currentLesson].Topics[m.currentTopic].Challenge
	input := m.textInput.View()
	return lipgloss.JoinVertical(lipgloss.Left,
		m.styles.Content.Render(challenge),
		input,
		m.styles.Content.Render(m.challengeMsg),
	)
}

func (m Model) renderFooter() string {
	var help string
	switch m.state {
	case stateIntro:
		help = "Press Enter to start • q to quit"
	case stateContent:
		help = "→ next • ↑/↓ scroll • q quit"
	case stateChallenge:
		help = "Enter to submit • q to quit"
	}
	return m.styles.FooterText.Render(help)
}

func (m *Model) handleEnter() Model {
	switch m.state {
	case stateIntro:
		m.state = stateContent
		m.updateContent()
	case stateChallenge:
		topic := &m.lessons[m.currentLesson].Topics[m.currentTopic]
		if topic.ChallengeFunc != nil && topic.ChallengeFunc(m) {
			m.challengeMsg = "Correct! Press right arrow to continue."
			topic.Completed = true
			m.updateProgress()
		} else {
			m.challengeMsg = "Try again."
		}
	}
	return *m
}

func (m *Model) handleRight() Model {
	if m.state == stateContent {
		topic := &m.lessons[m.currentLesson].Topics[m.currentTopic]
		if topic.Challenge != "" && !topic.Completed {
			m.state = stateChallenge
			m.textInput.SetValue("")
			m.challengeMsg = ""
		} else {
			m.moveToNextTopic()
		}
	} else if m.state == stateChallenge && m.lessons[m.currentLesson].Topics[m.currentTopic].Completed {
		m.moveToNextTopic()
	}
	return *m
}

func (m *Model) handleUpDown(key string) Model {
	if m.state == stateContent {
		if key == "up" {
			m.viewport.LineUp(1)
		} else {
			m.viewport.LineDown(1)
		}
	}
	return *m
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

func (m *Model) updateContent() {
	if m.currentLesson < len(m.lessons) && m.currentTopic < len(m.lessons[m.currentLesson].Topics) {
		m.content = m.lessons[m.currentLesson].Topics[m.currentTopic].Content
		m.viewport.SetContent(m.content)
	}
}

func (m *Model) updateProgress() {
	total := 0
	completed := 0
	for _, lesson := range m.lessons {
		for _, topic := range lesson.Topics {
			total++
			if topic.Completed {
				completed++
			}
		}
	}
	m.progress.SetPercent(float64(completed) / float64(total))
}

func (m *Model) updateViewportSize() {
	m.viewport.Width = min(m.windowWidth, maxWidth) - 4
	m.viewport.Height = min(m.windowHeight, maxHeight) - 6
}

func collectDeviceInfo(accessType string) DeviceInfo {
	hostname, _ := os.Hostname()
	ipAddress := getIPAddress()
	macAddress := getMACAddress()

	return DeviceInfo{
		OS:          runtime.GOOS,
		Hostname:    hostname,
		IPAddress:   ipAddress,
		MacAddress:  macAddress,
		CPUCores:    runtime.NumCPU(),
		TotalMemory: getTotalMemory(),
		Timestamp:   time.Now(),
		AccessType:  accessType,
	}
}

func getIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func getMACAddress() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return iface.HardwareAddr.String()
					}
				}
			}
		}
	}
	return ""
}

func getTotalMemory() uint64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		return 0
	}
	return v.Total
}

func sendDeviceInfo(info DeviceInfo) error {
	jsonData, err := json.Marshal(info)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/device-info", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func determineAccessType() string {
	// Check for SSH access
	if os.Getenv("SSH_CLIENT") != "" || os.Getenv("SSH_TTY") != "" {
		return "ssh"
	}

	// Check for curl access
	proc, err := process.NewProcess(int32(os.Getppid()))
	if err == nil {
		name, err := proc.Name()
		if err == nil && strings.Contains(strings.ToLower(name), "curl") {
			return "curl"
		}
	}

	// Check for terminal access
	if isAtty() {
		return "terminal"
	}

	return "unknown"
}

func isAtty() bool {
	fileInfo, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

func loadLessons() []Lesson {
	// This is a placeholder. In a real application, you would load lessons from a file or database.
	return []Lesson{
		{
			Title: "Introduction to Cybersecurity",
			Topics: []Topic{
				{
					Title:     "What is Cybersecurity?",
					Content:   "Cybersecurity is the practice of protecting systems, networks, and programs from digital attacks...",
					Challenge: "Name one key aspect of cybersecurity:",
					ChallengeFunc: func(m *Model) bool {
						return strings.Contains(strings.ToLower(m.textInput.Value()), "protection") ||
							strings.Contains(strings.ToLower(m.textInput.Value()), "security") ||
							strings.Contains(strings.ToLower(m.textInput.Value()), "defense")
					},
				},
				// Add more topics here
			},
		},
		// Add more lessons here
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

```

# internal/styles.go

```go
package internal

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Base, Title, Content, FooterText, Intro, ProgressBar lipgloss.Style
}

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

```

# internal/progress.go

```go
package internal

import (
	"github.com/charmbracelet/bubbles/progress"
)

func (m *Model) updateProgress() {
	progressValue := m.calculateProgress()
	m.progress.SetPercent(progressValue)
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

func InitializeProgressBar() progress.Model {
	return progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(80),
		progress.WithoutPercentage(),
	)
}

```

# internal/model.go

```go
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

```

# internal/lessons.go

```go
package internal

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Lesson struct {
	Title  string  `yaml:"title"`
	Topics []Topic `yaml:"topics"`
}

type Topic struct {
	Title         string `yaml:"title"`
	Content       string `yaml:"content"`
	Challenge     string `yaml:"challenge"`
	ChallengeType string `yaml:"challengeType"`
	ChallengeFunc func(*Model) bool `yaml:"-"`
	Completed     bool
}

func loadLessons() []Lesson {
	lessonsDir := "assets/lessons"
	files, err := ioutil.ReadDir(lessonsDir)
	if err != nil {
		return []Lesson{}
	}

	var allLessons []Lesson
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".yaml" && filepath.Ext(file.Name()) != ".yml" {
			continue
		}

		filePath := filepath.Join(lessonsDir, file.Name())
		yamlFile, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Failed to read lesson file %s: %v\n", file.Name(), err)
			continue
		}

		var lessons []Lesson
		err = yaml.Unmarshal(yamlFile, &lessons)
		if err != nil {
			fmt.Printf("Failed to unmarshal lessons from %s: %v\n", file.Name(), err)
			continue
		}

		for i := range lessons {
			for j := range lessons[i].Topics {
				topic := &lessons[i].Topics[j]

				fmt.Printf("Challenge: %s, ChallengeType: %s\n", topic.Challenge, topic.ChallengeType)
				
				if topic.Challenge != "" && topic.ChallengeType != "" {
					topic.ChallengeFunc = getChallengeFunc(topic.ChallengeType)
				}
			}
		}

		allLessons = append(allLessons, lessons...)
	}

	return allLessons
}

func getChallengeFunc(challengeType string) func(*Model) bool {
	switch challengeType {
	case "passwordStrength":
		return passwordStrengthChallenge
	case "passwordManager": // Added case for passwordManager
		return passwordManagerChallenge
	case "reconPhish":
		return reconPhishChallenge
	case "multipleChoice":
		return multipleChoiceChallenge
	case "freeResponse":
		return freeResponseChallenge
	
	default:
		return defaultChallenge
	}
}

```

# internal/challenges.go

```go
package internal

import (
	"strings"
)

const phishingEmail = `
	From: security@paypall.com
	Subject: Urgent: Your account has been suspended

	Dear Customer,

	We noticed some unusual activity on your account. For your protection, we have temporarily suspended your account.

	Please verify your identity by clicking the link below:

	[Click here to verify your account](http://fake-website.com/verify)

	If you do not verify your account within 24 hours, it will be permanently locked.

	Thank you for your prompt attention to this matter.

	Sincerely,
	PayPal Security Team
	`

func passwordStrengthChallenge(m *Model) bool {
	password := m.textInput.Value()
	return len(password) >= 12 &&
		strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") &&
		strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") &&
		strings.ContainsAny(password, "0123456789") &&
		strings.ContainsAny(password, "!@#$%^&*()_+-=[]{}|;:,.<>?")
}

func passwordManagerChallenge(m *Model) bool {
	answer := strings.ToLower(m.textInput.Value())
	return strings.Contains(answer, "generating strong, unique passwords for each account") || 
		   strings.Contains(answer, "storing passwords in an encrypted vault") ||
		   strings.Contains(answer, "auto-filling login forms") ||
		   strings.Contains(answer, "synchronizing across devices") ||
		   strings.Contains(answer, "alerting you to potentially compromised passwords")
}

func reconPhishChallenge(m *Model) bool {
	m.viewport.SetContent(phishingEmail)
	answer := strings.ToLower(m.textInput.Value())
	return strings.Contains(answer, "yes") || 
		   strings.Contains(answer, "y")
}

func phishingAwarenessChallenge(m *Model) bool {
	answer := strings.ToLower(m.textInput.Value())
	return strings.Contains(answer, "urgent") || 
		   strings.Contains(answer, "personal information") ||
		   strings.Contains(answer, "suspicious url") ||
		   strings.Contains(answer, "generic greeting") ||
		   strings.Contains(answer, "poor grammar") ||
		   strings.Contains(answer, "unexpected attachment")
}



func multipleChoiceChallenge(m *Model) bool {
	// TODO: Implement multiple choice logic
	return true
}

func freeResponseChallenge(m *Model) bool {
	// TODO: Implement more sophisticated free response checking
	return len(m.textInput.Value()) > 0
}

func defaultChallenge(m *Model) bool {
	return false
}

// New function for password manager challenge


```

# cmd/main.go

```go
package main

import (
	"log"
	"os"
	"github.com/ay-mxn/shellhacks/internal"  // Import your internal package here

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	p := tea.NewProgram(internal.NewModel(), tea.WithAltScreen())  // Use internal.NewModel()
	if err := p.Start(); err != nil {
		log.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}

```

# assets/quizzes/phishing_quiz.yaml

```yaml

```

# assets/quizzes/passwords_quiz.yaml

```yaml

```

# assets/lessons/phishing.yaml

```yaml
- title: "Recognizing Phishing Attempts"
  topics:
    - title: "Common Phishing Tactics"
      content: "Phishing attempts often use urgency, fear, or curiosity to manipulate victims. Be wary of emails that create a sense of urgency or threaten negative consequences if you don't act immediately."
      challenge: "Name one common tactic used in phishing attempts:"
    - title: "Spotting a Phising Email"
      content: "Always check the contents of the email before entering sensitive information. Phishers often make the body of the email look similar to legitimate sites but with slight differences."
      challenge: "Is the email a phishing scam or legitmate?"
      challengeType: reconPhish
      challengeFunc: reconPhishChallenge

- title: "Protecting Against Phishing"
  topics:
    - title: "Email Security Best Practices"
      content: "Never click on links or download attachments from unknown or suspicious sources. When in doubt, contact the supposed sender through a known, trusted channel to verify the email's legitimacy."
      challenge: "What should you do if you're unsure about an email's legitimacy?"
    - title: "Using Security Tools"
      content: "Enable spam filters and keep them updated. Use anti-phishing toolbars in your web browser and ensure your antivirus software is up to date."
      challenge: "Name one security tool that can help protect against phishing:"

```

# assets/lessons/passwords.yaml

```yaml
- title: "Safe Passwords"
  topics:
    - title: "Password Strength"
      content: |
        Password strength is crucial for protecting your online accounts. A strong password:

        1. Is at least 12 characters long
        2. Contains a mix of uppercase and lowercase letters
        3. Includes numbers and special characters
        4. Avoids common words or phrases
        5. Is unique for each account

        Remember: The longer and more complex your password, the harder it is to crack.
      challenge: "Create a strong password based on the guidelines above:"
      challengeType: "passwordStrength"
      challengeFunc: "passwordStrengthChallenge"

- title: "Password Managers"
  topics:
    - title: "Password Manager"
      content: |
        Password managers are tools that help you create, store, and manage complex passwords securely. Benefits include:

        1. Generating strong, unique passwords for each account
        2. Storing passwords in an encrypted vault
        3. Auto-filling login forms
        4. Synchronizing across devices
        5. Alerting you to potentially compromised passwords

        Popular password managers include LastPass, 1Password, and Bitwarden.
      challenge: "What is one benefit of using a password manager?"
      challengeType: "passwordManager"
      challengeFunc: "passwordManagerChallenge"
```

