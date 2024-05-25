package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20*10
func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	var sl []int

	sl = append(sl, 1)

	fmt.Println(sl)
}
