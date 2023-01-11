package main

import "fmt"

// A var statement can be at package or function level. We see both in this example.
var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
