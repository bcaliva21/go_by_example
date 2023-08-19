package main

import "fmt"

type base struct {
	num int
	hello string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}


func main () {
	 co := container {
		 base: base{
			 num: 1,
			 hello: "hola",
		 },
		 str: "some name",
	 }

	 fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	 fmt.Printf("%v\n", co.hello)

	 fmt.Println("also num:", co.base.num)

	 fmt.Println("describe: ", co.describe())


	 type describer interface {
		 describe() string
	 }

	 var d describer = co
	 fmt.Println("describer:", d.describe())
}

