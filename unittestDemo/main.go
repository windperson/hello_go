package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func subset(first, second []int) bool {
	set := make(map[int]int)
	for _, value := range second {
		set[value]++
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}

func main() {
	fmt.Println(subset([]int{1, 2, 3}, []int{1, 2, 3, 4}))
	fmt.Println(subset([]int{1, 2, 2}, []int{1, 2, 3, 4}))

	for i := 1; i <= 3; i++ {
		fmt.Println("1st:", EncodeToString(6))
		fmt.Println("2nd:", EncodeToString(6))
		fmt.Println("3rd:", EncodeToString(6))
		fmt.Println("4th:", EncodeToString(6))
	}
}

//EncodeToString generate random number of string based on input digit
func EncodeToString(maxDigit int) string {
	b := make([]byte, maxDigit)
	n, err := io.ReadAtLeast(rand.Reader, b, maxDigit)
	if n != maxDigit {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
