package main

import "fmt"

func main(){
	print_range(2, 87)
	fmt.Println(operation(add, 2, 3))
}

func print_range(start, end int) (sum int){ //this is a named return value, so we don't need to declare it in the function body
	for i:=start; i<end+1; i++{
		fmt.Print(i, " ")
		sum+=i
	}

	return
}

func add(num1 int, num2 int) int{
	return num1+num2
}

func operation(op func(int, int) int, x int, y int) (result int){
	result  = op(x, y)
	return
}
