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
