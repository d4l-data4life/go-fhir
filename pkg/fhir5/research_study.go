// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ResearchStudy represents a process where a researcher or organization plans and then executes a series of steps intended to increase the field of healthcare-related knowledge
type ResearchStudy struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ResearchStudy"

	// The date on which the research study was approved by the research ethics board
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// Classification for the study
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Contact details to assist a user in learning more about or engaging with the study
	Contact []ContactDetail `json:"contact,omitempty"`

	// A full description of how the study is being conducted
	Description *string `json:"description,omitempty"`

	// The date on which the research study was terminated
	EndDate *string `json:"endDate,omitempty"`

	// The facility where the study is being conducted
	Facility []common.Reference `json:"facility,omitempty"`

	// The condition that is the focus of the research study
	Focus []common.CodeableConcept `json:"focus,omitempty"`

	// Identifiers assigned to this research study by the sponsor or other systems
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A key word to identify the study
	Keyword []common.CodeableConcept `json:"keyword,omitempty"`

	// The location where the study is being conducted
	Location []common.CodeableConcept `json:"location,omitempty"`

	// A description and/or code explaining the premature termination of the study
	Note []Annotation `json:"note,omitempty"`

	// The party responsible for the execution of the study
	Protocol []common.Reference `json:"protocol,omitempty"`

	// A description and/or code explaining the premature termination of the study
	ReasonStopped *common.CodeableConcept `json:"reasonStopped,omitempty"`

	// Related artifacts such as publications, references, and other resources
	RelatedArtifact []interface{} `json:"relatedArtifact,omitempty"`

	// The party responsible for the execution of the study
	Sponsor *common.Reference `json:"sponsor,omitempty"`

	// The date on which the research study began
	StartDate *string `json:"startDate,omitempty"`

	// The current state of the study
	Status string `json:"status"`

	// The title of the study
	Title *string `json:"title,omitempty"`

	// The URL where the study can be found
	URL *string `json:"url,omitempty"`

	// The version of the study
	Version *string `json:"version,omitempty"`

	// The subjects that are part of the study
	Enrollment []common.Reference `json:"enrollment,omitempty"`

	// The period during which the study is conducted
	Period *common.Period `json:"period,omitempty"`

	// The phase of the study
	Phase *common.CodeableConcept `json:"phase,omitempty"`

	// The primary purpose of the study
	PrimaryPurposeType *common.CodeableConcept `json:"primaryPurposeType,omitempty"`

	// The sites where the study is being conducted
	Site []common.Reference `json:"site,omitempty"`

	// The design of the study
	StudyDesign []interface{} `json:"studyDesign,omitempty"`

	// The target population for the study
	TargetPopulation []common.Reference `json:"targetPopulation,omitempty"`
}
