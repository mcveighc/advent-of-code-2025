package day3

import (
	"strconv"
)

func GetTotalOutputJoltage(banks []string) int {
	totalJoltage := 0

	for _, bank := range banks {
		largestJoltage := FindLargestJoltageInBank(bank)
		totalJoltage += largestJoltage
	}

	return totalJoltage
}

func FindLargestJoltageInBank(bank string) int {
	joltagePartOne := 0
	joltagePartTwo := 0

	for i := 1; i < len(bank); i++ {
		partOneIndex := i - 1
		partTwoIndex := i

		tempOne, _ := strconv.Atoi(string(bank[partOneIndex]))
		tempTwo, _ := strconv.Atoi(string(bank[partTwoIndex]))

		if tempOne > joltagePartOne {
			joltagePartOne = tempOne
			joltagePartTwo = 0 // Reset part two when part one updates
		}

		if tempTwo > joltagePartTwo {
			joltagePartTwo = tempTwo
		}
	}

	return joltagePartOne*10 + joltagePartTwo
}
