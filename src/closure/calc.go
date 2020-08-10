/**
 * @Author: Administrator
 * @Date: 2020-08-10 19:45
 * @Software: GoLand
 * @File: calc.go
 * @Version: //TODO
 * @Description: //TODO
 */

package main

import "fmt"

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func main() {
	f1,f2:=calc(10)
	fmt.Println(f1(1),f2(2))
	fmt.Println(f1(3),f2(4))
	fmt.Println(f1(5),f2(6))
}
