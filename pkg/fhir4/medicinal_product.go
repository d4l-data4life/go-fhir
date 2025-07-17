package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MedicinalProduct represents detailed definition of a medicinal product, typically for uses other than direct patient care
type MedicinalProduct struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicinalProduct"

	// Whether the Medicinal Product is subject to additional monitoring for regulatory reasons
	AdditionalMonitoringIndicator *common.CodeableConcept `json:"additionalMonitoringIndicator,omitempty"`

	// Supporting documentation, typically for regulatory submission
	AttachedDocument []common.Reference `json:"attachedDocument,omitempty"`

	// Clinical trials or studies that this product is involved in
	ClinicalTrial []common.Reference `json:"clinicalTrial,omitempty"`

	// The dose form for a single part product, or combined form of a multiple part product
	CombinedPharmaceuticalDoseForm *common.CodeableConcept `json:"combinedPharmaceuticalDoseForm,omitempty"`

	// A product specific contact, person (in a role), or an organization
	Contact []common.Reference `json:"contact,omitempty"`

	// Reference to another product, e.g. for linking authorised to investigational product
	CrossReference []common.Identifier `json:"crossReference,omitempty"`

	// If this medicine applies to human or veterinary uses
	Domain *common.Coding `json:"domain,omitempty"`

	// Business identifier for this product. Could be an MPID
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The legal status of supply of the medicinal product as classified by the regulator
	LegalStatusOfSupply *common.CodeableConcept `json:"legalStatusOfSupply,omitempty"`

	// An operation applied to the product, for manufacturing or administrative purpose
	ManufacturingBusinessOperation []MedicinalProductManufacturingBusinessOperation `json:"manufacturingBusinessOperation,omitempty"`

	// Marketing status of the medicinal product, in contrast to marketing authorization
	MarketingStatus []MedicinalProductMarketingStatus `json:"marketingStatus,omitempty"`

	// A master file for to the medicinal product (e.g. Pharmacovigilance System Master File)
	MasterFile []common.Reference `json:"masterFile,omitempty"`

	// The product's name, including full name and possibly coded parts
	Name []MedicinalProductName `json:"name"`

	// Package representation for the product
	PackagedMedicinalProduct []common.Reference `json:"packagedMedicinalProduct,omitempty"`

	// If authorised for use in children
	PaediatricUseIndicator *common.CodeableConcept `json:"paediatricUseIndicator,omitempty"`

	// Pharmaceutical aspects of product
	PharmaceuticalProduct []common.Reference `json:"pharmaceuticalProduct,omitempty"`

	// Allows the product to be classified by various systems
	ProductClassification []common.CodeableConcept `json:"productClassification,omitempty"`

	// Indicates if the medicinal product has an orphan designation for the treatment of a rare disease
	SpecialDesignation []MedicinalProductSpecialDesignation `json:"specialDesignation,omitempty"`

	// Whether the Medicinal Product is subject to special measures for regulatory reasons
	SpecialMeasures []string `json:"specialMeasures,omitempty"`

	// Regulatory type, e.g. Investigational or Authorized
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// MedicinalProductName represents the product's name, including full name and possibly coded parts
type MedicinalProductName struct {
	common.BackboneElement

	// Country where the name applies
	CountryLanguage []MedicinalProductNameCountryLanguage `json:"countryLanguage,omitempty"`

	// Coding words or phrases of the name
	NamePart []MedicinalProductNameNamePart `json:"namePart,omitempty"`

	// The full product name
	ProductName string `json:"productName"`
}

// MedicinalProductNameNamePart represents coding words or phrases of the name
type MedicinalProductNameNamePart struct {
	common.BackboneElement

	// A fragment of a product name
	Part string `json:"part"`

	// Identifying type for this part of the name (e.g. strength part)
	Type common.Coding `json:"type"`
}

// MedicinalProductNameCountryLanguage represents country where the name applies
type MedicinalProductNameCountryLanguage struct {
	common.BackboneElement

	// Country code for where this name applies
	Country common.CodeableConcept `json:"country"`

	// Jurisdiction code for where this name applies
	Jurisdiction *common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Language code for this name
	Language common.CodeableConcept `json:"language"`
}

// MedicinalProductManufacturingBusinessOperation represents an operation applied to the product, for manufacturing or administrative purpose
type MedicinalProductManufacturingBusinessOperation struct {
	common.BackboneElement

	// Regulatory authorization reference number
	AuthorisationReferenceNumber *common.Identifier `json:"authorisationReferenceNumber,omitempty"`

	// To indicate if this process is commercially confidential
	ConfidentialityIndicator *common.CodeableConcept `json:"confidentialityIndicator,omitempty"`

	// Regulatory authorization date
	EffectiveDate *string `json:"effectiveDate,omitempty"`

	// The manufacturer or establishment associated with the process
	Manufacturer []common.Reference `json:"manufacturer,omitempty"`

	// The type of manufacturing operation
	OperationType *common.CodeableConcept `json:"operationType,omitempty"`

	// A regulator which oversees the operation
	Regulator *common.Reference `json:"regulator,omitempty"`
}

// MedicinalProductSpecialDesignation represents indicates if the medicinal product has an orphan designation for the treatment of a rare disease
type MedicinalProductSpecialDesignation struct {
	common.BackboneElement

	// Date when the designation was granted
	Date *string `json:"date,omitempty"`

	// Identifier for the designation, or procedure number
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Condition for which the medicinal use applies
	IndicationCodeableConcept *common.CodeableConcept `json:"indicationCodeableConcept,omitempty"`

	// Condition for which the medicinal use applies
	IndicationReference *common.Reference `json:"indicationReference,omitempty"`

	// The intended use of the product, e.g. prevention, treatment
	IntendedUse *common.CodeableConcept `json:"intendedUse,omitempty"`

	// Animal species for which this applies
	Species *common.CodeableConcept `json:"species,omitempty"`

	// For example granted, pending, expired or withdrawn
	Status *common.CodeableConcept `json:"status,omitempty"`

	// The type of special designation, e.g. orphan drug, minor use
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// MedicinalProductMarketingStatus represents the marketing status describes the date when a medicinal product is actually put on the market
type MedicinalProductMarketingStatus struct {
	common.BackboneElement

	// The country in which the marketing authorisation has been granted
	Country common.CodeableConcept `json:"country"`

	// The date when the Medicinal Product is placed on the market
	DateRange common.Period `json:"dateRange"`

	// Where a Medicines Regulatory Agency has granted a marketing authorisation
	Jurisdiction *common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The date when the Medicinal Product is placed on the market
	RestoreDate *string `json:"restoreDate,omitempty"`

	// This attribute provides information on the status of the marketing of the medicinal product
	Status common.CodeableConcept `json:"status"`
}
