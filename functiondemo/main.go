package main

import "fmt"

func printWow(times int) {
	for i := 0; i < times; i++ {
		fmt.Println("Wow")
	}
}

func max(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func getMax(numbers ...int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}

	var max = numbers[0]
	for _, value := range numbers[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

func maxWithPointer(i int, j int, k *int) int {
	if i > j {
		*k = i
	} else {
		*k = j
	}
	fmt.Println("k inside func=", k)
	return *k
}

func main() {
	var i, j = 100, 15
	var ret int
	fmt.Println("&ret=", &ret)
	// fmt.Println("max of ",i," and ",j," is ",max(i,j))
	fmt.Println("max of ", i, " and ", j, " is ", maxWithPointer(i, j, &ret))

	var data0 = []int{0, 100, -100}
	max := getMax(data0...)
	fmt.Println("max of ", data0, " is:", max)

    var data1 = []int{-100}
    max = getMax(data1...)
    fmt.Println("max of ", data1, " is:", max)

    var data2 = []int{+100,-100}
    max = getMax(data2...)
    fmt.Println("max of ", data2, " is:", max)
}
