package main

import "fmt"
import "math"

func main() {
	largestPrime := 1
	number := 600851475143
	upTo := int(math.Sqrt(float64(number)))
	numbers := make([]bool, upTo)
	numbers[0], numbers[1] = true, true
	for i := 1; i < upTo; i++ {
		if numbers[i] {
			continue
		}
		if number%i == 0 {
			largestPrime = i
		}
		for j := i; j < upTo; j += i {
			numbers[j] = true
		}
	}
	fmt.Println(largestPrime)
}
