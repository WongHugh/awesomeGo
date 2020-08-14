/**
 * @Author: Administrator
 * @Date: 2020-08-14 8:35
 * @Software: GoLand
 * @File: logtest.go
 * @Version: //TODO
 * @Description: //TODO
 */

package main

import logs "awesomeGo/src/logging"

func main() {
	log1 := logs.NewLogging()
	log1.Error("debug test")
}
