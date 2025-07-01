package main

import "fmt"

// --- Types for type assertion examples ---
type Command interface{}
type Start struct{ ID int }
type Stop struct{ Reason string }

// Example 1: Type assertion (correct and error)
func typeAssertionBasic() {
	fmt.Println("\n--- Example 1: Type assertion (correct and error) ---")
	var i interface{} = "hello"
	str, ok := i.(string)
	fmt.Println("Type assertion to string:", str, ok)
	f, ok := i.(float64)
	fmt.Println("Type assertion to float64:", f, ok)
	// Uncomment to see panic:
	// _ = i.(int)
	fmt.Println("Uncomment code to see panic on failed assertion.")
}

// Example 2: Type switch in real-world logging
func typeSwitchLogging() {
	fmt.Println("\n--- Example 2: Type switch for logging (real-world) ---")
	log := func(val interface{}) {
		switch v := val.(type) {
		case string:
			fmt.Println("Log string:", v)
		case int:
			fmt.Println("Log int:", v)
		case error:
			fmt.Println("Log error:", v)
		default:
			fmt.Println("Log unknown type:", v)
		}
	}
	log("hello")
	log(42)
	log(fmt.Errorf("something went wrong"))
}

// Example 3: Type assertion anti-pattern (not checking ok)
func typeAssertionAntiPattern() {
	fmt.Println("\n--- Example 3: Type assertion anti-pattern (not checking ok) ---")
	// var i interface{} = 123
	// This will panic if uncommented:
	// s := i.(string)
	// fmt.Println(s)
	fmt.Println("Uncomment code to see panic when not checking ok.")
}

// Example 4: Type switch for dynamic dispatch (business case)
func typeSwitchBusinessDispatch() {
	fmt.Println("\n--- Example 4: Type switch for dynamic dispatch (business) ---")
	dispatch := func(cmd Command) {
		switch v := cmd.(type) {
		case Start:
			fmt.Println("Start command, ID:", v.ID)
		case Stop:
			fmt.Println("Stop command, Reason:", v.Reason)
		default:
			fmt.Println("Unknown command")
		}
	}
	dispatch(Start{ID: 1})
	dispatch(Stop{Reason: "done"})
	dispatch("invalid")
}

// TypeAssertionExamples runs all type assertion examples
func TypeAssertionExamples() {
	typeAssertionBasic()
	typeSwitchLogging()
	typeAssertionAntiPattern()
	typeSwitchBusinessDispatch()
}
