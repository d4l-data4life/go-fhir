# Go FHIR Type Definitions

This project contains Go type definitions for HL7 FHIR (Fast Healthcare Interoperability Resources) converted from the original TypeScript definitions. It supports all major FHIR versions from R2 through R5.

## Overview

FHIR is a standard for exchanging healthcare information electronically. This Go library provides strongly-typed data structures for working with FHIR resources across different versions.

### Supported FHIR Versions

- **FHIR R2** (DSTU2, version 1.0.2) - `fhir2` package
- **FHIR R3** (STU3, version 3.0.2) - `fhir3` package  
- **FHIR R4** (version 4.0.1) - `fhir4` package
- **FHIR R4B** (version 4.3.0) - `fhir4b` package
- **FHIR R5** (version 5.0.0) - `fhir5` package

## Project Structure

```
go-fhir/
├── common/         # Shared base types and utilities
│   └── types.go
├── fhir2/          # FHIR R2/DSTU2 definitions
│   └── datatypes.go
├── fhir3/          # FHIR R3/STU3 definitions
│   └── datatypes.go
├── fhir4/          # FHIR R4 definitions
│   ├── datatypes.go
│   └── resources.go
├── fhir4b/         # FHIR R4B definitions
│   └── datatypes.go
├── fhir5/          # FHIR R5 definitions
│   └── datatypes.go
├── go.mod
└── README.md
```

## Installation

```bash
go get github.com/go-fhir/go-fhir
```

## Usage

### Basic Example - Creating a Patient Resource

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/go-fhir/go-fhir/fhir4"
    "github.com/go-fhir/go-fhir/common"
)

func main() {
    // Create a new Patient resource
    patient := &fhir4.Patient{
        DomainResource: fhir4.DomainResource{
            Resource: fhir4.Resource{
                ResourceType: "Patient",
                ID:          stringPtr("patient-example"),
            },
        },
        Active: boolPtr(true),
        Name: []fhir4.HumanName{
            {
                Use:    &fhir4.NameUseOfficial,
                Family: stringPtr("Doe"),
                Given:  []string{"John", "Michael"},
            },
        },
        Gender:    &fhir4.AdministrativeGenderMale,
        BirthDate: stringPtr("1990-01-01"),
        Address: []fhir4.Address{
            {
                Use:        &fhir4.AddressUseHome,
                Line:       []string{"123 Main Street"},
                City:       stringPtr("Anytown"),
                State:      stringPtr("NY"),
                PostalCode: stringPtr("12345"),
                Country:    stringPtr("US"),
            },
        },
    }

    // Marshal to JSON
    jsonData, err := json.MarshalIndent(patient, "", "  ")
    if err != nil {
        panic(err)
    }
    
    fmt.Println(string(jsonData))
}

// Helper functions for pointer creation
func stringPtr(s string) *string { return &s }
func boolPtr(b bool) *bool { return &b }
```

### Creating an Observation Resource

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/go-fhir/go-fhir/fhir4"
    "github.com/go-fhir/go-fhir/common"
)

func main() {
    observation := &fhir4.Observation{
        DomainResource: fhir4.DomainResource{
            Resource: fhir4.Resource{
                ResourceType: "Observation",
                ID:          stringPtr("blood-pressure-example"),
            },
        },
        Status: fhir4.ObservationStatusFinal,
        Category: []common.CodeableConcept{
            {
                Coding: []common.Coding{
                    {
                        System:  stringPtr("http://terminology.hl7.org/CodeSystem/observation-category"),
                        Code:    stringPtr("vital-signs"),
                        Display: stringPtr("Vital Signs"),
                    },
                },
            },
        },
        Code: common.CodeableConcept{
            Coding: []common.Coding{
                {
                    System:  stringPtr("http://loinc.org"),
                    Code:    stringPtr("85354-9"),
                    Display: stringPtr("Blood pressure panel"),
                },
            },
        },
        Subject: &common.Reference{
            Reference: stringPtr("Patient/patient-example"),
        },
        Component: []fhir4.ObservationComponent{
            {
                Code: common.CodeableConcept{
                    Coding: []common.Coding{
                        {
                            System:  stringPtr("http://loinc.org"),
                            Code:    stringPtr("8480-6"),
                            Display: stringPtr("Systolic blood pressure"),
                        },
                    },
                },
                ValueQuantity: &common.Quantity{
                    Value:  float64Ptr(120),
                    Unit:   stringPtr("mmHg"),
                    System: stringPtr("http://unitsofmeasure.org"),
                    Code:   stringPtr("mm[Hg]"),
                },
            },
            {
                Code: common.CodeableConcept{
                    Coding: []common.Coding{
                        {
                            System:  stringPtr("http://loinc.org"),
                            Code:    stringPtr("8462-4"),
                            Display: stringPtr("Diastolic blood pressure"),
                        },
                    },
                },
                ValueQuantity: &common.Quantity{
                    Value:  float64Ptr(80),
                    Unit:   stringPtr("mmHg"),
                    System: stringPtr("http://unitsofmeasure.org"),
                    Code:   stringPtr("mm[Hg]"),
                },
            },
        },
    }

    jsonData, err := json.MarshalIndent(observation, "", "  ")
    if err != nil {
        panic(err)
    }
    
    fmt.Println(string(jsonData))
}

func float64Ptr(f float64) *float64 { return &f }
func stringPtr(s string) *string { return &s }
```

### Working with Different FHIR Versions

```go
// Using FHIR R4
import "github.com/go-fhir/go-fhir/fhir4"

patient4 := &fhir4.Patient{
    // R4 Patient structure
}

// Using FHIR R5  
import "github.com/go-fhir/go-fhir/fhir5"

// Note: R5 uses DataType as base instead of Element
address5 := &fhir5.Address{
    // R5 Address structure with additional fields
}

// Using FHIR R2 (older version with fewer fields)
import "github.com/go-fhir/go-fhir/fhir2"

address2 := &fhir2.Address{
    // R2 Address structure - no billing use option
}
```

## Key Differences Between FHIR Versions

### Type Hierarchy Changes
- **R2**: Uses `Element` as base type
- **R3**: Uses `Element` as base type, adds `billing` to Address.use
- **R4**: Uses `Element` as base type, mature feature set
- **R4B**: Minor corrections to R4, same structure
- **R5**: Uses `DataType` as base type, adds new fields like `Attachment.duration`, `Attachment.frames`, etc.

### ContactPoint System Values
- **R2**: `phone`, `fax`, `email`, `pager`, `other`
- **R3+**: Adds `url`, `sms` options

### Address Use Values  
- **R2**: `home`, `work`, `temp`, `old`
- **R3+**: Adds `billing` option

## Common Patterns

### Pointer Helpers
Since Go requires pointers for optional fields, consider creating helper functions:

```go
func StringPtr(s string) *string { return &s }
func BoolPtr(b bool) *bool { return &b }
func IntPtr(i int) *int { return &i }
func Float64Ptr(f float64) *float64 { return &f }
```

### Resource Validation
All resources include required fields as non-pointer types and optional fields as pointers:

```go
// Required field (non-pointer)
observation.Status = fhir4.ObservationStatusFinal

// Optional field (pointer)
observation.Subject = &common.Reference{
    Reference: StringPtr("Patient/123"),
}
```

## JSON Serialization

All types include appropriate JSON tags for marshaling/unmarshaling:

```go
// Serialize to JSON
jsonBytes, err := json.Marshal(patient)

// Deserialize from JSON  
var patient fhir4.Patient
err := json.Unmarshal(jsonBytes, &patient)
```

## Contributing

This project was converted from the official FHIR TypeScript definitions. When FHIR specifications are updated, the TypeScript definitions should be reconverted to Go.

### Conversion Notes
- TypeScript interfaces → Go structs
- Optional fields (`field?: type`) → Go pointers (`*type`)
- Union types (`'a'|'b'|'c'`) → Go string constants
- Arrays remain arrays
- `undefined` values → `nil` pointers

## License

This project follows the same license as the original FHIR TypeScript definitions (MIT).

## Resources

- [HL7 FHIR Specification](http://hl7.org/fhir/)
- [Original TypeScript Definitions](https://github.com/DefinitelyTyped/DefinitelyTyped/tree/master/types/fhir)
- [FHIR R4 Documentation](http://hl7.org/fhir/R4/)
- [FHIR R5 Documentation](http://hl7.org/fhir/R5/) 
