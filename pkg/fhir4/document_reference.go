// Package fhir4 contains FHIR R4 (version 4.0.1) resource definitions
package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// DocumentReference represents a reference to a document
type DocumentReference struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DocumentReference"

	// Master Version Specific Identifier
	MasterIdentifier *common.Identifier `json:"masterIdentifier,omitempty"`

	// Other identifiers for the document
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// current | superseded | entered-in-error
	Status DocumentReferenceStatus `json:"status"`

	// preliminary | final | amended | entered-in-error
	DocStatus *DocumentReferenceDocStatus `json:"docStatus,omitempty"`

	// Kind of document (LOINC if possible)
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Categorization of document
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Who/what is the subject of the document
	Subject *common.Reference `json:"subject,omitempty"`

	// When this document reference was created
	Date *string `json:"date,omitempty"`

	// Who and/or what authored the document
	Author []common.Reference `json:"author,omitempty"`

	// Who/what authenticated the document
	Authenticator *common.Reference `json:"authenticator,omitempty"`

	// Organization which maintains the document
	Custodian *common.Reference `json:"custodian,omitempty"`

	// Relationships to other documents
	RelatesTo []DocumentReferenceRelatesTo `json:"relatesTo,omitempty"`

	// Human-readable description
	Description *string `json:"description,omitempty"`

	// Document security-tags
	SecurityLabel []common.CodeableConcept `json:"securityLabel,omitempty"`

	// Document referenced
	Content []DocumentReferenceContent `json:"content"`

	// Clinical context of document
	Context *DocumentReferenceContext `json:"context,omitempty"`
}

// DocumentReferenceStatus represents the status of the document reference
type DocumentReferenceStatus string

const (
	DocumentReferenceStatusCurrent        DocumentReferenceStatus = "current"
	DocumentReferenceStatusSuperseded     DocumentReferenceStatus = "superseded"
	DocumentReferenceStatusEnteredInError DocumentReferenceStatus = "entered-in-error"
)

// DocumentReferenceDocStatus represents the document status
type DocumentReferenceDocStatus string

const (
	DocumentReferenceDocStatusPreliminary    DocumentReferenceDocStatus = "preliminary"
	DocumentReferenceDocStatusFinal          DocumentReferenceDocStatus = "final"
	DocumentReferenceDocStatusAmended        DocumentReferenceDocStatus = "amended"
	DocumentReferenceDocStatusEnteredInError DocumentReferenceDocStatus = "entered-in-error"
)

// DocumentReferenceRelatesTo represents relationships to other documents
type DocumentReferenceRelatesTo struct {
	common.BackboneElement

	// replaces | transforms | signs | appends
	Code DocumentReferenceRelatesToCode `json:"code"`

	// Target of the relationship
	Target common.Reference `json:"target"`
}

// DocumentReferenceRelatesToCode represents the code for relationships to other documents
type DocumentReferenceRelatesToCode string

const (
	DocumentReferenceRelatesToCodeReplaces   DocumentReferenceRelatesToCode = "replaces"
	DocumentReferenceRelatesToCodeTransforms DocumentReferenceRelatesToCode = "transforms"
	DocumentReferenceRelatesToCodeSigns      DocumentReferenceRelatesToCode = "signs"
	DocumentReferenceRelatesToCodeAppends    DocumentReferenceRelatesToCode = "appends"
)

// DocumentReferenceContent represents content of the document reference
type DocumentReferenceContent struct {
	common.BackboneElement

	// Where to access the document
	Attachment Attachment `json:"attachment"`

	// Format/content rules for the document
	Format *common.Coding `json:"format,omitempty"`
}

// DocumentReferenceContext represents clinical context of the document
type DocumentReferenceContext struct {
	common.BackboneElement

	// Context of the document content
	Encounter []common.Reference `json:"encounter,omitempty"`

	// Main clinical acts documented
	Event []common.CodeableConcept `json:"event,omitempty"`

	// Time of service that is being documented
	Period *common.Period `json:"period,omitempty"`

	// Kind of facility where patient was seen
	FacilityType *common.CodeableConcept `json:"facilityType,omitempty"`

	// Additional details about where the content was created (e.g. clinical specialty)
	PracticeSetting *common.CodeableConcept `json:"practiceSetting,omitempty"`

	// Patient demographics from source
	SourcePatientInfo *common.Reference `json:"sourcePatientInfo,omitempty"`

	// Related identifiers or resources
	Related []common.Reference `json:"related,omitempty"`
}
