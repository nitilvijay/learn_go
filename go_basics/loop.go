package main

import "fmt"

func main() {
	val := true

	if val == true {
		fmt.Println("If condition worked")
	} else {
		fmt.Println("Else condition")
	}

	//While using for
	var i int
	for i<10{
		fmt.Println(i)
		i++
	}

	//For loop using range
	a := [...]int{3,2,3,4,5}
	for idx, val := range a{   //The val var is in diff scope
		fmt.Println(idx, val)
	}
}
	