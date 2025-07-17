package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ValueSet represents a set of codes drawn from one or more code systems
type ValueSet struct {
	DomainResource

	ResourceType string `json:"resourceType"` // Always "ValueSet"

	Compose      *ValueSetCompose         `json:"compose,omitempty"`
	Contact      []common.ContactDetail   `json:"contact,omitempty"`
	Copyright    *string                  `json:"copyright,omitempty"`
	Date         *string                  `json:"date,omitempty"`
	Description  *string                  `json:"description,omitempty"`
	Expansion    *ValueSetExpansion       `json:"expansion,omitempty"`
	Experimental *bool                    `json:"experimental,omitempty"`
	Identifier   []common.Identifier      `json:"identifier,omitempty"`
	Immutable    *bool                    `json:"immutable,omitempty"`
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`
	Name         *string                  `json:"name,omitempty"`
	Publisher    *string                  `json:"publisher,omitempty"`
	Purpose      *string                  `json:"purpose,omitempty"`
	Status       string                   `json:"status"` // "draft" | "active" | "retired" | "unknown"
	Title        *string                  `json:"title,omitempty"`
	Url          *string                  `json:"url,omitempty"`
	UseContext   []common.UsageContext    `json:"useContext,omitempty"`
	Version      *string                  `json:"version,omitempty"`
}

// ValueSetCompose represents the Content Logical Definition (CLD) for a ValueSet
type ValueSetCompose struct {
	common.BackboneElement

	Exclude    []ValueSetComposeInclude `json:"exclude,omitempty"`
	Include    []ValueSetComposeInclude `json:"include"`
	Inactive   *bool                    `json:"inactive,omitempty"`
	LockedDate *string                  `json:"lockedDate,omitempty"`
}

type ValueSetComposeInclude struct {
	common.BackboneElement

	Concept  []ValueSetComposeIncludeConcept `json:"concept,omitempty"`
	Filter   []ValueSetComposeIncludeFilter  `json:"filter,omitempty"`
	System   *string                         `json:"system,omitempty"`
	ValueSet []string                        `json:"valueSet,omitempty"`
	Version  *string                         `json:"version,omitempty"`
}

type ValueSetComposeIncludeConcept struct {
	common.BackboneElement

	Code        string                                     `json:"code"`
	Designation []ValueSetComposeIncludeConceptDesignation `json:"designation,omitempty"`
	Display     *string                                    `json:"display,omitempty"`
}

type ValueSetComposeIncludeConceptDesignation struct {
	common.BackboneElement

	Language *string        `json:"language,omitempty"`
	Use      *common.Coding `json:"use,omitempty"`
	Value    string         `json:"value"`
}

type ValueSetComposeIncludeFilter struct {
	common.BackboneElement

	Op       string `json:"op"` // "=" | "is-a" | "descendent-of" | "is-not-a" | "regex" | "in" | "not-in" | "generalizes" | "exists"
	Property string `json:"property"`
	Value    string `json:"value"`
}

// ValueSetExpansion represents the expansion of a ValueSet
type ValueSetExpansion struct {
	common.BackboneElement

	Contains   []ValueSetExpansionContains  `json:"contains,omitempty"`
	Identifier *string                      `json:"identifier,omitempty"`
	Parameter  []ValueSetExpansionParameter `json:"parameter,omitempty"`
	Timestamp  string                       `json:"timestamp"`
	Total      *int                         `json:"total,omitempty"`
	Offset     *int                         `json:"offset,omitempty"`
}

type ValueSetExpansionParameter struct {
	common.BackboneElement

	Name         string   `json:"name"`
	ValueString  *string  `json:"valueString,omitempty"`
	ValueBoolean *bool    `json:"valueBoolean,omitempty"`
	ValueInteger *int     `json:"valueInteger,omitempty"`
	ValueDecimal *float64 `json:"valueDecimal,omitempty"`
	ValueUri     *string  `json:"valueUri,omitempty"`
	ValueCode    *string  `json:"valueCode,omitempty"`
}

type ValueSetExpansionContains struct {
	common.BackboneElement

	Abstract *bool                       `json:"abstract,omitempty"`
	Code     *string                     `json:"code,omitempty"`
	Contains []ValueSetExpansionContains `json:"contains,omitempty"`
	Display  *string                     `json:"display,omitempty"`
	Inactive *bool                       `json:"inactive,omitempty"`
	System   *string                     `json:"system,omitempty"`
	Version  *string                     `json:"version,omitempty"`
}
