package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)

	count := 0
	for _, char := range s {
		if char == '1' {
			count++
		}
	}
	fmt.Println(count)
}
