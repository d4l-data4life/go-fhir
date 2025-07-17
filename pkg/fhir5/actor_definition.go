// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ActorType represents whether the actor is a person or system
type ActorType string

const (
	ActorTypePerson ActorType = "person"
	ActorTypeSystem ActorType = "system"
)

// ActorDefinition represents a description of an actor
type ActorDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ActorDefinition"

	// The capability statement for the actor (if the concept is applicable)
	Capabilities *string `json:"capabilities,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// Use and/or publishing restrictions
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// Date this was last changed
	Date *string `json:"date,omitempty"`

	// A url that identifies the definition of this actor in another IG
	DerivedFrom []string `json:"derivedFrom,omitempty"`

	// Natural language description of the actor
	Description *string `json:"description,omitempty"`

	// What the actor does (or is expected to do)
	Documentation *string `json:"documentation,omitempty"`

	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`

	// Additional identifier for the actor definition
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Intended jurisdiction for actor definition (if applicable)
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Name for this actor definition (computer friendly)
	Name *string `json:"name,omitempty"`

	// Name of the publisher (organization or individual)
	Publisher *string `json:"publisher,omitempty"`

	// Why this actor definition is defined
	Purpose *string `json:"purpose,omitempty"`

	// A reference to additional documentation about the actor
	Reference []string `json:"reference,omitempty"`

	// draft | active | retired | unknown
	Status RequestStatus `json:"status"`

	// Name for this actor definition (human friendly)
	Title *string `json:"title,omitempty"`

	// Whether the actor represents a human or an application
	Type ActorType `json:"type"`

	// Canonical identifier for this actor definition
	URL *string `json:"url,omitempty"`

	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`

	// Business version of the actor definition
	Version *string `json:"version,omitempty"`

	// How to compare versions
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}
