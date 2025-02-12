package electionday
import ("fmt")
// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
    var pntInitialVotes *int = &initialVotes
    return pntInitialVotes
	panic("Please implement the NewVoteCounter() function")
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
    if counter == nil {
        return 0
    } else {
        return *counter
    }
	panic("Please implement the VoteCount() function")
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
    var ner *ElectionResult = &ElectionResult {
        Name: candidateName,
        Votes: votes,
    }
    return ner
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
    if _,val := results[candidate]; val {
        results[candidate] -= 1
    }
}





package electionday

// ElectionResult represents an election result.
type ElectionResult struct {
	// Name is the name of the candidate.
	Name string
	// Votes is the total number of votes the candidate had.
	Votes int
}
