package innValidator

import (
  "testing"
)

var intStringToIntArrayFixtures = []struct {
  string string
  array  []int
  error  bool
}{
  {string: "0", array: []int{0}, error: true},
  {string: "01", array: []int{0, 1}, error: false},
  {string: "", array: []int{}, error: true},
  {string: "foo", array: []int{}, error: true},
  {string: "foo42", array: []int{}, error: true},
  {string: "42foo", array: []int{}, error: true},
}

func Test_intStringToIntArray(t *testing.T) {
  for _, value := range intStringToIntArrayFixtures {
    array, _ := intStringToIntArray(value.string)

    if !isSliceEqual(array, value.array) {
      t.Errorf("Huston, we have problem")
    }
  }
}

func isSliceEqual(left, right []int) bool {
  if len(left) != len(right) {
    return false
  }

  for index := range left {
    if left[index] != right[index] {
      return false
    }
  }

  return true
}