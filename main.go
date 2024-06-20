package main

import (
	"core-engine/helpers"
	"core-engine/models"
	"log"
	"path/filepath"
)

func main() {
	// Define the path to the YAML file
	yamlFile := filepath.Join("sample-pipeline.yaml")

	// Read and unmarshal the YAML file
	var ci models.UnifiedCI
	err := helpers.ReadYAML(yamlFile, &ci)
	if err != nil {
		log.Fatalf("Error reading and unmarshaling YAML file: %v", err)
	}

	for _, job := range ci.Jobs {
		err := helpers.ExecuteJob(job, ci.Variables)
		if err != nil {
			log.Fatalf("Error executing job: %v", err)
		}
	}
}
