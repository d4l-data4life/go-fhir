// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import "github.com/d4l-data4life/go-fhir/pkg/common"

// ResearchSubjectProgress represents a state or milestone in a research subject's journey
// See: https://hl7.org/fhir/researchsubject.html
// Generated from js/r5.d.ts

type ResearchSubjectProgress struct {
	common.BackboneElement

	// The date when the state ended
	EndDate *string `json:"endDate,omitempty"`

	// The milestone achieved
	Milestone *common.CodeableConcept `json:"milestone,omitempty"`

	// The reason for the state change
	Reason *common.CodeableConcept `json:"reason,omitempty"`

	// The date when the state started
	StartDate *string `json:"startDate,omitempty"`

	// The current state of the subject
	SubjectState *common.CodeableConcept `json:"subjectState,omitempty"`

	// The aspect of the subject's journey that the state refers to
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ResearchSubject represents a participant or object which is the recipient of investigative activities in a research study
// See: https://hl7.org/fhir/researchsubject.html
// Generated from js/r5.d.ts

type ResearchSubject struct {
	DomainResource

	ResourceType string `json:"resourceType"` // Always "ResearchSubject"

	// The name of the arm in the study the subject actually followed
	ActualComparisonGroup *string `json:"actualComparisonGroup,omitempty"`

	// The name of the arm assigned to the subject
	AssignedComparisonGroup *string `json:"assignedComparisonGroup,omitempty"`

	// A record of the patient's informed agreement to participate in the study
	Consent []common.Reference `json:"consent,omitempty"`

	// Identifiers assigned to this research subject for a study
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The dates the subject began and ended their participation in the study
	Period *common.Period `json:"period,omitempty"`

	// The subject's progress through states and milestones
	Progress []ResearchSubjectProgress `json:"progress,omitempty"`

	// The publication state of the resource (not of the subject)
	Status string `json:"status"`

	// Reference to the study the subject is participating in
	Study common.Reference `json:"study"`

	// The record of the person, animal or other entity involved in the study
	Subject common.Reference `json:"subject"`
}
