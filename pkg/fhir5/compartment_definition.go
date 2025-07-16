// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// CompartmentDefinitionResource represents a resource that may be in a compartment
type CompartmentDefinitionResource struct {
	common.BackboneElement

	Code          string   `json:"code"`
	Param         []string `json:"param,omitempty"`
	Documentation *string  `json:"documentation,omitempty"`
}

// CompartmentDefinition represents a compartment definition
type CompartmentDefinition struct {
	DomainResource

	ResourceType string `json:"resourceType"` // Always "CompartmentDefinition"

	Url          string                          `json:"url"`
	Version      *string                         `json:"version,omitempty"`
	Name         string                          `json:"name"`
	Status       PublicationStatus               `json:"status"`
	Experimental *bool                           `json:"experimental,omitempty"`
	Date         string                          `json:"date"`
	Publisher    *string                         `json:"publisher,omitempty"`
	Contact      []ContactDetail                 `json:"contact,omitempty"`
	Description  *string                         `json:"description,omitempty"`
	Purpose      *string                         `json:"purpose,omitempty"`
	Code         string                          `json:"code"`
	Search       bool                            `json:"search"`
	Resource     []CompartmentDefinitionResource `json:"resource,omitempty"`
}
