package referenceType

import "fmt"

type Vertex struct {
	Lat, Lon float64
}

//var o map[string]Vertex

func DoMap() {
	m := make(map[string]int)
	fmt.Println(m)
	m["A"] = 1
	m["B"] = 2
	m["C"] = 3
	fmt.Println(m)
	v := m["B"]
	fmt.Println(v)
	fmt.Println(len(m))
	delete(m, "C")
	fmt.Println(m["C"])
	_, e := m["C"]
	fmt.Println(e)

	n := map[string]int{"harish": 1, "harini": 2, "srinikethan": 3}
	fmt.Println(n["srinikethan"])

	o := make(map[string]Vertex)
	o["srirangam"] = Vertex{10.87, 78.68}
	o["chidambaram"] = Vertex{11.4, 79.7}
	fmt.Println(o)
}
