package main

import "fmt"

func main() {
	var a [100]int
	fatia := a[:10]

	fmt.Println(fatia)
	fmt.Println(len(fatia))
}
