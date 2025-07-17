package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SubstanceSpecification represents the detailed description of a substance, typically at a summary level
type SubstanceSpecification struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SubstanceSpecification"

	// Codes associated with the substance
	Code []SubstanceSpecificationCode `json:"code,omitempty"`

	// Textual comment about this record of a substance
	Comment *string `json:"comment,omitempty"`

	// Textual description of the substance
	Description *string `json:"description,omitempty"`

	// If the substance applies to only human or veterinary use
	Domain *common.CodeableConcept `json:"domain,omitempty"`

	// Identifier by which this substance is known
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Moiety, for structural modifications
	Moiety []SubstanceSpecificationMoiety `json:"moiety,omitempty"`

	// The molecular weight or weight range (for proteins, polymers or nucleic acids)
	MolecularWeight []SubstanceSpecificationStructureIsotopeMolecularWeight `json:"molecularWeight,omitempty"`

	// Names applicable to this substance
	Name []SubstanceSpecificationName `json:"name,omitempty"`

	// Data items specific to nucleic acids
	NucleicAcid *common.Reference `json:"nucleicAcid,omitempty"`

	// Data items specific to polymers
	Polymer *common.Reference `json:"polymer,omitempty"`

	// General specifications for this substance, including how it is related to other substances
	Property []SubstanceSpecificationProperty `json:"property,omitempty"`

	// Data items specific to proteins
	Protein *common.Reference `json:"protein,omitempty"`

	// General information detailing this substance
	ReferenceInformation *common.Reference `json:"referenceInformation,omitempty"`

	// A link between this substance and another, with details of the relationship
	Relationship []SubstanceSpecificationRelationship `json:"relationship,omitempty"`

	// Supporting literature
	Source []common.Reference `json:"source,omitempty"`

	// Material or taxonomic/anatomical source for the substance
	SourceMaterial *common.Reference `json:"sourceMaterial,omitempty"`

	// Status of substance within the catalogue e.g. approved
	Status *common.CodeableConcept `json:"status,omitempty"`

	// Structural information
	Structure *SubstanceSpecificationStructure `json:"structure,omitempty"`

	// High level categorization, e.g. polymer or nucleic acid
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SubstanceSpecificationMoiety represents moiety, for structural modifications
type SubstanceSpecificationMoiety struct {
	common.BackboneElement

	// Quantitative value for this moiety
	AmountQuantity *common.Quantity `json:"amountQuantity,omitempty"`

	// Quantitative value for this moiety
	AmountString *string `json:"amountString,omitempty"`

	// Identifier by which this moiety substance is known
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Molecular formula
	MolecularFormula *string `json:"molecularFormula,omitempty"`

	// Textual name for this moiety substance
	Name *string `json:"name,omitempty"`

	// Optical activity type
	OpticalActivity *common.CodeableConcept `json:"opticalActivity,omitempty"`

	// Role that the moiety is playing
	Role *common.CodeableConcept `json:"role,omitempty"`

	// Stereochemistry type
	Stereochemistry *common.CodeableConcept `json:"stereochemistry,omitempty"`
}

// SubstanceSpecificationProperty represents general specifications for this substance, including how it is related to other substances
type SubstanceSpecificationProperty struct {
	common.BackboneElement

	// Quantitative value for this property
	AmountQuantity *common.Quantity `json:"amountQuantity,omitempty"`

	// Quantitative value for this property
	AmountString *string `json:"amountString,omitempty"`

	// A category for this property, e.g. Physical, Chemical, Enzymatic
	Category *common.CodeableConcept `json:"category,omitempty"`

	// Property type e.g. viscosity, pH, isoelectric point
	Code *common.CodeableConcept `json:"code,omitempty"`

	// A substance upon which a defining property depends (e.g. for solubility: in water, in alcohol)
	DefiningSubstanceReference *common.Reference `json:"definingSubstanceReference,omitempty"`

	// A substance upon which a defining property depends (e.g. for solubility: in water, in alcohol)
	DefiningSubstanceCodeableConcept *common.CodeableConcept `json:"definingSubstanceCodeableConcept,omitempty"`

	// Parameters that were used in the measurement of a property (e.g. for viscosity: measured at 20C with a pH of 7.1)
	Parameters *string `json:"parameters,omitempty"`
}

// SubstanceSpecificationStructureIsotopeMolecularWeight represents the molecular weight or weight range (for proteins, polymers or nucleic acids)
type SubstanceSpecificationStructureIsotopeMolecularWeight struct {
	common.BackboneElement

	// Used to capture quantitative values for a variety of elements
	Amount *common.Quantity `json:"amount,omitempty"`

	// The method by which the molecular weight was determined
	Method *common.CodeableConcept `json:"method,omitempty"`

	// Type of molecular weight such as exact, average (also known as. number average), weight average
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SubstanceSpecificationStructureIsotope represents applicable for single substances that contain a radionuclide or a non-natural isotopic ratio
type SubstanceSpecificationStructureIsotope struct {
	common.BackboneElement

	// Half life - for a non-natural nuclide
	HalfLife *common.Quantity `json:"halfLife,omitempty"`

	// Substance identifier for each non-natural or radioisotope
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// The molecular weight or weight range (for proteins, polymers or nucleic acids)
	MolecularWeight *SubstanceSpecificationStructureIsotopeMolecularWeight `json:"molecularWeight,omitempty"`

	// Substance name for each non-natural or radioisotope
	Name *common.CodeableConcept `json:"name,omitempty"`

	// The type of isotopic substitution present in a single substance
	Substitution *common.CodeableConcept `json:"substitution,omitempty"`
}

// SubstanceSpecificationStructureRepresentation represents molecular structural representation
type SubstanceSpecificationStructureRepresentation struct {
	common.BackboneElement

	// An attached file with the structural representation
	Attachment *common.Attachment `json:"attachment,omitempty"`

	// The structural representation as text string in a format e.g. InChI, SMILES, MOLFILE, CDX
	Representation *string `json:"representation,omitempty"`

	// The type of structure (e.g. Full, Partial, Representative)
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SubstanceSpecificationStructure represents structural information
type SubstanceSpecificationStructure struct {
	common.BackboneElement

	// Applicable for single substances that contain a radionuclide or a non-natural isotopic ratio
	Isotope []SubstanceSpecificationStructureIsotope `json:"isotope,omitempty"`

	// Molecular formula
	MolecularFormula *string `json:"molecularFormula,omitempty"`

	// Specified per moiety according to the Hill system, i.e. first C, then H, then alphabetical, each moiety separated by a dot
	MolecularFormulaByMoiety *string `json:"molecularFormulaByMoiety,omitempty"`

	// The molecular weight or weight range (for proteins, polymers or nucleic acids)
	MolecularWeight *SubstanceSpecificationStructureIsotopeMolecularWeight `json:"molecularWeight,omitempty"`

	// Optical activity type
	OpticalActivity *common.CodeableConcept `json:"opticalActivity,omitempty"`

	// Molecular structural representation
	Representation []SubstanceSpecificationStructureRepresentation `json:"representation,omitempty"`

	// Supporting literature
	Source []common.Reference `json:"source,omitempty"`

	// Stereochemistry type
	Stereochemistry *common.CodeableConcept `json:"stereochemistry,omitempty"`
}

// SubstanceSpecificationCode represents codes associated with the substance
type SubstanceSpecificationCode struct {
	common.BackboneElement

	// The specific code
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Any comment can be provided in this field, if necessary
	Comment *string `json:"comment,omitempty"`

	// Supporting literature
	Source []common.Reference `json:"source,omitempty"`

	// Status of the code assignment
	Status *common.CodeableConcept `json:"status,omitempty"`

	// The date at which the code status is changed as part of the terminology maintenance
	StatusDate *string `json:"statusDate,omitempty"`
}

// SubstanceSpecificationNameOfficial represents details of the official nature of this name
type SubstanceSpecificationNameOfficial struct {
	common.BackboneElement

	// Which authority uses this official name
	Authority *common.CodeableConcept `json:"authority,omitempty"`

	// Date of official name change
	Date *string `json:"date,omitempty"`

	// The status of the official name
	Status *common.CodeableConcept `json:"status,omitempty"`
}

// SubstanceSpecificationName represents names applicable to this substance
type SubstanceSpecificationName struct {
	common.BackboneElement

	// The use context of this name for example if there is a different name a drug active ingredient as opposed to a food colour additive
	Domain []common.CodeableConcept `json:"domain,omitempty"`

	// The jurisdiction where this name applies
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Language of the name
	Language []common.CodeableConcept `json:"language,omitempty"`

	// The actual name
	Name string `json:"name"`

	// Details of the official nature of this name
	Official []SubstanceSpecificationNameOfficial `json:"official,omitempty"`

	// If this is the preferred name for this substance
	Preferred *bool `json:"preferred,omitempty"`

	// Supporting literature
	Source []common.Reference `json:"source,omitempty"`

	// The status of the name
	Status *common.CodeableConcept `json:"status,omitempty"`

	// Name type, e.g. 'systematic',  'scientific, 'brand'
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SubstanceSpecificationRelationship represents a link between this substance and another, with details of the relationship
type SubstanceSpecificationRelationship struct {
	common.BackboneElement

	// A pointer to another substance, as a resource or just a representational code
	SubstanceReference *common.Reference `json:"substanceReference,omitempty"`

	// A pointer to another substance, as a resource or just a representational code
	SubstanceCodeableConcept *common.CodeableConcept `json:"substanceCodeableConcept,omitempty"`

	// For example "salt to parent", "active moiety", "starting material"
	Relationship *common.CodeableConcept `json:"relationship,omitempty"`

	// For example where an enzyme strongly bonds to a particular substance, this is a defining relationship for that enzyme, out of several possible substance relationships
	IsDefining *bool `json:"isDefining,omitempty"`

	// A numeric factor for the relationship, for instance to express that the salt of a substance has some percentage of the active moiety
	AmountQuantity *common.Quantity `json:"amountQuantity,omitempty"`

	// A numeric factor for the relationship, for instance to express that the salt of a substance has some percentage of the active moiety
	AmountRatio *common.Ratio `json:"amountRatio,omitempty"`

	// A numeric factor for the relationship, for instance to express that the salt of a substance has some percentage of the active moiety
	AmountString *string `json:"amountString,omitempty"`

	// For use when the numeric
	AmountRatioLowLimit *common.Ratio `json:"amountRatioLowLimit,omitempty"`

	// An operator for the amount, for example "average", "approximately", "less than"
	AmountType *common.CodeableConcept `json:"amountType,omitempty"`

	// Supporting literature
	Source []common.Reference `json:"source,omitempty"`
}
