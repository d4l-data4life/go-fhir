// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// RequestOrchestration represents a set of related requests that can be used to capture and track the state of a request
type RequestOrchestration struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "RequestOrchestration"

	// The actions, if any, produced by the evaluation of the artifact
	Action []interface{} `json:"action,omitempty"`

	// The author of the request orchestration
	Author *common.Reference `json:"author,omitempty"`

	// A plan or request that is fulfilled in whole or in part by this request orchestration
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The date when this request orchestration was created
	Created *string `json:"created,omitempty"`

	// A group of related requests that can be used to capture and track the state of a request
	GroupIdentifier *common.Identifier `json:"groupIdentifier,omitempty"`

	// Business identifier for the request orchestration
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A reference to a formal or informal definition of the request orchestration
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// A reference to a formal or informal definition of the request orchestration
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// The intended performer of the request orchestration
	Intent string `json:"intent"`

	// A description of the request orchestration
	Note []Annotation `json:"note,omitempty"`

	// The priority of the request orchestration
	Priority *string `json:"priority,omitempty"`

	// A reference to a formal or informal definition of the request orchestration
	Replaces []common.Reference `json:"replaces,omitempty"`

	// The current state of the request orchestration
	Status string `json:"status"`

	// The subject of the request orchestration
	Subject *common.Reference `json:"subject,omitempty"`

	// A human-readable narrative that contains a summary of the request orchestration
	Text *Narrative `json:"text,omitempty"`
}
