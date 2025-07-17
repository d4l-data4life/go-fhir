// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// CommunicationRequestPayload represents text, attachment(s), or resource(s) to be communicated
type CommunicationRequestPayload struct {
	common.BackboneElement

	ContentAttachment      *Attachment             `json:"contentAttachment,omitempty"`
	ContentReference       *common.Reference       `json:"contentReference,omitempty"`
	ContentCodeableConcept *common.CodeableConcept `json:"contentCodeableConcept,omitempty"`
}

// CommunicationRequest represents a request to convey information
type CommunicationRequest struct {
	DomainResource

	ResourceType string `json:"resourceType"` // Always "CommunicationRequest"

	About               []common.Reference            `json:"about,omitempty"`
	AuthoredOn          *string                       `json:"authoredOn,omitempty"`
	BasedOn             []common.Reference            `json:"basedOn,omitempty"`
	Category            []common.CodeableConcept      `json:"category,omitempty"`
	DoNotPerform        *bool                         `json:"doNotPerform,omitempty"`
	Encounter           *common.Reference             `json:"encounter,omitempty"`
	GroupIdentifier     *common.Identifier            `json:"groupIdentifier,omitempty"`
	Identifier          []common.Identifier           `json:"identifier,omitempty"`
	InformationProvider []common.Reference            `json:"informationProvider,omitempty"`
	Intent              CommunicationRequestIntent    `json:"intent"`
	Medium              []common.CodeableConcept      `json:"medium,omitempty"`
	Note                []Annotation                  `json:"note,omitempty"`
	OccurrenceDateTime  *string                       `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod    *common.Period                `json:"occurrencePeriod,omitempty"`
	Payload             []CommunicationRequestPayload `json:"payload,omitempty"`
	Priority            *CommunicationRequestPriority `json:"priority,omitempty"`
	Reason              []CodeableReference           `json:"reason,omitempty"`
	Recipient           []common.Reference            `json:"recipient,omitempty"`
	Replaces            []common.Reference            `json:"replaces,omitempty"`
	Requester           *common.Reference             `json:"requester,omitempty"`
	Sender              *common.Reference             `json:"sender,omitempty"`
	Status              CommunicationRequestStatus    `json:"status"`
	StatusReason        *common.CodeableConcept       `json:"statusReason,omitempty"`
	Subject             *common.Reference             `json:"subject,omitempty"`
	Topic               *common.CodeableConcept       `json:"topic,omitempty"`
}

type CommunicationRequestIntent string

const (
	CommunicationRequestIntentProposal      CommunicationRequestIntent = "proposal"
	CommunicationRequestIntentPlan          CommunicationRequestIntent = "plan"
	CommunicationRequestIntentDirective     CommunicationRequestIntent = "directive"
	CommunicationRequestIntentOrder         CommunicationRequestIntent = "order"
	CommunicationRequestIntentOriginalOrder CommunicationRequestIntent = "original-order"
	CommunicationRequestIntentReflexOrder   CommunicationRequestIntent = "reflex-order"
	CommunicationRequestIntentFillerOrder   CommunicationRequestIntent = "filler-order"
	CommunicationRequestIntentInstanceOrder CommunicationRequestIntent = "instance-order"
	CommunicationRequestIntentOption        CommunicationRequestIntent = "option"
)

type CommunicationRequestPriority string

const (
	CommunicationRequestPriorityRoutine CommunicationRequestPriority = "routine"
	CommunicationRequestPriorityUrgent  CommunicationRequestPriority = "urgent"
	CommunicationRequestPriorityASAP    CommunicationRequestPriority = "asap"
	CommunicationRequestPriorityStat    CommunicationRequestPriority = "stat"
)

type CommunicationRequestStatus string

const (
	CommunicationRequestStatusDraft          CommunicationRequestStatus = "draft"
	CommunicationRequestStatusActive         CommunicationRequestStatus = "active"
	CommunicationRequestStatusOnHold         CommunicationRequestStatus = "on-hold"
	CommunicationRequestStatusRevoked        CommunicationRequestStatus = "revoked"
	CommunicationRequestStatusCompleted      CommunicationRequestStatus = "completed"
	CommunicationRequestStatusEnteredInError CommunicationRequestStatus = "entered-in-error"
	CommunicationRequestStatusUnknown        CommunicationRequestStatus = "unknown"
)
