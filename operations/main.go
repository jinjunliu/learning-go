package main

import (
	"fmt"
	"math"
)

func main() {
	// arithmetic operators
	fmt.Println("Arithmetic operators")
	var a int = 10
	var b int = 4
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	var b2 float64 = 4
	fmt.Println(float64(a) / b2) // type conversion, not automatic
	// relational operators
	fmt.Println("Relational operators")
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
	// logical operators
	fmt.Println("Logical operators")
	var c bool = true
	var d bool = false
	fmt.Println(c && d)
	fmt.Println(c || d)
	fmt.Println(!c)
	// bitwise operators
	fmt.Println("Bitwise operators")
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
	// assignment operators
	fmt.Println("Assignment operators")
	var e int = 10
	e += 10
	fmt.Println(e)
	e -= 10
	fmt.Println(e)
	e *= 10
	fmt.Println(e)
	e /= 10
	fmt.Println(e)
	e %= 10
	fmt.Println(e)
	e <<= 2
	fmt.Println(e)
	e >>= 2
	fmt.Println(e)
	e &= 2
	fmt.Println(e)
	e |= 2
	fmt.Println(e)
	e ^= 2
	fmt.Println(e)

	// math package
	fmt.Println("Math package")
	var f float64 = 3.14
	fmt.Println(math.Floor(f))

	i1, i2, i3 := 10, 20, 30
	intSum := i1 + i2 + i3
	fmt.Println(intSum)

	f1, f2, f3 := 10.5, 20.1, 30.3
	floatSum := f1 + f2 + f3
	fmt.Println(floatSum)
	fmt.Println(math.Round(floatSum))

	circRadius := 15.5
	circCircumference := circRadius * math.Pi * 2
	fmt.Printf("Circumference: %.2f\n", circCircumference)

}
