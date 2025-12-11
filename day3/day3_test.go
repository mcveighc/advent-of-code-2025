package day3

import (
	"bufio"
	"os"
	"testing"
)

func TestFindLargestJoltageInBank(t *testing.T) {
	bank := "987654321111111"
	expected := 98

	actual := FindLargestJoltageInBank(bank)

	if actual != expected {
		t.Errorf("FindLargestJoltageInBank(%q) = %d; want %d", bank, actual, expected)
	}
}

func TestFindLargestJoltageAtBoundaryInBank(t *testing.T) {
	bank := "811111111111119"
	expected := 89

	actual := FindLargestJoltageInBank(bank)

	if actual != expected {
		t.Errorf("FindLargestJoltageInBank(%q) = %d; want %d", bank, actual, expected)
	}
}

func TestFindLargestJoltageWithLastTwoNumbers(t *testing.T) {
	bank := "234234234234278"
	expected := 78

	actual := FindLargestJoltageInBank(bank)

	if actual != expected {
		t.Errorf("FindLargestJoltageInBank(%q) = %d; want %d", bank, actual, expected)
	}
}

func TestFindLargestJoltageWithPartTwoReset(t *testing.T) {
	bank := "818181911112111"
	expected := 92

	actual := FindLargestJoltageInBank(bank)

	if actual != expected {
		t.Errorf("FindLargestJoltageInBank(%q) = %d; want %d", bank, actual, expected)
	}
}

func TestGetTotalOutputJoltageReturnsExpectedTotal(t *testing.T) {
	banks := []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}
	expected := 357

	actual := GetTotalOutputJoltage(banks)

	if actual != expected {
		t.Errorf("GetTotalOutputJoltage(%q) = %d; want %d", banks, actual, expected)
	}
}

func TestGetTotalOutputJoltageFromInputReturnsExpectedTotal(t *testing.T) {
	banks := getInputBanks()
	expected := 357

	actual := GetTotalOutputJoltage(banks)

	if actual != expected {
		t.Errorf("GetTotalOutputJoltage(%q) = %d; want %d", banks, actual, expected)
	}
}

func getInputBanks() []string {
	var rotations []string

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rotations = append(rotations, scanner.Text())
	}
	return rotations
}
