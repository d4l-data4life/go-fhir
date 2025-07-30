// Package fhir4 contains FHIR R4 (version 4.0.1) resource definitions
package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Composition represents a set of healthcare-related information that is assembled together into a single logical package
type Composition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Composition"

	// Logical identifier for the composition
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// The workflow/clinical status of the composition
	Status CompositionStatus `json:"status"`

	// Kind of composition (LOINC if possible)
	Type common.CodeableConcept `json:"type"`

	// Categorization of Composition
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Who and/or what the composition is about
	Subject *common.Reference `json:"subject,omitempty"`

	// Context of the Composition
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Composition editing time
	Date string `json:"date"`

	// Who and/or what authored the composition
	Author []common.Reference `json:"author"`

	Title string `json:"title"`

	// Who attested the composition
	Attester []CompositionAttester `json:"attester,omitempty"`

	// Organization which maintains the composition
	Custodian *common.Reference `json:"custodian,omitempty"`

	// Relationships to other compositions/documents
	RelatesTo []CompositionRelatesTo `json:"relatesTo,omitempty"`

	// The clinical service(s) being documented
	Event []CompositionEvent `json:"event,omitempty"`

	// The root of the sections that make up the composition
	Section []CompositionSection `json:"section,omitempty"`
}

// CompositionStatus represents the workflow/clinical status of the composition
type CompositionStatus string

const (
	CompositionStatusPreliminary    CompositionStatus = "preliminary"
	CompositionStatusFinal          CompositionStatus = "final"
	CompositionStatusAmended        CompositionStatus = "amended"
	CompositionStatusEnteredInError CompositionStatus = "entered-in-error"
)

// CompositionAttester represents an attester of the composition
type CompositionAttester struct {
	common.BackboneElement

	// personal | professional | legal | official
	Mode []CompositionAttesterMode `json:"mode"`

	// When the composition was attested
	Time *string `json:"time,omitempty"`

	// Who attested the composition
	Party *common.Reference `json:"party,omitempty"`
}

// CompositionAttesterMode represents the mode of attestation
type CompositionAttesterMode string

const (
	CompositionAttesterModePersonal     CompositionAttesterMode = "personal"
	CompositionAttesterModeProfessional CompositionAttesterMode = "professional"
	CompositionAttesterModeLegal        CompositionAttesterMode = "legal"
	CompositionAttesterModeOfficial     CompositionAttesterMode = "official"
)

// CompositionRelatesTo represents relationships to other compositions/documents
type CompositionRelatesTo struct {
	common.BackboneElement

	// replaces | transforms | signs | appends
	Code CompositionRelatesToCode `json:"code"`

	// Target of the relationship
	TargetIdentifier *common.Identifier `json:"targetIdentifier,omitempty"`
	TargetReference  *common.Reference  `json:"targetReference,omitempty"`
}

// CompositionRelatesToCode represents the code for relationships to other compositions/documents
type CompositionRelatesToCode string

const (
	CompositionRelatesToCodeReplaces   CompositionRelatesToCode = "replaces"
	CompositionRelatesToCodeTransforms CompositionRelatesToCode = "transforms"
	CompositionRelatesToCodeSigns      CompositionRelatesToCode = "signs"
	CompositionRelatesToCodeAppends    CompositionRelatesToCode = "appends"
)

// CompositionEvent represents a clinical service being documented
type CompositionEvent struct {
	common.BackboneElement

	// Code(s) that apply to the event being documented
	Code []common.CodeableConcept `json:"code,omitempty"`

	// The period covered by the documentation
	Period *common.Period `json:"period,omitempty"`

	// The event(s) being documented
	Detail []common.Reference `json:"detail,omitempty"`
}

// CompositionSection represents a section in the composition
type CompositionSection struct {
	common.BackboneElement

	// Label for section (e.g. for ToC)
	Title *string `json:"title,omitempty"`

	// Classification of section (recommended)
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Who and/or what authored the section
	Author []common.Reference `json:"author,omitempty"`

	// Who/what the section is about, when it is not about the subject of composition
	Focus *common.Reference `json:"focus,omitempty"`

	// Text summary of the section, for human interpretation
	Text *Narrative `json:"text,omitempty"`

	// working | snapshot | changes
	Mode *SectionMode `json:"mode,omitempty"`

	// Order of section entries
	OrderedBy *common.CodeableConcept `json:"orderedBy,omitempty"`

	// A reference to data that supports this section
	Entry []common.Reference `json:"entry,omitempty"`

	// Why the section is empty
	EmptyReason *common.CodeableConcept `json:"emptyReason,omitempty"`

	// Nested Section
	Section []CompositionSection `json:"section,omitempty"`
}

// SectionMode represents the mode of a section
type SectionMode string

const (
	SectionModeWorking  SectionMode = "working"
	SectionModeSnapshot SectionMode = "snapshot"
	SectionModeChanges  SectionMode = "changes"
)
