// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// Basic represents a resource that is used to represent concepts not yet defined in FHIR
type Basic struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Basic"

	// Indicates who was responsible for creating the resource instance
	Author *common.Reference `json:"author,omitempty"`

	// The type of resource being represented - this defines the meaning of the resource and cannot be ignored
	Code common.CodeableConcept `json:"code"`

	// Identifies when the resource was first created
	Created *string `json:"created,omitempty"`

	// Identifier assigned to the resource for business purposes, outside the context of FHIR
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Optional as not all potential resources will have subjects
	Subject *common.Reference `json:"subject,omitempty"`
}
