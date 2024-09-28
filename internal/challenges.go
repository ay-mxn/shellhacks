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

