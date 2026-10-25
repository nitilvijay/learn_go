package main

import "fmt"

func main(){
	arr := [5]int{1,2,3}
	fmt.Println(arr, len(arr))

	arr1 := [...]int{1,2,3,4}
	fmt.Println(len(arr1))

	arr2 := [5]string{1:"happy", 4:"world"}
	fmt.Println(arr2)
}