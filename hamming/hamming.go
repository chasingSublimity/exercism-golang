// Package hamming exposes the Distance function
package hamming

import "errors"

// Distance calculates the Hamming distance of two nucleotide strings of equal length
// If the two strings are of unequal length, an error is returned
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Error: nucleotide strings must be of equal length")
	}

	var d int
	for i := range a {
		if a[i] != b[i] {
			d++
		}
	}

	return d, nil

}
