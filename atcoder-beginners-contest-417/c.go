package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &numbers[i])
	}

	count := 0
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if j-i == numbers[i-1]+numbers[j-1] {
				count++
			}
		}
	}

	fmt.Println(count)
}
