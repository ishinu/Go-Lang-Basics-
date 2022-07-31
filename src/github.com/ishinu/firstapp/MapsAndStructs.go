package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
}

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   // Ember Animal struct.
	SpeedKPH float32
	CanFly   bool
}

type Library struct {
	BookName   string `required max:"100"`
	BookAuthor string
}

func main() {
	statePopulations := map[string]int{ // Syntax to declare map.
		"Texas":    1213123,
		"Florida":  123132,
		"New York": 123123,
		"Ohio":     1231321,
	}
	// Keys and values pairs. As mentioned above, key must be string and values must be int.
	// Slices can't be key but arrays can.
	// m := map[[]int]string{}  Error since slices can't be key.
	m := map[[3]int]string{} // Arrays can't be.
	fmt.Println(statePopulations, m)
	a := make(map[string]int) // Using make() to create map.
	a = map[string]int{
		"One": 1,
		"Two": 2,
	}
	fmt.Println(a)
	fmt.Println(a["One"])
	a["Three"] = 3 // Add a new key|value pair.
	fmt.Println(a)
	delete(a, "Two") // Delete an existing key|value pair.
	fmt.Println(a)
	pop, ok := a["one"] // Intentionally misspelled 'One'
	fmt.Println(pop, ok)
	pop1, ok := a["One"] // To check the presence of a key, we use 'ok'(just another variable not keyword )  which tells is it present in map or not.
	fmt.Println(pop1, ok)
	_, ok = a["Three"] // By using '_' symbol, we can leave key out and just check the presence.
	fmt.Println(ok)
	// Just like slices, manipulating one map which is just addressing another map will change the underlying map too.
	b := a
	delete(b, "One")
	fmt.Println(b)
	fmt.Println(a) // Deleted from underlying map too.

	aDoctor := Doctor{
		Number:    3,
		ActorName: "Rechil Nakstu",
		Companions: []string{
			"Aerin",
			"Ieros",
			"Kespa",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.ActorName)
	fmt.Println(aDoctor.Companions[1])
	// Using positional syntax, we can instantiate our struct.

	bDoctor := Doctor{
		1,
		"Yiasa Iisnt",
		// In case another value gets added in struct called 'episodes []string'
		// We need to add '[]string{}' to compensate here in initialization.
		// Need to make sure ordering remains unchanged.
		[]string{
			"Iuiada",
		},
	}
	// Suggested to declare field names explicitely instead of positional syntax.
	fmt.Println(bDoctor)
	// Instead of setting a type struct , we can do in main()
	cDoctor := struct{ name string }{name: "Inncaa"} // Anonymous struct
	fmt.Println(cDoctor)
	dDoctor := cDoctor
	dDoctor.name = "Aneana" // Changing name won't change original struct values. Although we can use & operator to pass address.
	fmt.Println(dDoctor)
	fmt.Println(cDoctor) // Maintains the original value since structs passes copy not address.
	eDoctor := &cDoctor  // eDoctor is a pointer to the struct
	eDoctor.name = "Baslsa"
	fmt.Println(eDoctor)
	fmt.Println(cDoctor)
	// Go lang doesn't support inheritance but does supports composition.
	o := Bird{} // Creating instance of struct and assigning value.
	o.Name = "Nassaa"
	o.Origin = "Ohio"
	o.SpeedKPH = 20
	o.CanFly = true
	fmt.Println(o)
	u := Bird{
		Animal:   Animal{Name: "Naaaebe", Origin: "New York"},
		SpeedKPH: 42,
		CanFly:   true,
	}
	fmt.Println(u)

	// Validation for which we imported reflect package.
	z := reflect.TypeOf(Library{})
	field, _ := z.FieldByName("BookName")
	fmt.Println(field.Tag)
}
