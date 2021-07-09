package main

import "fmt"

func calc(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	fmt.Println("Go test tutorial")
	result := calc(2)
	fmt.Println(result)
}
