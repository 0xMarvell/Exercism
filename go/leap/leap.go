// Package leap checks if the inputted year is a (Gregorian Calender) leap year or not
package leap

// IsLeapYear checks if the inputted year is
// a (Gregorian Calender) leap year or not.
func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%4 == 0) && (year%100 == 0 && year%400 == 0)
}
