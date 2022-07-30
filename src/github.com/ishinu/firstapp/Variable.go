package main

import (
	"fmt"
	"strconv"
)

var outside int = 9

func main() {
	var f int = 80
	var o float32 = 2.3
	var k float32 = 2.2
	var j string = "40"
	i := 20
	fmt.Printf("%v,%T\n", k, k)
	fmt.Printf("%v,%T\n", j, j)
	fmt.Println(outside) // Variable declared outside main fn.
	fmt.Println(i, j)
	var l string
	l = strconv.Itoa(i) // Int to String
	fmt.Printf("%v, %T\n", l, l)
	var h float32
	h = float32(f) // Int to Float32
	fmt.Printf("%v, %T\n", h, h)
	var x int
	x = int(o) // Float32 to Int
	fmt.Printf("%v, %T\n", x, x)
}
