package main

import "fmt"

func main() {
	var aVar int32 = 15
	fmt.Println("Address of a", &aVar)
	aVar += 1
	b := aVar
	fmt.Println("new value of a=", aVar, "and address of a=", &aVar)
	fmt.Println("new value of b=", b, "and address of a=", &b)
}
