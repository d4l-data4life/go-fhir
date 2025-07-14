// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// DiagnosticReportStatus represents the status of a diagnostic report
type DiagnosticReportStatus string

const (
	DiagnosticReportStatusRegistered     DiagnosticReportStatus = "registered"
	DiagnosticReportStatusPartial        DiagnosticReportStatus = "partial"
	DiagnosticReportStatusPreliminary    DiagnosticReportStatus = "preliminary"
	DiagnosticReportStatusModified       DiagnosticReportStatus = "modified"
	DiagnosticReportStatusFinal          DiagnosticReportStatus = "final"
	DiagnosticReportStatusAmended        DiagnosticReportStatus = "amended"
	DiagnosticReportStatusCorrected      DiagnosticReportStatus = "corrected"
	DiagnosticReportStatusAppended       DiagnosticReportStatus = "appended"
	DiagnosticReportStatusCancelled      DiagnosticReportStatus = "cancelled"
	DiagnosticReportStatusEnteredInError DiagnosticReportStatus = "entered-in-error"
	DiagnosticReportStatusUnknown        DiagnosticReportStatus = "unknown"
)

// DiagnosticReportSupportingInfo represents supporting information used in the creation of the report
type DiagnosticReportSupportingInfo struct {
	common.BackboneElement

	// The reference for the supporting information in the diagnostic report
	Reference common.Reference `json:"reference"`

	// The code value for the role of the supporting information in the diagnostic report
	Type common.CodeableConcept `json:"type"`
}

// DiagnosticReportMedia represents key images or data associated with the report
type DiagnosticReportMedia struct {
	common.BackboneElement

	// The comment should be displayed with the image or data
	Comment *string `json:"comment,omitempty"`

	// Reference to the image or data source
	Link common.Reference `json:"link"`
}

// DiagnosticReport represents diagnostic test results and interpretation
type DiagnosticReport struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DiagnosticReport"

	// Note: Usually there is one test request for each result, however in some circumstances multiple test requests may be represented using a single test result resource
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Multiple categories are allowed using various categorization schemes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// A code or name that describes this diagnostic report
	Code common.CodeableConcept `json:"code"`

	// The Composition provides structure to the content of the DiagnosticReport
	Composition *common.Reference `json:"composition,omitempty"`

	// Concise and clinically contextualized summary conclusion (interpretation/impression) of the diagnostic report
	Conclusion *string `json:"conclusion,omitempty"`

	// One or more codes that represent the summary conclusion (interpretation/impression) of the diagnostic report
	ConclusionCode []common.CodeableConcept `json:"conclusionCode,omitempty"`

	// If the diagnostic procedure was performed on the patient, this is the time it was performed
	EffectiveDateTime *string        `json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *common.Period `json:"effectivePeriod,omitempty"`

	// This will typically be the encounter the event occurred within
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Usually assigned by the Information System of the diagnostic service provider (filler id)
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// May be different from the update time of the resource itself, because that is the status of the record
	Issued *string `json:"issued,omitempty"`

	// A list of key images or data associated with this report
	Media []DiagnosticReportMedia `json:"media,omitempty"`

	// May include general statements about the diagnostic report, or statements about significant, unexpected or unreliable results values
	Note []Annotation `json:"note,omitempty"`

	// This is not necessarily the source of the atomic data items or the entity that interpreted the results
	Performer []common.Reference `json:"performer,omitempty"`

	// "application/pdf" is recommended as the most reliable and interoperable in this context
	PresentedForm []Attachment `json:"presentedForm,omitempty"`

	// Observations can contain observations
	Result []common.Reference `json:"result,omitempty"`

	// Might not be the same entity that takes responsibility for the clinical report
	ResultsInterpreter []common.Reference `json:"resultsInterpreter,omitempty"`

	// If the specimen is sufficiently specified with a code in the test result name, then this additional data may be redundant
	Specimen []common.Reference `json:"specimen,omitempty"`

	// The status of the diagnostic report
	Status DiagnosticReportStatus `json:"status"`

	// For laboratory-type studies like GenomeStudy, type resources will be used for tracking additional metadata
	Study []common.Reference `json:"study,omitempty"`

	// The subject of the report. Usually, but not always, this is a patient
	Subject *common.Reference `json:"subject,omitempty"`

	// This backbone element contains supporting information that was used in the creation of the report
	SupportingInfo []DiagnosticReportSupportingInfo `json:"supportingInfo,omitempty"`
}
