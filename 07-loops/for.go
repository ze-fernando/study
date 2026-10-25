package main

import "fmt"

func main()  {
	var qtd int

	fmt.Print("Digite quantas vezes quer receber elogio: ")
	fmt.Scanf("%d", &qtd)

	for i := 0; i<qtd; i++{
		fmt.Println("Tu é foda papai")
	}
}
