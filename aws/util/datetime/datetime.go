package datetime

import (
	"time"
)

// Common ISO 8601 Time formats
const (
	ISO8601BasicFormat      = "20060102T150405Z"
	ISO8601BasicFormatShort = "20060102"
)

// Helper for JSON API DateTime.
type JsonDate float64

// Convert the JSON data type date to time.Time.
func (d JsonDate) Time() time.Time {
	return time.Unix(int64(d), 0)
}
