package main

import (
	"fmt"
	"reflect"
	"strings"
)

// StructReflectionExamples demonstrates struct reflection and tag parsing
func StructReflectionExamples() {
	fmt.Println("\n=== Struct Reflection & Tags Examples ===")

	// Example 1: Basic struct field reflection
	basicStructReflection()

	// Example 2: Struct tag parsing and validation
	structTagParsing()

	// Example 3: Embedded struct handling
	embeddedStructs()

	// Example 4: Dynamic struct field modification
	dynamicFieldModification()

	// Example 5: Struct tag-based JSON marshaling simulation
	tagBasedProcessing()

	// Example 6: Struct field iteration patterns
	structFieldIteration()
}

// Example 1: Basic struct field reflection
func basicStructReflection() {
	fmt.Println("\n--- Example 1: Basic Struct Field Reflection ---")

	person := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
	}

	v := reflect.ValueOf(person)
	t := reflect.TypeOf(person)

	fmt.Printf("Struct: %s\n", t.Name())
	fmt.Printf("Number of fields: %d\n", t.NumField())

	// Iterate through all fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fmt.Printf("Field %d:\n", i)
		fmt.Printf("  Name: %s\n", field.Name)
		fmt.Printf("  Type: %v\n", field.Type)
		fmt.Printf("  Kind: %v\n", field.Type.Kind())
		fmt.Printf("  Tag: %q\n", field.Tag)
		fmt.Printf("  Value: %v\n", value.Interface())
		fmt.Printf("  CanSet: %v\n", value.CanSet())
		fmt.Printf("  IsExported: %v\n", field.PkgPath == "")
		fmt.Println()
	}

	// Access specific field by name
	if nameField, found := t.FieldByName("Name"); found {
		fmt.Printf("Found Name field: %v\n", nameField.Type)
		nameValue := v.FieldByName("Name")
		fmt.Printf("Name value: %v\n", nameValue.String())
	}
}

// Example 2: Struct tag parsing and validation
func structTagParsing() {
	fmt.Println("\n--- Example 2: Struct Tag Parsing ---")

	user := User{
		ID:       1,
		Username: "john_doe",
		Password: "secret123",
	}

	t := reflect.TypeOf(user)

	fmt.Printf("Analyzing struct tags for: %s\n", t.Name())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		fmt.Printf("\nField: %s\n", field.Name)

		// Parse different tag types
		if jsonTag := field.Tag.Get("json"); jsonTag != "" {
			fmt.Printf("  JSON tag: %s\n", jsonTag)

			// Parse JSON tag options
			parts := strings.Split(jsonTag, ",")
			jsonName := parts[0]
			fmt.Printf("  JSON name: %s\n", jsonName)

			if len(parts) > 1 {
				fmt.Printf("  JSON options: %v\n", parts[1:])
			}
		}

		if dbTag := field.Tag.Get("db"); dbTag != "" {
			fmt.Printf("  DB tag: %s\n", dbTag)
		}

		if validateTag := field.Tag.Get("validate"); validateTag != "" {
			fmt.Printf("  Validate tag: %s\n", validateTag)
			// Parse validation rules
			rules := strings.Split(validateTag, ",")
			for _, rule := range rules {
				if strings.Contains(rule, "=") {
					parts := strings.Split(rule, "=")
					fmt.Printf("    Rule: %s = %s\n", parts[0], parts[1])
				} else {
					fmt.Printf("    Rule: %s\n", rule)
				}
			}
		}

		// Get the full tag
		fmt.Printf("  Full tag: %q\n", string(field.Tag))
	}
}

// Example 3: Embedded struct handling
func embeddedStructs() {
	fmt.Println("\n--- Example 3: Embedded Structs ---")

	employee := Employee{
		Person: Person{
			Name:  "Bob",
			Age:   25,
			Email: "bob@company.com",
		},
		ID:       100,
		Position: "Developer",
		Salary:   75000.0,
	}

	v := reflect.ValueOf(employee)
	t := reflect.TypeOf(employee)

	fmt.Printf("Employee struct analysis:\n")
	fmt.Printf("Number of fields: %d\n", t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fmt.Printf("\nField %d: %s\n", i, field.Name)
		fmt.Printf("  Type: %v\n", field.Type)
		fmt.Printf("  Anonymous: %v\n", field.Anonymous)
		fmt.Printf("  Value: %v\n", value.Interface())

		// If it's an embedded struct, explore its fields
		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			fmt.Printf("  Embedded struct fields:\n")
			embeddedType := field.Type
			embeddedValue := value

			for j := 0; j < embeddedType.NumField(); j++ {
				embeddedField := embeddedType.Field(j)
				embeddedFieldValue := embeddedValue.Field(j)
				fmt.Printf("    %s: %v\n", embeddedField.Name, embeddedFieldValue.Interface())
			}
		}
	}

	// Access embedded fields directly
	fmt.Printf("\nDirect access to embedded fields:\n")
	if nameField := v.FieldByName("Name"); nameField.IsValid() {
		fmt.Printf("Name (from embedded Person): %v\n", nameField.String())
	}
}

// Example 4: Dynamic field modification
func dynamicFieldModification() {
	fmt.Println("\n--- Example 4: Dynamic Field Modification ---")

	person := &Person{
		Name:  "Charlie",
		Age:   35,
		Email: "charlie@example.com",
	}

	v := reflect.ValueOf(person).Elem() // Get the pointed-to value
	t := reflect.TypeOf(*person)        // Get the type of the pointed-to value

	fmt.Printf("Original person: %+v\n", *person)
	fmt.Printf("Type of person: %v\n", t)

	// Modify fields by name
	modifications := map[string]interface{}{
		"Name":  "Charles",
		"Age":   36,
		"Email": "charles@example.com",
	}

	for fieldName, newValue := range modifications {
		field := v.FieldByName(fieldName)
		if !field.IsValid() {
			fmt.Printf("Field %s not found\n", fieldName)
			continue
		}

		if !field.CanSet() {
			fmt.Printf("Field %s cannot be set\n", fieldName)
			continue
		}

		// Convert the value to the correct type
		newVal := reflect.ValueOf(newValue)
		if field.Type() == newVal.Type() {
			field.Set(newVal)
			fmt.Printf("Set %s to %v\n", fieldName, newValue)
		} else if newVal.Type().ConvertibleTo(field.Type()) {
			converted := newVal.Convert(field.Type())
			field.Set(converted)
			fmt.Printf("Set %s to %v (converted from %T)\n", fieldName, converted.Interface(), newValue)
		} else {
			fmt.Printf("Cannot convert %T to %v for field %s\n", newValue, field.Type(), fieldName)
		}
	}

	fmt.Printf("Modified person: %+v\n", *person)

	// Demonstrate field type-specific setters
	fmt.Printf("\nUsing type-specific setters:\n")
	ageField := v.FieldByName("Age")
	if ageField.IsValid() && ageField.CanSet() && ageField.Kind() == reflect.Int {
		ageField.SetInt(40)
		fmt.Printf("Set age using SetInt: %d\n", person.Age)
	}

	nameField := v.FieldByName("Name")
	if nameField.IsValid() && nameField.CanSet() && nameField.Kind() == reflect.String {
		nameField.SetString("Chuck")
		fmt.Printf("Set name using SetString: %s\n", person.Name)
	}
}

// Example 5: Tag-based JSON marshaling simulation
func tagBasedProcessing() {
	fmt.Println("\n--- Example 5: Tag-based Processing ---")

	user := User{
		ID:       1,
		Username: "jane_doe",
		Password: "supersecret",
	}

	v := reflect.ValueOf(user)
	t := reflect.TypeOf(user)

	// Simulate JSON marshaling based on tags
	jsonData := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			continue
		}

		// Parse JSON tag
		parts := strings.Split(jsonTag, ",")
		jsonName := parts[0]

		// Skip fields marked with "-"
		if jsonName == "-" {
			fmt.Printf("Skipping field %s (marked with json:\"-\")\n", field.Name)
			continue
		}

		// Use field name if JSON name is empty
		if jsonName == "" {
			jsonName = field.Name
		}

		// Check for omitempty option
		omitEmpty := false
		for _, option := range parts[1:] {
			if option == "omitempty" {
				omitEmpty = true
				break
			}
		}

		// Add to JSON data if not omitting empty values
		fieldValue := value.Interface()
		if !omitEmpty || !isZeroValue(value) {
			jsonData[jsonName] = fieldValue
			fmt.Printf("Added field %s as '%s': %v\n", field.Name, jsonName, fieldValue)
		} else {
			fmt.Printf("Omitted empty field %s\n", field.Name)
		}
	}

	fmt.Printf("\nSimulated JSON data: %+v\n", jsonData)
}

// Example 6: Struct field iteration patterns
func structFieldIteration() {
	fmt.Println("\n--- Example 6: Struct Field Iteration Patterns ---")

	config := Config{
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			Username: "admin",
			Password: "secret",
			Database: "myapp",
		},
		Server: ServerConfig{
			Host: "0.0.0.0",
			Port: 8080,
		},
		Redis: RedisConfig{
			Host: "localhost",
			Port: 6379,
		},
	}

	// Pattern 1: Recursive struct traversal
	fmt.Printf("Pattern 1: Recursive struct traversal\n")
	traverseStruct(reflect.ValueOf(config), 0)

	// Pattern 2: Flat field collection
	fmt.Printf("\nPattern 2: Flat field collection\n")
	fields := collectAllFields(reflect.ValueOf(config), "")
	for path, value := range fields {
		fmt.Printf("  %s: %v\n", path, value)
	}
}

// Helper function to check if a value is zero
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.String:
		return v.String() == ""
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return v.IsNil()
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isZeroValue(v.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !isZeroValue(v.Field(i)) {
				return false
			}
		}
		return true
	}
	return false
}

// Helper function for recursive struct traversal
func traverseStruct(v reflect.Value, depth int) {
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return
	}

	t := v.Type()
	indent := strings.Repeat("  ", depth)

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fmt.Printf("%s%s: %v (type: %v)\n", indent, field.Name, value.Interface(), field.Type)

		// Recursively traverse nested structs
		if value.Kind() == reflect.Struct {
			traverseStruct(value, depth+1)
		} else if value.Kind() == reflect.Ptr && value.Type().Elem().Kind() == reflect.Struct {
			traverseStruct(value, depth+1)
		}
	}
}

// Helper function to collect all fields in a flat structure
func collectAllFields(v reflect.Value, prefix string) map[string]interface{} {
	result := make(map[string]interface{})

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return result
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return result
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fieldPath := field.Name
		if prefix != "" {
			fieldPath = prefix + "." + field.Name
		}

		if value.Kind() == reflect.Struct {
			// Recursively collect nested struct fields
			nested := collectAllFields(value, fieldPath)
			for k, v := range nested {
				result[k] = v
			}
		} else {
			result[fieldPath] = value.Interface()
		}
	}

	return result
}
