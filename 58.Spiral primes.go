package main

import (
	"fmt"
	"math"
)

const N = 1000000000

func main() {
	primes := getPrimesLessThanN()
	primesMap := make(map[int]bool)
	for _, x := range primes {
		primesMap[x] = true
	}

	var totalPrimes = 8
	var totalDiagonal = 13
	var currentSquare = 7
	for true {
		for i := 1; i <= 4; i++ {
			diagonal := currentSquare*currentSquare + (currentSquare+1)*i
			if primesMap[diagonal] {
				totalPrimes++
			}
		}
		totalDiagonal += 4
		currentSquare += 2
		if totalPrimes*100.0/totalDiagonal < 10 {
			fmt.Println(currentSquare)
			break
		}
	}

}

func getPrimesLessThanN() []int {
	var x, y, n int
	nsqrt := math.Sqrt(N)

	is_prime := [N]bool{}

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= N && (n%12 == 1 || n%12 == 5) {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) + y*y
			if n <= N && n%12 == 7 {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= N && n%12 == 11 {
				is_prime[n] = !is_prime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if is_prime[n] {
			for y = n * n; y < N; y += n * n {
				is_prime[y] = false
			}
		}
	}

	is_prime[2] = true
	is_prime[3] = true

	primes := make([]int, 0, 1270606)
	for x = 0; x < len(is_prime)-1; x++ {
		if is_prime[x] {
			primes = append(primes, x)
		}
	}

	return primes
}
