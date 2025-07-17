package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ImmunizationStatus represents the status of an immunization
type ImmunizationStatus string

const (
	ImmunizationStatusCompleted      ImmunizationStatus = "completed"
	ImmunizationStatusEnteredInError ImmunizationStatus = "entered-in-error"
	ImmunizationStatusNotDone        ImmunizationStatus = "not-done"
)

// ImmunizationPerformer represents who performed the immunization event
type ImmunizationPerformer struct {
	common.BackboneElement

	// When the individual practitioner who performed the action is known, it is best to send
	Actor common.Reference `json:"actor"`

	// Describes the type of performance (e.g. ordering provider, administering provider, etc.)
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// ImmunizationProgramEligibility represents a patient's eligibility for a funding program
type ImmunizationProgramEligibility struct {
	common.BackboneElement

	// Indicates which program the patient had their eligility evaluated for
	Program common.CodeableConcept `json:"program"`

	// Indicates the patient's eligility status for for a specific payment program
	ProgramStatus common.CodeableConcept `json:"programStatus"`
}

// ImmunizationReaction represents a reaction to the immunization
type ImmunizationReaction struct {
	common.BackboneElement

	// Date of reaction to the immunization
	Date *string `json:"date,omitempty"`

	// Details of the reaction
	Manifestation *CodeableReference `json:"manifestation,omitempty"`

	// Self-reported indicator
	Reported *bool `json:"reported,omitempty"`
}

// ImmunizationProtocolApplied represents the protocol being followed by the provider
type ImmunizationProtocolApplied struct {
	common.BackboneElement

	// Indicates the authority who published the protocol (e.g. ACIP) that is being followed
	Authority *common.Reference `json:"authority,omitempty"`

	// The use of an integer is preferred if known
	DoseNumber string `json:"doseNumber"`

	// One possible path to achieve presumed immunity against a disease
	Series *string `json:"series,omitempty"`

	// The use of an integer is preferred if known
	SeriesDoses *string `json:"seriesDoses,omitempty"`

	// The vaccine preventable disease the dose is being administered against
	TargetDisease []common.CodeableConcept `json:"targetDisease,omitempty"`
}

// Immunization represents the event of a patient being administered a vaccine
type Immunization struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Immunization"

	// An indication of which product was administered to the patient
	AdministeredProduct *CodeableReference `json:"administeredProduct,omitempty"`

	// Allows tracing of an authorization for the Immunization
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The quantity of vaccine product that was administered
	DoseQuantity *common.Quantity `json:"doseQuantity,omitempty"`

	// The visit or admission or other contact between patient and health care provider
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Date vaccine batch expires
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// Indicates the source of the vaccine actually administered
	FundingSource *common.CodeableConcept `json:"fundingSource,omitempty"`

	// A unique identifier assigned to this immunization record
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Typically will not be populated if primarySource = True
	InformationSource *CodeableReference `json:"informationSource,omitempty"`

	// Typically, the recognition of the dose being sub-potent is retrospective
	IsSubpotent *bool `json:"isSubpotent,omitempty"`

	// The service delivery location where the vaccine administration occurred
	Location *common.Reference `json:"location,omitempty"`

	// Lot number of the vaccine product
	LotNumber *string `json:"lotNumber,omitempty"`

	// Name of vaccine manufacturer
	Manufacturer *CodeableReference `json:"manufacturer,omitempty"`

	// Extra information about the immunization
	Note []Annotation `json:"note,omitempty"`

	// When immunizations are given a specific date and time should always be known
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`
	OccurrenceString   *string `json:"occurrenceString,omitempty"`

	// The patient who either received or did not receive the immunization
	Patient common.Reference `json:"patient"`

	// Indicates who performed the immunization event
	Performer []ImmunizationPerformer `json:"performer,omitempty"`

	// Reflects the "reliability" of the content
	PrimarySource *bool `json:"primarySource,omitempty"`

	// Indicates a patient's eligibility for a funding program
	ProgramEligibility []ImmunizationProgramEligibility `json:"programEligibility,omitempty"`

	// The protocol (set of recommendations) being followed by the provider
	ProtocolApplied []ImmunizationProtocolApplied `json:"protocolApplied,omitempty"`

	// A reaction may be an indication of an allergy or intolerance
	Reaction []ImmunizationReaction `json:"reaction,omitempty"`

	// Describes why the immunization occurred in coded or textual form
	Reason []CodeableReference `json:"reason,omitempty"`

	// The path by which the vaccine product is taken into the body
	Route *common.CodeableConcept `json:"route,omitempty"`

	// Body site where vaccine was administered
	Site *common.CodeableConcept `json:"site,omitempty"`

	// Will generally be set to show that the immunization has been completed or not done
	Status ImmunizationStatus `json:"status"`

	// This is generally only used for the status of "not-done"
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// Reason why a dose is considered to be subpotent
	SubpotentReason []common.CodeableConcept `json:"subpotentReason,omitempty"`

	// Additional information that is relevant to the immunization
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`

	// The code for the administered vaccine
	VaccineCode common.CodeableConcept `json:"vaccineCode"`
}
