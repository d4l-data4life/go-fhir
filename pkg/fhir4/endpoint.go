package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Endpoint represents the technical details of an endpoint that can be used for electronic services
type Endpoint struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Endpoint"

	// For rest-hook, and websocket, the end-point must be an http: or https: URL
	Address string `json:"address"`

	// For additional connectivity details for the protocol, extensions will be used at this point
	ConnectionType common.Coding `json:"connectionType"`

	// Contact details for a human to contact about the subscription
	Contact []common.ContactPoint `json:"contact,omitempty"`

	// Exactly what these mean depends on the channel type
	Header []string `json:"header,omitempty"`

	// Identifier for the organization that is used to identify the endpoint across multiple disparate systems
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// This property is not typically used when searching for Endpoint resources for usage
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// A friendly name that this endpoint can be referred to with
	Name *string `json:"name,omitempty"`

	// Sending the payload has obvious security consequences
	PayloadMimeType []string `json:"payloadMimeType,omitempty"`

	// The payloadFormat describes the serialization format of the data
	PayloadType []common.CodeableConcept `json:"payloadType"`

	// The interval during which the endpoint is expected to be operational
	Period *common.Period `json:"period,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the endpoint as not currently valid
	Status EndpointStatus `json:"status"`
}

// EndpointStatus represents the status of the endpoint
type EndpointStatus string

const (
	EndpointStatusActive         EndpointStatus = "active"
	EndpointStatusSuspended      EndpointStatus = "suspended"
	EndpointStatusError          EndpointStatus = "error"
	EndpointStatusOff            EndpointStatus = "off"
	EndpointStatusEnteredInError EndpointStatus = "entered-in-error"
	EndpointStatusTest           EndpointStatus = "test"
)
