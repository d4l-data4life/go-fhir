// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SpecimenFeature represents a physical feature or landmark on a specimen
type SpecimenFeature struct {
	common.BackboneElement
	Description string                 `json:"description"`
	Type        common.CodeableConcept `json:"type"`
}

// SpecimenCollection represents details concerning the specimen collection
type SpecimenCollection struct {
	common.BackboneElement
	BodySite                     *CodeableReference      `json:"bodySite,omitempty"`
	CollectedDateTime            *string                 `json:"collectedDateTime,omitempty"`
	CollectedPeriod              *common.Period          `json:"collectedPeriod,omitempty"`
	Collector                    *common.Reference       `json:"collector,omitempty"`
	Device                       *CodeableReference      `json:"device,omitempty"`
	Duration                     *Duration               `json:"duration,omitempty"`
	FastingStatusCodeableConcept *common.CodeableConcept `json:"fastingStatusCodeableConcept,omitempty"`
	FastingStatusDuration        *Duration               `json:"fastingStatusDuration,omitempty"`
	Method                       *common.CodeableConcept `json:"method,omitempty"`
	Procedure                    *common.Reference       `json:"procedure,omitempty"`
	Quantity                     *common.Quantity        `json:"quantity,omitempty"`
}

// SpecimenProcessing represents details concerning processing and processing steps for the specimen
type SpecimenProcessing struct {
	common.BackboneElement
	Additive     []common.Reference      `json:"additive,omitempty"`
	Description  *string                 `json:"description,omitempty"`
	Method       *common.CodeableConcept `json:"method,omitempty"`
	TimeDateTime *string                 `json:"timeDateTime,omitempty"`
	TimePeriod   *common.Period          `json:"timePeriod,omitempty"`
}

// SpecimenContainer represents the container holding the specimen
type SpecimenContainer struct {
	common.BackboneElement
	Device           common.Reference  `json:"device"`
	Location         *common.Reference `json:"location,omitempty"`
	SpecimenQuantity *common.Quantity  `json:"specimenQuantity,omitempty"`
}

// Specimen represents a sample to be used for analysis
type Specimen struct {
	DomainResource
	ResourceType        string                   `json:"resourceType"` // Always "Specimen"
	AccessionIdentifier *common.Identifier       `json:"accessionIdentifier,omitempty"`
	Collection          *SpecimenCollection      `json:"collection,omitempty"`
	Combined            *string                  `json:"combined,omitempty"`
	Condition           []common.CodeableConcept `json:"condition,omitempty"`
	Container           []SpecimenContainer      `json:"container,omitempty"`
	Feature             []SpecimenFeature        `json:"feature,omitempty"`
	Identifier          []common.Identifier      `json:"identifier,omitempty"`
	Note                []Annotation             `json:"note,omitempty"`
	Processing          []SpecimenProcessing     `json:"processing,omitempty"`
	Parent              []common.Reference       `json:"parent,omitempty"`
	ReceivedTime        *string                  `json:"receivedTime,omitempty"`
	Request             []common.Reference       `json:"request,omitempty"`
	Status              *string                  `json:"status,omitempty"`
	Subject             *common.Reference        `json:"subject,omitempty"`
	Type                *common.CodeableConcept  `json:"type,omitempty"`
}
