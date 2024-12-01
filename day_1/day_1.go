package day1

import (
	"advent-of-code-2024/tools"
	"math"
	"sort"
	"strconv"
	"strings"
)

// Split number pairs for each row into two slices
// Order the two lists by number ascending
// Compare each index pairs of the two slices and determine the greater number
// If a > b then do: a - b as the result
// If b > a then do: b - a as the result
// If a == b then take 0 as the result
// Store result for each comparison into a slice
// At the end sum up the numbers of the result slice to get the total distance

func getLocationIdListsOrdered(input string) ([]int, []int, error) {
	locationIdStrings := strings.Fields(input)

	locationIdList1 := make([]int, 0, len(locationIdStrings)/2)
	locationIdList2 := make([]int, 0, len(locationIdStrings)/2)

	for i, idString := range locationIdStrings {
		id, convErr := strconv.Atoi(idString)
		if convErr != nil {
			return nil, nil, convErr
		}

		if i%2 == 0 {
			locationIdList1 = append(locationIdList1, id)
		} else {
			locationIdList2 = append(locationIdList2, id)
		}
	}

	sort.Ints(locationIdList1)
	sort.Ints(locationIdList2)

	return locationIdList1, locationIdList2, nil
}

func CalculateTotalLocationDistances(input string) (int, error) {
	locationIdList1, locationIdList2, err := getLocationIdListsOrdered(input)
	if err != nil {
		return 0, err
	}

	var totalLocationDinstance int

	for i := range locationIdList1 {
		id1 := locationIdList1[i]
		id2 := locationIdList2[i]

		totalLocationDinstance += int(math.Abs(float64(id1 - id2)))
	}

	return totalLocationDinstance, err
}

func CalculateTotalSimilarityScore(input string) (int, error) {
	locationIdList1, locationIdList2, err := getLocationIdListsOrdered(input)
	if err != nil {
		return 0, err
	}

	var totalSimilarityScore int

	for _, locationId := range locationIdList1 {
		totalSimilarityScore += locationId * tools.CountOccurrences(locationId, locationIdList2)
	}

	return totalSimilarityScore, nil
}
