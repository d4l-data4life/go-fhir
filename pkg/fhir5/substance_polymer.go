// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SubstancePolymerMonomerSetStartingMaterial represents the starting materials used in the synthesis of the polymer
type SubstancePolymerMonomerSetStartingMaterial struct {
	common.BackboneElement

	// A percentage
	Amount *common.Quantity `json:"amount,omitempty"`

	// Substance high level category, e.g. chemical substance
	Category *common.CodeableConcept `json:"category,omitempty"`

	// The type of substance for this starting material
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Used to specify whether the attribute described is a defining element for the unique identification of the polymer
	IsDefining *bool `json:"isDefining,omitempty"`
}

// SubstancePolymerMonomerSet represents monomer set information
type SubstancePolymerMonomerSet struct {
	common.BackboneElement

	// Captures the type of ratio to the entire polymer, e.g. Monomer/Polymer ratio, SRU/Polymer Ratio
	RatioType *common.CodeableConcept `json:"ratioType,omitempty"`

	// The starting materials - monomer(s) used in the synthesis of the polymer
	StartingMaterial []SubstancePolymerMonomerSetStartingMaterial `json:"startingMaterial,omitempty"`
}

// SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation represents degree of polymerisation information
type SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation struct {
	common.BackboneElement

	// An average amount of polymerisation
	Average *int `json:"average,omitempty"`

	// A high expected limit of the amount
	High *int `json:"high,omitempty"`

	// A low expected limit of the amount
	Low *int `json:"low,omitempty"`

	// The type of the degree of polymerisation shall be described, e.g. SRU/Polymer Ratio
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SubstancePolymerRepeatRepeatUnitStructuralRepresentation represents a graphical structure for this SRU
type SubstancePolymerRepeatRepeatUnitStructuralRepresentation struct {
	common.BackboneElement

	// An attached file with the structural representation
	Attachment *Attachment `json:"attachment,omitempty"`

	// The format of the representation e.g. InChI, SMILES, MOLFILE, CDX, SDF, PDB, mmCIF
	Format *common.CodeableConcept `json:"format,omitempty"`

	// The structural representation as text string in a standard format
	Representation *string `json:"representation,omitempty"`

	// The type of structure (e.g. Full, Partial, Representative)
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SubstancePolymerRepeatRepeatUnit represents an SRU - Structural Repeat Unit
type SubstancePolymerRepeatRepeatUnit struct {
	common.BackboneElement

	// Number of repeats of this unit
	Amount *int `json:"amount,omitempty"`

	// Applies to homopolymer and block co-polymers where the degree of polymerisation within a block can be described
	DegreeOfPolymerisation []SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation `json:"degreeOfPolymerisation,omitempty"`

	// The orientation of the polymerisation, e.g. head-tail, head-head, random
	Orientation *common.CodeableConcept `json:"orientation,omitempty"`

	// A graphical structure for this SRU
	StructuralRepresentation []SubstancePolymerRepeatRepeatUnitStructuralRepresentation `json:"structuralRepresentation,omitempty"`

	// Structural repeat units are essential elements for defining polymers
	Unit *string `json:"unit,omitempty"`
}

// SubstancePolymerRepeat represents specifications and quantifies the repeated units and their configuration
type SubstancePolymerRepeat struct {
	common.BackboneElement

	// A representation of an (average) molecular formula from a polymer
	AverageMolecularFormula *string `json:"averageMolecularFormula,omitempty"`

	// An SRU - Structural Repeat Unit
	RepeatUnit []SubstancePolymerRepeatRepeatUnit `json:"repeatUnit,omitempty"`

	// How the quantitative amount of Structural Repeat Units is captured (e.g. Exact, Numeric, Average)
	RepeatUnitAmountType *common.CodeableConcept `json:"repeatUnitAmountType,omitempty"`
}

// SubstancePolymer represents properties of a substance specific to it being a polymer
type SubstancePolymer struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SubstancePolymer"

	// Overall type of the polymer
	Class *common.CodeableConcept `json:"class,omitempty"`

	// Describes the copolymer sequence type (polymer connectivity)
	CopolymerConnectivity []common.CodeableConcept `json:"copolymerConnectivity,omitempty"`

	// Polymer geometry, e.g. linear, branched, cross-linked, network or dendritic
	Geometry *common.CodeableConcept `json:"geometry,omitempty"`

	// A business identifier for this polymer, but typically this is handled by a SubstanceDefinition identifier
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Modification information
	Modification *string `json:"modification,omitempty"`

	// Monomer set information
	MonomerSet []SubstancePolymerMonomerSet `json:"monomerSet,omitempty"`

	// Specifies and quantifies the repeated units and their configuration
	Repeat []SubstancePolymerRepeat `json:"repeat,omitempty"`
}
