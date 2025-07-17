// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SubstanceNucleicAcidSubunitLinkage represents the linkages between sugar residues
type SubstanceNucleicAcidSubunitLinkage struct {
	common.BackboneElement

	// The entity that links the sugar residues together should also be captured
	Connectivity *string `json:"connectivity,omitempty"`

	// Each linkage will be registered as a fragment and have an ID
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Each linkage will be registered as a fragment and have at least one name
	Name *string `json:"name,omitempty"`

	// Residues shall be captured as described
	ResidueSite *string `json:"residueSite,omitempty"`
}

// SubstanceNucleicAcidSubunitSugar represents sugar ID information
type SubstanceNucleicAcidSubunitSugar struct {
	common.BackboneElement

	// The Substance ID of the sugar or sugar-like component that make up the nucleotide
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// The name of the sugar or sugar-like component that make up the nucleotide
	Name *string `json:"name,omitempty"`

	// The residues that contain a given sugar will be captured
	ResidueSite *string `json:"residueSite,omitempty"`
}

// SubstanceNucleicAcidSubunit represents subunits of nucleic acids
type SubstanceNucleicAcidSubunit struct {
	common.BackboneElement

	// The nucleotide present at the 5' terminal shall be specified based on a controlled vocabulary
	FivePrime *common.CodeableConcept `json:"fivePrime,omitempty"`

	// The length of the sequence shall be captured
	Length *int `json:"length,omitempty"`

	// The linkages between sugar residues will also be captured
	Linkage []SubstanceNucleicAcidSubunitLinkage `json:"linkage,omitempty"`

	// Actual nucleotide sequence notation from 5' to 3' end using standard single letter codes
	Sequence *string `json:"sequence,omitempty"`

	// Sequence attachment
	SequenceAttachment *Attachment `json:"sequenceAttachment,omitempty"`

	// Index of linear sequences of nucleic acids in order of decreasing length
	Subunit *int `json:"subunit,omitempty"`

	// Sugar information
	Sugar []SubstanceNucleicAcidSubunitSugar `json:"sugar,omitempty"`

	// The nucleotide present at the 3' terminal shall be specified based on a controlled vocabulary
	ThreePrime *common.CodeableConcept `json:"threePrime,omitempty"`
}

// SubstanceNucleicAcid represents nucleic acids defined by three distinct elements
type SubstanceNucleicAcid struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SubstanceNucleicAcid"

	// The area of hybridisation shall be described if applicable for double stranded RNA or DNA
	AreaOfHybridisation *string `json:"areaOfHybridisation,omitempty"`

	// The number of linear sequences of nucleotides linked through phosphodiester bonds shall be described
	NumberOfSubunits *int `json:"numberOfSubunits,omitempty"`

	// Oligo nucleotide type
	OligoNucleotideType *common.CodeableConcept `json:"oligoNucleotideType,omitempty"`

	// The type of the sequence shall be specified based on a controlled vocabulary
	SequenceType *common.CodeableConcept `json:"sequenceType,omitempty"`

	// Subunits are listed in order of decreasing length
	Subunit []SubstanceNucleicAcidSubunit `json:"subunit,omitempty"`
}
