package tools

import (
	"gopkg.in/yaml.v3"
	"os"
)

var RequirementsValue *Requirements

func init() {
	requirements, err := ReadInfoFromYamlFile()
	if err != nil {
		panic(err)
	}
	RequirementsValue = requirements
}

func ReadInfoFromYamlFile() (*Requirements, error) {
	// Load the file
	f, err := os.ReadFile("requirements.yaml")
	if err != nil {
		return nil, err
	}
	info := &Requirements{}
	if err := yaml.Unmarshal(f, &info); err != nil {
		return nil, err
	}
	return info, nil
}
