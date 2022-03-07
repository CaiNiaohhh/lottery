package main

import "fmt"

func main() {
	Map := make(map[byte]int)

	Map['a'] = 1
	k, ok := Map['b']
	fmt.Println("k:", k, "ok:", ok)

}


