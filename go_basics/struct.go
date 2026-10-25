package main

import "fmt"

type Car struct{
	auto bool
	wheel int
	name string
}

type Owns struct{
	name string
	car Car
	age int
}

func main(){
	car1 := Car{}
	car1.auto = true
	car1.wheel = 4
	car1.name = "Tiguan"

	fmt.Printf("car1: %v\n", car1)

	var owner Owns
	fmt.Printf("owner.car: %v\n", owner)

	//initializing a nested struct
	owner1 := Owns{
		name: "nitil",
		age: 19,
		car: Car{
			auto: true,
			wheel: 4,
			name: "honda",
		},
	}
	fmt.Printf("owner1: %v\n", owner1)

	owner2 := Owns{
		name: "vijay",
		age: 48,
		car: car1,
	}
	fmt.Printf("owner2: %v\n", owner2)
}
