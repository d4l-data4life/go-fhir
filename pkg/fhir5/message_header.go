package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MessageHeaderDestination represents message destinations
type MessageHeaderDestination struct {
	common.BackboneElement

	// URL for the destination
	EndpointUrl *string `json:"endpointUrl,omitempty"`

	// Reference to the destination endpoint
	EndpointReference *common.Reference `json:"endpointReference,omitempty"`

	// Name of the destination
	Name *string `json:"name,omitempty"`

	// Specific person or department to receive the message
	Receiver *common.Reference `json:"receiver,omitempty"`

	// Target system for the message
	Target *common.Reference `json:"target,omitempty"`
}

// MessageHeaderSource represents message source
type MessageHeaderSource struct {
	common.BackboneElement

	// Contact information for the source
	Contact *ContactPoint `json:"contact,omitempty"`

	// URL for the source
	EndpointUrl *string `json:"endpointUrl,omitempty"`

	// Reference to the source endpoint
	EndpointReference *common.Reference `json:"endpointReference,omitempty"`

	// Name of the source
	Name *string `json:"name,omitempty"`

	// Software running at the source
	Software *string `json:"software,omitempty"`

	// Version of the software
	Version *string `json:"version,omitempty"`
}

// MessageHeaderResponseCode represents response codes
type MessageHeaderResponseCode string

const (
	MessageHeaderResponseCodeOK             MessageHeaderResponseCode = "ok"
	MessageHeaderResponseCodeTransientError MessageHeaderResponseCode = "transient-error"
	MessageHeaderResponseCodeFatalError     MessageHeaderResponseCode = "fatal-error"
)

// MessageHeaderResponse represents response information
type MessageHeaderResponse struct {
	common.BackboneElement

	// Response code
	Code MessageHeaderResponseCode `json:"code"`

	// Details about the response
	Details *common.Reference `json:"details,omitempty"`

	// Identifier of the message being responded to
	Identifier common.Identifier `json:"identifier"`
}

// MessageHeader represents the header for a message exchange
type MessageHeader struct {
	DomainResource

	// Resource type
	ResourceType string `json:"resourceType"`

	// Author of the message
	Author *common.Reference `json:"author,omitempty"`

	// Link to the message definition
	Definition *string `json:"definition,omitempty"`

	// Message destinations
	Destination []MessageHeaderDestination `json:"destination,omitempty"`

	// Event the message represents (coding)
	EventCoding *common.Coding `json:"eventCoding,omitempty"`

	// Event the message represents (canonical)
	EventCanonical *string `json:"eventCanonical,omitempty"`

	// The actual data content
	Focus []common.Reference `json:"focus,omitempty"`

	// Cause of the event
	Reason *common.CodeableConcept `json:"reason,omitempty"`

	// Response information
	Response *MessageHeaderResponse `json:"response,omitempty"`

	// Responsible party
	Responsible *common.Reference `json:"responsible,omitempty"`

	// Message sender
	Sender *common.Reference `json:"sender,omitempty"`

	// Message source
	Source MessageHeaderSource `json:"source"`
}
