package main

import "fmt"

const DEFAULT_STRING = "Hello, World"
func Hello(name string) string {
	if name != "" {
		return "Hello, " + name
	}
	return DEFAULT_STRING
}

func main() {
	me := "Bradley"
	fmt.Println(Hello(me))
}
