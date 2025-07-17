// Package fhir4 contains FHIR R4 (version 4.0.1) resource definitions
package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Immunization represents the event of a patient being administered a vaccine or a record of an immunization as reported by a patient, a clinician or another party
type Immunization struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Immunization"

	// Business identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// completed | entered-in-error | not-done
	Status ImmunizationStatus `json:"status"`

	// Reason not done
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// Vaccine product administered
	VaccineCode common.CodeableConcept `json:"vaccineCode"`

	// Who was immunized
	Patient common.Reference `json:"patient"`

	// Encounter immunization was part of
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Vaccine administration date
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`
	OccurrenceString   *string `json:"occurrenceString,omitempty"`

	// When the immunization was first captured in the subject's record
	Recorded *string `json:"recorded,omitempty"`

	// Indicates if record is from primary source
	PrimarySource *bool `json:"primarySource,omitempty"`

	// Indicates the source of the vaccine actually administered
	ReportOrigin *common.CodeableConcept `json:"reportOrigin,omitempty"`

	// Where immunization occurred
	Location *common.Reference `json:"location,omitempty"`

	// Vaccine manufacturer
	Manufacturer *common.Reference `json:"manufacturer,omitempty"`

	// Vaccine lot number
	LotNumber *string `json:"lotNumber,omitempty"`

	// Vaccine expiration date
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// Body site vaccine was administered
	Site *common.CodeableConcept `json:"site,omitempty"`

	// How vaccine entered body
	Route *common.CodeableConcept `json:"route,omitempty"`

	// Amount of vaccine administered
	DoseQuantity *common.Quantity `json:"doseQuantity,omitempty"`

	// Who performed event
	Performer []ImmunizationPerformer `json:"performer,omitempty"`

	// Additional immunization notes
	Note []Annotation `json:"note,omitempty"`

	// Why immunization occurred
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Why immunization occurred
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// Dose potency
	IsSubpotent *bool `json:"isSubpotent,omitempty"`

	// Reason for being subpotent
	SubpotentReason []common.CodeableConcept `json:"subpotentReason,omitempty"`

	// Educational material presented to patient
	Education []ImmunizationEducation `json:"education,omitempty"`

	// Patient eligibility for a vaccination program
	ProgramEligibility []common.CodeableConcept `json:"programEligibility,omitempty"`

	// Funding source for the vaccine
	FundingSource *common.CodeableConcept `json:"fundingSource,omitempty"`

	// Details of a reaction that follows immunization
	Reaction []ImmunizationReaction `json:"reaction,omitempty"`

	// Protocol followed by the provider
	ProtocolApplied []ImmunizationProtocolApplied `json:"protocolApplied,omitempty"`
}

// ImmunizationStatus represents the status of the immunization
type ImmunizationStatus string

const (
	ImmunizationStatusCompleted      ImmunizationStatus = "completed"
	ImmunizationStatusEnteredInError ImmunizationStatus = "entered-in-error"
	ImmunizationStatusNotDone        ImmunizationStatus = "not-done"
)

// ImmunizationPerformer represents who performed the immunization
type ImmunizationPerformer struct {
	common.BackboneElement

	// What type of performance was done
	Function *common.CodeableConcept `json:"function,omitempty"`

	// Individual or organization who was performing
	Actor common.Reference `json:"actor"`
}

// ImmunizationEducation represents educational material presented to patient
type ImmunizationEducation struct {
	common.BackboneElement

	// Educational material document identifier
	DocumentType *string `json:"documentType,omitempty"`

	// Educational material reference pointer
	Reference *string `json:"reference,omitempty"`

	// Educational material publication date
	PublicationDate *string `json:"publicationDate,omitempty"`

	// Educational material presentation date
	PresentationDate *string `json:"presentationDate,omitempty"`
}

// ImmunizationReaction represents a reaction that follows immunization
type ImmunizationReaction struct {
	common.BackboneElement

	// When reaction started
	Date *string `json:"date,omitempty"`

	// Additional information on reaction
	Detail *common.Reference `json:"detail,omitempty"`

	// Indicates self-reported reaction
	Reported *bool `json:"reported,omitempty"`
}

// ImmunizationProtocolApplied represents protocol followed by the provider
type ImmunizationProtocolApplied struct {
	common.BackboneElement

	// Name of vaccine series
	Series *string `json:"series,omitempty"`

	// Who is responsible for publishing the recommendations
	Authority *common.Reference `json:"authority,omitempty"`

	// Vaccine preventatable disease being targeted
	TargetDisease []common.CodeableConcept `json:"targetDisease,omitempty"`

	// Dose number within series
	DoseNumberPositiveInt *int    `json:"doseNumberPositiveInt,omitempty"`
	DoseNumberString      *string `json:"doseNumberString,omitempty"`

	// Recommended number of doses for immunity
	SeriesDosesPositiveInt *int    `json:"seriesDosesPositiveInt,omitempty"`
	SeriesDosesString      *string `json:"seriesDosesString,omitempty"`
}
