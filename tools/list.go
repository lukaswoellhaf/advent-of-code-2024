package tools

import "sort"

func CountOccurrences[T comparable](element T, slice []T) int {
	count := 0
	for _, v := range slice {
		if v == element {
			count++
		}
	}
	return count
}

func SliceIsSortedAscOrDesc[T any](slice []T, less func(a, b T) bool) bool {
	asc := sort.SliceIsSorted(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})

	desc := sort.SliceIsSorted(slice, func(i, j int) bool {
		return less(slice[j], slice[i])
	})

	return asc || desc
}
