// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// BiologicallyDerivedProductCollection represents how this product was collected
type BiologicallyDerivedProductCollection struct {
	common.BackboneElement

	// Time of product collection
	CollectedDateTime *string `json:"collectedDateTime,omitempty"`

	// Time of product collection
	CollectedPeriod *common.Period `json:"collectedPeriod,omitempty"`

	// Healthcare professional who is performing the collection
	Collector *common.Reference `json:"collector,omitempty"`

	// The patient or entity providing the product
	Source *common.Reference `json:"source,omitempty"`
}

// BiologicallyDerivedProductProperty represents property information for a biologically derived product
type BiologicallyDerivedProductProperty struct {
	common.BackboneElement

	// The element is identified by name and system URI in the type
	Type common.CodeableConcept `json:"type"`

	// The value should be provided as a boolean, integer, CodeableConcept, period, quantity, range, ratio, or attachment
	ValueBoolean *bool `json:"valueBoolean,omitempty"`

	// The value should be provided as a boolean, integer, CodeableConcept, period, quantity, range, ratio, or attachment
	ValueInteger *int `json:"valueInteger,omitempty"`

	// The value should be provided as a boolean, integer, CodeableConcept, period, quantity, range, ratio, or attachment
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`

	// The value should be provided as a boolean, integer, CodeableConcept, period, quantity, range, ratio, or attachment
	ValuePeriod *common.Period `json:"valuePeriod,omitempty"`

	// The value should be provided as a boolean, integer, CodeableConcept, period, quantity, range, ratio, or attachment
	ValueQuantity *common.Quantity `json:"valueQuantity,omitempty"`

	// The value should be provided as a boolean, integer, CodeableConcept, period, quantity, range, ratio, or attachment
	ValueRange *Range `json:"valueRange,omitempty"`

	// The value should be provided as a boolean, integer, CodeableConcept, period, quantity, range, ratio, or attachment
	ValueRatio *Ratio `json:"valueRatio,omitempty"`

	// The value should be provided as a boolean, integer, CodeableConcept, period, quantity, range, ratio, or attachment
	ValueString *string `json:"valueString,omitempty"`

	// The value should be provided as a boolean, integer, CodeableConcept, period, quantity, range, ratio, or attachment
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
}

// BiologicallyDerivedProduct represents a biologically derived product
type BiologicallyDerivedProduct struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "BiologicallyDerivedProduct"

	// Necessary to support mandatory requirements for traceability from donor/source to recipient
	BiologicalSourceEvent *common.Identifier `json:"biologicalSourceEvent,omitempty"`

	// How this product was collected
	Collection *BiologicallyDerivedProductCollection `json:"collection,omitempty"`

	// A unique identifier for an aliquot of a product
	Division *string `json:"division,omitempty"`

	// Date, and where relevant time, of expiration
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// This identifier should uniquely identify the product instance in the business domain
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// For products that have multiple collections
	Parent []common.Reference `json:"parent,omitempty"`

	// Processing facilities responsible for the labeling and distribution
	ProcessingFacility []common.Reference `json:"processingFacility,omitempty"`

	// Broad category of this product
	ProductCategory *common.Coding `json:"productCategory,omitempty"`

	// A codified value that systematically supports characterization and classification
	ProductCode *common.CodeableConcept `json:"productCode,omitempty"`

	// Whether the product is currently available
	ProductStatus *common.Coding `json:"productStatus,omitempty"`

	// Property can be used to provide information on a wide range of additional information
	Property []BiologicallyDerivedProductProperty `json:"property,omitempty"`

	// Request to obtain and/or infuse this biologically derived product
	Request []common.Reference `json:"request,omitempty"`

	// May be extracted from information held in the Product Description Code
	StorageTempRequirements *Range `json:"storageTempRequirements,omitempty"`
}
