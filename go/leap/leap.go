package leap

// IsLeapYear determines if a given year is a leap year
func IsLeapYear(year int) bool {
	// Must be a multiple of 4
	if year%4 == 0 {
		// Must not be a multiple of 100 unless the year is
		// a multiple of 400
		if year%100 != 0 || year%400 == 0 {
			return true
		}
	}
	return false
}
