// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// DetectedIssueEvidence represents evidence for a detected issue
type DetectedIssueEvidence struct {
	common.BackboneElement

	Code   []common.CodeableConcept `json:"code,omitempty"`
	Detail []common.Reference       `json:"detail,omitempty"`
}

// DetectedIssueMitigation represents mitigation for a detected issue
type DetectedIssueMitigation struct {
	common.BackboneElement

	Action common.CodeableConcept `json:"action"`
	Author *common.Reference      `json:"author,omitempty"`
	Date   *string                `json:"date,omitempty"`
	Note   []Annotation           `json:"note,omitempty"`
}

// DetectedIssue represents a detected clinical issue
type DetectedIssue struct {
	DomainResource

	ResourceType string `json:"resourceType"` // Always "DetectedIssue"

	Author             *common.Reference         `json:"author,omitempty"`
	Category           []common.CodeableConcept  `json:"category,omitempty"`
	Code               *common.CodeableConcept   `json:"code,omitempty"`
	Detail             *string                   `json:"detail,omitempty"`
	Encounter          *common.Reference         `json:"encounter,omitempty"`
	Evidence           []DetectedIssueEvidence   `json:"evidence,omitempty"`
	IdentifiedDateTime *string                   `json:"identifiedDateTime,omitempty"`
	IdentifiedPeriod   *common.Period            `json:"identifiedPeriod,omitempty"`
	Identifier         []common.Identifier       `json:"identifier,omitempty"`
	Implicated         []common.Reference        `json:"implicated,omitempty"`
	Mitigation         []DetectedIssueMitigation `json:"mitigation,omitempty"`
	Reference          *string                   `json:"reference,omitempty"`
	Severity           *DetectedIssueSeverity    `json:"severity,omitempty"`
	Status             DetectedIssueStatus       `json:"status"`
	Subject            *common.Reference         `json:"subject,omitempty"`
}

// DetectedIssueSeverity represents the severity of a detected issue
type DetectedIssueSeverity string

const (
	DetectedIssueSeverityHigh     DetectedIssueSeverity = "high"
	DetectedIssueSeverityModerate DetectedIssueSeverity = "moderate"
	DetectedIssueSeverityLow      DetectedIssueSeverity = "low"
)

// DetectedIssueStatus represents the status of a detected issue
type DetectedIssueStatus string

const (
	DetectedIssueStatusPreliminary    DetectedIssueStatus = "preliminary"
	DetectedIssueStatusFinal          DetectedIssueStatus = "final"
	DetectedIssueStatusEnteredInError DetectedIssueStatus = "entered-in-error"
	DetectedIssueStatusMitigated      DetectedIssueStatus = "mitigated"
)
