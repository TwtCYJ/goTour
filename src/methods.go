package main

import (
	"fmt"
	"image"
	"io"
	"math"
	"net/http"
	"os"
	"strings"
	"time"
)

type Abser interface {
	Abs() float64
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type Vertex2 struct {
	x, y float64
}

type Person struct {
	Name string
	Age  int
}

type MyError struct {
	When time.Time
	What string
}

type HelloWeb struct {
}

func (hw HelloWeb) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello!")
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"system error",
	}
}

func (p Person) String() string {
	return fmt.Sprintf("%v(%v years)", p.Name, p.Age)
}

func (v *Vertex2) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vertex2) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := &Vertex2{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(5)
	fmt.Println(v, v.Abs())

	v1 := Vertex2{3, 4}
	var a Abser
	a = f
	a = &v1
	//a = v1
	fmt.Println(a.Abs())

	var w Writer
	w = os.Stdout
	fmt.Fprintf(w, "hello, writer\n")

	y := Person{"arthur dent", 42}
	z := Person{"zaphod beeblebrox", 9001}
	fmt.Println(y, z)

	if err := run(); err != nil {
		fmt.Println(err)
	}

	r := strings.NewReader("hello, reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	//var hw HelloWeb
	//err := http.ListenAndServe("localhost:4000", hw)
	//if err != nil {
	//	log.Fatal(err)
	//}

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
