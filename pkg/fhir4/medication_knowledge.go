package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MedicationKnowledge represents information about a medication that is used to support knowledge about this medication
type MedicationKnowledge struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicationKnowledge"

	// Guidelines for the administration of the medication
	AdministrationGuidelines []MedicationKnowledgeAdministrationGuidelines `json:"administrationGuidelines,omitempty"`

	// This is the quantity of medication in a package
	Amount *common.Quantity `json:"amount,omitempty"`

	// Associated or related medications
	AssociatedMedication []common.Reference `json:"associatedMedication,omitempty"`

	// Depending on the context of use, the code that was actually selected by the user
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Potential clinical issue with or between medication(s)
	Contraindication []common.Reference `json:"contraindication,omitempty"`

	// The price of the medication
	Cost []MedicationKnowledgeCost `json:"cost,omitempty"`

	// When Medication is referenced from MedicationRequest, this is the ordered form
	DoseForm *common.CodeableConcept `json:"doseForm,omitempty"`

	// Specifies descriptive properties of the medicine, such as color, shape, imprints, etc
	DrugCharacteristic []MedicationKnowledgeDrugCharacteristic `json:"drugCharacteristic,omitempty"`

	// Identifies a particular constituent of interest in the product
	Ingredient []MedicationKnowledgeIngredient `json:"ingredient,omitempty"`

	// The intended or approved route of administration
	IntendedRoute []common.CodeableConcept `json:"intendedRoute,omitempty"`

	// The time course of drug absorption, distribution, metabolism and excretion of a medication from the body
	Kinetics []MedicationKnowledgeKinetics `json:"kinetics,omitempty"`

	// Describes the details of the manufacturer of the medication product
	Manufacturer *common.Reference `json:"manufacturer,omitempty"`

	// Categorization of the medication within a formulary or classification system
	MedicineClassification []MedicationKnowledgeMedicineClassification `json:"medicineClassification,omitempty"`

	// The program under which the medication is reviewed
	MonitoringProgram []MedicationKnowledgeMonitoringProgram `json:"monitoringProgram,omitempty"`

	// Associated documentation about the medication
	Monograph []MedicationKnowledgeMonograph `json:"monograph,omitempty"`

	// Information that only applies to packages (not products)
	Packaging *MedicationKnowledgePackaging `json:"packaging,omitempty"`

	// The instructions for preparing the medication
	PreparationInstruction *string `json:"preparationInstruction,omitempty"`

	// Category of the medication or product (e.g. branded product, therapeutic moeity, generic product, innovator product, etc.)
	ProductType []common.CodeableConcept `json:"productType,omitempty"`

	// Regulatory information about a medication
	Regulatory []MedicationKnowledgeRegulatory `json:"regulatory,omitempty"`

	// Associated or related knowledge about a medication
	RelatedMedicationKnowledge []MedicationKnowledgeRelatedMedicationKnowledge `json:"relatedMedicationKnowledge,omitempty"`

	// This status is intended to identify if the medication in a local system is in active use within a drug database or inventory
	Status *MedicationKnowledgeStatus `json:"status,omitempty"`

	// Additional names for a medication, for example, the name(s) given to a medication in different countries
	Synonym []string `json:"synonym,omitempty"`
}

// MedicationKnowledgeStatus represents the status of the medication knowledge
type MedicationKnowledgeStatus string

const (
	MedicationKnowledgeStatusActive         MedicationKnowledgeStatus = "active"
	MedicationKnowledgeStatusInactive       MedicationKnowledgeStatus = "inactive"
	MedicationKnowledgeStatusEnteredInError MedicationKnowledgeStatus = "entered-in-error"
)

// MedicationKnowledgeRelatedMedicationKnowledge represents associated or related knowledge about a medication
type MedicationKnowledgeRelatedMedicationKnowledge struct {
	common.BackboneElement

	// Associated documentation about the associated medication knowledge
	Reference []common.Reference `json:"reference"`

	// The category of the associated medication knowledge reference
	Type common.CodeableConcept `json:"type"`
}

// MedicationKnowledgeMonograph represents associated documentation about the medication
type MedicationKnowledgeMonograph struct {
	common.BackboneElement

	// Associated documentation about the medication
	Source *common.Reference `json:"source,omitempty"`

	// The category of documentation about the medication
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// MedicationKnowledgeIngredient represents identifies a particular constituent of interest in the product
type MedicationKnowledgeIngredient struct {
	common.BackboneElement

	// Indication of whether this ingredient affects the therapeutic action of the drug
	IsActive *bool `json:"isActive,omitempty"`

	// The actual ingredient - either a substance (simple ingredient) or another medication
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`

	// The actual ingredient - either a substance (simple ingredient) or another medication
	ItemReference *common.Reference `json:"itemReference,omitempty"`

	// Specifies how many (or how much) of the items there are in this Medication
	Strength *common.Ratio `json:"strength,omitempty"`
}

// MedicationKnowledgeCost represents the price of the medication
type MedicationKnowledgeCost struct {
	common.BackboneElement

	// The price of the medication
	Cost common.Money `json:"cost"`

	// The source or owner that assigns the price to the medication
	Source *string `json:"source,omitempty"`

	// The category of the cost information
	Type common.CodeableConcept `json:"type"`
}

// MedicationKnowledgeMonitoringProgram represents the program under which the medication is reviewed
type MedicationKnowledgeMonitoringProgram struct {
	common.BackboneElement

	// Name of the reviewing program
	Name *string `json:"name,omitempty"`

	// Type of program under which the medication is monitored
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// MedicationKnowledgeAdministrationGuidelines represents guidelines for the administration of the medication
type MedicationKnowledgeAdministrationGuidelines struct {
	common.BackboneElement

	// Dosage for the medication for the specific guidelines
	Dosage []MedicationKnowledgeAdministrationGuidelinesDosage `json:"dosage,omitempty"`

	// Indication for use that apply to the specific administration guidelines
	IndicationCodeableConcept *common.CodeableConcept `json:"indicationCodeableConcept,omitempty"`

	// Indication for use that apply to the specific administration guidelines
	IndicationReference *common.Reference `json:"indicationReference,omitempty"`

	// Characteristics of the patient that are relevant to the administration guidelines
	PatientCharacteristics []MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics `json:"patientCharacteristics,omitempty"`
}

// MedicationKnowledgeAdministrationGuidelinesDosage represents dosage for the medication for the specific guidelines
type MedicationKnowledgeAdministrationGuidelinesDosage struct {
	common.BackboneElement

	// Dosage for the medication for the specific guidelines
	Dosage []common.Dosage `json:"dosage"`

	// The type of dosage (for example, prophylaxis, maintenance, therapeutic, etc.)
	Type common.CodeableConcept `json:"type"`
}

// MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics represents characteristics of the patient that are relevant to the administration guidelines
type MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics struct {
	common.BackboneElement

	// The specific characteristic (e.g. height, weight, gender, etc.)
	CharacteristicCodeableConcept *common.CodeableConcept `json:"characteristicCodeableConcept,omitempty"`

	// The specific characteristic (e.g. height, weight, gender, etc.)
	CharacteristicQuantity *common.Quantity `json:"characteristicQuantity,omitempty"`
}

// MedicationKnowledgeMedicineClassification represents categorization of the medication within a formulary or classification system
type MedicationKnowledgeMedicineClassification struct {
	common.BackboneElement

	// The specific category of the medication within a formulary or classification system
	Classification []common.CodeableConcept `json:"classification,omitempty"`

	// The type of category for the medication (for example, therapeutic classification, therapeutic sub-classification)
	Type common.CodeableConcept `json:"type"`
}

// MedicationKnowledgePackaging represents information that only applies to packages (not products)
type MedicationKnowledgePackaging struct {
	common.BackboneElement

	// The cost of the packaged medication
	Cost []MedicationKnowledgeCost `json:"cost,omitempty"`

	// The number of product units the package would contain if fully loaded
	Quantity *common.Quantity `json:"quantity,omitempty"`
}

// MedicationKnowledgeDrugCharacteristic represents specifies descriptive properties of the medicine
type MedicationKnowledgeDrugCharacteristic struct {
	common.BackboneElement

	// A code expressing the type of characteristic
	Type common.CodeableConcept `json:"type"`

	// Description of the characteristic
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`

	// Description of the characteristic
	ValueQuantity *common.Quantity `json:"valueQuantity,omitempty"`

	// Description of the characteristic
	ValueString *string `json:"valueString,omitempty"`
}

// MedicationKnowledgeRegulatory represents regulatory information about a medication
type MedicationKnowledgeRegulatory struct {
	common.BackboneElement

	// The maximum number of units of the medication that can be dispensed in a period
	MaxDispense *MedicationKnowledgeRegulatoryMaxDispense `json:"maxDispense,omitempty"`

	// The authority that is specifying the regulations
	RegulatoryAuthority common.Reference `json:"regulatoryAuthority"`

	// Specifies the schedule of a medication in jurisdiction
	Schedule []MedicationKnowledgeRegulatorySchedule `json:"schedule,omitempty"`

	// Specifies if changes are allowed when dispensing a medication from a regulatory perspective
	Substitution []MedicationKnowledgeRegulatorySubstitution `json:"substitution,omitempty"`
}

// MedicationKnowledgeRegulatorySubstitution represents specifies if changes are allowed when dispensing a medication from a regulatory perspective
type MedicationKnowledgeRegulatorySubstitution struct {
	common.BackboneElement

	// Specifies the type of substitution allowed
	Allowed bool `json:"allowed"`

	// Specifies the type of substitution allowed
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// MedicationKnowledgeRegulatorySchedule represents specifies the schedule of a medication in jurisdiction
type MedicationKnowledgeRegulatorySchedule struct {
	common.BackboneElement

	// Specifies the specific drug schedule
	Schedule common.CodeableConcept `json:"schedule"`
}

// MedicationKnowledgeRegulatoryMaxDispense represents the maximum number of units of the medication that can be dispensed in a period
type MedicationKnowledgeRegulatoryMaxDispense struct {
	common.BackboneElement

	// The maximum number of units of the medication that can be dispensed
	Quantity common.Quantity `json:"quantity"`

	// The period that applies to the maximum number of units
	Period *common.Duration `json:"period,omitempty"`
}

// MedicationKnowledgeKinetics represents the time course of drug absorption, distribution, metabolism and excretion of a medication from the body
type MedicationKnowledgeKinetics struct {
	common.BackboneElement

	// The drug concentration measured at certain discrete points in time
	AreaUnderCurve []common.Quantity `json:"areaUnderCurve,omitempty"`

	// The time required for concentration in the body to decrease by half
	HalfLifePeriod *common.Duration `json:"halfLifePeriod,omitempty"`

	// The median lethal dose of a drug
	LethalDose50 []common.Quantity `json:"lethalDose50,omitempty"`
}
