package day2_test

import (
	day2 "advent-of-code-2024/day_2"
	"advent-of-code-2024/tools"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Part1(t *testing.T) {
	day2Part1Input, readErr := tools.ReadFile("../inputs/day_2_example.txt")
	assert.Equal(t, nil, readErr, "readErr should be nil")
	safetyCount, getSafetyCountErr := day2.CalculateSafetyOfReports(day2Part1Input)
	assert.Equal(t, nil, getSafetyCountErr)
	assert.Equal(t, 2, safetyCount)
}
