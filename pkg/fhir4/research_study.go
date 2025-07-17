package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ResearchStudy represents a process where a researcher or organization plans and then executes a series of steps intended to increase the field of healthcare-related knowledge
type ResearchStudy struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ResearchStudy"

	// Describes an expected sequence of events for one of the participants of a study
	Arm []ResearchStudyArm `json:"arm,omitempty"`

	// Codes categorizing the type of study such as investigational vs. observational, type of blinding, type of randomization, safety vs. efficacy, etc
	Category []common.CodeableConcept `json:"category,omitempty"`

	// The condition that is the focus of the study
	Condition []common.CodeableConcept `json:"condition,omitempty"`

	// Contact details to assist a user in learning more about or engaging with the study
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A full description of how the study is being conducted
	Description *string `json:"description,omitempty"`

	// The Group referenced should not generally enumerate specific subjects
	Enrollment []common.Reference `json:"enrollment,omitempty"`

	// The medication(s), food(s), therapy(ies), device(s) or other concerns or interventions that the study is seeking to gain more information about
	Focus []common.CodeableConcept `json:"focus,omitempty"`

	// Identifiers assigned to this research study by the sponsor or other systems
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Key terms to aid in searching for or filtering the study
	Keyword []common.CodeableConcept `json:"keyword,omitempty"`

	// Indicates a country, state or other region where the study is taking place
	Location []common.CodeableConcept `json:"location,omitempty"`

	// Comments made about the study by the performer, subject or other participants
	Note []common.Annotation `json:"note,omitempty"`

	// A goal that the study is aiming to achieve in terms of a scientific question to be answered by the analysis of data collected during the study
	Objective []ResearchStudyObjective `json:"objective,omitempty"`

	// A larger research study of which this particular study is a component or step
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Identifies the start date and the expected (or actual, depending on status) end date for the study
	Period *common.Period `json:"period,omitempty"`

	// The stage in the progression of a therapy from initial experimental use in humans in clinical trials to post-market evaluation
	Phase *common.CodeableConcept `json:"phase,omitempty"`

	// The type of study based upon the intent of the study's activities
	PrimaryPurposeType *common.CodeableConcept `json:"primaryPurposeType,omitempty"`

	// A researcher in a study who oversees multiple aspects of the study
	PrincipalInvestigator *common.Reference `json:"principalInvestigator,omitempty"`

	// The set of steps expected to be performed as part of the execution of the study
	Protocol []common.Reference `json:"protocol,omitempty"`

	// A description and/or code explaining the premature termination of the study
	ReasonStopped *common.CodeableConcept `json:"reasonStopped,omitempty"`

	// Citations, references and other related documents
	RelatedArtifact []common.RelatedArtifact `json:"relatedArtifact,omitempty"`

	// A facility in which study activities are conducted
	Site []common.Reference `json:"site,omitempty"`

	// An organization that initiates the investigation and is legally responsible for the study
	Sponsor *common.Reference `json:"sponsor,omitempty"`

	// The current state of the study
	Status ResearchStudyStatus `json:"status"`

	// A short, descriptive user-friendly label for the study
	Title *string `json:"title,omitempty"`
}

// ResearchStudyStatus represents the current state of the study
type ResearchStudyStatus string

const (
	ResearchStudyStatusActive                                    ResearchStudyStatus = "active"
	ResearchStudyStatusAdministrativelyCompleted                 ResearchStudyStatus = "administratively-completed"
	ResearchStudyStatusApproved                                  ResearchStudyStatus = "approved"
	ResearchStudyStatusClosedToAccrual                           ResearchStudyStatus = "closed-to-accrual"
	ResearchStudyStatusClosedToAccrualAndIntervention            ResearchStudyStatus = "closed-to-accrual-and-intervention"
	ResearchStudyStatusCompleted                                 ResearchStudyStatus = "completed"
	ResearchStudyStatusDisapproved                               ResearchStudyStatus = "disapproved"
	ResearchStudyStatusInReview                                  ResearchStudyStatus = "in-review"
	ResearchStudyStatusTemporarilyClosedToAccrual                ResearchStudyStatus = "temporarily-closed-to-accrual"
	ResearchStudyStatusTemporarilyClosedToAccrualAndIntervention ResearchStudyStatus = "temporarily-closed-to-accrual-and-intervention"
	ResearchStudyStatusWithdrawn                                 ResearchStudyStatus = "withdrawn"
)

// ResearchStudyArm represents describes an expected sequence of events for one of the participants of a study
type ResearchStudyArm struct {
	common.BackboneElement

	// A succinct description of the path through the study that would be followed by a subject adhering to this arm
	Description *string `json:"description,omitempty"`

	// Unique, human-readable label for this arm of the study
	Name string `json:"name"`

	// Categorization of study arm, e.g. experimental, active comparator, placebo comparater
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ResearchStudyObjective represents a goal that the study is aiming to achieve in terms of a scientific question to be answered by the analysis of data collected during the study
type ResearchStudyObjective struct {
	common.BackboneElement

	// Unique, human-readable label for this objective of the study
	Name *string `json:"name,omitempty"`

	// The kind of study objective
	Type *common.CodeableConcept `json:"type,omitempty"`
}
