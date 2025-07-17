// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// CarePlanStatus represents the status of a care plan
type CarePlanStatus string

const (
	CarePlanStatusDraft          CarePlanStatus = "draft"
	CarePlanStatusActive         CarePlanStatus = "active"
	CarePlanStatusOnHold         CarePlanStatus = "on-hold"
	CarePlanStatusRevoked        CarePlanStatus = "revoked"
	CarePlanStatusCompleted      CarePlanStatus = "completed"
	CarePlanStatusEnteredInError CarePlanStatus = "entered-in-error"
	CarePlanStatusUnknown        CarePlanStatus = "unknown"
)

// CarePlanIntent represents the intention/purpose of a care plan
type CarePlanIntent string

const (
	CarePlanIntentProposal  CarePlanIntent = "proposal"
	CarePlanIntentPlan      CarePlanIntent = "plan"
	CarePlanIntentOrder     CarePlanIntent = "order"
	CarePlanIntentOption    CarePlanIntent = "option"
	CarePlanIntentDirective CarePlanIntent = "directive"
)

// CarePlanActivity represents an action that has occurred or is a planned action to occur as part of the plan
type CarePlanActivity struct {
	common.BackboneElement

	// Note that this should not duplicate the activity status (e.g. completed or in progress)
	PerformedActivity []CodeableReference `json:"performedActivity,omitempty"`

	// The goal should be visible when the resource referenced by CarePlan.activity.plannedActivityReference is viewed independently
	PlannedActivityReference *common.Reference `json:"plannedActivityReference,omitempty"`

	// This element should NOT be used to describe the activity to be performed
	Progress []Annotation `json:"progress,omitempty"`
}

// CarePlan represents the intention of how one or more practitioners intend to deliver care for a particular patient, group or community
type CarePlan struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "CarePlan"

	// Identifies an action that has occurred or is a planned action to occur as part of the plan
	Activity []CarePlanActivity `json:"activity,omitempty"`

	// Use CarePlan.addresses.concept when a code sufficiently describes the concern
	Addresses []CodeableReference `json:"addresses,omitempty"`

	// A higher-level request resource (i.e. a plan, proposal or order) that is fulfilled in whole or in part by this care plan
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Identifies all people and organizations who are expected to be involved in the care envisioned by this plan
	CareTeam []common.Reference `json:"careTeam,omitempty"`

	// There may be multiple axes of categorization and one plan may serve multiple purposes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Collaborative care plans may have multiple contributors
	Contributor []common.Reference `json:"contributor,omitempty"`

	// Represents when this particular CarePlan record was created in the system
	Created *string `json:"created,omitempty"`

	// The custodian might or might not be a contributor
	Custodian *common.Reference `json:"custodian,omitempty"`

	// A description of the scope and nature of the plan
	Description *string `json:"description,omitempty"`

	// This will typically be the encounter the event occurred within
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Goal can be achieving a particular change or merely maintaining a current state or even slowing a decline
	Goal []common.Reference `json:"goal,omitempty"`

	// Business identifier for this care plan
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The URL pointing to a FHIR-defined protocol, guideline, questionnaire or other definition that is adhered to in whole or in part
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// This might be an HTML page, PDF, etc. or could just be a non-resolvable URI identifier
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// This element is labeled as a modifier because the intent alters when and how the resource is actually applicable
	Intent CarePlanIntent `json:"intent"`

	// General notes about the care plan not covered elsewhere
	Note []Annotation `json:"note,omitempty"`

	// Each care plan is an independent request, such that having a care plan be part of another care plan can cause issues
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Any activities scheduled as part of the plan should be constrained to the specified period
	Period *common.Period `json:"period,omitempty"`

	// The replacement could be because the initial care plan was immediately rejected or because the previous care plan was completed
	Replaces []common.Reference `json:"replaces,omitempty"`

	// The unknown code is not to be used to convey other statuses
	Status CarePlanStatus `json:"status"`

	// Identifies the patient or group whose intended care is described by the plan
	Subject common.Reference `json:"subject"`

	// Use "concern" to identify specific conditions addressed by the care plan
	SupportingInfo []common.Reference `json:"supportingInfo,omitempty"`

	// Human-friendly name for the care plan
	Title *string `json:"title,omitempty"`
}
