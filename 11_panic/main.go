package main

import (
	"fmt"
)

func safelyDo(work int) {
	defer func() {
		fmt.Printf("work %d is done\n", work)
		if r := recover(); r != nil {
			fmt.Printf("Recover %v", r)
		}
	}()

	do(work)
}

func do(work int) {
	fmt.Printf("success %d\n", work)
	panic(fmt.Sprintf("ups %d", work))
}

func main() {
	workChan := make(chan int)
	defer close(workChan)

	go server(workChan)
}
