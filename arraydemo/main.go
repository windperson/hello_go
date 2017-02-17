package main

import "fmt"

var gItems []string

func addItem(a string) {
	fmt.Println("g_items current Capacity", cap(gItems))
	gItems = append(gItems, a)
}

func printItems() {
	fmt.Println("g_items content:")
	for i := 0; i < len(gItems); i++ {
		fmt.Println(gItems[i])
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
