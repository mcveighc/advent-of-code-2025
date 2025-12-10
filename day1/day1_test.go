package day1

import (
	"bufio"
	"os"
	"testing"
)

func TestRotateLeft(t *testing.T) {
	dial := 50
	rotation := 3

	want := 47
	got := RotateLeft(dial, rotation)

	if got != want {
		t.Errorf("rotateDial(%q, %d) = %d; want %d", rotation, dial, got, want)
	}
}

func TestRotateLeftAtLeftBoundary(t *testing.T) {
	rotation := 1
	dial := 0

	want := 99
	got := RotateLeft(dial, rotation)

	if got != want {
		t.Errorf("rotateDial(%q, %d) = %d; want %d", rotation, dial, got, want)
	}
}

func TestRotateRight(t *testing.T) {
	rotation := 3
	dial := 50

	want := 53
	got := RotateRight(dial, rotation)

	if got != want {
		t.Errorf("rotateDial(%q, %d) = %d; want %d", rotation, dial, got, want)
	}
}

func TestRotateRightAtLeftBoundary(t *testing.T) {
	rotation := 1
	dial := 99
	want := 0
	got := RotateRight(rotation, dial)

	if got != want {
		t.Errorf("rotateDial(%q, %d) = %d; want %d", rotation, dial, got, want)
	}
}

func TestFindZeroesFromRotationsReturnsExpectedResult(t *testing.T) {
	rotations := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
	// rotations := []string{"L68", "L30", "R48"}
	dial := 50
	want := 6
	got := FindZeroesFromRotations(rotations, dial)

	if got != want {
		t.Errorf("FindZeroesFromRotations(%v, %d) = %d; want %d", rotations, dial, got, want)
	}
}

func TestFindZeroesFromRotationsWithInputFile(t *testing.T) {
	rotations := getInputRotations()
	dial := 50
	want := 6268
	got := FindZeroesFromRotations(rotations, dial)

	if got != want {
		t.Errorf("FindZeroesFromRotations(%v, %d) = %d; want %d", rotations, dial, got, want)
	}
}

func getInputRotations() []string {
	var rotations []string

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rotations = append(rotations, scanner.Text())
	}
	return rotations
}
