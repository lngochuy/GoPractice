package main

import (
	"bufio"
	"fmt"
	"os"
	"tma/lngochuy/gopractice/largenum"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input the first number: ")
	s1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read:", err)
		return
	}

	ln1, err := largenum.InitializeLargeNum(s1[:len(s1)-2])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Input the second number: ")
	s2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read:", err)
	}

	ln2, err := largenum.InitializeLargeNum(s2[:len(s2)-2])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Choose the operator (+-*/): ")
	operator, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid operator")
		return
	}

	switch operator[:len(operator)-2] {
	case "+":
		fmt.Printf("Result: %s + %s = %s", ln1, ln2, ln1.Add(ln2))
	case "-":
		fmt.Printf("Result: %s - %s = %s", ln1, ln2, ln1.Sub(ln2))
	case "*":
		fmt.Printf("Result: %s * %s = %s", ln1, ln2, ln1.Mul(ln2))
	case "/":
		fmt.Printf("Result: %s / %s = %s", ln1, ln2, ln1.Div(ln2, -10))
	default:
		fmt.Printf("Invalid operator %s", operator)
	}
}
