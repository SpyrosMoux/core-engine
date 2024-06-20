package models

// Job represents an individual job in the CI/CD pipeline
type Job struct {
	Name        string   `yaml:"name"`
	Runner      string   `yaml:"runner"`
	Environment string   `yaml:"environment"`
	Needs       []string `yaml:"needs,omitempty"`
	Steps       []Step   `yaml:"steps"`
}

// Step represents an individual step within a job
type Step struct {
	Name string `yaml:"name"`
	Run  string `yaml:"run"`
}

// UnifiedCI represents the top-level structure containing jobs
type UnifiedCI struct {
	Jobs []Job `yaml:"jobs"`
}
