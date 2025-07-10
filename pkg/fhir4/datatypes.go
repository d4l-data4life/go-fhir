// Package fhir4 contains FHIR R4 (version 4.0.1) type definitions
package fhir4

import (
	"time"

	"github.com/go-fhir/go-fhir/pkg/common"
)

// Address represents an address expressed using postal conventions
type Address struct {
	common.Element

	// The name of the city, town, suburb, village or other community or delivery center
	City *string `json:"city,omitempty"`

	// ISO 3166 3 letter codes can be used in place of a human readable country name
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
	common.Element

	// Organization is used when there's no need for specific attribution
	AuthorReference *common.Reference `json:"authorReference,omitempty"`
	AuthorString    *string           `json:"authorString,omitempty"`

	// The text of the annotation in markdown format
	Text string `json:"text"`

	// Indicates when this particular annotation was made
	Time *time.Time `json:"time,omitempty"`
}

// Attachment represents data content defined in other formats
type Attachment struct {
	common.Element

	// Identifies the type of the data in the attachment
	ContentType *string `json:"contentType,omitempty"`

	// The date that the attachment was first created
	Creation *time.Time `json:"creation,omitempty"`

	// The base64-encoded data
	Data *string `json:"data,omitempty"`

	// Hash of the data (sha-1, base64ed)
	Hash *string `json:"hash,omitempty"`

	// Human language of the content (BCP-47)
	Language *string `json:"language,omitempty"`

	// Number of bytes of content (if url provided)
	Size *int `json:"size,omitempty"`

	// Label to display in place of the data
	Title *string `json:"title,omitempty"`

	// Uri where the data can be found
	URL *string `json:"url,omitempty"`
}

// ContactDetail specifies contact information for a person or organization
type ContactDetail struct {
	common.Element

	// Name of an individual to contact
	Name *string `json:"name,omitempty"`

	// Contact details for individual or organization
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

// ContactPoint represents details for technology mediated contact points
type ContactPoint struct {
	common.Element

	// Time period when the contact point was/is in use
	Period *common.Period `json:"period,omitempty"`

	// Rank for determining order of preference
	Rank *int `json:"rank,omitempty"`

	// phone | fax | email | pager | url | sms | other
	System *ContactPointSystem `json:"system,omitempty"`

	// home | work | temp | old | mobile
	Use *ContactPointUse `json:"use,omitempty"`

	// The actual contact point details
	Value *string `json:"value,omitempty"`
}

// ContactPointSystem represents telecommunications form for contact point
type ContactPointSystem string

const (
	ContactPointSystemPhone ContactPointSystem = "phone"
	ContactPointSystemFax   ContactPointSystem = "fax"
	ContactPointSystemEmail ContactPointSystem = "email"
	ContactPointSystemPager ContactPointSystem = "pager"
	ContactPointSystemURL   ContactPointSystem = "url"
	ContactPointSystemSMS   ContactPointSystem = "sms"
	ContactPointSystemOther ContactPointSystem = "other"
)

// ContactPointUse represents the use of a contact point
type ContactPointUse string

const (
	ContactPointUseHome   ContactPointUse = "home"
	ContactPointUseWork   ContactPointUse = "work"
	ContactPointUseTemp   ContactPointUse = "temp"
	ContactPointUseOld    ContactPointUse = "old"
	ContactPointUseMobile ContactPointUse = "mobile"
)

// Contributor represents a contributor to the content of a knowledge asset
type Contributor struct {
	common.Element

	// Contact details to assist a user in finding and communicating with the contributor
	Contact []ContactDetail `json:"contact,omitempty"`

	// The name of the individual or organization responsible for the contribution
	Name string `json:"name"`

	// author | editor | reviewer | endorser
	Type ContributorType `json:"type"`
}

// ContributorType represents the type of contributor
type ContributorType string

const (
	ContributorTypeAuthor   ContributorType = "author"
	ContributorTypeEditor   ContributorType = "editor"
	ContributorTypeReviewer ContributorType = "reviewer"
	ContributorTypeEndorser ContributorType = "endorser"
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
