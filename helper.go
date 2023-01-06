package qris

import "strconv"

func inRange(valS, minS, maxS string) bool {
	val, err := strconv.Atoi(valS)
	if err != nil {
		return false
	}
	min, err := strconv.Atoi(minS)
	if err != nil {
		return false
	}
	max, err := strconv.Atoi(maxS)
	if err != nil {
		return false
	}

	return val >= min && val <= max
}
