package main

import (
	"testing"
)

func add(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	if add(1, 2) != 3 {
		t.FailNow()
	}
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
