package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

type Medication struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Medication"

	// Business identifier for this medication
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A code that specifies this medication
	Code *common.CodeableConcept `json:"code,omitempty"`

	// active | inactive | entered-in-error
	Status *MedicationStatus `json:"status,omitempty"`

	// Manufacturer of the medication
	Manufacturer *common.Reference `json:"manufacturer,omitempty"`

	// powder | tablets | capsule +
	Form *common.CodeableConcept `json:"form,omitempty"`

	// Amount of drug in package
	Amount *Ratio `json:"amount,omitempty"`

	// Active or inactive ingredient
	Ingredient []MedicationIngredient `json:"ingredient,omitempty"`

	// Details about packaged medications
	Batch *MedicationBatch `json:"batch,omitempty"`
}

type MedicationStatus string

const (
	MedicationStatusActive         MedicationStatus = "active"
	MedicationStatusInactive       MedicationStatus = "inactive"
	MedicationStatusEnteredInError MedicationStatus = "entered-in-error"
)

type MedicationIngredient struct {
	common.BackboneElement

	// The actual ingredient - either a substance (simple ingredient) or another medication
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
	ItemReference       *common.Reference       `json:"itemReference,omitempty"`

	// Indication of whether this ingredient affects the therapeutic action of the drug
	IsActive *bool `json:"isActive,omitempty"`

	// Quantity of ingredient present
	Strength *Ratio `json:"strength,omitempty"`
}

type MedicationBatch struct {
	common.BackboneElement

	// Identifier assigned to batch
	LotNumber *string `json:"lotNumber,omitempty"`

	// When batch will expire
	ExpirationDate *string `json:"expirationDate,omitempty"`
}
