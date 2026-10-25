package main

import "fmt"

// Task: Given a list of names, create a nested dictionary where the first key is each name’s 
// first letter, the second key is the name, and the value is how many times that name appears.

func getNamesCount(name []string) map[rune]map[string]int {
	map1 := make(map[rune]map[string]int)
	for _, n:= range name{
		firstLetter := n[0]

		if map1[rune(firstLetter)] == nil{
			map1[rune(firstLetter)] = make(map[string]int)
		}
		map1[rune(firstLetter)][n]++
	}
	fmt.Println(map1)
	return map1
}

func main(){
	names := []string{"Alice", "Bob", "Charlie", "Alice", "Bob", "Alic"}
	result := getNamesCount(names)
	fmt.Println(result)
}