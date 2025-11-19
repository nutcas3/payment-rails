package api

import (
	"fmt"
	"strings"
	"time"
)

// JSONTime is a custom time type that handles multiple time formats
type JSONTime struct {
	time.Time
}

const (
	// Add more formats as needed
	ctLayout1 = "2006-01-02T15:04:05"
	ctLayout2 = "2006-01-02"
)

func (jt *JSONTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		jt.Time = time.Time{}
		return nil
	}

	// Try standard RFC3339 first
	t, err := time.Parse(time.RFC3339, s)
	if err == nil {
		jt.Time = t
		return nil
	}

	// Try RFC3339Nano
	t, err = time.Parse(time.RFC3339Nano, s)
	if err == nil {
		jt.Time = t
		return nil
	}

	// Try custom layouts
	t, err = time.Parse(ctLayout1, s)
	if err == nil {
		jt.Time = t
		return nil
	}

	t, err = time.Parse(ctLayout2, s)
	if err == nil {
		jt.Time = t
		return nil
	}

	return fmt.Errorf("could not parse time: %s", s)
}

func (jt JSONTime) MarshalJSON() ([]byte, error) {
	if jt.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", jt.Time.Format(time.RFC3339))), nil
}
