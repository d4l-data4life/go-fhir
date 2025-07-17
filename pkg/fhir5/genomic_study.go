package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// GenomicStudyStatus represents the status of a genomic study
type GenomicStudyStatus string

const (
	GenomicStudyStatusRegistered     GenomicStudyStatus = "registered"
	GenomicStudyStatusAvailable      GenomicStudyStatus = "available"
	GenomicStudyStatusCancelled      GenomicStudyStatus = "cancelled"
	GenomicStudyStatusEnteredInError GenomicStudyStatus = "entered-in-error"
	GenomicStudyStatusUnknown        GenomicStudyStatus = "unknown"
)

// GenomicStudyAnalysisPerformer represents the performer of an analysis
type GenomicStudyAnalysisPerformer struct {
	common.BackboneElement

	// Performer for the analysis event
	Actor *common.Reference `json:"actor,omitempty"`

	// Indicates the role of the performer
	Role *common.CodeableConcept `json:"role,omitempty"`
}

// GenomicStudyAnalysisInput represents inputs for the analysis event
type GenomicStudyAnalysisInput struct {
	common.BackboneElement

	// File containing input data
	File *common.Reference `json:"file,omitempty"`

	// The analysis event or other GenomicStudy that generated this input file
	GeneratedByIdentifier *common.Identifier `json:"generatedByIdentifier,omitempty"`

	// The analysis event or other GenomicStudy that generated this input file
	GeneratedByReference *common.Reference `json:"generatedByReference,omitempty"`

	// Type of input data, e.g., BAM, CRAM, or FASTA
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// GenomicStudyAnalysisOutput represents outputs for the analysis event
type GenomicStudyAnalysisOutput struct {
	common.BackboneElement

	// File containing output data
	File *common.Reference `json:"file,omitempty"`

	// Type of output data, e.g., VCF, MAF, or BAM
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// GenomicStudyAnalysisDevice represents devices used for the analysis
type GenomicStudyAnalysisDevice struct {
	common.BackboneElement

	// Device used for the analysis
	Device *common.Reference `json:"device,omitempty"`

	// Specific function for the device used for the analysis
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// GenomicStudyAnalysis represents the details about a specific analysis
type GenomicStudyAnalysis struct {
	common.BackboneElement

	// Type of the genomic changes studied in the analysis
	ChangeType []common.CodeableConcept `json:"changeType,omitempty"`

	// The date of the analysis event
	Date *string `json:"date,omitempty"`

	// Devices used for the analysis
	Device []GenomicStudyAnalysisDevice `json:"device,omitempty"`

	// The focus of the analysis
	Focus []common.Reference `json:"focus,omitempty"`

	// The reference genome build that is used in this analysis
	GenomeBuild *common.CodeableConcept `json:"genomeBuild,omitempty"`

	// Identifiers for the analysis event
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Inputs for the analysis event
	Input []GenomicStudyAnalysisInput `json:"input,omitempty"`

	// The defined protocol that describes the analysis
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`

	// The URL pointing to an externally maintained protocol
	InstantiatesUri *string `json:"instantiatesUri,omitempty"`

	// Type of the methods used in the analysis
	MethodType []common.CodeableConcept `json:"methodType,omitempty"`

	// Any notes capture with the analysis event
	Note []Annotation `json:"note,omitempty"`

	// Outputs for the analysis event
	Output []GenomicStudyAnalysisOutput `json:"output,omitempty"`

	// Performer for the analysis event
	Performer []GenomicStudyAnalysisPerformer `json:"performer,omitempty"`

	// The protocol that was performed for the analysis event
	ProtocolPerformed *common.Reference `json:"protocolPerformed,omitempty"`

	// Genomic regions actually called in the analysis event
	RegionsCalled []common.Reference `json:"regionsCalled,omitempty"`

	// The genomic regions to be studied in the analysis
	RegionsStudied []common.Reference `json:"regionsStudied,omitempty"`

	// The specimen used in the analysis event
	Specimen []common.Reference `json:"specimen,omitempty"`

	// Name of the analysis event (human friendly)
	Title *string `json:"title,omitempty"`
}

// GenomicStudy represents a set of analyses performed to analyze and generate genomic data
type GenomicStudy struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "GenomicStudy"

	// The details about a specific analysis that was performed
	Analysis []GenomicStudyAnalysis `json:"analysis,omitempty"`

	// Event resources that the genomic study is based on
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Description of the genomic study
	Description *string `json:"description,omitempty"`

	// The healthcare event with which this genomics study is associated
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Identifiers for this genomic study
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The defined protocol that describes the study
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`

	// The URL pointing to an externally maintained protocol
	InstantiatesUri *string `json:"instantiatesUri,omitempty"`

	// Healthcare professionals who interpreted the genomic study
	Interpreter []common.Reference `json:"interpreter,omitempty"`

	// Comments related to the genomic study
	Note []Annotation `json:"note,omitempty"`

	// Why the genomic study was performed
	Reason []CodeableReference `json:"reason,omitempty"`

	// Healthcare professional who requested or referred the genomic study
	Referrer *common.Reference `json:"referrer,omitempty"`

	// When the genomic study was started
	StartDate *string `json:"startDate,omitempty"`

	// The status of the genomic study
	Status GenomicStudyStatus `json:"status"`

	// The primary subject of the genomic study
	Subject common.Reference `json:"subject"`

	// The type of the study
	Type []common.CodeableConcept `json:"type,omitempty"`
}
