package main

import "fmt"

const gCapicity int = 5

var g_items [gCapicity]string
var g_length = 0

func addItem(a string) {
	if g_length < gCapicity {
		g_items[g_length] = a
		g_length++
	} else {
		fmt.Println("Already full!")
	}
}

func printItems() {
	fmt.Println("g_items content:")
	for i := 0; i < g_length; i++ {
		fmt.Println(g_items[i])
	}
}

func main() {
	addItem("A item")
	addItem("B item")
	addItem("C item")
	addItem("D item")
	addItem("E item")
	addItem("F item")
	printItems()
}
