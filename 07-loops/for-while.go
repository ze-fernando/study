package main

import "fmt"

func main()  {
	var num int
	i := 0

	fmt.Print("Digite em qual número tu quer parar: ")
	fmt.Scanf("%d", &num)

	for i <= num {
		fmt.Println(i)
		i++
	}
}
