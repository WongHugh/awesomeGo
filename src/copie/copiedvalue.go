/**
 * @Author: Administrator
 * @Date: 2020-08-14 17:35
 * @Software: GoLand
 * @File: copiedvalue.go
 * @Version: //TODO
 * @Description: //TODO
 */

package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var countVal = new(atomic.Value)
	countVal.Store([]int{9, 7, 5, 3, 1})
	anotherStore(countVal)
	fmt.Printf("the count value:%+v \n", countVal.Load())
}

func anotherStore(coutVal *atomic.Value) {
	coutVal.Store([]int{2, 4, 6, 8})

}
