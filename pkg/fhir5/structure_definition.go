// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// StructureDefinitionMapping represents a mapping from one set of concepts to one or more other concepts
type StructureDefinitionMapping struct {
	common.BackboneElement
	Comment  *string `json:"comment,omitempty"`
	Identity string  `json:"identity"`
	Name     *string `json:"name,omitempty"`
	URI      *string `json:"uri,omitempty"`
}

// StructureDefinitionContext represents a context in which an extension can be used
type StructureDefinitionContext struct {
	common.BackboneElement
	Expression string `json:"expression"`
	Type       string `json:"type"`
}

// StructureDefinitionSnapshot represents a snapshot view of a structure definition
type StructureDefinitionSnapshot struct {
	common.BackboneElement
	Element []interface{} `json:"element"`
}

// StructureDefinitionDifferential represents a differential view of a structure definition
type StructureDefinitionDifferential struct {
	common.BackboneElement
	Element []interface{} `json:"element"`
}

// StructureDefinition represents a FHIR structure definition
type StructureDefinition struct {
	DomainResource
	ResourceType    string                           `json:"resourceType"` // Always "StructureDefinition"
	Abstract        bool                             `json:"abstract"`
	BaseDefinition  *string                          `json:"baseDefinition,omitempty"`
	Contact         []ContactDetail                  `json:"contact,omitempty"`
	Context         []StructureDefinitionContext     `json:"context,omitempty"`
	Differential    *StructureDefinitionDifferential `json:"differential,omitempty"`
	EffectivePeriod *common.Period                   `json:"effectivePeriod,omitempty"`
	Experimental    *bool                            `json:"experimental,omitempty"`
	Identifier      []common.Identifier              `json:"identifier,omitempty"`
	Jurisdiction    []common.CodeableConcept         `json:"jurisdiction,omitempty"`
	LastReviewDate  *string                          `json:"lastReviewDate,omitempty"`
	Mapping         []StructureDefinitionMapping     `json:"mapping,omitempty"`
	Name            string                           `json:"name"`
	Publisher       *string                          `json:"publisher,omitempty"`
	Purpose         *string                          `json:"purpose,omitempty"`
	Reviewer        []ContactDetail                  `json:"reviewer,omitempty"`
	Snapshot        *StructureDefinitionSnapshot     `json:"snapshot,omitempty"`
	Status          string                           `json:"status"`
	Type            string                           `json:"type"`
	Url             string                           `json:"url"`
	Usage           *string                          `json:"usage,omitempty"`
	UseContext      []UsageContext                   `json:"useContext,omitempty"`
	Version         *string                          `json:"version,omitempty"`
	FhirVersion     string                           `json:"fhirVersion"`
	Kind            string                           `json:"kind"`
	Primitive       *bool                            `json:"primitive,omitempty"`
	Extension       []string                         `json:"extension,omitempty"`
	Key             *string                          `json:"key,omitempty"`
	Target          []string                         `json:"target,omitempty"`
}
