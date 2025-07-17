package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MedicinalProductDefinitionContact represents a product specific contact
type MedicinalProductDefinitionContact struct {
	common.BackboneElement

	// A product specific contact, person (in a role), or an organization
	Contact common.Reference `json:"contact"`

	// Allows the contact to be classified
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// MedicinalProductDefinitionNamePart represents coding words or phrases of the name
type MedicinalProductDefinitionNamePart struct {
	common.BackboneElement

	// A fragment of a product name
	Part string `json:"part"`

	// Identifying type for this part of the name
	Type common.CodeableConcept `json:"type"`
}

// MedicinalProductDefinitionNameUsage represents country and jurisdiction where the name applies
type MedicinalProductDefinitionNameUsage struct {
	common.BackboneElement

	// Country code for where this name applies
	Country common.CodeableConcept `json:"country"`

	// Jurisdiction code for where this name applies
	Jurisdiction *common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Language code for this name
	Language common.CodeableConcept `json:"language"`
}

// MedicinalProductDefinitionName represents the product's name
type MedicinalProductDefinitionName struct {
	common.BackboneElement

	// Coding words or phrases of the name
	Part []MedicinalProductDefinitionNamePart `json:"part,omitempty"`

	// The full product name
	ProductName string `json:"productName"`

	// Type of product name, such as rINN, BAN, Proprietary, Non-Proprietary
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Country and jurisdiction where the name applies, and associated language
	Usage []MedicinalProductDefinitionNameUsage `json:"usage,omitempty"`
}

// MedicinalProductDefinitionCrossReference represents reference to another product
type MedicinalProductDefinitionCrossReference struct {
	common.BackboneElement

	// Reference to another product
	Product CodeableReference `json:"product"`

	// The type of relationship
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// MedicinalProductDefinitionOperation represents a manufacturing or administrative process
type MedicinalProductDefinitionOperation struct {
	common.BackboneElement

	// Specifies whether this particular business or manufacturing process is considered proprietary or confidential
	ConfidentialityIndicator *common.CodeableConcept `json:"confidentialityIndicator,omitempty"`

	// Date range of applicability
	EffectiveDate *common.Period `json:"effectiveDate,omitempty"`

	// The organization or establishment responsible for the particular process or step
	Organization []common.Reference `json:"organization,omitempty"`

	// The type of manufacturing operation
	Type *CodeableReference `json:"type,omitempty"`
}

// MedicinalProductDefinitionCharacteristic represents key product features
type MedicinalProductDefinitionCharacteristic struct {
	common.BackboneElement

	// A code expressing the type of characteristic
	Type common.CodeableConcept `json:"type"`

	// Description of the characteristic
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueMarkdown        *string                 `json:"valueMarkdown,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueInteger         *int                    `json:"valueInteger,omitempty"`
	ValueDate            *string                 `json:"valueDate,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
}

// MedicinalProductDefinition represents a medicinal product
type MedicinalProductDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicinalProductDefinition"

	// Whether the Medicinal Product is subject to additional monitoring for regulatory reasons
	AdditionalMonitoringIndicator *common.CodeableConcept `json:"additionalMonitoringIndicator,omitempty"`

	// Additional information or supporting documentation about the medicinal product
	AttachedDocument []common.Reference `json:"attachedDocument,omitempty"`

	// Allows the key product features to be recorded
	Characteristic []MedicinalProductDefinitionCharacteristic `json:"characteristic,omitempty"`

	// Allows the product to be classified by various systems, commonly WHO ATC
	Classification []common.CodeableConcept `json:"classification,omitempty"`

	// Clinical trials or studies that this product is involved in
	ClinicalTrial []common.Reference `json:"clinicalTrial,omitempty"`

	// A code that this product is known by
	Code []common.Coding `json:"code,omitempty"`

	// The dose form for a single part product, or combined form of a multiple part product
	CombinedPharmaceuticalDoseForm *common.CodeableConcept `json:"combinedPharmaceuticalDoseForm,omitempty"`

	// Types of medicinal manufactured items and/or devices that this product consists of
	ComprisedOf []common.Reference `json:"comprisedOf,omitempty"`

	// A product specific contact, person (in a role), or an organization
	Contact []MedicinalProductDefinitionContact `json:"contact,omitempty"`

	// Reference to another product
	CrossReference []MedicinalProductDefinitionCrossReference `json:"crossReference,omitempty"`

	// General description of this product
	Description *string `json:"description,omitempty"`

	// If this medicine applies to human or veterinary uses
	Domain *common.CodeableConcept `json:"domain,omitempty"`

	// Business identifier for this product
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Any component of the drug product which is not the chemical entity defined as the drug substance
	Impurity []CodeableReference `json:"impurity,omitempty"`

	// Description of indication(s) for this product
	Indication *string `json:"indication,omitempty"`

	// The ingredients of this medicinal product
	Ingredient []common.CodeableConcept `json:"ingredient,omitempty"`

	// The legal status of supply of the medicinal product as classified by the regulator
	LegalStatusOfSupply *common.CodeableConcept `json:"legalStatusOfSupply,omitempty"`

	// Marketing status of the medicinal product
	MarketingStatus []interface{} `json:"marketingStatus,omitempty"`

	// A master file for the medicinal product
	MasterFile []common.Reference `json:"masterFile,omitempty"`

	// The product's name, including full name and possibly coded parts
	Name []MedicinalProductDefinitionName `json:"name"`

	// A manufacturing or administrative process or step associated with the medicinal product
	Operation []MedicinalProductDefinitionOperation `json:"operation,omitempty"`

	// Package type for the product
	PackagedMedicinalProduct []common.CodeableConcept `json:"packagedMedicinalProduct,omitempty"`

	// If authorised for use in children, or infants, neonates etc.
	PediatricUseIndicator *common.CodeableConcept `json:"pediatricUseIndicator,omitempty"`

	// The path by which the product is taken into or makes contact with the body
	Route []common.CodeableConcept `json:"route,omitempty"`

	// Whether the Medicinal Product is subject to special measures for regulatory reasons
	SpecialMeasures []common.CodeableConcept `json:"specialMeasures,omitempty"`

	// The status within the lifecycle of this product record
	Status *common.CodeableConcept `json:"status,omitempty"`

	// The date at which the given status became applicable
	StatusDate *string `json:"statusDate,omitempty"`

	// Regulatory type, e.g. Investigational or Authorized
	Type *common.CodeableConcept `json:"type,omitempty"`

	// A business identifier relating to a specific version of the product
	Version *string `json:"version,omitempty"`
}
