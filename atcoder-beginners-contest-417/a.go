package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	var s string
	fmt.Scanf("%s", &s)

	s = s[a:]
	s = s[0 : len(s)-b]

	fmt.Println(s)
}
