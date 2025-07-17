package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// EffectEvidenceSynthesis represents the EffectEvidenceSynthesis resource
type EffectEvidenceSynthesis struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "EffectEvidenceSynthesis"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance of the content
	Author []common.ContactDetail `json:"author,omitempty"`

	// A description of the certainty of the effect estimate
	Certainty []EffectEvidenceSynthesisCertainty `json:"certainty,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the effect evidence synthesis and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the effect evidence synthesis was built
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []common.ContactDetail `json:"editor,omitempty"`

	// The estimated effect of the exposure variant
	EffectEstimate []EffectEvidenceSynthesisEffectEstimate `json:"effectEstimate,omitempty"`

	// The effective period for a effect evidence synthesis determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// An individual or organization responsible for officially endorsing the content for use in some setting
	Endorser []common.ContactDetail `json:"endorser,omitempty"`

	// A reference to a EvidenceVariable resource that defines the exposure for the research
	Exposure common.Reference `json:"exposure"`

	// A reference to a EvidenceVariable resource that defines the comparison exposure for the research
	ExposureAlternative common.Reference `json:"exposureAlternative"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the effect evidence synthesis to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// A human-readable string to clarify or explain concepts about the resource
	Note []common.Annotation `json:"note,omitempty"`

	// A reference to a EvidenceVariable resource that defines the outcome for the research
	Outcome common.Reference `json:"outcome"`

	// A reference to a EvidenceVariable resource that defines the population for the research
	Population common.Reference `json:"population"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// Each related artifact is either an attachment, or a reference to another resource, but not both
	RelatedArtifact []common.RelatedArtifact `json:"relatedArtifact,omitempty"`

	// A description of the results for each exposure considered in the effect estimate
	ResultsByExposure []EffectEvidenceSynthesisResultsByExposure `json:"resultsByExposure,omitempty"`

	// An individual or organization primarily responsible for review of some aspect of the content
	Reviewer []common.ContactDetail `json:"reviewer,omitempty"`

	// A description of the size of the sample involved in the synthesis
	SampleSize *EffectEvidenceSynthesisSampleSize `json:"sampleSize,omitempty"`

	// Allows filtering of effect evidence synthesiss that are appropriate for use versus not
	Status EffectEvidenceSynthesisStatus `json:"status"`

	// Type of study eg randomized trial
	StudyType *common.CodeableConcept `json:"studyType,omitempty"`

	// Type of synthesis eg meta-analysis
	SynthesisType *common.CodeableConcept `json:"synthesisType,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Descriptive topics related to the content of the EffectEvidenceSynthesis
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different effect evidence synthesis instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// EffectEvidenceSynthesisStatus represents the status of the effect evidence synthesis
type EffectEvidenceSynthesisStatus string

const (
	EffectEvidenceSynthesisStatusDraft   EffectEvidenceSynthesisStatus = "draft"
	EffectEvidenceSynthesisStatusActive  EffectEvidenceSynthesisStatus = "active"
	EffectEvidenceSynthesisStatusRetired EffectEvidenceSynthesisStatus = "retired"
	EffectEvidenceSynthesisStatusUnknown EffectEvidenceSynthesisStatus = "unknown"
)

// EffectEvidenceSynthesisSampleSize represents a description of the size of the sample involved in the synthesis
type EffectEvidenceSynthesisSampleSize struct {
	common.BackboneElement

	// Human-readable summary of population estimate
	Description *string `json:"description,omitempty"`

	// Number of participants in the group
	NumberOfParticipants *int `json:"numberOfParticipants,omitempty"`

	// Number of studies included in this evidence synthesis
	NumberOfStudies *int `json:"numberOfStudies,omitempty"`

	// Known or expected number of participants known to have the exposure
	KnownDataCount *int `json:"knownDataCount,omitempty"`
}

// EffectEvidenceSynthesisResultsByExposure represents a description of the results for each exposure considered in the effect estimate
type EffectEvidenceSynthesisResultsByExposure struct {
	common.BackboneElement

	// Human-readable summary of results by exposure
	Description *string `json:"description,omitempty"`

	// Whether these results are for the exposure state or alternative exposure state
	ExposureState *EffectEvidenceSynthesisResultsByExposureExposureState `json:"exposureState,omitempty"`

	// Variant exposure states for the population
	VariantState *common.CodeableConcept `json:"variantState,omitempty"`

	// Risk evidence synthesis
	RiskEvidenceSynthesis common.Reference `json:"riskEvidenceSynthesis"`
}

// EffectEvidenceSynthesisResultsByExposureExposureState represents whether these results are for the exposure state or alternative exposure state
type EffectEvidenceSynthesisResultsByExposureExposureState string

const (
	EffectEvidenceSynthesisResultsByExposureExposureStateExposure            EffectEvidenceSynthesisResultsByExposureExposureState = "exposure"
	EffectEvidenceSynthesisResultsByExposureExposureStateExposureAlternative EffectEvidenceSynthesisResultsByExposureExposureState = "exposure-alternative"
)

// EffectEvidenceSynthesisEffectEstimate represents the estimated effect of the exposure variant
type EffectEvidenceSynthesisEffectEstimate struct {
	common.BackboneElement

	// Human-readable summary of effect estimate
	Description *string `json:"description,omitempty"`

	// Type of efffect estimate
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Variant exposure states for the population
	VariantState *common.CodeableConcept `json:"variantState,omitempty"`

	// Point estimate
	Value *float64 `json:"value,omitempty"`

	// What unit is the outcome described in
	UnitOfMeasure *common.CodeableConcept `json:"unitOfMeasure,omitempty"`

	// How precise the estimate is
	PrecisionEstimate []EffectEvidenceSynthesisEffectEstimatePrecisionEstimate `json:"precisionEstimate,omitempty"`
}

// EffectEvidenceSynthesisEffectEstimatePrecisionEstimate represents how precise the estimate is
type EffectEvidenceSynthesisEffectEstimatePrecisionEstimate struct {
	common.BackboneElement

	// Type of precision estimate
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Level of confidence interval
	Level *float64 `json:"level,omitempty"`

	// Lower bound
	From *float64 `json:"from,omitempty"`

	// Upper bound
	To *float64 `json:"to,omitempty"`
}

// EffectEvidenceSynthesisCertainty represents a description of the certainty of the effect estimate
type EffectEvidenceSynthesisCertainty struct {
	common.BackboneElement

	// Certainty rating
	Rating []common.CodeableConcept `json:"rating,omitempty"`

	// Used for footnotes or explanatory notes
	Note []common.Annotation `json:"note,omitempty"`

	// A component that contributes to the overall certainty
	CertaintySubcomponent []EffectEvidenceSynthesisCertaintyCertaintySubcomponent `json:"certaintySubcomponent,omitempty"`
}

// EffectEvidenceSynthesisCertaintyCertaintySubcomponent represents a component that contributes to the overall certainty
type EffectEvidenceSynthesisCertaintyCertaintySubcomponent struct {
	common.BackboneElement

	// Type of subcomponent of certainty rating
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Subcomponent certainty rating
	Rating []common.CodeableConcept `json:"rating,omitempty"`

	// Used for footnotes or explanatory notes
	Note []common.Annotation `json:"note,omitempty"`
}
