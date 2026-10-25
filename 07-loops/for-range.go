package main

import "fmt"

func main()  {
	var name string

	fmt.Print("Digite seu nome: ")
	fmt.Scanf("%s", &name)

	fmt.Println("byte | string")
	for _, b := range name {
		// fmt.Println(i, name[i], string(name[i]))
		fmt.Println(b, string(b))
	}
}
