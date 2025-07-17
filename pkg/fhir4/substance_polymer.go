package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SubstancePolymer represents todo
type SubstancePolymer struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SubstancePolymer"

	// Todo
	Class *common.CodeableConcept `json:"class,omitempty"`

	// Todo
	CopolymerConnectivity []common.CodeableConcept `json:"copolymerConnectivity,omitempty"`

	// Todo
	Geometry *common.CodeableConcept `json:"geometry,omitempty"`

	// Todo
	Modification []string `json:"modification,omitempty"`

	// Todo
	MonomerSet []SubstancePolymerMonomerSet `json:"monomerSet,omitempty"`

	// Todo
	Repeat []SubstancePolymerRepeat `json:"repeat,omitempty"`
}

// SubstancePolymerMonomerSet represents todo
type SubstancePolymerMonomerSet struct {
	common.BackboneElement

	// Todo
	RatioType *common.CodeableConcept `json:"ratioType,omitempty"`

	// Todo
	StartingMaterial []SubstancePolymerMonomerSetStartingMaterial `json:"startingMaterial,omitempty"`
}

// SubstancePolymerMonomerSetStartingMaterial represents todo
type SubstancePolymerMonomerSetStartingMaterial struct {
	common.BackboneElement

	// Todo
	Amount *SubstanceAmount `json:"amount,omitempty"`

	// Todo
	IsDefining *bool `json:"isDefining,omitempty"`

	// Todo
	Material *common.CodeableConcept `json:"material,omitempty"`

	// Todo
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SubstancePolymerRepeat represents todo
type SubstancePolymerRepeat struct {
	common.BackboneElement

	// Todo
	AverageMolecularFormula *string `json:"averageMolecularFormula,omitempty"`

	// Todo
	NumberOfUnits *int `json:"numberOfUnits,omitempty"`

	// Todo
	RepeatUnit []SubstancePolymerRepeatRepeatUnit `json:"repeatUnit,omitempty"`

	// Todo
	RepeatUnitAmountType *common.CodeableConcept `json:"repeatUnitAmountType,omitempty"`
}

// SubstancePolymerRepeatRepeatUnit represents todo
type SubstancePolymerRepeatRepeatUnit struct {
	common.BackboneElement

	// Todo
	Amount *SubstanceAmount `json:"amount,omitempty"`

	// Todo
	DegreeOfPolymerisation []SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation `json:"degreeOfPolymerisation,omitempty"`

	// Todo
	OrientationOfPolymerisation *common.CodeableConcept `json:"orientationOfPolymerisation,omitempty"`

	// Todo
	RepeatUnit *string `json:"repeatUnit,omitempty"`

	// Todo
	StructuralRepresentation []SubstancePolymerRepeatRepeatUnitStructuralRepresentation `json:"structuralRepresentation,omitempty"`
}

// SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation represents todo
type SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation struct {
	common.BackboneElement

	// Todo
	Amount *SubstanceAmount `json:"amount,omitempty"`

	// Todo
	Degree *common.CodeableConcept `json:"degree,omitempty"`
}

// SubstancePolymerRepeatRepeatUnitStructuralRepresentation represents todo
type SubstancePolymerRepeatRepeatUnitStructuralRepresentation struct {
	common.BackboneElement

	// Todo
	Attachment *common.Attachment `json:"attachment,omitempty"`

	// Todo
	Representation *string `json:"representation,omitempty"`

	// Todo
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SubstanceAmount represents todo
type SubstanceAmount struct {
	common.BackboneElement

	// Todo
	AmountQuantity *common.Quantity `json:"amountQuantity,omitempty"`

	// Todo
	AmountString *string `json:"amountString,omitempty"`

	// Todo
	AmountType *common.CodeableConcept `json:"amountType,omitempty"`

	// Todo
	AmountText *string `json:"amountText,omitempty"`
}
