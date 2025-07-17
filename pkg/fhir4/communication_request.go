package fhir4

import (
	"encoding/json"

	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// CommunicationRequest represents a FHIR R4 CommunicationRequest resource
type CommunicationRequest struct {
	DomainResource

	ResourceType       string                        `json:"resourceType"`
	Identifier         []common.Identifier           `json:"identifier,omitempty"`
	BasedOn            []common.Reference            `json:"basedOn,omitempty"`
	Replaces           []common.Reference            `json:"replaces,omitempty"`
	GroupIdentifier    *common.Identifier            `json:"groupIdentifier,omitempty"`
	Status             string                        `json:"status"`
	StatusReason       *common.CodeableConcept       `json:"statusReason,omitempty"`
	Category           []common.CodeableConcept      `json:"category,omitempty"`
	Priority           *string                       `json:"priority,omitempty"`
	DoNotPerform       *bool                         `json:"doNotPerform,omitempty"`
	Medium             []common.CodeableConcept      `json:"medium,omitempty"`
	Subject            *common.Reference             `json:"subject,omitempty"`
	About              []common.Reference            `json:"about,omitempty"`
	Encounter          *common.Reference             `json:"encounter,omitempty"`
	Payload            []CommunicationRequestPayload `json:"payload,omitempty"`
	OccurrenceDateTime *string                       `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *common.Period                `json:"occurrencePeriod,omitempty"`
	AuthoredOn         *string                       `json:"authoredOn,omitempty"`
	Requester          *common.Reference             `json:"requester,omitempty"`
	Recipient          []common.Reference            `json:"recipient,omitempty"`
	Sender             *common.Reference             `json:"sender,omitempty"`
	ReasonCode         []common.CodeableConcept      `json:"reasonCode,omitempty"`
	ReasonReference    []common.Reference            `json:"reasonReference,omitempty"`
	Note               []common.Annotation           `json:"note,omitempty"`
}

type CommunicationRequestPayload struct {
	common.BackboneElement
	ContentString     *string            `json:"contentString,omitempty"`
	ContentAttachment *common.Attachment `json:"contentAttachment,omitempty"`
	ContentReference  *common.Reference  `json:"contentReference,omitempty"`
}

func (c *CommunicationRequest) MarshalJSON() ([]byte, error) {
	type Alias CommunicationRequest
	return json.Marshal(&struct {
		*Alias
		ResourceType string `json:"resourceType"`
	}{
		Alias:        (*Alias)(c),
		ResourceType: "CommunicationRequest",
	})
}
