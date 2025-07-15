// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

// Instruction represents instruction details for the service request
type Instruction struct {
	DataType

	// Describes the instruction
	InstructionText *string `json:"instructionText,omitempty"`
}
