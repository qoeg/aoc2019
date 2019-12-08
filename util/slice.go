package util

import (
	"errors"
)

// Contains checks a slice for an item.
// If found, it returns true and the item's index.
// If not found, it returns false and -1.
func Contains(slice interface{}, item interface{}) (bool, int) {
	if s, ok := slice.([]interface{}); ok {
		for i := range s {
			if s[i] == item {
				return true, i
			}
		}
		return false, -1
	}

	return false, -1
}

// Interact allows a caller to take an action on an item in a slice if the item exists
func Interact(slice interface{}, item interface{}, act func(int, interface{})) error {
	if s, ok := slice.([]interface{}); ok {
		for i, v := range s {
			if v == item {
				act(i, v)
			}
		}
	} else {
		return errors.New("util: unrecognized slice")
	}
	
	return nil
}

// Permutations generates all of the permutations for the elements in the given slice
func Permutations(slice []int) [][]int {
	result := [][]int{}

	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			result = append(result, append([]int{}, a...))
			return
		}

		for i := k; i < len(slice); i++ {
			a[k], a[i] = a[i], a[k]
			rc(a, k+1)
			a[k], a[i] = a[i], a[k]
		}
	}
	rc(slice, 0)
	
	return result
}
