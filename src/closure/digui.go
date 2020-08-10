package main

import "fmt"

func F(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * F(n-1)

}

func main() {
	d := F(6)
	fmt.Println(d)
}
