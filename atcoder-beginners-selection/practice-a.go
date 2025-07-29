package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanf("%d", &a)

	var b, c int
	fmt.Scanf("%d %d", &b, &c)a

	var s string
	fmt.Scanf("%s", &s)

	fmt.Printf("%d %s\n", a+b+c, s)
}
