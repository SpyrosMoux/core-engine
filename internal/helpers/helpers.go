package helpers

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"os/exec"
	"spyrosmoux/core-engine/internal/models"
	"strings"
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
func ExecuteStep(step models.Step, variables map[string]string) error {
	fmt.Printf("Executing step: %s\n", step.Name)
	command := SubstituteVariables(step.Run, variables)
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing step '%s': %v, output: %s", step.Name, err, string(output))
	}
	fmt.Printf("Output: %s\n", string(output))
	return nil
}

// ExecuteJob executes all steps in a job
func ExecuteJob(job models.Job, variables map[string]string) error {
	fmt.Printf("Executing job: %s\n", job.Name)
	for _, step := range job.Steps {
		err := ExecuteStep(step, variables)
		if err != nil {
			return err
		}
	}
	return nil
}

// SubstituteVariables substitutes variables in the command
func SubstituteVariables(command string, variables map[string]string) string {
	for key, value := range variables {
		placeholder := fmt.Sprintf("${%s}", key)
		command = strings.ReplaceAll(command, placeholder, value)
	}
	return command
}

func ReadYAMLFromString(yamlData string, out interface{}) error {
	return yaml.Unmarshal([]byte(yamlData), out)
}
