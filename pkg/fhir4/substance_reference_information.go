package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SubstanceReferenceInformation represents todo
type SubstanceReferenceInformation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SubstanceReferenceInformation"

	// Todo
	Classification []SubstanceReferenceInformationClassification `json:"classification,omitempty"`

	// Todo
	Comment *string `json:"comment,omitempty"`

	// Todo
	Gene []SubstanceReferenceInformationGene `json:"gene,omitempty"`

	// Todo
	GeneElement []SubstanceReferenceInformationGeneElement `json:"geneElement,omitempty"`

	// Todo
	Target []SubstanceReferenceInformationTarget `json:"target,omitempty"`
}

// SubstanceReferenceInformationGene represents todo
type SubstanceReferenceInformationGene struct {
	common.BackboneElement

	// Todo
	Gene *common.CodeableConcept `json:"gene,omitempty"`

	// Todo
	GeneSequenceOrigin *common.CodeableConcept `json:"geneSequenceOrigin,omitempty"`

	// Todo
	Source []common.Reference `json:"source,omitempty"`
}

// SubstanceReferenceInformationGeneElement represents todo
type SubstanceReferenceInformationGeneElement struct {
	common.BackboneElement

	// Todo
	Element *common.Identifier `json:"element,omitempty"`

	// Todo
	Source []common.Reference `json:"source,omitempty"`

	// Todo
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SubstanceReferenceInformationClassification represents todo
type SubstanceReferenceInformationClassification struct {
	common.BackboneElement

	// Todo
	Classification *common.CodeableConcept `json:"classification,omitempty"`

	// Todo
	Domain *common.CodeableConcept `json:"domain,omitempty"`

	// Todo
	Source []common.Reference `json:"source,omitempty"`

	// Todo
	Subtype []common.CodeableConcept `json:"subtype,omitempty"`
}

// SubstanceReferenceInformationTarget represents todo
type SubstanceReferenceInformationTarget struct {
	common.BackboneElement

	// Todo
	AmountQuantity *common.Quantity `json:"amountQuantity,omitempty"`

	// Todo
	AmountRange *common.Range `json:"amountRange,omitempty"`

	// Todo
	AmountString *string `json:"amountString,omitempty"`

	// Todo
	AmountType *common.CodeableConcept `json:"amountType,omitempty"`

	// Todo
	Interaction *common.CodeableConcept `json:"interaction,omitempty"`

	// Todo
	Organism *common.CodeableConcept `json:"organism,omitempty"`

	// Todo
	OrganismType *common.CodeableConcept `json:"organismType,omitempty"`

	// Todo
	Source []common.Reference `json:"source,omitempty"`

	// Todo
	Target *common.Identifier `json:"target,omitempty"`

	// Todo
	Type *common.CodeableConcept `json:"type,omitempty"`
}
