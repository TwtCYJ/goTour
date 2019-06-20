package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

var c, python, java bool
var i, j = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	big   = 1 << 100
	small = big >> 99
)

func main() {
	fmt.Println("hello, world")
	fmt.Println("number is", rand.Intn(10))
	fmt.Println("next", math.Nextafter(2, 3))
	fmt.Println(math.Pi)
	fmt.Println(add(2, 3))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	var i int
	fmt.Println(i, c, python, java)
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	k := 3
	fmt.Println(i, j, k, c, python, java)
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	var x, y int = 3, 4
	var t float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(t)
	fmt.Println(x, y, z)

	v := 0.1
	fmt.Printf("v is of type %T\n", v)

	fmt.Println(needInt(small))
	fmt.Println(neddFloat(small))
	fmt.Println(neddFloat(big))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum1 := 1
	for ; sum1 < 1000; {
		sum1 += sum1
	}
	fmt.Println(sum1)

	sum2 := 2
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20))

	fmt.Print("go runs on  ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("os x")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s", os)
	}

	fmt.Println("when`s saturday")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today.")
	case today + 1:
		fmt.Println("tomorrow.")
	case today + 2:
		fmt.Println("in two days.")
	default:
		fmt.Println("too far away.")
	}

	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("morinning")
	case time.Hour() < 17:
		fmt.Println("afternoon")
	default:
		fmt.Println("evening")
	}

	defer fmt.Println("wait")
	fmt.Println("go now")

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")


}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func needInt(x int) int {
	return x*10 + 1
}

func neddFloat(x float64) float64 {
	return x * 0.1
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func swap(x string, y string) (string, string) {
	return y, x
}

func add(x int, y int) int {
	return x + y
}

func add1(x, y int) int {
	return x + y
}
