package main

import (
	"fmt"
	"time"
)

func fibonacci(n uint64) uint64 {
	if n >= 3 {
		return fibonacci(n-1) + fibonacci(n-2)
	}
	return 1

}

func main() {
	start := time.Now()
	d := fibonacci(50)
	total := time.Since(start)
	fmt.Println(d)
	fmt.Println(total)
}
