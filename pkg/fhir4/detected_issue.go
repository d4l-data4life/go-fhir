package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// DetectedIssue represents the DetectedIssue resource
type DetectedIssue struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DetectedIssue"

	// Individual or device responsible for the issue being raised
	Author *common.Reference `json:"author,omitempty"`

	// Identifies the general type of issue identified
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Should focus on information not covered elsewhere as discrete data
	Detail *string `json:"detail,omitempty"`

	// Supporting evidence or manifestations that provide the basis for identifying the detected issue
	Evidence []DetectedIssueEvidence `json:"evidence,omitempty"`

	// The date or period when the detected issue was initially identified
	IdentifiedDateTime *string `json:"identifiedDateTime,omitempty"`

	// The date or period when the detected issue was initially identified
	IdentifiedPeriod *common.Period `json:"identifiedPeriod,omitempty"`

	// Business identifier associated with the detected issue record
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// There's an implicit constraint on the number of implicated resources based on DetectedIssue.type
	Implicated []common.Reference `json:"implicated,omitempty"`

	// Indicates an action that has been taken or is committed to reduce or eliminate the likelihood of the risk
	Mitigation []DetectedIssueMitigation `json:"mitigation,omitempty"`

	// Indicates the patient whose record the detected issue is associated with
	Patient *common.Reference `json:"patient,omitempty"`

	// The literature, knowledge-base or similar reference that describes the propensity for the detected issue identified
	Reference *string `json:"reference,omitempty"`

	// Indicates the degree of importance associated with the identified issue based on the potential impact on the patient
	Severity *DetectedIssueSeverity `json:"severity,omitempty"`

	// This element is labeled as a modifier because the status contains the codes cancelled and entered-in-error that mark the issue as not currently valid
	Status DetectedIssueStatus `json:"status"`
}

// DetectedIssueSeverity represents the severity of the detected issue
type DetectedIssueSeverity string

const (
	DetectedIssueSeverityHigh     DetectedIssueSeverity = "high"
	DetectedIssueSeverityModerate DetectedIssueSeverity = "moderate"
	DetectedIssueSeverityLow      DetectedIssueSeverity = "low"
)

// DetectedIssueStatus represents the status of the detected issue
type DetectedIssueStatus string

const (
	DetectedIssueStatusRegistered     DetectedIssueStatus = "registered"
	DetectedIssueStatusPreliminary    DetectedIssueStatus = "preliminary"
	DetectedIssueStatusFinal          DetectedIssueStatus = "final"
	DetectedIssueStatusAmended        DetectedIssueStatus = "amended"
	DetectedIssueStatusCorrected      DetectedIssueStatus = "corrected"
	DetectedIssueStatusCancelled      DetectedIssueStatus = "cancelled"
	DetectedIssueStatusEnteredInError DetectedIssueStatus = "entered-in-error"
	DetectedIssueStatusUnknown        DetectedIssueStatus = "unknown"
)

// DetectedIssueEvidence represents supporting evidence or manifestations
type DetectedIssueEvidence struct {
	common.BackboneElement

	// A manifestation that led to the recording of this detected issue
	Code []common.CodeableConcept `json:"code,omitempty"`

	// Links to resources that constitute evidence for the detected issue such as a GuidanceResponse or MeasureReport
	Detail []common.Reference `json:"detail,omitempty"`
}

// DetectedIssueMitigation represents an action that has been taken or is committed to reduce or eliminate the likelihood of the risk
type DetectedIssueMitigation struct {
	common.BackboneElement

	// The "text" component can be used for detail or when no appropriate code exists
	Action common.CodeableConcept `json:"action"`

	// Identifies the practitioner who determined the mitigation and takes responsibility for the mitigation step occurring
	Author *common.Reference `json:"author,omitempty"`

	// This might not be the same as when the mitigating step was actually taken
	Date *string `json:"date,omitempty"`
}
