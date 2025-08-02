package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	numbers_a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &numbers_a[i])
	}

	numbers_b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &numbers_b[i])
	}

	for _, num_b := range numbers_b {
		for i, num_a := range numbers_a {
			if num_a == num_b {
				numbers_a = append(numbers_a[:i], numbers_a[i+1:]...)
				break
			}
		}
	}
	if len(numbers_a) != 0 {
		r := make([]string, len(numbers_a))
		for i, v := range numbers_a {
			r[i] = strconv.Itoa(v)
		}
		fmt.Println(strings.Join(r, " "))
	}
}
