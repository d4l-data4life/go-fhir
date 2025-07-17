package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// BiologicallyDerivedProductStatus represents the status of the biologically derived product
type BiologicallyDerivedProductStatus string

const (
	BiologicallyDerivedProductStatusAvailable   BiologicallyDerivedProductStatus = "available"
	BiologicallyDerivedProductStatusUnavailable BiologicallyDerivedProductStatus = "unavailable"
)

// BiologicallyDerivedProductCategory represents the category of the biologically derived product
type BiologicallyDerivedProductCategory string

const (
	BiologicallyDerivedProductCategoryOrgan           BiologicallyDerivedProductCategory = "organ"
	BiologicallyDerivedProductCategoryTissue          BiologicallyDerivedProductCategory = "tissue"
	BiologicallyDerivedProductCategoryFluid           BiologicallyDerivedProductCategory = "fluid"
	BiologicallyDerivedProductCategoryCells           BiologicallyDerivedProductCategory = "cells"
	BiologicallyDerivedProductCategoryBiologicalAgent BiologicallyDerivedProductCategory = "biologicalAgent"
)

// BiologicallyDerivedProductCollection represents collection information
type BiologicallyDerivedProductCollection struct {
	common.BackboneElement

	// Healthcare professional who is performing the collection
	Collector *common.Reference `json:"collector,omitempty"`

	// The patient or entity, such as a hospital or vendor in the case of a processed/manipulated/manufactured product, providing the product
	Source *common.Reference `json:"source,omitempty"`

	// Time of product collection
	CollectedDateTime *string        `json:"collectedDateTime,omitempty"`
	CollectedPeriod   *common.Period `json:"collectedPeriod,omitempty"`
}

// BiologicallyDerivedProductProcessing represents processing information
type BiologicallyDerivedProductProcessing struct {
	common.BackboneElement

	// Description of of processing
	Description *string `json:"description,omitempty"`

	// Procesing code
	Procedure *common.CodeableConcept `json:"procedure,omitempty"`

	// Substance added during processing
	Additive *common.Reference `json:"additive,omitempty"`

	// Time of processing
	TimeDateTime *string        `json:"timeDateTime,omitempty"`
	TimePeriod   *common.Period `json:"timePeriod,omitempty"`
}

// BiologicallyDerivedProductManipulation represents manipulation information
type BiologicallyDerivedProductManipulation struct {
	common.BackboneElement

	// Description of manipulation
	Description *string `json:"description,omitempty"`

	// Time of manipulation
	TimeDateTime *string        `json:"timeDateTime,omitempty"`
	TimePeriod   *common.Period `json:"timePeriod,omitempty"`
}

// BiologicallyDerivedProductStorage represents storage information
type BiologicallyDerivedProductStorage struct {
	common.BackboneElement

	// Description of storage
	Description *string `json:"description,omitempty"`

	// Storage temperature
	Temperature *float64 `json:"temperature,omitempty"`

	// Temperature scale used
	Scale *string `json:"scale,omitempty"`

	// Storage timeperiod
	Duration *common.Period `json:"duration,omitempty"`
}

// BiologicallyDerivedProduct represents a material substance originating from a biological entity intended to be transplanted or infused into another (possibly the same) biological entity
type BiologicallyDerivedProduct struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "BiologicallyDerivedProduct"

	// This records identifiers associated with this biologically derived product that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Broad category of this product
	ProductCategory *BiologicallyDerivedProductCategory `json:"productCategory,omitempty"`

	// A code that identifies the kind of this biologically derived product (SNOMED Ctcode)
	ProductCode *common.CodeableConcept `json:"productCode,omitempty"`

	// Whether the product is currently available
	Status *BiologicallyDerivedProductStatus `json:"status,omitempty"`

	// Procedure request to obtain this biologically derived product
	Request []common.Reference `json:"request,omitempty"`

	// Number of discrete units within this product
	Quantity *int `json:"quantity,omitempty"`

	// Parent product
	Parent []common.Reference `json:"parent,omitempty"`

	// How this product was collected
	Collection *BiologicallyDerivedProductCollection `json:"collection,omitempty"`

	// Any processing of the product during collection
	Processing []BiologicallyDerivedProductProcessing `json:"processing,omitempty"`

	// Any manipulation of product post-collection
	Manipulation *BiologicallyDerivedProductManipulation `json:"manipulation,omitempty"`

	// Product storage
	Storage []BiologicallyDerivedProductStorage `json:"storage,omitempty"`
}
