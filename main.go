package main

import (
	"fmt"
)

// Map -
func Map(s []int, f func(int) int) []int {
	r := make([]int, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Reduce -
func Reduce(s []int, initializer int, f func(int, int) int) int {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filter -
func Filter(s []int, f func(int) bool) []int {
	var r []int
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}

	f := func(i int) int {
		return i * i
	}

	g := func(i, j int) int {
		return i + j
	}

	h := func(i int) bool {
		return i > 5
	}

	fmt.Println(Map(s, f))
	fmt.Println(Reduce(s, 0, g))
	fmt.Println(Filter(s, h))
}
