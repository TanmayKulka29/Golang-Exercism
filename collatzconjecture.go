package collatzconjecture
import "errors"
func CollatzConjecture(n int) (int, error) {
    if n < 0 {
		return 0, errors.New("n cannot be less than 0")
	}
    if n == 0 {
		return 0, errors.New("n cannot be 0")
	}
    cnt := 0
    for cnt=0;;cnt++ {
        if n==1 {
        	break
        }
    	if n%2 == 0 {
            n/=2
        } else {
            n*=3
            n+=1
        }   
    }
    return cnt, nil
}





package collatzconjecture

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 7a8722a Reorder keys (#1960)

var testCases = []struct {
	description string
	input       int
	expectError bool
	expected    int
}{
	{
		description: "zero steps for one",
		input:       1,
		expectError: false,
		expected:    0,
	},
	{
		description: "divide if even",
		input:       16,
		expectError: false,
		expected:    4,
	},
	{
		description: "even and odd steps",
		input:       12,
		expectError: false,
		expected:    9,
	},
	{
		description: "large number of even and odd steps",
		input:       1000000,
		expectError: false,
		expected:    152,
	},
	{
		description: "zero is an error",
		input:       0,
		expectError: true,
		expected:    0,
	},
	{
		description: "negative value is an error",
		input:       -15,
		expectError: true,
		expected:    0,
	},
}
