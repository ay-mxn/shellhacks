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
