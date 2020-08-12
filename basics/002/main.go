package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

// test function for square root
func Sqrt(x float64) float64 {
	const eps float64 = 1e-15

	prevZ := 0.0
	z := 1.0
	iter := 0

	for diff := math.Abs(z - prevZ); diff > eps; iter++ {
		prevZ = z
		z -= (z*z - x) / (2 * z)
		diff = math.Abs(z - prevZ)
	}
	fmt.Println("Iterations:", iter)

	return z
}

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	sum = 1

	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)

	// infinite loop
	// for {
	// 	fmt.Println("inifinite")
	// }

	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	v := 10.0
	fmt.Println("My fun:", Sqrt(v), "math.Sqrt:", math.Sqrt(v))

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	defer fmt.Println("world")
	fmt.Println("hello")

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
