package main

import "fmt"

type Operands struct {
	n1 int
	n2 int
}

func main() {
	var operands Operands = Operands{}
	var operators = map[int]func(int, int) int{
		1: add,
		2: subtract,
		3: multiply,
		4: divide,
	}
	for {
		var choice int = getUserChoice()
		if choice < 0 || choice >= 5 {
			break
		}
		getOperands(&operands)
		result := operators[choice](operands.n1, operands.n2)
		fmt.Printf("result : %d\n", result)
	}
}

func getOperands(operands *Operands) {
	fmt.Println("Enter the operands :")
	fmt.Scanf("%d %d", &operands.n1, &operands.n2)
}

func getUserChoice() int {
	fmt.Println("Select the operation :")
	fmt.Println("1 : Add")
	fmt.Println("2 : Subtract")
	fmt.Println("3 : Multiply")
	fmt.Println("4 : Divide")
	fmt.Println("5 : Exit")
	var choice int
	fmt.Scanf("%d", &choice)
	return choice
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
