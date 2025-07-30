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

	alice_sum := 0
	for i := 0; i < n; i += 2 {
		alice_sum += numbers[i]
	}

	bob_sum := 0
	for i := 1; i < n; i += 2 {
		bob_sum += numbers[i]
	}

	fmt.Println(alice_sum - bob_sum)
}
