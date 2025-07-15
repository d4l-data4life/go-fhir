// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// BiologicallyDerivedProductCollection represents collection details for BiologicallyDerivedProduct
type BiologicallyDerivedProductCollection struct {
	common.BackboneElement

	// Healthcare professional who is performing the collection
	Collector *common.Reference `json:"collector,omitempty"`

	// The patient or entity, such as a hospital or vendor in the case of a processed/manipulated/manufactured product, providing the product
	Source *common.Reference `json:"source,omitempty"`

	// Time of product collection
	CollectedDateTime *string `json:"collectedDateTime,omitempty"`

	// Time of product collection
	CollectedPeriod *common.Period `json:"collectedPeriod,omitempty"`
}

// BiologicallyDerivedProductProperty represents property for BiologicallyDerivedProduct
type BiologicallyDerivedProductProperty struct {
	common.BackboneElement

	// Code that specifies the property being defined
	Type common.CodeableConcept `json:"type"`

	// Property values
	Value interface{} `json:"value,omitempty"`
}

// BiologicallyDerivedProduct represents a material substance originating from a biological entity intended to be transplanted or infused into another (possibly the same) biological entity
type BiologicallyDerivedProduct struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "BiologicallyDerivedProduct"

	// How this product was collected
	Collection *BiologicallyDerivedProductCollection `json:"collection,omitempty"`

	// Any manipulation of product post-collection that is intended to alter the product
	Manipulation interface{} `json:"manipulation,omitempty"`

	// Any processing of the product during collection that does not change the fundamental nature of the product
	Processing []interface{} `json:"processing,omitempty"`

	// Product storage
	Storage []interface{} `json:"storage,omitempty"`

	// Whether this product is currently available
	Availability *string `json:"availability,omitempty"`

	// Product category
	Category []common.CodeableConcept `json:"category,omitempty"`

	// What this biologically derived product is
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Description of the product
	Description *string `json:"description,omitempty"`

	// External ids for this item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The parent biologically derived product
	Parent []common.Reference `json:"parent,omitempty"`

	// Procedure request
	Request []common.Reference `json:"request,omitempty"`

	// available | unavailable | entered-in-error
	Status *string `json:"status,omitempty"`

	// The subject of the product
	Subject *common.Reference `json:"subject,omitempty"`

	// Property can be used to provide information on a wide range of additional information specific to a particular biologicallyDerivedProduct
	Property []BiologicallyDerivedProductProperty `json:"property,omitempty"`
}
