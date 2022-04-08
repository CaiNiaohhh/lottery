package main

import "fmt"

type test interface {
	Add(x, y int) int
}

type a struct {
}

type b struct {
}

func (c a) Add(x, y int) int {
	return x + y
}

func (c b) Add(x, y int) float64 {
	return float64(x + y)
}


func main() {
	var x test
	x = a{}
	k := x.Add(1, 2)
	fmt.Println(k)

}