package mylib

import (
	"fmt"
	"reflect"
)

func Variable() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x1 := 2                      //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x1)

	// Multiple Declaration
	var a, b, c, d int = 1, 3, 5, 7
	fmt.Println(a, b, c, d)

	var x, y = 6, "Hello" // x=6 & y=Hello
	fmt.Println(x, y)

	// Constant variable
	const PI = 3.14
	// Multiple Constants Declaration
	const (
		A int     = 1
		B float32 = 3.14
		C string  = "Hi!"
	)
	print(reflect.TypeOf(A).String())
}

/*
	int- stores integers (whole numbers), such as 123 or -123
	float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
	string - stores text, such as "Hello World". String values are surrounded by double quotes
	bool- stores values with two states: true or false
*/
