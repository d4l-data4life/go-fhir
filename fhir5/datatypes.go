// Package fhir5 contains FHIR R5 (version 5.0.0) type definitions
package fhir5

import (
	"time"

	"github.com/go-fhir/go-fhir/common"
)

// DataType is the base type for all data types in FHIR R5 (note: this differs from R4 which used Element)
type DataType struct {
	common.Element
}

// Address represents an address expressed using postal conventions
// Note: R5 has additional fields compared to R4
type Address struct {
	DataType

	// The name of the city, town, suburb, village or other community or delivery center
	City *string `json:"city,omitempty"`

	// ISO 3166 2- or 3- letter codes MAY be used in place of a human readable country name
	Country *string `json:"country,omitempty"`

	// District is sometimes known as county
	District *string `json:"district,omitempty"`

	// House number, apartment number, street name, street direction, P.O. Box number, delivery hints
	Line []string `json:"line,omitempty"`

	// Time period when address was/is in use
	Period *common.Period `json:"period,omitempty"`

	// A postal code designating a region defined by the postal service
	PostalCode *string `json:"postalCode,omitempty"`

	// Sub-unit of a country with limited sovereignty in a federally organized country
	State *string `json:"state,omitempty"`

	// Text representation of the address
	Text *string `json:"text,omitempty"`

	// postal | physical | both - The definition of Address
	Type *AddressType `json:"type,omitempty"`

	// home | work | temp | old | billing - purpose of this address
	Use *AddressUse `json:"use,omitempty"`
}

// AddressType represents the type of address
type AddressType string

const (
	AddressTypePostal   AddressType = "postal"
	AddressTypePhysical AddressType = "physical"
	AddressTypeBoth     AddressType = "both"
)

// AddressUse represents the use of an address
type AddressUse string

const (
	AddressUseHome    AddressUse = "home"
	AddressUseWork    AddressUse = "work"
	AddressUseTemp    AddressUse = "temp"
	AddressUseOld     AddressUse = "old"
	AddressUseBilling AddressUse = "billing"
)

// Age represents a duration of time during which an organism (or a process) has existed
type Age struct {
	common.Quantity
}

// Annotation represents a text note which also contains information about who made the statement and when
type Annotation struct {
	DataType

	// Organization is used when there's no need for specific attribution
	AuthorReference *common.Reference `json:"authorReference,omitempty"`
	AuthorString    *string           `json:"authorString,omitempty"`

	// The text of the annotation in markdown format
	Text string `json:"text"`

	// Indicates when this particular annotation was made
	Time *time.Time `json:"time,omitempty"`
}

// Attachment represents data content defined in other formats
// Note: R5 has additional fields like duration, frames, height, width, pages compared to R4
type Attachment struct {
	DataType

	// Identifies the type of the data in the attachment
	ContentType *string `json:"contentType,omitempty"`

	// The date that the attachment was first created
	Creation *time.Time `json:"creation,omitempty"`

	// The base64-encoded data
	Data *string `json:"data,omitempty"`

	// The duration might differ from occurrencePeriod if recording was paused
	Duration *int `json:"duration,omitempty"`

	// Number of frames if not supplied, value may be unknown
	Frames *int `json:"frames,omitempty"`

	// Hash of the data (sha-1, base64ed)
	Hash *string `json:"hash,omitempty"`

	// Height of the image in pixels (photo/video)
	Height *int `json:"height,omitempty"`

	// Human language of the content (BCP-47)
	Language *string `json:"language,omitempty"`

	// The number of pages when printed
	Pages *int `json:"pages,omitempty"`

	// Number of bytes of content (if url provided) - Note: string in R5, int in R4
	Size *string `json:"size,omitempty"`

	// Label to display in place of the data
	Title *string `json:"title,omitempty"`

	// Uri where the data can be found
	URL *string `json:"url,omitempty"`

	// Width of the image in pixels (photo/video)
	Width *int `json:"width,omitempty"`
}

// Availability represents availability data for an item (new in R5)
type Availability struct {
	DataType

	// Times the item is available
	AvailableTime []AvailabilityAvailableTime `json:"availableTime,omitempty"`

	// Not available during this time due to provided reason
	NotAvailableTime []AvailabilityNotAvailableTime `json:"notAvailableTime,omitempty"`
}

// AvailabilityAvailableTime represents times the item is available
type AvailabilityAvailableTime struct {
	common.Element

	// Always available? i.e. 24 hour service
	AllDay *bool `json:"allDay,omitempty"`

	// Closing time of day (ignored if allDay = true)
	AvailableEndTime *string `json:"availableEndTime,omitempty"`

	// Opening time of day (ignored if allDay = true)
	AvailableStartTime *string `json:"availableStartTime,omitempty"`

	// mon | tue | wed | thu | fri | sat | sun
	DaysOfWeek []DaysOfWeek `json:"daysOfWeek,omitempty"`
}

// AvailabilityNotAvailableTime represents not available during this time due to provided reason
type AvailabilityNotAvailableTime struct {
	common.Element

	// Reason presented to the user explaining why time not available
	Description *string `json:"description,omitempty"`

	// Service not available during this period
	During *common.Period `json:"during,omitempty"`
}

// DaysOfWeek represents days of the week
type DaysOfWeek string

const (
	DaysOfWeekMon DaysOfWeek = "mon"
	DaysOfWeekTue DaysOfWeek = "tue"
	DaysOfWeekWed DaysOfWeek = "wed"
	DaysOfWeekThu DaysOfWeek = "thu"
	DaysOfWeekFri DaysOfWeek = "fri"
	DaysOfWeekSat DaysOfWeek = "sat"
	DaysOfWeekSun DaysOfWeek = "sun"
)

// Count represents a measured amount (or an amount that can potentially be measured)
type Count struct {
	common.Quantity
}

// Distance represents a length - a value with a unit that is a physical distance
type Distance struct {
	common.Quantity
}

// Duration represents a length of time
type Duration struct {
	common.Quantity
}

// Range represents a set of values bounded by low and high
type Range struct {
	DataType

	// Low limit
	Low *common.Quantity `json:"low,omitempty"`

	// High limit
	High *common.Quantity `json:"high,omitempty"`
}

// Ratio represents a relationship of two Quantity values
type Ratio struct {
	DataType

	// Numerator value
	Numerator *common.Quantity `json:"numerator,omitempty"`

	// Denominator value
	Denominator *common.Quantity `json:"denominator,omitempty"`
}
