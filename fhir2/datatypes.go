// Package fhir2 contains FHIR R2 (version 1.0.2) type definitions
package fhir2

import (
	"time"

	"github.com/go-fhir/go-fhir/common"
)

// Address represents an address expressed using postal conventions
// Note: R2 is the oldest version with fewer fields than later versions
type Address struct {
	common.Element

	// The name of the city, town, village or other community or delivery center
	City *string `json:"city,omitempty"`

	// Country - a nation as commonly understood or generally accepted
	Country *string `json:"country,omitempty"`

	// The name of the administrative area (county)
	District *string `json:"district,omitempty"`

	// This component contains the house number, apartment number, street name, street direction
	Line []string `json:"line,omitempty"`

	// Allows addresses to be placed in historical context
	Period *common.Period `json:"period,omitempty"`

	// A postal code designating a region defined by the postal service
	PostalCode *string `json:"postalCode,omitempty"`

	// Sub-unit of a country with limited sovereignty in a federally organized country
	State *string `json:"state,omitempty"`

	// A renderable, unencoded form
	Text *string `json:"text,omitempty"`

	// postal | physical | both - Distinguishes between physical and mailing addresses
	Type *AddressType `json:"type,omitempty"`

	// home | work | temp | old - purpose of this address
	Use *AddressUse `json:"use,omitempty"`
}

// AddressType represents the type of address (R2 version)
type AddressType string

const (
	AddressTypePostal   AddressType = "postal"
	AddressTypePhysical AddressType = "physical"
	AddressTypeBoth     AddressType = "both"
)

// AddressUse represents the use of an address (R2 version - no billing option)
type AddressUse string

const (
	AddressUseHome AddressUse = "home"
	AddressUseWork AddressUse = "work"
	AddressUseTemp AddressUse = "temp"
	AddressUseOld  AddressUse = "old"
)

// Age represents a duration of time during which an organism has existed
type Age struct {
	common.Quantity
}

// Annotation represents a text note with information about who made the statement and when
type Annotation struct {
	common.Element

	// The individual responsible for making the annotation
	AuthorReference *common.Reference `json:"authorReference,omitempty"`
	AuthorString    *string           `json:"authorString,omitempty"`

	// The text of the annotation
	Text string `json:"text"`

	// Indicates when this particular annotation was made
	Time *time.Time `json:"time,omitempty"`
}

// Attachment represents data content defined in other formats
type Attachment struct {
	common.Element

	// Processors of the data need to be able to know how to interpret the data
	ContentType *string `json:"contentType,omitempty"`

	// This is often tracked as an integrity issue for use of the attachment
	Creation *time.Time `json:"creation,omitempty"`

	// The data needs to able to be transmitted inline
	Data *string `json:"data,omitempty"`

	// Included so that applications can verify that the contents have not changed
	Hash *string `json:"hash,omitempty"`

	// Users need to be able to choose between the languages in a set of attachments
	Language *string `json:"language,omitempty"`

	// Representing the size allows applications to determine whether they should fetch the content
	Size *int `json:"size,omitempty"`

	// Applications need a label to display to a human user in place of the actual data
	Title *string `json:"title,omitempty"`

	// The data needs to be transmitted by reference
	URL *string `json:"url,omitempty"`
}

// ContactPoint represents details for technology mediated contact points
type ContactPoint struct {
	common.Element

	// Time period when the contact point was/is in use
	Period *common.Period `json:"period,omitempty"`

	// Specifies a preferred order in which to use a set of contacts
	Rank *int `json:"rank,omitempty"`

	// phone | fax | email | pager | other - Telecommunications form for contact point
	System *ContactPointSystem `json:"system,omitempty"`

	// home | work | temp | old | mobile - The use of a contact point
	Use *ContactPointUse `json:"use,omitempty"`

	// The actual contact point details
	Value *string `json:"value,omitempty"`
}

// ContactPointSystem represents telecommunications form for contact point (R2 version)
type ContactPointSystem string

const (
	ContactPointSystemPhone ContactPointSystem = "phone"
	ContactPointSystemFax   ContactPointSystem = "fax"
	ContactPointSystemEmail ContactPointSystem = "email"
	ContactPointSystemPager ContactPointSystem = "pager"
	ContactPointSystemOther ContactPointSystem = "other"
	// Note: R2 doesn't have "url" or "sms" options that were added in later versions
)

// ContactPointUse represents the use of a contact point (R2 version)
type ContactPointUse string

const (
	ContactPointUseHome   ContactPointUse = "home"
	ContactPointUseWork   ContactPointUse = "work"
	ContactPointUseTemp   ContactPointUse = "temp"
	ContactPointUseOld    ContactPointUse = "old"
	ContactPointUseMobile ContactPointUse = "mobile"
)

// Count represents a measured amount
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
	common.Element

	// Low limit
	Low *common.Quantity `json:"low,omitempty"`

	// High limit
	High *common.Quantity `json:"high,omitempty"`
}

// Ratio represents a relationship of two Quantity values
type Ratio struct {
	common.Element

	// Numerator value
	Numerator *common.Quantity `json:"numerator,omitempty"`

	// Denominator value
	Denominator *common.Quantity `json:"denominator,omitempty"`
}
