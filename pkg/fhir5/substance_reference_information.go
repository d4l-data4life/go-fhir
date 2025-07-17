// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SubstanceReferenceInformationGene represents gene information
type SubstanceReferenceInformationGene struct {
	common.BackboneElement

	// Gene information
	Gene *common.CodeableConcept `json:"gene,omitempty"`

	// Gene sequence origin
	GeneSequenceOrigin *common.CodeableConcept `json:"geneSequenceOrigin,omitempty"`

	// Source information
	Source []common.Reference `json:"source,omitempty"`
}

// SubstanceReferenceInformationGeneElement represents gene element information
type SubstanceReferenceInformationGeneElement struct {
	common.BackboneElement

	// Element information
	Element *common.Identifier `json:"element,omitempty"`

	// Source information
	Source []common.Reference `json:"source,omitempty"`

	// Type information
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SubstanceReferenceInformationTarget represents target information
type SubstanceReferenceInformationTarget struct {
	common.BackboneElement

	// Amount quantity
	AmountQuantity *common.Quantity `json:"amountQuantity,omitempty"`

	// Amount range
	AmountRange *Range `json:"amountRange,omitempty"`

	// Amount string
	AmountString *string `json:"amountString,omitempty"`

	// Amount type
	AmountType *common.CodeableConcept `json:"amountType,omitempty"`

	// Interaction information
	Interaction *common.CodeableConcept `json:"interaction,omitempty"`

	// Organism information
	Organism *common.CodeableConcept `json:"organism,omitempty"`

	// Organism type
	OrganismType *common.CodeableConcept `json:"organismType,omitempty"`

	// Source information
	Source []common.Reference `json:"source,omitempty"`

	// Target information
	Target *common.Identifier `json:"target,omitempty"`

	// Type information
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SubstanceReferenceInformation represents reference information for substances
type SubstanceReferenceInformation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SubstanceReferenceInformation"

	// Comment information
	Comment *string `json:"comment,omitempty"`

	// Gene information
	Gene []SubstanceReferenceInformationGene `json:"gene,omitempty"`

	// Gene element information
	GeneElement []SubstanceReferenceInformationGeneElement `json:"geneElement,omitempty"`

	// Target information
	Target []SubstanceReferenceInformationTarget `json:"target,omitempty"`
}
