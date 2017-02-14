package main

import "fmt"

var g_items []string

func addItem(a string) {
	fmt.Println("g_items current Capacity",cap(g_items))
    g_items = append(g_items,a)
}

func printItems() {
	fmt.Println("g_items content:")
	for i := 0; i < len(g_items); i++ {
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
