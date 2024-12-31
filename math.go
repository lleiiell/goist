package goist

import (
	"fmt"
	"math"
	"strconv"
)

// IdMask Generate irregular, reversible masks
func IdMask(id int64) (str string) {
	return fmt.Sprintf("%s%s", strconv.FormatInt(id, 36), RandStr(3))
}

func IdUnmask(mask string) (id int64) {
	runeSli := []rune(mask)
	if len(runeSli) <= 3 {
		return 0
	}

	id, _ = strconv.ParseInt(string(runeSli[:len(runeSli)-3]), 36, 64)
	return
}

func Round2NearestHalf(num float64) float64 {
	return math.Round(num*2) / 2
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64 |
	~complex64 | ~complex128
}

func Sum[T Number](numbers []T) T {
	var total T

	for _, v := range numbers {
		total += v
	}

	return total
}
