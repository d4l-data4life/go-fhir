package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// TestPlanDependency represents the required criteria to execute the test plan
type TestPlanDependency struct {
	common.BackboneElement

	// A textual description of the criterium - what is needed for the dependency to be considered met
	Description *string `json:"description,omitempty"`

	// Predecessor test plans - those that are expected to be successfully performed as a dependency for the execution of this test plan
	Predecessor *common.Reference `json:"predecessor,omitempty"`
}

// TestPlanTestCaseDependency represents the required criteria to execute the test case
type TestPlanTestCaseDependency struct {
	common.BackboneElement

	// Description of the criteria
	Description *string `json:"description,omitempty"`

	// Link to predecessor test plans
	Predecessor *common.Reference `json:"predecessor,omitempty"`
}

// TestPlanTestCaseTestRunScript represents the test cases in a structured language
type TestPlanTestCaseTestRunScript struct {
	common.BackboneElement

	// The language for the test cases e.g. 'gherkin', 'testscript'
	Language *common.CodeableConcept `json:"language,omitempty"`

	// The actual content of the cases - references to TestScripts or externally defined content
	SourceString *string `json:"sourceString,omitempty"`

	// The actual content of the cases - references to TestScripts or externally defined content
	SourceReference *common.Reference `json:"sourceReference,omitempty"`
}

// TestPlanTestCaseTestRun represents the actual test to be executed
type TestPlanTestCaseTestRun struct {
	common.BackboneElement

	// The narrative description of the tests
	Narrative *string `json:"narrative,omitempty"`

	// The test cases in a structured language e.g. gherkin, Postman, or FHIR TestScript
	Script *TestPlanTestCaseTestRunScript `json:"script,omitempty"`
}

// TestPlanTestCase represents the individual test cases that are part of this plan
type TestPlanTestCase struct {
	common.BackboneElement

	// The required criteria to execute the test case - e.g. preconditions, previous tests
	Dependency []TestPlanTestCaseDependency `json:"dependency,omitempty"`

	// The actual test to be executed
	TestRun []TestPlanTestCaseTestRun `json:"testRun,omitempty"`
}

// TestPlan represents a plan for executing testing on an artifact and/or the requirements for a collection of test cases
type TestPlan struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "TestPlan"

	// The category of the Test Plan - can be acceptance, unit, performance, etc
	Category []common.CodeableConcept `json:"category,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the test plan and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// The (c) symbol should NOT be included in this string
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The date is often not tracked until the resource is published, but may be present on draft content
	Date *string `json:"date,omitempty"`

	// The required criteria to execute the test plan - e.g. preconditions, previous tests
	Dependency []TestPlanDependency `json:"dependency,omitempty"`

	// This description can be used to capture details such as comments about misuse, instructions for clinical use and interpretation, literature references, examples from the paper world, etc
	Description *string `json:"description,omitempty"`

	// The threshold or criteria for the test plan to be considered successfully executed - narrative
	ExitCriteria *string `json:"exitCriteria,omitempty"`

	// Allows filtering of test plans that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type, and can then identify this test plan outside of FHIR, where it is not possible to use the logical URI
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the test plan to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the test plan
	Purpose *string `json:"purpose,omitempty"`

	// What is being tested with this Test Plan - a conformance resource, or narrative criteria, or an external reference
	Scope []common.Reference `json:"scope,omitempty"`

	// Allows filtering of test plans that are appropriate for use versus not
	Status TestPlanStatus `json:"status"`

	// The individual test cases that are part of this plan, when they they are made explicit
	TestCase []TestPlanTestCase `json:"testCase,omitempty"`

	// A description of test tools to be used in the test plan
	TestTools *string `json:"testTools,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different test plan instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`

	// If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}

// TestPlanStatus represents the status of a test plan
type TestPlanStatus string

const (
	TestPlanStatusDraft   TestPlanStatus = "draft"
	TestPlanStatusActive  TestPlanStatus = "active"
	TestPlanStatusRetired TestPlanStatus = "retired"
	TestPlanStatusUnknown TestPlanStatus = "unknown"
)
