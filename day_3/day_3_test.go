package day3_test

import (
	day3 "advent-of-code-2024/day_3"
	"advent-of-code-2024/tools"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3Part1(t *testing.T) {
	day3Part1Input, readErr := tools.ReadFile("../inputs/day_3_example.txt")
	assert.Equal(t, nil, readErr)
	instructionResults, getInstructionResultsErr := day3.ExecuteInstructions(day3Part1Input)
	assert.Equal(t, nil, getInstructionResultsErr)
	assert.Equal(t, 161, instructionResults)
}
