package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// TestReportParticipant represents a participant in the test execution
type TestReportParticipant struct {
	common.BackboneElement

	// The display name of the participant
	Display *string `json:"display,omitempty"`

	// The type of participant
	Type TestReportParticipantType `json:"type"`

	// The uri of the participant
	Uri string `json:"uri"`
}

// TestReportParticipantType represents the type of participant
type TestReportParticipantType string

const (
	TestReportParticipantTypeTestEngine TestReportParticipantType = "test-engine"
	TestReportParticipantTypeClient     TestReportParticipantType = "client"
	TestReportParticipantTypeServer     TestReportParticipantType = "server"
)

// TestReportSetupActionOperation represents the operation performed
type TestReportSetupActionOperation struct {
	common.BackboneElement

	// A link to further details on the result
	Detail *string `json:"detail,omitempty"`

	// An explanatory message associated with the result
	Message *string `json:"message,omitempty"`

	// The result of this operation
	Result TestReportActionResult `json:"result"`
}

// TestReportSetupActionAssertRequirement represents requirements for assertions
type TestReportSetupActionAssertRequirement struct {
	common.BackboneElement

	// Link or reference providing traceability to the testing requirement for this test
	LinkUri *string `json:"linkUri,omitempty"`

	// Link or reference providing traceability to the testing requirement for this test
	LinkCanonical *string `json:"linkCanonical,omitempty"`
}

// TestReportSetupActionAssert represents the results of the assertion performed
type TestReportSetupActionAssert struct {
	common.BackboneElement

	// A link to further details on the result
	Detail *string `json:"detail,omitempty"`

	// An explanatory message associated with the result
	Message *string `json:"message,omitempty"`

	// TestScript and TestReport instances are typically based on known, defined test requirements and documentation
	Requirement []TestReportSetupActionAssertRequirement `json:"requirement,omitempty"`

	// The result of this assertion
	Result TestReportActionResult `json:"result"`
}

// TestReportSetupAction represents a single action performed as part of the test setup
type TestReportSetupAction struct {
	common.BackboneElement

	// The operation performed
	Operation *TestReportSetupActionOperation `json:"operation,omitempty"`

	// The results of the assertion performed on the previous operations
	Assert *TestReportSetupActionAssert `json:"assert,omitempty"`
}

// TestReportSetup represents the results of the series of required setup operations before the tests were executed
type TestReportSetup struct {
	common.BackboneElement

	// Action would contain either an operation or an assertion
	Action []TestReportSetupAction `json:"action"`
}

// TestReportTestAction represents a single action performed as part of the test
type TestReportTestAction struct {
	common.BackboneElement

	// The operation performed
	Operation *TestReportSetupActionOperation `json:"operation,omitempty"`

	// The results of the assertion performed on the previous operations
	Assert *TestReportSetupActionAssert `json:"assert,omitempty"`
}

// TestReportTest represents a test executed from the test script
type TestReportTest struct {
	common.BackboneElement

	// The name of this test used for tracking/logging purposes by test engines
	Name *string `json:"name,omitempty"`

	// A short description of the test used by test engines for tracking and reporting purposes
	Description *string `json:"description,omitempty"`

	// Action would contain either an operation or an assertion
	Action []TestReportTestAction `json:"action"`
}

// TestReportTeardownAction represents a single action performed as part of the test teardown
type TestReportTeardownAction struct {
	common.BackboneElement

	// The teardown operation performed
	Operation TestReportSetupActionOperation `json:"operation"`
}

// TestReportTeardown represents the results of the series of operations required to clean up after all the tests were executed
type TestReportTeardown struct {
	common.BackboneElement

	// The teardown action will only contain an operation
	Action []TestReportTeardownAction `json:"action"`
}

// TestReportActionResult represents the result of a test action
type TestReportActionResult string

const (
	TestReportActionResultPass    TestReportActionResult = "pass"
	TestReportActionResultSkip    TestReportActionResult = "skip"
	TestReportActionResultFail    TestReportActionResult = "fail"
	TestReportActionResultWarning TestReportActionResult = "warning"
	TestReportActionResultError   TestReportActionResult = "error"
)

// TestReportResult represents the overall result of the test execution
type TestReportResult string

const (
	TestReportResultPass    TestReportResult = "pass"
	TestReportResultFail    TestReportResult = "fail"
	TestReportResultPending TestReportResult = "pending"
)

// TestReportStatus represents the status of the test report
type TestReportStatus string

const (
	TestReportStatusCompleted      TestReportStatus = "completed"
	TestReportStatusInProgress     TestReportStatus = "in-progress"
	TestReportStatusWaiting        TestReportStatus = "waiting"
	TestReportStatusStopped        TestReportStatus = "stopped"
	TestReportStatusEnteredInError TestReportStatus = "entered-in-error"
)

// TestReport represents a summary of information based on the results of executing a TestScript
type TestReport struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "TestReport"

	// Identifier for the TestReport assigned for external purposes outside the context of FHIR
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Additional specific dates may be added as extensions
	Issued *string `json:"issued,omitempty"`

	// Not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// A participant in the test execution, either the execution engine, a client, or a server
	Participant []TestReportParticipant `json:"participant,omitempty"`

	// The pass and fail result represents a completed test script execution
	Result TestReportResult `json:"result"`

	// The final score (percentage of tests passed) resulting from the execution of the TestScript
	Score *float64 `json:"score,omitempty"`

	// The results of the series of required setup operations before the tests were executed
	Setup *TestReportSetup `json:"setup,omitempty"`

	// The status represents where the execution is currently within the test script execution life cycle
	Status TestReportStatus `json:"status"`

	// The results of the series of operations required to clean up after all the tests were executed
	Teardown *TestReportTeardown `json:"teardown,omitempty"`

	// A test executed from the test script
	Test []TestReportTest `json:"test,omitempty"`

	// Usually an organization, but may be an individual
	Tester *string `json:"tester,omitempty"`

	// Ideally this is an absolute URL that is used to identify the version-specific TestScript that was executed
	TestScript string `json:"testScript"`
}
