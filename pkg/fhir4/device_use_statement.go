package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// DeviceUseStatement represents a record of a device being used by a patient
type DeviceUseStatement struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DeviceUseStatement"

	// A plan, proposal or order that is fulfilled in whole or in part by this DeviceUseStatement
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Indicates the anatomic location on the subject's body where the device was used (i.e. the target)
	BodySite *common.CodeableConcept `json:"bodySite,omitempty"`

	// The most common use cases for deriving a DeviceUseStatement comes from creating it from a request or from an observation or a claim
	DerivedFrom []common.Reference `json:"derivedFrom,omitempty"`

	// The details of the device used
	Device common.Reference `json:"device"`

	// An external identifier for this statement such as an IRI
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Details about the device statement that were not represented at all or sufficiently in one of the attributes provided in a class
	Note []common.Annotation `json:"note,omitempty"`

	// Reason or justification for the use of the device
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Indicates another resource whose existence justifies this DeviceUseStatement
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// The time at which the statement was made/recorded
	RecordedOn *string `json:"recordedOn,omitempty"`

	// Who reported the device was being used by the patient
	Source *common.Reference `json:"source,omitempty"`

	// DeviceUseStatement is a statement at a point in time
	Status DeviceUseStatementStatus `json:"status"`

	// The patient who used the device
	Subject common.Reference `json:"subject"`

	// How often the device was used
	TimingTiming *common.Timing `json:"timingTiming,omitempty"`

	// How often the device was used
	TimingPeriod *common.Period `json:"timingPeriod,omitempty"`

	// How often the device was used
	TimingDateTime *string `json:"timingDateTime,omitempty"`
}

// DeviceUseStatementStatus represents the status of the device use statement
type DeviceUseStatementStatus string

const (
	DeviceUseStatementStatusActive         DeviceUseStatementStatus = "active"
	DeviceUseStatementStatusCompleted      DeviceUseStatementStatus = "completed"
	DeviceUseStatementStatusEnteredInError DeviceUseStatementStatus = "entered-in-error"
	DeviceUseStatementStatusIntended       DeviceUseStatementStatus = "intended"
	DeviceUseStatementStatusStopped        DeviceUseStatementStatus = "stopped"
	DeviceUseStatementStatusOnHold         DeviceUseStatementStatus = "on-hold"
)
