// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// ClinicalUseDefinitionContraindicationOtherTherapy represents information about the use of the medicinal product
type ClinicalUseDefinitionContraindicationOtherTherapy struct {
	common.BackboneElement

	// The type of relationship between the medicinal product indication or contraindication and another therapy
	RelationshipType common.CodeableConcept `json:"relationshipType"`

	// Reference to a specific medication (active substance, medicinal product or class of products, biological, food etc.)
	Treatment CodeableReference `json:"treatment"`
}

// ClinicalUseDefinitionContraindication represents specifics for when this is a contraindication
type ClinicalUseDefinitionContraindication struct {
	common.BackboneElement

	// The expression may be inlined or may be a reference to a named expression within a logic library
	Applicability *Expression `json:"applicability,omitempty"`

	// A comorbidity (concurrent condition) or coinfection
	Comorbidity []CodeableReference `json:"comorbidity,omitempty"`

	// The status of the disease or symptom for the contraindication, for example "chronic" or "metastatic"
	DiseaseStatus *CodeableReference `json:"diseaseStatus,omitempty"`

	// The situation that is being documented as contraindicating against this item
	DiseaseSymptomProcedure *CodeableReference `json:"diseaseSymptomProcedure,omitempty"`

	// The indication which this is a contraindication for
	Indication []common.Reference `json:"indication,omitempty"`

	// Information about the use of the medicinal product in relation to other therapies
	OtherTherapy []ClinicalUseDefinitionContraindicationOtherTherapy `json:"otherTherapy,omitempty"`
}

// ClinicalUseDefinitionIndication represents specifics for when this is an indication
type ClinicalUseDefinitionIndication struct {
	common.BackboneElement

	// The expression may be inlined or may be a reference to a named expression within a logic library
	Applicability *Expression `json:"applicability,omitempty"`

	// A comorbidity (concurrent condition) or coinfection as part of the indication
	Comorbidity []CodeableReference `json:"comorbidity,omitempty"`

	// The status of the disease or symptom for the indication, for example "chronic" or "metastatic"
	DiseaseStatus *CodeableReference `json:"diseaseStatus,omitempty"`

	// The situation that is being documented as an indication for this item
	DiseaseSymptomProcedure *CodeableReference `json:"diseaseSymptomProcedure,omitempty"`

	// Timing or duration information, that may be associated with use with the indicated condition
	DurationRange *Range `json:"durationRange,omitempty"`

	// Timing or duration information, that may be associated with use with the indicated condition
	DurationString *string `json:"durationString,omitempty"`

	// The intended effect, aim or strategy to be achieved
	IntendedEffect *CodeableReference `json:"intendedEffect,omitempty"`

	// Information about the use of the medicinal product in relation to other therapies
	OtherTherapy []ClinicalUseDefinitionContraindicationOtherTherapy `json:"otherTherapy,omitempty"`

	// An unwanted side effect or negative outcome that may happen if you use the drug
	UndesirableEffect []common.Reference `json:"undesirableEffect,omitempty"`
}

// ClinicalUseDefinitionInteractionInteractant represents the specific medication, product, food, substance etc.
type ClinicalUseDefinitionInteractionInteractant struct {
	common.BackboneElement

	// The specific medication, product, food, substance etc. or laboratory test that interacts
	ItemReference *common.Reference `json:"itemReference,omitempty"`

	// The specific medication, product, food, substance etc. or laboratory test that interacts
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

// ClinicalUseDefinitionInteraction represents specifics for when this is an interaction
type ClinicalUseDefinitionInteraction struct {
	common.BackboneElement

	// The effect of the interaction, for example "reduced gastric absorption of primary medication"
	Effect *CodeableReference `json:"effect,omitempty"`

	// The incidence of the interaction, e.g. theoretical, observed
	Incidence *common.CodeableConcept `json:"incidence,omitempty"`

	// The specific medication, product, food, substance etc. or laboratory test that interacts
	Interactant []ClinicalUseDefinitionInteractionInteractant `json:"interactant,omitempty"`

	// Actions for managing the interaction
	Management []common.CodeableConcept `json:"management,omitempty"`

	// The type of the interaction e.g. drug-drug interaction, drug-food interaction, drug-lab test interaction
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ClinicalUseDefinitionUndesirableEffect represents the possible undesirable effects
type ClinicalUseDefinitionUndesirableEffect struct {
	common.BackboneElement

	// High level classification of the effect
	Classification *common.CodeableConcept `json:"classification,omitempty"`

	// How often the effect is seen
	FrequencyOfOccurrence *common.CodeableConcept `json:"frequencyOfOccurrence,omitempty"`

	// The situation in which the undesirable effect may manifest
	SymptomConditionEffect *CodeableReference `json:"symptomConditionEffect,omitempty"`
}

// ClinicalUseDefinitionWarning represents a critical piece of information about environmental, health or physical risks
type ClinicalUseDefinitionWarning struct {
	common.BackboneElement

	// A coded or unformatted textual definition of this warning
	Code *common.CodeableConcept `json:"code,omitempty"`

	// A textual definition of this warning, with formatting
	Description *string `json:"description,omitempty"`
}

// ClinicalUseDefinition represents a single issue - either an indication, contraindication, interaction or an undesirable effect
type ClinicalUseDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ClinicalUseDefinition"

	// A categorisation of the issue, primarily for dividing warnings into subject heading areas
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Specifics for when this is a contraindication
	Contraindication *ClinicalUseDefinitionContraindication `json:"contraindication,omitempty"`

	// Business identifier for this issue
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Specifics for when this is an indication
	Indication *ClinicalUseDefinitionIndication `json:"indication,omitempty"`

	// Specifics for when this is an interaction
	Interaction *ClinicalUseDefinitionInteraction `json:"interaction,omitempty"`

	// Logic used by the clinical use definition
	Library []string `json:"library,omitempty"`

	// The population group to which this applies
	Population []common.Reference `json:"population,omitempty"`

	// Whether this is a current issue or one that has been retired etc
	Status *common.CodeableConcept `json:"status,omitempty"`

	// The medication, product, substance, device, procedure etc. for which this is an indication
	Subject []common.Reference `json:"subject,omitempty"`

	// indication | contraindication | interaction | undesirable-effect | warning
	Type ClinicalUseDefinitionType `json:"type"`

	// Describe the possible undesirable effects (negative outcomes) from the use of the medicinal product as treatment
	UndesirableEffect *ClinicalUseDefinitionUndesirableEffect `json:"undesirableEffect,omitempty"`

	// A critical piece of information about environmental, health or physical risks or hazards
	Warning *ClinicalUseDefinitionWarning `json:"warning,omitempty"`
}

// ClinicalUseDefinitionType represents the type of clinical use definition
type ClinicalUseDefinitionType string

const (
	ClinicalUseDefinitionTypeIndication        ClinicalUseDefinitionType = "indication"
	ClinicalUseDefinitionTypeContraindication  ClinicalUseDefinitionType = "contraindication"
	ClinicalUseDefinitionTypeInteraction       ClinicalUseDefinitionType = "interaction"
	ClinicalUseDefinitionTypeUndesirableEffect ClinicalUseDefinitionType = "undesirable-effect"
	ClinicalUseDefinitionTypeWarning           ClinicalUseDefinitionType = "warning"
)
