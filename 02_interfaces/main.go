package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please select an example to run:")
		fmt.Println("1 - Interface Basics")
		fmt.Println("2 - Implicit Implementation & Polymorphism")
		fmt.Println("3 - Type Assertion & Type Switch")
		fmt.Println("4 - Interface Composition")
		fmt.Println("5 - Empty Interface & Any")
		fmt.Println("Usage: go run main.go <example_number>")
		return
	}

	example, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid example number")
		return
	}

	switch example {
	case 1:
		fmt.Println("=== Interface Basics ===")
		InterfaceBasicExamples()
	case 2:
		fmt.Println("=== Implicit Implementation & Polymorphism ===")
		ImplicitImplExamples()
	case 3:
		fmt.Println("=== Type Assertion & Type Switch ===")
		TypeAssertionExamples()
	case 4:
		fmt.Println("=== Interface Composition ===")
		InterfaceCompositionExamples()
	case 5:
		fmt.Println("=== Empty Interface & Any ===")
		EmptyInterfaceExamples()
	default:
		fmt.Println("Unknown example number")
	}
}
