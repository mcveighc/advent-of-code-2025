package day2

import (
	"strconv"
	"strings"
)

type IdRange struct {
	Start int
	End   int
}

func GetInvalidIdTotalFromRanges(idRanges []string) int {
	total := 0

	for _, idRange := range idRanges {
		total += GetInvalidIdTotal(idRange)
	}

	return total
}

func GetInvalidIdTotal(idRange string) int {
	total := 0
	invalidIdArr := GetInvalidIdsInRange(idRange)

	for _, idValue := range invalidIdArr {
		total += idValue
	}

	return total

}

func IsInvalidId(id int) bool {
	idStr := strconv.Itoa(id)
	mid := len(idStr) / 2

	return idStr[:mid] == idStr[mid:]
}

func GetInvalidIdsInRange(idRangeStr string) []int {
	result := []int{}
	idRange := GetIdRange(idRangeStr)
	for i := idRange.Start; i <= idRange.End; i++ {
		if IsInvalidId(i) {
			result = append(result, i)
		}
	}
	return result
}

func GetIdRange(idRange string) IdRange {
	idRangeParts := strings.SplitN(idRange, "-", 2)

	idRangePartsStart, _ := strconv.Atoi(idRangeParts[0])
	idPartsRangeEnd, _ := strconv.Atoi(idRangeParts[1])

	return IdRange{Start: idRangePartsStart, End: idPartsRangeEnd}
}
