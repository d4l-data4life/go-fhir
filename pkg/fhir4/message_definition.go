package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MessageDefinition represents defines the characteristics of a message that can be shared between systems
type MessageDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MessageDefinition"

	// This indicates an application level response to "close" a transaction implicit in a particular request message
	AllowedResponse []MessageDefinitionAllowedResponse `json:"allowedResponse,omitempty"`

	// The MessageDefinition that is the basis for the contents of this resource
	Base *string `json:"base,omitempty"`

	// The impact of the content of the message
	Category *MessageDefinitionCategory `json:"category,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the message definition and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date string `json:"date"`

	// This description can be used to capture details such as why the message definition was built
	Description *string `json:"description,omitempty"`

	// Event code or link to the EventDefinition
	EventCoding *common.Coding `json:"eventCoding,omitempty"`

	// Event code or link to the EventDefinition
	EventUri *string `json:"eventUri,omitempty"`

	// Allows filtering of message definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Identifies the resource (or resources) that are being addressed by the event
	Focus []MessageDefinitionFocus `json:"focus,omitempty"`

	// Canonical reference to a GraphDefinition
	Graph []string `json:"graph,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the message definition to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// It should be possible to use MessageDefinition to describe a message to be used by certain steps in a particular protocol
	Parent []string `json:"parent,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the message definition
	Purpose *string `json:"purpose,omitempty"`

	// A MessageDefinition that is superseded by this definition
	Replaces []string `json:"replaces,omitempty"`

	// This enables the capability currently available through MSH-16 (Application Level acknowledgement) in HL7 Version 2
	ResponseRequired *MessageDefinitionResponseRequired `json:"responseRequired,omitempty"`

	// Allows filtering of message definitions that are appropriate for use versus not
	Status MessageDefinitionStatus `json:"status"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different message definition instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// MessageDefinitionStatus represents the status of the message definition
type MessageDefinitionStatus string

const (
	MessageDefinitionStatusDraft   MessageDefinitionStatus = "draft"
	MessageDefinitionStatusActive  MessageDefinitionStatus = "active"
	MessageDefinitionStatusRetired MessageDefinitionStatus = "retired"
	MessageDefinitionStatusUnknown MessageDefinitionStatus = "unknown"
)

// MessageDefinitionCategory represents the impact of the content of the message
type MessageDefinitionCategory string

const (
	MessageDefinitionCategoryConsequence  MessageDefinitionCategory = "consequence"
	MessageDefinitionCategoryCurrency     MessageDefinitionCategory = "currency"
	MessageDefinitionCategoryNotification MessageDefinitionCategory = "notification"
)

// MessageDefinitionResponseRequired represents the response requirement
type MessageDefinitionResponseRequired string

const (
	MessageDefinitionResponseRequiredAlways    MessageDefinitionResponseRequired = "always"
	MessageDefinitionResponseRequiredOnError   MessageDefinitionResponseRequired = "on-error"
	MessageDefinitionResponseRequiredNever     MessageDefinitionResponseRequired = "never"
	MessageDefinitionResponseRequiredOnSuccess MessageDefinitionResponseRequired = "on-success"
)

// MessageDefinitionFocus represents identifies the resource (or resources) that are being addressed by the event
type MessageDefinitionFocus struct {
	common.BackboneElement

	// Multiple focuses addressing different resources may occasionally occur
	Code string `json:"code"`

	// Identifies the maximum number of resources of this type that must be pointed to by a message
	Max *string `json:"max,omitempty"`

	// Identifies the minimum number of resources of this type that must be pointed to by a message
	Min int `json:"min"`

	// This should be present for most message definitions
	Profile *string `json:"profile,omitempty"`
}

// MessageDefinitionAllowedResponse represents an application level response to "close" a transaction
type MessageDefinitionAllowedResponse struct {
	common.BackboneElement

	// A reference to the message definition that must be adhered to by this supported response
	Message string `json:"message"`

	// Provides a description of the circumstances in which this response should be used
	Situation *string `json:"situation,omitempty"`
}
