package main

import (
	"fmt"
)

func calcSum(r []int) int {
	count := 0

	defer func() {
		fmt.Printf("elements count: %d\n", count)
	}()

	sum := 0
	for _, rr := range r {
		sum += rr
		count++
	}

	return sum
}

func cc() (i int) {
	return 1
}

func main() {
	// r := []int{1, 2, 3, 4, 5}
	// sum := calcSum(r)

	sum := cc()
	fmt.Println(sum)
}
