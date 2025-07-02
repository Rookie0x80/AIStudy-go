package main

import (
	"context"
	"fmt"
	"reflect"
)

// FunctionMethodReflectionExamples demonstrates function and method reflection
func FunctionMethodReflectionExamples() {
	fmt.Println("\n=== Function & Method Reflection Examples ===")

	// Example 1: Function type reflection
	functionTypeReflection()

	// Example 2: Method reflection and invocation
	methodReflection()

	// Example 3: Dynamic function calls with parameters
	dynamicFunctionCalls()

	// Example 4: Function signature validation
	functionSignatureValidation()

	// Example 5: Handler registration pattern
	handlerRegistration()

	// Example 6: Method discovery and documentation
	methodDiscovery()
}

// Example 1: Function type reflection
func functionTypeReflection() {
	fmt.Println("\n--- Example 1: Function Type Reflection ---")

	// Various function types to analyze
	functions := map[string]interface{}{
		"Simple function":   simpleAdd,
		"Multiple returns":  divideWithError,
		"Variadic function": sum,
		"Handler function":  sampleHandler,
		"Method function":   (SimpleCalculator{}).Add,
	}

	for name, fn := range functions {
		fmt.Printf("\nAnalyzing: %s\n", name)
		analyzeFunction(fn)
	}
}

// Example 2: Method reflection and invocation
func methodReflection() {
	fmt.Println("\n--- Example 2: Method Reflection ---")

	calc := SimpleCalculator{Name: "TestCalc"}
	v := reflect.ValueOf(calc)
	t := reflect.TypeOf(calc)

	fmt.Printf("Object: %s\n", t.Name())
	fmt.Printf("Number of methods: %d\n", t.NumMethod())

	// List all methods
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("\nMethod %d: %s\n", i, method.Name)
		fmt.Printf("  Type: %v\n", method.Type)

		// Get method value and call it
		methodValue := v.Method(i)
		if method.Name == "Add" {
			// Call Add method with parameters
			args := []reflect.Value{
				reflect.ValueOf(15),
				reflect.ValueOf(25),
			}
			results := methodValue.Call(args)
			fmt.Printf("  Add(15, 25) = %v\n", results[0].Interface())
		}
	}

	// Call method by name
	fmt.Printf("\nCalling methods by name:\n")
	methodNames := []string{"Add", "Subtract", "Multiply", "Divide"}

	for _, methodName := range methodNames {
		method := v.MethodByName(methodName)
		if !method.IsValid() {
			fmt.Printf("%s method not found\n", methodName)
			continue
		}

		args := []reflect.Value{
			reflect.ValueOf(100),
			reflect.ValueOf(20),
		}

		results := method.Call(args)
		if len(results) == 1 {
			fmt.Printf("%s(100, 20) = %v\n", methodName, results[0].Interface())
		} else if len(results) == 2 {
			// Handle error return
			result := results[0].Interface()
			err := results[1].Interface()
			fmt.Printf("%s(100, 20) = %v, error: %v\n", methodName, result, err)
		}
	}
}

// Example 3: Dynamic function calls with parameters
func dynamicFunctionCalls() {
	fmt.Println("\n--- Example 3: Dynamic Function Calls ---")

	// Function registry
	functions := map[string]interface{}{
		"add":       simpleAdd,
		"divide":    divideWithError,
		"sum":       sum,
		"greet":     greetUser,
		"handleCtx": sampleHandler,
	}

	// Test calls with different parameters
	testCases := []struct {
		name string
		args []interface{}
	}{
		{"add", []interface{}{5, 3}},
		{"divide", []interface{}{10, 2}},
		{"divide", []interface{}{10, 0}}, // Error case
		{"sum", []interface{}{1, 2, 3, 4, 5}},
		{"greet", []interface{}{"Alice"}},
		{"handleCtx", []interface{}{context.Background(), "test data"}},
	}

	for _, testCase := range testCases {
		fmt.Printf("\nCalling %s with args: %v\n", testCase.name, testCase.args)

		fn, exists := functions[testCase.name]
		if !exists {
			fmt.Printf("Function %s not found\n", testCase.name)
			continue
		}

		callFunction(fn, testCase.args)
	}
}

// Example 4: Function signature validation
func functionSignatureValidation() {
	fmt.Println("\n--- Example 4: Function Signature Validation ---")

	// Define expected signatures
	signatures := map[string]FunctionSignature{
		"Handler": {
			Name: "Handler",
			In:   []reflect.Type{reflect.TypeOf((*context.Context)(nil)).Elem(), reflect.TypeOf("")},
			Out:  []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()},
		},
		"Calculator": {
			Name: "Calculator",
			In:   []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)},
			Out:  []reflect.Type{reflect.TypeOf(0)},
		},
	}

	// Test functions against signatures
	testFunctions := map[string]interface{}{
		"sampleHandler":   sampleHandler,
		"simpleAdd":       simpleAdd,
		"divideWithError": divideWithError,
		"invalidFunc":     func(s string) string { return s },
	}

	for funcName, fn := range testFunctions {
		fmt.Printf("\nValidating function: %s\n", funcName)

		for sigName, expectedSig := range signatures {
			if validateFunctionSignature(fn, expectedSig) {
				fmt.Printf("  ✓ Matches %s signature\n", sigName)
			} else {
				fmt.Printf("  ✗ Does not match %s signature\n", sigName)
			}
		}
	}
}

// Example 5: Handler registration pattern
func handlerRegistration() {
	fmt.Println("\n--- Example 5: Handler Registration Pattern ---")

	registry := NewHandlerRegistry()

	// Register different types of handlers
	registry.Register("user.create", createUserHandler)
	registry.Register("user.update", updateUserHandler)
	registry.Register("system.health", healthCheckHandler)

	// List registered handlers
	fmt.Println("Registered handlers:")
	for name := range registry.handlers {
		fmt.Printf("  - %s\n", name)
	}

	// Execute handlers
	testData := []struct {
		event string
		data  string
	}{
		{"user.create", "new user data"},
		{"user.update", "update user data"},
		{"system.health", ""},
		{"unknown.event", "some data"},
	}

	for _, test := range testData {
		fmt.Printf("\nExecuting handler for: %s\n", test.event)
		err := registry.Execute(test.event, context.Background(), test.data)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

// Example 6: Method discovery and documentation
func methodDiscovery() {
	fmt.Println("\n--- Example 6: Method Discovery ---")

	// Analyze different types
	types := []interface{}{
		SimpleCalculator{},
		&User{},
		LogEventHandler{},
	}

	for _, obj := range types {
		fmt.Printf("\nAnalyzing type: %T\n", obj)
		documentType(obj)
	}
}

// Helper functions

func simpleAdd(a, b int) int {
	return a + b
}

func divideWithError(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func greetUser(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func sampleHandler(ctx context.Context, data string) error {
	fmt.Printf("Handler received: %s\n", data)
	return nil
}

func analyzeFunction(fn interface{}) {
	t := reflect.TypeOf(fn)
	if t.Kind() != reflect.Func {
		fmt.Printf("  Not a function\n")
		return
	}

	fmt.Printf("  Function type: %v\n", t)
	fmt.Printf("  Number of inputs: %d\n", t.NumIn())
	fmt.Printf("  Number of outputs: %d\n", t.NumOut())
	fmt.Printf("  Is variadic: %v\n", t.IsVariadic())

	// Input parameters
	for i := 0; i < t.NumIn(); i++ {
		paramType := t.In(i)
		fmt.Printf("    In[%d]: %v\n", i, paramType)
	}

	// Output parameters
	for i := 0; i < t.NumOut(); i++ {
		returnType := t.Out(i)
		fmt.Printf("    Out[%d]: %v\n", i, returnType)
	}
}

func callFunction(fn interface{}, args []interface{}) {
	v := reflect.ValueOf(fn)
	t := reflect.TypeOf(fn)

	if t.Kind() != reflect.Func {
		fmt.Printf("  Error: not a function\n")
		return
	}

	// Prepare arguments
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	// Check if we have the right number of arguments
	expectedArgs := t.NumIn()
	if t.IsVariadic() {
		if len(args) < expectedArgs-1 {
			fmt.Printf("  Error: not enough arguments (expected at least %d, got %d)\n", expectedArgs-1, len(args))
			return
		}
	} else if len(args) != expectedArgs {
		fmt.Printf("  Error: wrong number of arguments (expected %d, got %d)\n", expectedArgs, len(args))
		return
	}

	// Call function and handle results
	results := v.Call(in)
	fmt.Printf("  Results: ")
	for i, result := range results {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%v", result.Interface())
	}
	fmt.Println()
}

// Function signature validation
type FunctionSignature struct {
	Name string
	In   []reflect.Type
	Out  []reflect.Type
}

func validateFunctionSignature(fn interface{}, expected FunctionSignature) bool {
	t := reflect.TypeOf(fn)
	if t.Kind() != reflect.Func {
		return false
	}

	// Check input parameters
	if t.NumIn() != len(expected.In) {
		return false
	}
	for i, expectedType := range expected.In {
		if t.In(i) != expectedType {
			return false
		}
	}

	// Check output parameters
	if t.NumOut() != len(expected.Out) {
		return false
	}
	for i, expectedType := range expected.Out {
		if t.Out(i) != expectedType {
			return false
		}
	}

	return true
}

// Handler registry implementation
type HandlerRegistry struct {
	handlers map[string]reflect.Value
}

func NewHandlerRegistry() *HandlerRegistry {
	return &HandlerRegistry{
		handlers: make(map[string]reflect.Value),
	}
}

func (hr *HandlerRegistry) Register(event string, handler interface{}) error {
	v := reflect.ValueOf(handler)
	t := reflect.TypeOf(handler)

	// Validate handler signature
	if t.Kind() != reflect.Func {
		return fmt.Errorf("handler must be a function")
	}

	// Expected: func(context.Context, string) error
	expectedSig := FunctionSignature{
		In:  []reflect.Type{reflect.TypeOf((*context.Context)(nil)).Elem(), reflect.TypeOf("")},
		Out: []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()},
	}

	if !validateFunctionSignature(handler, expectedSig) {
		return fmt.Errorf("handler must have signature: func(context.Context, string) error")
	}

	hr.handlers[event] = v
	return nil
}

func (hr *HandlerRegistry) Execute(event string, ctx context.Context, data string) error {
	handler, exists := hr.handlers[event]
	if !exists {
		return fmt.Errorf("no handler registered for event: %s", event)
	}

	args := []reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(data),
	}

	results := handler.Call(args)
	if len(results) > 0 {
		if err, ok := results[0].Interface().(error); ok {
			return err
		}
	}

	return nil
}

// Sample handlers
func createUserHandler(ctx context.Context, data string) error {
	fmt.Printf("Creating user with data: %s\n", data)
	return nil
}

func updateUserHandler(ctx context.Context, data string) error {
	fmt.Printf("Updating user with data: %s\n", data)
	return nil
}

func healthCheckHandler(ctx context.Context, data string) error {
	fmt.Printf("Health check OK\n")
	return nil
}

// Type documentation
func documentType(obj interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	// Handle pointers
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	fmt.Printf("  Type: %s\n", t.Name())
	fmt.Printf("  Kind: %s\n", t.Kind())
	fmt.Printf("  Package: %s\n", t.PkgPath())

	// Methods
	numMethods := t.NumMethod()
	if numMethods > 0 {
		fmt.Printf("  Methods (%d):\n", numMethods)
		for i := 0; i < numMethods; i++ {
			method := t.Method(i)
			fmt.Printf("    %s%s\n", method.Name, formatFunctionSignature(method.Type))
		}
	}

	// Fields (if struct)
	if t.Kind() == reflect.Struct {
		numFields := t.NumField()
		if numFields > 0 {
			fmt.Printf("  Fields (%d):\n", numFields)
			for i := 0; i < numFields; i++ {
				field := t.Field(i)
				exported := ""
				if field.PkgPath == "" {
					exported = " (exported)"
				}
				fmt.Printf("    %s %s%s\n", field.Name, field.Type, exported)
			}
		}
	}
}

func formatFunctionSignature(t reflect.Type) string {
	if t.Kind() != reflect.Func {
		return ""
	}

	sig := "("
	for i := 0; i < t.NumIn(); i++ {
		if i > 0 {
			sig += ", "
		}
		sig += t.In(i).String()
	}
	sig += ")"

	if t.NumOut() > 0 {
		sig += " ("
		for i := 0; i < t.NumOut(); i++ {
			if i > 0 {
				sig += ", "
			}
			sig += t.Out(i).String()
		}
		sig += ")"
	}

	return sig
}
