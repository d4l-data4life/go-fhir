package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// NamingSystemStatus represents status of naming system
type NamingSystemStatus string

const (
	NamingSystemStatusDraft   NamingSystemStatus = "draft"
	NamingSystemStatusActive  NamingSystemStatus = "active"
	NamingSystemStatusRetired NamingSystemStatus = "retired"
	NamingSystemStatusUnknown NamingSystemStatus = "unknown"
)

// NamingSystemKind represents kind of naming system
type NamingSystemKind string

const (
	NamingSystemKindCodesystem NamingSystemKind = "codesystem"
	NamingSystemKindIdentifier NamingSystemKind = "identifier"
	NamingSystemKindRoot       NamingSystemKind = "root"
)

// NamingSystemUniqueIdType represents type of unique id
type NamingSystemUniqueIdType string

const (
	NamingSystemUniqueIdTypeOID          NamingSystemUniqueIdType = "oid"
	NamingSystemUniqueIdTypeUUID         NamingSystemUniqueIdType = "uuid"
	NamingSystemUniqueIdTypeURI          NamingSystemUniqueIdType = "uri"
	NamingSystemUniqueIdTypeIRIStem      NamingSystemUniqueIdType = "iri-stem"
	NamingSystemUniqueIdTypeV2CSMnemonic NamingSystemUniqueIdType = "v2csmnemonic"
	NamingSystemUniqueIdTypeOther        NamingSystemUniqueIdType = "other"
)

// NamingSystemUniqueId represents unique identifier for naming system
type NamingSystemUniqueId struct {
	common.BackboneElement

	// Indicates whether this identifier is endorsed by the official owner
	Authoritative *bool `json:"authoritative,omitempty"`

	// Comment about the identifier
	Comment *string `json:"comment,omitempty"`

	// The period of time this identifier is valid
	Period *common.Period `json:"period,omitempty"`

	// Indicates whether this identifier is the "preferred" identifier
	Preferred *bool `json:"preferred,omitempty"`

	// The type of identifier
	Type NamingSystemUniqueIdType `json:"type"`

	// The string that should be sent over the wire
	Value string `json:"value"`
}

// NamingSystem represents a namespace that issues unique symbols
type NamingSystem struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "NamingSystem"

	// The approval date
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance
	Author []ContactDetail `json:"author,omitempty"`

	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`

	// Copyright statement
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The date when the naming system was published
	Date string `json:"date"`

	// Natural language description of the naming system
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence
	Editor []ContactDetail `json:"editor,omitempty"`

	// When the naming system is expected to be used
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// Who endorsed the naming system
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Intended jurisdiction for naming system
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Indicates the purpose for the naming system
	Kind NamingSystemKind `json:"kind"`

	// Last review date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// Name for this naming system
	Name string `json:"name"`

	// Name of the publisher
	Publisher *string `json:"publisher,omitempty"`

	// Why this naming system is defined
	Purpose *string `json:"purpose,omitempty"`

	// Additional documentation, citations, etc.
	RelatedArtifact []interface{} `json:"relatedArtifact,omitempty"`

	// Organization that is responsible for issuing identifiers or codes
	Responsible *string `json:"responsible,omitempty"`

	// Who reviewed the naming system
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// draft | active | retired | unknown
	Status NamingSystemStatus `json:"status"`

	// Human-readable title
	Title *string `json:"title,omitempty"`

	// E.g. Education, Treatment, Assessment, etc.
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// Categorizes a naming system for easier search
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Indicates how the system may be identified when referenced
	UniqueId []NamingSystemUniqueId `json:"uniqueId"`

	// Canonical identifier for this naming system
	Url *string `json:"url,omitempty"`

	// Provides guidance on the use of the namespace
	Usage *string `json:"usage,omitempty"`

	// The context that the content is intended to support
	UseContext []interface{} `json:"useContext,omitempty"`
}
