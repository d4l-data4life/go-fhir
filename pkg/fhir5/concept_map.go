// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ConceptMapAdditionalAttribute represents additional attributes for mapping data
type ConceptMapAdditionalAttribute struct {
	common.BackboneElement

	Code        string  `json:"code"`
	URI         *string `json:"uri,omitempty"`
	Description *string `json:"description,omitempty"`
	Type        string  `json:"type"`
}

// ConceptMapProperty represents properties for mapping
type ConceptMapProperty struct {
	common.BackboneElement

	Code        string  `json:"code"`
	URI         *string `json:"uri,omitempty"`
	Description *string `json:"description,omitempty"`
	Type        string  `json:"type"`
}

// ConceptMapGroupElementTargetProperty represents a property value for mapping
type ConceptMapGroupElementTargetProperty struct {
	common.BackboneElement

	Code          string         `json:"code"`
	ValueCoding   *common.Coding `json:"valueCoding,omitempty"`
	ValueString   *string        `json:"valueString,omitempty"`
	ValueInteger  *int           `json:"valueInteger,omitempty"`
	ValueBoolean  *bool          `json:"valueBoolean,omitempty"`
	ValueDateTime *string        `json:"valueDateTime,omitempty"`
	ValueDecimal  *float64       `json:"valueDecimal,omitempty"`
	ValueCode     *string        `json:"valueCode,omitempty"`
}

// ConceptMapGroupElementTargetDependsOn represents dependencies for mapping
type ConceptMapGroupElementTargetDependsOn struct {
	common.BackboneElement

	Attribute     string           `json:"attribute"`
	ValueCode     *string          `json:"valueCode,omitempty"`
	ValueCoding   *common.Coding   `json:"valueCoding,omitempty"`
	ValueString   *string          `json:"valueString,omitempty"`
	ValueBoolean  *bool            `json:"valueBoolean,omitempty"`
	ValueQuantity *common.Quantity `json:"valueQuantity,omitempty"`
	ValueSet      *string          `json:"valueSet,omitempty"`
}

// ConceptMapGroupElementTarget represents a target element in mapping
type ConceptMapGroupElementTarget struct {
	common.BackboneElement

	Code         *string                                 `json:"code,omitempty"`
	Comment      *string                                 `json:"comment,omitempty"`
	DependsOn    []ConceptMapGroupElementTargetDependsOn `json:"dependsOn,omitempty"`
	Display      *string                                 `json:"display,omitempty"`
	Product      []ConceptMapGroupElementTargetDependsOn `json:"product,omitempty"`
	Property     []ConceptMapGroupElementTargetProperty  `json:"property,omitempty"`
	Relationship ConceptMapRelationship                  `json:"relationship"`
	ValueSet     *string                                 `json:"valueSet,omitempty"`
}

// ConceptMapGroupElement represents a source element in mapping
type ConceptMapGroupElement struct {
	common.BackboneElement

	Code     *string                        `json:"code,omitempty"`
	Display  *string                        `json:"display,omitempty"`
	NoMap    *bool                          `json:"noMap,omitempty"`
	Target   []ConceptMapGroupElementTarget `json:"target,omitempty"`
	ValueSet *string                        `json:"valueSet,omitempty"`
}

// ConceptMapGroupUnmapped represents unmapped elements
type ConceptMapGroupUnmapped struct {
	common.BackboneElement

	Code         *string                 `json:"code,omitempty"`
	Display      *string                 `json:"display,omitempty"`
	Mode         ConceptMapUnmappedMode  `json:"mode"`
	OtherMap     *string                 `json:"otherMap,omitempty"`
	Relationship *ConceptMapRelationship `json:"relationship,omitempty"`
	ValueSet     *string                 `json:"valueSet,omitempty"`
}

// ConceptMapGroup represents a group of mappings
type ConceptMapGroup struct {
	common.BackboneElement

	Element  []ConceptMapGroupElement `json:"element"`
	Source   *string                  `json:"source,omitempty"`
	Target   *string                  `json:"target,omitempty"`
	Unmapped *ConceptMapGroupUnmapped `json:"unmapped,omitempty"`
}

// ConceptMap represents a statement of relationships from one set of concepts to another
type ConceptMap struct {
	DomainResource

	ResourceType string `json:"resourceType"` // Always "ConceptMap"

	AdditionalAttribute    []ConceptMapAdditionalAttribute `json:"additionalAttribute,omitempty"`
	ApprovalDate           *string                         `json:"approvalDate,omitempty"`
	Author                 []ContactDetail                 `json:"author,omitempty"`
	Contact                []ContactDetail                 `json:"contact,omitempty"`
	Copyright              *string                         `json:"copyright,omitempty"`
	CopyrightLabel         *string                         `json:"copyrightLabel,omitempty"`
	Date                   *string                         `json:"date,omitempty"`
	Description            *string                         `json:"description,omitempty"`
	Editor                 []ContactDetail                 `json:"editor,omitempty"`
	EffectivePeriod        *common.Period                  `json:"effectivePeriod,omitempty"`
	Endorser               []ContactDetail                 `json:"endorser,omitempty"`
	Experimental           *bool                           `json:"experimental,omitempty"`
	Group                  []ConceptMapGroup               `json:"group,omitempty"`
	Identifier             []common.Identifier             `json:"identifier,omitempty"`
	Jurisdiction           []common.CodeableConcept        `json:"jurisdiction,omitempty"`
	LastReviewDate         *string                         `json:"lastReviewDate,omitempty"`
	Name                   *string                         `json:"name,omitempty"`
	Property               []ConceptMapProperty            `json:"property,omitempty"`
	Publisher              *string                         `json:"publisher,omitempty"`
	Purpose                *string                         `json:"purpose,omitempty"`
	RelatedArtifact        []RelatedArtifact               `json:"relatedArtifact,omitempty"`
	Reviewer               []ContactDetail                 `json:"reviewer,omitempty"`
	SourceScopeUri         *string                         `json:"sourceScopeUri,omitempty"`
	SourceScopeCanonical   *string                         `json:"sourceScopeCanonical,omitempty"`
	Status                 PublicationStatus               `json:"status"`
	TargetScopeUri         *string                         `json:"targetScopeUri,omitempty"`
	TargetScopeCanonical   *string                         `json:"targetScopeCanonical,omitempty"`
	Title                  *string                         `json:"title,omitempty"`
	Topic                  []common.CodeableConcept        `json:"topic,omitempty"`
	URL                    *string                         `json:"url,omitempty"`
	UseContext             []UsageContext                  `json:"useContext,omitempty"`
	Version                *string                         `json:"version,omitempty"`
	VersionAlgorithmString *string                         `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding                  `json:"versionAlgorithmCoding,omitempty"`
}

// ConceptMapRelationship represents the relationship between concepts
type ConceptMapRelationship string

const (
	ConceptMapRelationshipRelatedTo                  ConceptMapRelationship = "related-to"
	ConceptMapRelationshipEquivalent                 ConceptMapRelationship = "equivalent"
	ConceptMapRelationshipSourceIsNarrowerThanTarget ConceptMapRelationship = "source-is-narrower-than-target"
	ConceptMapRelationshipSourceIsBroaderThanTarget  ConceptMapRelationship = "source-is-broader-than-target"
	ConceptMapRelationshipNotRelatedTo               ConceptMapRelationship = "not-related-to"
)

// ConceptMapUnmappedMode represents the mode for unmapped elements
type ConceptMapUnmappedMode string

const (
	ConceptMapUnmappedModeUseSourceCode ConceptMapUnmappedMode = "use-source-code"
	ConceptMapUnmappedModeFixed         ConceptMapUnmappedMode = "fixed"
	ConceptMapUnmappedModeOtherMap      ConceptMapUnmappedMode = "other-map"
)
