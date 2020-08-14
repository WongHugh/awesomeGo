/**
 * @Author: Administrator
 * @Date: 2020-08-11 8:23
 * @Software: GoLand
 * @File: 01interface.go
 * @Version:
 * @Description:  接口
 */

package main

import (
	"fmt"
)

type Speaker interface {
	run()
}

type dog struct {
	name string
}

func newDog(name string) *dog {
	return &dog{name: name}
}

func (d *dog) run() {
	fmt.Printf("小狗%s跑起来了\n", d.name)
}

type cat struct {
	name string
}

func newCat(name string) *cat {
	return &cat{name: name}
}

func (c *cat) run() {
	fmt.Printf("小猫跳起来了%s\n", c.name)
}

func gogo(s Speaker) {
	s.run()
}
func main() {
	var c = newCat("lili")
	//var d = newDog("pp")

	var s Speaker
	s = c
	s.run()

}
