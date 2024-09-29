# view.go

```go
package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	s := m.styles

	header := m.appBoundaryView("Security Quiz")

	// Form (left side)
	formView := m.form.View()

	// Status (right side)
	status := s.Status.Render(
		s.StatusHeader.Render("Current Lesson") + "\n" +
			m.lessons[m.currentLessonIndex].Title + "\n\n" +
			s.StatusHeader.Render("Current Topic") + "\n" +
			m.lessons[m.currentLessonIndex].Topics[m.currentTopicIndex].Title,
	)

	body := lipgloss.JoinHorizontal(lipgloss.Top, formView, status)

	progressBar := m.renderProgressBar()

	footer := s.Help.Render("q: quit • enter: submit")

	return s.Base.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			header,
			body,
			footer,
			progressBar,
		),
	)
}

func (m Model) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(indigo),
	)
}

func (m Model) renderProgressBar() string {
	return m.styles.ProgressBar.Render(m.progress.View())
}

```

# update.go

```go
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

```

# styles.go

```go
package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green  = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
)

type Styles struct {
	Base, HeaderText, Status, StatusHeader, Highlight, Help, ProgressBar lipgloss.Style
}

func NewStyles() *Styles {
	s := &Styles{}
	s.Base = lipgloss.NewStyle().Padding(1, 4, 0, 1)
	s.HeaderText = lipgloss.NewStyle().Foreground(indigo).Bold(true).Padding(0, 1, 0, 2)
	s.Status = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(indigo).PaddingLeft(1).MarginTop(1)
	s.StatusHeader = lipgloss.NewStyle().Foreground(green).Bold(true)
	s.Highlight = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	s.Help = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	s.ProgressBar = lipgloss.NewStyle().Padding(0, 0, 1, 0)
	return s
}

```

# model.go

```go
package main

import (
	"github.com/charmbracelet/huh"
)

func buildForm(m *Model) *huh.Form {
	if m.quizCompleted {
		return buildCompletionForm(m)
	}

	lesson := m.lessons[m.currentLessonIndex]
	topic := lesson.Topics[m.currentTopicIndex]

	return huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title(lesson.Title+" - "+topic.Title).
				Description(topic.Content),
			huh.NewInput().
				Title("Challenge: "+topic.Challenge).
				Placeholder("Type your answer here").
				Validate(func(s string) error {
					return validateAnswer(m, s)
				}),
		),
	).WithWidth(60).WithShowHelp(false)
}

func buildCompletionForm(m *Model) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("Quiz Completed!").
				Description("Congratulations! You've completed the security quiz."),
			huh.NewConfirm().
				Title("Would you like to restart the quiz?").
				Affirmative("Yes").
				Negative("No"),
		),
	).WithWidth(60).WithShowHelp(false)
}

func validateAnswer(m *Model, answer string) error {
	// Implement your answer validation logic here
	m.lessons[m.currentLessonIndex].Topics[m.currentTopicIndex].Completed = true
	return nil
}

```

# main.go

```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v2"
	"github.com/ay-mxn/shellhacks/internal"
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
	func main() {
	// Collect and send device info
	err := internal.CollectAndSendDeviceInfo()
	if err != nil {
		log.Printf("Failed to collect and send device info: %v", err)
		// Note: We're continuing with the program even if this fails
	}

	p := tea.NewProgram(NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
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

	// Collect and send device info
	log.Println("Collecting and sending device info...")
	err = internal.CollectAndSendDeviceInfo()
	if err != nil {
		log.Printf("Failed to collect and send device info: %v", err)
		// Note: We're continuing with the program even if this fails
	}

	log.Println("Starting quiz application...")
	p := tea.NewProgram(NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Printf("Error running program: %v", err)
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}


```

# lessons.go

```go
package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

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

func loadLessons() []Lesson {
	// This is a placeholder implementation. Replace with actual YAML parsing logic.
	lessons := []Lesson{
		{
			Title: "Sample Lesson",
			Topics: []Topic{
				{
					Title:     "Sample Topic",
					Content:   "This is sample content.",
					Challenge: "This is a sample challenge.",
				},
			},
		},
	}

	// Here you would typically load lessons from YAML files
	// If no lessons are loaded, we'll use the placeholder lesson

	return lessons
}

```

# go.mod

```mod
module github.com/ay-mxn/shellhacks

go 1.22.3

require (
	github.com/charmbracelet/bubbles v0.20.0
	github.com/charmbracelet/bubbletea v1.1.1
	github.com/charmbracelet/huh v0.6.0
	github.com/charmbracelet/lipgloss v0.13.0
	github.com/mattn/go-sqlite3 v1.14.23
	github.com/shirou/gopsutil/v3 v3.24.5
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/catppuccin/go v0.2.0 // indirect
	github.com/charmbracelet/harmonica v0.2.0 // indirect
	github.com/charmbracelet/x/ansi v0.2.3 // indirect
	github.com/charmbracelet/x/exp/strings v0.0.0-20240722160745-212f7b056ed0 // indirect
	github.com/charmbracelet/x/term v0.2.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/erikgeiser/coninput v0.0.0-20211004153227-1c3628e74d0f // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-localereader v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mitchellh/hashstructure/v2 v2.0.2 // indirect
	github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6 // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/muesli/termenv v0.15.3-0.20240618155329-98d742f6907a // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
)

```

# debug.log

```log

```

# internal/server.go

```go
package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mattn/go-sqlite3"
)

type Server struct {
	db *SQLiteDB
}

func NewServer() (*Server, error) {
	log.Println("Initializing server...")
	db, err := NewSQLiteDB("./system_info.db")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %v", err)
	}
	log.Println("Database initialized successfully")
	return &Server{db: db}, nil
}

func (s *Server) Start() {
	log.Println("Setting up HTTP routes...")
	http.HandleFunc("/beacon", s.logRequest(s.BeaconHandler))
	http.HandleFunc("/fetch", s.logRequest(s.FetchHandler))

	addr := ":8080"
	log.Printf("Server is starting on http://localhost%s", addr)
	go func() {
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
}

func (s *Server) BeaconHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var info DeviceInfo
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	info.LastBeaconTime = time.Now()

	err = s.db.SaveSystemInfo(info)
	if err != nil {
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	log.Printf("Received and stored beacon from device: %s", info.ID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data received and stored successfully"))
}

func (s *Server) FetchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	info, err := s.db.GetSystemInfo(id)
	if err != nil {
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		return
	}

	if info == nil {
		http.Error(w, "ID not found", http.StatusNotFound)
		return
	}

	log.Printf("Fetched data for device: %s", id)
	json.NewEncoder(w).Encode(info)
}

func (s *Server) logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("Received %s request for %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s request for %s in %v", r.Method, r.URL.Path, time.Since(startTime))
	}
}

type SQLiteDB struct {
	*sqlite3.SQLiteConn
}

func NewSQLiteDB(dbPath string) (*SQLiteDB, error) {
	log.Printf("Opening SQLite database at %s", dbPath)
	conn, err := sqlite3.OpenConn(dbPath, sqlite3.SQLITE_OPEN_READWRITE|sqlite3.SQLITE_OPEN_CREATE)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	if err := createTable(conn); err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	log.Println("SQLite database initialized successfully")
	return &SQLiteDB{SQLiteConn: conn}, nil
}

func createTable(conn *sqlite3.SQLiteConn) error {
	log.Println("Creating system_info table if it doesn't exist")
	_, err := conn.Exec(`
		CREATE TABLE IF NOT EXISTS system_info (
			id TEXT PRIMARY KEY,
			username TEXT,
			os TEXT,
			ram_total INTEGER,
			cpu_cores INTEGER,
			last_beacon_time DATETIME
		)
	`, nil)
	return err
}

func (db *SQLiteDB) SaveSystemInfo(info DeviceInfo) error {
	log.Printf("Saving system info for ID: %s", info.ID)
	_, err := db.Exec(`
		INSERT OR REPLACE INTO system_info 
		(id, username, os, ram_total, cpu_cores, last_beacon_time) 
		VALUES (?, ?, ?, ?, ?, ?)
	`, info.ID, info.Username, info.OS, info.RAMTotal, info.CPUCores, info.LastBeaconTime)
	if err != nil {
		log.Printf("Error saving system info: %v", err)
	}
	return err
}

func (db *SQLiteDB) GetSystemInfo(id string) (*DeviceInfo, error) {
	log.Printf("Fetching system info for ID: %s", id)
	var info DeviceInfo
	err := db.QueryRow(`
		SELECT id, username, os, ram_total, cpu_cores, last_beacon_time 
		FROM system_info WHERE id = ?
	`, id).Scan(&info.ID, &info.Username, &info.OS, &info.RAMTotal, &info.CPUCores, &info.LastBeaconTime)

	if err == sqlite3.ErrNoRows {
		log.Printf("No system info found for ID: %s", id)
		return nil, nil
	} else if err != nil {
		log.Printf("Error fetching system info: %v", err)
		return nil, err
	}

	log.Printf("Successfully fetched system info for ID: %s", id)
	return &info, nil
}

```

# internal/device_info.go

```go
package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type DeviceInfo struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	OS             string    `json:"os"`
	RAMTotal       int       `json:"ram_total"`
	CPUCores       int       `json:"cpu_cores"`
	LastBeaconTime time.Time `json:"last_beacon_time"`
}

func CollectAndSendDeviceInfo() error {
	log.Println("Collecting device information...")
	info, err := collectDeviceInfo()
	if err != nil {
		return fmt.Errorf("failed to collect device info: %v", err)
	}

	log.Println("Sending device information to server...")
	return sendDeviceInfo(info)
}

func collectDeviceInfo() (DeviceInfo, error) {
	username := os.Getenv("USER")
	if username == "" {
		username = os.Getenv("USERNAME")
	}

	v, err := mem.VirtualMemory()
	if err != nil {
		return DeviceInfo{}, fmt.Errorf("failed to get virtual memory info: %v", err)
	}

	h, err := host.Info()
	if err != nil {
		return DeviceInfo{}, fmt.Errorf("failed to get host info: %v", err)
	}

	log.Printf("Collected device info: ID=%s, Username=%s, OS=%s, RAM=%d MB, CPUs=%d",
		h.HostID, username, runtime.GOOS, v.Total/1024/1024, runtime.NumCPU())

	return DeviceInfo{
		ID:             h.HostID,
		Username:       username,
		OS:             runtime.GOOS,
		RAMTotal:       int(v.Total / 1024 / 1024), // Convert to MB
		CPUCores:       runtime.NumCPU(),
		LastBeaconTime: time.Now(),
	}, nil
}

func sendDeviceInfo(info DeviceInfo) error {
	jsonData, err := json.Marshal(info)
	if err != nil {
		return fmt.Errorf("failed to marshal device info: %v", err)
	}

	resp, err := http.Post("http://localhost:8080/beacon", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send device info: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned non-OK status: %s", resp.Status)
	}

	log.Println("Device information sent successfully")
	return nil
}

```

# assets/lessons/sample_lessons.yaml

```yaml
- title: "Password Security"
  topics:
    - title: "Creating Strong Passwords"
      content: |
        Strong passwords are essential for protecting your online accounts. Here are some tips for creating strong passwords:
        1. Use a mix of uppercase and lowercase letters, numbers, and symbols.
        2. Make your password at least 12 characters long.
        3. Avoid using personal information or common words.
        4. Use a unique password for each account.
      challenge: "Create a strong password following the guidelines above."

    - title: "Password Managers"
      content: |
        Password managers are tools that help you create, store, and manage complex passwords securely. Benefits include:
        - Generating strong, unique passwords for each account
        - Securely storing your passwords in an encrypted vault
        - Auto-filling login forms for convenience
        - Syncing across multiple devices
      challenge: "Name two benefits of using a password manager."

    - title: "Two-Factor Authentication"
      content: |
        Two-factor authentication (2FA) adds an extra layer of security to your accounts. It requires two different forms of identification:
        1. Something you know (like a password)
        2. Something you have (like a phone or security key)
        This makes it much harder for attackers to gain unauthorized access to your accounts.
      challenge: "Explain why two-factor authentication is more secure than just using a password."

- title: "Phishing Awareness"
  topics:
    - title: "Recognizing Phishing Emails"
      content: |
        Phishing emails are attempts to trick you into revealing sensitive information. Common signs of phishing include:
        - Urgent or threatening language
        - Requests for personal information
        - Suspicious attachments or links
        - Spelling and grammar errors
        - Generic greetings
      challenge: "List three signs that an email might be a phishing attempt."

    - title: "URL Safety"
      content: |
        Before clicking on links in emails or messages, it's important to verify the URL. Here are some tips:
        - Hover over links to see the actual destination
        - Check for subtle misspellings in domain names
        - Look for 'https://' at the beginning of the URL
        - Be wary of shortened URLs
      challenge: "What should you do before clicking on a link in an email?"

    - title: "Reporting Phishing Attempts"
      content: |
        If you receive a suspected phishing email, it's important to report it. This helps protect others and improves email security. Steps to report phishing:
        1. Don't click on any links or download attachments
        2. Forward the email to your IT department or security team
        3. Delete the email from your inbox
      challenge: "Describe the steps you should take if you receive a suspected phishing email."

- title: "Safe Browsing Habits"
  topics:
    - title: "HTTPS Importance"
      content: |
        HTTPS (Hypertext Transfer Protocol Secure) is crucial for secure internet browsing. It provides:
        - Encryption of data transmitted between your browser and websites
        - Authentication of website identity
        - Protection against man-in-the-middle attacks
        Always look for the padlock icon in your browser's address bar to ensure you're on a secure connection.
      challenge: "Explain why it's important to use HTTPS when browsing the internet."

    - title: "Public Wi-Fi Safety"
      content: |
        Using public Wi-Fi networks can be risky. To stay safe:
        - Avoid accessing sensitive information (e.g., online banking)
        - Use a VPN to encrypt your traffic
        - Ensure your device's firewall is active
        - Disable file sharing and automatic connections to Wi-Fi networks
      challenge: "Name two precautions you should take when using public Wi-Fi."

    - title: "Browser Extensions Security"
      content: |
        While browser extensions can be useful, they can also pose security risks. To use extensions safely:
        - Only install extensions from official web stores
        - Read reviews and check permissions before installing
        - Regularly review and remove unused extensions
        - Keep extensions updated
      challenge: "What should you check before installing a browser extension?"

- title: "Device Security"
  topics:
    - title: "Operating System Updates"
      content: |
        Keeping your operating system up-to-date is crucial for security. Updates often include:
        - Patches for known vulnerabilities
        - Improvements to system stability
        - New security features
        Enable automatic updates whenever possible to ensure you're always protected.
      challenge: "Why is it important to keep your operating system updated?"

    - title: "Antivirus Software"
      content: |
        Antivirus software helps protect your device from malware. Key features include:
        - Real-time scanning of files and programs
        - Web protection against malicious websites
        - Automatic updates to detect new threats
        - Scheduled full system scans
        Choose a reputable antivirus program and keep it updated.
      challenge: "List two key features of antivirus software."

    - title: "Device Encryption"
      content: |
        Device encryption protects your data if your device is lost or stolen. It scrambles the data on your device, making it unreadable without the correct key or password. Many modern operating systems offer built-in encryption options:
        - Windows: BitLocker
        - macOS: FileVault
        - Android and iOS: Enabled by default on most devices
      challenge: "Explain the purpose of device encryption and name one encryption tool."

- title: "Social Media Privacy"
  topics:
    - title: "Privacy Settings"
      content: |
        Managing your social media privacy settings is crucial. Key steps include:
        - Reviewing and adjusting who can see your posts and personal information
        - Limiting the information visible on your public profile
        - Controlling which apps and websites have access to your account
        - Using privacy checkup tools provided by the platform
      challenge: "Name three things you should review in your social media privacy settings."

    - title: "Safe Sharing Practices"
      content: |
        Be mindful of what you share on social media:
        - Avoid posting sensitive personal information (e.g., address, phone number)
        - Think twice before sharing your location
        - Be cautious about sharing photos that might reveal too much about your daily routines
        - Consider the potential long-term consequences of your posts
      challenge: "List two types of information you should avoid sharing on social media."

    - title: "Recognizing Social Engineering"
      content: |
        Social engineering attacks often target social media users. Be aware of:
        - Friend requests from unknown individuals
        - Messages asking for personal information or money
        - Contests or giveaways that seem too good to be true
        - Pressure to act quickly on unexpected offers
        Always verify the identity of individuals or organizations contacting you through social media.
      challenge: "Describe two signs that a social media interaction might be a social engineering attempt."

```

