package main

import "fmt"

func main() {
	var a [3]string
	a[0] = "chen"
	a[1] = "ying"
	a[2] = "jie"

	fmt.Println(a[0], a[1], a[2])
	fmt.Println(a)

	p := []int{2, 3, 4, 5, 6, 7}
	fmt.Println("p == ", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	fmt.Println("p[1:4] ==", p[1:4])
	fmt.Println("p[:3]", p[:3])
	fmt.Println("p[4:]", p[4:])

	v1 := make([]int, 5)
	printSlice("v1", v1)
	v2 := make([]int, 0, 5)
	printSlice("v2", v2)
	v3 := v2[:2]
	printSlice("v3", v3)
	v4 := v3[2:5]
	printSlice("v4", v4)

	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

	var y []int
	printSlice("y", y)

	y = append(y, 0)
	printSlice("y", y)
	y = append(y, 1)
	printSlice("y", y)
	y = append(y, 2, 3, 4)
	printSlice("y", y)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow1 := make([]int, 10)
	for i := range pow1 {
		pow1[i] = 1 << uint(i)
	}

	for _, value := range pow1 {
		fmt.Printf("%d\n", value)
	}

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
