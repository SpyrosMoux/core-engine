package helpers

import (
	"gopkg.in/yaml.v2"
)

func ReadYAMLFromString(yamlData string, out interface{}) error {
	return yaml.Unmarshal([]byte(yamlData), out)
}
