// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

// EvidenceVariable represents the EvidenceVariable resource
type EvidenceVariable struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "EvidenceVariable"

	// (Add all fields from the EvidenceVariable struct here)
}

// (Add related types: EvidenceVariableCharacteristic, etc. as needed)
