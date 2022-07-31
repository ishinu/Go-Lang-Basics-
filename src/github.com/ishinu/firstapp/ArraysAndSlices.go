package main

import "fmt"

func main() {
	grade := [3]int{97, 92, 12}
	grades := [...]int{97, 92, 12}
	var students [3]string
	var childrens [3]string
	fmt.Printf("Grades : %v\n", grade)
	fmt.Printf("Grades : %v\n", grades)
	fmt.Printf("Students : %v\n", students)
	students[0] = "Pume"
	fmt.Printf("Students : %v\n", students)
	childrens[0] = "Krism"
	childrens[1] = "Rosh"
	childrens[2] = "Oaash"
	fmt.Printf("Childrens: %v\n", childrens)
	fmt.Printf("Childrens[2]: %v\n", childrens[2])
	fmt.Printf("Number of Childrens: %v\n", len(childrens))
	// An example from linear algebra of linear matrix.
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)
	// Another way to declare similar values.
	var identityMatrixv2 [3][3]int
	identityMatrixv2[0] = [3]int{1, 0, 0}
	identityMatrixv2[1] = [3]int{0, 1, 0}
	identityMatrixv2[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrixv2)
	// Assigning a variable values of an array. ( copy of array )
	a := [...]int{1, 2, 3}
	b := a  // Gets the copy of a and not pointing to the address of 'a' array.
	c := &a // Points to address of 'a' array
	b[1] = 10
	c[1] = 20
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c) // Changes 'a' array too.
	// Array inability to have their size determined at compile time led creators to introduce slice.
	d := []int{1, 2, 3} // A Slice ( no size passed during declaration )
	fmt.Println(d)
	fmt.Printf("Length: %v\n", len(d))
	fmt.Printf("Capacity: %v\n", cap(d)) // cap denotes Capacity
	e := d                               // By default points to address of 'd' array. No need to put '&' operation.
	e[2] = 30
	fmt.Println(e)
	fmt.Println(d)
	r := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // [...] Also valid since slices can be made from array/slice
	r1 := r[:]
	r2 := r[3:]  // Slice from 4th element to end
	r3 := r[:6]  // Slice first 6 elements.
	r4 := r[3:6] // Slice 4th, 5th & 6th elements. ( from index 3, upto but excluding index 6 )
	r[5] = 40    // Changes all the slices. ( All points to same address )
	fmt.Println(r)
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
	// Build in make() in Go lang.
	f := make([]int, 3, 100)
	fmt.Println(f)
	fmt.Printf("Length: %v\n", len(f))
	fmt.Printf("Capacity: %v\n", cap(f))
	// Changing capacity of slice after declaration using append()
	x := []int{}
	fmt.Println(x)
	fmt.Printf("Length: %v\n", len(x))
	fmt.Printf("Capacity: %v\n", cap(x))
	x = append(x, 200)
	x = append(x, 2, 3, 4, 5)
	// x = append(x, []int{2, 3, 4, 5})  Won't work since can't append slice of int to int. Check line 65 that x is int type.
	x = append(x, []int{2, 3, 4, 5}...) // Now works since (...) spread operator assigns individually rather all together. Slice concatenation.
	fmt.Println(x)
	fmt.Printf("Length: %v\n", len(x))
	fmt.Printf("Capacity: %v\n", cap(x))
	// apend is pushing elements in stack.
	k := x[1:] // One way to pop from beginning of slice.
	fmt.Println(k)
	k = x[:len(x)-1] // Wayt to pop from end of slice.
	fmt.Println(k)
	k = append(x[:4], x[5:]...) // Removes element from middle.
	fmt.Println(k)
	fmt.Println(x) // Unexpected changes in underlying array since all points to same location.
}
