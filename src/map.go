package main

import "fmt"

type Vertex1 struct {
	Lat, Long float64
}

var m map[string]Vertex1

var m1 = map[string]Vertex1{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	m = make(map[string]Vertex1)
	m["bell labs"] = Vertex1{40.68433, -74.3996}
	fmt.Println(m["bell labs"])

	fmt.Println(m1)

	m2 := make(map[string]int)
	m2["1"] = 1
	fmt.Println("m2[\"1\"]=", m2["1"])
	m2["1"] = 2
	fmt.Println("m2[\"1\"]=", m2["1"])

	delete(m2, "1")
	fmt.Println("m2 = ", m2["1"])

	v, ok := m2["1"]
	fmt.Println("m2[\"1\"]= ", v, "present?", ok)

}
