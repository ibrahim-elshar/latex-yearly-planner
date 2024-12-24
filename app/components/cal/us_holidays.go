package cal

import (
	"time"
)

// getUSHolidays returns a list of US holidays for a given year
func getUSHolidays(year int) Holidays {
	holidays := Holidays{
		// Fixed date holidays
		{Name: "New Year's Day", Date: time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local), IsAllDay: true},
		{Name: "Juneteenth", Date: time.Date(year, time.June, 19, 0, 0, 0, 0, time.Local), IsAllDay: true},
		{Name: "Independence Day", Date: time.Date(year, time.July, 4, 0, 0, 0, 0, time.Local), IsAllDay: true},
		{Name: "Veterans Day", Date: time.Date(year, time.November, 11, 0, 0, 0, 0, time.Local), IsAllDay: true},
		{Name: "Christmas", Date: time.Date(year, time.December, 25, 0, 0, 0, 0, time.Local), IsAllDay: true},
	}

	// Add holidays that fall on specific weekdays
	holidays = append(holidays,
		getNthWeekdayHoliday(year, time.January, time.Monday, 3, "M L King Day"),           // 3rd Monday in January
		getNthWeekdayHoliday(year, time.February, time.Monday, 3, "Presidents' Day"),       // 3rd Monday in February
		getLastWeekdayHoliday(year, time.May, time.Monday, "Memorial Day"),                 // Last Monday in May
		getNthWeekdayHoliday(year, time.September, time.Monday, 1, "Labor Day"),           // 1st Monday in September
		getNthWeekdayHoliday(year, time.October, time.Monday, 2, "Columbus Day"),          // 2nd Monday in October
		getNthWeekdayHoliday(year, time.November, time.Thursday, 4, "Thanksgiving Day"),   // 4th Thursday in November
	)

	return holidays
}

// getNthWeekdayHoliday returns a holiday that occurs on the nth occurrence of a weekday in a month
func getNthWeekdayHoliday(year int, month time.Month, weekday time.Weekday, n int, name string) *Holiday {
	// Start with the first day of the month
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	
	// Find the first occurrence of the weekday
	for date.Weekday() != weekday {
		date = date.AddDate(0, 0, 1)
	}
	
	// Add weeks to get to the nth occurrence
	date = date.AddDate(0, 0, (n-1)*7)
	
	return &Holiday{
		Name:     name,
		Date:     date,
		IsAllDay: true,
	}
}

// getLastWeekdayHoliday returns a holiday that occurs on the last occurrence of a weekday in a month
func getLastWeekdayHoliday(year int, month time.Month, weekday time.Weekday, name string) *Holiday {
	// Start with the last day of the month
	date := time.Date(year, month+1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
	
	// Find the last occurrence of the weekday
	for date.Weekday() != weekday {
		date = date.AddDate(0, 0, -1)
	}
	
	return &Holiday{
		Name:     name,
		Date:     date,
		IsAllDay: true,
	}
} 