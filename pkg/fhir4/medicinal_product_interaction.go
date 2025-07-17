package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MedicinalProductInteraction represents the interactions of the medicinal product with other medicinal products, or other forms of interactions
type MedicinalProductInteraction struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicinalProductInteraction"

	// The interaction described
	Description *string `json:"description,omitempty"`

	// The effect of the interaction, for example "reduced gastric absorption of primary medication"
	Effect *common.CodeableConcept `json:"effect,omitempty"`

	// The incidence of the interaction, e.g. theoretical, observed
	Incidence *common.CodeableConcept `json:"incidence,omitempty"`

	// The specific medication, food or laboratory test that interacts
	Interactant []MedicinalProductInteractionInteractant `json:"interactant,omitempty"`

	// Actions for managing the interaction
	Management *common.CodeableConcept `json:"management,omitempty"`

	// The medication for which this is a described interaction
	Subject []common.Reference `json:"subject,omitempty"`

	// The type of the interaction e.g. drug-drug interaction, drug-food interaction, drug-lab test interaction
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// MedicinalProductInteractionInteractant represents the specific medication, food or laboratory test that interacts
type MedicinalProductInteractionInteractant struct {
	common.BackboneElement

	// The specific medication, food or laboratory test that interacts
	ItemReference *common.Reference `json:"itemReference,omitempty"`

	// The specific medication, food or laboratory test that interacts
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
}
