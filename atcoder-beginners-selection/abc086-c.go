package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	times := make([]int, n)
	positions := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&times[i], &positions[i][0], &positions[i][1])
	}

	// 最初の位置は(0, 0)で、時刻は0
	currentTime := 0
	currentPosition := [2]int{0, 0}
	for i := 0; i < n; i++ {
		nextTime := times[i]
		nextPosition := positions[i]

		// 時間の差
		timeDiff := nextTime - currentTime
		// 位置の差
		xDiff := nextPosition[0] - currentPosition[0]
		yDiff := nextPosition[1] - currentPosition[1]

		// 移動に必要な時間は、xDiffとyDiffの絶対値の和
		moveTime := int(math.Abs(float64(xDiff))) + int(math.Abs(float64(yDiff)))

		// 移動に必要な時間が、与えられた時間の差より大きい場合は不可能
		// また、移動に必要な時間と時間の差の偶奇が異なる場合も不可能
		if moveTime > timeDiff || (timeDiff-moveTime)%2 != 0 {
			fmt.Println("No")
			return
		}

		currentTime = nextTime
		currentPosition = nextPosition
	}
	fmt.Println("Yes")
}
