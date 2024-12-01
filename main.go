package main

import (
	day1 "advent-of-code-2024/day_1"
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
	fmt.Print("--- Day 1: Historian Hysteria ---")
	fmt.Print("--- Part 1 ---")
	fmt.Printf("Result: %d", totalDistance)
}
