package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// PackagedProductDefinitionPackagingProperty represents general characteristics of this item
type PackagedProductDefinitionPackagingProperty struct {
	common.BackboneElement

	// A code expressing the type of characteristic
	Type common.CodeableConcept `json:"type"`

	// A value for the characteristic
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueDate            *string                 `json:"valueDate,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
}

// PackagedProductDefinitionPackagingContainedItem represents the item(s) within the packaging
type PackagedProductDefinitionPackagingContainedItem struct {
	common.BackboneElement

	// The number of this type of item within this packaging
	Amount *common.Quantity `json:"amount,omitempty"`

	// The actual item(s) of medication, as manufactured
	Item CodeableReference `json:"item"`
}

// PackagedProductDefinitionPackaging represents a packaging item, as a container for medically related items
type PackagedProductDefinitionPackaging struct {
	common.BackboneElement

	// A possible alternate material for this part of the packaging
	AlternateMaterial []common.CodeableConcept `json:"alternateMaterial,omitempty"`

	// Is this a part of the packaging rather than the packaging itself
	ComponentPart *bool `json:"componentPart,omitempty"`

	// The item(s) within the packaging
	ContainedItem []PackagedProductDefinitionPackagingContainedItem `json:"containedItem,omitempty"`

	// A business identifier that is specific to this particular part of the packaging
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Manufacturer of this packaging item
	Manufacturer []common.Reference `json:"manufacturer,omitempty"`

	// Material type of the package item
	Material []common.CodeableConcept `json:"material,omitempty"`

	// Allows containers within containers
	Packaging []PackagedProductDefinitionPackaging `json:"packaging,omitempty"`

	// General characteristics of this item
	Property []PackagedProductDefinitionPackagingProperty `json:"property,omitempty"`
}

// PackagedProductDefinition represents a medically related item or items, in a container or package
type PackagedProductDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "PackagedProductDefinition"

	// Additional information or supporting documentation
	AttachedDocument []common.Reference `json:"attachedDocument,omitempty"`

	// Allows the key product features to be recorded
	Characteristic []PackagedProductDefinitionPackagingProperty `json:"characteristic,omitempty"`

	// A total of the complete count of contained items
	ContainedItemQuantity []common.Quantity `json:"containedItemQuantity,omitempty"`

	// Identifies if the package contains different items
	CopackagedIndicator *bool `json:"copackagedIndicator,omitempty"`

	// Textual description
	Description *string `json:"description,omitempty"`

	// A unique identifier for this package as whole
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The legal status of supply of the packaged item
	LegalStatusOfSupply []interface{} `json:"legalStatusOfSupply,omitempty"`

	// Manufacturer of this package type
	Manufacturer []common.Reference `json:"manufacturer,omitempty"`

	// Allows specifying that an item is on the market for sale
	MarketingStatus []interface{} `json:"marketingStatus,omitempty"`

	// A name for this package
	Name *string `json:"name,omitempty"`

	// The product this package model relates to
	PackageFor []common.Reference `json:"packageFor,omitempty"`

	// A packaging item, as a container for medically related items
	Packaging *PackagedProductDefinitionPackaging `json:"packaging,omitempty"`

	// The status within the lifecycle of this item
	Status *common.CodeableConcept `json:"status,omitempty"`

	// The date at which the given status became applicable
	StatusDate *string `json:"statusDate,omitempty"`

	// A high level category
	Type *common.CodeableConcept `json:"type,omitempty"`
}
