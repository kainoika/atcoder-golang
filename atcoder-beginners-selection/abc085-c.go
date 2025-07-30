package main

import (
	"fmt"
)

func main() {
	var n, y int
	fmt.Scanf("%d %d", &n, &y)

	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			// kは一意に定まる(nからiとjを引いた残りの数)ため、ループは不要
			// ループにしないことで計算量がO(n^3)からO(n^2)に減る
			k := n - i - j
			if 10000*i+5000*j+1000*k == y {
				fmt.Printf("%d %d %d\n", i, j, k)
				return
			}
		}
	}

	fmt.Println("-1 -1 -1")
}
