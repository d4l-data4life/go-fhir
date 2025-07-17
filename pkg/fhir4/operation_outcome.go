package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// OperationOutcome represents a collection of error, warning, or information messages that result from a system action
type OperationOutcome struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "OperationOutcome"

	// An error, warning, or information message that results from a system action
	Issue []OperationOutcomeIssue `json:"issue"`
}

// OperationOutcomeIssue represents an error, warning, or information message that results from a system action
type OperationOutcomeIssue struct {
	common.BackboneElement

	// Describes the type of the issue
	Code string `json:"code"`

	// A human readable description of the error issue SHOULD be placed in details.text
	Details *common.CodeableConcept `json:"details,omitempty"`

	// This may be a description of how a value is erroneous, a stack dump to help trace the issue or other troubleshooting information
	Diagnostics *string `json:"diagnostics,omitempty"`

	// The root of the FHIRPath is the resource or bundle that generated OperationOutcome
	Expression []string `json:"expression,omitempty"`

	// The root of the XPath is the resource or bundle that generated OperationOutcome
	Location []string `json:"location,omitempty"`

	// This is labeled as "Is Modifier" because applications should not confuse hints and warnings with errors
	Severity OperationOutcomeIssueSeverity `json:"severity"`
}

// OperationOutcomeIssueSeverity represents this is labeled as "Is Modifier" because applications should not confuse hints and warnings with errors
type OperationOutcomeIssueSeverity string

const (
	OperationOutcomeIssueSeverityFatal       OperationOutcomeIssueSeverity = "fatal"
	OperationOutcomeIssueSeverityError       OperationOutcomeIssueSeverity = "error"
	OperationOutcomeIssueSeverityWarning     OperationOutcomeIssueSeverity = "warning"
	OperationOutcomeIssueSeverityInformation OperationOutcomeIssueSeverity = "information"
)
