package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	var divisors = make([]int, 21)
	for _, number := range numbers {
		var localDivisors = make([]int, number)
		for i := 2; number!=1; {
			if number%i == 0 {
				localDivisors[i-1] += 1
				number /= i
				i = 2
			} else {
				i++
			}
		}
		for index, number := range localDivisors {
			if divisors[index] < number {
				divisors[index] = number
			}
		}
	}

	total := 1
	for number, occurances := range divisors {
		if occurances != 0 {
			total *= int(math.Pow(float64(number+1),float64(occurances)))
		}
	}
	fmt.Println(total)
}
