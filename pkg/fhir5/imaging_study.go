package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// ImagingStudyStatus represents the status of an imaging study
type ImagingStudyStatus string

const (
	ImagingStudyStatusRegistered     ImagingStudyStatus = "registered"
	ImagingStudyStatusAvailable      ImagingStudyStatus = "available"
	ImagingStudyStatusCancelled      ImagingStudyStatus = "cancelled"
	ImagingStudyStatusEnteredInError ImagingStudyStatus = "entered-in-error"
	ImagingStudyStatusUnknown        ImagingStudyStatus = "unknown"
)

// ImagingStudySeriesPerformer represents who or what performed the series
type ImagingStudySeriesPerformer struct {
	common.BackboneElement

	// Indicates who or what performed the series
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of involvement of the performer in the series
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// ImagingStudySeriesInstance represents a single SOP instance within the series
type ImagingStudySeriesInstance struct {
	common.BackboneElement

	// The number of instance in the series
	Number *int `json:"number,omitempty"`

	// DICOM instance type
	SopClass common.Coding `json:"sopClass"`

	// Particularly for post-acquisition analytic objects
	Title *string `json:"title,omitempty"`

	// See DICOM PS3.3 C.12.1
	Uid string `json:"uid"`
}

// ImagingStudySeries represents each study has one or more series of images or other content
type ImagingStudySeries struct {
	common.BackboneElement

	// The anatomic structures examined
	BodySite *CodeableReference `json:"bodySite,omitempty"`

	// A description of the series
	Description *string `json:"description,omitempty"`

	// Typical endpoint types include DICOM WADO-RS
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// A single SOP instance within the series
	Instance []ImagingStudySeriesInstance `json:"instance,omitempty"`

	// The laterality of the anatomic structures examined
	Laterality *common.CodeableConcept `json:"laterality,omitempty"`

	// The distinct modality for this series
	Modality common.CodeableConcept `json:"modality"`

	// The numeric identifier of this series in the study
	Number *int `json:"number,omitempty"`

	// Number of SOP Instances in the Study
	NumberOfInstances *int `json:"numberOfInstances,omitempty"`

	// If the person who performed the series is not known, their Organization may be recorded
	Performer []ImagingStudySeriesPerformer `json:"performer,omitempty"`

	// The specimen imaged
	Specimen []common.Reference `json:"specimen,omitempty"`

	// The date and time the series was started
	Started *string `json:"started,omitempty"`

	// See DICOM PS3.3 C.7.3
	Uid string `json:"uid"`
}

// ImagingStudy represents the content produced in a DICOM imaging study
type ImagingStudy struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ImagingStudy"

	// A list of the diagnostic requests that resulted in this imaging study
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The Imaging Manager description of the study
	Description *string `json:"description,omitempty"`

	// This will typically be the encounter the event occurred within
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Typical endpoint types include DICOM WADO-RS
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// See discussion under Imaging Study Implementation Notes
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The principal physical location where the ImagingStudy was performed
	Location *common.Reference `json:"location,omitempty"`

	// A list of all the distinct values of series.modality
	Modality []common.CodeableConcept `json:"modality,omitempty"`

	// Per the recommended DICOM mapping
	Note []Annotation `json:"note,omitempty"`

	// Number of SOP Instances in Study
	NumberOfInstances *int `json:"numberOfInstances,omitempty"`

	// Number of Series in the Study
	NumberOfSeries *int `json:"numberOfSeries,omitempty"`

	// To link an ImagingStudy to an Encounter use encounter
	PartOf []common.Reference `json:"partOf,omitempty"`

	// This field corresponds to the DICOM Procedure Code Sequence
	Procedure []CodeableReference `json:"procedure,omitempty"`

	// Description of clinical condition indicating why the ImagingStudy was requested
	Reason []CodeableReference `json:"reason,omitempty"`

	// The requesting/referring physician
	Referrer *common.Reference `json:"referrer,omitempty"`

	// Each study has one or more series of images or other content
	Series []ImagingStudySeries `json:"series,omitempty"`

	// Date and time the study started
	Started *string `json:"started,omitempty"`

	// Unknown does not represent "other" - one of the defined statuses must apply
	Status ImagingStudyStatus `json:"status"`

	// QA phantoms can be recorded with a Device; multiple subjects can be recorded with a Group
	Subject common.Reference `json:"subject"`
}
