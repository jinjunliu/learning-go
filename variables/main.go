package main

import (
	"fmt"
)

// declare a constant
const aConstant int = 10

func main() {
	// print a constant
	fmt.Println(aConstant)
	// bool type
	var isTrue bool = true
	fmt.Println(isTrue)
	// string type
	var aString string = "Hello World"
	fmt.Println(aString)
	fmt.Printf("The variable is of type %T\n", aString)
	// implicit type: string
	var anotherString = "Hello World 2"
	fmt.Println(anotherString)
	// using semicolon
	myString := "Hello World 3"
	fmt.Println(myString)
	// int type
	var anInt int = 10
	fmt.Println(anInt)
	// int with default value
	var anotherInt int
	fmt.Println(anotherInt)
	// float type
	var aFloat float64 = 10.2
	fmt.Println(aFloat)
	// complex type
	var aComplex complex64 = 1 + 2i
	fmt.Println(aComplex)
}
