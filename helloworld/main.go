package main

import "fmt"

func main() {
	var aVar int32 = 15
	fmt.Println("Hello Go!\r\nAddress of a", &aVar)
	aVar += 1
	b := int(aVar)
	fmt.Println("new value of a=", aVar, "and address of a=", &aVar)
	fmt.Println("new value of b=", b, "and address of a=", &b)
}
