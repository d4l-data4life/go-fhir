package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ManufacturedItemDefinitionStatus represents the status of a manufactured item definition
type ManufacturedItemDefinitionStatus string

const (
	ManufacturedItemDefinitionStatusDraft   ManufacturedItemDefinitionStatus = "draft"
	ManufacturedItemDefinitionStatusActive  ManufacturedItemDefinitionStatus = "active"
	ManufacturedItemDefinitionStatusRetired ManufacturedItemDefinitionStatus = "retired"
	ManufacturedItemDefinitionStatusUnknown ManufacturedItemDefinitionStatus = "unknown"
)

// ManufacturedItemDefinitionProperty represents general characteristics of an item
type ManufacturedItemDefinitionProperty struct {
	common.BackboneElement

	// A code expressing the type of characteristic
	Type common.CodeableConcept `json:"type"`

	// A value for the characteristic
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueDate            *string                 `json:"valueDate,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueMarkdown        *string                 `json:"valueMarkdown,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
}

// ManufacturedItemDefinitionComponentConstituent represents a constituent of a component
type ManufacturedItemDefinitionComponentConstituent struct {
	common.BackboneElement

	// The measurable amount of the constituent within the component
	Amount []common.Quantity `json:"amount,omitempty"`

	// A reference to a constituent of the manufactured item
	Constituent CodeableReference `json:"constituent"`

	// A reference to a component of a constituent within this component
	Component []ManufacturedItemDefinitionComponent `json:"component,omitempty"`

	// The function of this constituent within the component
	Function []common.CodeableConcept `json:"function,omitempty"`

	// The physical location of the constituent within the component
	Location []common.CodeableConcept `json:"location,omitempty"`
}

// ManufacturedItemDefinitionComponent represents physical parts of the manufactured item
type ManufacturedItemDefinitionComponent struct {
	common.BackboneElement

	// The measurable amount of total quantity of all substances in the component
	Amount []common.Quantity `json:"amount,omitempty"`

	// A component that this component contains or is made from
	Component []ManufacturedItemDefinitionComponent `json:"component,omitempty"`

	// A reference to a constituent of the manufactured item
	Constituent []ManufacturedItemDefinitionComponentConstituent `json:"constituent,omitempty"`

	// The function of this component within the item
	Function []common.CodeableConcept `json:"function,omitempty"`

	// General characteristics of this component
	Property []ManufacturedItemDefinitionProperty `json:"property,omitempty"`

	// Defining type of the component e.g. shell, layer, ink
	Type common.CodeableConcept `json:"type"`
}

// ManufacturedItemDefinition represents the definition and characteristics of a medicinal manufactured item
type ManufacturedItemDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ManufacturedItemDefinition"

	// Physical parts of the manufactured item
	Component []ManufacturedItemDefinitionComponent `json:"component,omitempty"`

	// Unique identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The ingredients of this manufactured item
	Ingredient []common.CodeableConcept `json:"ingredient,omitempty"`

	// Dose form as manufactured and before any transformation into the pharmaceutical product
	ManufacturedDoseForm common.CodeableConcept `json:"manufacturedDoseForm"`

	// Manufacturer of the item, one of several possible
	Manufacturer []common.Reference `json:"manufacturer,omitempty"`

	// Allows specifying that an item is on the market for sale
	MarketingStatus []interface{} `json:"marketingStatus,omitempty"`

	// A descriptive name applied to this item
	Name *string `json:"name,omitempty"`

	// General characteristics of this item
	Property []ManufacturedItemDefinitionProperty `json:"property,omitempty"`

	// Allows filtering of manufactured items that are appropriate for use versus not
	Status ManufacturedItemDefinitionStatus `json:"status"`

	// The "real-world" units in which the quantity of the manufactured item is described
	UnitOfPresentation *common.CodeableConcept `json:"unitOfPresentation,omitempty"`
}
