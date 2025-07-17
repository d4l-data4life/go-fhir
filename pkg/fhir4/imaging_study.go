package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ImagingStudy represents a representation of the content produced in a DICOM imaging study
type ImagingStudy struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ImagingStudy"

	// A list of the diagnostic requests that resulted in this imaging study being performed
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The Imaging Manager description of the study
	Description *string `json:"description,omitempty"`

	// This will typically be the encounter the event occurred within, but some events may be initiated prior to or after the official completion of an encounter
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Typical endpoint types include DICOM WADO-RS, which is used to retrieve DICOM instances in native or rendered formats
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// See discussion under Imaging Study Implementation Notes for encoding of DICOM Study Instance UID
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Who read the study and interpreted the images or other content
	Interpreter []common.Reference `json:"interpreter,omitempty"`

	// The principal physical location where the ImagingStudy was performed
	Location *common.Reference `json:"location,omitempty"`

	// A list of all the series.modality values that are actual acquisition modalities
	Modality []common.Coding `json:"modality,omitempty"`

	// Per the recommended DICOM mapping, this element is derived from the Study Description attribute
	Note []common.Annotation `json:"note,omitempty"`

	// Number of SOP Instances in Study
	NumberOfInstances *int `json:"numberOfInstances,omitempty"`

	// Number of Series in the Study
	NumberOfSeries *int `json:"numberOfSeries,omitempty"`

	// The code for the performed procedure type
	ProcedureCode []common.CodeableConcept `json:"procedureCode,omitempty"`

	// The procedure which this ImagingStudy was part of
	ProcedureReference *common.Reference `json:"procedureReference,omitempty"`

	// Description of clinical condition indicating why the ImagingStudy was requested
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Indicates another resource whose existence justifies this Study
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// The requesting/referring physician
	Referrer *common.Reference `json:"referrer,omitempty"`

	// Each study has one or more series of images or other content
	Series []ImagingStudySeries `json:"series,omitempty"`

	// Date and time the study started
	Started *string `json:"started,omitempty"`

	// Unknown does not represent "other" - one of the defined statuses must apply
	Status ImagingStudyStatus `json:"status"`

	// QA phantoms can be recorded with a Device; multiple subjects (such as mice) can be recorded with a Group
	Subject common.Reference `json:"subject"`
}

// ImagingStudyStatus represents the status of the imaging study
type ImagingStudyStatus string

const (
	ImagingStudyStatusRegistered     ImagingStudyStatus = "registered"
	ImagingStudyStatusAvailable      ImagingStudyStatus = "available"
	ImagingStudyStatusCancelled      ImagingStudyStatus = "cancelled"
	ImagingStudyStatusEnteredInError ImagingStudyStatus = "entered-in-error"
	ImagingStudyStatusUnknown        ImagingStudyStatus = "unknown"
)

// ImagingStudySeries represents each study has one or more series of images or other content
type ImagingStudySeries struct {
	common.BackboneElement

	// The anatomic structures examined
	BodySite *common.Coding `json:"bodySite,omitempty"`

	// A description of the series
	Description *string `json:"description,omitempty"`

	// Typical endpoint types include DICOM WADO-RS, which is used to retrieve DICOM instances in native or rendered formats
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// A single SOP instance within the series, e.g. an image, or presentation state
	Instance []ImagingStudySeriesInstance `json:"instance,omitempty"`

	// The laterality of the (possibly paired) anatomic structures examined
	Laterality *common.Coding `json:"laterality,omitempty"`

	// The modality of this series sequence
	Modality common.Coding `json:"modality"`

	// The numeric identifier of this series in the study
	Number *int `json:"number,omitempty"`

	// Number of SOP Instances in the Study
	NumberOfInstances *int `json:"numberOfInstances,omitempty"`

	// If the person who performed the series is not known, their Organization may be recorded
	Performer []ImagingStudySeriesPerformer `json:"performer,omitempty"`

	// The specimen imaged, e.g., for whole slide imaging of a biopsy
	Specimen []common.Reference `json:"specimen,omitempty"`

	// The date and time the series was started
	Started *string `json:"started,omitempty"`

	// See DICOM PS3.3 C.7.3
	Uid string `json:"uid"`
}

// ImagingStudySeriesInstance represents a single SOP instance within the series
type ImagingStudySeriesInstance struct {
	common.BackboneElement

	// The number of instance in the series
	Number *int `json:"number,omitempty"`

	// DICOM instance type
	SopClass common.Coding `json:"sopClass"`

	// Particularly for post-acquisition analytic objects, such as SR, presentation states, value mapping, etc
	Title *string `json:"title,omitempty"`

	// See DICOM PS3.3 C.12.1
	Uid string `json:"uid"`
}

// ImagingStudySeriesPerformer represents if the person who performed the series is not known, their Organization may be recorded
type ImagingStudySeriesPerformer struct {
	common.BackboneElement

	// Indicates who or what performed the series
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of involvement of the performer in the series
	Function *common.CodeableConcept `json:"function,omitempty"`
}
