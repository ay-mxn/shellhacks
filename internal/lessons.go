package internal

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

//go:embed assets/lessons/*.yaml
var lessonsFS embed.FS

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
	var allLessons []Lesson

	err := fs.WalkDir(lessonsFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error walking directory: %v\n", err)
			return err
		}


		if filepath.Ext(d.Name()) != ".yaml" && filepath.Ext(d.Name()) != ".yml" {
			return nil
		}

		yamlFile, err := lessonsFS.ReadFile(path)
		if err != nil {
			fmt.Printf("Failed to read lesson file %s: %v\n", path, err)
			return nil
		}

		var lessons []Lesson
		err = yaml.Unmarshal(yamlFile, &lessons)
		if err != nil {
			fmt.Printf("Failed to unmarshal lessons from %s: %v\n", path, err)
			return nil
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
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through embedded files: %v\n", err)
	}

	return allLessons
}

func getChallengeFunc(challengeType string) func(*Model) bool {
	switch challengeType {
	case "passwordStrength":
		return passwordStrengthChallenge
	case "passwordManager":
		return passwordManagerChallenge
	case "reconPhish":
		return reconPhishChallenge
	case "phishingAwareness":
		return phishingAwarenessChallenge
	case "dataCollection":
		return dataCollectionChallenge
	case "socialMediaPrivacy":
		return socialMediaPrivacyChallenge
	case "vpnUsage":
		return vpnUsageChallenge
	case "httpsImportance":
		return httpsImportanceChallenge
	case "fakeSite":
		return fakeSiteChallenge
	case "browserSettings":
		return browserSettingsChallenge
	case "deviceEncryption":
		return deviceEncryptionChallenge
	case "softwareUpdate":
		return softwareUpdateChallenge
	case "secureWifi":
		return secureWifiChallenge
	default:
		return nil
	}
}
