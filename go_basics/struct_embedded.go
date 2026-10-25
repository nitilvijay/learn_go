package main

import "fmt"

type Person struct{
	name string
	age int
	marks
}
type marks struct{
	eng int
	math int
	sci int
}

func main(){
	//Initializing the embedded struct
	p1 := Person{
		name: "nitil",
		age: 19,
		marks: marks{
			eng: 98,
			math: 93,
			sci: 97,
		},
	}
	fmt.Printf("p1: %v\n", p1.eng)
}