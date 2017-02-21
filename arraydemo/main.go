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

func makeslice() []int {
	s := make([]int, 0, 10)
	s = append(s, 100)
	return s
}

func makeMap() map[string]int {
	var m = make(map[string]int)
	m["a"] = 1
	return m
}

func main() {
	addItem("A item")
	addItem("B item")
	addItem("C item")
	addItem("D item")
	addItem("E item")
	addItem("F item")
	printItems()

	var s = makeslice()
	println(s[0])

	var m = makeMap()
	fmt.Println(m)
	var x, ok = m["b"]
	fmt.Println(x, ok)
}
