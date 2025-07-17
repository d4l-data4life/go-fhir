package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// TestScript represents a structured collection of tests against a FHIR server or client implementation
type TestScript struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "TestScript"

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the test script and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the test script
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the test script was built
	Description *string `json:"description,omitempty"`

	// The purpose of this element is to define the profile of a destination element used elsewhere in the script
	Destination []TestScriptDestination `json:"destination,omitempty"`

	// Allows filtering of test scripts that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Fixture in the test script - by reference (uri). All fixtures are required for the test script to execute
	Fixture []TestScriptFixture `json:"fixture,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the test script to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The required capability must exist and are assumed to function correctly on the FHIR server being tested
	Metadata *TestScriptMetadata `json:"metadata,omitempty"`

	// The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly
	Name string `json:"name"`

	// The purpose of this element is to define the profile of an origin element used elsewhere in the script
	Origin []TestScriptOrigin `json:"origin,omitempty"`

	// See http://build.fhir.org/resourcelist.html for complete list of resource types
	Profile []common.Reference `json:"profile,omitempty"`

	// Usually an organization but may be an individual. The publisher (or steward) of the test script is the organization or individual primarily responsible for the maintenance and upkeep of the test script
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the test script. Instead, it provides traceability of 'why' the resource is either needed or 'why' it is defined as it is
	Purpose *string `json:"purpose,omitempty"`

	// A series of required setup operations before tests are executed
	Setup *TestScriptSetup `json:"setup,omitempty"`

	// Allows filtering of test scripts that are appropriate for use versus not
	Status string `json:"status"` // "draft" | "active" | "retired" | "unknown"

	// A series of operations required to clean up after all the tests are executed (successfully or otherwise)
	Teardown *TestScriptTeardown `json:"teardown,omitempty"`

	// A test in this script
	Test []TestScriptTest `json:"test,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url string `json:"url"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// Variables would be set based either on XPath/JSONPath expressions against fixtures
	Variable []TestScriptVariable `json:"variable,omitempty"`

	// There may be different test script instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// TestScriptOrigin represents the purpose of this element is to define the profile of an origin element used elsewhere in the script
type TestScriptOrigin struct {
	common.BackboneElement

	// A given origin index (e.g. 1) can appear only once in the list
	Index int `json:"index"`

	// Must be a "sender"/"client" profile
	Profile common.Coding `json:"profile"`
}

// TestScriptDestination represents the purpose of this element is to define the profile of a destination element used elsewhere in the script
type TestScriptDestination struct {
	common.BackboneElement

	// A given destination index (e.g. 1) can appear only once in the list
	Index int `json:"index"`

	// Must be a "receiver"/"server" profile
	Profile common.Coding `json:"profile"`
}

// TestScriptMetadataLink represents a link to the FHIR specification that this test is covering
type TestScriptMetadataLink struct {
	common.BackboneElement

	// Short description of the link
	Description *string `json:"description,omitempty"`

	// URL to a particular requirement or feature within the FHIR specification
	Url string `json:"url"`
}

// TestScriptMetadataCapability represents when the metadata capabilities section is defined at TestScript.metadata
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

	// When the metadata capabilities section is defined at TestScript.metadata
	Capability []TestScriptMetadataCapability `json:"capability"`

	// A link to the FHIR specification that this test is covering
	Link []TestScriptMetadataLink `json:"link,omitempty"`
}

// TestScriptFixture represents fixture in the test script - by reference (uri)
type TestScriptFixture struct {
	common.BackboneElement

	// Whether or not to implicitly create the fixture during setup
	Autocreate bool `json:"autocreate"`

	// Whether or not to implicitly delete the fixture during teardown
	Autodelete bool `json:"autodelete"`

	// See http://build.fhir.org/resourcelist.html for complete list of resource types
	Resource *common.Reference `json:"resource,omitempty"`
}

// TestScriptVariable represents variables would be set based either on XPath/JSONPath expressions against fixtures
type TestScriptVariable struct {
	common.BackboneElement

	// The purpose of this element is to allow for a pre-defined value that can be used as a default or as an override value
	DefaultValue *string `json:"defaultValue,omitempty"`

	// A free text natural language description of the variable and its purpose
	Description *string `json:"description,omitempty"`

	// If headerField is defined, then the variable will be evaluated against the headers that sourceId is pointing to
	Expression *string `json:"expression,omitempty"`

	// If headerField is defined, then the variable will be evaluated against the headers that sourceId is pointing to
	HeaderField *string `json:"headerField,omitempty"`

	// If headerField is defined, then the variable will be evaluated against the headers that sourceId is pointing to
	Hint *string `json:"hint,omitempty"`

	// If headerField is defined, then the variable will be evaluated against the headers that sourceId is pointing to
	Path *string `json:"path,omitempty"`

	// If headerField is defined, then the variable will be evaluated against the headers that sourceId is pointing to
	SourceId *string `json:"sourceId,omitempty"`
}

// TestScriptSetup represents a series of required setup operations before tests are executed
type TestScriptSetup struct {
	common.BackboneElement

	// Action would contain either an operation or an assertion but not both
	Action []TestScriptSetupAction `json:"action"`

	// The metadata capabilities section is defined at TestScript.setup.metadata
	Metadata *TestScriptMetadata `json:"metadata,omitempty"`
}

// TestScriptSetupAction represents action would contain either an operation or an assertion but not both
type TestScriptSetupAction struct {
	common.BackboneElement

	// The results of the assertion performed on the previous operations
	Assert *TestScriptSetupActionAssert `json:"assert,omitempty"`

	// The operation to perform
	Operation *TestScriptSetupActionOperation `json:"operation,omitempty"`
}

// TestScriptSetupActionOperation represents the operation to perform
type TestScriptSetupActionOperation struct {
	common.BackboneElement

	// The type of operation to perform
	Type *common.Coding `json:"type,omitempty"`

	// The resource type
	Resource *string `json:"resource,omitempty"`

	// The label would be used for tracking/logging purposes by test engines
	Label *string `json:"label,omitempty"`

	// The description would be used by test engines for tracking and reporting purposes
	Description *string `json:"description,omitempty"`

	// The mime-type to use for RESTful operations in the 'Accept' header
	Accept *string `json:"accept,omitempty"`

	// The mime-type to use for RESTful operations in the 'Content-Type' header
	ContentType *string `json:"contentType,omitempty"`

	// The server where the request message is destined for
	Destination *int `json:"destination,omitempty"`

	// Whether or not to implicitly send the request url in encoded format
	EncodeRequestUrl *bool `json:"encodeRequestUrl,omitempty"`

	// The server where the request message originates from
	Origin *int `json:"origin,omitempty"`

	// Path plus parameters after [type]
	Params *string `json:"params,omitempty"`

	// Header elements would be used to set HTTP headers
	RequestHeader []TestScriptSetupActionOperationRequestHeader `json:"requestHeader,omitempty"`

	// The fixture id (maybe new) to map to the request
	RequestId *string `json:"requestId,omitempty"`

	// The type of the resource
	ResponseId *string `json:"responseId,omitempty"`

	// The id of the fixture used as the body of a PUT or POST request
	SourceId *string `json:"sourceId,omitempty"`

	// Complete request URL
	TargetId *string `json:"targetId,omitempty"`

	// The URL for the request
	Url *string `json:"url,omitempty"`
}

// TestScriptSetupActionOperationRequestHeader represents header elements would be used to set HTTP headers
type TestScriptSetupActionOperationRequestHeader struct {
	common.BackboneElement

	// The HTTP header field e.g. "Accept"
	Field string `json:"field"`

	// The value of the header e.g. "application/fhir+xml"
	Value string `json:"value"`
}

// TestScriptSetupActionAssert represents the results of the assertion performed on the previous operations
type TestScriptSetupActionAssert struct {
	common.BackboneElement

	// The label would be used for tracking/logging purposes by test engines
	Label *string `json:"label,omitempty"`

	// The description would be used by test engines for tracking and reporting purposes
	Description *string `json:"description,omitempty"`

	// The direction to use for the assertion
	Direction *string `json:"direction,omitempty"` // "response" | "request"

	// Id of the source fixture to be evaluated
	CompareToSourceId *string `json:"compareToSourceId,omitempty"`

	// The fluentpath expression to evaluate against the source fixture
	CompareToSourceExpression *string `json:"compareToSourceExpression,omitempty"`

	// XPath or JSONPath expression to evaluate against the source fixture
	CompareToSourcePath *string `json:"compareToSourcePath,omitempty"`

	// The mime-type to use for RESTful operations in the 'Accept' header
	ContentType *string `json:"contentType,omitempty"`

	// The fluentpath expression to be evaluated
	Expression *string `json:"expression,omitempty"`

	// The HTTP header field name e.g. 'Location'
	HeaderField *string `json:"headerField,omitempty"`

	// The ID of a fixture
	MinimumId *string `json:"minimumId,omitempty"`

	// Whether or not the test execution will perform the validation
	NavigationLinks *bool `json:"navigationLinks,omitempty"`

	// The operator type defines the conditional behavior of the assert
	Operator *string `json:"operator,omitempty"` // "equals" | "notEquals" | "in" | "notIn" | "greaterThan" | "lessThan" | "empty" | "notEmpty" | "contains" | "notContains" | "eval"

	// The XPath or JSONPath expression to evaluate against the fixture
	Path *string `json:"path,omitempty"`

	// The request method or HTTP operation code to compare against that used by the client
	RequestMethod *string `json:"requestMethod,omitempty"` // "delete" | "get" | "options" | "patch" | "post" | "put" | "head"

	// The value to compare to
	RequestURL *string `json:"requestURL,omitempty"`

	// The type of the resource
	Resource *string `json:"resource,omitempty"`

	// okay | created | noContent | notModified | bad | forbidden | notFound | methodNotAllowed | conflict | gone | preconditionFailed | unprocessable
	Response *string `json:"response,omitempty"` // "okay" | "created" | "noContent" | "notModified" | "bad" | "forbidden" | "notFound" | "methodNotAllowed" | "conflict" | "gone" | "preconditionFailed" | "unprocessable"

	// The value of the HTTP response code to be tested
	ResponseCode *string `json:"responseCode,omitempty"`

	// The TestScript.rule this assert will evaluate
	Rule *TestScriptSetupActionAssertRule `json:"rule,omitempty"`

	// The TestScript.ruleset this assert will evaluate
	Ruleset *TestScriptSetupActionAssertRuleset `json:"ruleset,omitempty"`

	// Fixture to evaluate the XPath/JSONPath expression or the headerField against
	SourceId *string `json:"sourceId,omitempty"`

	// The ID of the Profile to validate against
	ValidateProfileId *string `json:"validateProfileId,omitempty"`

	// The value to compare to
	Value *string `json:"value,omitempty"`

	// Will this assert produce a warning only on error
	WarningOnly *bool `json:"warningOnly,omitempty"`
}

// TestScriptSetupActionAssertRule represents the TestScript.rule this assert will evaluate
type TestScriptSetupActionAssertRule struct {
	common.BackboneElement

	// The TestScript.rule id value this assert will evaluate
	RuleId string `json:"ruleId"`

	// The referenced rule param
	Param []TestScriptSetupActionAssertRuleParam `json:"param,omitempty"`
}

// TestScriptSetupActionAssertRuleParam represents the referenced rule param
type TestScriptSetupActionAssertRuleParam struct {
	common.BackboneElement

	// Parameter name matching external assert rule parameter
	Name string `json:"name"`

	// Parameter value defined either explicitly or dynamically
	Value string `json:"value"`
}

// TestScriptSetupActionAssertRuleset represents the TestScript.ruleset this assert will evaluate
type TestScriptSetupActionAssertRuleset struct {
	common.BackboneElement

	// The TestScript.ruleset id value this assert will evaluate
	RulesetId string `json:"rulesetId"`

	// The referenced rule param
	Rule []TestScriptSetupActionAssertRulesetRule `json:"rule,omitempty"`
}

// TestScriptSetupActionAssertRulesetRule represents the referenced rule param
type TestScriptSetupActionAssertRulesetRule struct {
	common.BackboneElement

	// The TestScript.rule id value this assert will evaluate
	RuleId string `json:"ruleId"`

	// The referenced rule param
	Param []TestScriptSetupActionAssertRuleParam `json:"param,omitempty"`
}

// TestScriptTest represents a test in this script
type TestScriptTest struct {
	common.BackboneElement

	// The name of this test used for tracking/logging purposes by test engines
	Name *string `json:"name,omitempty"`

	// A short description of the test used by test engines for tracking and reporting purposes
	Description *string `json:"description,omitempty"`

	// Action would contain either an operation or an assertion but not both
	Action []TestScriptTestAction `json:"action"`
}

// TestScriptTestAction represents action would contain either an operation or an assertion but not both
type TestScriptTestAction struct {
	common.BackboneElement

	// The results of the assertion performed on the previous operations
	Assert *TestScriptSetupActionAssert `json:"assert,omitempty"`

	// The operation to perform
	Operation *TestScriptSetupActionOperation `json:"operation,omitempty"`
}

// TestScriptTeardown represents a series of operations required to clean up after all the tests are executed
type TestScriptTeardown struct {
	common.BackboneElement

	// The teardown action will only contain an operation
	Action []TestScriptTeardownAction `json:"action"`
}

// TestScriptTeardownAction represents the teardown action will only contain an operation
type TestScriptTeardownAction struct {
	common.BackboneElement

	// The teardown operation to perform
	Operation TestScriptSetupActionOperation `json:"operation"`
}
