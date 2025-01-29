package strain

// Implement the "Keep" and "Discard" function in this file.
func Keep[T any](elements []T, predicate func(T) bool) []T {
    keep := make([]T,0)
    for _, ele := range elements {
        if predicate(ele) {
            keep = append(keep, ele)
        }
    }
    return keep
}

func Discard[T any](elements []T, predicate func(T) bool) []T {
    disc := make([]T,0)
    for _, ele := range elements {
        if !predicate(ele) {
            disc = append(disc, ele)
        }
    }
    return disc
}
