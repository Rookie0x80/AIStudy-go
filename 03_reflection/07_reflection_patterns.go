package main

import (
	"fmt"
	"reflect"
	"strings"
)

// ReflectionPatternsExamples demonstrates reflection design patterns
func ReflectionPatternsExamples() {
	fmt.Println("\n=== Reflection Design Patterns Examples ===")

	// Example 1: Object Mapper Pattern
	objectMapperPattern()

	// Example 2: Dependency Injection Container
	dependencyInjectionPattern()

	// Example 3: Struct Validation Framework
	validationFrameworkPattern()

	// Example 4: Generic Serialization Framework
	serializationFrameworkPattern()

	// Example 5: Dynamic Configuration Binding
	configurationBindingPattern()

	// Example 6: Test Data Builder Pattern
	testDataBuilderPattern()
}

// Example 1: Object Mapper Pattern
func objectMapperPattern() {
	fmt.Println("\n--- Example 1: Object Mapper Pattern ---")

	mapper := NewObjectMapper()

	// Source and destination structs
	source := Person{
		Name:  "Alice Smith",
		Age:   30,
		Email: "alice@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
	}

	// Map to different structure
	type PersonDTO struct {
		FullName    string `map:"Name"`
		Age         int    `map:"Age"`
		EmailAddr   string `map:"Email"`
		StreetAddr  string `map:"Address.Street"`
		CityName    string `map:"Address.City"`
		CountryCode string `map:"Address.Country"`
	}

	var dto PersonDTO
	if err := mapper.Map(source, &dto); err != nil {
		fmt.Printf("Mapping error: %v\n", err)
		return
	}

	fmt.Printf("Source: %+v\n", source)
	fmt.Printf("Mapped DTO: %+v\n", dto)

	// Map back
	var person2 Person
	if err := mapper.Map(dto, &person2); err != nil {
		fmt.Printf("Reverse mapping error: %v\n", err)
		return
	}
	fmt.Printf("Reverse mapped: %+v\n", person2)
}

// Example 2: Dependency Injection Container
func dependencyInjectionPattern() {
	fmt.Println("\n--- Example 2: Dependency Injection Container ---")

	container := NewDIContainer()

	// Register services
	container.RegisterSingleton("database", &mockDatabase{})
	container.RegisterSingleton("logger", &mockLogger{})
	container.RegisterTransient("userService", func() interface{} {
		return &UserServiceImpl{}
	})

	// Create service with dependencies
	userService, err := container.Resolve("userService")
	if err != nil {
		fmt.Printf("Error resolving service: %v\n", err)
		return
	}

	fmt.Printf("Resolved service: %T\n", userService)

	// Inject dependencies
	if err := container.InjectDependencies(userService); err != nil {
		fmt.Printf("Error injecting dependencies: %v\n", err)
		return
	}

	fmt.Printf("Dependencies injected successfully\n")
}

// Example 3: Struct Validation Framework
func validationFrameworkPattern() {
	fmt.Println("\n--- Example 3: Validation Framework ---")

	validator := NewValidator()

	// Test valid user
	validUser := User{
		ID:       1,
		Username: "john_doe",
		Password: "securepass123",
	}

	if err := validator.Validate(validUser); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Printf("Valid user: %+v\n", validUser)
	}

	// Test invalid user
	invalidUser := User{
		ID:       0,
		Username: "jo",  // Too short
		Password: "123", // Too short
	}

	if err := validator.Validate(invalidUser); err != nil {
		fmt.Printf("Validation errors: %v\n", err)
	}
}

// Example 4: Generic Serialization Framework
func serializationFrameworkPattern() {
	fmt.Println("\n--- Example 4: Serialization Framework ---")

	serializer := NewGenericSerializer()

	person := Person{
		Name:  "Bob Johnson",
		Age:   25,
		Email: "bob@example.com",
	}

	// Serialize to map
	data, err := serializer.Serialize(person)
	if err != nil {
		fmt.Printf("Serialization error: %v\n", err)
		return
	}

	fmt.Printf("Serialized data: %v\n", data)

	// Deserialize back
	var person2 Person
	if err := serializer.Deserialize(data, &person2); err != nil {
		fmt.Printf("Deserialization error: %v\n", err)
		return
	}

	fmt.Printf("Deserialized: %+v\n", person2)
}

// Example 5: Dynamic Configuration Binding
func configurationBindingPattern() {
	fmt.Println("\n--- Example 5: Configuration Binding ---")

	binder := NewConfigBinder()

	// Sample configuration data
	configData := map[string]interface{}{
		"database.host":     "localhost",
		"database.port":     5432,
		"database.username": "admin",
		"database.password": "secret",
		"server.host":       "0.0.0.0",
		"server.port":       8080,
		"redis.host":        "redis-server",
		"redis.port":        6379,
	}

	var config Config
	if err := binder.Bind(configData, &config); err != nil {
		fmt.Printf("Binding error: %v\n", err)
		return
	}

	fmt.Printf("Bound configuration: %+v\n", config)
}

// Example 6: Test Data Builder Pattern
func testDataBuilderPattern() {
	fmt.Println("\n--- Example 6: Test Data Builder ---")

	builder := NewTestDataBuilder()

	// Build test person
	person := builder.Build(Person{}).(*Person)
	fmt.Printf("Generated person: %+v\n", *person)

	// Build test user
	user := builder.Build(User{}).(*User)
	fmt.Printf("Generated user: %+v\n", *user)

	// Build custom data
	customPerson := builder.BuildWithRules(Person{}, map[string]interface{}{
		"Name": "Custom Name",
		"Age":  99,
	}).(*Person)
	fmt.Printf("Custom person: %+v\n", *customPerson)
}

// Object Mapper Implementation
type ObjectMapper struct {
	mappings map[string]string
}

func NewObjectMapper() *ObjectMapper {
	return &ObjectMapper{
		mappings: make(map[string]string),
	}
}

func (om *ObjectMapper) Map(source, destination interface{}) error {
	srcVal := reflect.ValueOf(source)
	destVal := reflect.ValueOf(destination)

	if destVal.Kind() != reflect.Ptr {
		return fmt.Errorf("destination must be a pointer")
	}

	destVal = destVal.Elem()
	return om.mapValues(srcVal, destVal)
}

func (om *ObjectMapper) mapValues(src, dest reflect.Value) error {
	destType := dest.Type()

	for i := 0; i < destType.NumField(); i++ {
		field := destType.Field(i)
		destField := dest.Field(i)

		if !destField.CanSet() {
			continue
		}

		// Check for mapping tag
		mapTag := field.Tag.Get("map")
		if mapTag == "" {
			mapTag = field.Name
		}

		// Get source value
		srcValue, err := om.getNestedValue(src, mapTag)
		if err != nil {
			continue // Skip if source field not found
		}

		// Set destination field
		if err := om.setValue(destField, srcValue); err != nil {
			return fmt.Errorf("error setting field %s: %v", field.Name, err)
		}
	}

	return nil
}

func (om *ObjectMapper) getNestedValue(val reflect.Value, path string) (reflect.Value, error) {
	parts := strings.Split(path, ".")
	current := val

	for _, part := range parts {
		if current.Kind() == reflect.Ptr {
			current = current.Elem()
		}

		if current.Kind() != reflect.Struct {
			return reflect.Value{}, fmt.Errorf("not a struct")
		}

		current = current.FieldByName(part)
		if !current.IsValid() {
			return reflect.Value{}, fmt.Errorf("field %s not found", part)
		}
	}

	return current, nil
}

func (om *ObjectMapper) setValue(dest, src reflect.Value) error {
	if !dest.CanSet() {
		return fmt.Errorf("destination field cannot be set")
	}

	if src.Type().AssignableTo(dest.Type()) {
		dest.Set(src)
		return nil
	}

	if src.Type().ConvertibleTo(dest.Type()) {
		dest.Set(src.Convert(dest.Type()))
		return nil
	}

	return fmt.Errorf("cannot assign %v to %v", src.Type(), dest.Type())
}

// Dependency Injection Container
type DIContainer struct {
	singletons map[string]interface{}
	factories  map[string]func() interface{}
}

func NewDIContainer() *DIContainer {
	return &DIContainer{
		singletons: make(map[string]interface{}),
		factories:  make(map[string]func() interface{}),
	}
}

func (di *DIContainer) RegisterSingleton(name string, instance interface{}) {
	di.singletons[name] = instance
}

func (di *DIContainer) RegisterTransient(name string, factory func() interface{}) {
	di.factories[name] = factory
}

func (di *DIContainer) Resolve(name string) (interface{}, error) {
	if instance, exists := di.singletons[name]; exists {
		return instance, nil
	}

	if factory, exists := di.factories[name]; exists {
		return factory(), nil
	}

	return nil, fmt.Errorf("service %s not registered", name)
}

func (di *DIContainer) InjectDependencies(obj interface{}) error {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		if !fieldVal.CanSet() {
			continue
		}

		// Check for inject tag
		injectTag := field.Tag.Get("inject")
		if injectTag == "" {
			continue
		}

		dependency, err := di.Resolve(injectTag)
		if err != nil {
			return fmt.Errorf("failed to resolve dependency %s: %v", injectTag, err)
		}

		depVal := reflect.ValueOf(dependency)
		if depVal.Type().AssignableTo(fieldVal.Type()) {
			fieldVal.Set(depVal)
		}
	}

	return nil
}

// Mock services for DI example
type mockDatabase struct{}
type mockLogger struct{}

// Validation Framework
type ValidatorFramework struct {
	rules map[string]func(reflect.Value) error
}

func NewValidator() *ValidatorFramework {
	v := &ValidatorFramework{
		rules: make(map[string]func(reflect.Value) error),
	}

	// Register built-in rules
	v.rules["required"] = func(val reflect.Value) error {
		if val.IsZero() {
			return fmt.Errorf("field is required")
		}
		return nil
	}

	v.rules["min"] = func(val reflect.Value) error {
		// This is a simplified implementation
		return nil
	}

	return v
}

func (vf *ValidatorFramework) Validate(obj interface{}) error {
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	var errors []string

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		validateTag := field.Tag.Get("validate")
		if validateTag == "" {
			continue
		}

		rules := strings.Split(validateTag, ",")
		for _, rule := range rules {
			rule = strings.TrimSpace(rule)

			// Parse rule (e.g., "min=3")
			parts := strings.Split(rule, "=")
			ruleName := parts[0]

			if validator, exists := vf.rules[ruleName]; exists {
				if err := validator(fieldVal); err != nil {
					errors = append(errors, fmt.Sprintf("%s: %v", field.Name, err))
				}
			}

			// Handle specific rules
			if ruleName == "min" && len(parts) > 1 {
				if err := vf.validateMin(fieldVal, parts[1]); err != nil {
					errors = append(errors, fmt.Sprintf("%s: %v", field.Name, err))
				}
			}
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("validation errors: %s", strings.Join(errors, ", "))
	}

	return nil
}

func (vf *ValidatorFramework) validateMin(val reflect.Value, minStr string) error {
	switch val.Kind() {
	case reflect.String:
		if len(val.String()) < 3 { // Simplified - should parse minStr
			return fmt.Errorf("string too short")
		}
	case reflect.Int:
		if val.Int() < 1 { // Simplified
			return fmt.Errorf("value too small")
		}
	}
	return nil
}

// Generic Serialization Framework
type GenericSerializer struct{}

func NewGenericSerializer() *GenericSerializer {
	return &GenericSerializer{}
}

func (gs *GenericSerializer) Serialize(obj interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("only structs are supported")
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		jsonTag := field.Tag.Get("json")
		if jsonTag == "-" {
			continue
		}

		key := field.Name
		if jsonTag != "" {
			parts := strings.Split(jsonTag, ",")
			if parts[0] != "" {
				key = parts[0]
			}
		}

		result[key] = fieldVal.Interface()
	}

	return result, nil
}

func (gs *GenericSerializer) Deserialize(data map[string]interface{}, obj interface{}) error {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("destination must be a pointer")
	}

	val = val.Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		if !fieldVal.CanSet() {
			continue
		}

		jsonTag := field.Tag.Get("json")
		if jsonTag == "-" {
			continue
		}

		key := field.Name
		if jsonTag != "" {
			parts := strings.Split(jsonTag, ",")
			if parts[0] != "" {
				key = parts[0]
			}
		}

		if value, exists := data[key]; exists {
			srcVal := reflect.ValueOf(value)
			if srcVal.Type().AssignableTo(fieldVal.Type()) {
				fieldVal.Set(srcVal)
			} else if srcVal.Type().ConvertibleTo(fieldVal.Type()) {
				fieldVal.Set(srcVal.Convert(fieldVal.Type()))
			}
		}
	}

	return nil
}

// Configuration Binding
type ConfigBinder struct{}

func NewConfigBinder() *ConfigBinder {
	return &ConfigBinder{}
}

func (cb *ConfigBinder) Bind(data map[string]interface{}, obj interface{}) error {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("destination must be a pointer")
	}

	val = val.Elem()
	return cb.bindStruct(data, val, "")
}

func (cb *ConfigBinder) bindStruct(data map[string]interface{}, val reflect.Value, prefix string) error {
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		if !fieldVal.CanSet() {
			continue
		}

		mapTag := field.Tag.Get("mapstructure")
		if mapTag == "" {
			mapTag = strings.ToLower(field.Name)
		}

		key := mapTag
		if prefix != "" {
			key = prefix + "." + mapTag
		}

		if fieldVal.Kind() == reflect.Struct {
			// Recursively bind nested structs
			if err := cb.bindStruct(data, fieldVal, key); err != nil {
				return err
			}
		} else {
			// Bind simple field
			if value, exists := data[key]; exists {
				srcVal := reflect.ValueOf(value)
				if srcVal.Type().AssignableTo(fieldVal.Type()) {
					fieldVal.Set(srcVal)
				} else if srcVal.Type().ConvertibleTo(fieldVal.Type()) {
					fieldVal.Set(srcVal.Convert(fieldVal.Type()))
				}
			}
		}
	}

	return nil
}

// Test Data Builder
type TestDataBuilder struct {
	generators map[reflect.Kind]func() interface{}
}

func NewTestDataBuilder() *TestDataBuilder {
	tdb := &TestDataBuilder{
		generators: make(map[reflect.Kind]func() interface{}),
	}

	// Register default generators
	tdb.generators[reflect.String] = func() interface{} {
		return "test_string"
	}
	tdb.generators[reflect.Int] = func() interface{} {
		return 42
	}
	tdb.generators[reflect.Bool] = func() interface{} {
		return true
	}

	return tdb
}

func (tdb *TestDataBuilder) Build(prototype interface{}) interface{} {
	typ := reflect.TypeOf(prototype)
	val := reflect.New(typ).Elem()

	tdb.populateStruct(val)
	return val.Addr().Interface()
}

func (tdb *TestDataBuilder) BuildWithRules(prototype interface{}, rules map[string]interface{}) interface{} {
	obj := tdb.Build(prototype)
	val := reflect.ValueOf(obj).Elem()

	for fieldName, value := range rules {
		field := val.FieldByName(fieldName)
		if field.IsValid() && field.CanSet() {
			srcVal := reflect.ValueOf(value)
			if srcVal.Type().AssignableTo(field.Type()) {
				field.Set(srcVal)
			}
		}
	}

	return obj
}

func (tdb *TestDataBuilder) populateStruct(val reflect.Value) {
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		// field := typ.Field(i) // Get the field type
		fieldVal := val.Field(i)

		if !fieldVal.CanSet() {
			continue
		}

		if fieldVal.Kind() == reflect.Struct {
			tdb.populateStruct(fieldVal)
		} else if generator, exists := tdb.generators[fieldVal.Kind()]; exists {
			generated := generator()
			srcVal := reflect.ValueOf(generated)
			if srcVal.Type().AssignableTo(fieldVal.Type()) {
				fieldVal.Set(srcVal)
			}
		}
	}
}
