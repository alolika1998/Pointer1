package main

import "fmt"

func main() {
	size := new(int)
	fmt.Printf("Size value: %v, Type is %T, Address is %a\n", *size, size, size)
	*size = 85
	fmt.Println("New Size: ", *size)
}
