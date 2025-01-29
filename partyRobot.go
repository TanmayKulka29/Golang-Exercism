package partyrobot
import (
    "fmt"
    "strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
    return "Welcome to my party, " + name + "!"
	panic("Please implement the Welcome function")
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
    ageStr := strconv.Itoa(age)
    return "Happy birthday " + name + "! You are now " +ageStr+" years old!"
	panic("Please implement the HappyBirthday function")
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
    tableStr := fmt.Sprintf("%03d", table)
	distanceStr := fmt.Sprintf("%.1f", distance)
    return "Welcome to my party, " + name + "!\n" + "You have been assigned to table "+tableStr+". Your table is "+direction+", exactly "+distanceStr+" meters from here.\n"+"You will be sitting next to "+neighbor+"."
	panic("Please implement the AssignTable function")
}
