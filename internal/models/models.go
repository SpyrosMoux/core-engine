package models

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