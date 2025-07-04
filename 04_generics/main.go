package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	example, err := strconv.Atoi(os.Args[1])
	if err != nil || example < 1 || example > 8 {
		fmt.Printf("Invalid example number: %s\n", os.Args[1])
		printHelp()
		return
	}

	fmt.Printf("🚀 Running Go Generics Example %d\n", example)
	fmt.Println(strings.Repeat("=", 50))

	switch example {
	case 1:
		fmt.Println("📚 Basic Generics - Functions, Type Parameters, Type Inference")
		runBasicGenericsExample()
	case 2:
		fmt.Println("🔒 Constraints - Built-in, Custom, Type Sets")
		runConstraintsExample()
	case 3:
		fmt.Println("🏗️ Generic Types - Structs, Interfaces, Methods")
		runGenericTypesExample()
	case 4:
		fmt.Println("🔧 Advanced Constraints - Type Sets, Unions, Approximations")
		runAdvancedConstraintsExample()
	case 5:
		fmt.Println("📦 Generic Containers - Stack, Queue, Map, Tree")
		runGenericContainersExample()
	case 6:
		fmt.Println("⚡ Generic Algorithms - Sort, Search, Transform, Aggregate")
		runGenericAlgorithmsExample()
	case 7:
		fmt.Println("🎨 Design Patterns - Factory, Builder, Decorator with Generics")
		runDesignPatternsExample()
	case 8:
		fmt.Println("✨ Best Practices - Performance, Pitfalls, Code Organization")
		runBestPracticesExample()
	}
}

func printHelp() {
	fmt.Println("🧬 Go Generics Programming Examples")
	fmt.Println("====================================")
	fmt.Println()
	fmt.Println("📖 Available Examples:")
	fmt.Println("  1 - Basic Generics (Functions, Type Parameters, Type Inference)")
	fmt.Println("  2 - Constraints (Built-in, Custom, Type Sets)")
	fmt.Println("  3 - Generic Types (Structs, Interfaces, Methods)")
	fmt.Println("  4 - Advanced Constraints (Type Sets, Unions, Approximations)")
	fmt.Println("  5 - Generic Containers (Stack, Queue, Map, Tree)")
	fmt.Println("  6 - Generic Algorithms (Sort, Search, Transform, Aggregate)")
	fmt.Println("  7 - Design Patterns (Factory, Builder, Decorator)")
	fmt.Println("  8 - Best Practices (Performance, Pitfalls, Organization)")
	fmt.Println()
	fmt.Println("🚀 Usage:")
	fmt.Println("  go run . <example_number>")
	fmt.Println()
	fmt.Println("💡 Examples:")
	fmt.Println("  go run . 1    # Basic generics")
	fmt.Println("  go run . 5    # Generic containers")
	fmt.Println("  go run . 8    # Best practices")
	fmt.Println()
	fmt.Println("📋 Learning Path:")
	fmt.Println("  Basic (1-3) → Intermediate (4-6) → Advanced (7-8)")
}

// Example function placeholders (to be implemented in separate files)
// runBasicGenericsExample is implemented in 01_basic_generics.go

// runConstraintsExample is implemented in 02_constraints.go

// runGenericTypesExample is implemented in 03_generic_types.go

// runAdvancedConstraintsExample is implemented in 04_advanced_constraints.go

// runGenericContainersExample is implemented in 05_generic_containers.go

// runGenericAlgorithmsExample is implemented in 06_generic_algorithms.go

// runDesignPatternsExample is implemented in 07_design_patterns.go

// runBestPracticesExample is implemented in 08_best_practices.go
