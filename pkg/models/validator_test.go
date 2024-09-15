package models

import (
	"os"
	"testing"
)

func TestValidator(t *testing.T) {
	yamlData, err := os.ReadFile("../../sample-pipeline.yaml")
	if err != nil {
		t.Fatal(err)
	}

	_, err = ValidateYAMLStructure(yamlData)
	if err != nil {
		t.Fatal(err)
	}
}
