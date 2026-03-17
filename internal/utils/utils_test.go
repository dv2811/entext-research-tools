package utils

import (
	"testing"
)

func isEqual[T comparable](result, target []T) bool {
	if len(result) != len(target) {
		return false
	}
	for indx := range target {
		if result[indx] != target[indx] {
			return false
		}
	}
	return true
}

// return an unique
func TestUniqueFilter(t *testing.T) {
	testCases := [][]string{
		[]string{"cat", "cat", "dog", "meow", "duck"},
		[]string{"cat", "meow", "dog", "meow", "duck"},
	}
	expected := [][]string{
		[]string{"cat", "dog", "meow", "duck"},
		[]string{"cat", "meow", "dog", "duck"},
	}
	for indx := range testCases {
		result := ApplyUniqueFilter(testCases[indx])
		if !isEqual(result, expected[indx]) {
			t.Errorf("Values mismatch. Test case %d, result: %v, expected %v", indx, result, expected[indx])
		}
	}
}
