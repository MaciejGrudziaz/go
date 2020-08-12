package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, java, python = 1, true, "NO!"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println("Sqrt of 2 is", math.Sqrt(2))
	fmt.Println("Random number:", rand.Intn(10))
	fmt.Println("2 + 2 =", add(2, 2))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	_1, _2 := split(90)
	fmt.Println("Split results are", _1, _2)

	i := 5
	fmt.Println(i, c, java, python)

	v := "test"
	fmt.Printf("Value is of type %T\n", v)

	const ConstVal = 123
	fmt.Println("const value", ConstVal)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
