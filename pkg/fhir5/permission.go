package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// PermissionStatus represents the status of a permission
type PermissionStatus string

const (
	PermissionStatusActive         PermissionStatus = "active"
	PermissionStatusEnteredInError PermissionStatus = "entered-in-error"
	PermissionStatusInactive       PermissionStatus = "inactive"
)

// PermissionJustification represents the regulatory grounds for the permission
type PermissionJustification struct {
	common.BackboneElement

	// The regulatory grounds for the permission
	Basis []common.CodeableConcept `json:"basis,omitempty"`

	// Evidence for the permission
	Evidence []common.Reference `json:"evidence,omitempty"`
}

// PermissionSubjectCharacteristic represents characteristics of the subject
type PermissionSubjectCharacteristic struct {
	common.BackboneElement

	// Whether the characteristic is excluded or included
	Exclude *bool `json:"exclude,omitempty"`

	// The period during which the characteristic is valid
	Period *common.Period `json:"period,omitempty"`

	// The type of characteristic
	Type common.CodeableConcept `json:"type"`

	// The value of the characteristic
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueExpression      *Expression             `json:"valueExpression,omitempty"`
}

// PermissionSubject represents the subject of the permission
type PermissionSubject struct {
	common.BackboneElement

	// Characteristics of the subject
	Characteristic []PermissionSubjectCharacteristic `json:"characteristic,omitempty"`

	// References to the subject
	Reference []common.Reference `json:"reference,omitempty"`
}

// Permission represents a permission
type Permission struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Permission"

	// The person or entity that asserts the permission
	Asserter *common.Reference `json:"asserter,omitempty"`

	// The date that permission was asserted
	Date []*string `json:"date,omitempty"`

	// The regulatory grounds for the permission
	Justification *PermissionJustification `json:"justification,omitempty"`

	// The person or entity that is the custodian of the permission
	ManagingEntity *common.Reference `json:"managingEntity,omitempty"`

	// A note that describes or explains the permission
	Note []Annotation `json:"note,omitempty"`

	// The status of the permission
	Status PermissionStatus `json:"status"`

	// The subject of the permission
	Subject []PermissionSubject `json:"subject,omitempty"`

	// The period of time during which the permission is granted
	ValidityPeriod *common.Period `json:"validityPeriod,omitempty"`
}
