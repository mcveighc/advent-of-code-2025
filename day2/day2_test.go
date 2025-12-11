package day2

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestIsInvalidIdWithInvalidSeqOfTwoReturnsTrue(t *testing.T) {
	// Assemble
	id := 11
	expected := true

	// Act
	actual := IsInvalidId(id)

	// Assert
	if actual != expected {
		t.Errorf("IsInvalidId(%d) = %t; expected %t", id, actual, expected)
	}
}

func TestIsInvalidIdWithInvalidSeqOfThreeReturnsTrue(t *testing.T) {
	// Assemble
	id := 123123123
	expected := true

	// Act
	actual := IsInvalidId(id)

	// Assert
	if actual != expected {
		t.Errorf("IsInvalidId(%d) = %t; expected %t", id, actual, expected)
	}
}

func TestIsInvalidIdWithInvalidSeqOfFiveReturnsTrue(t *testing.T) {
	// Assemble
	id := 2121212121
	expected := true

	// Act
	actual := IsInvalidId(id)

	// Assert
	if actual != expected {
		t.Errorf("IsInvalidId(%d) = %t; expected %t", id, actual, expected)
	}
}

func TestIsInvalidIdWithInvalidSeqOfSevenReturnsTrue(t *testing.T) {
	// Assemble
	id := 21212121212121
	expected := true

	actual := IsInvalidId(id)

	if actual != expected {
		t.Errorf("IsInvalidId(%d) = %t; expected %t", id, actual, expected)
	}
}

func TestIsInvalidIdReturnsFalseForInvalidId(t *testing.T) {
	// Assemble
	id := 12
	expected := false

	// Act
	actual := IsInvalidId(id)

	// Assert
	if actual != expected {
		t.Errorf("IsInvalidId(%d) = %t; expected %t", id, actual, expected)
	}
}

func TestGetInvalidIdsInRangeReturnsExpectedResult(t *testing.T) {
	// Assemble
	idRange := "95-115"
	expected := []int{99}

	// Act
	actual := GetInvalidIdsInRange(idRange)

	// Assert
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}

func TestGetIdRangeReturnsExpectedResult(t *testing.T) {
	idRange := "11-22"
	expected := IdRange{Start: 11, End: 22}

	actual := GetIdRange(idRange)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}

func TestGetInvalidIdTotalReturnsExpectedResult(t *testing.T) {
	idRange := "11-22"
	expected := 33

	actual := GetInvalidIdTotal(idRange)

	if actual != expected {
		t.Errorf("Expected %d, actual %d", expected, actual)
	}
}

func TestGetInvalidIdTotalFromRangesReturnsExpectedResult(t *testing.T) {
	ranges := []string{"11-22", "95-115", "998-1012", "1188511880-1188511890", "222220-222224",
		"1698522-1698528", "446443-446449", "38593856-38593862", "565653-565659",
		"824824821-824824827", "2121212118-2121212124"}
	expected := 4174379265

	actual := GetInvalidIdTotalFromRanges(ranges)

	if actual != expected {
		t.Errorf("Expected %d, actual %d", expected, actual)
	}
}

func TestGetInvalidIdtotalFromInputReturnsExpectedResult(t *testing.T) {
	ranges := getInputIdRanges()
	expected := 41662374059

	actual := GetInvalidIdTotalFromRanges(ranges)

	if actual != expected {
		t.Errorf("Expected %d, actual %d", expected, actual)
	}
}

func getInputIdRanges() []string {
	content, _ := os.ReadFile("input.txt")

	// Convert to string and split by comma
	line := strings.TrimSpace(string(content))
	values := strings.Split(line, ",")

	// Optional: trim whitespace from each value
	for i, v := range values {
		values[i] = strings.TrimSpace(v)
	}

	fmt.Println("Values:", values)
	fmt.Println("Count:", len(values))

	return values
}
