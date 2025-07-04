package main

import (
	"fmt"
	"math"
	"strconv"
)

// ==========================================
// Built-in Constraints
// ==========================================

// Example using 'any' constraint (equivalent to interface{})
func PrintAny[T any](value T) {
	fmt.Printf("Value: %v (Type: %T)\n", value, value)
}

// Example using 'comparable' constraint
func Equal[T comparable](a, b T) bool {
	return a == b
}

// Find returns the index of the first occurrence of value in slice
func Find[T comparable](slice []T, value T) int {
	for i, item := range slice {
		if item == value {
			return i
		}
	}
	return -1
}

// ==========================================
// Custom Constraints - Basic
// ==========================================

// Stringer represents types that can be converted to string
type Stringer interface {
	String() string
}

// DisplayInfo prints information about any Stringer
func DisplayInfo[T Stringer](item T) {
	fmt.Printf("Display: %s\n", item.String())
}

// Person implements Stringer interface
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (age %d)", p.Name, p.Age)
}

// ==========================================
// Numeric Constraints
// ==========================================

// Number constraint for all numeric types
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Integer constraint for integer types only
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Float constraint for floating-point types only
type Float interface {
	~float32 | ~float64
}

// Signed integer constraint
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned integer constraint
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Add returns the sum of two numbers
func Add[T Number](a, b T) T {
	return a + b
}

// Multiply returns the product of two numbers
func Multiply[T Number](a, b T) T {
	return a * b
}

// Abs returns the absolute value (works only with signed numbers)
func Abs[T Signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Sqrt returns the square root (works only with float types)
func Sqrt[T Float](x T) T {
	return T(math.Sqrt(float64(x)))
}

// ==========================================
// Complex Constraints - Method Sets
// ==========================================

// Reader interface for readable types
type Reader interface {
	Read() string
}

// Writer interface for writable types
type Writer interface {
	Write(string)
}

// ReadWriter combines Reader and Writer
type ReadWriter interface {
	Reader
	Writer
}

// Buffer implements both Reader and Writer
type Buffer struct {
	data string
}

func (b *Buffer) Read() string {
	return b.data
}

func (b *Buffer) Write(s string) {
	b.data += s
}

// ProcessReadWriter works with types that can both read and write
func ProcessReadWriter[T ReadWriter](rw T) {
	fmt.Println("Current content:", rw.Read())
	rw.Write(" [processed]")
	fmt.Println("After processing:", rw.Read())
}

// ==========================================
// Type Sets and Union Types
// ==========================================

// StringOrInt constraint allows only string or int types
type StringOrInt interface {
	string | int
}

// Convert handles conversion between string and int
func Convert[T StringOrInt](value T) any {
	switch v := any(value).(type) {
	case string:
		if num, err := strconv.Atoi(v); err == nil {
			return num
		}
		return 0
	case int:
		return strconv.Itoa(v)
	default:
		return nil
	}
}

// BasicType constraint for basic Go types
type BasicType interface {
	bool | string |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 |
		complex64 | complex128
}

// IsZero checks if a value is the zero value for its type
func IsZero[T BasicType](value T) bool {
	var zero T
	return value == zero
}

// ==========================================
// Underlying Type Constraints (~)
// ==========================================

// MyInt is a custom type based on int
type MyInt int

// MyString is a custom type based on string
type MyString string

// IntLike accepts int and types with underlying type int
type IntLike interface {
	~int
}

// StringLike accepts string and types with underlying type string
type StringLike interface {
	~string
}

// DoubleInt doubles an int-like value
func DoubleInt[T IntLike](x T) T {
	return x * 2
}

// UpperString converts a string-like value to uppercase
func UpperString[T StringLike](s T) T {
	return T(fmt.Sprintf("%s", string(s)))
}

// ==========================================
// Advanced Constraint Combinations
// ==========================================

// Serializable represents types that can be serialized
type Serializable interface {
	comparable
	Stringer
}

// SerializableMap is a generic map for serializable types
type SerializableMap[K comparable, V Serializable] map[K]V

// NewSerializableMap creates a new serializable map
func NewSerializableMap[K comparable, V Serializable]() SerializableMap[K, V] {
	return make(SerializableMap[K, V])
}

// Add adds a key-value pair to the map
func (sm SerializableMap[K, V]) Add(key K, value V) {
	sm[key] = value
}

// String returns a string representation of the map
func (sm SerializableMap[K, V]) String() string {
	result := "SerializableMap{"
	first := true
	for k, v := range sm {
		if !first {
			result += ", "
		}
		result += fmt.Sprintf("%v: %s", k, v.String())
		first = false
	}
	result += "}"
	return result
}

// ==========================================
// Constraint Interfaces with Type Parameters
// ==========================================

// Container represents a generic container
type Container[T any] interface {
	Add(T)
	Get() T
	Size() int
}

// SliceContainer implements Container for slices
type SliceContainer[T any] struct {
	items []T
}

func NewSliceContainer[T any]() *SliceContainer[T] {
	return &SliceContainer[T]{items: make([]T, 0)}
}

func (sc *SliceContainer[T]) Add(item T) {
	sc.items = append(sc.items, item)
}

func (sc *SliceContainer[T]) Get() T {
	if len(sc.items) == 0 {
		var zero T
		return zero
	}
	return sc.items[len(sc.items)-1]
}

func (sc *SliceContainer[T]) Size() int {
	return len(sc.items)
}

// ProcessContainer works with any Container implementation
func ProcessContainer[T any, C Container[T]](container C, items []T) {
	for _, item := range items {
		container.Add(item)
	}
	fmt.Printf("Container size after adding %d items: %d\n", len(items), container.Size())
	fmt.Printf("Last item: %v\n", container.Get())
}

// ==========================================
// Constraint Validation Examples
// ==========================================

// ValidatingConstraint ensures T implements both comparable and Stringer
type ValidatingConstraint interface {
	comparable
	Stringer
}

// ValidatedStore stores items that are both comparable and stringable
type ValidatedStore[T ValidatingConstraint] struct {
	items []T
}

func NewValidatedStore[T ValidatingConstraint]() *ValidatedStore[T] {
	return &ValidatedStore[T]{items: make([]T, 0)}
}

func (vs *ValidatedStore[T]) Add(item T) {
	// We can compare because T is comparable
	for _, existing := range vs.items {
		if existing == item {
			fmt.Printf("Item already exists: %s\n", item.String())
			return
		}
	}
	// We can call String() because T implements Stringer
	vs.items = append(vs.items, item)
	fmt.Printf("Added: %s\n", item.String())
}

func (vs *ValidatedStore[T]) List() {
	fmt.Println("Store contents:")
	for i, item := range vs.items {
		fmt.Printf("  %d: %s\n", i, item.String())
	}
}

// ==========================================
// Error Examples and Anti-Patterns
// ==========================================

// ❌ DON'T: Over-constraining when you don't need specific operations
// func BadExample[T string | int | float64](value T) T {
//     return value  // We're not using any operations that require these specific types
// }

// ✅ DO: Use appropriate constraints for what you actually need
func GoodExample[T any](value T) T {
	return value // any is sufficient since we're just returning the value
}

// ❌ DON'T: Constraining to exact types when underlying types should work
// func BadIntProcess[T int](x T) T {
//     return x * 2  // This won't work with MyInt type
// }

// ✅ DO: Use underlying type constraints
func GoodIntProcess[T ~int](x T) T {
	return x * 2 // This works with int, MyInt, and other int-based types
}

// ==========================================
// Main Example Function
// ==========================================

func runConstraintsExample() {
	fmt.Println("\n🔸 Built-in Constraints")

	// any constraint
	PrintAny(42)
	PrintAny("hello")
	PrintAny([]int{1, 2, 3})

	// comparable constraint
	fmt.Println("Equal(5, 5):", Equal(5, 5))
	fmt.Println("Equal(\"hello\", \"world\"):", Equal("hello", "world"))

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Find 3 in numbers:", Find(numbers, 3))
	fmt.Println("Find 6 in numbers:", Find(numbers, 6))

	fmt.Println("\n🔸 Custom Constraints")
	person := Person{Name: "Alice", Age: 30}
	DisplayInfo(person)

	fmt.Println("\n🔸 Numeric Constraints")
	fmt.Println("Add(5, 3):", Add(5, 3))
	fmt.Println("Add(2.5, 1.7):", Add(2.5, 1.7))
	fmt.Println("Multiply(4, 7):", Multiply(4, 7))

	fmt.Println("Abs(-42):", Abs(-42))
	fmt.Println("Sqrt(16.0):", Sqrt(16.0))
	fmt.Println("Sqrt(2.0):", Sqrt(2.0))

	fmt.Println("\n🔸 Method Set Constraints")
	buffer := &Buffer{data: "initial content"}
	ProcessReadWriter(buffer)

	fmt.Println("\n🔸 Union Type Constraints")
	fmt.Println("Convert string to int:", Convert("123"))
	fmt.Println("Convert int to string:", Convert(456))

	fmt.Println("IsZero(0):", IsZero(0))
	fmt.Println("IsZero(42):", IsZero(42))
	fmt.Println("IsZero(\"\"):", IsZero(""))
	fmt.Println("IsZero(\"hello\"):", IsZero("hello"))

	fmt.Println("\n🔸 Underlying Type Constraints")
	var myInt MyInt = 21
	var myString MyString = "hello world"

	fmt.Println("DoubleInt(21):", DoubleInt(21))
	fmt.Println("DoubleInt(MyInt(21)):", DoubleInt(myInt))
	fmt.Println("UpperString(MyString):", UpperString(myString))

	fmt.Println("\n🔸 Advanced Constraint Combinations")
	serMap := NewSerializableMap[string, Person]()
	serMap.Add("person1", Person{Name: "Bob", Age: 25})
	serMap.Add("person2", Person{Name: "Carol", Age: 35})
	fmt.Println("Serializable map:", serMap.String())

	fmt.Println("\n🔸 Generic Containers with Constraints")
	container := NewSliceContainer[int]()
	ProcessContainer(container, []int{1, 2, 3, 4, 5})

	fmt.Println("\n🔸 Validation Constraints")
	store := NewValidatedStore[Person]()
	store.Add(Person{Name: "Dave", Age: 40})
	store.Add(Person{Name: "Eve", Age: 28})
	store.Add(Person{Name: "Dave", Age: 40}) // Duplicate
	store.List()

	fmt.Println("\n✅ Constraints examples completed!")
}
