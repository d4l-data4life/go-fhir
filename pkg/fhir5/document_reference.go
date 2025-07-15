// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// DocumentReferenceStatus represents the status of a document reference
type DocumentReferenceStatus string

const (
	DocumentReferenceStatusCurrent        DocumentReferenceStatus = "current"
	DocumentReferenceStatusSuperseded     DocumentReferenceStatus = "superseded"
	DocumentReferenceStatusEnteredInError DocumentReferenceStatus = "entered-in-error"
)

// DocumentReferenceAttester represents who attested the document
type DocumentReferenceAttester struct {
	common.BackboneElement

	// The type of attestation the authenticator offers
	Mode common.CodeableConcept `json:"mode"`

	// Who attested the document in the specified way
	Party *common.Reference `json:"party,omitempty"`

	// When the document was attested by the party
	Time *string `json:"time,omitempty"`
}

// DocumentReferenceRelatesTo represents relationships this document has with other documents
type DocumentReferenceRelatesTo struct {
	common.BackboneElement

	// The type of relationship between documents
	Code common.CodeableConcept `json:"code"`

	// The target document of this relationship
	Target common.Reference `json:"target"`
}

// DocumentReferenceContentProfile represents the profile information for document content
type DocumentReferenceContentProfile struct {
	common.BackboneElement

	// Profile information as a coding, URI, or canonical reference
	ValueCoding    *common.Coding `json:"valueCoding,omitempty"`
	ValueUri       *string        `json:"valueUri,omitempty"`
	ValueCanonical *string        `json:"valueCanonical,omitempty"`
}

// DocumentReferenceContent represents the content of the document
type DocumentReferenceContent struct {
	common.BackboneElement

	// The document or URL of the document with critical metadata
	Attachment Attachment `json:"attachment"`

	// Format/content rules for the document
	Profile []DocumentReferenceContentProfile `json:"profile,omitempty"`
}

// DocumentReference represents a reference to a document of any kind
type DocumentReference struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DocumentReference"

	// Who attested the document
	Attester []DocumentReferenceAttester `json:"attester,omitempty"`

	// Who and/or what authored the document
	Author []common.Reference `json:"author,omitempty"`

	// Procedure that is fulfilled in whole or part by the creation of this document
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Anatomic structures included in the document
	BodySite []CodeableReference `json:"bodySite,omitempty"`

	// Categorization of document
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Document referenced
	Content []DocumentReferenceContent `json:"content"`

	// Context of the document encounter
	Context []common.Reference `json:"context,omitempty"`

	// Organization which maintains the document
	Custodian *common.Reference `json:"custodian,omitempty"`

	// When this document reference was created
	Date *string `json:"date,omitempty"`

	// Human-readable description
	Description *string `json:"description,omitempty"`

	// Document security-tags
	Facility *common.CodeableConcept `json:"facility,omitempty"`

	// Business identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Time period over which the service described by the document was provided
	Period *common.Period `json:"period,omitempty"`

	// Additional details about where the content was created
	PracticeSetting *common.CodeableConcept `json:"practiceSetting,omitempty"`

	// Relationships to other documents
	RelatesTo []DocumentReferenceRelatesTo `json:"relatesTo,omitempty"`

	// Document security-tags
	SecurityLabel []common.CodeableConcept `json:"securityLabel,omitempty"`

	// Current | superseded | entered-in-error
	Status DocumentReferenceStatus `json:"status"`

	// Who/what is the subject of the document
	Subject *common.Reference `json:"subject,omitempty"`

	// Kind of document
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Document version
	Version *string `json:"version,omitempty"`
}
