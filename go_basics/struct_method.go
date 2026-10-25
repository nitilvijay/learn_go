package main

import "fmt"

type Shape struct{
	length int
	breadth int
}
func (s Shape) area()(a int){
	return s.breadth*s.length
}

func main(){
	s1 := Shape{
		length: 10,
		breadth: 20,
	}

	fmt.Printf("s1: %v\n", s1)
	
	fmt.Println(s1.area())
}