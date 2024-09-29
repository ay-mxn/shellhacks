package ui

import (
	"gopkg.in/yaml.v2"
	"os"
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
	// Load lessons from YAML file
	yamlFile, err := os.ReadFile("assets/lessons/sample_lessons.yaml")
	if err != nil {
		panic(err)
	}

	var lessons []Lesson
	err = yaml.Unmarshal(yamlFile, &lessons)
	if err != nil {
		panic(err)
	}

	return lessons
}
