package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) == len(b) {
		hammingDistance := 0
		for i := range a {
			if a[i] != b[i] {
				hammingDistance++
			}
		}
		return hammingDistance, nil
	} else {
		return 0, errors.New("The DNA strands match. One may be longer than the other.")
	}
}
