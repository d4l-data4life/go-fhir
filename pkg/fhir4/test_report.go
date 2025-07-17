package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// TestReport represents a summary of information based on the results of executing a TestScript
type TestReport struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "TestReport"

	// Identifier for the TestScript assigned for external purposes outside the context of FHIR
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Additional specific dates may be added as extensions
	Issued *string `json:"issued,omitempty"`

	// Not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// A participant in the test execution, either the execution engine, a client, or a server
	Participant []TestReportParticipant `json:"participant,omitempty"`

	// The pass and fail result represents a completed test script execution. The pending result represents a test script execution that has not yet started or is currently in progress
	Result string `json:"result"` // "pass" | "fail" | "pending"

	// The final score (percentage of tests passed) resulting from the execution of the TestScript
	Score *float64 `json:"score,omitempty"`

	// The results of the series of required setup operations before the tests were executed
	Setup *TestReportSetup `json:"setup,omitempty"`

	// The status represents where the execution is currently within the test script execution life cycle
	Status string `json:"status"` // "completed" | "in-progress" | "waiting" | "stopped" | "entered-in-error"

	// The results of the series of operations required to clean up after all the tests were executed (successfully or otherwise)
	Teardown *TestReportTeardown `json:"teardown,omitempty"`

	// A test executed from the test script
	Test []TestReportTest `json:"test,omitempty"`

	// Usually an organization, but may be an individual. This item SHOULD be populated unless the information is available from context
	Tester *string `json:"tester,omitempty"`

	// Ideally this is an absolute URL that is used to identify the version-specific TestScript that was executed, matching the TestScript.url
	TestScript common.Reference `json:"testScript"`
}

// TestReportParticipant represents a participant in the test execution, either the execution engine, a client, or a server
type TestReportParticipant struct {
	common.BackboneElement

	// The display name of the participant
	Display *string `json:"display,omitempty"`

	// The type of participant
	Type string `json:"type"` // "test-engine" | "client" | "server"

	// The uri of the participant. An absolute URL is preferred
	Uri string `json:"uri"`
}

// TestReportSetupActionOperation represents the operation performed
type TestReportSetupActionOperation struct {
	common.BackboneElement

	// A link to further details on the result
	Detail *string `json:"detail,omitempty"`

	// An explanatory message associated with the result
	Message *string `json:"message,omitempty"`

	// The result of this operation
	Result string `json:"result"` // "pass" | "skip" | "fail" | "warning" | "error"
}

// TestReportSetupActionAssert represents the results of the assertion performed on the previous operations
type TestReportSetupActionAssert struct {
	common.BackboneElement

	// A link to further details on the result
	Detail *string `json:"detail,omitempty"`

	// An explanatory message associated with the result
	Message *string `json:"message,omitempty"`

	// The result of this assertion
	Result string `json:"result"` // "pass" | "skip" | "fail" | "warning" | "error"
}

// TestReportSetupAction represents an action should contain either an operation or an assertion but not both
type TestReportSetupAction struct {
	common.BackboneElement

	// The results of the assertion performed on the previous operations
	Assert *TestReportSetupActionAssert `json:"assert,omitempty"`

	// The operation performed
	Operation *TestReportSetupActionOperation `json:"operation,omitempty"`
}

// TestReportSetup represents the results of the series of required setup operations before the tests were executed
type TestReportSetup struct {
	common.BackboneElement

	// An action should contain either an operation or an assertion but not both. It can contain any number of variables
	Action []TestReportSetupAction `json:"action"`
}

// TestReportTest represents a test executed from the test script
type TestReportTest struct {
	common.BackboneElement

	// A test operation or assert that was performed
	Action []TestReportTestAction `json:"action"`

	// A description of the test's purpose
	Description *string `json:"description,omitempty"`

	// The name of this test used for tracking/logging purposes by test engines
	Name *string `json:"name,omitempty"`
}

// TestReportTestAction represents a test operation or assert that was performed
type TestReportTestAction struct {
	common.BackboneElement

	// The results of the assertion performed on the previous operations
	Assert *TestReportSetupActionAssert `json:"assert,omitempty"`

	// The operation performed
	Operation *TestReportSetupActionOperation `json:"operation,omitempty"`
}

// TestReportTeardown represents the results of the series of operations required to clean up after all the tests were executed
type TestReportTeardown struct {
	common.BackboneElement

	// The teardown action will only contain an operation
	Action []TestReportTeardownAction `json:"action"`
}

// TestReportTeardownAction represents the teardown action will only contain an operation
type TestReportTeardownAction struct {
	common.BackboneElement

	// The teardown operation performed
	Operation TestReportSetupActionOperation `json:"operation"`
}
