package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// NamingSystem represents a curated namespace that issues unique symbols within that namespace
type NamingSystem struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "NamingSystem"

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date string `json:"date"`

	// This description can be used to capture details such as why the naming system was built
	Description *string `json:"description,omitempty"`

	// It may be possible for the naming system to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Indicates the purpose for the naming system - what kinds of things does it make unique?
	Kind NamingSystemKind `json:"kind"`

	// The name is not expected to be globally unique
	Name string `json:"name"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This is the primary organization
	Responsible *string `json:"responsible,omitempty"`

	// Allows filtering of naming systems that are appropriate for use versus not
	Status NamingSystemStatus `json:"status"`

	// This will most commonly be used for identifier namespaces
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Multiple identifiers may exist, either due to duplicate registration, regional rules, needs of different communication technologies, etc
	UniqueId []NamingSystemUniqueId `json:"uniqueId"`

	// Provides guidance on the use of the namespace
	Usage *string `json:"usage,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`
}

// NamingSystemKind represents the purpose for the naming system
type NamingSystemKind string

const (
	NamingSystemKindCodesystem NamingSystemKind = "codesystem"
	NamingSystemKindIdentifier NamingSystemKind = "identifier"
	NamingSystemKindRoot       NamingSystemKind = "root"
)

// NamingSystemStatus represents the status of the naming system
type NamingSystemStatus string

const (
	NamingSystemStatusDraft   NamingSystemStatus = "draft"
	NamingSystemStatusActive  NamingSystemStatus = "active"
	NamingSystemStatusRetired NamingSystemStatus = "retired"
	NamingSystemStatusUnknown NamingSystemStatus = "unknown"
)

// NamingSystemUniqueId represents multiple identifiers may exist
type NamingSystemUniqueId struct {
	common.BackboneElement

	// e.g. "must be used in Germany" or "was initially published in error with this value"
	Comment *string `json:"comment,omitempty"`

	// Within a registry, a given identifier should only be "active" for a single namespace at a time
	Period *common.Period `json:"period,omitempty"`

	// Indicates whether this identifier is the "preferred" identifier of this type
	Preferred *bool `json:"preferred,omitempty"`

	// Different identifier types may be used in different types of communications
	Type NamingSystemUniqueIdType `json:"type"`

	// If the value is a URI intended for use as FHIR system identifier
	Value string `json:"value"`
}

// NamingSystemUniqueIdType represents different identifier types
type NamingSystemUniqueIdType string

const (
	NamingSystemUniqueIdTypeOid   NamingSystemUniqueIdType = "oid"
	NamingSystemUniqueIdTypeUuid  NamingSystemUniqueIdType = "uuid"
	NamingSystemUniqueIdTypeUri   NamingSystemUniqueIdType = "uri"
	NamingSystemUniqueIdTypeOther NamingSystemUniqueIdType = "other"
)
