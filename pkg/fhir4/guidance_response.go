package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// GuidanceResponse represents a guidance response is the formal response to a guidance request
type GuidanceResponse struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "GuidanceResponse"

	// If the evaluation could not be completed due to lack of information, or additional information would potentially result in a more accurate response
	DataRequirement []common.DataRequirement `json:"dataRequirement,omitempty"`

	// This will typically be the encounter the event occurred within, but some activities may be initiated prior to or after the official completion of an encounter
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Messages resulting from the evaluation of the artifact or artifacts
	EvaluationMessage []common.Reference `json:"evaluationMessage,omitempty"`

	// Allows a service to provide unique, business identifiers for the response
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// An identifier, CodeableConcept or canonical reference to the guidance that was requested
	ModuleUri *string `json:"moduleUri,omitempty"`

	// An identifier, CodeableConcept or canonical reference to the guidance that was requested
	ModuleCanonical *string `json:"moduleCanonical,omitempty"`

	// An identifier, CodeableConcept or canonical reference to the guidance that was requested
	ModuleCodeableConcept *common.CodeableConcept `json:"moduleCodeableConcept,omitempty"`

	// Provides a mechanism to communicate additional information about the response
	Note []common.Annotation `json:"note,omitempty"`

	// Indicates when the guidance response was processed
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`

	// The output parameters of the evaluation, if any
	OutputParameters *common.Reference `json:"outputParameters,omitempty"`

	// Provides a reference to the device that performed the guidance
	Performer *common.Reference `json:"performer,omitempty"`

	// Describes the reason for the guidance response in coded or textual form
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Indicates the reason the request was initiated
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// The identifier of the request associated with this response
	RequestIdentifier *common.Identifier `json:"requestIdentifier,omitempty"`

	// The actions, if any, produced by the evaluation of the artifact
	Result *common.Reference `json:"result,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status GuidanceResponseStatus `json:"status"`

	// The patient for which the request was processed
	Subject *common.Reference `json:"subject,omitempty"`
}

// GuidanceResponseStatus represents the status of the guidance response
type GuidanceResponseStatus string

const (
	GuidanceResponseStatusSuccess        GuidanceResponseStatus = "success"
	GuidanceResponseStatusDataRequested  GuidanceResponseStatus = "data-requested"
	GuidanceResponseStatusDataRequired   GuidanceResponseStatus = "data-required"
	GuidanceResponseStatusInProgress     GuidanceResponseStatus = "in-progress"
	GuidanceResponseStatusFailure        GuidanceResponseStatus = "failure"
	GuidanceResponseStatusEnteredInError GuidanceResponseStatus = "entered-in-error"
)
