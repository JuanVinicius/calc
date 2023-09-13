package main

import (
	"fmt"
)

func mult(x int, y int) int {
	return x * y
}

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func div(fx float64, fy float64) float64 {
	return fx / fy
}

//multi := x * y
//soma := x + y
//sub := x - y
//div := x / y

func main() {
	x := 7
	y := 7
	var fx float64 = float64(x)
	var fy float64 = float64(y)

	var calc string = "sub"

	switch calc {
	case "mult":
		fmt.Printf("Multiplicando: \t%v\n", mult(x, y))

	case "add":
		fmt.Printf("Somando: \t%v\n", add(x, y))

	case "sub":
		fmt.Printf("Subtraindo: \t%v\n", sub(x, y))

	case "div":
		fmt.Printf("Dividindo: \t%v\n", div(fx, fy))
	}

}
