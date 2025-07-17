package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MessageHeader represents the header for a message exchange that is either requesting or responding to an action
type MessageHeader struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MessageHeader"

	// Usually only for the request but can be used in a response
	Author *common.Reference `json:"author,omitempty"`

	// Permanent link to the MessageDefinition for this message
	Definition *string `json:"definition,omitempty"`

	// There SHOULD be at least one destination, but in some circumstances, the source system is unaware of any particular destination system
	Destination []MessageHeaderDestination `json:"destination,omitempty"`

	// Usually only for the request but can be used in a response
	Enterer *common.Reference `json:"enterer,omitempty"`

	// The time of the event will be found in the focus resource
	EventCoding *common.Coding `json:"eventCoding,omitempty"`

	// The time of the event will be found in the focus resource
	EventUri *string `json:"eventUri,omitempty"`

	// The data is defined where the transaction type is defined
	Focus []common.Reference `json:"focus,omitempty"`

	// Coded indication of the cause for the event - indicates a reason for the occurrence of the event that is a focus of this message
	Reason *common.CodeableConcept `json:"reason,omitempty"`

	// Information about the message that this message is a response to
	Response *MessageHeaderResponse `json:"response,omitempty"`

	// Usually only for the request but can be used in a response
	Responsible *common.Reference `json:"responsible,omitempty"`

	// Use case is for where a (trusted) sending system is responsible for multiple organizations
	Sender *common.Reference `json:"sender,omitempty"`

	// The source application from which this message originated
	Source MessageHeaderSource `json:"source"`
}

// MessageHeaderDestination represents there SHOULD be at least one destination
type MessageHeaderDestination struct {
	common.BackboneElement

	// The id may be a non-resolvable URI for systems that do not use standard network-based addresses
	Endpoint string `json:"endpoint"`

	// Human-readable name for the target system
	Name *string `json:"name,omitempty"`

	// Allows data conveyed by a message to be addressed to a particular person or department
	Receiver *common.Reference `json:"receiver,omitempty"`

	// Identifies the target end system in situations where the initial message transmission is to an intermediary system
	Target *common.Reference `json:"target,omitempty"`
}

// MessageHeaderSource represents the source application from which this message originated
type MessageHeaderSource struct {
	common.BackboneElement

	// An e-mail, phone, website or other contact point to use to resolve issues with message communications
	Contact *common.ContactPoint `json:"contact,omitempty"`

	// The id may be a non-resolvable URI for systems that do not use standard network-based addresses
	Endpoint string `json:"endpoint"`

	// Human-readable name for the source system
	Name *string `json:"name,omitempty"`

	// May include configuration or other information useful in debugging
	Software *string `json:"software,omitempty"`

	// Can convey versions of multiple systems in situations where a message passes through multiple hands
	Version *string `json:"version,omitempty"`
}

// MessageHeaderResponse represents information about the message that this message is a response to
type MessageHeaderResponse struct {
	common.BackboneElement

	// This is a generic response to the request message
	Code MessageHeaderResponseCode `json:"code"`

	// This SHALL be contained in the bundle
	Details *common.Reference `json:"details,omitempty"`

	// The MessageHeader.id of the message to which this message is a response
	Identifier string `json:"identifier"`
}

// MessageHeaderResponseCode represents the response code
type MessageHeaderResponseCode string

const (
	MessageHeaderResponseCodeOk             MessageHeaderResponseCode = "ok"
	MessageHeaderResponseCodeTransientError MessageHeaderResponseCode = "transient-error"
	MessageHeaderResponseCodeFatalError     MessageHeaderResponseCode = "fatal-error"
)
