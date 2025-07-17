package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// TestScriptOrigin represents the purpose of this element is to define the profile of an origin element used elsewhere in the script
type TestScriptOrigin struct {
	common.BackboneElement

	// A given origin index can appear only once in the list
	Index int `json:"index"`

	// Must be a "sender"/"client" profile
	Profile common.Coding `json:"profile"`

	// If provided, the test engine is not expected to prompt for or accept external input of this value
	Url *string `json:"url,omitempty"`
}

// TestScriptDestination represents the purpose of this element is to define the profile of a destination element used elsewhere in the script
type TestScriptDestination struct {
	common.BackboneElement

	// A given destination index can appear only once in the list
	Index int `json:"index"`

	// Must be a "receiver"/"server" profile
	Profile common.Coding `json:"profile"`

	// If provided, the test engine is not expected to prompt for or accept external input of this value
	Url *string `json:"url,omitempty"`
}

// TestScriptMetadataLink represents a link to the FHIR specification that this test is covering
type TestScriptMetadataLink struct {
	common.BackboneElement

	// Short description of the link
	Description *string `json:"description,omitempty"`

	// URL to a particular requirement or feature within the FHIR specification
	Url string `json:"url"`
}

// TestScriptMetadataCapability represents the required capability must exist and are assumed to function correctly on the FHIR server being tested
type TestScriptMetadataCapability struct {
	common.BackboneElement

	// The conformance statement of the server has to contain at a minimum the contents of the reference pointed to by this element
	Capabilities string `json:"capabilities"`

	// Description of the capabilities that this test script is requiring the server to support
	Description *string `json:"description,omitempty"`

	// Which server these requirements apply to
	Destination *int `json:"destination,omitempty"`

	// Links to the FHIR specification that describes this interaction and the resources involved in more detail
	Link []string `json:"link,omitempty"`

	// Which origin server these requirements apply to
	Origin []int `json:"origin,omitempty"`

	// Whether or not the test execution will require the given capabilities of the server in order for this test script to execute
	Required bool `json:"required"`

	// Whether or not the test execution will validate the given capabilities of the server in order for this test script to execute
	Validated bool `json:"validated"`
}

// TestScriptMetadata represents the required capability must exist and are assumed to function correctly on the FHIR server being tested
type TestScriptMetadata struct {
	common.BackboneElement

	// The required capability must exist and are assumed to function correctly on the FHIR server being tested
	Capability []TestScriptMetadataCapability `json:"capability"`

	// A link to the FHIR specification that this test is covering
	Link []TestScriptMetadataLink `json:"link,omitempty"`
}

// TestScriptScope represents the scope indicates a conformance artifact that is tested by the test(s) within this test case
type TestScriptScope struct {
	common.BackboneElement

	// The specific conformance artifact being tested
	Artifact string `json:"artifact"`

	// The expectation of whether the test must pass for the system to be considered conformant with the artifact
	Conformance *common.CodeableConcept `json:"conformance,omitempty"`

	// The phase of testing for this artifact
	Phase *common.CodeableConcept `json:"phase,omitempty"`
}

// TestScriptFixture represents fixture in the test script - by reference (uri)
type TestScriptFixture struct {
	common.BackboneElement

	// Whether or not to implicitly create the fixture during setup
	Autocreate bool `json:"autocreate"`

	// Whether or not to implicitly delete the fixture during teardown
	Autodelete bool `json:"autodelete"`

	// Reference to the resource (contained in the test script)
	Resource *common.Reference `json:"resource,omitempty"`
}

// TestScriptVariable represents variables would be set based either on XPath/JSONPath expressions against fixtures
type TestScriptVariable struct {
	common.BackboneElement

	// Descriptive name for this variable
	Name string `json:"name"`

	// Default, hard-coded, or user-defined value for this variable
	DefaultValue *string `json:"defaultValue,omitempty"`

	// A free-text natural language description of the variable and its purpose
	Description *string `json:"description,omitempty"`

	// The FHIRPath expression against the fixture body
	Expression *string `json:"expression,omitempty"`

	// HTTP header field name for source extraction
	HeaderField *string `json:"headerField,omitempty"`

	// Hint help text for default value to enter
	Hint *string `json:"hint,omitempty"`

	// XPath or JSONPath against the fixture body
	Path *string `json:"path,omitempty"`

	// Source of the value for this variable
	SourceID *string `json:"sourceId,omitempty"`
}

// TestScriptSetupActionOperation represents the operation to perform
type TestScriptSetupActionOperation struct {
	common.BackboneElement

	// The type of operation to perform
	Type *common.Coding `json:"type,omitempty"`

	// Resource type
	Resource *string `json:"resource,omitempty"`

	// Label would be used for tracking/logging purposes by test engines
	Label *string `json:"label,omitempty"`

	// The description would be used by test engines for tracking and reporting purposes
	Description *string `json:"description,omitempty"`

	// The mime-type to use for RESTful operations in the 'Accept' header
	Accept *string `json:"accept,omitempty"`

	// The mime-type to use for RESTful operations in the 'Content-Type' header
	ContentType *string `json:"contentType,omitempty"`

	// The server where the request message originates from
	Destination *int `json:"destination,omitempty"`

	// Whether or not to implicitly send the request url in encoded format
	EncodeRequestUrl *bool `json:"encodeRequestUrl,omitempty"`

	// The server where the request message is destined for
	Origin *int `json:"origin,omitempty"`

	// Path plus parameters after [type]
	Params *string `json:"params,omitempty"`

	// Explicitly defined path parameters
	RequestHeader []TestScriptSetupActionOperationRequestHeader `json:"requestHeader,omitempty"`

	// The fixture id (maybe new) for the request body
	RequestID *string `json:"requestId,omitempty"`

	// The type of the resource
	ResponseID *string `json:"responseId,omitempty"`

	// The fixture id for the response body
	SourceID *string `json:"sourceId,omitempty"`

	// Id of fixture used for extracting the [id],  [type], and [vid] for GET requests
	TargetID *string `json:"targetId,omitempty"`

	// Complete request URL
	Url *string `json:"url,omitempty"`
}

// TestScriptSetupActionOperationRequestHeader represents explicitly defined path parameters
type TestScriptSetupActionOperationRequestHeader struct {
	common.BackboneElement

	// HTTP header field name
	Field string `json:"field"`

	// HTTP header field value
	Value string `json:"value"`
}

// TestScriptSetupActionAssert represents the assertion to perform
type TestScriptSetupActionAssert struct {
	common.BackboneElement

	// The label would be used for tracking/logging purposes by test engines
	Label *string `json:"label,omitempty"`

	// The description would be used by test engines for tracking and reporting purposes
	Description *string `json:"description,omitempty"`

	// The direction to use for the assertion
	Direction *TestScriptSetupActionAssertDirection `json:"direction,omitempty"`

	// Id of the source fixture used as the contents to be evaluated by either the "source/expression" or "sourceId/path" definition
	CompareToSourceID *string `json:"compareToSourceId,omitempty"`

	// The fluentpath expression to evaluate against the source fixture
	CompareToSourceExpression *string `json:"compareToSourceExpression,omitempty"`

	// XPath or JSONPath expression to evaluate against the source fixture
	CompareToSourcePath *string `json:"compareToSourcePath,omitempty"`

	// The content-type or mime-type to use for RESTful operations in the 'Content-Type' header
	ContentType *string `json:"contentType,omitempty"`

	// The fluentpath expression to be evaluated
	Expression *string `json:"expression,omitempty"`

	// HTTP header field name
	HeaderField *string `json:"headerField,omitempty"`

	// The ID of a fixture
	MinimumID *string `json:"minimumId,omitempty"`

	// Whether or not the test execution performs validation on the specified navigation links
	NavigationLinks *bool `json:"navigationLinks,omitempty"`

	// The operator type defines the conditional behavior of the assert
	Operator *TestScriptSetupActionAssertOperator `json:"operator,omitempty"`

	// XPath or JSONPath expression
	Path *string `json:"path,omitempty"`

	// The request method or HTTP operation code to compare against that used by the client system under test
	RequestMethod *string `json:"requestMethod,omitempty"`

	// The value to compare to
	RequestURL *string `json:"requestURL,omitempty"`

	// The type of the resource
	Resource *string `json:"resource,omitempty"`

	// okay | created | noContent | notModified | bad | forbidden | notFound | methodNotAllowed | conflict | gone | preconditionFailed | unprocessable
	Response *TestScriptSetupActionAssertResponse `json:"response,omitempty"`

	// The value of the HTTP response code to be tested
	ResponseCode *string `json:"responseCode,omitempty"`

	// The TestScript.rule this assert will evaluate
	Rule *TestScriptSetupActionAssertRule `json:"rule,omitempty"`

	// The TestScript.ruleset this assert will evaluate
	Ruleset *TestScriptSetupActionAssertRuleset `json:"ruleset,omitempty"`

	// Fixture to evaluate the XPath/JSONPath expression or the headerField against
	SourceID *string `json:"sourceId,omitempty"`

	// The ID of the Profile to validate against
	ValidateProfileID *string `json:"validateProfileId,omitempty"`

	// The value to compare to
	Value *string `json:"value,omitempty"`

	// Will this assert produce a warning only on error
	WarningOnly *bool `json:"warningOnly,omitempty"`
}

// TestScriptSetupActionAssertDirection represents the direction to use for the assertion
type TestScriptSetupActionAssertDirection string

const (
	TestScriptSetupActionAssertDirectionResponse TestScriptSetupActionAssertDirection = "response"
	TestScriptSetupActionAssertDirectionRequest  TestScriptSetupActionAssertDirection = "request"
)

// TestScriptSetupActionAssertOperator represents the operator type defines the conditional behavior of the assert
type TestScriptSetupActionAssertOperator string

const (
	TestScriptSetupActionAssertOperatorEquals      TestScriptSetupActionAssertOperator = "equals"
	TestScriptSetupActionAssertOperatorNotEquals   TestScriptSetupActionAssertOperator = "notEquals"
	TestScriptSetupActionAssertOperatorIn          TestScriptSetupActionAssertOperator = "in"
	TestScriptSetupActionAssertOperatorNotIn       TestScriptSetupActionAssertOperator = "notIn"
	TestScriptSetupActionAssertOperatorGreaterThan TestScriptSetupActionAssertOperator = "greaterThan"
	TestScriptSetupActionAssertOperatorLessThan    TestScriptSetupActionAssertOperator = "lessThan"
	TestScriptSetupActionAssertOperatorEmpty       TestScriptSetupActionAssertOperator = "empty"
	TestScriptSetupActionAssertOperatorNotEmpty    TestScriptSetupActionAssertOperator = "notEmpty"
	TestScriptSetupActionAssertOperatorContains    TestScriptSetupActionAssertOperator = "contains"
	TestScriptSetupActionAssertOperatorNotContains TestScriptSetupActionAssertOperator = "notContains"
	TestScriptSetupActionAssertOperatorEval        TestScriptSetupActionAssertOperator = "eval"
)

// TestScriptSetupActionAssertResponse represents the response type
type TestScriptSetupActionAssertResponse string

const (
	TestScriptSetupActionAssertResponseOkay               TestScriptSetupActionAssertResponse = "okay"
	TestScriptSetupActionAssertResponseCreated            TestScriptSetupActionAssertResponse = "created"
	TestScriptSetupActionAssertResponseNoContent          TestScriptSetupActionAssertResponse = "noContent"
	TestScriptSetupActionAssertResponseNotModified        TestScriptSetupActionAssertResponse = "notModified"
	TestScriptSetupActionAssertResponseBad                TestScriptSetupActionAssertResponse = "bad"
	TestScriptSetupActionAssertResponseForbidden          TestScriptSetupActionAssertResponse = "forbidden"
	TestScriptSetupActionAssertResponseNotFound           TestScriptSetupActionAssertResponse = "notFound"
	TestScriptSetupActionAssertResponseMethodNotAllowed   TestScriptSetupActionAssertResponse = "methodNotAllowed"
	TestScriptSetupActionAssertResponseConflict           TestScriptSetupActionAssertResponse = "conflict"
	TestScriptSetupActionAssertResponseGone               TestScriptSetupActionAssertResponse = "gone"
	TestScriptSetupActionAssertResponsePreconditionFailed TestScriptSetupActionAssertResponse = "preconditionFailed"
	TestScriptSetupActionAssertResponseUnprocessable      TestScriptSetupActionAssertResponse = "unprocessable"
)

// TestScriptSetupActionAssertRule represents the TestScript.rule this assert will evaluate
type TestScriptSetupActionAssertRule struct {
	common.BackboneElement

	// The TestScript.rule id value this assert will evaluate
	RuleID string `json:"ruleId"`

	// The referenced rule within the external ruleset template
	Param []TestScriptSetupActionAssertRuleParam `json:"param,omitempty"`
}

// TestScriptSetupActionAssertRuleParam represents the referenced rule within the external ruleset template
type TestScriptSetupActionAssertRuleParam struct {
	common.BackboneElement

	// Parameter name matching external assert rule parameter
	Name string `json:"name"`

	// Parameter value defined either explicitly or dynamically
	Value *string `json:"value,omitempty"`
}

// TestScriptSetupActionAssertRuleset represents the TestScript.ruleset this assert will evaluate
type TestScriptSetupActionAssertRuleset struct {
	common.BackboneElement

	// The TestScript.ruleset id value this assert will evaluate
	RulesetID string `json:"rulesetId"`

	// The referenced rule within the external ruleset template
	Rule []TestScriptSetupActionAssertRulesetRule `json:"rule,omitempty"`
}

// TestScriptSetupActionAssertRulesetRule represents the referenced rule within the external ruleset template
type TestScriptSetupActionAssertRulesetRule struct {
	common.BackboneElement

	// Id of the referenced rule within the external ruleset template
	RuleID string `json:"ruleId"`

	// The referenced rule within the external ruleset template
	Param []TestScriptSetupActionAssertRuleParam `json:"param,omitempty"`
}

// TestScriptSetupAction represents a single action performed as part of the test setup
type TestScriptSetupAction struct {
	common.BackboneElement

	// The operation to perform
	Operation *TestScriptSetupActionOperation `json:"operation,omitempty"`

	// The assertion to perform
	Assert *TestScriptSetupActionAssert `json:"assert,omitempty"`
}

// TestScriptSetup represents a series of required setup operations before tests are executed
type TestScriptSetup struct {
	common.BackboneElement

	// Action would contain either an operation or an assertion
	Action []TestScriptSetupAction `json:"action"`
}

// TestScriptTestAction represents a single action performed as part of the test
type TestScriptTestAction struct {
	common.BackboneElement

	// The operation to perform
	Operation *TestScriptSetupActionOperation `json:"operation,omitempty"`

	// The assertion to perform
	Assert *TestScriptSetupActionAssert `json:"assert,omitempty"`
}

// TestScriptTest represents a test in this script
type TestScriptTest struct {
	common.BackboneElement

	// The name of this test used for tracking/logging purposes by test engines
	Name *string `json:"name,omitempty"`

	// A short description of the test used by test engines for tracking and reporting purposes
	Description *string `json:"description,omitempty"`

	// Action would contain either an operation or an assertion
	Action []TestScriptTestAction `json:"action"`
}

// TestScriptTeardownAction represents a single action performed as part of the test teardown
type TestScriptTeardownAction struct {
	common.BackboneElement

	// The teardown operation to perform
	Operation TestScriptSetupActionOperation `json:"operation"`
}

// TestScriptTeardown represents a series of operations required to clean up after all the tests are executed
type TestScriptTeardown struct {
	common.BackboneElement

	// The teardown action will only contain an operation
	Action []TestScriptTeardownAction `json:"action"`
}

// TestScriptStatus represents the status of a test script
type TestScriptStatus string

const (
	TestScriptStatusDraft   TestScriptStatus = "draft"
	TestScriptStatusActive  TestScriptStatus = "active"
	TestScriptStatusRetired TestScriptStatus = "retired"
	TestScriptStatusUnknown TestScriptStatus = "unknown"
)

// TestScript represents a structured set of tests against a FHIR server or client implementation to determine compliance against the FHIR specification
type TestScript struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "TestScript"

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the test script and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// The (c) symbol should NOT be included in this string
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The date is often not tracked until the resource is published, but may be present on draft content
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as comments about misuse, instructions for clinical use and interpretation, literature references, examples from the paper world, etc
	Description *string `json:"description,omitempty"`

	// The purpose of this element is to define the profile of a destination element used elsewhere in the script
	Destination []TestScriptDestination `json:"destination,omitempty"`

	// Allows filtering of test scripts that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Fixture in the test script - by reference (uri)
	Fixture []TestScriptFixture `json:"fixture,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type, and can then identify this test script outside of FHIR, where it is not possible to use the logical URI
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the test script to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The required capability must exist and are assumed to function correctly on the FHIR server being tested
	Metadata *TestScriptMetadata `json:"metadata,omitempty"`

	// The name is not expected to be globally unique
	Name string `json:"name"`

	// The purpose of this element is to define the profile of an origin element used elsewhere in the script
	Origin []TestScriptOrigin `json:"origin,omitempty"`

	// See the Resource List for complete list of resource types
	Profile []string `json:"profile,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the test script
	Purpose *string `json:"purpose,omitempty"`

	// The scope indicates a conformance artifact that is tested by the test(s) within this test case
	Scope []TestScriptScope `json:"scope,omitempty"`

	// A series of required setup operations before tests are executed
	Setup *TestScriptSetup `json:"setup,omitempty"`

	// Allows filtering of test scripts that are appropriate for use versus not
	Status TestScriptStatus `json:"status"`

	// A series of operations required to clean up after all the tests are executed
	Teardown *TestScriptTeardown `json:"teardown,omitempty"`

	// A test in this script
	Test []TestScriptTest `json:"test,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// Variables would be set based either on XPath/JSONPath expressions against fixtures
	Variable []TestScriptVariable `json:"variable,omitempty"`

	// There may be different test script instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`

	// If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}
