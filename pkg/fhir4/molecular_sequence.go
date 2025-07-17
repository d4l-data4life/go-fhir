package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MolecularSequence represents raw data describing a biological sequence
type MolecularSequence struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MolecularSequence"

	// Whether the sequence is numbered starting at 0 (0-based numbering or coordinates, inclusive start, exclusive end) or starting at 1 (1-based numbering, inclusive start and inclusive end)
	CoordinateSystem int `json:"coordinateSystem"`

	// The method for sequencing, for example, chip information
	Device *common.Reference `json:"device,omitempty"`

	// A unique identifier for this particular sequence instance
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Sequence that was observed
	ObservedSeq *string `json:"observedSeq,omitempty"`

	// The patient whose sequencing results are described by this resource
	Patient *common.Reference `json:"patient,omitempty"`

	// The organization or lab that should be responsible for this result
	Performer *common.Reference `json:"performer,omitempty"`

	// Pointer to next atomic sequence which at most contains one variant
	Pointer []common.Reference `json:"pointer,omitempty"`

	// An experimental feature attribute that defines the quality of the feature in a quantitative way
	Quality []MolecularSequenceQuality `json:"quality,omitempty"`

	// The number of copies of the sequence of interest (RNASeq)
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Coverage (read depth or depth) is the average number of reads representing a given nucleotide in the reconstructed sequence
	ReadCoverage *int `json:"readCoverage,omitempty"`

	// A sequence that is used as a reference to describe variants that are present in a sequence analyzed
	ReferenceSeq *MolecularSequenceReferenceSeq `json:"referenceSeq,omitempty"`

	// Configurations of the external repository
	Repository []MolecularSequenceRepository `json:"repository,omitempty"`

	// Specimen used for sequencing
	Specimen *common.Reference `json:"specimen,omitempty"`

	// Information about chromosome structure variation
	StructureVariant []MolecularSequenceStructureVariant `json:"structureVariant,omitempty"`

	// Amino Acid Sequence/ DNA Sequence / RNA Sequence
	Type *MolecularSequenceType `json:"type,omitempty"`

	// The definition of variant here originates from Sequence ontology
	Variant []MolecularSequenceVariant `json:"variant,omitempty"`
}

// MolecularSequenceType represents the type of sequence
type MolecularSequenceType string

const (
	MolecularSequenceTypeAA  MolecularSequenceType = "aa"
	MolecularSequenceTypeDNA MolecularSequenceType = "dna"
	MolecularSequenceTypeRNA MolecularSequenceType = "rna"
)

// MolecularSequenceReferenceSeq represents a sequence that is used as a reference to describe variants
type MolecularSequenceReferenceSeq struct {
	common.BackboneElement

	// Structural unit composed of a nucleic acid molecule which controls its own replication
	Chromosome *common.CodeableConcept `json:"chromosome,omitempty"`

	// The Genome Build used for reference, following GRCh build versions e.g. 'GRCh 37'
	GenomeBuild *string `json:"genomeBuild,omitempty"`

	// A relative reference to a DNA strand based on gene orientation
	Orientation *MolecularSequenceReferenceSeqOrientation `json:"orientation,omitempty"`

	// Reference identifier of reference sequence submitted to NCBI
	ReferenceSeqId *common.CodeableConcept `json:"referenceSeqId,omitempty"`

	// A pointer to another MolecularSequence entity as reference sequence
	ReferenceSeqPointer *common.Reference `json:"referenceSeqPointer,omitempty"`

	// A string like "ACGT"
	ReferenceSeqString *string `json:"referenceSeqString,omitempty"`

	// An absolute reference to a strand
	Strand *MolecularSequenceReferenceSeqStrand `json:"strand,omitempty"`

	// End position of the window on the reference sequence
	WindowEnd *int `json:"windowEnd,omitempty"`

	// Start position of the window on the reference sequence
	WindowStart *int `json:"windowStart,omitempty"`
}

// MolecularSequenceReferenceSeqOrientation represents a relative reference to a DNA strand
type MolecularSequenceReferenceSeqOrientation string

const (
	MolecularSequenceReferenceSeqOrientationSense     MolecularSequenceReferenceSeqOrientation = "sense"
	MolecularSequenceReferenceSeqOrientationAntisense MolecularSequenceReferenceSeqOrientation = "antisense"
)

// MolecularSequenceReferenceSeqStrand represents an absolute reference to a strand
type MolecularSequenceReferenceSeqStrand string

const (
	MolecularSequenceReferenceSeqStrandWatson MolecularSequenceReferenceSeqStrand = "watson"
	MolecularSequenceReferenceSeqStrandCrick  MolecularSequenceReferenceSeqStrand = "crick"
)

// MolecularSequenceVariant represents the definition of variant here originates from Sequence ontology
type MolecularSequenceVariant struct {
	common.BackboneElement

	// Extended CIGAR string for aligning the sequence with reference bases
	Cigar *string `json:"cigar,omitempty"`

	// End position of the variant on the reference sequence
	End *int `json:"end,omitempty"`

	// An allele is one of a set of coexisting sequence variants of a gene
	ObservedAllele *string `json:"observedAllele,omitempty"`

	// An allele is one of a set of coexisting sequence variants of a gene
	ReferenceAllele *string `json:"referenceAllele,omitempty"`

	// Start position of the variant on the reference sequence
	Start *int `json:"start,omitempty"`

	// A pointer to an Observation containing variant information
	VariantPointer *common.Reference `json:"variantPointer,omitempty"`
}

// MolecularSequenceQuality represents an experimental feature attribute that defines the quality of the feature
type MolecularSequenceQuality struct {
	common.BackboneElement

	// End position of the sequence
	End *int `json:"end,omitempty"`

	// Harmonic mean of Recall and Precision
	FScore *float64 `json:"fScore,omitempty"`

	// The number of false positives where the non-REF alleles in the Truth and Query Call Sets match
	GtFP *int `json:"gtFP,omitempty"`

	// Which method is used to get sequence quality
	Method *common.CodeableConcept `json:"method,omitempty"`

	// QUERY.TP / (QUERY.TP + QUERY.FP)
	Precision *float64 `json:"precision,omitempty"`

	// False positives, i.e. the number of sites in the Query Call Set for which there is no path through the Truth Call Set
	QueryFP *int `json:"queryFP,omitempty"`

	// True positives, from the perspective of the query data
	QueryTP *int `json:"queryTP,omitempty"`

	// TRUTH.TP / (TRUTH.TP + TRUTH.FN)
	Recall *float64 `json:"recall,omitempty"`

	// Receiver Operator Characteristic (ROC) Curve to give sensitivity/specificity tradeoff
	Roc *MolecularSequenceQualityRoc `json:"roc,omitempty"`

	// The score of an experimentally derived feature such as a p-value
	Score *common.Quantity `json:"score,omitempty"`

	// Gold standard sequence used for comparing against
	StandardSequence *common.CodeableConcept `json:"standardSequence,omitempty"`

	// Start position of the sequence
	Start *int `json:"start,omitempty"`

	// False negatives, i.e. the number of sites in the Truth Call Set for which there is no path through the Query Call Set
	TruthFN *int `json:"truthFN,omitempty"`

	// True positives, from the perspective of the truth data
	TruthTP *int `json:"truthTP,omitempty"`

	// INDEL / SNP / Undefined variant
	Type MolecularSequenceQualityType `json:"type"`
}

// MolecularSequenceQualityType represents the type of quality
type MolecularSequenceQualityType string

const (
	MolecularSequenceQualityTypeIndel   MolecularSequenceQualityType = "indel"
	MolecularSequenceQualityTypeSNP     MolecularSequenceQualityType = "snp"
	MolecularSequenceQualityTypeUnknown MolecularSequenceQualityType = "unknown"
)

// MolecularSequenceQualityRoc represents Receiver Operator Characteristic (ROC) Curve
type MolecularSequenceQualityRoc struct {
	common.BackboneElement

	// Calculated fScore if the GQ score threshold was set to "score" field value
	FMeasure []float64 `json:"fMeasure,omitempty"`

	// The number of false negatives if the GQ score threshold was set to "score" field value
	NumFN []int `json:"numFN,omitempty"`

	// The number of false positives if the GQ score threshold was set to "score" field value
	NumFP []int `json:"numFP,omitempty"`

	// The number of true positives if the GQ score threshold was set to "score" field value
	NumTP []int `json:"numTP,omitempty"`

	// Calculated precision if the GQ score threshold was set to "score" field value
	Precision []float64 `json:"precision,omitempty"`

	// Individual data point representing the GQ (genotype quality) score threshold
	Score []int `json:"score,omitempty"`

	// Calculated sensitivity if the GQ score threshold was set to "score" field value
	Sensitivity []float64 `json:"sensitivity,omitempty"`
}

// MolecularSequenceRepository represents configurations of the external repository
type MolecularSequenceRepository struct {
	common.BackboneElement

	// Id of the variant in this external repository
	DatasetId *string `json:"datasetId,omitempty"`

	// URI of an external repository which contains further details about the genetics data
	Name *string `json:"name,omitempty"`

	// Id of the read in this external repository
	ReadsetId *string `json:"readsetId,omitempty"`

	// Click and see / RESTful API / Need login to see / RESTful API with authentication / Other ways to see resource
	Type MolecularSequenceRepositoryType `json:"type"`

	// URI of an external repository which contains further details about the genetics data
	Url *string `json:"url,omitempty"`

	// Id of the variantset in this external repository
	VariantsetId *string `json:"variantsetId,omitempty"`
}

// MolecularSequenceRepositoryType represents the type of repository
type MolecularSequenceRepositoryType string

const (
	MolecularSequenceRepositoryTypeDirectlink MolecularSequenceRepositoryType = "directlink"
	MolecularSequenceRepositoryTypeOpenapi    MolecularSequenceRepositoryType = "openapi"
	MolecularSequenceRepositoryTypeLogin      MolecularSequenceRepositoryType = "login"
	MolecularSequenceRepositoryTypeOauth      MolecularSequenceRepositoryType = "oauth"
	MolecularSequenceRepositoryTypeOther      MolecularSequenceRepositoryType = "other"
)

// MolecularSequenceStructureVariant represents information about chromosome structure variation
type MolecularSequenceStructureVariant struct {
	common.BackboneElement

	// Used to indicate if the outer and inner start-end values have the same meaning
	Exact *bool `json:"exact,omitempty"`

	// Structural variant inner
	Inner *MolecularSequenceStructureVariantInner `json:"inner,omitempty"`

	// Length of the variant chromosome
	Length *int `json:"length,omitempty"`

	// Structural variant outer
	Outer *MolecularSequenceStructureVariantOuter `json:"outer,omitempty"`

	// Information about chromosome structure variation
	VariantType *common.CodeableConcept `json:"variantType,omitempty"`
}

// MolecularSequenceStructureVariantOuter represents structural variant outer
type MolecularSequenceStructureVariantOuter struct {
	common.BackboneElement

	// Structural variant outer end
	End *int `json:"end,omitempty"`

	// Structural variant outer start
	Start *int `json:"start,omitempty"`
}

// MolecularSequenceStructureVariantInner represents structural variant inner
type MolecularSequenceStructureVariantInner struct {
	common.BackboneElement

	// Structural variant inner end
	End *int `json:"end,omitempty"`

	// Structural variant inner start
	Start *int `json:"start,omitempty"`
}
