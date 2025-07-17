package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

type MessageDefinitionStatus string

const (
	MessageDefinitionStatusDraft   MessageDefinitionStatus = "draft"
	MessageDefinitionStatusActive  MessageDefinitionStatus = "active"
	MessageDefinitionStatusRetired MessageDefinitionStatus = "retired"
	MessageDefinitionStatusUnknown MessageDefinitionStatus = "unknown"
)

type MessageDefinitionCategory string

type MessageDefinitionResponseRequired string

type MessageDefinitionFocus struct {
	common.BackboneElement

	Code    string  `json:"code"`
	Max     *string `json:"max,omitempty"`
	Min     int     `json:"min"`
	Profile *string `json:"profile,omitempty"`
}

type MessageDefinitionAllowedResponse struct {
	common.BackboneElement

	Message   string  `json:"message"`
	Situation *string `json:"situation,omitempty"`
}

type MessageDefinition struct {
	DomainResource

	ResourceType string `json:"resourceType"` // Always "MessageDefinition"

	AllowedResponse  []MessageDefinitionAllowedResponse `json:"allowedResponse,omitempty"`
	Base             *string                            `json:"base,omitempty"`
	Category         *MessageDefinitionCategory         `json:"category,omitempty"`
	Contact          []ContactDetail                    `json:"contact,omitempty"`
	Copyright        *string                            `json:"copyright,omitempty"`
	CopyrightLabel   *string                            `json:"copyrightLabel,omitempty"`
	Date             string                             `json:"date"`
	Description      *string                            `json:"description,omitempty"`
	EventCoding      *common.Coding                     `json:"eventCoding,omitempty"`
	EventUri         *string                            `json:"eventUri,omitempty"`
	Experimental     *bool                              `json:"experimental,omitempty"`
	Focus            []MessageDefinitionFocus           `json:"focus,omitempty"`
	Graph            *string                            `json:"graph,omitempty"`
	Identifier       []common.Identifier                `json:"identifier,omitempty"`
	Jurisdiction     []common.CodeableConcept           `json:"jurisdiction,omitempty"`
	Name             *string                            `json:"name,omitempty"`
	Parent           []string                           `json:"parent,omitempty"`
	Publisher        *string                            `json:"publisher,omitempty"`
	Purpose          *string                            `json:"purpose,omitempty"`
	Replaces         []string                           `json:"replaces,omitempty"`
	ResponseRequired *MessageDefinitionResponseRequired `json:"responseRequired,omitempty"`
	Status           MessageDefinitionStatus            `json:"status"`
	Title            *string                            `json:"title,omitempty"`
	Url              *string                            `json:"url,omitempty"`
	UseContext       []UsageContext                     `json:"useContext,omitempty"`
	Version          *string                            `json:"version,omitempty"`
}
