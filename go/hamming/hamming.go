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
	}
	return 0, errors.New("strands do not match")
}
