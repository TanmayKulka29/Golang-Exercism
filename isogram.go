package isogram
import ("strings")
func IsIsogram(word string) bool {
    word = strings.ToLower(word)
	mp := make(map[byte]int)
    for i:=0;i<len(word);i++ {
        ch := word[i]
        if ch == ' ' || ch == '-' {
            continue
        }
        if _, ok := mp[ch]; ok {
            return false
        }
        mp[ch]+=1
    }
    return true
}





package isogram

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: d137db1 Format using prettier (#1917)

var testCases = []struct {
	description string
	input       string
	expected    bool
}{
	{
		description: "empty string",
		input:       "",
		expected:    true,
	},
	{
		description: "isogram with only lower case characters",
		input:       "isogram",
		expected:    true,
	},
	{
		description: "word with one duplicated character",
		input:       "eleven",
		expected:    false,
	},
	{
		description: "word with one duplicated character from the end of the alphabet",
		input:       "zzyzx",
		expected:    false,
	},
	{
		description: "longest reported english isogram",
		input:       "subdermatoglyphic",
		expected:    true,
	},
	{
		description: "word with duplicated character in mixed case",
		input:       "Alphabet",
		expected:    false,
	},
	{
		description: "word with duplicated character in mixed case, lowercase first",
		input:       "alphAbet",
		expected:    false,
	},
	{
		description: "hypothetical isogrammic word with hyphen",
		input:       "thumbscrew-japingly",
		expected:    true,
	},
	{
		description: "hypothetical word with duplicated character following hyphen",
		input:       "thumbscrew-jappingly",
		expected:    false,
	},
	{
		description: "isogram with duplicated hyphen",
		input:       "six-year-old",
		expected:    true,
	},
	{
		description: "made-up name that is an isogram",
		input:       "Emily Jung Schwartzkopf",
		expected:    true,
	},
	{
		description: "duplicated character in the middle",
		input:       "accentor",
		expected:    false,
	},
	{
		description: "same first and last characters",
		input:       "angola",
		expected:    false,
	},
	{
		description: "word with duplicated character and with two hyphens",
		input:       "up-to-date",
		expected:    false,
	},
}
