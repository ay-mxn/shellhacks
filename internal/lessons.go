package internal

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Lesson struct {
	Title  string  `yaml:"title"`
	Topics []Topic `yaml:"topics"`
}

type Topic struct {
	Title         string `yaml:"title"`
	Content       string `yaml:"content"`
	Challenge     string `yaml:"challenge"`
	ChallengeType string `yaml:"challengeType"`
	ChallengeFunc func(*Model) bool `yaml:"-"`
	Completed     bool
}

func loadLessons() []Lesson {
	lessonsDir := "assets/lessons"
	files, err := ioutil.ReadDir(lessonsDir)
	if err != nil {
		return []Lesson{}
	}

	var allLessons []Lesson
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".yaml" && filepath.Ext(file.Name()) != ".yml" {
			continue
		}

		filePath := filepath.Join(lessonsDir, file.Name())
		yamlFile, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Failed to read lesson file %s: %v\n", file.Name(), err)
			continue
		}

		var lessons []Lesson
		err = yaml.Unmarshal(yamlFile, &lessons)
		if err != nil {
			fmt.Printf("Failed to unmarshal lessons from %s: %v\n", file.Name(), err)
			continue
		}

		for i := range lessons {
			for j := range lessons[i].Topics {
				topic := &lessons[i].Topics[j]
				
				if topic.Challenge != "" && topic.ChallengeType != "" {
					topic.ChallengeFunc = getChallengeFunc(topic.ChallengeType)
				}
			}
		}

		allLessons = append(allLessons, lessons...)
	}

	return allLessons
}

func getChallengeFunc(challengeType string) func(*Model) bool {
	switch challengeType {
	case "passwordStrength":
		return passwordStrengthChallenge
	case "multipleChoice":
		return multipleChoiceChallenge
	case "freeResponse":
		return freeResponseChallenge
	default:
		return defaultChallenge
	}
}
