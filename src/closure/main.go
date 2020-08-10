/**
 * @Author: Administrator
 * @Date: 2020-08-10 19:07
 * @Software: GoLand
 * @File: main.go
 * @Description: //TODO
 */

package main

import (
	"fmt"
	"strings"
)

func makeSuffixFuc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFuc(".jpg")
	txtFunc := makeSuffixFuc(".txt")
	fmt.Println(jpgFunc("test"))
	fmt.Println(jpgFunc("test1.jpg"))
	fmt.Println(txtFunc("txt.txt"))
}
