package main

import (
	"fmt"
	"strings"
)

// Ordered represents types that can be ordered
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// ==========================================
// Basic Generic Functions
// ==========================================

// BasicMin returns the minimum of two values of any ordered type
func BasicMin[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// BasicMax returns the maximum of two values of any ordered type
func BasicMax[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Print demonstrates a generic function with any type
func Print[T any](value T) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// ==========================================
// Type Inference Examples
// ==========================================

// Swap demonstrates type inference - no need to specify types explicitly
func Swap[T any](a, b T) (T, T) {
	return b, a
}

// Contains checks if a slice contains a specific value
func Contains[T comparable](slice []T, value T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// Map applies a transformation function to each element in a slice
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = fn(item)
	}
	return result
}

// Filter creates a new slice with elements that satisfy the predicate
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// ==========================================
// Multiple Type Parameters
// ==========================================

// Pair represents a generic pair of values
type Pair[T, U any] struct {
	First  T
	Second U
}

// NewPair creates a new pair (demonstrates type inference)
func NewPair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{First: first, Second: second}
}

// Transform applies functions to both elements of a pair
func Transform[T, U, V, W any](
	pair Pair[T, U],
	fn1 func(T) V,
	fn2 func(U) W,
) Pair[V, W] {
	return Pair[V, W]{
		First:  fn1(pair.First),
		Second: fn2(pair.Second),
	}
}

// ==========================================
// Real-World Examples
// ==========================================

// Optional represents a value that may or may not be present (like Rust's Option)
type Optional[T any] struct {
	value    T
	hasValue bool
}

// Some creates an Optional with a value
func Some[T any](value T) Optional[T] {
	return Optional[T]{value: value, hasValue: true}
}

// None creates an empty Optional
func None[T any]() Optional[T] {
	return Optional[T]{hasValue: false}
}

// IsSome checks if the Optional has a value
func (o Optional[T]) IsSome() bool {
	return o.hasValue
}

// IsNone checks if the Optional is empty
func (o Optional[T]) IsNone() bool {
	return !o.hasValue
}

// Unwrap returns the value (panics if None)
func (o Optional[T]) Unwrap() T {
	if !o.hasValue {
		panic("called Unwrap on a None value")
	}
	return o.value
}

// UnwrapOr returns the value or a default
func (o Optional[T]) UnwrapOr(defaultValue T) T {
	if o.hasValue {
		return o.value
	}
	return defaultValue
}

// ==========================================
// Advanced Function Signatures
// ==========================================

// Reduce applies a function to accumulate slice elements into a single value
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, item := range slice {
		result = fn(result, item)
	}
	return result
}

// Chain allows chaining operations on slices
func Chain[T any](slice []T) *ChainOp[T] {
	return &ChainOp[T]{data: slice}
}

type ChainOp[T any] struct {
	data []T
}

func (c *ChainOp[T]) Map(fn func(T) T) *ChainOp[T] {
	c.data = Map(c.data, fn)
	return c
}

func (c *ChainOp[T]) Filter(predicate func(T) bool) *ChainOp[T] {
	c.data = Filter(c.data, predicate)
	return c
}

func (c *ChainOp[T]) Collect() []T {
	return c.data
}

// ==========================================
// Common Mistakes and Pitfalls
// ==========================================

// ❌ DON'T: Over-constraining when any would work
// func BadPrint[T string | int](value T) {
//     fmt.Println(value)
// }

// ✅ DO: Use any when no specific operations are needed
func GoodPrint[T any](value T) {
	fmt.Println(value)
}

// ❌ DON'T: Unnecessary type specifications when inference works
// func callWithExplicitTypes() {
//     result := BasicMin[int](5, 3)  // Unnecessary [int]
// }

// ✅ DO: Let type inference work
func callWithInference() {
	result := BasicMin(5, 3) // Type inferred as int
	fmt.Println("Min result:", result)
}

// ==========================================
// Main Example Function
// ==========================================

func runBasicGenericsExample() {
	fmt.Println("\n🔸 Basic Generic Functions")
	fmt.Println("Min(5, 3):", BasicMin(5, 3))
	fmt.Println("Min(3.14, 2.71):", BasicMin(3.14, 2.71))
	fmt.Println("Min(\"apple\", \"banana\"):", BasicMin("apple", "banana"))

	fmt.Println("\nMax(5, 3):", BasicMax(5, 3))
	fmt.Println("Max(3.14, 2.71):", BasicMax(3.14, 2.71))

	fmt.Println("\n🔸 Generic Print Function")
	Print(42)
	Print("Hello, Generics!")
	Print([]int{1, 2, 3})

	fmt.Println("\n🔸 Type Inference")
	a, b := Swap("hello", "world")
	fmt.Printf("Swapped: %s, %s\n", a, b)

	x, y := Swap(10, 20)
	fmt.Printf("Swapped: %d, %d\n", x, y)

	fmt.Println("\n🔸 Slice Operations")
	numbers := []int{1, 2, 3, 4, 5}
	words := []string{"hello", "world", "go", "generics"}

	fmt.Println("Contains 3 in numbers:", Contains(numbers, 3))
	fmt.Println("Contains \"go\" in words:", Contains(words, "go"))

	// Map: double each number
	doubled := Map(numbers, func(n int) int { return n * 2 })
	fmt.Println("Doubled numbers:", doubled)

	// Map: uppercase each word
	uppercased := Map(words, strings.ToUpper)
	fmt.Println("Uppercased words:", uppercased)

	// Map: transform int to string
	numberStrings := Map(numbers, func(n int) string {
		return fmt.Sprintf("num-%d", n)
	})
	fmt.Println("Number strings:", numberStrings)

	// Filter: even numbers only
	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Println("Even numbers:", evens)

	// Filter: words with length > 2
	longWords := Filter(words, func(s string) bool { return len(s) > 2 })
	fmt.Println("Long words:", longWords)

	fmt.Println("\n🔸 Multiple Type Parameters")
	pair1 := NewPair("hello", 42)
	fmt.Printf("Pair: (%v, %v)\n", pair1.First, pair1.Second)

	// Transform pair
	transformed := Transform(
		pair1,
		strings.ToUpper,
		func(n int) float64 { return float64(n) * 1.5 },
	)
	fmt.Printf("Transformed: (%v, %v)\n", transformed.First, transformed.Second)

	fmt.Println("\n🔸 Optional Type (Rust-like)")
	someValue := Some(42)
	noneValue := None[int]()

	fmt.Println("Some value exists:", someValue.IsSome())
	fmt.Println("Some value:", someValue.Unwrap())

	fmt.Println("None value exists:", noneValue.IsSome())
	fmt.Println("None with default:", noneValue.UnwrapOr(99))

	fmt.Println("\n🔸 Functional Programming")
	// Sum using Reduce
	sum := Reduce(numbers, 0, func(acc, n int) int { return acc + n })
	fmt.Println("Sum:", sum)

	// Product using Reduce
	product := Reduce(numbers, 1, func(acc, n int) int { return acc * n })
	fmt.Println("Product:", product)

	// Concatenate strings
	concat := Reduce(words, "", func(acc, s string) string {
		if acc == "" {
			return s
		}
		return acc + "-" + s
	})
	fmt.Println("Concatenated:", concat)

	fmt.Println("\n🔸 Method Chaining")
	result := Chain([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		Filter(func(n int) bool { return n%2 == 0 }). // Keep evens
		Map(func(n int) int { return n * n }).        // Square them
		Collect()
	fmt.Println("Chained operations (even squares):", result)

	fmt.Println("\n🔸 Type Inference Best Practices")
	callWithInference()

	fmt.Println("\n✅ Basic generics examples completed!")
}
