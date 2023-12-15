package slicex

import (
	"fmt"
)

type INumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func NumberMax[T INumber](numbers []T) (T, error) {
	if len(numbers) == 0 {
		var zero T
		return zero, fmt.Errorf("empty slice")
	}
	max := numbers[0]
	for _, v := range numbers[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}

func NumberMin[T INumber](numbers []T) (T, error) {
	if len(numbers) == 0 {
		var zero T
		return zero, fmt.Errorf("empty slice")
	}
	min := numbers[0]
	for _, v := range numbers[1:] {
		if v < min {
			min = v
		}
	}
	return min, nil
}
