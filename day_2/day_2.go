package day2

import (
	"advent-of-code-2024/tools"
	"math"
	"strconv"
	"strings"
)

func getReports(input string) ([][]int, error) {
	reportStrings := strings.Split(input, "\n")
	reports := make([][]int, len(reportStrings))

	for i, reportString := range reportStrings {
		levels, err := getReportLevels(reportString)
		if err != nil {
			return nil, err
		}
		reports[i] = levels
	}

	return reports, nil
}

func getReportLevels(reportString string) ([]int, error) {
	levelStrings := strings.Fields(reportString)
	levels := make([]int, 0, len(levelStrings))

	for _, levelString := range levelStrings {
		level, convErr := strconv.Atoi(levelString)
		if convErr != nil {
			return nil, convErr
		}
		levels = append(levels, level)
	}

	return levels, nil
}

func CalculateSafetyOfReports(input string) (int, error) {
	reports, err := getReports(input)
	if err != nil {
		return 0, err
	}

	var safetyCount int

	for i := range reports {
		isReportSafe := true

		levels := reports[i]
		lessThenFunc := func(a, b int) bool { return a < b }
		if !tools.SliceIsSortedAscOrDesc(levels, lessThenFunc) {
			isReportSafe = false
		}

		for j := 0; j < len(reports[i])-1; j++ {
			currentLevel := reports[i][j]
			nextLevel := reports[i][j+1]
			difference := int(math.Abs(float64(currentLevel - nextLevel)))
			if isSafe := difference >= 1 && difference <= 3; !isSafe {
				isReportSafe = false
				break
			}
		}
		if isReportSafe {
			safetyCount++
		}
	}

	return safetyCount, nil
}
