package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please select an example to run:")
		fmt.Println("1 - Goroutines Basic Examples")
		fmt.Println("2 - Channels Communication Examples")
		fmt.Println("3 - Select Multiplexing Examples")
		fmt.Println("4 - Context Management Examples")
		fmt.Println("5 - Sync Package Synchronization Primitives Examples")
		fmt.Println("6 - Atomic Operations Examples")
		fmt.Println("7 - Actor Model Examples")
		fmt.Println("8 - CSP Pattern Examples")
		fmt.Println("9 - Future/Promise Pattern Examples")
		fmt.Println("10 - Reactive Programming Examples")
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
		fmt.Println("=== Goroutines Basic Examples ===")
		GoroutinesExamples()
	case 2:
		fmt.Println("=== Channels Communication Examples ===")
		ChannelsExamples()
	case 3:
		fmt.Println("=== Select Multiplexing Examples ===")
		SelectExamples()
	case 4:
		fmt.Println("=== Context Management Examples ===")
		ContextExamples()
	case 5:
		fmt.Println("=== Sync Package Synchronization Primitives Examples ===")
		SyncExamples()
	case 6:
		fmt.Println("=== Atomic Operations Examples ===")
		AtomicExamples()
	case 7:
		fmt.Println("=== Actor Model Examples ===")
		ActorExamples()
	case 8:
		fmt.Println("=== CSP Pattern Examples ===")
		CSPExamples()
	case 9:
		fmt.Println("=== Future/Promise Pattern Examples ===")
		FutureExamples()
	case 10:
		fmt.Println("=== Reactive Programming Examples ===")
		ReactiveExamples()
	default:
		fmt.Println("Unknown example number")
	}
}
