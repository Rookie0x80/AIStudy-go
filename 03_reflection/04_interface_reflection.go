package main

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

// InterfaceReflectionExamples demonstrates interface reflection
func InterfaceReflectionExamples() {
	fmt.Println("\n=== Interface Reflection Examples ===")

	// Example 1: Basic interface type checking
	basicInterfaceReflection()

	// Example 2: Dynamic type assertion and extraction
	dynamicTypeAssertion()

	// Example 3: Interface implementation checking
	interfaceImplementationChecking()

	// Example 4: Empty interface handling
	emptyInterfaceHandling()

	// Example 5: Interface composition analysis
	interfaceComposition()

	// Example 6: Runtime interface satisfaction
	runtimeInterfaceSatisfaction()
}

// Example 1: Basic interface type checking
func basicInterfaceReflection() {
	fmt.Println("\n--- Example 1: Basic Interface Reflection ---")

	// Various interface values
	var w io.Writer = &strings.Builder{}
	var calc Calculator = SimpleCalculator{Name: "TestCalc"}
	var handler EventHandler = LogEventHandler{Name: "Logger"}
	var empty interface{} = "hello world"

	interfaces := map[string]interface{}{
		"io.Writer":    w,
		"Calculator":   calc,
		"EventHandler": handler,
		"empty":        empty,
	}

	for name, iface := range interfaces {
		fmt.Printf("\nAnalyzing interface: %s\n", name)
		analyzeInterface(iface)
	}
}

// Example 2: Dynamic type assertion and extraction
func dynamicTypeAssertion() {
	fmt.Println("\n--- Example 2: Dynamic Type Assertion ---")

	// Collection of different types stored as interfaces
	values := []interface{}{
		42,
		"hello",
		3.14,
		true,
		[]int{1, 2, 3},
		map[string]int{"key": 42},
		Person{Name: "Alice", Age: 30},
		&Person{Name: "Bob", Age: 25},
		SimpleCalculator{Name: "Calc"},
	}

	for i, val := range values {
		fmt.Printf("\nValue %d: %v\n", i+1, val)

		v := reflect.ValueOf(val)
		t := reflect.TypeOf(val)

		fmt.Printf("  Static type: %v\n", t)
		fmt.Printf("  Dynamic type: %v\n", v.Type())
		fmt.Printf("  Kind: %v\n", v.Kind())

		// Perform type-specific operations
		performTypeSpecificOperations(val)

		// Try to extract concrete type
		extractConcreteType(val)
	}
}

// Example 3: Interface implementation checking
func interfaceImplementationChecking() {
	fmt.Println("\n--- Example 3: Interface Implementation Checking ---")

	// Define test types
	types := []interface{}{
		Person{},
		&Person{},
		SimpleCalculator{},
		&SimpleCalculator{},
		LogEventHandler{},
		&LogEventHandler{},
		strings.Builder{},
		&strings.Builder{},
		42,
		"string",
	}

	// Define interfaces to check against
	interfaceTypes := map[string]reflect.Type{
		"io.Writer":    reflect.TypeOf((*io.Writer)(nil)).Elem(),
		"Calculator":   reflect.TypeOf((*Calculator)(nil)).Elem(),
		"EventHandler": reflect.TypeOf((*EventHandler)(nil)).Elem(),
		"fmt.Stringer": reflect.TypeOf((*fmt.Stringer)(nil)).Elem(),
	}

	fmt.Printf("Interface implementation matrix:\n")
	fmt.Printf("%-20s", "Type\\Interface")
	for ifaceName := range interfaceTypes {
		fmt.Printf("%-15s", ifaceName)
	}
	fmt.Println()

	for _, obj := range types {
		t := reflect.TypeOf(obj)
		fmt.Printf("%-20s", t.String())

		for _, ifaceType := range interfaceTypes {
			implements := t.Implements(ifaceType)
			if implements {
				fmt.Printf("%-15s", "✓")
			} else {
				fmt.Printf("%-15s", "✗")
			}
		}
		fmt.Println()
	}
}

// Example 4: Empty interface handling
func emptyInterfaceHandling() {
	fmt.Println("\n--- Example 4: Empty Interface Handling ---")

	// Various types stored in empty interfaces
	var items []interface{} = []interface{}{
		nil,
		42,
		"hello",
		3.14,
		true,
		[]string{"a", "b", "c"},
		map[string]interface{}{"key": "value"},
		Person{Name: "Alice"},
		func() { fmt.Println("function") },
		make(chan int),
	}

	for i, item := range items {
		fmt.Printf("\nItem %d:\n", i+1)
		processEmptyInterface(item)
	}
}

// Example 5: Interface composition analysis
func interfaceComposition() {
	fmt.Println("\n--- Example 5: Interface Composition Analysis ---")

	// Analyze ReadWriter interface (composed of Reader and Writer)
	rwType := reflect.TypeOf((*ReadWriter)(nil)).Elem()
	fmt.Printf("Analyzing interface: %s\n", rwType.Name())
	fmt.Printf("Number of methods: %d\n", rwType.NumMethod())

	// List all methods in the composed interface
	fmt.Printf("Interface methods:\n")
	for i := 0; i < rwType.NumMethod(); i++ {
		method := rwType.Method(i)
		fmt.Printf("  %s%s\n", method.Name, formatFunctionSignature(method.Type))
	}

	// Analyze component interfaces
	fmt.Printf("\nComponent interfaces:\n")
	readerType := reflect.TypeOf((*Reader)(nil)).Elem()
	writerType := reflect.TypeOf((*Writer)(nil)).Elem()

	fmt.Printf("  Reader interface (%d methods):\n", readerType.NumMethod())
	for i := 0; i < readerType.NumMethod(); i++ {
		method := readerType.Method(i)
		fmt.Printf("    %s%s\n", method.Name, formatFunctionSignature(method.Type))
	}

	fmt.Printf("  Writer interface (%d methods):\n", writerType.NumMethod())
	for i := 0; i < writerType.NumMethod(); i++ {
		method := writerType.Method(i)
		fmt.Printf("    %s%s\n", method.Name, formatFunctionSignature(method.Type))
	}

	// Check what types implement this interface
	testTypes := []reflect.Type{
		reflect.TypeOf(&strings.Builder{}),
		reflect.TypeOf(Person{}),
		reflect.TypeOf(SimpleCalculator{}),
	}

	fmt.Printf("\nImplementation check:\n")
	for _, t := range testTypes {
		implements := t.Implements(rwType)
		fmt.Printf("  %s implements ReadWriter: %v\n", t, implements)
	}
}

// Example 6: Runtime interface satisfaction
func runtimeInterfaceSatisfaction() {
	fmt.Println("\n--- Example 6: Runtime Interface Satisfaction ---")

	// Create an interface satisfaction checker
	checker := NewInterfaceSatisfactionChecker()

	// Register some interfaces
	checker.RegisterInterface("Calculator", reflect.TypeOf((*Calculator)(nil)).Elem())
	checker.RegisterInterface("EventHandler", reflect.TypeOf((*EventHandler)(nil)).Elem())
	checker.RegisterInterface("io.Writer", reflect.TypeOf((*io.Writer)(nil)).Elem())

	// Test various objects
	testObjects := []interface{}{
		SimpleCalculator{Name: "Calc"},
		&SimpleCalculator{Name: "CalcPtr"},
		LogEventHandler{Name: "Handler"},
		&LogEventHandler{Name: "HandlerPtr"},
		&strings.Builder{},
		Person{Name: "Alice"},
		42,
	}

	for _, obj := range testObjects {
		fmt.Printf("\nTesting object: %T\n", obj)
		satisfiedInterfaces := checker.GetSatisfiedInterfaces(obj)
		if len(satisfiedInterfaces) > 0 {
			fmt.Printf("  Satisfies: %v\n", satisfiedInterfaces)
		} else {
			fmt.Printf("  Satisfies no registered interfaces\n")
		}
	}
}

// Helper functions

func analyzeInterface(iface interface{}) {
	v := reflect.ValueOf(iface)
	t := reflect.TypeOf(iface)

	fmt.Printf("  Value: %v\n", iface)
	fmt.Printf("  Type: %v\n", t)
	fmt.Printf("  Kind: %v\n", t.Kind())

	// Check if it's actually an interface type
	if t.Kind() == reflect.Interface {
		fmt.Printf("  Is interface type: true\n")
		fmt.Printf("  Number of methods: %d\n", t.NumMethod())

		// List interface methods
		for i := 0; i < t.NumMethod(); i++ {
			method := t.Method(i)
			fmt.Printf("    Method %d: %s\n", i, method.Name)
		}
	} else {
		fmt.Printf("  Is interface type: false\n")
		fmt.Printf("  Concrete type: %v\n", t)
	}

	// Check if value is nil (only for types that can be nil)
	if v.IsValid() {
		switch v.Kind() {
		case reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func, reflect.Slice, reflect.Map:
			fmt.Printf("  Is nil: %v\n", v.IsNil())
		default:
			fmt.Printf("  Is nil: N/A (cannot be nil)\n")
		}
		fmt.Printf("  Is zero: %v\n", v.IsZero())
	} else {
		fmt.Printf("  Value is invalid\n")
	}
}

func performTypeSpecificOperations(val interface{}) {
	v := reflect.ValueOf(val)

	switch v.Kind() {
	case reflect.String:
		s := v.String()
		fmt.Printf("  String length: %d\n", len(s))
	case reflect.Int:
		i := v.Int()
		fmt.Printf("  Integer value: %d, is even: %v\n", i, i%2 == 0)
	case reflect.Float64:
		f := v.Float()
		fmt.Printf("  Float value: %.2f, is integer: %v\n", f, f == float64(int64(f)))
	case reflect.Bool:
		b := v.Bool()
		fmt.Printf("  Boolean value: %v\n", b)
	case reflect.Slice:
		fmt.Printf("  Slice length: %d, capacity: %d\n", v.Len(), v.Cap())
	case reflect.Map:
		fmt.Printf("  Map length: %d\n", v.Len())
		if v.Len() > 0 {
			fmt.Printf("  Keys: %v\n", v.MapKeys())
		}
	case reflect.Struct:
		t := v.Type()
		fmt.Printf("  Struct with %d fields\n", t.NumField())
	case reflect.Ptr:
		if !v.IsNil() {
			fmt.Printf("  Pointer to: %v\n", v.Elem().Type())
		} else {
			fmt.Printf("  Nil pointer\n")
		}
	default:
		fmt.Printf("  Unhandled kind: %v\n", v.Kind())
	}
}

func extractConcreteType(val interface{}) {
	v := reflect.ValueOf(val)
	t := reflect.TypeOf(val)

	fmt.Printf("  Extracting concrete type:\n")

	// If it's an interface, get the dynamic type
	if t.Kind() == reflect.Interface {
		if !v.IsNil() {
			concrete := v.Elem()
			concreteType := concrete.Type()
			fmt.Printf("    Dynamic type: %v\n", concreteType)
			fmt.Printf("    Dynamic value: %v\n", concrete.Interface())
		} else {
			fmt.Printf("    Nil interface\n")
		}
	} else {
		fmt.Printf("    Already concrete type: %v\n", t)
	}

	// Try type assertions for common types
	switch val := val.(type) {
	case string:
		fmt.Printf("    As string: %q\n", val)
	case int:
		fmt.Printf("    As int: %d\n", val)
	case Person:
		fmt.Printf("    As Person: %+v\n", val)
	case *Person:
		fmt.Printf("    As *Person: %+v\n", val)
	default:
		fmt.Printf("    Unknown concrete type\n")
	}
}

func processEmptyInterface(item interface{}) {
	v := reflect.ValueOf(item)

	if !v.IsValid() {
		fmt.Printf("  Invalid value (nil)\n")
		return
	}

	t := v.Type()
	fmt.Printf("  Type: %v\n", t)
	fmt.Printf("  Kind: %v\n", v.Kind())
	fmt.Printf("  Value: %v\n", item)

	// Check if nil
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func, reflect.Slice, reflect.Map:
		fmt.Printf("  Is nil: %v\n", v.IsNil())
	}

	// Type switch simulation using reflection
	switch v.Kind() {
	case reflect.String:
		fmt.Printf("  Processing as string: length = %d\n", v.Len())
	case reflect.Int:
		fmt.Printf("  Processing as int: value * 2 = %d\n", v.Int()*2)
	case reflect.Float64:
		fmt.Printf("  Processing as float: squared = %.2f\n", v.Float()*v.Float())
	case reflect.Bool:
		fmt.Printf("  Processing as bool: negated = %v\n", !v.Bool())
	case reflect.Slice:
		fmt.Printf("  Processing as slice: length = %d\n", v.Len())
		if v.Type().Elem().Kind() == reflect.String {
			fmt.Printf("    String slice contents: ")
			for i := 0; i < v.Len(); i++ {
				if i > 0 {
					fmt.Printf(", ")
				}
				fmt.Printf("%q", v.Index(i).String())
			}
			fmt.Println()
		}
	case reflect.Map:
		fmt.Printf("  Processing as map: %d keys\n", v.Len())
	case reflect.Struct:
		fmt.Printf("  Processing as struct: %d fields\n", t.NumField())
		if t.Name() == "Person" {
			nameField := v.FieldByName("Name")
			if nameField.IsValid() {
				fmt.Printf("    Person name: %s\n", nameField.String())
			}
		}
	case reflect.Func:
		fmt.Printf("  Processing as function: %v\n", t)
	case reflect.Chan:
		fmt.Printf("  Processing as channel: %v\n", t)
	}
}

func formatMethodSignature(method reflect.Method) string {
	return formatFunctionSignature(method.Type)
}

// Interface satisfaction checker
type InterfaceSatisfactionChecker struct {
	interfaces map[string]reflect.Type
}

func NewInterfaceSatisfactionChecker() *InterfaceSatisfactionChecker {
	return &InterfaceSatisfactionChecker{
		interfaces: make(map[string]reflect.Type),
	}
}

func (isc *InterfaceSatisfactionChecker) RegisterInterface(name string, ifaceType reflect.Type) {
	if ifaceType.Kind() != reflect.Interface {
		panic(fmt.Sprintf("Type %v is not an interface", ifaceType))
	}
	isc.interfaces[name] = ifaceType
}

func (isc *InterfaceSatisfactionChecker) GetSatisfiedInterfaces(obj interface{}) []string {
	t := reflect.TypeOf(obj)
	var satisfied []string

	for name, ifaceType := range isc.interfaces {
		if t.Implements(ifaceType) {
			satisfied = append(satisfied, name)
		}
	}

	return satisfied
}

func (isc *InterfaceSatisfactionChecker) CheckSatisfaction(obj interface{}, interfaceName string) bool {
	ifaceType, exists := isc.interfaces[interfaceName]
	if !exists {
		return false
	}

	t := reflect.TypeOf(obj)
	return t.Implements(ifaceType)
}
