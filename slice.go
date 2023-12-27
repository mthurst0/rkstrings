package rkstrings

import "fmt"

// RemoveFromSlice removes the element at idx efficiently by
// moving the last element. Order of the slice is not preserved.
func RemoveFromSlice(s []string, idx int) ([]string, error) {
	if idx >= len(s) || idx < 0 {
		return nil, fmt.Errorf("bad index: %d", idx)
	}
	s[idx] = s[len(s)-1]
	return s[:len(s)-1], nil
}

// RemoveFromSlicePreserveOrder removes the element at idx from the slice
// and preserves order.
func RemoveFromSlicePreserveOrder(s []string, idx int) ([]string, error) {
	if idx >= len(s) || idx < 0 {
		return nil, fmt.Errorf("bad index: %d", idx)
	}
	return append(s[:idx], s[idx+1:]...), nil
}
