package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%d", &a)

	var b int
	fmt.Scanf("%d", &b)

	var c int
	fmt.Scanf("%d", &c)

	var x int
	fmt.Scanf("%d", &x)

	// 500 円玉をA 枚、100 円玉をB 枚、50 円玉をC 枚持っているものとする
	// 合計金額が丁度X円になるパターンが幾つあるか算出する
	count := 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if 500*i+100*j+50*k == x {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
