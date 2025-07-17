package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Subscription represents the subscription resource is used to define a push-based subscription from a server to another system
type Subscription struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Subscription"

	// Details where to send notifications when resources are received that meet the criteria
	Channel SubscriptionChannel `json:"channel"`

	// Contact details for a human to contact about the subscription
	Contact []common.ContactPoint `json:"contact,omitempty"`

	// The rules are search criteria (without the [base] part)
	Criteria string `json:"criteria"`

	// The server is permitted to deviate from this time but should observe it
	End *string `json:"end,omitempty"`

	// A record of the last error that occurred when the server processed a notification
	Error *string `json:"error,omitempty"`

	// A description of why this subscription is defined
	Reason string `json:"reason"`

	// A client can only submit subscription resources in the requested or off state
	Status SubscriptionStatus `json:"status"`
}

// SubscriptionStatus represents a client can only submit subscription resources in the requested or off state
type SubscriptionStatus string

const (
	SubscriptionStatusRequested SubscriptionStatus = "requested"
	SubscriptionStatusActive    SubscriptionStatus = "active"
	SubscriptionStatusError     SubscriptionStatus = "error"
	SubscriptionStatusOff       SubscriptionStatus = "off"
)

// SubscriptionChannel represents details where to send notifications when resources are received that meet the criteria
type SubscriptionChannel struct {
	common.BackboneElement

	// For rest-hook, and websocket, the end-point must be an http: or https: URL
	Endpoint *string `json:"endpoint,omitempty"`

	// Exactly what these mean depend on the channel type
	Header []string `json:"header,omitempty"`

	// Sending the payload has obvious security implications
	Payload *string `json:"payload,omitempty"`

	// The type of channel to send notifications on
	Type SubscriptionChannelType `json:"type"`
}

// SubscriptionChannelType represents the type of channel to send notifications on
type SubscriptionChannelType string

const (
	SubscriptionChannelTypeRestHook  SubscriptionChannelType = "rest-hook"
	SubscriptionChannelTypeWebsocket SubscriptionChannelType = "websocket"
	SubscriptionChannelTypeEmail     SubscriptionChannelType = "email"
	SubscriptionChannelTypeSms       SubscriptionChannelType = "sms"
	SubscriptionChannelTypeMessage   SubscriptionChannelType = "message"
)
