package main

import "fmt"

// --- Types for empty interface examples ---
type AnySlice []interface{}

// Example 1: Empty interface for generic containers
func emptyInterfaceGenericContainer() {
	fmt.Println("\n--- Example 1: Empty interface for generic container ---")
	var s AnySlice = []interface{}{123, "abc", []int{1, 2, 3}}
	for i, v := range s {
		fmt.Printf("s[%d] = %v (type %T)\n", i, v, v)
	}
}

// Example 2: Type switch for dynamic dispatch (real-world)
func emptyInterfaceTypeSwitch() {
	fmt.Println("\n--- Example 2: Type switch for dynamic dispatch ---")
	process := func(val interface{}) {
		switch v := val.(type) {
		case int:
			fmt.Println("Processing int:", v)
		case string:
			fmt.Println("Processing string:", v)
		case []int:
			fmt.Println("Processing []int:", v)
		default:
			fmt.Println("Unknown type:", v)
		}
	}
	process(123)
	process("abc")
	process([]int{4, 5, 6})
}

// Example 3: Real-world - JSON decode to map[string]interface{}
func emptyInterfaceJSONDecode() {
	fmt.Println("\n--- Example 3: JSON decode to map[string]interface{} (simulated) ---")
	// Simulate a JSON decode result
	var data map[string]interface{} = map[string]interface{}{
		"name": "Alice",
		"age":  30,
		"tags": []interface{}{1, "go", true},
	}
	for k, v := range data {
		fmt.Printf("key=%s, value=%v (type %T)\n", k, v, v)
	}
}

// Example 4: Error - type safety pitfall with any
func emptyInterfaceTypeSafetyPitfall() {
	fmt.Println("\n--- Example 4: Type safety pitfall with any ---")
	// var x interface{} = 123
	// s := x.(string) // Uncomment to see panic: interface conversion: int is not string
	fmt.Println("Uncomment code to see panic when type assertion fails.")
}

// Example 5: Real-world - Reflection with empty interface
func emptyInterfaceReflection() {
	fmt.Println("\n--- Example 5: Reflection with empty interface (simulated) ---")
	// In real code, use reflect.TypeOf/ValueOf to inspect interface{}
	var x interface{} = []int{1, 2, 3}
	fmt.Printf("Type: %T, Value: %v\n", x, x)
}

// EmptyInterfaceExamples runs all empty interface examples
func EmptyInterfaceExamples() {
	emptyInterfaceGenericContainer()
	emptyInterfaceTypeSwitch()
	emptyInterfaceJSONDecode()
	emptyInterfaceTypeSafetyPitfall()
	emptyInterfaceReflection()
}
