package main

import (
	"fmt"
	"os"
)

func readFile() (string, error) {
	handle, err := os.Open("pointer_deref.c")
	defer handle.Close() // Will execute everytime before the function returns
	//Useful when the function has multiple return statements and you want to ensure that the file is closed before returning
	
	if err != nil {
		return "", err
	}
	return "File read successfully", nil
}

func main() {
	fmt.Println(readFile())
}