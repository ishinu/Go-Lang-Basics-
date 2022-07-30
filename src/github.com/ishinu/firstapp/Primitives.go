package main

import "fmt"

func main() {
	// 1. Boolean
	var n bool = true
	var f bool = false
	fmt.Printf("%v,%T\n", n, n)
	fmt.Printf("%v,%T\n", f, f)
	a := 1 == 1 // Expression and auto initialization.
	b := 1 == 2
	fmt.Printf("%v,%T\n", a, a)
	fmt.Printf("%v,%T\n", b, b)
	var m bool // By default 0 value, For boolean it's false.
	// Boolean is its own type. Not an alias for integer
	fmt.Printf("%v,%T\n", m, m)

	// 2. Integers
	var p uint16 = 42
	fmt.Printf("%v,%T\n", p, p)
	x := 10
	y := 3
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)
	// Adding int + int8 is not allowed. Type conversion req.
	// Go hopes to stay far away from implicit data conversion

	// Float
	var d float32 = 3.14
	d = 2.1e14
	fmt.Printf("%v,%T\n", d, d)
	var j float64 = 3.14
	j = 13.8e82
	fmt.Printf("%v,%T\n", j, j)

	// String ( UTF-8 )
	s := "Hello there, i am a string"
	fmt.Printf("%v,%T\n", s, s)
	fmt.Printf("%v,%T\n", s[4], s[4])
	fmt.Printf("%v,%T\n", string(s[4]), string(s[4]))
	s2 := " and you are human being"
	fmt.Printf("%v,%T\n", s+s2, s+s2) // Concatenation.

	// Slice
	c := []byte(s)
	fmt.Printf("%v,%T\n", c, c) // uint8 is a type alias of byte
	// Runes ( UTF-32 )
	r := 'a'                    // or var r rune = 'a'
	fmt.Printf("%v,%T\n", r, r) // runes are type alias for int32
}
