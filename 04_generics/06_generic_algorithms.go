package main

import (
	"fmt"
	"sort"
	"strings"
)

// ==========================================
// Generic Sorting Algorithms
// ==========================================

// Sortable represents types that can be ordered
type Sortable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

// BubbleSort implements bubble sort for any sortable type
func BubbleSort[T Sortable](slice []T) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// QuickSort implements quicksort for any sortable type
func QuickSort[T Sortable](slice []T) {
	if len(slice) < 2 {
		return
	}
	quickSortHelper(slice, 0, len(slice)-1)
}

func quickSortHelper[T Sortable](slice []T, low, high int) {
	if low < high {
		pivot := partition(slice, low, high)
		quickSortHelper(slice, low, pivot-1)
		quickSortHelper(slice, pivot+1, high)
	}
}

func partition[T Sortable](slice []T, low, high int) int {
	pivot := slice[high]
	i := low - 1

	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

// SortBy sorts a slice using a custom comparison function
func SortBy[T any](slice []T, less func(a, b T) bool) {
	sort.Slice(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
}

// ==========================================
// Generic Search Algorithms
// ==========================================

// LinearSearch finds the first occurrence of a value
func LinearSearch[T comparable](slice []T, target T) int {
	for i, item := range slice {
		if item == target {
			return i
		}
	}
	return -1
}

// BinarySearch finds a value in a sorted slice
func BinarySearch[T Sortable](slice []T, target T) int {
	left, right := 0, len(slice)-1

	for left <= right {
		mid := left + (right-left)/2

		if slice[mid] == target {
			return mid
		} else if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// FindAll returns all indices where the predicate is true
func FindAll[T any](slice []T, predicate func(T) bool) []int {
	var indices []int
	for i, item := range slice {
		if predicate(item) {
			indices = append(indices, i)
		}
	}
	return indices
}

// FindFirst returns the first element that satisfies the predicate
func FindFirst[T any](slice []T, predicate func(T) bool) (T, bool) {
	for _, item := range slice {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

// ==========================================
// Generic Transformation Algorithms
// ==========================================

// MapSlice applies a function to each element
func MapSlice[T, U any](slice []T, mapper func(T) U) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = mapper(item)
	}
	return result
}

// FilterSlice returns elements that satisfy the predicate
func FilterSlice[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// ReduceSlice reduces a slice to a single value
func ReduceSlice[T, U any](slice []T, initial U, reducer func(U, T) U) U {
	result := initial
	for _, item := range slice {
		result = reducer(result, item)
	}
	return result
}

// FlatMap applies a function and flattens the result
func FlatMap[T, U any](slice []T, mapper func(T) []U) []U {
	var result []U
	for _, item := range slice {
		result = append(result, mapper(item)...)
	}
	return result
}

// Partition splits a slice into two based on a predicate
func Partition[T any](slice []T, predicate func(T) bool) ([]T, []T) {
	var trueSlice, falseSlice []T
	for _, item := range slice {
		if predicate(item) {
			trueSlice = append(trueSlice, item)
		} else {
			falseSlice = append(falseSlice, item)
		}
	}
	return trueSlice, falseSlice
}

// ==========================================
// Generic Aggregation Algorithms
// ==========================================

// Sum calculates the sum of numeric values
func SumSlice[T Number](slice []T) T {
	var sum T
	for _, item := range slice {
		sum += item
	}
	return sum
}

// Product calculates the product of numeric values
func Product[T Number](slice []T) T {
	if len(slice) == 0 {
		return T(0)
	}

	product := T(1)
	for _, item := range slice {
		product *= item
	}
	return product
}

// Min finds the minimum value
func MinSlice[T Sortable](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}

	min := slice[0]
	for _, item := range slice[1:] {
		if item < min {
			min = item
		}
	}
	return min, true
}

// Max finds the maximum value
func MaxSlice[T Sortable](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}

	max := slice[0]
	for _, item := range slice[1:] {
		if item > max {
			max = item
		}
	}
	return max, true
}

// Average calculates the average of numeric values
func AverageSlice[T Number](slice []T) T {
	if len(slice) == 0 {
		return T(0)
	}
	return SumSlice(slice) / T(len(slice))
}

// ==========================================
// Generic Collection Operations
// ==========================================

// Reverse reverses a slice in place
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Unique removes duplicate elements (preserves order)
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	var result []T

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

// GroupBy groups elements by a key function
func GroupBy[T any, K comparable](slice []T, keyFunc func(T) K) map[K][]T {
	groups := make(map[K][]T)

	for _, item := range slice {
		key := keyFunc(item)
		groups[key] = append(groups[key], item)
	}
	return groups
}

// Zip combines two slices into pairs
func Zip[T, U any](slice1 []T, slice2 []U) []struct {
	First  T
	Second U
} {
	minLen := len(slice1)
	if len(slice2) < minLen {
		minLen = len(slice2)
	}

	result := make([]struct {
		First  T
		Second U
	}, minLen)
	for i := 0; i < minLen; i++ {
		result[i] = struct {
			First  T
			Second U
		}{slice1[i], slice2[i]}
	}
	return result
}

// ==========================================
// Functional Programming Patterns
// ==========================================

// Predicate represents a boolean function
type Predicate[T any] func(T) bool

// And combines predicates with logical AND
func And[T any](predicates ...Predicate[T]) Predicate[T] {
	return func(item T) bool {
		for _, pred := range predicates {
			if !pred(item) {
				return false
			}
		}
		return true
	}
}

// Or combines predicates with logical OR
func Or[T any](predicates ...Predicate[T]) Predicate[T] {
	return func(item T) bool {
		for _, pred := range predicates {
			if pred(item) {
				return true
			}
		}
		return false
	}
}

// Not negates a predicate
func Not[T any](predicate Predicate[T]) Predicate[T] {
	return func(item T) bool {
		return !predicate(item)
	}
}

// ==========================================
// Pipeline Operations
// ==========================================

// Pipeline represents a data processing pipeline
type Pipeline[T any] struct {
	data []T
}

// NewPipeline creates a new pipeline
func NewPipeline[T any](data []T) *Pipeline[T] {
	return &Pipeline[T]{data: data}
}

// Map applies a transformation
func (p *Pipeline[T]) Map(mapper func(T) T) *Pipeline[T] {
	p.data = MapSlice(p.data, mapper)
	return p
}

// Filter applies a filter
func (p *Pipeline[T]) Filter(predicate func(T) bool) *Pipeline[T] {
	p.data = FilterSlice(p.data, predicate)
	return p
}

// Sort sorts the data
func (p *Pipeline[T]) Sort(less func(a, b T) bool) *Pipeline[T] {
	SortBy(p.data, less)
	return p
}

// Take takes the first n elements
func (p *Pipeline[T]) Take(n int) *Pipeline[T] {
	if n > len(p.data) {
		n = len(p.data)
	}
	p.data = p.data[:n]
	return p
}

// Skip skips the first n elements
func (p *Pipeline[T]) Skip(n int) *Pipeline[T] {
	if n > len(p.data) {
		p.data = []T{}
	} else {
		p.data = p.data[n:]
	}
	return p
}

// Collect returns the final result
func (p *Pipeline[T]) Collect() []T {
	return p.data
}

// ==========================================
// Mathematical Algorithms
// ==========================================

// Fibonacci generates fibonacci numbers
func Fibonacci[T Number](n int) []T {
	if n <= 0 {
		return []T{}
	}
	if n == 1 {
		return []T{T(0)}
	}

	fib := make([]T, n)
	fib[0] = T(0)
	fib[1] = T(1)

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

// IsPrime checks if a number is prime
func IsPrime[T Integer](n T) bool {
	if n < 2 {
		return false
	}

	for i := T(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Range generates a range of numbers
func Range[T Number](start, end, step T) []T {
	var result []T
	if step > 0 {
		for i := start; i < end; i += step {
			result = append(result, i)
		}
	} else if step < 0 {
		for i := start; i > end; i += step {
			result = append(result, i)
		}
	}
	return result
}

// ==========================================
// String Algorithms
// ==========================================

// JoinStrings joins elements into a string
func JoinStrings[T any](slice []T, separator string, toString func(T) string) string {
	if len(slice) == 0 {
		return ""
	}

	parts := make([]string, len(slice))
	for i, item := range slice {
		parts[i] = toString(item)
	}
	return strings.Join(parts, separator)
}

// ==========================================
// Main Example Function
// ==========================================

func runGenericAlgorithmsExample() {
	fmt.Println("\n🔸 Generic Sorting Algorithms")

	// Bubble sort
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Original: %v\n", numbers)

	bubbleNums := make([]int, len(numbers))
	copy(bubbleNums, numbers)
	BubbleSort(bubbleNums)
	fmt.Printf("Bubble sort: %v\n", bubbleNums)

	// Quick sort
	quickNums := make([]int, len(numbers))
	copy(quickNums, numbers)
	QuickSort(quickNums)
	fmt.Printf("Quick sort: %v\n", quickNums)

	// Sort strings
	words := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("Original words: %v\n", words)
	BubbleSort(words)
	fmt.Printf("Sorted words: %v\n", words)

	fmt.Println("\n🔸 Custom Sorting")

	// Sort by length
	names := []string{"Alice", "Bob", "Charlie", "Dave"}
	SortBy(names, func(a, b string) bool {
		return len(a) < len(b)
	})
	fmt.Printf("Sorted by length: %v\n", names)

	fmt.Println("\n🔸 Generic Search Algorithms")

	sortedNums := []int{1, 3, 5, 7, 9, 11, 13, 15}
	fmt.Printf("Sorted array: %v\n", sortedNums)

	// Linear search
	fmt.Printf("Linear search for 7: index %d\n", LinearSearch(sortedNums, 7))
	fmt.Printf("Linear search for 4: index %d\n", LinearSearch(sortedNums, 4))

	// Binary search
	fmt.Printf("Binary search for 7: index %d\n", BinarySearch(sortedNums, 7))
	fmt.Printf("Binary search for 4: index %d\n", BinarySearch(sortedNums, 4))

	// Find all even numbers
	evenIndices := FindAll(sortedNums, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Even number indices: %v\n", evenIndices)

	// Find first number > 10
	if num, found := FindFirst(sortedNums, func(n int) bool { return n > 10 }); found {
		fmt.Printf("First number > 10: %d\n", num)
	}

	fmt.Println("\n🔸 Generic Transformations")

	// Map: square numbers
	squares := MapSlice([]int{1, 2, 3, 4, 5}, func(n int) int { return n * n })
	fmt.Printf("Squares: %v\n", squares)

	// Map: string lengths
	lengths := MapSlice([]string{"hello", "world", "go"}, func(s string) int { return len(s) })
	fmt.Printf("String lengths: %v\n", lengths)

	// Filter: even numbers
	evens := FilterSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Even numbers: %v\n", evens)

	// Reduce: sum
	sum := ReduceSlice([]int{1, 2, 3, 4, 5}, 0, func(acc, n int) int { return acc + n })
	fmt.Printf("Sum: %d\n", sum)

	// FlatMap: words to characters
	chars := FlatMap([]string{"hi", "go"}, func(s string) []rune { return []rune(s) })
	fmt.Printf("Characters: %v\n", chars)

	// Partition: even/odd
	even, odd := Partition([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Even: %v, Odd: %v\n", even, odd)

	fmt.Println("\n🔸 Generic Aggregations")

	testNums := []int{1, 2, 3, 4, 5}
	fmt.Printf("Numbers: %v\n", testNums)
	fmt.Printf("Sum: %d\n", SumSlice(testNums))
	fmt.Printf("Product: %d\n", Product(testNums))
	fmt.Printf("Average: %d\n", AverageSlice(testNums))

	if min, ok := MinSlice(testNums); ok {
		fmt.Printf("Min: %d\n", min)
	}
	if max, ok := MaxSlice(testNums); ok {
		fmt.Printf("Max: %d\n", max)
	}

	fmt.Println("\n🔸 Collection Operations")

	// Reverse
	reverseTest := []string{"a", "b", "c", "d"}
	fmt.Printf("Original: %v\n", reverseTest)
	Reverse(reverseTest)
	fmt.Printf("Reversed: %v\n", reverseTest)

	// Unique
	duplicates := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
	unique := Unique(duplicates)
	fmt.Printf("Duplicates: %v\n", duplicates)
	fmt.Printf("Unique: %v\n", unique)

	// GroupBy
	items := []string{"apple", "apricot", "banana", "blueberry", "cherry"}
	groups := GroupBy(items, func(s string) rune { return []rune(s)[0] })
	fmt.Printf("Grouped by first letter: %v\n", groups)

	// Zip
	nums1 := []int{1, 2, 3}
	strs1 := []string{"a", "b", "c"}
	zipped := Zip(nums1, strs1)
	fmt.Printf("Zipped: %v\n", zipped)

	fmt.Println("\n🔸 Functional Programming")

	// Predicate combinators
	isEven := func(n int) bool { return n%2 == 0 }
	isPositive := func(n int) bool { return n > 0 }
	isGreaterThan5 := func(n int) bool { return n > 5 }

	testData := []int{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8}

	evenAndPositive := FilterSlice(testData, And(isEven, isPositive))
	fmt.Printf("Even and positive: %v\n", evenAndPositive)

	evenOrGreaterThan5 := FilterSlice(testData, Or(isEven, isGreaterThan5))
	fmt.Printf("Even or > 5: %v\n", evenOrGreaterThan5)

	notEven := FilterSlice(testData, Not(isEven))
	fmt.Printf("Not even: %v\n", notEven)

	fmt.Println("\n🔸 Pipeline Operations")

	pipelineData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	result := NewPipeline(pipelineData).
		Filter(func(n int) bool { return n%2 == 0 }). // Keep evens
		Map(func(n int) int { return n * n }).        // Square them
		Sort(func(a, b int) bool { return a > b }).   // Sort descending
		Take(3).                                      // Take first 3
		Collect()

	fmt.Printf("Pipeline result: %v\n", result)

	fmt.Println("\n🔸 Mathematical Algorithms")

	// Fibonacci
	fib := Fibonacci[int](10)
	fmt.Printf("Fibonacci(10): %v\n", fib)

	// Prime checking
	primes := FilterSlice(Range(2, 20, 1), IsPrime[int])
	fmt.Printf("Primes 2-19: %v\n", primes)

	// Range generation
	range1 := Range(0, 10, 2)
	fmt.Printf("Range(0, 10, 2): %v\n", range1)

	range2 := Range(10, 0, -3)
	fmt.Printf("Range(10, 0, -3): %v\n", range2)

	fmt.Println("\n🔸 String Operations")

	// Join with custom toString
	joined := JoinStrings([]int{1, 2, 3, 4, 5}, " -> ", func(n int) string {
		return fmt.Sprintf("num(%d)", n)
	})
	fmt.Printf("Joined: %s\n", joined)

	fmt.Println("\n✅ Generic algorithms examples completed!")
}
