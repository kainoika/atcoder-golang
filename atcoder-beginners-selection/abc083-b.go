package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	sum := 0
	for i := 1; i <= n; i++ {
		digitSum := 0
		num := i
		for num > 0 {
			digitSum += num % 10
			num /= 10
		}
		if digitSum >= a && digitSum <= b {
			sum += i
		}
	}
	fmt.Println(sum)

}
