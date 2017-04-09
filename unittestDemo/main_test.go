package main

import (
	"fmt"
	"testing"
	//install assert lib using
	//go get -u github.com/stretchr/testify/assert
	"github.com/stretchr/testify/assert"
)

func add(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	if add(1, 2) != 3 {
		t.FailNow()
	}
	assert.Equal(t, 3, add(1, 2))
}

func sub(x, y int) int {
	return x - y
}

func TestSub(t *testing.T) {
	if sub(2, 2) != 0 {
		t.FailNow()
	}
}

func multiply(x, y int) int {
	return x * y
}

func TestMultiply(t *testing.T) {
	if multiply(1, 2) != 2 {
		t.FailNow()
	}
}

func TestArraysEqual(t *testing.T) {
	a1 := [...]int{1, 2, 3}
	a2 := [...]int{1, 2, 3}
	fmt.Println("a1=", a1, ", a2=", a2)
	fmt.Printf("a1 address: %p, a2 address: %p\n", &a1, &a2)
	fmt.Println("Is a1 and a2 are equal? :", a1 == a2)
	sliceA1 := a1[1:3]
	sliceA2 := a2[1:3]
	fmt.Println("sliceA1=", sliceA1, ", sliceA2=", sliceA2)
	//fmt.Println("Is a2 and sliceA2 are equal? :", a2 == sliceA2) //will get compile error: mismatched type [3]int and []int

	assert := assert.New(t)
	assert.Equal(a1, a2, "using assert.Equal() to exam a1, a2 arrays")
	assert.EqualValues(a1, a2, "using assert.EqualValues() to exam a1, a2 arrays")
	assert.Equal(sliceA1, sliceA2, "using assert.Equal() to exam slice of a1 & slice of a2")
	assert.NotEqual(a1, sliceA2, "using assert.NotEqual() to exam a1 array and slice of a2")

	for _, val := range sliceA2 {
		assert.Contains(a2, val, "using assert.Contains() to exam a2 array and slice of a2")
		assert.Contains(a1, val, "using assert.Contains() to exam a1 array and slice of a2")
	}

	for _, val := range a2 {
		assert.Contains(a1, val, "using assert.Contains() to exam a1, a2 arrays")
	}
}
