package main

import "fmt"

func main(){
	var _ int64 = 10
	num1 := 23.1
	num1 = 11
	fmt.Printf("This is the type %p\n", &num1)
}