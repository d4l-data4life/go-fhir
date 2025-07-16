// Package fhir5 contains FHIR R5 resource definitions
package fhir5

// QuestionnaireResponseStatus represents the status of a questionnaire response (FHIR R5)
type QuestionnaireResponseStatus string

const (
	QuestionnaireResponseStatusInProgress     QuestionnaireResponseStatus = "in-progress"
	QuestionnaireResponseStatusCompleted      QuestionnaireResponseStatus = "completed"
	QuestionnaireResponseStatusAmended        QuestionnaireResponseStatus = "amended"
	QuestionnaireResponseStatusEnteredInError QuestionnaireResponseStatus = "entered-in-error"
	QuestionnaireResponseStatusStopped        QuestionnaireResponseStatus = "stopped"
)
