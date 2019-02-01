package hamming

import "errors"

// Distance calculates the hamming distance of two strings with the same length
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("length of strings do not match")
	}

	sum := 0
	// Only test against length of a since both are the same length
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			sum++
		}
	}
	return sum, nil
}
