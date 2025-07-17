package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// DocumentManifest represents a collection of documents compiled for a purpose
type DocumentManifest struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DocumentManifest"

	// Not necessarily who did the actual data entry (i.e. typist) or who was the source (informant)
	Author []common.Reference `json:"author,omitempty"`

	// When used for XDS the intended focus of the DocumentManifest is for the reference to target to be a set of DocumentReference Resources
	Content []common.Reference `json:"content"`

	// Creation time is used for tracking, organizing versions and searching
	Created *string `json:"created,omitempty"`

	// What the document is about, rather than a terse summary of the document
	Description *string `json:"description,omitempty"`

	// Other identifiers associated with the document manifest, including version independent identifiers
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A single identifier that uniquely identifies this manifest
	MasterIdentifier *common.Identifier `json:"masterIdentifier,omitempty"`

	// How the recipient receives the document set or is notified of it is up to the implementation
	Recipient []common.Reference `json:"recipient,omitempty"`

	// May be identifiers or resources that caused the DocumentManifest to be created
	Related []DocumentManifestRelated `json:"related,omitempty"`

	// Identifies the source system, application, or software that produced the document manifest
	Source *string `json:"source,omitempty"`

	// This element is labeled as a modifier because the status contains the codes that mark the manifest as not currently valid
	Status DocumentManifestStatus `json:"status"`

	// Who or what the set of documents is about
	Subject *common.Reference `json:"subject,omitempty"`

	// Specifies the kind of this set of documents (e.g. Patient Summary, Discharge Summary, Prescription, etc.)
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// DocumentManifestStatus represents the status of the document manifest
type DocumentManifestStatus string

const (
	DocumentManifestStatusCurrent        DocumentManifestStatus = "current"
	DocumentManifestStatusSuperseded     DocumentManifestStatus = "superseded"
	DocumentManifestStatusEnteredInError DocumentManifestStatus = "entered-in-error"
)

// DocumentManifestRelated represents identifiers or resources that caused the DocumentManifest to be created
type DocumentManifestRelated struct {
	common.BackboneElement

	// If both identifier and ref elements are present they shall refer to the same thing
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// If both identifier and ref elements are present they shall refer to the same thing
	Ref *common.Reference `json:"ref,omitempty"`
}
