package main

import "fmt"

func main(){
	// []int creates a slice; [...]int creates an array whose size is inferred from the literal.
	//Different ways of creating slice
	s1 := []int{1,2,3}
	fmt.Println(s1)

	a1 := [...]int{1,2,3,4,5}
	s2 := a1[0:3]
	fmt.Println(s2)

	s3 := make([]int, 5, 10)
	fmt.Println(cap(s3), len(s3))
	fmt.Println(s3)
	
}