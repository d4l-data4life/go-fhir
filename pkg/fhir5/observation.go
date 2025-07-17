// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ObservationStatus represents the status of an observation
type ObservationStatus string

const (
	ObservationStatusRegistered     ObservationStatus = "registered"
	ObservationStatusPreliminary    ObservationStatus = "preliminary"
	ObservationStatusFinal          ObservationStatus = "final"
	ObservationStatusAmended        ObservationStatus = "amended"
	ObservationStatusCorrected      ObservationStatus = "corrected"
	ObservationStatusCancelled      ObservationStatus = "cancelled"
	ObservationStatusEnteredInError ObservationStatus = "entered-in-error"
	ObservationStatusUnknown        ObservationStatus = "unknown"
)

// ObservationTriggeredByType represents the type of trigger for an observation
type ObservationTriggeredByType string

const (
	ObservationTriggeredByTypeReflex ObservationTriggeredByType = "reflex"
	ObservationTriggeredByTypeRepeat ObservationTriggeredByType = "repeat"
	ObservationTriggeredByTypeReRun  ObservationTriggeredByType = "re-run"
)

// ObservationTriggeredBy represents observations that triggered this observation
type ObservationTriggeredBy struct {
	common.BackboneElement

	// Reference to the triggering observation
	Observation common.Reference `json:"observation"`

	// Provides the reason why this observation was performed as a result of the observation(s) referenced
	Reason *string `json:"reason,omitempty"`

	// The type of trigger
	Type ObservationTriggeredByType `json:"type"`
}

// ObservationReferenceRange represents reference range for observation values
type ObservationReferenceRange struct {
	common.BackboneElement

	// The age at which this reference range is applicable
	Age *Range `json:"age,omitempty"`

	// This SHOULD be populated if there is more than one range
	AppliesTo []common.CodeableConcept `json:"appliesTo,omitempty"`

	// The value of the high bound of the reference range
	High *common.Quantity `json:"high,omitempty"`

	// The value of the low bound of the reference range
	Low *common.Quantity `json:"low,omitempty"`

	// The value of the normal value of the reference range
	NormalValue *common.CodeableConcept `json:"normalValue,omitempty"`

	// Text based reference range in an observation
	Text *string `json:"text,omitempty"`

	// This SHOULD be populated if there is more than one range
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ObservationComponent represents component observations
type ObservationComponent struct {
	common.BackboneElement

	// Describes what was observed
	Code common.CodeableConcept `json:"code"`

	// For exceptional values in measurements
	DataAbsentReason *common.CodeableConcept `json:"dataAbsentReason,omitempty"`

	// Historically used for laboratory results (known as 'abnormal flag')
	Interpretation []common.CodeableConcept `json:"interpretation,omitempty"`

	// Reference ranges for the component
	ReferenceRange []ObservationReferenceRange `json:"referenceRange,omitempty"`

	// Component observation values - choice of types
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueInteger         *int                    `json:"valueInteger,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueRatio           *Ratio                  `json:"valueRatio,omitempty"`
	ValueSampledData     *common.SampledData     `json:"valueSampledData,omitempty"`
	ValueTime            *string                 `json:"valueTime,omitempty"`
	ValueDateTime        *string                 `json:"valueDateTime,omitempty"`
	ValuePeriod          *common.Period          `json:"valuePeriod,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
}

// Observation represents measurements and simple assertions made about a patient, device or other subject
type Observation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Observation"

	// A plan, proposal or order that is fulfilled in whole or in part by this event
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Only used if not implicit in code found in Observation.code
	BodySite *common.CodeableConcept `json:"bodySite,omitempty"`

	// Only used if not implicit in code found in Observation.code or bodySite is used
	BodyStructure *common.Reference `json:"bodyStructure,omitempty"`

	// In addition to the required category valueset, this element allows various categorization schemes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Describes what was observed
	Code common.CodeableConcept `json:"code"`

	// For a discussion on the ways Observations can be assembled in groups together
	Component []ObservationComponent `json:"component,omitempty"`

	// Null or exceptional values can be represented two ways in FHIR Observations
	DataAbsentReason *common.CodeableConcept `json:"dataAbsentReason,omitempty"`

	// All the reference choices that are listed in this element can represent clinical observations
	DerivedFrom []common.Reference `json:"derivedFrom,omitempty"`

	// Note that this is not meant to represent a device involved in the transmission of the result
	Device *common.Reference `json:"device,omitempty"`

	// At least a date should be present unless this observation is a historical report
	EffectiveDateTime *string        `json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *common.Period `json:"effectivePeriod,omitempty"`
	EffectiveTiming   *Timing        `json:"effectiveTiming,omitempty"`
	EffectiveInstant  *string        `json:"effectiveInstant,omitempty"`

	// This will typically be the encounter the event occurred within
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Typically, an observation is made about the subject
	Focus []common.Reference `json:"focus,omitempty"`

	// When using this element, an observation will typically have either a value or a set of related resources
	HasMember []common.Reference `json:"hasMember,omitempty"`

	// A unique identifier assigned to this observation
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// ObservationDefinition can be referenced by its canonical url using instantiatesCanonical
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`

	// ObservationDefinition can be referenced by its canonical url using instantiatesCanonical
	InstantiatesReference *common.Reference `json:"instantiatesReference,omitempty"`

	// Historically used for laboratory results (known as 'abnormal flag')
	Interpretation []common.CodeableConcept `json:"interpretation,omitempty"`

	// For Observations that don't require review and verification
	Issued *string `json:"issued,omitempty"`

	// Only used if not implicit in code for Observation.code
	Method *common.CodeableConcept `json:"method,omitempty"`

	// May include general statements about the observation
	Note []Annotation `json:"note,omitempty"`

	// To link an Observation to an Encounter use `encounter`
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Who was responsible for asserting the observed value as "true"
	Performer []common.Reference `json:"performer,omitempty"`

	// Most observations only have one generic reference range
	ReferenceRange []ObservationReferenceRange `json:"referenceRange,omitempty"`

	// Should only be used if not implicit in code found in `Observation.code`
	Specimen *common.Reference `json:"specimen,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status ObservationStatus `json:"status"`

	// One would expect this element to be a cardinality of 1..1
	Subject *common.Reference `json:"subject,omitempty"`

	// Identifies the observation(s) that triggered the performance of this observation
	TriggeredBy []ObservationTriggeredBy `json:"triggeredBy,omitempty"`

	// An observation may have a single value here, both a value and a set of related or component values, or only a set of related or component values
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueInteger         *int                    `json:"valueInteger,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueRatio           *Ratio                  `json:"valueRatio,omitempty"`
	ValueSampledData     *common.SampledData     `json:"valueSampledData,omitempty"`
	ValueTime            *string                 `json:"valueTime,omitempty"`
	ValueDateTime        *string                 `json:"valueDateTime,omitempty"`
	ValuePeriod          *common.Period          `json:"valuePeriod,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
}
