package main

import "fmt"

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func greet(name string) func() {
	return func() {
		fmt.Printf("hello %s\n", name)
		return
	}
}


func main() {

	firstInt := counter()

	fmt.Println(firstInt())
	fmt.Println(firstInt())
	fmt.Println(firstInt())

	secondInt := counter()
	fmt.Println(secondInt())

	bob := greet("bob")
	bob()

	sally := greet("sally")
	sally()
	sally()
	sally()

}


