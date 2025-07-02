package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please select an example to run:")
		fmt.Println("1 - Type System Reflection")
		fmt.Println("2 - Struct Reflection & Tags")
		fmt.Println("3 - Function & Method Reflection")
		fmt.Println("4 - Interface Reflection")
		fmt.Println("5 - Value Operations & Creation")
		fmt.Println("6 - Common Mistakes & Error Handling")
		fmt.Println("7 - Reflection Design Patterns")
		fmt.Println("8 - Advanced Topics")
		fmt.Println("Usage: go run . <example_number>")
		return
	}

	example, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid example number")
		return
	}

	switch example {
	case 1:
		fmt.Println("=== Type System Reflection ===")
		TypeReflectionExamples()
	case 2:
		fmt.Println("=== Struct Reflection & Tags ===")
		StructReflectionExamples()
	case 3:
		fmt.Println("=== Function & Method Reflection ===")
		FunctionMethodReflectionExamples()
	case 4:
		fmt.Println("=== Interface Reflection ===")
		InterfaceReflectionExamples()
	case 5:
		fmt.Println("=== Value Operations & Creation ===")
		ValueOperationsExamples()
	case 6:
		fmt.Println("=== Common Mistakes & Error Handling ===")
		CommonMistakesExamples()
	case 7:
		fmt.Println("=== Reflection Design Patterns ===")
		ReflectionPatternsExamples()
	case 8:
		fmt.Println("=== Advanced Topics ===")
		AdvancedTopicsExamples()
	default:
		fmt.Println("Unknown example number")
	}
}
