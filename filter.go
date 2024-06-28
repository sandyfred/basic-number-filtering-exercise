package filter

import (
	"math"
)

type Condition func(int) bool

func FilterAll(numbers []int, conditions []Condition) (res []int) {
	for _, number := range numbers {
		all := true
		for _, condition := range conditions {
			if !condition(number) {
				all = false
				break
			}
		}
		if all {
			res = append(res, number)
		}
	}

	return
}

func FilterAny(numbers []int, conditions []Condition) (res []int) {
	for _, number := range numbers {
		all := false
		for _, condition := range conditions {
			if condition(number) {
				all = true
				break
			}
		}
		if all {
			res = append(res, number)
		}
	}

	return
}

func Prime(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func Odd(number int) bool {
	return number%2 != 0
}

func Even(number int) bool {
	return number%2 == 0
}

func MultipleOf(divisor int) Condition {
	return func(number int) bool {
		return number%divisor == 0
	}
}

func GreaterThan(target int) Condition {
	return func(number int) bool {
		return number > target
	}
}

func LessThan(target int) Condition {
	return func(number int) bool {
		return number < target
	}
}
