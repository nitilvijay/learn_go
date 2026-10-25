package main

import "fmt"

func function1() func() string {
	return func() string {
		return "Hello World"
	}
}

func main() {
	f := function1() //f is a function that returns a function
	fmt.Println(f()) //Calling the function f which returns "Hello World"
}