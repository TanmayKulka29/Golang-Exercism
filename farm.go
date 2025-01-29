package thefarm
import (
    "errors"
    "fmt"
)
// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, noc int) (float64, error) {
    fa, err := fc.FodderAmount(noc)
    if err != nil {
        return 0, err
    }
    ff, err := fc.FatteningFactor()
    if err != nil {
        return 0, err
    }
    val := float64(fa*ff)/float64(noc)
    return val, nil
}
// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, noc int) (float64, error) {
    if noc>0 {
        return DivideFood(fc, noc)
    } else {
        return 0, errors.New("invalid number of cows")
    }
}

type InvalidCowsError struct {
    message string
    noc int
}

func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.noc, e.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(noc int) error {
    if noc < 0 {
        return &InvalidCowsError {
            noc: noc,
            message: "there are no negative cows",
        }
    } else if noc == 0 {
        return &InvalidCowsError {
            noc: noc,
            message: "no cows don't need food",
        }
    } else {
        return nil
    }
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.





package thefarm

// This file contains types used in the exercise and tests and should not be modified.

// FodderCalculator provides helper methods to determine the optimal
// amount of fodder to feed cows.
type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}
