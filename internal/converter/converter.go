package converter

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

// JSONToYAML converts a JSON byte slice to a YAML byte slice.
func JSONToYAML(data []byte) ([]byte, error) {
    var intermediate map[string]interface{}
    if err := json.Unmarshal(data, &intermediate); err != nil {
        return nil, fmt.Errorf("failed to parse JSON: %w", err)
    }

    yamlData, err := yaml.Marshal(intermediate)
    if err != nil {
        return nil, fmt.Errorf("failed to encode YAML: %w", err)
    }
    return yamlData, nil
}

// YAMLToJSON converts a YAML byte slice to a pretty-printed JSON byte slice.
func YAMLToJSON(data []byte) ([]byte, error) {
    var intermediate map[string]interface{}
    if err := yaml.Unmarshal(data, &intermediate); err != nil {
        return nil, fmt.Errorf("failed to parse YAML: %w", err)
    }

    jsonData, err := json.MarshalIndent(intermediate, "", "  ")
    if err != nil {
        return nil, fmt.Errorf("failed to encode JSON: %w", err)
    }
    return jsonData, nil
}
