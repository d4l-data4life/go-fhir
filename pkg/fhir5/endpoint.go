// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// EndpointStatus represents the status of an endpoint
type EndpointStatus string

const (
	EndpointStatusActive         EndpointStatus = "active"
	EndpointStatusSuspended      EndpointStatus = "suspended"
	EndpointStatusError          EndpointStatus = "error"
	EndpointStatusOff            EndpointStatus = "off"
	EndpointStatusEnteredInError EndpointStatus = "entered-in-error"
)

// EndpointPayload represents the payload type and format for the endpoint
type EndpointPayload struct {
	common.BackboneElement

	// MIME types supported by the endpoint
	MimeType []string `json:"mimeType,omitempty"`

	// The payload type
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// Endpoint represents a technical endpoint providing access to services
type Endpoint struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Endpoint"

	// The technical base address for connecting to this endpoint
	Address string `json:"address"`

	// Connection type details
	ConnectionType []common.CodeableConcept `json:"connectionType"`

	// Contact details for endpoint
	Contact []ContactPoint `json:"contact,omitempty"`

	// Description of the endpoint
	Description *string `json:"description,omitempty"`

	// Environment type
	EnvironmentType []common.CodeableConcept `json:"environmentType,omitempty"`

	// Usage depends on the channel type
	Header []string `json:"header,omitempty"`

	// Identifier for the endpoint
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Organization that manages this endpoint
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// A name that this endpoint can be referred to with
	Name *string `json:"name,omitempty"`

	// Set of payloads that are provided by this endpoint
	Payload []EndpointPayload `json:"payload,omitempty"`

	// Interval during which the endpoint is expected to be operational
	Period *common.Period `json:"period,omitempty"`

	// Active | suspended | error | off | entered-in-error
	Status EndpointStatus `json:"status"`
}
