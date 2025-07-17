// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// CompositionAttester represents a participant who has attested to the accuracy of the composition
type CompositionAttester struct {
	common.BackboneElement

	Mode  []string          `json:"mode"`
	Time  *string           `json:"time,omitempty"`
	Party *common.Reference `json:"party,omitempty"`
}

// CompositionEvent represents an event the composition documents
type CompositionEvent struct {
	common.BackboneElement

	Code   []common.CodeableConcept `json:"code,omitempty"`
	Period *common.Period           `json:"period,omitempty"`
	Detail []common.Reference       `json:"detail,omitempty"`
}

// CompositionRelatesTo represents relationships to other compositions
type CompositionRelatesTo struct {
	common.BackboneElement

	Code             string             `json:"code"`
	TargetIdentifier *common.Identifier `json:"targetIdentifier,omitempty"`
	TargetReference  *common.Reference  `json:"targetReference,omitempty"`
}

// CompositionSection represents a section within the composition
type CompositionSection struct {
	common.BackboneElement

	Title       *string                 `json:"title,omitempty"`
	Code        *common.CodeableConcept `json:"code,omitempty"`
	Author      []common.Reference      `json:"author,omitempty"`
	Focus       *common.Reference       `json:"focus,omitempty"`
	Text        *Narrative              `json:"text,omitempty"`
	Mode        *string                 `json:"mode,omitempty"`
	OrderedBy   *common.CodeableConcept `json:"orderedBy,omitempty"`
	Entry       []common.Reference      `json:"entry,omitempty"`
	EmptyReason *common.CodeableConcept `json:"emptyReason,omitempty"`
	Section     []CompositionSection    `json:"section,omitempty"`
}

// Composition represents a set of healthcare-related information that is assembled together
type Composition struct {
	DomainResource

	ResourceType string `json:"resourceType"` // Always "Composition"

	Identifier      *common.Identifier       `json:"identifier,omitempty"`
	Status          CompositionStatus        `json:"status"`
	Type            common.CodeableConcept   `json:"type"`
	Category        []common.CodeableConcept `json:"category,omitempty"`
	Subject         *common.Reference        `json:"subject,omitempty"`
	Encounter       *common.Reference        `json:"encounter,omitempty"`
	Date            string                   `json:"date"`
	Author          []common.Reference       `json:"author"`
	Title           string                   `json:"title"`
	Confidentiality *string                  `json:"confidentiality,omitempty"`
	Attester        []CompositionAttester    `json:"attester,omitempty"`
	Custodian       *common.Reference        `json:"custodian,omitempty"`
	RelatesTo       []CompositionRelatesTo   `json:"relatesTo,omitempty"`
	Event           []CompositionEvent       `json:"event,omitempty"`
	Section         []CompositionSection     `json:"section,omitempty"`
}

type CompositionStatus string

const (
	CompositionStatusPreliminary    CompositionStatus = "preliminary"
	CompositionStatusFinal          CompositionStatus = "final"
	CompositionStatusAmended        CompositionStatus = "amended"
	CompositionStatusEnteredInError CompositionStatus = "entered-in-error"
)
