package fhir5

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/d4l-data4life/go-fhir/pkg/test_utils"
)

func TestAccount_Serialization(t *testing.T) {
	exampleFiles := []string{
		"../../examples-fhir/fhir5/account-example.json",
		"../../examples-fhir/fhir5/account-example-with-guarantor.json",
	}
	for _, file := range exampleFiles {
		data, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("failed to read example file %s: %v", file, err)
		}
		var account Account
		if err := json.Unmarshal(data, &account); err != nil {
			t.Errorf("failed to unmarshal %s: %v", file, err)
		}
		_, err = json.MarshalIndent(account, "", "  ")
		if err != nil {
			t.Errorf("failed to marshal %s: %v", file, err)
		}
		// Optionally, compare original and serialized JSON for round-trip integrity
	}
}

func TestAccount_Deserialization(t *testing.T) {
	jsonDir := "../../examples-fhir/fhir5-json"
	files := []string{
		"account-example.json",
		"account-example-with-guarantor.json",
	}
	for _, fname := range files {
		path := filepath.Join(jsonDir, fname)
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read json file %s: %v", path, err)
		}
		var account Account
		if err := json.Unmarshal(data, &account); err != nil {
			t.Errorf("failed to unmarshal %s: %v", path, err)
		}
	}
}

func TestAccount_RoundTripSerialization(t *testing.T) {
	exampleFiles := []string{
		"../../examples-fhir/fhir5/account-example.json",
		"../../examples-fhir/fhir5/account-example-with-guarantor.json",
	}
	for _, file := range exampleFiles {
		originalData, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("failed to read example file %s: %v", file, err)
		}

		var account Account
		test_utils.CompareRoundTripSerialization(t, originalData, &account, file)
	}
}
