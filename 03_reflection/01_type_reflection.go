package main

import (
	"fmt"
	"reflect"
)

// TypeReflectionExamples demonstrates type system reflection
func TypeReflectionExamples() {
	fmt.Println("\n=== Type System Reflection Examples ===")

	// Example 1: Basic TypeOf and ValueOf
	basicTypeOfValueOf()

	// Example 2: Kind vs Type distinction
	kindVsType()

	// Example 3: Type information and properties
	typeInformation()

	// Example 4: Method sets and type methods
	methodSets()

	// Example 5: Type comparison and conversion
	typeComparison()

	// Example 6: Working with pointers and indirection
	pointerReflection()
}

// Example 1: Basic TypeOf and ValueOf operations
func basicTypeOfValueOf() {
	fmt.Println("\n--- Example 1: Basic TypeOf and ValueOf ---")

	var x int = 42
	var s string = "hello"
	var p *Person = &Person{Name: "Alice", Age: 30}

	// Get type information
	fmt.Printf("Type of x: %v\n", reflect.TypeOf(x))
	fmt.Printf("Type of s: %v\n", reflect.TypeOf(s))
	fmt.Printf("Type of p: %v\n", reflect.TypeOf(p))

	// Get value information
	vx := reflect.ValueOf(x)
	vs := reflect.ValueOf(s)
	vp := reflect.ValueOf(p)

	fmt.Printf("Value of x: %v, Type: %v\n", vx, vx.Type())
	fmt.Printf("Value of s: %v, Type: %v\n", vs, vs.Type())
	fmt.Printf("Value of p: %v, Type: %v\n", vp, vp.Type())

	// Check if value is valid and can be interfaced
	fmt.Printf("vx is valid: %v, can interface: %v\n", vx.IsValid(), vx.CanInterface())
	if vx.CanInterface() {
		fmt.Printf("vx interface: %v\n", vx.Interface())
	}
}

// Example 2: Understanding Kind vs Type
func kindVsType() {
	fmt.Println("\n--- Example 2: Kind vs Type ---")

	var examples = []interface{}{
		int(42),
		string("hello"),
		[]int{1, 2, 3},
		map[string]int{"a": 1},
		Person{Name: "Bob"},
		&Person{Name: "Charlie"},
		Calculator(nil),
	}

	for i, example := range examples {
		t := reflect.TypeOf(example)
		if t == nil {
			fmt.Printf("Example %d: nil type\n", i+1)
			continue
		}

		fmt.Printf("Example %d:\n", i+1)
		fmt.Printf("  Type: %v\n", t)
		fmt.Printf("  Kind: %v\n", t.Kind())
		fmt.Printf("  String: %s\n", t.String())

		// Demonstrate the difference
		if t.Kind() == reflect.Ptr {
			fmt.Printf("  Pointer to: %v\n", t.Elem())
			fmt.Printf("  Elem Kind: %v\n", t.Elem().Kind())
		}
		if t.Kind() == reflect.Interface {
			fmt.Printf("  Interface methods: %d\n", t.NumMethod())
		}
		fmt.Println()
	}
}

// Example 3: Detailed type information
func typeInformation() {
	fmt.Println("\n--- Example 3: Type Information ---")

	t := reflect.TypeOf(Person{})
	fmt.Printf("Type: %v\n", t)
	fmt.Printf("Kind: %v\n", t.Kind())
	fmt.Printf("Name: %v\n", t.Name())
	fmt.Printf("PkgPath: %v\n", t.PkgPath())
	fmt.Printf("Size: %d bytes\n", t.Size())
	fmt.Printf("Align: %d bytes\n", t.Align())
	fmt.Printf("Comparable: %v\n", t.Comparable())

	// Check if type implements interfaces
	writerType := reflect.TypeOf((*Writer)(nil)).Elem()
	fmt.Printf("Person implements Writer: %v\n", t.Implements(writerType))

	// Check assignability
	addressType := reflect.TypeOf(Address{})
	fmt.Printf("Person assignable to Address: %v\n", t.AssignableTo(addressType))
	fmt.Printf("Address assignable to Person: %v\n", addressType.AssignableTo(t))

	// Check convertibility
	intType := reflect.TypeOf(int(0))
	int32Type := reflect.TypeOf(int32(0))
	fmt.Printf("int convertible to int32: %v\n", intType.ConvertibleTo(int32Type))
}

// Example 4: Method sets and reflection
func methodSets() {
	fmt.Println("\n--- Example 4: Method Sets ---")

	calc := SimpleCalculator{Name: "MyCalc"}
	t := reflect.TypeOf(calc)
	v := reflect.ValueOf(calc)

	fmt.Printf("Type: %v\n", t)
	fmt.Printf("Number of methods: %d\n", t.NumMethod())

	// Iterate through methods
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("Method %d: %s\n", i, method.Name)
		fmt.Printf("  Type: %v\n", method.Type)
		fmt.Printf("  In: %d, Out: %d\n", method.Type.NumIn(), method.Type.NumOut())

		// Print parameter types
		for j := 0; j < method.Type.NumIn(); j++ {
			fmt.Printf("    In[%d]: %v\n", j, method.Type.In(j))
		}

		// Print return types
		for j := 0; j < method.Type.NumOut(); j++ {
			fmt.Printf("    Out[%d]: %v\n", j, method.Type.Out(j))
		}
	}

	// Get method by name
	if addMethod, found := t.MethodByName("Add"); found {
		fmt.Printf("\nFound Add method: %v\n", addMethod.Type)

		// Call the method using reflection
		methodValue := v.MethodByName("Add")
		if methodValue.IsValid() {
			args := []reflect.Value{
				reflect.ValueOf(10),
				reflect.ValueOf(20),
			}
			results := methodValue.Call(args)
			fmt.Printf("10 + 20 = %v\n", results[0].Interface())
		}
	}
}

// Example 5: Type comparison and conversion
func typeComparison() {
	fmt.Println("\n--- Example 5: Type Comparison ---")

	var examples = []interface{}{
		int(42),
		int32(42),
		float64(42.0),
		"42",
		Person{Name: "test"},
		&Person{Name: "test"},
	}

	intType := reflect.TypeOf(int(0))

	for i, example := range examples {
		t := reflect.TypeOf(example)
		v := reflect.ValueOf(example)

		fmt.Printf("Example %d: %v (Type: %v)\n", i+1, example, t)

		// Type equality
		fmt.Printf("  Equal to int: %v\n", t == intType)

		// Assignability
		fmt.Printf("  Assignable to int: %v\n", t.AssignableTo(intType))

		// Convertibility
		fmt.Printf("  Convertible to int: %v\n", t.ConvertibleTo(intType))

		// Try conversion if possible
		if v.IsValid() && t.ConvertibleTo(intType) {
			converted := v.Convert(intType)
			fmt.Printf("  Converted value: %v\n", converted.Interface())
		}

		fmt.Println()
	}
}

// Example 6: Pointer reflection and indirection
func pointerReflection() {
	fmt.Println("\n--- Example 6: Pointer Reflection ---")

	person := Person{Name: "Alice", Age: 30}
	personPtr := &person

	// Direct value
	vPerson := reflect.ValueOf(person)
	tPerson := reflect.TypeOf(person)
	fmt.Printf("Direct value - Type: %v, Kind: %v\n", tPerson, tPerson.Kind())
	fmt.Printf("  Can addr: %v, Can set: %v\n", vPerson.CanAddr(), vPerson.CanSet())

	// Pointer value
	vPtr := reflect.ValueOf(personPtr)
	tPtr := reflect.TypeOf(personPtr)
	fmt.Printf("Pointer value - Type: %v, Kind: %v\n", tPtr, tPtr.Kind())
	fmt.Printf("  Can addr: %v, Can set: %v\n", vPtr.CanAddr(), vPtr.CanSet())

	// Indirect (dereference)
	if vPtr.Kind() == reflect.Ptr {
		vIndirect := vPtr.Elem()
		tIndirect := tPtr.Elem()
		fmt.Printf("Indirect value - Type: %v, Kind: %v\n", tIndirect, tIndirect.Kind())
		fmt.Printf("  Can addr: %v, Can set: %v\n", vIndirect.CanAddr(), vIndirect.CanSet())

		// Now we can modify fields
		if vIndirect.CanSet() {
			nameField := vIndirect.FieldByName("Name")
			if nameField.IsValid() && nameField.CanSet() {
				oldName := nameField.String()
				nameField.SetString("Modified Alice")
				fmt.Printf("  Changed name from '%s' to '%s'\n", oldName, nameField.String())
			}
		}
	}

	fmt.Printf("Final person value: %+v\n", person)
	fmt.Printf("Final person pointer value: %+v\n", *personPtr)

	// Working with multiple levels of indirection
	doublePtr := &personPtr
	vDoublePtr := reflect.ValueOf(doublePtr)
	fmt.Printf("Double pointer type: %v\n", vDoublePtr.Type())

	// Navigate through pointers
	current := vDoublePtr
	level := 0
	for current.Kind() == reflect.Ptr {
		fmt.Printf("  Level %d: Type %v, Kind %v\n", level, current.Type(), current.Kind())
		current = current.Elem()
		level++
	}
	fmt.Printf("  Final: Type %v, Kind %v\n", current.Type(), current.Kind())
}
