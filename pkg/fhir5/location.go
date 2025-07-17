// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// LocationStatus represents the status of a location
type LocationStatus string

const (
	LocationStatusActive    LocationStatus = "active"
	LocationStatusSuspended LocationStatus = "suspended"
	LocationStatusInactive  LocationStatus = "inactive"
)

// LocationMode represents the mode of a location
type LocationMode string

const (
	LocationModeInstance LocationMode = "instance"
	LocationModeKind     LocationMode = "kind"
)

// LocationPosition represents the position of a location
type LocationPosition struct {
	common.BackboneElement

	// Altitude with the same coordinate system as KML
	Altitude *float64 `json:"altitude,omitempty"`

	// Latitude with the same coordinate system as KML
	Latitude float64 `json:"latitude"`

	// Longitude with the same coordinate system as KML
	Longitude float64 `json:"longitude"`
}

// Location represents a place where services are provided and resources may be stored
type Location struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Location"

	// Address of the location
	Address *Address `json:"address,omitempty"`

	// Alternate names that the location is known as
	Alias []string `json:"alias,omitempty"`

	// Collection of characteristics that describe the location
	Characteristic []common.CodeableConcept `json:"characteristic,omitempty"`

	// Official contact details for the location
	Contact []ExtendedContactDetail `json:"contact,omitempty"`

	// Additional details about the location
	Description *string `json:"description,omitempty"`

	// Technical endpoints providing access to services operated for the location
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// Physical form of the location
	Form *common.CodeableConcept `json:"form,omitempty"`

	// What days/times during a week is this location usually open
	HoursOfOperation []Availability `json:"hoursOfOperation,omitempty"`

	// Unique code or number identifying the location to its users
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Organization responsible for provisioning and upkeep
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// Whether this location is a class of locations or a specific location
	Mode *LocationMode `json:"mode,omitempty"`

	// Name of the location as used by humans
	Name *string `json:"name,omitempty"`

	// The operational status of the location
	OperationalStatus *common.Coding `json:"operationalStatus,omitempty"`

	// Another location this location is physically part of
	PartOf *common.Reference `json:"partOf,omitempty"`

	// The absolute geographic location
	Position *LocationPosition `json:"position,omitempty"`

	// Whether the location is still in active use
	Status *LocationStatus `json:"status,omitempty"`

	// Type of function performed at the location
	Type []common.CodeableConcept `json:"type,omitempty"`

	// Connection details for virtual services
	VirtualService []VirtualServiceDetail `json:"virtualService,omitempty"`
}
