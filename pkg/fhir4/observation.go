// Package fhir4 contains FHIR R4 (version 4.0.1) resource definitions
package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Observation represents measurements and simple assertions made about a patient
type Observation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Observation"

	// Business Identifier for observation
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Fulfills plan, proposal or order
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Part of referenced event
	PartOf []common.Reference `json:"partOf,omitempty"`

	// registered | preliminary | final | amended | corrected | cancelled | entered-in-error | unknown
	Status ObservationStatus `json:"status"`

	// Classification of type of observation
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Type of observation (code / type)
	Code common.CodeableConcept `json:"code"`

	// Who and/or what the observation is about
	Subject *common.Reference `json:"subject,omitempty"`

	// What the observation is about, when it is not about the subject of record
	Focus []common.Reference `json:"focus,omitempty"`

	// Healthcare event during which this observation is made
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Clinically relevant time/time-period for observation
	EffectiveDateTime *string        `json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *common.Period `json:"effectivePeriod,omitempty"`
	EffectiveTiming   *Timing        `json:"effectiveTiming,omitempty"`
	EffectiveInstant  *string        `json:"effectiveInstant,omitempty"`

	// Date/Time this version was made available
	Issued *string `json:"issued,omitempty"`

	// Who is responsible for the observation
	Performer []common.Reference `json:"performer,omitempty"`

	// Actual result
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueInteger         *int                    `json:"valueInteger,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueRatio           *Ratio                  `json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData            `json:"valueSampledData,omitempty"`
	ValueTime            *string                 `json:"valueTime,omitempty"`
	ValueDateTime        *string                 `json:"valueDateTime,omitempty"`
	ValuePeriod          *common.Period          `json:"valuePeriod,omitempty"`

	// Why the result is missing
	DataAbsentReason *common.CodeableConcept `json:"dataAbsentReason,omitempty"`

	// High, low, normal, etc.
	Interpretation []common.CodeableConcept `json:"interpretation,omitempty"`

	// Comments about the observation
	Note []Annotation `json:"note,omitempty"`

	// Observed body part
	BodySite *common.CodeableConcept `json:"bodySite,omitempty"`

	// How it was done
	Method *common.CodeableConcept `json:"method,omitempty"`

	// Specimen used for this observation
	Specimen *common.Reference `json:"specimen,omitempty"`

	// (Measurement) Device
	Device *common.Reference `json:"device,omitempty"`

	// Provides guide for interpretation
	ReferenceRange []ObservationReferenceRange `json:"referenceRange,omitempty"`

	// Related resource that belongs to the Observation group
	HasMember []common.Reference `json:"hasMember,omitempty"`

	// Related measurements the observation is made from
	DerivedFrom []common.Reference `json:"derivedFrom,omitempty"`

	// Component results
	Component []ObservationComponent `json:"component,omitempty"`
}

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

// ObservationReferenceRange represents reference range for an observation
type ObservationReferenceRange struct {
	common.BackboneElement

	// Low Range, if relevant
	Low *common.Quantity `json:"low,omitempty"`

	// High Range, if relevant
	High *common.Quantity `json:"high,omitempty"`

	// Reference range qualifier
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Reference range population
	AppliesTo []common.CodeableConcept `json:"appliesTo,omitempty"`

	// Applicable age range, if relevant
	Age *Range `json:"age,omitempty"`

	// Text based reference range in an observation
	Text *string `json:"text,omitempty"`
}

// ObservationComponent represents a component observation
type ObservationComponent struct {
	common.BackboneElement

	// Type of component observation (code / type)
	Code common.CodeableConcept `json:"code"`

	// Actual component result
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueInteger         *int                    `json:"valueInteger,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueRatio           *Ratio                  `json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData            `json:"valueSampledData,omitempty"`
	ValueTime            *string                 `json:"valueTime,omitempty"`
	ValueDateTime        *string                 `json:"valueDateTime,omitempty"`
	ValuePeriod          *common.Period          `json:"valuePeriod,omitempty"`

	// Why the component result is missing
	DataAbsentReason *common.CodeableConcept `json:"dataAbsentReason,omitempty"`

	// High, low, normal, etc.
	Interpretation []common.CodeableConcept `json:"interpretation,omitempty"`

	// Provides guide for interpretation of component result
	ReferenceRange []ObservationReferenceRange `json:"referenceRange,omitempty"`
}
