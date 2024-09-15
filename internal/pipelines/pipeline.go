package pipelines

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spyrosmoux/core-engine/internal/logger"

	"github.com/spyrosmoux/core-engine/pkg/models"
)

// ExecuteStep executes a single step
func ExecuteStep(step models.Step, variables map[string]string) error {
	logger.Log(logger.InfoLevel, "Executing step: "+step.Name)
	command := SubstituteVariables(step.Run, variables)
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing step '%s': %v, output: %s", step.Name, err, string(output))
	}
	logger.Log(logger.InfoLevel, "Output: "+string(output))
	return nil
}

// ExecuteJob executes all steps in a job
func ExecuteJob(job models.Job, variables map[string]string) error {
	logger.Log(logger.InfoLevel, "Executing job: "+job.Name)
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

// PrepareRun creates a _work directory which the job will use as context
func PrepareRun() {
	// TODO(spyrosmoux) create a unique dir for the run based on the unique build number (build number -> generated by the API)
	err := os.Mkdir("_work", 0755)
	if err != nil {
		logger.Log(logger.FatalLevel, "Error creating temporary directory: "+err.Error())
	}

	err = os.Chdir("_work")
	if err != nil {
		logger.Log(logger.FatalLevel, "Error switching context: "+err.Error())
	}
}

// CleanupRun deletes the '_work' directory and all its contents
func CleanupRun() {
	err := os.Chdir("../")
	if err != nil {
		logger.Log(logger.FatalLevel, "Error switching context: "+err.Error())
	}

	err = os.RemoveAll("_work")
	if err != nil {
		logger.Log(logger.FatalLevel, "Error removing temporary directory '_work': "+err.Error())
	}
	logger.Log(logger.InfoLevel, "Temporary directory '_work' removed successfully")
}

// RunJob prepares, executes and cleans-up a run
func RunPipeline(job string) error {
	ci, err := models.ValidateYAMLStructure([]byte(job))
	if err != nil {
		logger.Log(logger.ErrorLevel, "Error validating yaml structure with error: "+err.Error())
		return err
	}

	PrepareRun()

	for _, job := range ci.Jobs {
		logger.Log(logger.InfoLevel, "Running job: "+job.Name)
		err := ExecuteJob(job, ci.Variables)
		if err != nil {
			CleanupRun()
			logger.Log(logger.ErrorLevel, "Error executing job: "+err.Error())
			return err
		}
	}

	CleanupRun()

	return nil
}
