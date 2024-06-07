package mylib

import (
	"fmt"
	"reflect"
)

func Array() {
	// var array_name = [length]datatype{values}
	var arr1 = [3]int{1, 2, 3}      // Fixed size
	arr2 := [...]int{4, 5, 6, 7, 8} //Inferred size

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(reflect.TypeOf(arr1).String())

	arr3 := [5]int{1: 10, 2: 40} // index defined for value. Output: [0 10 40 0 0]
	fmt.Println(arr3)

	// Size
	arr4 := [4]string{"Volvo", "BMW", "Ford"}
	fmt.Println(len(arr4)) // 4

	// Slice
	var slice []int = []int{1, 2, 3} //not fixed size
	print(slice)

}
