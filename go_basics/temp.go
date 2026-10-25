package main

import "fmt"

func main(){
	arr := []int{1,2,3}
	fmt.Println(arr, len(arr))

	//prac slice
	s1 := [8]int{1,2,3}
	fmt.Println(s1)

	s2 := make([]int, 5, 10)
	fmt.Println(s2)
}