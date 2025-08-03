package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	count := make(map[int]int)
	for i := 0; i < n; i++ {
		count[i+1+A[i]]++
	}
	fmt.Println(count)

	ans := 0
	for j := 0; j < n; j++ {
		ans += count[j+1-A[j]]
	}
	fmt.Println(ans)
}
