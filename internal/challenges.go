package internal

import (
	"strings"
)

func passwordStrengthChallenge(m *Model) bool {
	password := m.textInput.Value()
	return len(password) >= 12 &&
		strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") &&
		strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") &&
		strings.ContainsAny(password, "0123456789") &&
		strings.ContainsAny(password, "!@#$%^&*()_+-=[]{}|;:,.<>?")
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
