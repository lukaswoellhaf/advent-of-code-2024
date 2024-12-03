package day3

import (
	"regexp"
	"strconv"
)

func filterValidInstructions(input string) ([]string, error) {
	reg, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		return nil, err
	}

	return reg.FindAllString(input, -1), nil
}

func parseInstruction(instruction string) (int, error) {
	reg, err := regexp.Compile(`\d{1,3}`)
	if err != nil {
		return 0, err
	}

	numberStrings := reg.FindAllString(instruction, -1)
	if len(numberStrings) != 2 {
		return 0, err
	}

	var numbers []int

	for _, numberString := range numberStrings {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			return 0, nil
		}
		numbers = append(numbers, number)
	}

	return numbers[0] * numbers[1], nil
}

func ExecuteInstructions(input string) (int, error) {
	instructions, err := filterValidInstructions(input)
	if err != nil {
		return 0, err
	}

	instructionResults := 0
	for _, instruction := range instructions {
		result, err := parseInstruction(instruction)
		if err != nil {
			return 0, err
		}
		instructionResults += result
	}

	return instructionResults, nil
}
