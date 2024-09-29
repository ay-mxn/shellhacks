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
