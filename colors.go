package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	res := []string {"black","brown","red","orange","yellow","green","blue","violet","grey","white"}
    return res
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	mp := map[string]int {
        "black": 0,
        "brown": 1,
        "red": 2,
        "orange": 3,
        "yellow": 4,
        "green": 5,
        "blue": 6,
        "violet": 7,
        "grey": 8,
        "white": 9,
    }
    return mp[color]
}





package resistorcolor

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: d137db1 Format using prettier (#1917)

type colorCodeTestCase struct {
	description string
	input       string
	expected    int
}

type colorsTestCase struct {
	description string
	expected    []string
}

var colorCodeTestCases = []colorCodeTestCase{
	{
		description: "Black",
		input:       "black",
		expected:    0,
	},
	{
		description: "White",
		input:       "white",
		expected:    9,
	},
	{
		description: "Orange",
		input:       "orange",
		expected:    3,
	},
}

var colorsTestCases = []colorsTestCase{
	{
		description: "Colors",
		expected: []string{
			"black",
			"brown",
			"red",
			"orange",
			"yellow",
			"green",
			"blue",
			"violet",
			"grey",
			"white",
		},
	},
}
