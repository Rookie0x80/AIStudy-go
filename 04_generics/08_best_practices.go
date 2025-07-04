package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"time"
)

// ==========================================
// Best Practice 1: Prefer Specific Constraints
// ==========================================

// ❌ BAD: Too generic constraint
func badAdd[T any](a, b T) T {
	// This won't compile - no addition operation for 'any'
	// return a + b
	var zero T
	return zero
}

// ✅ GOOD: Specific constraint for the operation
func goodAdd[T Number](a, b T) T {
	return a + b
}

// ❌ BAD: Overly restrictive
func badCompare[T int | float64](a, b T) bool {
	return a > b
}

// ✅ GOOD: Use built-in constraints when available
func goodCompare[T Sortable](a, b T) bool {
	return a > b
}

// ==========================================
// Best Practice 2: Minimize Type Parameters
// ==========================================

// ❌ BAD: Too many type parameters
func badProcess[T any, U any, V any, W any](
	input T, transformer func(T) U, validator func(U) V, finalizer func(V) W,
) W {
	return finalizer(validator(transformer(input)))
}

// ✅ GOOD: Consolidate and simplify
func goodProcess[T any, R any](input T, pipeline func(T) R) R {
	return pipeline(input)
}

// ==========================================
// Best Practice 3: Interface Segregation
// ==========================================

// ❌ BAD: Fat interface
type BadContainer[T any] interface {
	Add(T)
	Remove(T) bool
	Contains(T) bool
	Size() int
	Clear()
	ToSlice() []T
	Filter(func(T) bool) BadContainer[T]
	Map(func(T) T) BadContainer[T]
	Reduce(func(T, T) T) T
}

// ✅ GOOD: Small, focused interfaces
type BasicContainer[T any] interface {
	Add(T)
	Size() int
}

type Queryable[T any] interface {
	Contains(T) bool
	ToSlice() []T
}

type Transformable[T any] interface {
	Filter(func(T) bool) []T
	Map(func(T) T) []T
}

// ==========================================
// Best Practice 4: Avoid Generic Structs for Simple Cases
// ==========================================

// ❌ BAD: Overuse of generics for simple wrapper
type BadWrapper[T any] struct {
	value T
}

func (w BadWrapper[T]) Get() T {
	return w.value
}

// ✅ GOOD: Use generics only when you need type flexibility
type EventHandler[T any] struct {
	handlers []func(T)
}

func (eh *EventHandler[T]) Subscribe(handler func(T)) {
	eh.handlers = append(eh.handlers, handler)
}

func (eh *EventHandler[T]) Publish(event T) {
	for _, handler := range eh.handlers {
		handler(event)
	}
}

// ==========================================
// Best Practice 5: Type Inference Friendly Design
// ==========================================

// ✅ GOOD: Design for type inference
func createSlice[T any](items ...T) []T {
	return items // Type can be inferred from usage
}

func createMap[K comparable, V any]() map[K]V {
	return make(map[K]V) // Return type makes inference possible
}

// Helper function with explicit types when needed
func createTypedMap[K comparable, V any](pairs map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range pairs {
		result[k] = v
	}
	return result
}

// ==========================================
// Best Practice 6: Performance Considerations
// ==========================================

// Memory-efficient generic slice operations
type OptimizedSlice[T any] struct {
	data []T
	cap  int
}

func NewOptimizedSlice[T any](capacity int) *OptimizedSlice[T] {
	return &OptimizedSlice[T]{
		data: make([]T, 0, capacity),
		cap:  capacity,
	}
}

func (s *OptimizedSlice[T]) Append(item T) {
	if len(s.data) < s.cap {
		s.data = append(s.data, item)
	} else {
		// Handle overflow - could resize or drop old items
		copy(s.data, s.data[1:])
		s.data[len(s.data)-1] = item
	}
}

func (s *OptimizedSlice[T]) Get(index int) T {
	return s.data[index]
}

func (s *OptimizedSlice[T]) Len() int {
	return len(s.data)
}

// Pre-allocated worker pool pattern
type WorkerPool[T any, R any] struct {
	workers    int
	jobQueue   chan T
	resultChan chan R
	processor  func(T) R
	wg         sync.WaitGroup
}

func NewWorkerPool[T any, R any](workers int, processor func(T) R) *WorkerPool[T, R] {
	return &WorkerPool[T, R]{
		workers:    workers,
		jobQueue:   make(chan T, workers*2), // Buffer to avoid blocking
		resultChan: make(chan R, workers*2),
		processor:  processor,
	}
}

func (wp *WorkerPool[T, R]) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}
}

func (wp *WorkerPool[T, R]) worker() {
	defer wp.wg.Done()
	for job := range wp.jobQueue {
		result := wp.processor(job)
		wp.resultChan <- result
	}
}

func (wp *WorkerPool[T, R]) Submit(job T) {
	wp.jobQueue <- job
}

func (wp *WorkerPool[T, R]) GetResult() R {
	return <-wp.resultChan
}

func (wp *WorkerPool[T, R]) Close() {
	close(wp.jobQueue)
	wp.wg.Wait()
	close(wp.resultChan)
}

// ==========================================
// Best Practice 7: Error Handling Patterns
// ==========================================

// BestPracticeResult type for better error handling (avoiding conflict)
type BestPracticeResult[T any] struct {
	value T
	err   error
}

func BPOk[T any](value T) BestPracticeResult[T] {
	return BestPracticeResult[T]{value: value, err: nil}
}

func BPErr[T any](err error) BestPracticeResult[T] {
	var zero T
	return BestPracticeResult[T]{value: zero, err: err}
}

func (r BestPracticeResult[T]) IsOk() bool {
	return r.err == nil
}

func (r BestPracticeResult[T]) IsErr() bool {
	return r.err != nil
}

func (r BestPracticeResult[T]) Unwrap() (T, error) {
	return r.value, r.err
}

func (r BestPracticeResult[T]) UnwrapOr(defaultValue T) T {
	if r.IsOk() {
		return r.value
	}
	return defaultValue
}

// MapBP result to another type
func MapBP[T, U any](r BestPracticeResult[T], fn func(T) U) BestPracticeResult[U] {
	if r.IsErr() {
		return BPErr[U](r.err)
	}
	return BPOk(fn(r.value))
}

// Chain operations that might fail
func AndThenBP[T, U any](r BestPracticeResult[T], fn func(T) BestPracticeResult[U]) BestPracticeResult[U] {
	if r.IsErr() {
		return BPErr[U](r.err)
	}
	return fn(r.value)
}

// ==========================================
// Best Practice 8: Testing Patterns
// ==========================================

// Testable generic functions with dependency injection
type BPValidator[T any] interface {
	Validate(T) error
}

type BPProcessor[T any] interface {
	Process(T) T
}

type BPPipeline[T any] struct {
	validator BPValidator[T]
	processor BPProcessor[T]
}

func NewBPPipelineWithDeps[T any](validator BPValidator[T], processor BPProcessor[T]) *BPPipeline[T] {
	return &BPPipeline[T]{
		validator: validator,
		processor: processor,
	}
}

func (p *BPPipeline[T]) Execute(input T) BestPracticeResult[T] {
	if err := p.validator.Validate(input); err != nil {
		return BPErr[T](err)
	}
	processed := p.processor.Process(input)
	return BPOk(processed)
}

// Mock implementations for testing
type MockBPValidator[T any] struct {
	shouldFail bool
}

func (m MockBPValidator[T]) Validate(T) error {
	if m.shouldFail {
		return fmt.Errorf("validation failed")
	}
	return nil
}

type MockBPProcessor[T any] struct {
	transform func(T) T
}

func (m MockBPProcessor[T]) Process(input T) T {
	if m.transform != nil {
		return m.transform(input)
	}
	return input
}

// ==========================================
// Best Practice 9: Documentation and Naming
// ==========================================

// ✅ GOOD: Clear naming and documentation
// BPSafeMap provides thread-safe access to a generic map (avoiding conflict).
// Type parameter K must be comparable for use as map key.
// Type parameter V can be any type.
type BPSafeMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

// NewBPSafeMap creates a new thread-safe map.
func NewBPSafeMap[K comparable, V any]() *BPSafeMap[K, V] {
	return &BPSafeMap[K, V]{
		data: make(map[K]V),
	}
}

// Set stores a key-value pair in the map.
func (sm *BPSafeMap[K, V]) Set(key K, value V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

// Get retrieves a value by key, returning the value and whether it exists.
func (sm *BPSafeMap[K, V]) Get(key K) (V, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, exists := sm.data[key]
	return value, exists
}

// ==========================================
// Best Practice 10: Avoiding Common Pitfalls
// ==========================================

// ❌ BAD: Generic method on non-generic type
type BadService struct{}

// This won't compile - methods can't introduce new type parameters
// func (s BadService) Process[T any](data T) T { return data }

// ✅ GOOD: Generic type with methods or generic functions
type GoodService[T any] struct {
	data T
}

func (s GoodService[T]) Process() T {
	return s.data
}

// Or use generic functions
func ProcessData[T any](service interface{}, data T) T {
	// Process the data
	return data
}

// ❌ BAD: Using type assertion with generics unnecessarily
func badTypeAssertion[T any](value interface{}) T {
	return value.(T) // Runtime panic risk
}

// ✅ GOOD: Use type constraints instead
func goodTypeConstraint[T Number](value T) T {
	return value // Compile-time safety
}

// ==========================================
// Performance Benchmarking Examples
// ==========================================

// Benchmark generic vs interface performance
type NumberInterface interface {
	~int | ~float64
}

func genericSum[T NumberInterface](numbers []T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func interfaceSum(numbers []interface{}) float64 {
	var sum float64
	for _, n := range numbers {
		switch v := n.(type) {
		case int:
			sum += float64(v)
		case float64:
			sum += v
		}
	}
	return sum
}

// Memory usage comparison function
func memoryUsageDemo[T any](size int, factory func() T) {
	runtime.GC()
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	// Create slice of generic type
	slice := make([]T, size)
	for i := range slice {
		slice[i] = factory()
	}

	runtime.GC()
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	fmt.Printf("Memory used for %d items of type %s: %d bytes\n",
		size, reflect.TypeOf(slice[0]).String(), m2.Alloc-m1.Alloc)
}

// ==========================================
// Main Example Function
// ==========================================

func runBestPracticesExample() {
	fmt.Println("\n🔸 Type Constraint Best Practices")

	// Good constraint usage
	result1 := goodAdd(10, 20)
	result2 := goodAdd(3.14, 2.86)
	fmt.Printf("Good addition: %d, %.2f\n", result1, result2)

	// Good comparison
	fmt.Printf("Good comparison: %t\n", goodCompare(10, 5))

	fmt.Println("\n🔸 Type Parameter Minimization")

	// Simple pipeline usage
	processedResult := goodProcess("hello", func(s string) string {
		return s + " world!"
	})
	fmt.Printf("Processed result: %s\n", processedResult)

	fmt.Println("\n🔸 Interface Design")

	// Event handler example
	eventHandler := &EventHandler[string]{}
	eventHandler.Subscribe(func(event string) {
		fmt.Printf("Handler 1 received: %s\n", event)
	})
	eventHandler.Subscribe(func(event string) {
		fmt.Printf("Handler 2 received: %s\n", event)
	})

	eventHandler.Publish("Important event!")

	fmt.Println("\n🔸 Type Inference Friendly Design")

	// Type inference examples
	numbers := createSlice(1, 2, 3, 4, 5)
	words := createSlice("hello", "world")
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Words: %v\n", words)

	// Create maps with inference
	intMap := createMap[string, int]()
	intMap["one"] = 1
	intMap["two"] = 2
	fmt.Printf("Int map: %v\n", intMap)

	fmt.Println("\n🔸 Performance Optimizations")

	// Optimized slice usage
	optimizedSlice := NewOptimizedSlice[int](3)
	for i := 0; i < 5; i++ {
		optimizedSlice.Append(i)
		fmt.Printf("Added %d, slice: ", i)
		for j := 0; j < optimizedSlice.Len(); j++ {
			fmt.Printf("%d ", optimizedSlice.Get(j))
		}
		fmt.Println()
	}

	fmt.Println("\n🔸 Worker Pool Pattern")

	// Worker pool demonstration
	pool := NewWorkerPool(3, func(n int) int {
		time.Sleep(10 * time.Millisecond) // Simulate work
		return n * n
	})

	pool.Start()

	// Submit jobs
	for i := 1; i <= 5; i++ {
		pool.Submit(i)
	}

	// Collect results
	fmt.Print("Worker pool results: ")
	for i := 0; i < 5; i++ {
		result := pool.GetResult()
		fmt.Printf("%d ", result)
	}
	fmt.Println()

	pool.Close()

	fmt.Println("\n🔸 Error Handling with Result Type")

	// Result type examples
	successResult := BPOk("success")
	errorResult := BPErr[string](fmt.Errorf("something went wrong"))

	fmt.Printf("Success result is ok: %t\n", successResult.IsOk())
	fmt.Printf("Error result is ok: %t\n", errorResult.IsOk())

	// Map operations
	mapped := MapBP(successResult, func(s string) int { return len(s) })
	fmt.Printf("Mapped success: %v\n", mapped.UnwrapOr(0))

	mappedErr := MapBP(errorResult, func(s string) int { return len(s) })
	fmt.Printf("Mapped error: %v\n", mappedErr.UnwrapOr(-1))

	// Chain operations
	chained := AndThenBP(successResult, func(s string) BestPracticeResult[int] {
		if len(s) > 0 {
			return BPOk(len(s))
		}
		return BPErr[int](fmt.Errorf("empty string"))
	})

	if value, err := chained.Unwrap(); err == nil {
		fmt.Printf("Chained result: %d\n", value)
	} else {
		fmt.Printf("Chained error: %v\n", err)
	}

	fmt.Println("\n🔸 Testing with Dependency Injection")

	// Create pipeline with mocks
	validator := MockBPValidator[string]{shouldFail: false}
	processor := MockBPProcessor[string]{
		transform: func(s string) string { return "processed_" + s },
	}

	pipeline := NewBPPipelineWithDeps[string](validator, processor)
	result := pipeline.Execute("test_data")

	if result.IsOk() {
		value, _ := result.Unwrap()
		fmt.Printf("Pipeline result: %s\n", value)
	}

	// Test with failing validator
	failingValidator := MockBPValidator[string]{shouldFail: true}
	failingPipeline := NewBPPipelineWithDeps[string](failingValidator, processor)
	failResult := failingPipeline.Execute("test_data")

	if failResult.IsErr() {
		_, err := failResult.Unwrap()
		fmt.Printf("Pipeline error: %v\n", err)
	}

	fmt.Println("\n🔸 Thread-Safe Generic Collections")

	// Safe map usage
	safeMap := NewBPSafeMap[string, int]()
	safeMap.Set("count", 42)
	safeMap.Set("total", 100)

	if value, exists := safeMap.Get("count"); exists {
		fmt.Printf("Safe map value: %d\n", value)
	}

	fmt.Println("\n🔸 Performance Comparison")

	// Prepare data for benchmarking
	size := 10000
	intNumbers := make([]int, size)
	interfaceNumbers := make([]interface{}, size)

	for i := 0; i < size; i++ {
		intNumbers[i] = i
		interfaceNumbers[i] = i
	}

	// Benchmark generic version
	start := time.Now()
	genericResult := genericSum(intNumbers)
	genericTime := time.Since(start)

	// Benchmark interface version
	start = time.Now()
	interfaceResult := interfaceSum(interfaceNumbers)
	interfaceTime := time.Since(start)

	fmt.Printf("Generic sum (%v): %d\n", genericTime, genericResult)
	fmt.Printf("Interface sum (%v): %.0f\n", interfaceTime, interfaceResult)
	fmt.Printf("Generic is %.2fx faster\n", float64(interfaceTime)/float64(genericTime))

	fmt.Println("\n🔸 Memory Usage Analysis")

	// Compare memory usage of different types
	memoryUsageDemo(1000, func() int { return 42 })
	memoryUsageDemo(1000, func() string { return "hello" })
	memoryUsageDemo(1000, func() interface{} { return 42 })

	fmt.Println("\n📋 Generic Programming Best Practices Summary:")
	fmt.Println("1. ✅ Use specific constraints, avoid 'any' when possible")
	fmt.Println("2. ✅ Minimize type parameters, prefer composability")
	fmt.Println("3. ✅ Design small, focused interfaces")
	fmt.Println("4. ✅ Make functions type-inference friendly")
	fmt.Println("5. ✅ Consider performance implications")
	fmt.Println("6. ✅ Use Result types for better error handling")
	fmt.Println("7. ✅ Design for testability with dependency injection")
	fmt.Println("8. ✅ Document type constraints clearly")
	fmt.Println("9. ✅ Avoid common pitfalls (type assertions, method generics)")
	fmt.Println("10. ✅ Benchmark generic vs non-generic implementations")

	fmt.Println("\n✅ Best practices examples completed!")
}
