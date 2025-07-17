package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// DeviceUsageAdherence represents how or if the device is being used
type DeviceUsageAdherence struct {
	common.BackboneElement

	// The reason for the adherence
	Reason common.CodeableConcept `json:"reason"`
}

// DeviceUsage represents a record of a device being used by a patient
type DeviceUsage struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DeviceUsage"

	// This indicates how or if the device is being used
	Adherence *DeviceUsageAdherence `json:"adherence,omitempty"`

	// A plan, proposal or order that is fulfilled in whole or in part by this DeviceUsage
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Indicates the anatomic location on the subject's body where the device was used
	BodySite *CodeableReference `json:"bodySite,omitempty"`

	// This attribute indicates a category for the statement
	Category []common.CodeableConcept `json:"category,omitempty"`

	// The encounter or episode of care that establishes the context for this device use statement
	Context *common.Reference `json:"context,omitempty"`

	// The time at which the statement was recorded by informationSource
	DateAsserted *string `json:"dateAsserted,omitempty"`

	// The most common use cases for deriving a DeviceUsage comes from creating it from a request or from an observation or a claim
	DerivedFrom []common.Reference `json:"derivedFrom,omitempty"`

	// Code or Reference to device used
	Device CodeableReference `json:"device"`

	// An external identifier for this statement such as an IRI
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Who reported the device was being used by the patient
	InformationSource *common.Reference `json:"informationSource,omitempty"`

	// Details about the device statement that were not represented at all or sufficiently in one of the attributes
	Note []Annotation `json:"note,omitempty"`

	// The patient who used the device
	Patient common.Reference `json:"patient"`

	// When the status is not done, the reason code indicates why it was not done
	Reason []CodeableReference `json:"reason,omitempty"`

	// DeviceUseStatement is a statement at a point in time
	Status DeviceUsageStatus `json:"status"`

	// How often the device was used
	TimingTiming   *Timing        `json:"timingTiming,omitempty"`
	TimingPeriod   *common.Period `json:"timingPeriod,omitempty"`
	TimingDateTime *string        `json:"timingDateTime,omitempty"`

	// The reason for asserting the usage status
	UsageReason []common.CodeableConcept `json:"usageReason,omitempty"`

	// The status of the device usage, for example always, sometimes, never
	UsageStatus *common.CodeableConcept `json:"usageStatus,omitempty"`
}

// DeviceUsageStatus represents the status of the device usage statement
type DeviceUsageStatus string

const (
	DeviceUsageStatusActive         DeviceUsageStatus = "active"
	DeviceUsageStatusCompleted      DeviceUsageStatus = "completed"
	DeviceUsageStatusNotDone        DeviceUsageStatus = "not-done"
	DeviceUsageStatusEnteredInError DeviceUsageStatus = "entered-in-error"
	DeviceUsageStatusIntended       DeviceUsageStatus = "intended"
	DeviceUsageStatusStopped        DeviceUsageStatus = "stopped"
	DeviceUsageStatusOnHold         DeviceUsageStatus = "on-hold"
)
