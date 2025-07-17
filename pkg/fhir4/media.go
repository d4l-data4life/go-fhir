package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Media represents a photo, video, or audio recording acquired or used in healthcare
type Media struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Media"

	// A procedure that is fulfilled in whole or in part by the creation of this media
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Only used if not implicit in code found in Observation.code
	BodySite *common.CodeableConcept `json:"bodySite,omitempty"`

	// Recommended content types: image/jpeg, image/png, image/tiff, video/mpeg, audio/mp4, application/dicom
	Content common.Attachment `json:"content"`

	// The date and time(s) at which the media was collected
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The date and time(s) at which the media was collected
	CreatedPeriod *common.Period `json:"createdPeriod,omitempty"`

	// An extension should be used if further typing of the device is needed
	Device *common.Reference `json:"device,omitempty"`

	// The name of the device / manufacturer of the device that was used to make the recording
	DeviceName *string `json:"deviceName,omitempty"`

	// The duration might differ from occurrencePeriod if recording was paused
	Duration *float64 `json:"duration,omitempty"`

	// This will typically be the encounter the media occurred within
	Encounter *common.Reference `json:"encounter,omitempty"`

	// If the number of frames is not supplied, the value may be unknown
	Frames *int `json:"frames,omitempty"`

	// Height of the image in pixels (photo/video)
	Height *int `json:"height,omitempty"`

	// The identifier label and use can be used to determine what kind of identifier it is
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be the same as the lastUpdated time of the resource itself
	Issued *string `json:"issued,omitempty"`

	// Details of the type of the media - usually, how it was acquired (what type of device)
	Modality *common.CodeableConcept `json:"modality,omitempty"`

	// Not to be used for observations, conclusions, etc
	Note []common.Annotation `json:"note,omitempty"`

	// The person who administered the collection of the image
	Operator *common.Reference `json:"operator,omitempty"`

	// Not to be used to link an event to an Encounter - use Media.encounter for that
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Textual reasons can be captured using reasonCode.text
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// A nominal state-transition diagram can be found in the Event pattern documentation
	Status MediaStatus `json:"status"`

	// Who/What this Media is a record of
	Subject *common.Reference `json:"subject,omitempty"`

	// A code that classifies whether the media is an image, video or audio recording or some other media category
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The name of the imaging view e.g. Lateral or Antero-posterior (AP)
	View *common.CodeableConcept `json:"view,omitempty"`

	// Width of the image in pixels (photo/video)
	Width *int `json:"width,omitempty"`
}

// MediaStatus represents the status of the media
type MediaStatus string

const (
	MediaStatusPreparation    MediaStatus = "preparation"
	MediaStatusInProgress     MediaStatus = "in-progress"
	MediaStatusNotDone        MediaStatus = "not-done"
	MediaStatusOnHold         MediaStatus = "on-hold"
	MediaStatusStopped        MediaStatus = "stopped"
	MediaStatusCompleted      MediaStatus = "completed"
	MediaStatusEnteredInError MediaStatus = "entered-in-error"
	MediaStatusUnknown        MediaStatus = "unknown"
)
