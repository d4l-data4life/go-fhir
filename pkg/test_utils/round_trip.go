package test_utils

import (
	"encoding/json"
	"testing"
)

// CompareRoundTripSerialization compares original JSON with the result of
// deserializing and then serializing back to JSON. This helps verify that
// resource structs preserve all data correctly during round-trip operations.
func CompareRoundTripSerialization(t *testing.T, originalData []byte, resource interface{}, resourceName string) {
	// Deserialize the original JSON
	if err := json.Unmarshal(originalData, resource); err != nil {
		t.Errorf("failed to unmarshal %s: %v", resourceName, err)
		return
	}

	// Serialize back to JSON
	serializedData, err := json.MarshalIndent(resource, "", "  ")
	if err != nil {
		t.Errorf("failed to marshal %s: %v", resourceName, err)
		return
	}

	// Compare the original and serialized JSON
	// Note: We need to normalize both JSONs for comparison since field order might differ
	var originalMap, serializedMap map[string]interface{}

	if err := json.Unmarshal(originalData, &originalMap); err != nil {
		t.Errorf("failed to unmarshal original JSON to map %s: %v", resourceName, err)
		return
	}

	if err := json.Unmarshal(serializedData, &serializedMap); err != nil {
		t.Errorf("failed to unmarshal serialized JSON to map %s: %v", resourceName, err)
		return
	}

	// Convert back to JSON for comparison (this normalizes field order)
	originalNormalized, err := json.Marshal(originalMap)
	if err != nil {
		t.Errorf("failed to marshal original map %s: %v", resourceName, err)
		return
	}

	serializedNormalized, err := json.Marshal(serializedMap)
	if err != nil {
		t.Errorf("failed to marshal serialized map %s: %v", resourceName, err)
		return
	}

	if string(originalNormalized) != string(serializedNormalized) {
		t.Errorf("round-trip serialization failed for %s: original and serialized JSON do not match", resourceName)
		t.Logf("Original: %s", string(originalNormalized))
		t.Logf("Serialized: %s", string(serializedNormalized))
	}
}
