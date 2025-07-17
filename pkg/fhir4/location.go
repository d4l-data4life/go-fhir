package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Location represents details and position information for a physical place
type Location struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Location"

	// Unique code or number identifying the location to its users
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// active | suspended | inactive
	Status *LocationStatus `json:"status,omitempty"`

	// Name of the location as used by humans
	Name *string `json:"name,omitempty"`

	// A list of alternate names that the location is known as, or was known as in the past
	Alias []string `json:"alias,omitempty"`

	// Additional details about the location that could be displayed as further information
	Description *string `json:"description,omitempty"`

	// instance | kind - Indicates whether a resource instance represents a specific location or a class of locations
	Mode *LocationMode `json:"mode,omitempty"`

	// Indicates the type of function performed at the location
	Type []common.CodeableConcept `json:"type,omitempty"`

	// Contact details of the location
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Physical location
	Address *Address `json:"address,omitempty"`

	// Physical form of the location
	PhysicalType *common.CodeableConcept `json:"physicalType,omitempty"`

	// The absolute geographic location
	Position *LocationPosition `json:"position,omitempty"`

	// Organization responsible for provisioning and upkeep
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// Another Location of which this Location is physically a part of
	PartOf *common.Reference `json:"partOf,omitempty"`

	// What days/times during a week is this location usually open
	HoursOfOperation []LocationHoursOfOperation `json:"hoursOfOperation,omitempty"`

	// Description of availability exceptions
	AvailabilityExceptions *string `json:"availabilityExceptions,omitempty"`

	// Technical endpoints providing access to services operated for the location
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// The operational status covers operation values most relevant to beds
	OperationalStatus *common.Coding `json:"operationalStatus,omitempty"`
}

// LocationStatus represents the status of the location
type LocationStatus string

const (
	LocationStatusActive    LocationStatus = "active"
	LocationStatusSuspended LocationStatus = "suspended"
	LocationStatusInactive  LocationStatus = "inactive"
)

// LocationMode represents the mode of the location
type LocationMode string

const (
	LocationModeInstance LocationMode = "instance"
	LocationModeKind     LocationMode = "kind"
)

// LocationPosition represents the absolute geographic location
type LocationPosition struct {
	common.BackboneElement

	// Longitude with WGS84 datum
	Longitude float64 `json:"longitude"`

	// Latitude with WGS84 datum
	Latitude float64 `json:"latitude"`

	// Altitude with WGS84 datum
	Altitude *float64 `json:"altitude,omitempty"`
}

// LocationHoursOfOperation represents the hours of operation for the location
type LocationHoursOfOperation struct {
	common.BackboneElement

	// Indicates which days of the week are available between the start and end Times
	DaysOfWeek []DaysOfWeek `json:"daysOfWeek,omitempty"`

	// The Location is open all day
	AllDay *bool `json:"allDay,omitempty"`

	// Time that the Location opens
	OpeningTime *string `json:"openingTime,omitempty"`

	// Time that the Location closes
	ClosingTime *string `json:"closingTime,omitempty"`
}
