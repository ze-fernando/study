package main

import "fmt"

func main()  {
	var status int

	fmt.Print("Digite o status (1 - 4) ")
	fmt.Scanf("%d", &status)

	// 1 - Backlog
	// 2 - InProgress
	// 3 - Completed
	// 4 - Blocked

	switch status {
	case 1:
		fmt.Println("Backlog")
	case 2:
		fmt.Println("In Progress")
	case 3:
		fmt.Println("Completed")
	case 4:
		fmt.Println("Blocked")
	default:
		fmt.Println("Opção inválida")
	}
}
