package main

import "fmt"

func main()  {
	var idade int

	fmt.Print("Qual sua idade? ")
	fmt.Scanf("%d", &idade)

	if idade > 21 {
		fmt.Println("Adulto")

	} else if idade > 18 && idade < 21 {

		fmt.Println("Jovem")
	} else if idade > 12 && idade < 16 {
		fmt.Println("Adolescente")
	} else{
		fmt.Println("Criança")
	}
	
}
