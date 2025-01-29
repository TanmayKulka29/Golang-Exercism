package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    sum := 0
    for i:=0;i<len(birdsPerDay);i++ {
        sum += birdsPerDay[i]
    }
    return sum
	panic("Please implement the TotalBirdCount() function")
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    sum := 0
    week -= 1
	start := 7*week
    for i:=start;i<(start+7);i++ {
        sum+=birdsPerDay[i]
    }
    return sum
	panic("Please implement the BirdsInWeek() function")
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i:=0;i<len(birdsPerDay);i=i+2 {
        birdsPerDay[i]+=1
    }
    return birdsPerDay
	panic("Please implement the FixBirdCountLog() function")
}
