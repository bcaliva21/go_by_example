package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func apiCall() (string, string) {
	return "SUCCESS", "ERROR"
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	res, err := apiCall()
	fmt.Println(res)
	fmt.Println(err)

	_, c := vals()
	fmt.Println(c)
}

