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
