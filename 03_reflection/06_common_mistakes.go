package main

import (
	"fmt"
	"reflect"
)

// CommonMistakesExamples demonstrates common reflection mistakes and how to avoid them
func CommonMistakesExamples() {
	fmt.Println("\n=== Common Mistakes & Error Handling Examples ===")

	// Example 1: Nil pointer and nil interface pitfalls
	nilPointerPitfalls()

	// Example 2: CanSet() violations and how to fix them
	canSetViolations()

	// Example 3: Type assertion failures
	typeAssertionFailures()

	// Example 4: Kind vs Type confusion
	kindTypeConfusion()

	// Example 5: Panic-safe reflection operations
	panicSafeOperations()

	// Example 6: Performance pitfalls and best practices
	performancePitfalls()
}

// Example 1: Nil pointer and nil interface pitfalls
func nilPointerPitfalls() {
	fmt.Println("\n--- Example 1: Nil Pointer Pitfalls ---")

	// Common mistake: Calling methods on nil interface
	fmt.Printf("Mistake 1: Nil interface operations\n")
	demonstrateNilInterfaceMistakes()

	// Common mistake: Nil pointer dereference
	fmt.Printf("\nMistake 2: Nil pointer dereference\n")
	demonstrateNilPointerMistakes()

	// Common mistake: Checking nil incorrectly
	fmt.Printf("\nMistake 3: Incorrect nil checking\n")
	demonstrateIncorrectNilChecking()

	// Safe practices
	fmt.Printf("\nSafe practices:\n")
	demonstrateSafeNilHandling()
}

// Example 2: CanSet() violations
func canSetViolations() {
	fmt.Println("\n--- Example 2: CanSet() Violations ---")

	// Common mistake: Trying to set unexported fields
	fmt.Printf("Mistake 1: Setting unexported fields\n")
	demonstrateUnexportedFieldMistake()

	// Common mistake: Setting non-addressable values
	fmt.Printf("\nMistake 2: Setting non-addressable values\n")
	demonstrateNonAddressableMistake()

	// Common mistake: Setting through non-pointer
	fmt.Printf("\nMistake 3: Setting through non-pointer\n")
	demonstrateNonPointerMistake()

	// Safe practices
	fmt.Printf("\nSafe practices:\n")
	demonstrateSafeSetOperations()
}

// Example 3: Type assertion failures
func typeAssertionFailures() {
	fmt.Println("\n--- Example 3: Type Assertion Failures ---")

	// Common mistake: Unsafe type assertions
	fmt.Printf("Mistake 1: Unsafe type assertions\n")
	demonstrateUnsafeTypeAssertions()

	// Common mistake: Wrong type expectations
	fmt.Printf("\nMistake 2: Wrong type expectations\n")
	demonstrateWrongTypeExpectations()

	// Safe practices
	fmt.Printf("\nSafe practices:\n")
	demonstrateSafeTypeAssertions()
}

// Example 4: Kind vs Type confusion
func kindTypeConfusion() {
	fmt.Println("\n--- Example 4: Kind vs Type Confusion ---")

	// Common mistake: Using Kind when Type is needed
	fmt.Printf("Mistake 1: Kind vs Type confusion\n")
	demonstrateKindTypeConfusion()

	// Common mistake: Wrong type comparisons
	fmt.Printf("\nMistake 2: Wrong type comparisons\n")
	demonstrateWrongTypeComparisons()

	// Safe practices
	fmt.Printf("\nSafe practices:\n")
	demonstrateCorrectTypeHandling()
}

// Example 5: Panic-safe reflection operations
func panicSafeOperations() {
	fmt.Println("\n--- Example 5: Panic-Safe Operations ---")

	// Demonstrate safe wrappers for common operations
	fmt.Printf("Safe reflection wrappers:\n")
	demonstrateSafeReflectionWrappers()

	// Error handling patterns
	fmt.Printf("\nError handling patterns:\n")
	demonstrateErrorHandlingPatterns()
}

// Example 6: Performance pitfalls
func performancePitfalls() {
	fmt.Println("\n--- Example 6: Performance Pitfalls ---")

	// Common mistake: Repeated reflection operations
	fmt.Printf("Mistake 1: Repeated reflection in loops\n")
	demonstrateReflectionInLoops()

	// Common mistake: Not caching reflection results
	fmt.Printf("\nMistake 2: Not caching reflection results\n")
	demonstrateNoCaching()

	// Best practices
	fmt.Printf("\nBest practices:\n")
	demonstratePerformanceBestPractices()
}

// Helper functions for demonstrations

func demonstrateNilInterfaceMistakes() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  PANIC caught: %v\n", r)
		}
	}()

	var nilInterface interface{}

	fmt.Printf("  Attempting to get type of nil interface...\n")
	t := reflect.TypeOf(nilInterface)
	if t == nil {
		fmt.Printf("  TypeOf(nil) returns nil - this is correct\n")
	}

	// This would panic if not handled
	fmt.Printf("  Attempting unsafe operations on nil...\n")
	// Uncomment to see panic:
	// v := reflect.ValueOf(nilInterface)
	// fmt.Printf("Value: %v\n", v.Interface()) // PANIC!
}

func demonstrateNilPointerMistakes() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  PANIC caught: %v\n", r)
		}
	}()

	var nilPerson *Person

	v := reflect.ValueOf(nilPerson)
	fmt.Printf("  Nil pointer value: %v\n", v)
	fmt.Printf("  Is nil: %v\n", v.IsNil())

	// This would panic:
	fmt.Printf("  Attempting to dereference nil pointer...\n")
	// Uncomment to see panic:
	// elem := v.Elem() // PANIC!
	// fmt.Printf("Element: %v\n", elem)
}

func demonstrateIncorrectNilChecking() {
	// Common mistake: checking interface{} == nil
	var person *Person = nil
	var iface interface{} = person

	fmt.Printf("  person == nil: %v\n", person == nil)
	fmt.Printf("  iface == nil: %v\n", iface == nil) // This is false!

	// Correct way using reflection
	v := reflect.ValueOf(iface)
	fmt.Printf("  reflect.ValueOf(iface).IsNil(): %v\n", v.IsNil())
}

func demonstrateSafeNilHandling() {
	examples := []interface{}{
		nil,
		(*Person)(nil),
		Person{},
		&Person{},
	}

	for i, example := range examples {
		fmt.Printf("  Example %d: ", i+1)
		if err := safeTypeAnalysis(example); err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Safe analysis completed\n")
		}
	}
}

func demonstrateUnexportedFieldMistake() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  PANIC caught: %v\n", r)
		}
	}()

	// Create a struct with unexported fields
	type privateStruct struct {
		Public  string
		private string
	}

	obj := privateStruct{Public: "public", private: "private"}
	v := reflect.ValueOf(&obj).Elem()

	// This works
	publicField := v.FieldByName("Public")
	fmt.Printf("  Public field CanSet: %v\n", publicField.CanSet())

	// This won't work
	privateField := v.FieldByName("private")
	fmt.Printf("  Private field CanSet: %v\n", privateField.CanSet())

	// Attempting to set private field would panic:
	// privateField.SetString("modified") // PANIC!
}

func demonstrateNonAddressableMistake() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  PANIC caught: %v\n", r)
		}
	}()

	person := Person{Name: "Alice", Age: 30}
	v := reflect.ValueOf(person) // Non-pointer value

	nameField := v.FieldByName("Name")
	fmt.Printf("  Non-addressable field CanSet: %v\n", nameField.CanSet())
	fmt.Printf("  Non-addressable field CanAddr: %v\n", nameField.CanAddr())

	// This would panic:
	// nameField.SetString("Bob") // PANIC!
}

func demonstrateNonPointerMistake() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  PANIC caught: %v\n", r)
		}
	}()

	person := Person{Name: "Alice", Age: 30}

	// Wrong: passing value
	v1 := reflect.ValueOf(person)
	fmt.Printf("  Value CanSet: %v\n", v1.CanSet())

	// Correct: passing pointer and getting element
	v2 := reflect.ValueOf(&person).Elem()
	fmt.Printf("  Pointer.Elem() CanSet: %v\n", v2.CanSet())
}

func demonstrateSafeSetOperations() {
	person := &Person{Name: "Alice", Age: 30}

	if err := safeSetField(person, "Name", "Bob"); err != nil {
		fmt.Printf("  Error setting Name: %v\n", err)
	} else {
		fmt.Printf("  Successfully set Name to: %s\n", person.Name)
	}

	if err := safeSetField(person, "Age", 35); err != nil {
		fmt.Printf("  Error setting Age: %v\n", err)
	} else {
		fmt.Printf("  Successfully set Age to: %d\n", person.Age)
	}

	// Try to set invalid field
	if err := safeSetField(person, "InvalidField", "value"); err != nil {
		fmt.Printf("  Expected error for invalid field: %v\n", err)
	}
}

func demonstrateUnsafeTypeAssertions() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  PANIC caught: %v\n", r)
		}
	}()

	var value interface{} = "hello"
	v := reflect.ValueOf(value)

	// Unsafe: direct Interface() call without checking
	fmt.Printf("  Attempting unsafe type assertion...\n")
	// This works but is unsafe if v was invalid:
	result := v.Interface().(string)
	fmt.Printf("  Result: %s\n", result)

	// Unsafe with wrong type:
	fmt.Printf("  Attempting wrong type assertion...\n")
	// This would panic:
	// wrongResult := v.Interface().(int) // PANIC!
}

func demonstrateWrongTypeExpectations() {
	values := []interface{}{
		42,
		"hello",
		3.14,
		Person{Name: "Alice"},
	}

	for i, val := range values {
		v := reflect.ValueOf(val)
		fmt.Printf("  Value %d (%v): ", i+1, val)

		// Wrong: assuming all values are strings
		if v.Kind() == reflect.String {
			fmt.Printf("is string: %s\n", v.String())
		} else {
			fmt.Printf("not a string (kind: %v)\n", v.Kind())
		}
	}
}

func demonstrateSafeTypeAssertions() {
	values := []interface{}{
		"hello",
		42,
		Person{Name: "Alice"},
		nil,
	}

	for i, val := range values {
		fmt.Printf("  Value %d: ", i+1)
		if result, err := safeStringExtraction(val); err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("String: %s\n", result)
		}
	}
}

func demonstrateKindTypeConfusion() {
	type MyInt int
	type MyString string

	values := []interface{}{
		int(42),
		MyInt(42),
		string("hello"),
		MyString("hello"),
	}

	for i, val := range values {
		v := reflect.ValueOf(val)
		t := reflect.TypeOf(val)

		fmt.Printf("  Value %d: %v\n", i+1, val)
		fmt.Printf("    Kind: %v\n", v.Kind())
		fmt.Printf("    Type: %v\n", t)

		// Wrong: using Kind for type-specific logic
		if v.Kind() == reflect.Int {
			fmt.Printf("    Treated as int (but might be MyInt!)\n")
		}

		// Correct: using Type for exact type checking
		if t == reflect.TypeOf(int(0)) {
			fmt.Printf("    Exactly int type\n")
		} else if t == reflect.TypeOf(MyInt(0)) {
			fmt.Printf("    Exactly MyInt type\n")
		}
	}
}

func demonstrateWrongTypeComparisons() {
	type CustomString string

	val1 := "hello"
	val2 := CustomString("hello")

	t1 := reflect.TypeOf(val1)
	t2 := reflect.TypeOf(val2)

	fmt.Printf("  string vs CustomString:\n")
	fmt.Printf("    Same kind: %v\n", t1.Kind() == t2.Kind())
	fmt.Printf("    Same type: %v\n", t1 == t2)
	fmt.Printf("    Convertible: %v\n", t1.ConvertibleTo(t2))
	fmt.Printf("    Assignable: %v\n", t1.AssignableTo(t2))
}

func demonstrateCorrectTypeHandling() {
	values := []interface{}{
		int(42),
		int32(42),
		int64(42),
		float32(3.14),
		float64(3.14),
	}

	for i, val := range values {
		v := reflect.ValueOf(val)
		fmt.Printf("  Value %d (%v): ", i+1, val)

		// Correct: handle by kind for similar operations
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Printf("integer type, value: %d\n", v.Int())
		case reflect.Float32, reflect.Float64:
			fmt.Printf("float type, value: %.2f\n", v.Float())
		default:
			fmt.Printf("other type\n")
		}
	}
}

func demonstrateSafeReflectionWrappers() {
	person := &Person{Name: "Alice", Age: 30}

	// Safe field access
	if name, err := safeGetField(person, "Name"); err != nil {
		fmt.Printf("  Error getting Name: %v\n", err)
	} else {
		fmt.Printf("  Name: %v\n", name)
	}

	// Safe method call
	calc := SimpleCalculator{Name: "TestCalc"}
	if result, err := safeMethodCall(calc, "Add", 10, 20); err != nil {
		fmt.Printf("  Error calling Add: %v\n", err)
	} else {
		fmt.Printf("  Add result: %v\n", result)
	}
}

func demonstrateErrorHandlingPatterns() {
	// Pattern 1: Defensive checking
	fmt.Printf("  Pattern 1: Defensive checking\n")
	defensiveReflectionOperation(nil)
	defensiveReflectionOperation(&Person{Name: "Alice"})

	// Pattern 2: Error wrapping
	fmt.Printf("  Pattern 2: Error wrapping\n")
	if err := complexReflectionOperation("invalid"); err != nil {
		fmt.Printf("    Error: %v\n", err)
	}
}

func demonstrateReflectionInLoops() {
	// Bad: reflection in loop
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	fmt.Printf("  Bad practice - reflection in loop:\n")
	for i, person := range people {
		v := reflect.ValueOf(person)
		nameField := v.FieldByName("Name") // Reflection lookup every iteration!
		fmt.Printf("    Person %d: %s\n", i+1, nameField.String())
	}
}

func demonstrateNoCaching() {
	person := Person{Name: "Alice", Age: 30}

	fmt.Printf("  Bad practice - no caching:\n")
	for i := 0; i < 3; i++ {
		v := reflect.ValueOf(person)
		t := reflect.TypeOf(person)
		nameField, _ := t.FieldByName("Name") // Lookup every time!
		fmt.Printf("    Iteration %d: Field %s\n", i+1, nameField.Name)
		fmt.Printf("    Iteration %d: Value %v\n", i+1, v.FieldByName("Name"))
	}
}

func demonstratePerformanceBestPractices() {
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	fmt.Printf("  Good practice - caching:\n")

	// Cache reflection metadata
	personType := reflect.TypeOf(Person{})
	nameFieldIndex := -1
	for i := 0; i < personType.NumField(); i++ {
		if personType.Field(i).Name == "Name" {
			nameFieldIndex = i
			break
		}
	}

	// Use cached metadata
	for i, person := range people {
		v := reflect.ValueOf(person)
		nameField := v.Field(nameFieldIndex) // Direct field access!
		fmt.Printf("    Person %d: %s\n", i+1, nameField.String())
	}
}

// Safe wrapper functions

func safeTypeAnalysis(value interface{}) error {
	if value == nil {
		fmt.Printf("nil value\n")
		return nil
	}

	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	if !v.IsValid() {
		return fmt.Errorf("invalid value")
	}

	fmt.Printf("type: %v, kind: %v", t, v.Kind())

	// Safe nil checking for pointer types
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func, reflect.Slice, reflect.Map:
		if v.IsNil() {
			fmt.Printf(", is nil: true")
		} else {
			fmt.Printf(", is nil: false")
		}
	}

	fmt.Printf("\n")
	return nil
}

func safeSetField(obj interface{}, fieldName string, value interface{}) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("object must be a pointer")
	}

	v = v.Elem()
	if !v.IsValid() {
		return fmt.Errorf("invalid object")
	}

	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("field %s not found", fieldName)
	}

	if !field.CanSet() {
		return fmt.Errorf("field %s cannot be set", fieldName)
	}

	newValue := reflect.ValueOf(value)
	if !newValue.Type().AssignableTo(field.Type()) {
		if !newValue.Type().ConvertibleTo(field.Type()) {
			return fmt.Errorf("value type %v not assignable to field type %v", newValue.Type(), field.Type())
		}
		newValue = newValue.Convert(field.Type())
	}

	field.Set(newValue)
	return nil
}

func safeGetField(obj interface{}, fieldName string) (interface{}, error) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if !v.IsValid() {
		return nil, fmt.Errorf("invalid object")
	}

	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, fmt.Errorf("field %s not found", fieldName)
	}

	if !field.CanInterface() {
		return nil, fmt.Errorf("field %s cannot be accessed", fieldName)
	}

	return field.Interface(), nil
}

func safeMethodCall(obj interface{}, methodName string, args ...interface{}) (interface{}, error) {
	v := reflect.ValueOf(obj)
	method := v.MethodByName(methodName)

	if !method.IsValid() {
		return nil, fmt.Errorf("method %s not found", methodName)
	}

	// Prepare arguments
	argValues := make([]reflect.Value, len(args))
	for i, arg := range args {
		argValues[i] = reflect.ValueOf(arg)
	}

	// Call method
	results := method.Call(argValues)
	if len(results) == 0 {
		return nil, nil
	}

	return results[0].Interface(), nil
}

func safeStringExtraction(value interface{}) (string, error) {
	if value == nil {
		return "", fmt.Errorf("nil value")
	}

	v := reflect.ValueOf(value)
	if !v.IsValid() {
		return "", fmt.Errorf("invalid value")
	}

	switch v.Kind() {
	case reflect.String:
		return v.String(), nil
	default:
		// Try to convert to string
		stringType := reflect.TypeOf("")
		if v.Type().ConvertibleTo(stringType) {
			converted := v.Convert(stringType)
			return converted.String(), nil
		}
		return "", fmt.Errorf("value of type %v cannot be converted to string", v.Type())
	}
}

func defensiveReflectionOperation(obj interface{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("    Recovered from panic: %v\n", r)
		}
	}()

	if obj == nil {
		fmt.Printf("    Safely handled nil object\n")
		return
	}

	v := reflect.ValueOf(obj)
	if !v.IsValid() {
		fmt.Printf("    Safely handled invalid value\n")
		return
	}

	fmt.Printf("    Successfully processed: %v\n", v.Type())
}

func complexReflectionOperation(input interface{}) error {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Errorf("reflection panic: %v", r))
		}
	}()

	v := reflect.ValueOf(input)
	if !v.IsValid() {
		return fmt.Errorf("invalid input")
	}

	// Simulate complex operation that might fail
	if v.Kind() != reflect.String {
		return fmt.Errorf("expected string, got %v", v.Kind())
	}

	return nil
}
