package main

import "fmt"

func arithmetic(num1, num2 int, op_func func(int, int) int) int{
	return op_func(num1, num2)
}

func add(num1, num2 int) int{
	return num1+num2
}

func mul(num1, num2 int) int{
	return num1*num2
}

func main(){
	var num1, num2 int
	fmt.Println("Enter the two numbers")
	fmt.Scanln(&num1, &num2)
	fmt.Println("Addition:", arithmetic(num1, num2, add))
	fmt.Println("Multiplication:", arithmetic(num1, num2, mul))
}