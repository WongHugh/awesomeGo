/**
 * @Author: Administrator
 * @Date: 2020-08-14 8:25
 * @Software: GoLand
 * @File: logPacket.go
 * @Description: //TODO
 */

package logging

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

type Ilogging interface {
	Error(s string)
}

type Logging struct{}

func (l Logging) Error(s string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("runtime error")
		return
	}
	now := time.Now()
	timestamp := now.Format("2006/01/02 15:04:05:000")
	call := runtime.FuncForPC(pc).Name()
	baseFile := path.Base(file)
	fmt.Printf("%s Error %v line:%v %v %v", timestamp, baseFile, line, call, s)

}

func NewLogging() *Logging {
	return &Logging{}
}
