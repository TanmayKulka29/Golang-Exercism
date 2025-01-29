package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File[] bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    cnt := 0
    for _, val := range cb[file] {
        if val {
            cnt+=1
        }
    }
    return cnt
	panic("Please implement CountInFile()")
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if rank<1 || rank>8 {
        return 0
    }
    cnt := 0
    for _, file := range cb {
        if file[rank-1] {
            cnt+=1
        }
    }
    return cnt
	panic("Please implement CountInRank()")
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    rows := len(cb)
    cols := 0
    for _, file := range cb {
        cols = len(file)
        break
    }
    return rows*cols
	panic("Please implement CountAll()")
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    cnt := 0
    for _,file := range cb {
        for _, val := range file {
            if val {
                cnt += 1
            }
        }
    }
    return cnt
	panic("Please implement CountOccupied()")
}
