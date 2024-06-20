package helpers

import (
	"core-engine/models"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"os/exec"
)

// ReadYAML reads and unmarshals the YAML file into the given struct
func ReadYAML(filePath string, out interface{}) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, out)
	if err != nil {
		return err
	}

	return nil
}

// ExecuteStep executes a single step
func ExecuteStep(step models.Step) error {
	fmt.Printf("Executing step: %s\n", step.Name)
	cmd := exec.Command("sh", "-c", step.Run)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing step '%s': %v, output: %s", step.Name, err, string(output))
	}
	fmt.Printf("Output: %s\n", string(output))
	return nil
}

// ExecuteJob executes all steps in a job
func ExecuteJob(job models.Job) error {
	fmt.Printf("Executing job: %s\n", job.Name)
	for _, step := range job.Steps {
		err := ExecuteStep(step)
		if err != nil {
			return err
		}
	}
	return nil
}
