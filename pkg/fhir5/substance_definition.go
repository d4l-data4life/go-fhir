// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SubstanceDefinitionMoiety represents moiety, for structural modifications
type SubstanceDefinitionMoiety struct {
	common.BackboneElement

	// Quantitative value for this moiety
	AmountQuantity *common.Quantity `json:"amountQuantity,omitempty"`
	AmountString   *string          `json:"amountString,omitempty"`

	// Identifier by which this moiety substance is known
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// The measurement type of the quantitative value
	MeasurementType *common.CodeableConcept `json:"measurementType,omitempty"`

	// Molecular formula for this moiety of this substance, typically using the Hill system
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

// SubstanceDefinitionCharacterization represents general specifications for this substance
type SubstanceDefinitionCharacterization struct {
	common.BackboneElement

	// The description or justification in support of the interpretation of the data file
	Description *string `json:"description,omitempty"`

	// The data produced by the analytical instrument or a pictorial representation of that data
	File []Attachment `json:"file,omitempty"`

	// Describes the nature of the chemical entity and explains, for instance, whether this is a base or a salt form
	Form *common.CodeableConcept `json:"form,omitempty"`

	// The method used to elucidate the characterization of the drug substance
	Technique *common.CodeableConcept `json:"technique,omitempty"`
}

// SubstanceDefinitionProperty represents general specifications for this substance
type SubstanceDefinitionProperty struct {
	common.BackboneElement

	// A code expressing the type of property
	Type common.CodeableConcept `json:"type"`

	// A value for the property
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueDate            *string                 `json:"valueDate,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
}

// SubstanceDefinitionMolecularWeight represents the average mass of a molecule of a compound
type SubstanceDefinitionMolecularWeight struct {
	common.BackboneElement

	// Used to capture quantitative values for a variety of elements
	Amount common.Quantity `json:"amount"`

	// The method by which the molecular weight was determined
	Method *common.CodeableConcept `json:"method,omitempty"`

	// Type of molecular weight such as exact, average (also known as. number average), weight average
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SubstanceDefinitionStructureRepresentation represents a depiction of the structure of the substance
type SubstanceDefinitionStructureRepresentation struct {
	common.BackboneElement

	// An attached file with the structural representation
	Document *common.Reference `json:"document,omitempty"`

	// The format of the representation e.g. InChI, SMILES, MOLFILE, CDX, SDF, PDB, mmCIF
	Format *common.CodeableConcept `json:"format,omitempty"`

	// The structural representation as a text string in a standard format
	Representation *string `json:"representation,omitempty"`

	// The kind of structural representation (e.g. full, partial)
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SubstanceDefinitionStructure represents structural information
type SubstanceDefinitionStructure struct {
	common.BackboneElement

	// An expression which states the number and type of atoms present in a molecule of a substance
	MolecularFormula *string `json:"molecularFormula,omitempty"`

	// Specified per moiety according to the Hill system, i.e. first C, then H, then alphabetical, each moiety separated by a dot
	MolecularFormulaByMoiety *string `json:"molecularFormulaByMoiety,omitempty"`

	// The molecular weight or weight range (for proteins, polymers or nucleic acids)
	MolecularWeight *SubstanceDefinitionMolecularWeight `json:"molecularWeight,omitempty"`

	// Stereochemistry type
	Stereochemistry *common.CodeableConcept `json:"stereochemistry,omitempty"`

	// Optical activity type
	OpticalActivity *common.CodeableConcept `json:"opticalActivity,omitempty"`

	// Applicable for single substances that contain a radionuclide or a non-natural isotopic ratio
	Isotope []SubstanceDefinitionStructureIsotope `json:"isotope,omitempty"`

	// Molecular structural representation
	Representation []SubstanceDefinitionStructureRepresentation `json:"representation,omitempty"`
}

// SubstanceDefinitionStructureIsotope represents isotope information
type SubstanceDefinitionStructureIsotope struct {
	common.BackboneElement

	// Substance identifier for each non-natural or radioisotope
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Substance name for each non-natural or radioisotope
	Name *common.CodeableConcept `json:"name,omitempty"`

	// The type of isotopic substitution present in a single substance
	Substitution *common.CodeableConcept `json:"substitution,omitempty"`

	// Half life - for a non-natural nuclide, a measure of the instability
	HalfLife *common.Quantity `json:"halfLife,omitempty"`

	// Molecular weight
	MolecularWeight *SubstanceDefinitionMolecularWeight `json:"molecularWeight,omitempty"`
}

// SubstanceDefinitionCode represents codes associated with the substance
type SubstanceDefinitionCode struct {
	common.BackboneElement

	// The specific code
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Status of the code assignment
	Status *common.CodeableConcept `json:"status,omitempty"`

	// The date at which the code status is changed as part of the terminology maintenance
	StatusDate *string `json:"statusDate,omitempty"`

	// Any comment can be provided in this field, if necessary
	Note []Annotation `json:"note,omitempty"`

	// Supporting literature
	Source []common.Reference `json:"source,omitempty"`
}

// SubstanceDefinitionName represents names applicable to this substance
type SubstanceDefinitionName struct {
	common.BackboneElement

	// The actual name
	Name string `json:"name"`

	// Name type
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The status of the name
	Status *common.CodeableConcept `json:"status,omitempty"`

	// If this is the preferred name for this substance
	Preferred *bool `json:"preferred,omitempty"`

	// Human language that the name is written in
	Language []common.CodeableConcept `json:"language,omitempty"`

	// The use context of this name for example if there is a different name a drug active ingredient as opposed to a food colour additive
	Domain []common.CodeableConcept `json:"domain,omitempty"`

	// The jurisdiction where this name applies
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// A synonym of this particular name, by which the substance is also known
	Synonym []SubstanceDefinitionName `json:"synonym,omitempty"`

	// A translation for this name into another human language
	Translation []SubstanceDefinitionName `json:"translation,omitempty"`

	// Details of the official nature of this name
	Official []SubstanceDefinitionNameOfficial `json:"official,omitempty"`

	// Supporting literature
	Source []common.Reference `json:"source,omitempty"`
}

// SubstanceDefinitionNameOfficial represents details of the official nature of this name
type SubstanceDefinitionNameOfficial struct {
	common.BackboneElement

	// Which authority uses this official name
	Authority *common.CodeableConcept `json:"authority,omitempty"`

	// The status of the official name
	Status *common.CodeableConcept `json:"status,omitempty"`

	// Date of official name change
	Date *string `json:"date,omitempty"`
}

// SubstanceDefinitionRelationship represents a link between this substance and another
type SubstanceDefinitionRelationship struct {
	common.BackboneElement

	// A pointer to another substance, as a resource or just a representational code
	SubstanceDefinitionCodeableConcept *common.CodeableConcept `json:"substanceDefinitionCodeableConcept,omitempty"`
	SubstanceDefinitionReference       *common.Reference       `json:"substanceDefinitionReference,omitempty"`

	// For example "salt to parent", "active moiety", "starting material"
	Type common.CodeableConcept `json:"type"`

	// For example where an enzyme strongly bonds onto a particular substance, this is a defining relationship for that enzyme, out of several possible substance relationships
	IsDefining *bool `json:"isDefining,omitempty"`

	// A numeric factor for the relationship, for instance to express that the salt of a substance has some percentage of the active substance in relation to some other
	AmountQuantity *common.Quantity `json:"amountQuantity,omitempty"`
	AmountRatio    *Ratio           `json:"amountRatio,omitempty"`
	AmountString   *string          `json:"amountString,omitempty"`

	// For use when the numeric
	RatioHighLimitAmount *common.Quantity `json:"ratioHighLimitAmount,omitempty"`
	RatioHighLimitRatio  *Ratio           `json:"ratioHighLimitRatio,omitempty"`

	// An operator for the amount, for example "average", "approximately", "less than"
	Comparator *common.CodeableConcept `json:"comparator,omitempty"`

	// Supporting literature
	Source []common.Reference `json:"source,omitempty"`
}

// SubstanceDefinitionSourceMaterial represents material or taxonomic/anatomical source for the substance
type SubstanceDefinitionSourceMaterial struct {
	common.BackboneElement

	// A classification that provides the origin of the raw material
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The genus of an organism shall be specified
	Genus *common.CodeableConcept `json:"genus,omitempty"`

	// The species of an organism shall be specified
	Species *common.CodeableConcept `json:"species,omitempty"`

	// An anatomical origin of the source material within an organism
	Part *common.CodeableConcept `json:"part,omitempty"`

	// The country or countries where the plant material is harvested or the countries where the plasma is sourced from
	CountryOfOrigin []common.CodeableConcept `json:"countryOfOrigin,omitempty"`
}

// SubstanceDefinition represents the detailed description of a substance
type SubstanceDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SubstanceDefinition"

	// General specifications for this substance
	Characterization []SubstanceDefinitionCharacterization `json:"characterization,omitempty"`

	// A high level categorization, e.g. polymer or nucleic acid, or food, chemical, biological
	Classification []common.CodeableConcept `json:"classification,omitempty"`

	// Codes associated with the substance
	Code []SubstanceDefinitionCode `json:"code,omitempty"`

	// Textual description of the substance
	Description *string `json:"description,omitempty"`

	// If the substance applies to human or veterinary use
	Domain *common.CodeableConcept `json:"domain,omitempty"`

	// The quality standard, established benchmark, to which substance complies
	Grade []common.CodeableConcept `json:"grade,omitempty"`

	// Identifier by which this substance is known
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Supporting literature
	InformationSource []common.Reference `json:"informationSource,omitempty"`

	// The entity that creates, makes, produces or fabricates the substance
	Manufacturer []common.Reference `json:"manufacturer,omitempty"`

	// Moiety, for structural modifications
	Moiety []SubstanceDefinitionMoiety `json:"moiety,omitempty"`

	// The average mass of a molecule of a compound compared to 1/12 the mass of carbon 12
	MolecularWeight []SubstanceDefinitionMolecularWeight `json:"molecularWeight,omitempty"`

	// Names applicable to this substance
	Name []SubstanceDefinitionName `json:"name,omitempty"`

	// Textual comment about the substance's catalogue or registry record
	Note []Annotation `json:"note,omitempty"`

	// Data items specific to nucleic acids
	NucleicAcid *common.Reference `json:"nucleicAcid,omitempty"`

	// Data items specific to polymers
	Polymer *common.Reference `json:"polymer,omitempty"`

	// General specifications for this substance
	Property []SubstanceDefinitionProperty `json:"property,omitempty"`

	// Data items specific to proteins
	Protein *common.Reference `json:"protein,omitempty"`

	// General information detailing this substance
	ReferenceInformation *common.Reference `json:"referenceInformation,omitempty"`

	// A link between this substance and another, with details of the relationship
	Relationship []SubstanceDefinitionRelationship `json:"relationship,omitempty"`

	// Material or taxonomic/anatomical source for the substance
	SourceMaterial *SubstanceDefinitionSourceMaterial `json:"sourceMaterial,omitempty"`

	// Status of substance within the catalogue e.g. active, retired
	Status *common.CodeableConcept `json:"status,omitempty"`

	// Structural information
	Structure *SubstanceDefinitionStructure `json:"structure,omitempty"`

	// An entity that is the source for the substance
	Supplier []common.Reference `json:"supplier,omitempty"`

	// A business level version identifier of the substance
	Version *string `json:"version,omitempty"`
}
