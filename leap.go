// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	if year%4 == 0 {
        if year%100!=0 {
            return true
        } else if year%100==0 && year%400 == 0 {
            return true
        } else {
            return false
        }
    }
    return false
}





package leap

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: a2c75d2 leap: fix typo (#1726)

var testCases = []struct {
	description string
	year        int
	expected    bool
}{
	{
		description: "year not divisible by 4 in common year",
		year:        2015,
		expected:    false,
	},
	{
		description: "year divisible by 2, not divisible by 4 in common year",
		year:        1970,
		expected:    false,
	},
	{
		description: "year divisible by 4, not divisible by 100 in leap year",
		year:        1996,
		expected:    true,
	},
	{
		description: "year divisible by 4 and 5 is still a leap year",
		year:        1960,
		expected:    true,
	},
	{
		description: "year divisible by 100, not divisible by 400 in common year",
		year:        2100,
		expected:    false,
	},
	{
		description: "year divisible by 100 but not by 3 is still not a leap year",
		year:        1900,
		expected:    false,
	},
	{
		description: "year divisible by 400 is leap year",
		year:        2000,
		expected:    true,
	},
	{
		description: "year divisible by 400 but not by 125 is still a leap year",
		year:        2400,
		expected:    true,
	},
	{
		description: "year divisible by 200, not divisible by 400 in common year",
		year:        1800,
		expected:    false,
	},
}
