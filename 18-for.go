package main

import "fmt"

// Go has only one looping construct, the for loop.
/*
The basic for loop has three components separated by semicolons:

the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration
The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.
*/
/*
Note: Unlike other languages like C, Java, or JavaScript
 there are no parentheses surrounding the three components of the for statement
 and the braces { } are always required.
*/

func forWithInit() int {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	return sum
}

// The init and post statements are optional.
// At that point you can drop the semicolons: C's while is spelled for in Go.
func forInitPostOptional() int {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	return sum
}

/*
	If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
	To exist from this loop condition, you can use break or return based on the condition that needed for the program.
*/
// func forForever() {
// 	for {
// 	}
// }

func main() {
	fmt.Println(forWithInit())
	fmt.Println(forInitPostOptional())
}
