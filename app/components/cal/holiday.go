package cal

import (
	"time"
)

// Holiday represents a single holiday
type Holiday struct {
	Name     string
	Date     time.Time
	IsAllDay bool
}

// Holidays is a collection of Holiday objects
type Holidays []*Holiday

// ForDate returns a holiday that falls on the given date, or nil if none exists
func (h Holidays) ForDate(date time.Time) *Holiday {
	for _, holiday := range h {
		if holiday.Date.Year() == date.Year() &&
			holiday.Date.Month() == date.Month() &&
			holiday.Date.Day() == date.Day() {
			return holiday
		}
	}
	return nil
} 