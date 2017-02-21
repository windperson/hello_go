package main

import "fmt"
// DebugMode is to define running in debug or not.
var DebugMode = true

func main() {
	f()
	fmt.Println("Returned normally from f().")
}

func f() {
	if DebugMode {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f() and getting r=", r)
			}
		}()
	}

	fmt.Println("calling g().")
	g(0)
	fmt.Println("Return normally from g().")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer func() {
		fmt.Printf("Defer in g(%d)\n", i)
	}()
	fmt.Printf("Printing in g(%d)\n", i)
	g(i + 1)
}
