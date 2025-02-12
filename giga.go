// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}





package gigasecond

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 5fc501b Remove unneeded nesting (#1798)

var addCases = []struct {
	description string
	in          string
	want        string
}{
	{
		description: "date only specification of time",
		in:          "2011-04-25",
		want:        "2043-01-01T01:46:40",
	},
	{
		description: "second test for date only specification of time",
		in:          "1977-06-13",
		want:        "2009-02-19T01:46:40",
	},
	{
		description: "third test for date only specification of time",
		in:          "1959-07-19",
		want:        "1991-03-27T01:46:40",
	},
	{
		description: "full time specified",
		in:          "2015-01-24T22:00:00",
		want:        "2046-10-02T23:46:40",
	},
	{
		description: "full time with day roll-over",
		in:          "2015-01-24T23:59:59",
		want:        "2046-10-03T01:46:39",
	},
}
