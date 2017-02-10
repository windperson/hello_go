package main

import (
	"fmt"

	strutil "github.com/windperson/hello_go/stringutil"
)

func main() {
	var str = "!oG ,olleH"
	fmt.Println("str=", str)
	var revStr = strutil.Reverse(str)
	fmt.Println("revStr=", revStr)
}
