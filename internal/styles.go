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
