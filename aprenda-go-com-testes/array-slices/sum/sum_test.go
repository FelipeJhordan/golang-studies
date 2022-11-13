package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		result := Sum(numbers)

		expected := 15

		if expected != result {
			t.Errorf("result %d, expected %d, data %v", result, expected, numbers)
		}
	})

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		result := Sum(numbers)

		expect := 6

		if result != expect {
			t.Errorf("result %d, expected %d, data %v", result, expect, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expect := []int{3, 9}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("result %v expect %v", result, expect)
	}
}

func TestSumAllRest(t *testing.T) {
	verifySums := func(t *testing.T, result, expect []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("result %v, expect %v", result, expect)
		}
	}
	t.Run("do the sums of some slices", func(t *testing.T) {
		result := SumAllRest([]int{1, 2}, []int{0, 9})
		expect := []int{2, 9}

		verifySums(t, result, expect)
	})

	t.Run("sum empty slices of security way", func(t *testing.T) {
		result := SumAllRest([]int{}, []int{3, 4, 5})
		expect := []int{0, 9}

		verifySums(t, result, expect)
	})
}
