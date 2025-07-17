// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SubstanceSourceMaterialFractionDescription represents fraction description information
type SubstanceSourceMaterialFractionDescription struct {
	common.BackboneElement

	// This element is capturing information about the fraction of a plant part, or human plasma for fractionation
	Fraction *string `json:"fraction,omitempty"`

	// The specific type of the material constituting the component
	MaterialType *common.CodeableConcept `json:"materialType,omitempty"`
}

// SubstanceSourceMaterialOrganismAuthor represents author type information
type SubstanceSourceMaterialOrganismAuthor struct {
	common.BackboneElement

	// The author of an organism species shall be specified
	AuthorDescription *string `json:"authorDescription,omitempty"`

	// The type of author of an organism species shall be specified
	AuthorType *common.CodeableConcept `json:"authorType,omitempty"`
}

// SubstanceSourceMaterialOrganismHybrid represents hybrid species information
type SubstanceSourceMaterialOrganismHybrid struct {
	common.BackboneElement

	// The hybrid type of an organism shall be specified
	HybridType *common.CodeableConcept `json:"hybridType,omitempty"`

	// The identifier of the maternal species constituting the hybrid organism
	MaternalOrganismId *string `json:"maternalOrganismId,omitempty"`

	// The name of the maternal species constituting the hybrid organism
	MaternalOrganismName *string `json:"maternalOrganismName,omitempty"`

	// The identifier of the paternal species constituting the hybrid organism
	PaternalOrganismId *string `json:"paternalOrganismId,omitempty"`

	// The name of the paternal species constituting the hybrid organism
	PaternalOrganismName *string `json:"paternalOrganismName,omitempty"`
}

// SubstanceSourceMaterialOrganismOrganismGeneral represents organism general information
type SubstanceSourceMaterialOrganismOrganismGeneral struct {
	common.BackboneElement

	// The class of an organism shall be specified
	Class *common.CodeableConcept `json:"class,omitempty"`

	// The kingdom of an organism shall be specified
	Kingdom *common.CodeableConcept `json:"kingdom,omitempty"`

	// The order of an organism shall be specified
	Order *common.CodeableConcept `json:"order,omitempty"`

	// The phylum of an organism shall be specified
	Phylum *common.CodeableConcept `json:"phylum,omitempty"`
}

// SubstanceSourceMaterialOrganism represents the organism which the substance is derived from
type SubstanceSourceMaterialOrganism struct {
	common.BackboneElement

	// Author type information
	Author []SubstanceSourceMaterialOrganismAuthor `json:"author,omitempty"`

	// The family of an organism shall be specified
	Family *common.CodeableConcept `json:"family,omitempty"`

	// The genus of an organism shall be specified
	Genus *common.CodeableConcept `json:"genus,omitempty"`

	// Hybrid species information
	Hybrid *SubstanceSourceMaterialOrganismHybrid `json:"hybrid,omitempty"`

	// The intraspecific description of an organism shall be specified
	IntraspecificDescription *string `json:"intraspecificDescription,omitempty"`

	// The Intraspecific type of an organism shall be specified
	IntraspecificType *common.CodeableConcept `json:"intraspecificType,omitempty"`

	// Organism general information
	OrganismGeneral *SubstanceSourceMaterialOrganismOrganismGeneral `json:"organismGeneral,omitempty"`

	// The species of an organism shall be specified
	Species *common.CodeableConcept `json:"species,omitempty"`
}

// SubstanceSourceMaterialPartDescription represents part description information
type SubstanceSourceMaterialPartDescription struct {
	common.BackboneElement

	// Entity of anatomical origin of source material within an organism
	Part *common.CodeableConcept `json:"part,omitempty"`

	// The detailed anatomic location when the part can be extracted from different anatomical locations of the organism
	PartLocation *common.CodeableConcept `json:"partLocation,omitempty"`
}

// SubstanceSourceMaterial represents source material for substances
type SubstanceSourceMaterial struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SubstanceSourceMaterial"

	// The country where the plant material is harvested or the countries where the plasma is sourced from
	CountryOfOrigin []common.CodeableConcept `json:"countryOfOrigin,omitempty"`

	// Stage of life for animals, plants, insects and microorganisms
	DevelopmentStage *common.CodeableConcept `json:"developmentStage,omitempty"`

	// Many complex materials are fractions of parts of plants, animals, or minerals
	FractionDescription []SubstanceSourceMaterialFractionDescription `json:"fractionDescription,omitempty"`

	// The place/region where the plant is harvested or the places/regions where the animal source material has its habitat
	GeographicalLocation []string `json:"geographicalLocation,omitempty"`

	// This subclause describes the organism which the substance is derived from
	Organism *SubstanceSourceMaterialOrganism `json:"organism,omitempty"`

	// The unique identifier associated with the source material parent organism
	OrganismId *common.Identifier `json:"organismId,omitempty"`

	// The organism accepted Scientific name shall be provided based on the organism taxonomy
	OrganismName *string `json:"organismName,omitempty"`

	// The parent of the herbal drug
	ParentSubstanceId []common.Identifier `json:"parentSubstanceId,omitempty"`

	// The parent substance of the Herbal Drug, or Herbal preparation
	ParentSubstanceName []string `json:"parentSubstanceName,omitempty"`

	// Part description information
	PartDescription []SubstanceSourceMaterialPartDescription `json:"partDescription,omitempty"`

	// General high level classification of the source material specific to the origin of the material
	SourceMaterialClass *common.CodeableConcept `json:"sourceMaterialClass,omitempty"`

	// The state of the source material when extracted
	SourceMaterialState *common.CodeableConcept `json:"sourceMaterialState,omitempty"`

	// The type of the source material shall be specified based on a controlled vocabulary
	SourceMaterialType *common.CodeableConcept `json:"sourceMaterialType,omitempty"`
}
