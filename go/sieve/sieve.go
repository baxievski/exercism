package sieve

// Sieve gives all primes up to `limit`
func Sieve(limit int) []int {
	primes := []int{}
	p := map[int]bool{}

	for i := 2; i <= limit; i++ {
		p[i] = true
	}

	for i := 2; i <= limit; i++ {
		if !p[i] {
			continue
		}
		primes = append(primes, i)
		for j := i; j <= limit; j += i {
			p[j] = false
		}
	}
	return primes
}
