/**
 * @Author: Administrator
 * @Date: 2020-08-13 10:11
 * @Software: GoLand
 * @File: 01file.go
 * @Version: //TODO
 * @Description: //TODO
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func file01() {
	fileObj, err := os.Open("./01file.go")
	if err != nil {
		fmt.Println("文件打开错误")
	}
	defer fileObj.Close()
	var temp = bufio.NewReader(fileObj)
	msg, err := temp.ReadString('\n')
	if err == io.EOF {
		fmt.Println("EOF read")

	}
	if err != nil {
		fmt.Println("err read")

	}
	fmt.Printf("%s\n", msg)

}

func file02() {
	fileObj, err := os.Open("./01file.go")
	if err != nil {
		fmt.Println("文件打开错误")
	}

	defer func() {
		recover()

		err := fileObj.Close()
		if err != nil {
			panic("文件关闭错误")
		}

	}()
	var temp = make([]byte, 100, 100)
	_, err = io.ReadFull(fileObj, temp)
	if err == io.EOF {
		fmt.Println("EOF read")

	}
	if err != nil {
		fmt.Println("err read")

	}
	fmt.Printf("%s\n", temp)

}

func file03() {
	//fileObj,err:=os.Open("./01file.go")
	//if err != nil {
	//	fmt.Println("文件打开错误")
	//}
	//defer fileObj.Close()
	//msg,err:=ioutil.ReadAll(fileObj)

	msg, err := ioutil.ReadFile("./01file.go")

	if err == io.EOF {
		fmt.Println("EOF read")

	}
	if err != nil {
		fmt.Println("err read")

	}
	fmt.Printf("%s\n", msg)

}

func main() {
	//file01()
	file02()
	//file03()
}
