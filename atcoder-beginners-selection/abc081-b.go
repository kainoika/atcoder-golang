package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &numbers[i])
	}

	count := 0

	for {
		allEven := true
		for i := 0; i < n; i++ {
			if numbers[i]%2 != 0 {
				allEven = false
				break
			}
		}
		if !allEven {
			break
		}
		for i := 0; i < n; i++ {
			numbers[i] /= 2
		}
		count++
	}

	fmt.Println(count)

}
