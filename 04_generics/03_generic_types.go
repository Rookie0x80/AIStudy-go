package main

import (
	"fmt"
	"sync"
)

// ==========================================
// Basic Generic Structs
// ==========================================

// KeyValuePair represents a generic key-value pair
type KeyValuePair[T, U any] struct {
	Key   T
	Value U
}

// NewKeyValuePair creates a new key-value pair
func NewKeyValuePair[T, U any](key T, value U) KeyValuePair[T, U] {
	return KeyValuePair[T, U]{Key: key, Value: value}
}

// Swap exchanges the key and value
func (p KeyValuePair[T, U]) Swap() KeyValuePair[U, T] {
	return KeyValuePair[U, T]{Key: p.Value, Value: p.Key}
}

// String returns a string representation
func (p KeyValuePair[T, U]) String() string {
	return fmt.Sprintf("(%v: %v)", p.Key, p.Value)
}

// ==========================================
// Generic Collections
// ==========================================

// Stack represents a generic stack
type Stack[T any] struct {
	items []T
	mu    sync.RWMutex
}

// NewStack creates a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, item)
}

// Pop removes and returns the top item
func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

// Peek returns the top item without removing it
func (s *Stack[T]) Peek() (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	return s.items[len(s.items)-1], true
}

// Size returns the number of items in the stack
func (s *Stack[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.items)
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// ==========================================
// Generic Linked List
// ==========================================

// Node represents a node in a linked list
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList represents a generic linked list
type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// NewLinkedList creates a new linked list
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Add appends an item to the end of the list
func (ll *LinkedList[T]) Add(value T) {
	newNode := &Node[T]{Value: value}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.Next = newNode
		ll.tail = newNode
	}
	ll.size++
}

// Prepend adds an item to the beginning of the list
func (ll *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value, Next: ll.head}
	ll.head = newNode

	if ll.tail == nil {
		ll.tail = newNode
	}
	ll.size++
}

// Get returns the value at the specified index
func (ll *LinkedList[T]) Get(index int) (T, bool) {
	if index < 0 || index >= ll.size {
		var zero T
		return zero, false
	}

	current := ll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Value, true
}

// Size returns the number of items in the list
func (ll *LinkedList[T]) Size() int {
	return ll.size
}

// ToSlice converts the linked list to a slice
func (ll *LinkedList[T]) ToSlice() []T {
	result := make([]T, 0, ll.size)
	current := ll.head

	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result
}

// ==========================================
// Generic Interfaces
// ==========================================

// GenericContainer represents a generic container interface
type GenericContainer[T any] interface {
	Add(T)
	Size() int
	IsEmpty() bool
}

// Iterable represents a generic iterable interface
type Iterable[T any] interface {
	ForEach(func(T))
}

// Collection combines GenericContainer and Iterable
type Collection[T any] interface {
	GenericContainer[T]
	Iterable[T]
}

// ==========================================
// Generic Map with Methods
// ==========================================

// SafeMap represents a thread-safe generic key-value map
type SafeMap[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex
}

// NewSafeMap creates a new thread-safe generic map
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

// Set stores a key-value pair
func (m *SafeMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

// Get retrieves a value by key
func (m *SafeMap[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, exists := m.data[key]
	return value, exists
}

// Delete removes a key-value pair
func (m *SafeMap[K, V]) Delete(key K) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.data[key]; exists {
		delete(m.data, key)
		return true
	}
	return false
}

// Keys returns all keys
func (m *SafeMap[K, V]) Keys() []K {
	m.mu.RLock()
	defer m.mu.RUnlock()

	keys := make([]K, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

// Values returns all values
func (m *SafeMap[K, V]) Values() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()

	values := make([]V, 0, len(m.data))
	for _, v := range m.data {
		values = append(values, v)
	}
	return values
}

// Size returns the number of key-value pairs
func (m *SafeMap[K, V]) Size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.data)
}

// ForEach iterates over all key-value pairs
func (m *SafeMap[K, V]) ForEach(fn func(K, V)) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for k, v := range m.data {
		fn(k, v)
	}
}

// ==========================================
// Generic Result Type (Rust-inspired)
// ==========================================

// Result represents either a success value or an error
type Result[T any] struct {
	value T
	err   error
	isOk  bool
}

// Ok creates a successful Result
func Ok[T any](value T) Result[T] {
	return Result[T]{value: value, isOk: true}
}

// Err creates an error Result
func Err[T any](err error) Result[T] {
	return Result[T]{err: err, isOk: false}
}

// IsOk checks if the result is successful
func (r Result[T]) IsOk() bool {
	return r.isOk
}

// IsErr checks if the result is an error
func (r Result[T]) IsErr() bool {
	return !r.isOk
}

// Unwrap returns the value (panics if error)
func (r Result[T]) Unwrap() T {
	if !r.isOk {
		panic(fmt.Sprintf("called Unwrap on error Result: %v", r.err))
	}
	return r.value
}

// UnwrapOr returns the value or a default
func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.isOk {
		return r.value
	}
	return defaultValue
}

// Error returns the error
func (r Result[T]) Error() error {
	return r.err
}

// Map transforms the value if Ok
func (r Result[T]) Map(fn func(T) T) Result[T] {
	if r.isOk {
		return Ok(fn(r.value))
	}
	return Err[T](r.err)
}

// ==========================================
// Generic Cache with TTL
// ==========================================

// CacheItem represents a cached item with expiration
type CacheItem[T any] struct {
	Value   T
	Expires int64
}

// Cache represents a generic cache with TTL
type Cache[K comparable, V any] struct {
	items map[K]CacheItem[V]
	mu    sync.RWMutex
}

// NewCache creates a new cache
func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		items: make(map[K]CacheItem[V]),
	}
}

// Set stores a value with TTL (in seconds)
func (c *Cache[K, V]) Set(key K, value V, ttl int64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expires := int64(0)
	if ttl > 0 {
		expires = ttl // Simplified: would use time.Now().Unix() + ttl in real implementation
	}

	c.items[key] = CacheItem[V]{Value: value, Expires: expires}
}

// Get retrieves a value from cache
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.items[key]
	if !exists {
		var zero V
		return zero, false
	}

	// Simplified TTL check (would check against current time in real implementation)
	if item.Expires > 0 {
		// In real implementation: if time.Now().Unix() > item.Expires
		// For demo purposes, we'll just return the value
	}

	return item.Value, true
}

// ==========================================
// Method Receivers with Type Parameters
// ==========================================

// Vector represents a generic vector
type Vector[T any] []T

// NewVector creates a new vector
func NewVector[T any]() Vector[T] {
	return make(Vector[T], 0)
}

// Add appends an element to the vector
func (v *Vector[T]) Add(item T) {
	*v = append(*v, item)
}

// Get returns the element at index
func (v Vector[T]) Get(index int) (T, bool) {
	if index < 0 || index >= len(v) {
		var zero T
		return zero, false
	}
	return v[index], true
}

// Len returns the length of the vector
func (v Vector[T]) Len() int {
	return len(v)
}

// Filter creates a new vector with elements that satisfy the predicate
func (v Vector[T]) Filter(predicate func(T) bool) Vector[T] {
	result := NewVector[T]()
	for _, item := range v {
		if predicate(item) {
			result.Add(item)
		}
	}
	return result
}

// Map transforms each element and returns a new vector
func (v Vector[T]) Map(transform func(T) T) Vector[T] {
	result := make(Vector[T], len(v))
	for i, item := range v {
		result[i] = transform(item)
	}
	return result
}

// ==========================================
// Main Example Function
// ==========================================

func runGenericTypesExample() {
	fmt.Println("\n🔸 Basic Generic Structs")

	// KeyValuePair examples
	intStringPair := NewKeyValuePair(42, "hello")
	fmt.Printf("Original pair: %s\n", intStringPair.String())

	swapped := intStringPair.Swap()
	fmt.Printf("Swapped pair: %s\n", swapped.String())

	floatBoolPair := NewKeyValuePair(3.14, true)
	fmt.Printf("Float-bool pair: %s\n", floatBoolPair.String())

	fmt.Println("\n🔸 Generic Stack")

	// Integer stack
	intStack := NewStack[int]()
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Printf("Stack size: %d\n", intStack.Size())

	if top, ok := intStack.Peek(); ok {
		fmt.Printf("Top element: %d\n", top)
	}

	for !intStack.IsEmpty() {
		if item, ok := intStack.Pop(); ok {
			fmt.Printf("Popped: %d\n", item)
		}
	}

	// String stack
	stringStack := NewStack[string]()
	stringStack.Push("first")
	stringStack.Push("second")
	stringStack.Push("third")

	fmt.Printf("String stack size: %d\n", stringStack.Size())

	fmt.Println("\n🔸 Generic Linked List")

	list := NewLinkedList[string]()
	list.Add("apple")
	list.Add("banana")
	list.Prepend("cherry")

	fmt.Printf("List size: %d\n", list.Size())
	fmt.Printf("List contents: %v\n", list.ToSlice())

	if value, ok := list.Get(1); ok {
		fmt.Printf("Element at index 1: %s\n", value)
	}

	fmt.Println("\n🔸 Generic Safe Map")

	userMap := NewSafeMap[string, int]()
	userMap.Set("alice", 25)
	userMap.Set("bob", 30)
	userMap.Set("charlie", 35)

	if age, exists := userMap.Get("alice"); exists {
		fmt.Printf("Alice's age: %d\n", age)
	}

	fmt.Printf("Map size: %d\n", userMap.Size())
	fmt.Printf("All keys: %v\n", userMap.Keys())

	userMap.ForEach(func(name string, age int) {
		fmt.Printf("User %s is %d years old\n", name, age)
	})

	fmt.Println("\n🔸 Generic Result Type")

	// Success case
	successResult := Ok[string]("operation successful")
	fmt.Printf("Result is ok: %t\n", successResult.IsOk())
	fmt.Printf("Success value: %s\n", successResult.Unwrap())

	// Error case
	errorResult := Err[string](fmt.Errorf("something went wrong"))
	fmt.Printf("Result is error: %t\n", errorResult.IsErr())
	fmt.Printf("Error value with default: %s\n", errorResult.UnwrapOr("default value"))

	// Map transformation
	transformed := successResult.Map(func(s string) string {
		return fmt.Sprintf("Transformed: %s", s)
	})
	fmt.Printf("Transformed result: %s\n", transformed.Unwrap())

	fmt.Println("\n🔸 Generic Cache")

	cache := NewCache[string, int]()
	cache.Set("user:123", 25, 3600)
	cache.Set("user:456", 30, 7200)

	if age, found := cache.Get("user:123"); found {
		fmt.Printf("Cached age for user:123: %d\n", age)
	}

	fmt.Println("\n🔸 Generic Vector with Methods")

	numbers := NewVector[int]()
	numbers.Add(1)
	numbers.Add(2)
	numbers.Add(3)
	numbers.Add(4)
	numbers.Add(5)

	fmt.Printf("Vector length: %d\n", numbers.Len())

	// Filter even numbers
	evens := numbers.Filter(func(n int) bool { return n%2 == 0 })
	fmt.Printf("Even numbers: %v\n", []int(evens))

	// Double all numbers
	doubled := numbers.Map(func(n int) int { return n * 2 })
	fmt.Printf("Doubled numbers: %v\n", []int(doubled))

	fmt.Println("\n✅ Generic types examples completed!")
}
