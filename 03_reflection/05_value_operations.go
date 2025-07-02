package main

import (
	"fmt"
	"reflect"
)

// ValueOperationsExamples demonstrates value operations and creation
func ValueOperationsExamples() {
	fmt.Println("\n=== Value Operations & Creation Examples ===")

	// Example 1: Basic value operations
	basicValueOperations()

	// Example 2: Dynamic value creation
	dynamicValueCreation()

	// Example 3: Value modification and setting
	valueModification()

	// Example 4: Type conversion operations
	typeConversion()

	// Example 5: Slice and map operations
	sliceMapOperations()

	// Example 6: Zero value handling
	zeroValueHandling()
}

// Example 1: Basic value operations
func basicValueOperations() {
	fmt.Println("\n--- Example 1: Basic Value Operations ---")

	// Various value types
	values := []interface{}{
		42,
		3.14,
		"hello",
		true,
		[]int{1, 2, 3},
		map[string]int{"key": 42},
		Person{Name: "Alice", Age: 30},
	}

	for i, val := range values {
		fmt.Printf("\nValue %d: %v\n", i+1, val)

		v := reflect.ValueOf(val)
		analyzeValue(v)
	}
}

// Example 2: Dynamic value creation
func dynamicValueCreation() {
	fmt.Println("\n--- Example 2: Dynamic Value Creation ---")

	// Create values of different types dynamically
	types := []reflect.Type{
		reflect.TypeOf(int(0)),
		reflect.TypeOf(string("")),
		reflect.TypeOf(Person{}),
		reflect.TypeOf(&Person{}),
		reflect.TypeOf([]int{}),
		reflect.TypeOf(map[string]int{}),
		reflect.TypeOf(make(chan int)),
	}

	for _, t := range types {
		fmt.Printf("\nCreating value of type: %v\n", t)
		createAndInitializeValue(t)
	}
}

// Example 3: Value modification and setting
func valueModification() {
	fmt.Println("\n--- Example 3: Value Modification ---")

	// Create a modifiable struct
	person := &Person{
		Name:  "John",
		Age:   25,
		Email: "john@example.com",
	}

	fmt.Printf("Original person: %+v\n", *person)

	v := reflect.ValueOf(person).Elem()
	modifyStructFields(v)

	fmt.Printf("Modified person: %+v\n", *person)

	// Modify slice elements
	fmt.Printf("\nSlice modification:\n")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original slice: %v\n", numbers)

	sliceValue := reflect.ValueOf(numbers)
	modifySliceElements(sliceValue)
	fmt.Printf("Modified slice: %v\n", numbers)

	// Modify map values
	fmt.Printf("\nMap modification:\n")
	data := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("Original map: %v\n", data)

	mapValue := reflect.ValueOf(data)
	modifyMapValues(mapValue)
	fmt.Printf("Modified map: %v\n", data)
}

// Example 4: Type conversion operations
func typeConversion() {
	fmt.Println("\n--- Example 4: Type Conversion ---")

	// Source values of different types
	sourceValues := []interface{}{
		42,
		3.14,
		"123",
		true,
		int32(42),
		float32(3.14),
	}

	// Target types to convert to
	targetTypes := []reflect.Type{
		reflect.TypeOf(int(0)),
		reflect.TypeOf(int64(0)),
		reflect.TypeOf(float64(0)),
		reflect.TypeOf(string("")),
		reflect.TypeOf(bool(false)),
	}

	for _, source := range sourceValues {
		sourceValue := reflect.ValueOf(source)
		sourceType := sourceValue.Type()

		fmt.Printf("\nSource: %v (type: %v)\n", source, sourceType)

		for _, targetType := range targetTypes {
			convertValue(sourceValue, targetType)
		}
	}
}

// Example 5: Slice and map operations
func sliceMapOperations() {
	fmt.Println("\n--- Example 5: Slice and Map Operations ---")

	// Dynamic slice creation and manipulation
	fmt.Printf("Dynamic slice operations:\n")
	performSliceOperations()

	// Dynamic map creation and manipulation
	fmt.Printf("\nDynamic map operations:\n")
	performMapOperations()

	// Complex nested structure manipulation
	fmt.Printf("\nNested structure operations:\n")
	performNestedOperations()
}

// Example 6: Zero value handling
func zeroValueHandling() {
	fmt.Println("\n--- Example 6: Zero Value Handling ---")

	// Test zero values for different types
	types := []reflect.Type{
		reflect.TypeOf(int(0)),
		reflect.TypeOf(string("")),
		reflect.TypeOf(bool(false)),
		reflect.TypeOf([]int{}),
		reflect.TypeOf(map[string]int{}),
		reflect.TypeOf(Person{}),
		reflect.TypeOf(&Person{}),
		reflect.TypeOf(make(chan int)),
	}

	for _, t := range types {
		fmt.Printf("\nType: %v\n", t)
		demonstrateZeroValue(t)
	}
}

// Helper functions

func analyzeValue(v reflect.Value) {
	fmt.Printf("  Type: %v\n", v.Type())
	fmt.Printf("  Kind: %v\n", v.Kind())
	fmt.Printf("  Is valid: %v\n", v.IsValid())

	if v.IsValid() {
		fmt.Printf("  Is zero: %v\n", v.IsZero())
		fmt.Printf("  Can interface: %v\n", v.CanInterface())
		fmt.Printf("  Can addr: %v\n", v.CanAddr())
		fmt.Printf("  Can set: %v\n", v.CanSet())

		// Check for nil (if applicable)
		switch v.Kind() {
		case reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func, reflect.Slice, reflect.Map:
			fmt.Printf("  Is nil: %v\n", v.IsNil())
		}

		// Type-specific information
		switch v.Kind() {
		case reflect.Slice, reflect.Array:
			fmt.Printf("  Length: %d\n", v.Len())
			if v.Kind() == reflect.Slice {
				fmt.Printf("  Capacity: %d\n", v.Cap())
			}
		case reflect.Map:
			fmt.Printf("  Length: %d\n", v.Len())
		case reflect.String:
			fmt.Printf("  Length: %d\n", v.Len())
		case reflect.Struct:
			fmt.Printf("  Number of fields: %d\n", v.NumField())
		}
	}
}

func createAndInitializeValue(t reflect.Type) {
	var v reflect.Value

	switch t.Kind() {
	case reflect.Ptr:
		// Create pointer to new value
		v = reflect.New(t.Elem())
		fmt.Printf("  Created pointer: %v\n", v.Type())

		// Initialize the pointed-to value
		elem := v.Elem()
		if elem.CanSet() {
			initializeValue(elem)
			fmt.Printf("  Initialized value: %v\n", v.Interface())
		}

	case reflect.Slice:
		// Create slice with some capacity
		v = reflect.MakeSlice(t, 0, 3)
		fmt.Printf("  Created slice: %v (len=%d, cap=%d)\n", v.Type(), v.Len(), v.Cap())

		// Add some elements if it's a slice of basic types
		if t.Elem().Kind() == reflect.Int {
			for i := 0; i < 3; i++ {
				v = reflect.Append(v, reflect.ValueOf(i+1))
			}
			fmt.Printf("  With elements: %v\n", v.Interface())
		}

	case reflect.Map:
		// Create map
		v = reflect.MakeMap(t)
		fmt.Printf("  Created map: %v\n", v.Type())

		// Add some entries if it's a map[string]int
		if t.Key().Kind() == reflect.String && t.Elem().Kind() == reflect.Int {
			v.SetMapIndex(reflect.ValueOf("key1"), reflect.ValueOf(10))
			v.SetMapIndex(reflect.ValueOf("key2"), reflect.ValueOf(20))
			fmt.Printf("  With entries: %v\n", v.Interface())
		}

	case reflect.Chan:
		// Create channel
		v = reflect.MakeChan(t, 2)
		fmt.Printf("  Created channel: %v\n", v.Type())

	default:
		// Create zero value
		v = reflect.Zero(t)
		fmt.Printf("  Created zero value: %v\n", v.Interface())

		// Try to initialize if it's a struct
		if t.Kind() == reflect.Struct {
			initializeStruct(v)
		}
	}
}

func initializeValue(v reflect.Value) {
	switch v.Kind() {
	case reflect.Int:
		v.SetInt(42)
	case reflect.String:
		v.SetString("initialized")
	case reflect.Bool:
		v.SetBool(true)
	case reflect.Float64:
		v.SetFloat(3.14)
	case reflect.Struct:
		initializeStruct(v)
	}
}

func initializeStruct(v reflect.Value) {
	if !v.CanSet() {
		return
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		if !field.CanSet() {
			continue
		}

		fieldType := field.Type()
		switch fieldType.Kind() {
		case reflect.String:
			field.SetString(fmt.Sprintf("field_%d", i))
		case reflect.Int:
			field.SetInt(int64(i + 1))
		case reflect.Bool:
			field.SetBool(i%2 == 0)
		}
	}
}

func modifyStructFields(v reflect.Value) {
	t := v.Type()

	fmt.Printf("Modifying struct fields:\n")
	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if !field.CanSet() {
			fmt.Printf("  Field %s: cannot set\n", fieldType.Name)
			continue
		}

		fmt.Printf("  Field %s (%v): ", fieldType.Name, field.Type())

		switch field.Kind() {
		case reflect.String:
			oldVal := field.String()
			newVal := "Modified_" + oldVal
			field.SetString(newVal)
			fmt.Printf("%s -> %s\n", oldVal, newVal)

		case reflect.Int:
			oldVal := field.Int()
			newVal := oldVal + 10
			field.SetInt(newVal)
			fmt.Printf("%d -> %d\n", oldVal, newVal)

		default:
			fmt.Printf("unsupported type\n")
		}
	}
}

func modifySliceElements(v reflect.Value) {
	if v.Kind() != reflect.Slice {
		fmt.Printf("Not a slice\n")
		return
	}

	fmt.Printf("Modifying slice elements (if addressable):\n")
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		if elem.CanSet() && elem.Kind() == reflect.Int {
			oldVal := elem.Int()
			newVal := oldVal * 10
			elem.SetInt(newVal)
			fmt.Printf("  [%d]: %d -> %d\n", i, oldVal, newVal)
		}
	}
}

func modifyMapValues(v reflect.Value) {
	if v.Kind() != reflect.Map {
		fmt.Printf("Not a map\n")
		return
	}

	fmt.Printf("Modifying map values:\n")
	keys := v.MapKeys()
	for _, key := range keys {
		val := v.MapIndex(key)
		if val.Kind() == reflect.Int {
			oldVal := val.Int()
			newVal := oldVal * 100
			v.SetMapIndex(key, reflect.ValueOf(int(newVal)))
			fmt.Printf("  %v: %d -> %d\n", key.Interface(), oldVal, newVal)
		}
	}
}

func convertValue(source reflect.Value, targetType reflect.Type) {
	sourceType := source.Type()

	fmt.Printf("  %v -> %v: ", sourceType, targetType)

	if sourceType == targetType {
		fmt.Printf("same type\n")
		return
	}

	if !sourceType.ConvertibleTo(targetType) {
		fmt.Printf("not convertible\n")
		return
	}

	converted := source.Convert(targetType)
	fmt.Printf("%v -> %v\n", source.Interface(), converted.Interface())
}

func performSliceOperations() {
	// Create slice dynamically
	intSliceType := reflect.TypeOf([]int{})
	slice := reflect.MakeSlice(intSliceType, 0, 5)

	fmt.Printf("  Created slice: %v (len=%d, cap=%d)\n", slice.Interface(), slice.Len(), slice.Cap())

	// Append elements
	for i := 1; i <= 3; i++ {
		slice = reflect.Append(slice, reflect.ValueOf(i*10))
	}
	fmt.Printf("  After appending: %v (len=%d)\n", slice.Interface(), slice.Len())

	// Modify elements
	for i := 0; i < slice.Len(); i++ {
		elem := slice.Index(i)
		elem.SetInt(elem.Int() + 1)
	}
	fmt.Printf("  After modifying: %v\n", slice.Interface())

	// Create slice of structs
	personSliceType := reflect.TypeOf([]Person{})
	personSlice := reflect.MakeSlice(personSliceType, 0, 2)

	person1 := reflect.ValueOf(Person{Name: "Alice", Age: 30})
	person2 := reflect.ValueOf(Person{Name: "Bob", Age: 25})

	personSlice = reflect.Append(personSlice, person1, person2)
	fmt.Printf("  Person slice: %v\n", personSlice.Interface())
}

func performMapOperations() {
	// Create map dynamically
	mapType := reflect.TypeOf(map[string]int{})
	m := reflect.MakeMap(mapType)

	fmt.Printf("  Created map: %v\n", m.Interface())

	// Set values
	keys := []string{"one", "two", "three"}
	for i, key := range keys {
		m.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf((i+1)*10))
	}
	fmt.Printf("  After setting values: %v\n", m.Interface())

	// Iterate through map
	fmt.Printf("  Iterating through map:\n")
	iter := m.MapRange()
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("    %v: %v\n", key.Interface(), value.Interface())
	}

	// Delete a key
	m.SetMapIndex(reflect.ValueOf("two"), reflect.Value{})
	fmt.Printf("  After deleting 'two': %v\n", m.Interface())
}

func performNestedOperations() {
	// Create a nested structure: map[string][]Person
	mapType := reflect.TypeOf(map[string][]Person{})
	m := reflect.MakeMap(mapType)

	// Create slice of persons
	personSliceType := reflect.TypeOf([]Person{})
	slice1 := reflect.MakeSlice(personSliceType, 0, 2)
	slice1 = reflect.Append(slice1, reflect.ValueOf(Person{Name: "Alice", Age: 30}))
	slice1 = reflect.Append(slice1, reflect.ValueOf(Person{Name: "Bob", Age: 25}))

	slice2 := reflect.MakeSlice(personSliceType, 0, 1)
	slice2 = reflect.Append(slice2, reflect.ValueOf(Person{Name: "Charlie", Age: 35}))

	// Set map entries
	m.SetMapIndex(reflect.ValueOf("team1"), slice1)
	m.SetMapIndex(reflect.ValueOf("team2"), slice2)

	fmt.Printf("  Created nested structure: %v\n", m.Interface())

	// Access nested values
	team1 := m.MapIndex(reflect.ValueOf("team1"))
	if team1.IsValid() && team1.Kind() == reflect.Slice {
		fmt.Printf("  Team1 has %d members\n", team1.Len())
		for i := 0; i < team1.Len(); i++ {
			person := team1.Index(i)
			nameField := person.FieldByName("Name")
			if nameField.IsValid() {
				fmt.Printf("    Member %d: %s\n", i+1, nameField.String())
			}
		}
	}
}

func demonstrateZeroValue(t reflect.Type) {
	zero := reflect.Zero(t)

	fmt.Printf("  Zero value: %v\n", zero.Interface())
	fmt.Printf("  Is zero: %v\n", zero.IsZero())
	fmt.Printf("  Type: %v\n", zero.Type())

	// Check if zero value is nil for pointer types
	switch t.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func, reflect.Slice, reflect.Map:
		fmt.Printf("  Is nil: %v\n", zero.IsNil())
	}

	// Create non-zero value for comparison
	var nonZero reflect.Value
	switch t.Kind() {
	case reflect.Int:
		nonZero = reflect.ValueOf(42)
	case reflect.String:
		nonZero = reflect.ValueOf("hello")
	case reflect.Bool:
		nonZero = reflect.ValueOf(true)
	case reflect.Slice:
		nonZero = reflect.MakeSlice(t, 1, 1)
	case reflect.Map:
		nonZero = reflect.MakeMap(t)
		// Add an entry to make it non-zero
		if t.Key().Kind() == reflect.String && t.Elem().Kind() == reflect.Int {
			nonZero.SetMapIndex(reflect.ValueOf("key"), reflect.ValueOf(1))
		}
	case reflect.Ptr:
		nonZero = reflect.New(t.Elem())
	case reflect.Struct:
		nonZero = reflect.New(t).Elem()
		// Initialize a field to make it non-zero
		for i := 0; i < nonZero.NumField(); i++ {
			field := nonZero.Field(i)
			if field.CanSet() {
				switch field.Kind() {
				case reflect.String:
					field.SetString("test")
				case reflect.Int:
					field.SetInt(1)
				case reflect.Bool:
					field.SetBool(true)
				}
				break
			}
		}
	}

	if nonZero.IsValid() {
		fmt.Printf("  Non-zero value: %v\n", nonZero.Interface())
		fmt.Printf("  Non-zero is zero: %v\n", nonZero.IsZero())
	}
}
