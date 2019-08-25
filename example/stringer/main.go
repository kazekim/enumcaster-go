/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/kazekim/enumcaster-go/example/stringer/enum"
)

func main() {

	var wd enum.Weekday

	fmt.Println(wd)

	wd = enum.Saturday

	fmt.Println(wd)
}

