package main

import "fmt"

const o int64 = 2334

const (
	c1 = iota
	c2 = iota
	c3 = iota
	// Same result for
	// c2 ( With no assignment )
	// c3 ( With no assignment )
)

const (
	l2 = iota // iota is scoped to a constant block.
)

const (
	// enumerated constants
	// handle the 0 value error by assigning errorSpecialist = iota and leaving below constants unassigned.
	// Also _=iota would mean almost the same but unlike the above, it throws away 0 value assigned to it.
	_ = iota + 5 // for a fixed offset
	checkIota
	catSpecialist = iota
	dogSpecialist
	snakeSpecialist
)

const (
	_  = iota // Ignores the first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	// Constant
	// Immutable but can be shadowed. Shadowing means redeclaring an existing constant from outer scope
	// with a different value and a different type.
	// Value must be calculable, can't assign functions to evaluate for geting value.
	// External constant prefer PascalCase, Internal constant presfer CamelCase
	const myConst int = 42
	const a string = "Hey you, good day"
	const b float32 = 3.14
	const c bool = true
	const o int = 10
	var g int = 29
	fmt.Printf("%v,%T\n", myConst, myConst)
	fmt.Printf("%v,%T\n", a, a)
	fmt.Printf("%v,%T\n", b, b)
	fmt.Printf("%v,%T\n", c, c)
	fmt.Printf("%v,%T\n", o, o) // Change type since inner constant shadows outer constant.
	fmt.Printf("%v,%T\n", o+g, o+g)
	const l = iota
	fmt.Printf("%v,%T\n", l, l)
	fmt.Printf("%v,%T\n", c1, c1)
	fmt.Printf("%v,%T\n", c2, c2)
	fmt.Printf("%v,%T\n", c3, c3)
	fmt.Printf("%v,%T\n", l2, l2)
	var specialistType int = catSpecialist
	fmt.Printf("%v\n", specialistType == catSpecialist)
	fmt.Printf("%v\n", checkIota)
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)
}
