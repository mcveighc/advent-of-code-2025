package day1

import (
	"strconv"
)

func FindZeroesFromRotations(rotations []string, dial int) int {
	zeroCount := 0
	for _, rotation := range rotations {
		direction := rotation[0]
		dial = RotateDial(rotation, direction, dial)

		if dial == 0 {
			zeroCount++
		}
	}
	return zeroCount
}

func RotateDial(rotation string, direction byte, dial int) int {
	rotationTotal, err := strconv.Atoi(rotation[1:])

	switch direction {
	case 'L':
		dial = RotateLeft(dial, rotationTotal)
	case 'R':
		dial = RotateRight(dial, rotationTotal)
	default:
		panic(err)
	}
	return dial
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
