package internal

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

// CheckCompletion function checks if all lessons are completed and runs a placeholder function
func CheckCompletion(m *Model) {
    completedLessons := 0
    for _, lesson := range m.lessons {
        lessonCompleted := true
        for _, topic := range lesson.Topics {
            if !topic.Completed {
                lessonCompleted = false
                break
            }
        }
        if lessonCompleted {
            completedLessons++
        }
    }
    if completedLessons == len(m.lessons) {
        RunPlaceholderFunction()
    }
}

// RunPlaceholderFunction is a placeholder for the function to be implemented later
func RunPlaceholderFunction() {
    fmt.Println("Placeholder function. Implement this later with the actual functionality.")
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
            CheckCompletion(&m) // Added completion check here
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
        if topic.Completed {
            m.moveToNextTopic()
            m.updateProgressOnSlideChange()
            CheckCompletion(&m) // Added completion check here
        }
    }
    return m, nil
}
