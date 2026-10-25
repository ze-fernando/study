package main

import "fmt"

func main()  {
	fmt.Println("String em uma linha")
	
	fmt.Println(`String
em 
multiplas
linha`)

	fmt.Println("Zioncodes tem ", len("zioncodes"), "letras")
	
	fmt.Println(string("Zioncodes"[0]))
	fmt.Println("Zioncodes"[0])
}
