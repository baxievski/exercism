package prime

import "math"

// Nth returns the n-th prime number
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	primes := sieve(limit(n))
	if len(primes) < n {
		return 0, false
	}
	return primes[n-1], true
}

// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes#Algorithm_and_variants
func sieve(limit int) []int {
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

// https://en.wikipedia.org/wiki/Prime_number_theorem#Approximations_for_the_nth_prime_number
func limit(n int) int {
	if n < 5 {
		return 12
	}
	nf := float64(n)
	return int(nf*math.Log(nf) + nf*math.Log(math.Log(nf)))
}
