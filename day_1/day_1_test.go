package day1_test

import (
	day1 "advent-of-code-2024/day_1"
	"advent-of-code-2024/tools"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Part1(t *testing.T) {
	day1Part1Input, readErr := tools.ReadFile("../inputs/day_1_part_1_example.txt")
	assert.Equal(t, nil, readErr, "readErr should be nil")
	totalDistance, getTotalDistanceErr := day1.CalculateTotalLocationDistances(day1Part1Input)
	assert.Equal(t, nil, getTotalDistanceErr, "getTotalDistanceErr should be nil")
	assert.Equal(t, 11, totalDistance, "totalDistance should be 11")
}

func TestDay1Part2(t *testing.T) {
	day1Part2Input, readErr := tools.ReadFile("../inputs/day_1_part_1_example.txt")
	assert.Equal(t, nil, readErr, "readErr should be nil")
	totalSimilarityScore, getScoreErr := day1.CalculateTotalSimilarityScore(day1Part2Input)
	assert.Equal(t, nil, getScoreErr, "getScoreErr should be nil")
	assert.Equal(t, 31, totalSimilarityScore, "totalSimilarityScore should be 31")
}
