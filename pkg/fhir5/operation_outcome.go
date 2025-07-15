// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// OperationOutcomeIssueSeverity represents severity of operation outcome issue
type OperationOutcomeIssueSeverity string

const (
	OperationOutcomeIssueSeverityFatal       OperationOutcomeIssueSeverity = "fatal"
	OperationOutcomeIssueSeverityError       OperationOutcomeIssueSeverity = "error"
	OperationOutcomeIssueSeverityWarning     OperationOutcomeIssueSeverity = "warning"
	OperationOutcomeIssueSeverityInformation OperationOutcomeIssueSeverity = "information"
	OperationOutcomeIssueSeveritySuccess     OperationOutcomeIssueSeverity = "success"
)

// OperationOutcomeIssue represents an error, warning, or information message
type OperationOutcomeIssue struct {
	common.BackboneElement

	// Code values should align with the severity
	Code string `json:"code"`

	// A human readable description of the error issue
	Details *common.CodeableConcept `json:"details,omitempty"`

	// This may be a description of how a value is erroneous
	Diagnostics *string `json:"diagnostics,omitempty"`

	// The root of the FHIRPath is the resource or bundle that generated OperationOutcome
	Expression []string `json:"expression,omitempty"`

	// The root of the XPath is the resource or bundle that generated OperationOutcome
	Location []string `json:"location,omitempty"`

	// Indicates whether the issue indicates a variation from successful processing
	Severity OperationOutcomeIssueSeverity `json:"severity"`
}

// OperationOutcome represents a collection of error, warning, or information messages
type OperationOutcome struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "OperationOutcome"

	// An error, warning, or information message that results from a system action
	Issue []OperationOutcomeIssue `json:"issue"`
}
