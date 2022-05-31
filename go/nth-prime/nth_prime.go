package prime

// Nth returns the nth prime
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}

	count := 1
	lastPrime := 2
	for i := 3; count < n; i += 2 {
		prime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			lastPrime = i
			count++
		}
	}
	return lastPrime, true
}
