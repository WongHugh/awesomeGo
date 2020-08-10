package main

import (
	"fmt"
	"time"
)

func fibonacci02(n int) int {
	for i := 2; i <= n; i++ {
		fibarry[2] = fibarry[0] + fibarry[1]
		fibarry[0] = fibarry[1]
		fibarry[1] = fibarry[2]
	}
	return fibarry[2]
}

func main() {
	start := time.Now()
	d := fibonacci02(10)
	total := time.Since(start)
	fmt.Println(d)
	fmt.Println(total)
}
