package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

type MolecularSequenceType string

type MolecularSequenceRelativeStartingSequenceOrientation string

type MolecularSequenceRelativeStartingSequenceStrand string

type MolecularSequenceRelativeStartingSequence struct {
	common.BackboneElement

	Chromosome              *common.CodeableConcept                               `json:"chromosome,omitempty"`
	GenomeAssembly          *common.CodeableConcept                               `json:"genomeAssembly,omitempty"`
	Orientation             *MolecularSequenceRelativeStartingSequenceOrientation `json:"orientation,omitempty"`
	SequenceCodeableConcept *common.CodeableConcept                               `json:"sequenceCodeableConcept,omitempty"`
	SequenceString          *string                                               `json:"sequenceString,omitempty"`
	SequenceReference       *common.Reference                                     `json:"sequenceReference,omitempty"`
	Strand                  *MolecularSequenceRelativeStartingSequenceStrand      `json:"strand,omitempty"`
	WindowEnd               *int                                                  `json:"windowEnd,omitempty"`
	WindowStart             *int                                                  `json:"windowStart,omitempty"`
}

type MolecularSequenceRelativeEdit struct {
	common.BackboneElement

	End                 *int    `json:"end,omitempty"`
	ReplacedSequence    *string `json:"replacedSequence,omitempty"`
	ReplacementSequence *string `json:"replacementSequence,omitempty"`
	Start               *int    `json:"start,omitempty"`
}

type MolecularSequenceRelative struct {
	common.BackboneElement

	CoordinateSystem common.CodeableConcept                     `json:"coordinateSystem"`
	Edit             []MolecularSequenceRelativeEdit            `json:"edit,omitempty"`
	OrdinalPosition  *int                                       `json:"ordinalPosition,omitempty"`
	SequenceRange    *Range                                     `json:"sequenceRange,omitempty"`
	StartingSequence *MolecularSequenceRelativeStartingSequence `json:"startingSequence,omitempty"`
}

type MolecularSequence struct {
	DomainResource

	ResourceType string                      `json:"resourceType"` // Always "MolecularSequence"
	Device       *common.Reference           `json:"device,omitempty"`
	Focus        []common.Reference          `json:"focus,omitempty"`
	Formatted    []Attachment                `json:"formatted,omitempty"`
	Identifier   []common.Identifier         `json:"identifier,omitempty"`
	Literal      *string                     `json:"literal,omitempty"`
	Performer    *common.Reference           `json:"performer,omitempty"`
	Relative     []MolecularSequenceRelative `json:"relative,omitempty"`
	Specimen     *common.Reference           `json:"specimen,omitempty"`
	Subject      *common.Reference           `json:"subject,omitempty"`
	Type         *MolecularSequenceType      `json:"type,omitempty"`
}
