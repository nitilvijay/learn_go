package main

import "fmt"

func main(){
	//initializing a map
	map1 := map[string]int{}
	map1["nitil"] = 20
	fmt.Println(map1)
	delete(map1, "nitil")
	fmt.Println(map1)

	map2 := make(map[string][]int)
	fmt.Println(map2)
	fmt.Println(map2["athe"])

}