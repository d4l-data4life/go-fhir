package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// DeviceAssociationOperation represents the details about the device when it is in use
type DeviceAssociationOperation struct {
	common.BackboneElement

	// The individual performing the action enabled by the device
	Operator []common.Reference `json:"operator,omitempty"`

	// Begin and end dates and times for the device's operation
	Period *common.Period `json:"period,omitempty"`

	// Device operational condition corresponding to the association
	Status common.CodeableConcept `json:"status"`
}

// DeviceAssociation represents a record of association or dissociation of a device with a patient
type DeviceAssociation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DeviceAssociation"

	// Current anatomical location of the device in/on subject
	BodyStructure *common.Reference `json:"bodyStructure,omitempty"`

	// Describes the relationship between the device and subject
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Reference to the devices associated with the patient or group
	Device common.Reference `json:"device"`

	// Instance identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The details about the device when it is in use to describe its operation
	Operation []DeviceAssociationOperation `json:"operation,omitempty"`

	// Begin and end dates and times for the device association
	Period *common.Period `json:"period,omitempty"`

	// Indicates the state of the Device association
	Status common.CodeableConcept `json:"status"`

	// The reasons given for the current association status
	StatusReason []common.CodeableConcept `json:"statusReason,omitempty"`

	// The individual, group of individuals or device that the device is on or associated with
	Subject *common.Reference `json:"subject,omitempty"`
}
