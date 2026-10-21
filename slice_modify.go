package main

import "fmt"

func main(){
	s1 := make([]int, 5)
	fmt.Println(cap(s1))

	s1 = append(s1, 6)
	fmt.Println(cap(s1))

	s2 := []int{1,2,3}
	s3 := append(s1,s2...)

	fmt.Println(s3)


}