package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Group represents a defined collection of entities that may be discussed or acted upon collectively
type Group struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Group"

	// Indicates whether the record for the group is available for use or is merely being retained for historical purposes
	Active *bool `json:"active,omitempty"`

	// If true, indicates that the resource refers to a specific group of real individuals
	Actual bool `json:"actual"`

	// All the identified characteristics must be true for an entity to a member of the group
	Characteristic []GroupCharacteristic `json:"characteristic,omitempty"`

	// This would generally be omitted for Person resources
	Code *common.CodeableConcept `json:"code,omitempty"`

	// A unique business identifier for this group
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// This does not strictly align with ownership of a herd or flock, but may suffice to represent that relationship in simple cases
	ManagingEntity *common.Reference `json:"managingEntity,omitempty"`

	// Identifies the resource instances that are members of the group
	Member []GroupMember `json:"member,omitempty"`

	// A label assigned to the group for human identification and communication
	Name *string `json:"name,omitempty"`

	// Note that the quantity may be less than the number of members if some of the members are not active
	Quantity *int `json:"quantity,omitempty"`

	// Group members SHALL be of the appropriate resource type
	Type GroupType `json:"type"`
}

// GroupType represents the type of group
type GroupType string

const (
	GroupTypePerson       GroupType = "person"
	GroupTypeAnimal       GroupType = "animal"
	GroupTypePractitioner GroupType = "practitioner"
	GroupTypeDevice       GroupType = "device"
	GroupTypeMedication   GroupType = "medication"
	GroupTypeSubstance    GroupType = "substance"
)

// GroupCharacteristic represents all the identified characteristics must be true for an entity to a member of the group
type GroupCharacteristic struct {
	common.BackboneElement

	// A code that identifies the kind of trait being asserted
	Code common.CodeableConcept `json:"code"`

	// This is labeled as "Is Modifier" because applications cannot wrongly include excluded members as included or vice versa
	Exclude bool `json:"exclude"`

	// The period over which the characteristic is tested; e.g. the patient had an operation during the month of June
	Period *common.Period `json:"period,omitempty"`

	// For Range, it means members of the group have a value that falls somewhere within the specified range
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`

	// For Range, it means members of the group have a value that falls somewhere within the specified range
	ValueBoolean *bool `json:"valueBoolean,omitempty"`

	// For Range, it means members of the group have a value that falls somewhere within the specified range
	ValueQuantity *common.Quantity `json:"valueQuantity,omitempty"`

	// For Range, it means members of the group have a value that falls somewhere within the specified range
	ValueRange *common.Range `json:"valueRange,omitempty"`

	// For Range, it means members of the group have a value that falls somewhere within the specified range
	ValueReference *common.Reference `json:"valueReference,omitempty"`
}

// GroupMember represents identifies the resource instances that are members of the group
type GroupMember struct {
	common.BackboneElement

	// A reference to the entity that is a member of the group
	Entity common.Reference `json:"entity"`

	// A flag to indicate that the member is no longer in the group, but previously may have been a member
	Inactive *bool `json:"inactive,omitempty"`

	// The period that the member was in the group, if known
	Period *common.Period `json:"period,omitempty"`
}
