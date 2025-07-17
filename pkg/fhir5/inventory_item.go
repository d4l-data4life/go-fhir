package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// InventoryItemStatus represents the status of an inventory item
type InventoryItemStatus string

const (
	InventoryItemStatusActive         InventoryItemStatus = "active"
	InventoryItemStatusInactive       InventoryItemStatus = "inactive"
	InventoryItemStatusEnteredInError InventoryItemStatus = "entered-in-error"
	InventoryItemStatusUnknown        InventoryItemStatus = "unknown"
)

// InventoryItemName represents the item name(s) - the brand name, or common name, functional name, generic name
type InventoryItemName struct {
	common.BackboneElement

	// The language that the item name is expressed in
	Language string `json:"language"`

	// The name or designation that the item is given
	Name string `json:"name"`

	// The type of name e.g. 'brand-name', 'functional-name', 'common-name'
	NameType common.Coding `json:"nameType"`
}

// InventoryItemResponsibleOrganization represents organization(s) responsible for the product
type InventoryItemResponsibleOrganization struct {
	common.BackboneElement

	// An organization that has an association with the item, e.g. manufacturer, distributor, responsible, etc.
	Organization common.Reference `json:"organization"`

	// The role of the organization e.g. manufacturer, distributor, etc.
	Role common.CodeableConcept `json:"role"`
}

// InventoryItemDescription represents the descriptive characteristics of the inventory item
type InventoryItemDescription struct {
	common.BackboneElement

	// Textual description of the item
	Description *string `json:"description,omitempty"`

	// The language for the item description
	Language *string `json:"language,omitempty"`
}

// InventoryItemAssociation represents association with other items or products
type InventoryItemAssociation struct {
	common.BackboneElement

	// This attribute defined the type of association when establishing associations or relations between items
	AssociationType common.CodeableConcept `json:"associationType"`

	// The quantity of the related product in this product
	Quantity Ratio `json:"quantity"`

	// The related item or product
	RelatedItem common.Reference `json:"relatedItem"`
}

// InventoryItemCharacteristic represents the descriptive or identifying characteristics of the item
type InventoryItemCharacteristic struct {
	common.BackboneElement

	// The type of characteristic that is being defined
	CharacteristicType common.CodeableConcept `json:"characteristicType"`

	// The string value is used for characteristics that are descriptive and not codeable information
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueInteger         *int                    `json:"valueInteger,omitempty"`
	ValueDecimal         *float64                `json:"valueDecimal,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueUrl             *string                 `json:"valueUrl,omitempty"`
	ValueDateTime        *string                 `json:"valueDateTime,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueRatio           *Ratio                  `json:"valueRatio,omitempty"`
	ValueAnnotation      *Annotation             `json:"valueAnnotation,omitempty"`
	ValueAddress         *Address                `json:"valueAddress,omitempty"`
	ValueDuration        *Duration               `json:"valueDuration,omitempty"`
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
}

// InventoryItemInstance represents instances or occurrences of the product
type InventoryItemInstance struct {
	common.BackboneElement

	// The expiry date or date and time for the product
	Expiry *string `json:"expiry,omitempty"`

	// The identifier for the physical instance, typically a serial number
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The location that the item is associated with
	Location *common.Reference `json:"location,omitempty"`

	// The lot or batch number of the item
	LotNumber *string `json:"lotNumber,omitempty"`

	// The subject that the item is associated with
	Subject *common.Reference `json:"subject,omitempty"`
}

// InventoryItem represents a functional description of an inventory item used in inventory and supply-related workflows
type InventoryItem struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "InventoryItem"

	// Association with other items or products
	Association []InventoryItemAssociation `json:"association,omitempty"`

	// The base unit of measure - the unit in which the product is used or counted
	BaseUnit *common.CodeableConcept `json:"baseUnit,omitempty"`

	// Category or class of the item
	Category []common.CodeableConcept `json:"category,omitempty"`

	// The descriptive or identifying characteristics of the item
	Characteristic []InventoryItemCharacteristic `json:"characteristic,omitempty"`

	// Code designating the specific type of item
	Code []common.CodeableConcept `json:"code,omitempty"`

	// The descriptive characteristics of the inventory item
	Description *InventoryItemDescription `json:"description,omitempty"`

	// Business identifier for the inventory item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Instances or occurrences of the product
	Instance *InventoryItemInstance `json:"instance,omitempty"`

	// The usage status e.g. recalled, in use, discarded... This can be used to indicate that the items have been taken out of inventory
	InventoryStatus []common.CodeableConcept `json:"inventoryStatus,omitempty"`

	// The item name(s) - the brand name, or common name, functional name, generic name
	Name []InventoryItemName `json:"name,omitempty"`

	// Net content or amount present in the item
	NetContent *common.Quantity `json:"netContent,omitempty"`

	// Link to a product resource used in clinical workflows
	ProductReference *common.Reference `json:"productReference,omitempty"`

	// Organization(s) responsible for the product
	ResponsibleOrganization []InventoryItemResponsibleOrganization `json:"responsibleOrganization,omitempty"`

	// Status of the item entry
	Status InventoryItemStatus `json:"status"`
}
