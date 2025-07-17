// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SubstanceProteinSubunit represents subunits constituting the SubstanceProtein
type SubstanceProteinSubunit struct {
	common.BackboneElement

	// The modification at the C-terminal shall be specified
	CTerminalModification *string `json:"cTerminalModification,omitempty"`

	// Unique identifier for molecular fragment modification based on the ISO 11238 Substance ID
	CTerminalModificationId *common.Identifier `json:"cTerminalModificationId,omitempty"`

	// Length of linear sequences of amino acids contained in the subunit
	Length *int `json:"length,omitempty"`

	// The name of the fragment modified at the N-terminal of the SubstanceProtein shall be specified
	NTerminalModification *string `json:"nTerminalModification,omitempty"`

	// Unique identifier for molecular fragment modification based on the ISO 11238 Substance ID
	NTerminalModificationId *common.Identifier `json:"nTerminalModificationId,omitempty"`

	// The sequence information shall be provided enumerating the amino acids from N- to C-terminal end
	Sequence *string `json:"sequence,omitempty"`

	// The sequence information shall be provided enumerating the amino acids from N- to C-terminal end
	SequenceAttachment *Attachment `json:"sequenceAttachment,omitempty"`

	// Index of primary sequences of amino acids linked through peptide bonds in order of decreasing length
	Subunit *int `json:"subunit,omitempty"`
}

// SubstanceProtein represents a single unit of a linear amino acid sequence
type SubstanceProtein struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SubstanceProtein"

	// The disulphide bond between two cysteine residues either on the same subunit or on two different subunits
	DisulfideLinkage []string `json:"disulfideLinkage,omitempty"`

	// Number of linear sequences of amino acids linked through peptide bonds
	NumberOfSubunits *int `json:"numberOfSubunits,omitempty"`

	// The SubstanceProtein descriptive elements will only be used when a complete or partial amino acid sequence is available
	SequenceType *common.CodeableConcept `json:"sequenceType,omitempty"`

	// This subclause refers to the description of each subunit constituting the SubstanceProtein
	Subunit []SubstanceProteinSubunit `json:"subunit,omitempty"`
}
