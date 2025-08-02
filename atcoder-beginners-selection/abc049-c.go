package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)

	words := []string{"dreamer", "eraser", "dream", "erase"}

	for len(s) > 0 {
		matched := false
		for _, w := range words {
			if strings.HasSuffix(s, w) {
				s = s[:len(s)-len(w)]
				matched = true
				break
			}
		}
		if !matched {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
