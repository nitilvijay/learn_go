package main

import "fmt"

type Person struct {
	name    string
	age     int
	address struct {
		pin    int
		street string
	}
}

func main() {
	p := Person{
		name: "Nitil",
		age:  19,
		address: struct {
			pin    int
			street string
		}{
			pin:    637003,
			street: "west street",
		},
	}
	fmt.Println(p)
}
