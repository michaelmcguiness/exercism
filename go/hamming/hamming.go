// Package hamming provides primitives for calculating
// the Hamming Distance between two DNA strands.
package hamming

import (
	"errors"
)

// Distance calculate the Hamming Distance between two DNA strands.
func Distance(a, b string) (int, error) {
	dist := 0
	if len(a) != len(b) {
		return 0, errors.New("hamming distance is only defined for sequences of equal length")
	}
	for pos, char := range a {
		if char != int32(b[pos]) {
			dist++
		}
	}
	return dist, nil
}
