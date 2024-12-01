package tools

func CountOccurrences[T comparable](element T, slice []T) int {
	count := 0
	for _, v := range slice {
		if v == element {
			count++
		}
	}
	return count
}
