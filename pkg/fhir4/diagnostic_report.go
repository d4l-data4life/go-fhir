package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// DiagnosticReport represents the findings and interpretation of diagnostic tests
type DiagnosticReport struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DiagnosticReport"

	// Business identifier for report
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Details concerning a service requested
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// registered | partial | preliminary | final | amended | corrected | appended | cancelled | entered-in-error | unknown
	Status DiagnosticReportStatus `json:"status"`

	// Service category
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Name/Code for this diagnostic report
	Code common.CodeableConcept `json:"code"`

	// The subject of the report - usually, but not always, this is a patient
	Subject *common.Reference `json:"subject,omitempty"`

	// Health care event when test ordered
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Clinically relevant time/time-period for report
	EffectiveDateTime *string        `json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *common.Period `json:"effectivePeriod,omitempty"`

	// DateTime this version was made
	Issued *string `json:"issued,omitempty"`

	// Responsible Diagnostic Service
	Performer []common.Reference `json:"performer,omitempty"`

	// Primary result interpreter
	ResultsInterpreter []common.Reference `json:"resultsInterpreter,omitempty"`

	// Specimens this report is based on
	Specimen []common.Reference `json:"specimen,omitempty"`

	// Observations
	Result []common.Reference `json:"result,omitempty"`

	// Key images associated with this report
	Media []DiagnosticReportMedia `json:"media,omitempty"`

	// Clinical conclusion (interpretation) of test results
	Conclusion *string `json:"conclusion,omitempty"`

	// Codes for the clinical conclusion of test results
	ConclusionCode []common.CodeableConcept `json:"conclusionCode,omitempty"`

	// Entire report as attachment
	PresentedForm []Attachment `json:"presentedForm,omitempty"`
}

// DiagnosticReportStatus represents the status of the diagnostic report
type DiagnosticReportStatus string

const (
	DiagnosticReportStatusRegistered     DiagnosticReportStatus = "registered"
	DiagnosticReportStatusPartial        DiagnosticReportStatus = "partial"
	DiagnosticReportStatusPreliminary    DiagnosticReportStatus = "preliminary"
	DiagnosticReportStatusFinal          DiagnosticReportStatus = "final"
	DiagnosticReportStatusAmended        DiagnosticReportStatus = "amended"
	DiagnosticReportStatusCorrected      DiagnosticReportStatus = "corrected"
	DiagnosticReportStatusAppended       DiagnosticReportStatus = "appended"
	DiagnosticReportStatusCancelled      DiagnosticReportStatus = "cancelled"
	DiagnosticReportStatusEnteredInError DiagnosticReportStatus = "entered-in-error"
	DiagnosticReportStatusUnknown        DiagnosticReportStatus = "unknown"
)

// DiagnosticReportMedia represents key images associated with the report
type DiagnosticReportMedia struct {
	common.BackboneElement

	// Comment about the image (e.g. explanation)
	Comment *string `json:"comment,omitempty"`

	// Reference to the image source
	Link common.Reference `json:"link"`
}
