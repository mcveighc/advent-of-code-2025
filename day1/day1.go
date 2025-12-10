package day1

import (
	"strconv"
)

func FindZeroesFromRotations(rotations []string, dial int) int {
	zeroCount := 0
	for _, rotation := range rotations {
		startDial := dial

		direction := rotation[0]
		clicks, _ := strconv.Atoi(rotation[1:])

		// Add logic to count full 100 rotations
		zeroCount += clicks / 100

		switch direction {
		case 'L':
			dial = RotateLeft(startDial, clicks)
			if startDial != 0 && dial > startDial {
				zeroCount++
			}
		case 'R':
			dial = RotateRight(startDial, clicks)
			if dial != 0 && dial < startDial {
				zeroCount++
			}
		default:
			panic("Invalid direction")
		}

		if dial == 0 {
			zeroCount++
		}
	}
	return zeroCount
}

func RotateLeft(dial int, rotation int) int {
	projectedDialPosition := dial - rotation
	if projectedDialPosition < 0 {
		remainder := (-projectedDialPosition) % 100
		if remainder == 0 {
			return 0
		} else {
			return 100 - remainder
		}
	} else {
		return (dial - rotation) % 100
	}
}

func RotateRight(dial int, distance int) int {
	projectedDialPosition := dial + distance%100
	if projectedDialPosition > 99 {
		return projectedDialPosition - 100
	} else {
		return projectedDialPosition
	}
}
