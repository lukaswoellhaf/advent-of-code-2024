package main

import (
	day1 "advent-of-code-2024/day_1"
	day2 "advent-of-code-2024/day_2"
	"advent-of-code-2024/tools"
	"fmt"
	"log"
)

func main() {
	day1Part1Input, readErr := tools.ReadFile("./inputs/day_1_part_1_puzzle.txt")
	if readErr != nil {
		log.Fatal(readErr)
	}
	totalDistance, getSumErr := day1.CalculateTotalLocationDistances(day1Part1Input)
	if getSumErr != nil {
		log.Fatal(getSumErr)
	}
	fmt.Println("--- Day 1: Historian Hysteria ---")
	fmt.Println("--- Part 1 ---")
	fmt.Printf("Result: %d\n", totalDistance)

	day1Part2Input, readErr := tools.ReadFile("./inputs/day_1_part_2_puzzle.txt")
	if readErr != nil {
		log.Fatal(readErr)
	}
	totalSimilarityScore, getScoreErr := day1.CalculateTotalSimilarityScore(day1Part2Input)
	if getScoreErr != nil {
		log.Fatal(getScoreErr)
	}
	fmt.Println("--- Part 2 ---")
	fmt.Printf("Result: %d\n", totalSimilarityScore)

	day2Input, readErr := tools.ReadFile("./inputs/day_2_puzzle.txt")
	if readErr != nil {
		log.Fatal(readErr)
	}
	safetyCount, getSafetyCountErr := day2.CalculateSafetyOfReports(day2Input)
	if getSafetyCountErr != nil {
		log.Fatal(getSafetyCountErr)
	}
	fmt.Println("--- Day 2: Red-Nosed Reports---")
	fmt.Println("--- Part 1 ---")
	fmt.Printf("Result: %d\n", safetyCount)
}
