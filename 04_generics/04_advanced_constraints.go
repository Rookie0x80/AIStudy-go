package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ==========================================
// Type Sets and Complex Constraints
// ==========================================

// IntegerTypes defines a constraint for all integer types
type IntegerTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// NumericTypes defines a constraint for all numeric types
type NumericTypes interface {
	IntegerTypes | ~float32 | ~float64
}

// BasicTypes defines a constraint for basic Go types
type BasicTypes interface {
	~bool | ~string | NumericTypes
}

// StringLikeTypes defines string-like types
type StringLikeTypes interface {
	~string | ~[]byte | ~[]rune
}

// ==========================================
// Custom Types for Testing
// ==========================================

type CustomInt int
type CustomFloat float64
type CustomString string
type CustomBytes []byte

// ==========================================
// Advanced Constraint Functions
// ==========================================

// Sum works with any integer type (including custom ones)
func Sum[T IntegerTypes](values ...T) T {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum
}

// Average works with any numeric type
func Average[T NumericTypes](values ...T) T {
	if len(values) == 0 {
		return T(0)
	}

	var sum T
	for _, v := range values {
		sum += v
	}
	return sum / T(len(values))
}

// ToString converts various types to string
func ToString[T BasicTypes](value T) string {
	switch v := any(value).(type) {
	case bool:
		return strconv.FormatBool(v)
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// ConvertToString handles string-like types
func ConvertToString[T StringLikeTypes](value T) string {
	switch v := any(value).(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case []rune:
		return string(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// ==========================================
// Union Types with Methods
// ==========================================

// Identifier can be either string or int
type Identifier interface {
	~string | ~int
}

// User with generic ID
type User[ID Identifier] struct {
	ID   ID
	Name string
	Age  int
}

// NewUser creates a new user
func NewUser[ID Identifier](id ID, name string, age int) User[ID] {
	return User[ID]{ID: id, Name: name, Age: age}
}

// String returns string representation
func (u User[ID]) String() string {
	return fmt.Sprintf("User{ID: %v, Name: %s, Age: %d}", u.ID, u.Name, u.Age)
}

// GetIDAsString returns ID as string
func (u User[ID]) GetIDAsString() string {
	switch id := any(u.ID).(type) {
	case string:
		return id
	case int:
		return strconv.Itoa(id)
	default:
		return fmt.Sprintf("%v", id)
	}
}

// ==========================================
// Constrained Interfaces
// ==========================================

// Comparable types that can be ordered
type Orderable interface {
	comparable
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

// SortableSlice represents a slice that can be sorted
type SortableSlice[T Orderable] []T

// Sort sorts the slice in ascending order
func (s SortableSlice[T]) Sort() {
	// Simple bubble sort for demonstration
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

// Search finds the index of a value
func (s SortableSlice[T]) Search(value T) int {
	for i, v := range s {
		if v == value {
			return i
		}
	}
	return -1
}

// ==========================================
// Complex Type Constraints
// ==========================================

// AdvancedSerializable combines multiple constraints
type AdvancedSerializable interface {
	fmt.Stringer
	comparable
}

// Hashable represents types that can be hashed
type Hashable interface {
	comparable
}

// Cacheable represents types that can be cached
type Cacheable interface {
	AdvancedSerializable
	Hashable
}

// CacheEntry represents a cache entry
type CacheEntry[K Hashable, V Cacheable] struct {
	Key   K
	Value V
	Hash  uint64
}

// NewCacheEntry creates a new cache entry
func NewCacheEntry[K Hashable, V Cacheable](key K, value V) CacheEntry[K, V] {
	// Simple hash function for demonstration
	hash := uint64(0)
	keyStr := fmt.Sprintf("%v", key)
	for _, c := range keyStr {
		hash = hash*31 + uint64(c)
	}

	return CacheEntry[K, V]{
		Key:   key,
		Value: value,
		Hash:  hash,
	}
}

// ==========================================
// Type Inference with Constraints
// ==========================================

// Processor interface with constraint
type Processor[T any] interface {
	Process(T) T
}

// StringProcessor implements Processor for strings
type StringProcessor struct{}

func (sp StringProcessor) Process(s string) string {
	return strings.ToUpper(s)
}

// IntProcessor implements Processor for integers
type IntProcessor struct{}

func (ip IntProcessor) Process(i int) int {
	return i * 2
}

// ProcessBatch processes a batch of items
func ProcessBatch[T any, P Processor[T]](items []T, processor P) []T {
	result := make([]T, len(items))
	for i, item := range items {
		result[i] = processor.Process(item)
	}
	return result
}

// ==========================================
// Advanced Constraint Validation
// ==========================================

// Validator interface for validation
type Validator[T any] interface {
	Validate(T) bool
}

// RangeValidator validates numeric ranges
type RangeValidator[T NumericTypes] struct {
	Min, Max T
}

func (rv RangeValidator[T]) Validate(value T) bool {
	return value >= rv.Min && value <= rv.Max
}

// LengthValidator validates string length
type LengthValidator struct {
	MinLen, MaxLen int
}

func (lv LengthValidator) Validate(value string) bool {
	length := len(value)
	return length >= lv.MinLen && length <= lv.MaxLen
}

// ValidatedValue represents a validated value
type ValidatedValue[T any, V Validator[T]] struct {
	Value     T
	Validator V
	IsValid   bool
}

// NewValidatedValue creates a validated value
func NewValidatedValue[T any, V Validator[T]](value T, validator V) ValidatedValue[T, V] {
	return ValidatedValue[T, V]{
		Value:     value,
		Validator: validator,
		IsValid:   validator.Validate(value),
	}
}

// ==========================================
// Error Handling with Constraints
// ==========================================

// Result type with constraints
type ValidationResult[T any] struct {
	Value T
	Error string
	Valid bool
}

// ValidateAndConvert validates and converts values
func ValidateAndConvert[T NumericTypes](input string) ValidationResult[T] {
	// Try to parse as float64 first
	if f, err := strconv.ParseFloat(input, 64); err == nil {
		value := T(f)
		return ValidationResult[T]{
			Value: value,
			Valid: true,
		}
	}

	var zero T
	return ValidationResult[T]{
		Value: zero,
		Error: "invalid numeric format",
		Valid: false,
	}
}

// ==========================================
// Main Example Function
// ==========================================

func runAdvancedConstraintsExample() {
	fmt.Println("\n🔸 Type Sets and Integer Constraints")

	// Integer sum with various types
	fmt.Println("Sum of ints:", Sum(1, 2, 3, 4, 5))
	fmt.Println("Sum of CustomInt:", Sum(CustomInt(10), CustomInt(20), CustomInt(30)))

	// Custom integer types
	var a, b, c CustomInt = 5, 10, 15
	fmt.Println("Sum of custom ints:", Sum(a, b, c))

	fmt.Println("\n🔸 Numeric Type Constraints")

	// Average with different numeric types
	fmt.Println("Average of ints:", Average(1, 2, 3, 4, 5))
	fmt.Println("Average of floats:", Average(1.5, 2.5, 3.5))
	fmt.Println("Average of CustomFloat:", Average(CustomFloat(1.0), CustomFloat(2.0), CustomFloat(3.0)))

	fmt.Println("\n🔸 Basic Type Conversions")

	// ToString with various types
	fmt.Println("Bool to string:", ToString(true))
	fmt.Println("Int to string:", ToString(42))
	fmt.Println("Float to string:", ToString(3.14))
	fmt.Println("String to string:", ToString("hello"))
	fmt.Println("CustomInt to string:", ToString(CustomInt(100)))

	fmt.Println("\n🔸 String-like Types")

	// String-like conversions
	fmt.Println("String:", ConvertToString("hello"))
	fmt.Println("Bytes:", ConvertToString([]byte("world")))
	fmt.Println("Runes:", ConvertToString([]rune("golang")))
	fmt.Println("CustomString:", ConvertToString(CustomString("custom")))
	fmt.Println("CustomBytes:", ConvertToString(CustomBytes("bytes")))

	fmt.Println("\n🔸 Union Types with Generic Structs")

	// Users with different ID types
	stringUser := NewUser("user123", "Alice", 30)
	intUser := NewUser(456, "Bob", 25)

	fmt.Println("String ID user:", stringUser.String())
	fmt.Println("String ID as string:", stringUser.GetIDAsString())

	fmt.Println("Int ID user:", intUser.String())
	fmt.Println("Int ID as string:", intUser.GetIDAsString())

	fmt.Println("\n🔸 Sortable Slices with Constraints")

	// Integer slice
	intSlice := SortableSlice[int]{5, 2, 8, 1, 9, 3}
	fmt.Printf("Before sort: %v\n", []int(intSlice))
	intSlice.Sort()
	fmt.Printf("After sort: %v\n", []int(intSlice))
	fmt.Printf("Search for 8: index %d\n", intSlice.Search(8))

	// String slice
	stringSlice := SortableSlice[string]{"banana", "apple", "cherry", "date"}
	fmt.Printf("Before sort: %v\n", []string(stringSlice))
	stringSlice.Sort()
	fmt.Printf("After sort: %v\n", []string(stringSlice))

	fmt.Println("\n🔸 Complex Constraint Combinations")

	// Cache entries (requires multiple constraints)
	entry1 := NewCacheEntry("key1", stringUser)
	fmt.Printf("Cache entry: Key=%v, Hash=%d, Value=%s\n",
		entry1.Key, entry1.Hash, entry1.Value.String())

	fmt.Println("\n🔸 Processor Pattern with Constraints")

	// String processing
	strings := []string{"hello", "world", "golang"}
	stringProcessor := StringProcessor{}
	processedStrings := ProcessBatch(strings, stringProcessor)
	fmt.Printf("Processed strings: %v\n", processedStrings)

	// Integer processing
	numbers := []int{1, 2, 3, 4, 5}
	intProcessor := IntProcessor{}
	processedNumbers := ProcessBatch(numbers, intProcessor)
	fmt.Printf("Processed numbers: %v\n", processedNumbers)

	fmt.Println("\n🔸 Validation with Constraints")

	// Numeric range validation
	rangeValidator := RangeValidator[int]{Min: 1, Max: 100}
	validatedNum := NewValidatedValue(50, rangeValidator)
	fmt.Printf("Number 50 valid (1-100): %t\n", validatedNum.IsValid)

	invalidNum := NewValidatedValue(150, rangeValidator)
	fmt.Printf("Number 150 valid (1-100): %t\n", invalidNum.IsValid)

	// String length validation
	lengthValidator := LengthValidator{MinLen: 3, MaxLen: 10}
	validatedStr := NewValidatedValue("hello", lengthValidator)
	fmt.Printf("String 'hello' valid (3-10 chars): %t\n", validatedStr.IsValid)

	fmt.Println("\n🔸 Validation and Conversion")

	// Validate and convert strings to numbers
	result1 := ValidateAndConvert[int]("42")
	fmt.Printf("Convert '42' to int: valid=%t, value=%d\n", result1.Valid, result1.Value)

	result2 := ValidateAndConvert[float64]("3.14")
	fmt.Printf("Convert '3.14' to float64: valid=%t, value=%f\n", result2.Valid, result2.Value)

	result3 := ValidateAndConvert[int]("invalid")
	fmt.Printf("Convert 'invalid' to int: valid=%t, error=%s\n", result3.Valid, result3.Error)

	fmt.Println("\n✅ Advanced constraints examples completed!")
}
