// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// GoalLifecycleStatus represents the lifecycle status of a goal
type GoalLifecycleStatus string

const (
	GoalLifecycleStatusProposed       GoalLifecycleStatus = "proposed"
	GoalLifecycleStatusPlanned        GoalLifecycleStatus = "planned"
	GoalLifecycleStatusAccepted       GoalLifecycleStatus = "accepted"
	GoalLifecycleStatusActive         GoalLifecycleStatus = "active"
	GoalLifecycleStatusOnHold         GoalLifecycleStatus = "on-hold"
	GoalLifecycleStatusCompleted      GoalLifecycleStatus = "completed"
	GoalLifecycleStatusCancelled      GoalLifecycleStatus = "cancelled"
	GoalLifecycleStatusEnteredInError GoalLifecycleStatus = "entered-in-error"
	GoalLifecycleStatusRejected       GoalLifecycleStatus = "rejected"
)

// GoalTarget represents when multiple targets are present for a single goal instance
type GoalTarget struct {
	common.BackboneElement

	// A CodeableConcept with just a text would be used instead of a string if the field was usually coded
	DetailQuantity        *common.Quantity        `json:"detailQuantity,omitempty"`
	DetailRange           *Range                  `json:"detailRange,omitempty"`
	DetailCodeableConcept *common.CodeableConcept `json:"detailCodeableConcept,omitempty"`
	DetailString          *string                 `json:"detailString,omitempty"`
	DetailBoolean         *bool                   `json:"detailBoolean,omitempty"`
	DetailInteger         *int                    `json:"detailInteger,omitempty"`
	DetailRatio           *Ratio                  `json:"detailRatio,omitempty"`

	// Indicates either the date or the duration after start by which the goal should be met
	DueDate     *string   `json:"dueDate,omitempty"`
	DueDuration *Duration `json:"dueDuration,omitempty"`

	// The parameter whose value is being tracked, e.g. body weight, blood pressure, or hemoglobin A1c level
	Measure *common.CodeableConcept `json:"measure,omitempty"`
}

// Goal represents the intended objective(s) for a patient, group or organization care
type Goal struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Goal"

	// Describes the progression, or lack thereof, towards the goal against the target
	AchievementStatus *common.CodeableConcept `json:"achievementStatus,omitempty"`

	// The identified conditions and other health record elements that are intended to be addressed by the goal
	Addresses []common.Reference `json:"addresses,omitempty"`

	// Indicates a category the goal falls within
	Category []common.CodeableConcept `json:"category,omitempty"`

	// For example, getting a yellow fever vaccination for a planned trip is a goal that is designed to be completed
	Continuous *bool `json:"continuous,omitempty"`

	// If no code is available, use CodeableConcept.text
	Description common.CodeableConcept `json:"description"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// This element is labeled as a modifier because the lifecycleStatus contains codes that mark the resource as not currently valid
	LifecycleStatus GoalLifecycleStatus `json:"lifecycleStatus"`

	// May be used for progress notes, concerns or other related information that doesn't actually describe the goal itself
	Note []Annotation `json:"note,omitempty"`

	// Note that this should not duplicate the goal status
	Outcome []CodeableReference `json:"outcome,omitempty"`

	// Extensions are available to track priorities as established by each participant
	Priority *common.CodeableConcept `json:"priority,omitempty"`

	// This is the individual or team responsible for establishing the goal
	Source *common.Reference `json:"source,omitempty"`

	// The date or event after which the goal should begin being pursued
	StartDate            *string                 `json:"startDate,omitempty"`
	StartCodeableConcept *common.CodeableConcept `json:"startCodeableConcept,omitempty"`

	// To see the date for past statuses, query history
	StatusDate *string `json:"statusDate,omitempty"`

	// This will typically be captured for statuses such as rejected, on-hold or cancelled
	StatusReason *string `json:"statusReason,omitempty"`

	// Identifies the patient, group or organization for whom the goal is being established
	Subject common.Reference `json:"subject"`

	// When multiple targets are present for a single goal instance, all targets must be met for the overall goal to be met
	Target []GoalTarget `json:"target,omitempty"`
}
