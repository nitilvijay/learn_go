package main

import "fmt"

type Car struct{
	auto bool
	wheel int
	name string
}

type Bike struct{
	auto bool
	wheel int
	name string
}

func (c Car) display() string{
	return fmt.Sprintf("Car: %v, Auto: %v, Wheel: %v", c.name, c.auto, c.wheel)
}

// func (b Bike) display() string{
// 	return fmt.Sprintf("Bike: %v, Auto: %v, Wheel: %v", b.name, b.auto, b.wheel)
// }

type vehicle interface {
	display() string
}

func show(v vehicle){
	fmt.Println(v.display())
}

func main(){
	var v vehicle = Car{
		auto: true,
		wheel: 4,
		name: "Tiguan",
	}
	show(v) // Accessing the display method through interface is possible
	// fmt.Println(v.wheel) // Accessing the wheel field through interface is not possible
	// Type assertion
	c := v.(Car) // Type assertion to access the fields of the struct
	fmt.Println(c.wheel) // Accessing the wheel field through type assertion is possible

	//standard way to check if the type assertion is successful
	c1, ok := v.(Car)
	if ok {
		fmt.Println(c1.wheel)
	} else {
		fmt.Println("Type assertion failed")
	}
}