/**
 * @Author: Administrator
 * @Date: 2020-08-15 8:00
 * @Software: GoLand
 * @File: main.go
 * @Version: //TODO
 * @Description: //TODO
 */

package main

import (
	"fmt"
)

func main() {
	var ch1 = make(chan int, 1)
	for i := 0; i < 100; i++ {
		select {
		case x := <-ch1:
			fmt.Println(x)
		case ch1 <- i:
		}
	}
}
