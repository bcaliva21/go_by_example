package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	key := "k1"

	v1 := m[key]
	fmt.Println("v1: ", v1)

	v3 := m["k3"]
	fmt.Println("v3: ", v3)

	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	_, k2InMapBool := m["k2"]

	_, keyInMapBool := m[key]

	fmt.Println("k2Bool: ", k2InMapBool)
	fmt.Println("keyInMapBool: ", keyInMapBool)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)
}


