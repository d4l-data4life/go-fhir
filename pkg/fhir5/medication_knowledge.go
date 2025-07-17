package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

type MedicationKnowledgeStatus string

const (
	MedicationKnowledgeStatusActive         MedicationKnowledgeStatus = "active"
	MedicationKnowledgeStatusEnteredInError MedicationKnowledgeStatus = "entered-in-error"
	MedicationKnowledgeStatusInactive       MedicationKnowledgeStatus = "inactive"
)

// MedicationKnowledgeDefinitionalDrugCharacteristic represents descriptive properties of the medicine
type MedicationKnowledgeDefinitionalDrugCharacteristic struct {
	common.BackboneElement

	Type                 *common.CodeableConcept `json:"type,omitempty"`
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueBase64Binary    *string                 `json:"valueBase64Binary,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
}

// MedicationKnowledgeDefinitionalIngredient represents ingredients of the medication
type MedicationKnowledgeDefinitionalIngredient struct {
	common.BackboneElement

	Item                    CodeableReference       `json:"item"`
	IsActive                *bool                   `json:"isActive,omitempty"`
	StrengthRatio           *Ratio                  `json:"strengthRatio,omitempty"`
	StrengthCodeableConcept *common.CodeableConcept `json:"strengthCodeableConcept,omitempty"`
	StrengthQuantity        *common.Quantity        `json:"strengthQuantity,omitempty"`
}

// MedicationKnowledgeDefinitional represents definitional information about the medication
type MedicationKnowledgeDefinitional struct {
	common.BackboneElement

	Definition         []common.Reference                                  `json:"definition,omitempty"`
	DoseForm           *common.CodeableConcept                             `json:"doseForm,omitempty"`
	DrugCharacteristic []MedicationKnowledgeDefinitionalDrugCharacteristic `json:"drugCharacteristic,omitempty"`
	Ingredient         []MedicationKnowledgeDefinitionalIngredient         `json:"ingredient,omitempty"`
	IntendedRoute      []common.CodeableConcept                            `json:"intendedRoute,omitempty"`
}

// MedicationKnowledgeCost represents the price of the medication
type MedicationKnowledgeCost struct {
	common.BackboneElement

	CostMoney           *Money                  `json:"costMoney,omitempty"`
	CostCodeableConcept *common.CodeableConcept `json:"costCodeableConcept,omitempty"`
	EffectiveDate       []common.Period         `json:"effectiveDate,omitempty"`
	Source              *common.CodeableConcept `json:"source,omitempty"`
	Type                *common.CodeableConcept `json:"type,omitempty"`
}

// MedicationKnowledge represents information about a medication that is used to support knowledge
type MedicationKnowledge struct {
	DomainResource

	ResourceType               string                           `json:"resourceType"` // Always "MedicationKnowledge"
	AssociatedMedication       []common.Reference               `json:"associatedMedication,omitempty"`
	Author                     *common.Reference                `json:"author,omitempty"`
	ClinicalUseIssue           []common.Reference               `json:"clinicalUseIssue,omitempty"`
	Code                       *common.CodeableConcept          `json:"code,omitempty"`
	Cost                       []MedicationKnowledgeCost        `json:"cost,omitempty"`
	Definitional               *MedicationKnowledgeDefinitional `json:"definitional,omitempty"`
	Identifier                 []common.Identifier              `json:"identifier,omitempty"`
	IndicationGuideline        []interface{}                    `json:"indicationGuideline,omitempty"`
	MaxDispenseQuantity        *common.Quantity                 `json:"maxDispenseQuantity,omitempty"`
	MedicineClassification     []interface{}                    `json:"medicineClassification,omitempty"`
	MonitoringProgram          []interface{}                    `json:"monitoringProgram,omitempty"`
	Monograph                  []interface{}                    `json:"monograph,omitempty"`
	Packaging                  []interface{}                    `json:"packaging,omitempty"`
	PharmacokineticProperty    []interface{}                    `json:"pharmacokineticProperty,omitempty"`
	RegulatoryInformation      []interface{}                    `json:"regulatoryInformation,omitempty"`
	RelatedMedicationKnowledge []interface{}                    `json:"relatedMedicationKnowledge,omitempty"`
	Status                     *MedicationKnowledgeStatus       `json:"status,omitempty"`
	StorageGuideline           []interface{}                    `json:"storageGuideline,omitempty"`
}
