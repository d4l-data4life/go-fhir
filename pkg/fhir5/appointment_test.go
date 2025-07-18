package fhir5

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/d4l-data4life/go-fhir/pkg/test_utils"
)

func TestAppointment_Serialization(t *testing.T) {
	exampleFiles := []string{
		"testdata/fhir5-json/appointment-example.json",
		"testdata/fhir5-json/appointment-example-request.json",
		"testdata/fhir5-json/appointment-example2doctors.json",
	}

	for _, file := range exampleFiles {
		t.Run(filepath.Base(file), func(t *testing.T) {
			// Read the example file
			data, err := os.ReadFile(file)
			if err != nil {
				t.Fatalf("failed to read file %s: %v", file, err)
			}

			// Deserialize the JSON
			var appointment Appointment
			if err := json.Unmarshal(data, &appointment); err != nil {
				t.Fatalf("failed to unmarshal appointment: %v", err)
			}

			// Verify the resource type
			if appointment.ResourceType != "Appointment" {
				t.Errorf("expected resourceType to be 'Appointment', got '%s'", appointment.ResourceType)
			}

			// Serialize back to JSON
			serialized, err := json.MarshalIndent(appointment, "", "  ")
			if err != nil {
				t.Fatalf("failed to marshal appointment: %v", err)
			}

			// Verify that the serialized JSON is valid
			var verifyAppointment Appointment
			if err := json.Unmarshal(serialized, &verifyAppointment); err != nil {
				t.Fatalf("failed to unmarshal serialized appointment: %v", err)
			}

			t.Logf("Successfully serialized and deserialized appointment from %s", file)
		})
	}
}

func TestAppointment_Deserialization(t *testing.T) {
	exampleFiles := []string{
		"testdata/fhir5-json/appointment-example.json",
		"testdata/fhir5-json/appointment-example-request.json",
		"testdata/fhir5-json/appointment-example2doctors.json",
	}

	for _, file := range exampleFiles {
		t.Run(filepath.Base(file), func(t *testing.T) {
			// Read the example file
			data, err := os.ReadFile(file)
			if err != nil {
				t.Fatalf("failed to read file %s: %v", file, err)
			}

			// Deserialize the JSON
			var appointment Appointment
			if err := json.Unmarshal(data, &appointment); err != nil {
				t.Fatalf("failed to unmarshal appointment: %v", err)
			}

			// Verify basic fields are populated
			if appointment.ResourceType != "Appointment" {
				t.Errorf("expected resourceType to be 'Appointment', got '%s'", appointment.ResourceType)
			}

			// Verify that participants are present (they are required)
			if len(appointment.Participant) == 0 {
				t.Error("expected at least one participant")
			}

			t.Logf("Successfully deserialized appointment from %s", file)
		})
	}
}

func TestAppointment_RoundTripSerialization(t *testing.T) {
	exampleFiles := []string{
		"testdata/fhir5-json/appointment-example.json",
		"testdata/fhir5-json/appointment-example-request.json",
		"testdata/fhir5-json/appointment-example2doctors.json",
	}

	for _, file := range exampleFiles {
		t.Run(filepath.Base(file), func(t *testing.T) {
			// Read the example file
			data, err := os.ReadFile(file)
			if err != nil {
				t.Fatalf("failed to read file %s: %v", file, err)
			}

			// Perform round-trip serialization test
			var appointment Appointment
			test_utils.CompareRoundTripSerialization(t, data, &appointment, "Appointment")
		})
	}
}
