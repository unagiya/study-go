package main

import "fmt"

//go:generate stringer -type=CarType
const (
	Sedan CarType = iota + 1
	Hatchback
	MPV
	SUV
	Crossover
	Coupe
	Convertible
)

type CarType int

func main() {
	var t CarType
	t = Convertible
	fmt.Printf("%sã§ãã\n", t)

}
