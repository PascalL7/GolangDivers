package main

import (
	"fmt"
	"strings"
)

func main() {
	var W, H, N int
	fmt.Scan(&W, &H, &N)

	var x, y int
	fmt.Scan(&x, &y)

	var x1, y1, x2, y2 int
	x1, y1, x2, y2 = 0, 0, W -1, H -1


	for i := 0; i < N; i++ {
		var bombDir string
		fmt.Scan(&bombDir)

		if strings.Contains(bombDir, "U") {
			y2 = y - 1
		} else if strings.Contains(bombDir, "D") {
			y1 = y + 1
		}
	
		if strings.Contains(bombDir, "L") {
			x2 = x - 1
		} else if strings.Contains(bombDir, "R") {
			x1 = x + 1
		}
	
		x = x1 + (x2 - x1) / 2
		y = y1 + (y2 - y1) / 2
	
		fmt.Println(x, y)
	}
}

