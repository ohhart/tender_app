package utils

import (
	"errors"
	"math"
)

func SafeIntToUint(i int) (uint, error) {
	if i < 0 {
		return 0, errors.New("integer overflow: cannot convert negative value to uint")
	}
	return uint(i), nil
}

func SafeUint64ToUint(val uint64) (uint, error) {
	if val > uint64(math.MaxUint32) {
		return 0, errors.New("value is greater than the maximum value of uint")
	}
	return uint(val), nil
}
