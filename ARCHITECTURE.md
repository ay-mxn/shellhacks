# Digital Security Literacy CLI Architecture

## Tech Stack
- Language: Go
- CLI Framework: [Bubble Tea](https://github.com/charmbracelet/bubbletea) (for TUI)
- Styling: [Lipgloss](https://github.com/charmbracelet/lipgloss)
- Input Handling: [Bubbles](https://github.com/charmbracelet/bubbles)
- Storage: Local file system for progress and settings
- Content Management: YAML files for lesson content and quizzes

## Project Structure
```
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
```

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
