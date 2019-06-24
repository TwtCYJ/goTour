package main

import (
	"fmt"
	"time"
)

func main() {
	go say("world")
	say("hello")

	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	d := make(chan int, 2)
	d <- 1
	d <- 2
	//d <- 3
	fmt.Println(<-d)
	fmt.Println(<-d)

	e := make(chan int, 10)
	go fibonacci(cap(e), e)
	for i := range e {
		fmt.Println(i)
	}

	f := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-f)
		}
		quit <- 0
	}()
	fibonacci1(f, quit)

	tick := time.Tick(100 * time.Microsecond)
	boom := time.After(500 * time.Microsecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom!")
			return
		default:
			fmt.Println("      .")
			time.Sleep(50 * time.Microsecond)
		}
	}
}

func fibonacci1(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}
