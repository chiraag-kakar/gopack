package main

import (
	"fmt"
	operation "operation/Operation"
)

func main() {
	//Initalise the Init function with value of A,B
	e := operation.Init{
		9, 2,
	}
	fmt.Println(e.GetSquareSum())
}
