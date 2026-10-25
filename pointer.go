package main

import "fmt"

func main(){
	a := "hello"
	b := &a
	fmt.Println(a)  // Output: hello
	fmt.Println(b)  // Output: memory address of a
	fmt.Println(*b) // Output: hello

	*b = "world"
	fmt.Println(b+1)
}