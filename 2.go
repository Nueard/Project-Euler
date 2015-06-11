package main

import "fmt"

func main() {
	upTo := 4000000
	sum := 0
	prev := 1
	for i := 2; i < upTo; i, prev = i+prev, i {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
