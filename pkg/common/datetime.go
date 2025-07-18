package common

import (
	"fmt"
	"strings"
	"time"
)

// FHIR is known to use partial dateTime values (e.g., just "2016-01" or "2016") — the spec explicitly allows this.
// This is not supported by the time.Time type, so we need to handle it manually.
// We use the FHIRDateTime type to represent partial dateTime values.

type FHIRDateTime struct {
	Time      time.Time
	Precision string
}

// Supported layouts with increasing precision
var fhirLayouts = []struct {
	layout    string
	precision string
}{
	{time.RFC3339, "second"},
	{"2006-01-02", "day"},
	{"2006-01", "month"},
	{"2006", "year"},
}

// ParseDateTime parses a string into a FHIRDateTime, supporting partial dateTime values
// as specified in the FHIR specification.
func ParseDateTime(s string) (*FHIRDateTime, error) {
	s = strings.TrimSpace(s)
	for _, l := range fhirLayouts {
		t, err := time.Parse(l.layout, s)
		if err == nil {
			return &FHIRDateTime{
				Time:      t,
				Precision: l.precision,
			}, nil
		}
	}
	return nil, fmt.Errorf("FHIRDateTime: unable to parse %q", s)
}

func (f *FHIRDateTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	var err error
	for _, l := range fhirLayouts {
		f.Time, err = time.Parse(l.layout, s)
		if err == nil {
			f.Precision = l.precision
			return nil
		}
	}
	return fmt.Errorf("FHIRDateTime: unable to parse %q", s)
}

func (f FHIRDateTime) MarshalJSON() ([]byte, error) {
	var layout string
	for _, l := range fhirLayouts {
		if l.precision == f.Precision {
			layout = l.layout
			break
		}
	}
	if layout == "" {
		return nil, fmt.Errorf("FHIRDateTime: unknown precision %q", f.Precision)
	}
	return []byte(fmt.Sprintf(`"%s"`, f.Time.Format(layout))), nil
}
