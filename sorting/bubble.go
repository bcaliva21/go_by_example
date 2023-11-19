package main

import "fmt"

type Bubble struct {
	List []int
}

func (b *Bubble) Swap (i1, i2 int) {
	b.List[i1], b.List[i2] = b.List[i2], b.List[i1]
}

func (b *Bubble) Sort () {
	for i := len(b.List)-1; i > -1; i-- {
		for j := 0; j < i; j++ {
			if b.List[j] > b.List[j+1] {
				b.Swap(j,j+1)
			}
		}
	}
}

func main() {

	test := &Bubble{List: []int{10,9,8,7,6,5,4,3,2,1}}

	fmt.Println(test)
	
	test.Sort()

	fmt.Println(test)
}

