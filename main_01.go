package main

import ( 
	"fmt"
)

func main() {
	// different ways to print

	fmt.Print("Hello", "World", 1, 2, 3, true, false, 1.2, 2.3, 3.4)
	fmt.Print("\n") // for new line
	fmt.Println("Hello", "World", 1, 2, 3, true, false, 1.2, 2.3, 3.4)
	fmt.Printf("Hello %s %d %d %d %t %t %f %f %.2f", "World", 1, 2, 3, true, false, 1.2, 2.3, 3.4) //%.2f for 2 decimal places
	fmt.Printf("\n") // for new line

	// check the type of a variable
	fmt.Printf("%T \n", 1)
	fmt.Printf("%T \n", 1.2)
	fmt.Printf("%T \n", "Hello")
	fmt.Printf("%T \n", true)

	// assign a variable
	var a int = 10
	fmt.Println(a)
	var b float64 = 10.2
	fmt.Println(b)
	var c string = "Hello"
	fmt.Println(c)
	var d bool = true
	fmt.Println(d)

	// assign a variable with using var
	e := 10
	fmt.Println(e)

	// assign a variable with using user input
	var f int
	fmt.Scan(&f)
	fmt.Println(f)
	
	var g string
	fmt.Scan(&g)
	fmt.Println(g)

	// assign a variable with using user input with multiple values
	var h, i int
	fmt.Scan(&h, &i)
	fmt.Println(h, i)

	// take a user input with message
	var j int
	fmt.Println("Enter a number")
	fmt.Scan(&j)
	fmt.Println(j)
}
