package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// RiskEvidenceSynthesis represents the RiskEvidenceSynthesis resource describes the likelihood of an outcome in a population plus exposure state where the risk estimate is derived from a combination of research studies
type RiskEvidenceSynthesis struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "RiskEvidenceSynthesis"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individiual or organization primarily involved in the creation and maintenance of the content
	Author []common.ContactDetail `json:"author,omitempty"`

	// A description of the certainty of the risk estimate
	Certainty []RiskEvidenceSynthesisCertainty `json:"certainty,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the risk evidence synthesis and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the risk evidence synthesis was built
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []common.ContactDetail `json:"editor,omitempty"`

	// The effective period for a risk evidence synthesis determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// An individual or organization responsible for officially endorsing the content for use in some setting
	Endorser []common.ContactDetail `json:"endorser,omitempty"`

	// A reference to a EvidenceVariable resource that defines the exposure for the research
	Exposure *common.Reference `json:"exposure,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the risk evidence synthesis to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// A human-readable string to clarify or explain concepts about the resource
	Note []common.Annotation `json:"note,omitempty"`

	// A reference to a EvidenceVariable resomece that defines the outcome for the research
	Outcome common.Reference `json:"outcome"`

	// A reference to a EvidenceVariable resource that defines the population for the research
	Population common.Reference `json:"population"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// Each related artifact is either an attachment, or a reference to another resource, but not both
	RelatedArtifact []common.RelatedArtifact `json:"relatedArtifact,omitempty"`

	// An individual or organization primarily responsible for review of some aspect of the content
	Reviewer []common.ContactDetail `json:"reviewer,omitempty"`

	// The estimated risk of the outcome
	RiskEstimate *RiskEvidenceSynthesisRiskEstimate `json:"riskEstimate,omitempty"`

	// A description of the size of the sample involved in the synthesis
	SampleSize *RiskEvidenceSynthesisSampleSize `json:"sampleSize,omitempty"`

	// Allows filtering of risk evidence synthesiss that are appropriate for use versus not
	Status RiskEvidenceSynthesisStatus `json:"status"`

	// Type of study eg randomized trial
	StudyType *common.CodeableConcept `json:"studyType,omitempty"`

	// Type of synthesis eg meta-analysis
	SynthesisType *common.CodeableConcept `json:"synthesisType,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Descriptive topics related to the content of the RiskEvidenceSynthesis
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different risk evidence synthesis instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// RiskEvidenceSynthesisStatus represents allows filtering of risk evidence synthesiss that are appropriate for use versus not
type RiskEvidenceSynthesisStatus string

const (
	RiskEvidenceSynthesisStatusDraft   RiskEvidenceSynthesisStatus = "draft"
	RiskEvidenceSynthesisStatusActive  RiskEvidenceSynthesisStatus = "active"
	RiskEvidenceSynthesisStatusRetired RiskEvidenceSynthesisStatus = "retired"
	RiskEvidenceSynthesisStatusUnknown RiskEvidenceSynthesisStatus = "unknown"
)

// RiskEvidenceSynthesisSampleSize represents a description of the size of the sample involved in the synthesis
type RiskEvidenceSynthesisSampleSize struct {
	common.BackboneElement

	// Human-readable summary of sample size
	Description *string `json:"description,omitempty"`

	// Number of participants included in this evidence synthesis
	NumberOfParticipants *int `json:"numberOfParticipants,omitempty"`

	// Number of studies included in this evidence synthesis
	NumberOfStudies *int `json:"numberOfStudies,omitempty"`
}

// RiskEvidenceSynthesisRiskEstimate represents the estimated risk of the outcome
type RiskEvidenceSynthesisRiskEstimate struct {
	common.BackboneElement

	// The sample size for the group that was measured for this risk estimate
	DenominatorCount *int `json:"denominatorCount,omitempty"`

	// Human-readable summary of risk estimate
	Description *string `json:"description,omitempty"`

	// The number of group members with the outcome of interest
	NumeratorCount *int `json:"numeratorCount,omitempty"`

	// A description of the precision of the estimate for the effect
	PrecisionEstimate []RiskEvidenceSynthesisRiskEstimatePrecisionEstimate `json:"precisionEstimate,omitempty"`

	// Examples include proportion and mean
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Specifies the UCUM unit for the outcome
	UnitOfMeasure *common.CodeableConcept `json:"unitOfMeasure,omitempty"`

	// The point estimate of the risk estimate
	Value *float64 `json:"value,omitempty"`
}

// RiskEvidenceSynthesisRiskEstimatePrecisionEstimate represents a description of the precision of the estimate for the effect
type RiskEvidenceSynthesisRiskEstimatePrecisionEstimate struct {
	common.BackboneElement

	// Lower bound of confidence interval
	From *float64 `json:"from,omitempty"`

	// Use 95 for a 95% confidence interval
	Level *float64 `json:"level,omitempty"`

	// Upper bound of confidence interval
	To *float64 `json:"to,omitempty"`

	// Examples include confidence interval and interquartile range
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// RiskEvidenceSynthesisCertainty represents a description of the certainty of the risk estimate
type RiskEvidenceSynthesisCertainty struct {
	common.BackboneElement

	// A description of a component of the overall certainty
	CertaintySubcomponent []RiskEvidenceSynthesisCertaintyCertaintySubcomponent `json:"certaintySubcomponent,omitempty"`

	// A human-readable string to clarify or explain concepts about the resource
	Note []common.Annotation `json:"note,omitempty"`

	// A rating of the certainty of the effect estimate
	Rating []common.CodeableConcept `json:"rating,omitempty"`
}

// RiskEvidenceSynthesisCertaintyCertaintySubcomponent represents a description of a component of the overall certainty
type RiskEvidenceSynthesisCertaintyCertaintySubcomponent struct {
	common.BackboneElement

	// A human-readable string to clarify or explain concepts about the resource
	Note []common.Annotation `json:"note,omitempty"`

	// A rating of a subcomponent of rating certainty
	Rating []common.CodeableConcept `json:"rating,omitempty"`

	// Type of subcomponent of certainty rating
	Type *common.CodeableConcept `json:"type,omitempty"`
}
