package models

import "gopkg.in/yaml.v2"

// Job represents an individual job in the CI/CD pipeline
type Job struct {
	Name  string   `yaml:"name"`
	Needs []string `yaml:"needs,omitempty"`
	Steps []Step   `yaml:"steps"`
}

// Step represents an individual step within a job
type Step struct {
	Name string `yaml:"name"`
	Run  string `yaml:"run"`
}

// UnifiedCI represents the top-level structure containing jobs
type UnifiedCI struct {
	Variables map[string]string `yaml:"variables,omitempty"`
	Jobs      []Job             `yaml:"jobs"`
}

func ReadYAMLFromString(yamlData string, out UnifiedCI) error {
	return yaml.Unmarshal([]byte(yamlData), out)
}
