package main

import (
	"fmt"
)

const gs string = "Constant"

func main() {
	fmt.Println(gs)

	const num = 500000 //Implisit type declaration
	fmt.Println(num)

	const pi = 22 / 7
	fmt.Println(pi)
}
