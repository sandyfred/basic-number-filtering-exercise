package filter

import (
	"reflect"
	"testing"
)

func TestFilterAll(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	tests := []struct {
		name       string
		conditions []Condition
		want       []int
	}{
		{name: "Even", conditions: []Condition{Even}, want: []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}},
		{name: "Odd", conditions: []Condition{Odd}, want: []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}},
		{name: "Prime", conditions: []Condition{Prime}, want: []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{name: "OddPrime", conditions: []Condition{Odd, Prime}, want: []int{3, 5, 7, 11, 13, 17, 19}},
		{name: "EvenMultiplesOf5", conditions: []Condition{Even, MultipleOf(5)}, want: []int{10, 20}},
		{name: "OddMultiplesOf3GreaterThan10", conditions: []Condition{Odd, MultipleOf(3), GreaterThan(10)}, want: []int{15}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterAll(numbers, tt.conditions)
			checkArrays(t, got, tt.want)
		})
	}
}

func TestFilterAny(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	conditions := []Condition{LessThan(6), MultipleOf(3)}
	got := FilterAny(numbers, conditions)
	want := []int{1, 2, 3, 4, 5, 6, 9, 12, 15, 18}
	checkArrays(t, got, want)
}

func checkArrays(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
