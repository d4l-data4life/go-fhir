package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/d4l-data4life/go-fhir/pkg/common"
	"github.com/d4l-data4life/go-fhir/pkg/fhir4"
)

func main() {
	// Create a simple Patient example
	patient := createPatientExample()

	// Convert to JSON
	jsonData, err := json.MarshalIndent(patient, "", "  ")
	if err != nil {
		log.Fatal("Error marshaling patient:", err)
	}

	fmt.Println("FHIR R4 Patient Resource:")
	fmt.Println(string(jsonData))

	// Demonstrate unmarshaling
	var unmarshaledPatient fhir4.Patient
	err = json.Unmarshal(jsonData, &unmarshaledPatient)
	if err != nil {
		log.Fatal("Error unmarshaling patient:", err)
	}

	fmt.Printf("\nSuccessfully roundtrip marshaled/unmarshaled patient: %s\n", *unmarshaledPatient.ID)
}

func createPatientExample() *fhir4.Patient {
	return &fhir4.Patient{
		DomainResource: fhir4.DomainResource{
			Resource: fhir4.Resource{
				ResourceType: "Patient",
				ID:           stringPtr("example-patient-001"),
			},
		},
		Active: boolPtr(true),
		Identifier: []common.Identifier{
			{
				Use: identifierUsePtr(common.IdentifierUseUsual),
				Type: &common.CodeableConcept{
					Coding: []common.Coding{
						{
							System:  stringPtr("http://terminology.hl7.org/CodeSystem/v2-0203"),
							Code:    stringPtr("MR"),
							Display: stringPtr("Medical Record Number"),
						},
					},
				},
				System: stringPtr("http://hospital.example.org"),
				Value:  stringPtr("123456789"),
			},
		},
		Name: []fhir4.HumanName{
			{
				Use:    nameUsePtr(fhir4.NameUseOfficial),
				Family: stringPtr("Smith"),
				Given:  []string{"John", "David"},
				Prefix: []string{"Mr."},
			},
		},
		Telecom: []fhir4.ContactPoint{
			{
				System: contactPointSystemPtr(fhir4.ContactPointSystemPhone),
				Value:  stringPtr("+1-555-123-4567"),
				Use:    contactPointUsePtr(fhir4.ContactPointUseHome),
			},
			{
				System: contactPointSystemPtr(fhir4.ContactPointSystemEmail),
				Value:  stringPtr("john.smith@example.com"),
				Use:    contactPointUsePtr(fhir4.ContactPointUseHome),
			},
		},
		Gender:    administrativeGenderPtr(fhir4.AdministrativeGenderMale),
		BirthDate: stringPtr("1985-03-15"),
		Address: []fhir4.Address{
			{
				Use:        addressUsePtr(fhir4.AddressUseHome),
				Type:       addressTypePtr(fhir4.AddressTypePhysical),
				Line:       []string{"456 Oak Avenue", "Apt 2B"},
				City:       stringPtr("Springfield"),
				State:      stringPtr("IL"),
				PostalCode: stringPtr("62701"),
				Country:    stringPtr("US"),
				Period: &common.Period{
					Start: parseFHIRDateTimePtr("2020-01-01T00:00:00Z"),
				},
			},
		},
		MaritalStatus: &common.CodeableConcept{
			Coding: []common.Coding{
				{
					System:  stringPtr("http://terminology.hl7.org/CodeSystem/v3-MaritalStatus"),
					Code:    stringPtr("M"),
					Display: stringPtr("Married"),
				},
			},
		},
		Contact: []fhir4.PatientContact{
			{
				Relationship: []common.CodeableConcept{
					{
						Coding: []common.Coding{
							{
								System:  stringPtr("http://terminology.hl7.org/CodeSystem/v2-0131"),
								Code:    stringPtr("E"),
								Display: stringPtr("Emergency Contact"),
							},
						},
					},
				},
				Name: &fhir4.HumanName{
					Family: stringPtr("Smith"),
					Given:  []string{"Jane"},
				},
				Telecom: []fhir4.ContactPoint{
					{
						System: contactPointSystemPtr(fhir4.ContactPointSystemPhone),
						Value:  stringPtr("+1-555-987-6543"),
						Use:    contactPointUsePtr(fhir4.ContactPointUseMobile),
					},
				},
				Gender: administrativeGenderPtr(fhir4.AdministrativeGenderFemale),
			},
		},
		Communication: []fhir4.PatientCommunication{
			{
				Language: common.CodeableConcept{
					Coding: []common.Coding{
						{
							System:  stringPtr("urn:ietf:bcp:47"),
							Code:    stringPtr("en-US"),
							Display: stringPtr("English (United States)"),
						},
					},
				},
				Preferred: boolPtr(true),
			},
		},
	}
}

// Helper functions for creating pointers
func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func parseFHIRDateTimePtr(s string) *common.FHIRDateTime {
	fhirDateTime, err := common.ParseDateTime(s)
	if err != nil {
		return nil
	}
	return fhirDateTime
}

func identifierUsePtr(use common.IdentifierUse) *common.IdentifierUse {
	return &use
}

func nameUsePtr(use fhir4.NameUse) *fhir4.NameUse {
	return &use
}

func contactPointSystemPtr(system fhir4.ContactPointSystem) *fhir4.ContactPointSystem {
	return &system
}

func contactPointUsePtr(use fhir4.ContactPointUse) *fhir4.ContactPointUse {
	return &use
}

func administrativeGenderPtr(gender fhir4.AdministrativeGender) *fhir4.AdministrativeGender {
	return &gender
}

func addressUsePtr(use fhir4.AddressUse) *fhir4.AddressUse {
	return &use
}

func addressTypePtr(addrType fhir4.AddressType) *fhir4.AddressType {
	return &addrType
}
