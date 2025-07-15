// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// GroupType represents the type of group
type GroupType string

const (
	GroupTypePerson                      GroupType = "person"
	GroupTypeAnimal                      GroupType = "animal"
	GroupTypePractitioner                GroupType = "practitioner"
	GroupTypeDevice                      GroupType = "device"
	GroupTypeMedication                  GroupType = "medication"
	GroupTypeSubstance                   GroupType = "substance"
	GroupTypeOrganization                GroupType = "organization"
	GroupTypeLocation                    GroupType = "location"
	GroupTypePlanDefinition              GroupType = "plandefinition"
	GroupTypeHealthcareService           GroupType = "healthcareservice"
	GroupTypeNutritionProduct            GroupType = "nutritionproduct"
	GroupTypeBiologicallyDerivedProduct  GroupType = "biologicallyderivedproduct"
	GroupTypeSpecimen                    GroupType = "specimen"
	GroupTypeObservation                 GroupType = "observation"
	GroupTypeProcedure                   GroupType = "procedure"
	GroupTypeCarePlan                    GroupType = "careplan"
	GroupTypeCareTeam                    GroupType = "careteam"
	GroupTypeGoal                        GroupType = "goal"
	GroupTypeConsent                     GroupType = "consent"
	GroupTypeContract                    GroupType = "contract"
	GroupTypeCoverage                    GroupType = "coverage"
	GroupTypeCoverageEligibilityRequest  GroupType = "coverageeligibilityrequest"
	GroupTypeCoverageEligibilityResponse GroupType = "coverageeligibilityresponse"
	GroupTypeEnrollmentRequest           GroupType = "enrollmentrequest"
	GroupTypeEnrollmentResponse          GroupType = "enrollmentresponse"
	GroupTypeClaim                       GroupType = "claim"
	GroupTypeClaimResponse               GroupType = "claimresponse"
	GroupTypeInvoice                     GroupType = "invoice"
	GroupTypePaymentNotice               GroupType = "paymentnotice"
	GroupTypePaymentReconciliation       GroupType = "paymentreconciliation"
	GroupTypeChargeItem                  GroupType = "chargeitem"
	GroupTypeChargeItemDefinition        GroupType = "chargeitemdefinition"
	GroupTypeResearchStudy               GroupType = "researchstudy"
	GroupTypeResearchSubject             GroupType = "researchsubject"
	GroupTypeEncounter                   GroupType = "encounter"
	GroupTypeEpisodeOfCare               GroupType = "episodeofcare"
	GroupTypeCondition                   GroupType = "condition"
	GroupTypeAllergyIntolerance          GroupType = "allergyintolerance"
	GroupTypeFamilyMemberHistory         GroupType = "familymemberhistory"
	GroupTypeDetectedIssue               GroupType = "detectedissue"
	GroupTypeRiskAssessment              GroupType = "riskassessment"
	GroupTypeImmunization                GroupType = "immunization"
	GroupTypeImmunizationEvaluation      GroupType = "immunizationevaluation"
	GroupTypeImmunizationRecommendation  GroupType = "immunizationrecommendation"
	GroupTypeMedicationRequest           GroupType = "medicationrequest"
	GroupTypeMedicationDispense          GroupType = "medicationdispense"
	GroupTypeMedicationAdministration    GroupType = "medicationadministration"
	GroupTypeMedicationStatement         GroupType = "medicationstatement"
	GroupTypeMedicationKnowledge         GroupType = "medicationknowledge"
	GroupTypeMedicationUsage             GroupType = "medicationusage"
	GroupTypeCommunication               GroupType = "communication"
	GroupTypeCommunicationRequest        GroupType = "communicationrequest"
	GroupTypeDeviceRequest               GroupType = "devicerequest"
	GroupTypeDeviceUsage                 GroupType = "deviceusage"
	GroupTypeDeviceMetric                GroupType = "devicemetric"
	GroupTypeDeviceDefinition            GroupType = "devicedefinition"
	GroupTypeNutritionOrder              GroupType = "nutritionorder"
	GroupTypeNutritionIntake             GroupType = "nutritionintake"
	GroupTypeVisionPrescription          GroupType = "visionprescription"
	GroupTypeServiceRequest              GroupType = "servicerequest"
	GroupTypeSupplyRequest               GroupType = "supplyrequest"
	GroupTypeSupplyDelivery              GroupType = "supplydelivery"
	GroupTypeAppointment                 GroupType = "appointment"
	GroupTypeAppointmentResponse         GroupType = "appointmentresponse"
	GroupTypeSchedule                    GroupType = "schedule"
	GroupTypeSlot                        GroupType = "slot"
	GroupTypeVerificationResult          GroupType = "verificationresult"
	GroupTypeQuestionnaire               GroupType = "questionnaire"
	GroupTypeQuestionnaireResponse       GroupType = "questionnaireresponse"
	GroupTypeActivityDefinition          GroupType = "activitydefinition"
	GroupTypeTask                        GroupType = "task"
	GroupTypeRequestOrchestration        GroupType = "requestorchestration"
	GroupTypeLibrary                     GroupType = "library"
	GroupTypeMeasure                     GroupType = "measure"
	GroupTypeMeasureReport               GroupType = "measurereport"
	GroupTypeEvidence                    GroupType = "evidence"
	GroupTypeEvidenceReport              GroupType = "evidencereport"
	GroupTypeEvidenceVariable            GroupType = "evidencevariable"
	GroupTypeResearchDefinition          GroupType = "researchdefinition"
	GroupTypeResearchElementDefinition   GroupType = "researchelementdefinition"
	GroupTypeGroup                       GroupType = "group"
)

// GroupMembership represents the basis for membership in the Group
type GroupMembership string

const (
	GroupMembershipDefinitional GroupMembership = "definitional"
	GroupMembershipEnumerated   GroupMembership = "enumerated"
)

// GroupCharacteristic represents a code that identifies the kind of trait being asserted
type GroupCharacteristic struct {
	common.BackboneElement

	// A code that identifies the kind of trait being asserted
	Code common.CodeableConcept `json:"code"`

	// This is labeled as "Is Modifier" because applications cannot wrongly include excluded members
	Exclude bool `json:"exclude"`

	// The period over which the characteristic is tested
	Period *common.Period `json:"period,omitempty"`

	// For Range, it means members of the group have a value that falls somewhere within the specified range
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
}

// GroupMember represents a reference to the entity that is a member of the group
type GroupMember struct {
	common.BackboneElement

	// A reference to the entity that is a member of the group
	Entity common.Reference `json:"entity"`

	// A flag to indicate that the member is no longer in the group
	Inactive *bool `json:"inactive,omitempty"`

	// The period that the member was in the group, if known
	Period *common.Period `json:"period,omitempty"`
}

// Group represents a group of related resources
type Group struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Group"

	// Indicates whether the record for the group is available for use or is merely being retained for historical purposes
	Active *bool `json:"active,omitempty"`

	// All the identified characteristics must be true for an entity to a member of the group
	Characteristic []GroupCharacteristic `json:"characteristic,omitempty"`

	// This would generally be omitted for Person resources
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Explanation of what the group represents and how it is intended to be used
	Description *string `json:"description,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// This does not strictly align with ownership of a herd or flock, but may suffice to represent that relationship
	ManagingEntity *common.Reference `json:"managingEntity,omitempty"`

	// Identifies the resource instances that are members of the group
	Member []GroupMember `json:"member,omitempty"`

	// Basis for membership in the Group
	Membership GroupMembership `json:"membership"`

	// A label assigned to the group for human identification and communication
	Name *string `json:"name,omitempty"`

	// Note that the quantity may be less than the number of members if some of the members are not active
	Quantity *int `json:"quantity,omitempty"`

	// Group members SHALL be of the appropriate resource type
	Type GroupType `json:"type"`
}
