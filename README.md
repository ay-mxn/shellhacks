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
