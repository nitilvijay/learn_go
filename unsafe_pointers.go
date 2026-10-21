package main

import "fmt"
import "unsafe"

func main(){
	var p unsafe.Pointer = unsafe.Pointer(uintptr(0x2dc4877ac008))
	fmt.Print(*p)
}