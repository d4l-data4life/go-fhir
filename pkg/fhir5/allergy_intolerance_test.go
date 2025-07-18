package fhir5

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/d4l-data4life/go-fhir/pkg/test_utils"
)

func TestAllergyIntolerance_Serialization(t *testing.T) {
	exampleFiles := []string{
		"../../examples-fhir/fhir5/allergyintolerance-example.json",
		"../../examples-fhir/fhir5/allergyintolerance-fishallergy.json",
		"../../examples-fhir/fhir5/allergyintolerance-medication.json",
		"../../examples-fhir/fhir5/allergyintolerance-nka.json",
		"../../examples-fhir/fhir5/allergyintolerance-nkda.json",
		"../../examples-fhir/fhir5/allergyintolerance-nkla.json",
	}

	for _, file := range exampleFiles {
		t.Run(filepath.Base(file), func(t *testing.T) {
			// Read the example file
			data, err := os.ReadFile(file)
			if err != nil {
				t.Fatalf("failed to read file %s: %v", file, err)
			}

			// Deserialize the JSON
			var allergyIntolerance AllergyIntolerance
			if err := json.Unmarshal(data, &allergyIntolerance); err != nil {
				t.Fatalf("failed to unmarshal allergyIntolerance: %v", err)
			}

			// Verify the resource type
			if allergyIntolerance.ResourceType != "AllergyIntolerance" {
				t.Errorf("expected resourceType to be 'AllergyIntolerance', got '%s'", allergyIntolerance.ResourceType)
			}

			// Serialize back to JSON
			serialized, err := json.MarshalIndent(allergyIntolerance, "", "  ")
			if err != nil {
				t.Fatalf("failed to marshal allergyIntolerance: %v", err)
			}

			// Verify that the serialized JSON is valid
			var verifyAllergyIntolerance AllergyIntolerance
			if err := json.Unmarshal(serialized, &verifyAllergyIntolerance); err != nil {
				t.Fatalf("failed to unmarshal serialized allergyIntolerance: %v", err)
			}

			t.Logf("Successfully serialized and deserialized allergyIntolerance from %s", file)
		})
	}
}

func TestAllergyIntolerance_Deserialization(t *testing.T) {
	exampleFiles := []string{
		"../../examples-fhir/fhir5/allergyintolerance-example.json",
		"../../examples-fhir/fhir5/allergyintolerance-fishallergy.json",
		"../../examples-fhir/fhir5/allergyintolerance-medication.json",
		"../../examples-fhir/fhir5/allergyintolerance-nka.json",
		"../../examples-fhir/fhir5/allergyintolerance-nkda.json",
		"../../examples-fhir/fhir5/allergyintolerance-nkla.json",
	}

	for _, file := range exampleFiles {
		t.Run(filepath.Base(file), func(t *testing.T) {
			// Read the example file
			data, err := os.ReadFile(file)
			if err != nil {
				t.Fatalf("failed to read file %s: %v", file, err)
			}

			// Deserialize the JSON
			var allergyIntolerance AllergyIntolerance
			if err := json.Unmarshal(data, &allergyIntolerance); err != nil {
				t.Fatalf("failed to unmarshal allergyIntolerance: %v", err)
			}

			// Verify basic fields are populated
			if allergyIntolerance.ResourceType != "AllergyIntolerance" {
				t.Errorf("expected resourceType to be 'AllergyIntolerance', got '%s'", allergyIntolerance.ResourceType)
			}

			// Verify that patient is present (it is required)
			if allergyIntolerance.Patient.Reference == nil || *allergyIntolerance.Patient.Reference == "" {
				t.Error("expected patient reference to be present")
			}

			t.Logf("Successfully deserialized allergyIntolerance from %s", file)
		})
	}
}

func TestAllergyIntolerance_RoundTripSerialization(t *testing.T) {
	exampleFiles := []string{
		"../../examples-fhir/fhir5/allergyintolerance-example.json",
		"../../examples-fhir/fhir5/allergyintolerance-fishallergy.json",
		"../../examples-fhir/fhir5/allergyintolerance-medication.json",
		"../../examples-fhir/fhir5/allergyintolerance-nka.json",
		"../../examples-fhir/fhir5/allergyintolerance-nkda.json",
		"../../examples-fhir/fhir5/allergyintolerance-nkla.json",
	}

	for _, file := range exampleFiles {
		t.Run(filepath.Base(file), func(t *testing.T) {
			// Read the example file
			data, err := os.ReadFile(file)
			if err != nil {
				t.Fatalf("failed to read file %s: %v", file, err)
			}

			// Perform round-trip serialization test
			var allergyIntolerance AllergyIntolerance
			test_utils.CompareRoundTripSerialization(t, data, &allergyIntolerance, "AllergyIntolerance")
		})
	}
}
