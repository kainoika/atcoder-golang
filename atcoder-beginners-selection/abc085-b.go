package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &numbers[i])
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})

	count := 1
	for i := 0; i < n-1; i++ {
		if numbers[i] > numbers[i+1] {
			count++
		}
	}

	fmt.Println(count)
}
