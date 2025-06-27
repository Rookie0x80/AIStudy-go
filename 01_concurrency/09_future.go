package main

import (
	"fmt"
	"sync"
	"time"
)

// Future/Promise pattern examples

// Future represents the result of an asynchronous computation
type Future struct {
	result interface{}
	err    error
	done   chan struct{}
	mu     sync.RWMutex
}

// NewFuture creates a new Future
func NewFuture() *Future {
	return &Future{
		done: make(chan struct{}),
	}
}

// SetResult sets the Future result
func (f *Future) SetResult(result interface{}) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.result = result
	close(f.done)
}

// SetError sets the Future error
func (f *Future) SetError(err error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.err = err
	close(f.done)
}

// Get gets the Future result, blocking until completion
func (f *Future) Get() (interface{}, error) {
	<-f.done
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.result, f.err
}

// GetWithTimeout gets result with timeout
func (f *Future) GetWithTimeout(timeout time.Duration) (interface{}, error) {
	select {
	case <-f.done:
		f.mu.RLock()
		defer f.mu.RUnlock()
		return f.result, f.err
	case <-time.After(timeout):
		return nil, fmt.Errorf("timeout after %v", timeout)
	}
}

// IsDone checks if Future is completed
func (f *Future) IsDone() bool {
	select {
	case <-f.done:
		return true
	default:
		return false
	}
}

// Promise represents a Future that can set results
type Promise struct {
	*Future
}

// NewPromise creates a new Promise
func NewPromise() *Promise {
	return &Promise{
		Future: NewFuture(),
	}
}

// FutureExamples runs all Future/Promise examples
func FutureExamples() {
	fmt.Println("=== Future/Promise Pattern Examples ===")

	// Example 1: Basic Future
	basicFutureExample()

	// Example 2: Promise usage
	promiseExample()

	// Example 3: Asynchronous computation
	asyncComputationExample()

	// Example 4: Timeout handling
	futureTimeoutExample()

	// Example 5: Compose Future
	composeFutureExample()

	// Example 6: Parallel execution
	parallelFutureExample()

	// Example 7: Error handling
	errorHandlingExample()

	// Example 8: Comprehensive example
	comprehensiveFutureExample()
}

// Example 1: Basic Future
func basicFutureExample() {
	fmt.Println("\n--- Example 1: Basic Future ---")

	future := NewFuture()

	// Start asynchronous computation
	go func() {
		fmt.Println("Starting asynchronous computation...")
		time.Sleep(500 * time.Millisecond) // Simulate computation
		result := "Computation result: 42"
		fmt.Println("Computation completed")
		future.SetResult(result)
	}()

	// Main thread continues with other work
	fmt.Println("Main thread continues execution...")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main thread waiting for result...")

	// Get result
	result, err := future.Get()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Received result: %v\n", result)
	}
}

// Example 2: Promise usage
func promiseExample() {
	fmt.Println("\n--- Example 2: Promise Usage ---")

	promise := NewPromise()

	// Start asynchronous task
	go func() {
		fmt.Println("Promise: starting processing...")
		time.Sleep(300 * time.Millisecond)

		// Simulate success or failure
		if time.Now().Unix()%2 == 0 {
			promise.SetResult("Processing successful")
		} else {
			promise.SetError(fmt.Errorf("Processing failed"))
		}
	}()

	// Check status
	for i := 0; i < 5; i++ {
		if promise.IsDone() {
			fmt.Println("Promise completed")
			break
		}
		fmt.Printf("Promise status check %d: not completed\n", i+1)
		time.Sleep(100 * time.Millisecond)
	}

	// Get result
	result, err := promise.Get()
	if err != nil {
		fmt.Printf("Promise error: %v\n", err)
	} else {
		fmt.Printf("Promise result: %v\n", result)
	}
}

// Example 3: Asynchronous computation
func asyncComputationExample() {
	fmt.Println("\n--- Example 3: Asynchronous Computation ---")

	// Create multiple asynchronous computations
	futures := make([]*Future, 3)

	for i := 0; i < 3; i++ {
		future := NewFuture()
		futures[i] = future

		go func(id int, f *Future) {
			fmt.Printf("Calculator %d: starting computation\n", id)
			time.Sleep(time.Duration(id+1) * 200 * time.Millisecond)
			result := fmt.Sprintf("Calculator %d result: %d", id, id*id)
			fmt.Printf("Calculator %d: computation completed\n", id)
			f.SetResult(result)
		}(i, future)
	}

	// Wait for all computations to complete
	fmt.Println("Waiting for all computations to complete...")
	for i, future := range futures {
		result, err := future.Get()
		if err != nil {
			fmt.Printf("Calculator %d error: %v\n", i, err)
		} else {
			fmt.Printf("Calculator %d result: %v\n", i, result)
		}
	}
}

// Example 4: Timeout handling
func futureTimeoutExample() {
	fmt.Println("\n--- Example 4: Timeout Handling ---")

	future := NewFuture()

	// Start long-running task
	go func() {
		fmt.Println("Long-running task: starting...")
		time.Sleep(2 * time.Second) // Simulate long-running task
		future.SetResult("Long-running task completed")
	}()

	// Try to get result with timeout
	fmt.Println("Attempting to get result (timeout 1 second)...")
	result, err := future.GetWithTimeout(1 * time.Second)
	if err != nil {
		fmt.Printf("Timeout error: %v\n", err)

		// Continue waiting
		fmt.Println("Continuing to wait...")
		result, err = future.Get()
		if err != nil {
			fmt.Printf("Final error: %v\n", err)
		} else {
			fmt.Printf("Final result: %v\n", result)
		}
	} else {
		fmt.Printf("Result obtained in time: %v\n", result)
	}
}

// Example 5: Compose Future
func composeFutureExample() {
	fmt.Println("\n--- Example 5: Compose Future ---")

	// First Future: Get user ID
	userIDFuture := NewFuture()
	go func() {
		time.Sleep(100 * time.Millisecond)
		userIDFuture.SetResult(12345)
	}()

	// Second Future: Get user info based on user ID
	userInfoFuture := NewFuture()
	go func() {
		userID, err := userIDFuture.Get()
		if err != nil {
			userInfoFuture.SetError(err)
			return
		}

		time.Sleep(200 * time.Millisecond)
		userInfo := fmt.Sprintf("User info: ID=%v, Name=John, Age=30", userID)
		userInfoFuture.SetResult(userInfo)
	}()

	// Third Future: Get orders based on user info
	ordersFuture := NewFuture()
	go func() {
		userInfo, err := userInfoFuture.Get()
		if err != nil {
			ordersFuture.SetError(err)
			return
		}

		time.Sleep(150 * time.Millisecond)
		orders := fmt.Sprintf("Order info: %v order list", userInfo)
		ordersFuture.SetResult(orders)
	}()

	// Get final result
	result, err := ordersFuture.Get()
	if err != nil {
		fmt.Printf("Compose Future error: %v\n", err)
	} else {
		fmt.Printf("Compose Future result: %v\n", result)
	}
}

// Example 6: Parallel execution
func parallelFutureExample() {
	fmt.Println("\n--- Example 6: Parallel Execution ---")

	// Create multiple parallel tasks
	tasks := []func() (interface{}, error){
		func() (interface{}, error) {
			time.Sleep(300 * time.Millisecond)
			return "Task 1 completed", nil
		},
		func() (interface{}, error) {
			time.Sleep(200 * time.Millisecond)
			return "Task 2 completed", nil
		},
		func() (interface{}, error) {
			time.Sleep(400 * time.Millisecond)
			return "Task 3 completed", nil
		},
		func() (interface{}, error) {
			time.Sleep(100 * time.Millisecond)
			return "Task 4 completed", nil
		},
	}

	// Execute all tasks in parallel
	futures := make([]*Future, len(tasks))
	for i, task := range tasks {
		future := NewFuture()
		futures[i] = future

		go func(f *Future, t func() (interface{}, error)) {
			result, err := t()
			if err != nil {
				f.SetError(err)
			} else {
				f.SetResult(result)
			}
		}(future, task)
	}

	// Wait for all tasks to complete
	fmt.Println("Waiting for all parallel tasks to complete...")
	var wg sync.WaitGroup
	wg.Add(len(futures))

	for i, future := range futures {
		go func(id int, f *Future) {
			defer wg.Done()
			result, err := f.Get()
			if err != nil {
				fmt.Printf("Task %d error: %v\n", id+1, err)
			} else {
				fmt.Printf("Task %d result: %v\n", id+1, result)
			}
		}(i, future)
	}

	wg.Wait()
	fmt.Println("All parallel tasks completed")
}

// Example 7: Error handling
func errorHandlingExample() {
	fmt.Println("\n--- Example 7: Error Handling ---")

	// Create Future that may error
	errorFuture := NewFuture()
	go func() {
		time.Sleep(100 * time.Millisecond)
		// Simulate error
		errorFuture.SetError(fmt.Errorf("Simulated error"))
	}()

	// Create successful Future
	successFuture := NewFuture()
	go func() {
		time.Sleep(150 * time.Millisecond)
		successFuture.SetResult("Successful result")
	}()

	// Handle error Future
	result, err := errorFuture.Get()
	if err != nil {
		fmt.Printf("Error Future: %v\n", err)
		// Error recovery: use default value
		result = "Default value"
		fmt.Printf("Using default value: %v\n", result)
	} else {
		fmt.Printf("Error Future result: %v\n", result)
	}

	// Handle successful Future
	result, err = successFuture.Get()
	if err != nil {
		fmt.Printf("Success Future error: %v\n", err)
	} else {
		fmt.Printf("Success Future result: %v\n", result)
	}

	// Demonstrate error propagation
	fmt.Println("\nDemonstrating error propagation:")
	chainFuture := NewFuture()
	go func() {
		// First operation fails
		time.Sleep(50 * time.Millisecond)
		chainFuture.SetError(fmt.Errorf("Chained operation failed"))
	}()

	// Try to continue processing
	result, err = chainFuture.Get()
	if err != nil {
		fmt.Printf("Chained operation error: %v\n", err)
		// Error handling: log, retry, etc.
		fmt.Println("Executing error handling logic...")
	} else {
		fmt.Printf("Chained operation result: %v\n", result)
	}
}

// Example 8: Comprehensive example
func comprehensiveFutureExample() {
	fmt.Println("\n--- Example 8: Comprehensive Example ---")

	// Simulate a complex asynchronous processing flow
	type User struct {
		ID   int
		Name string
	}

	type Order struct {
		ID     int
		UserID int
		Items  []string
		Total  float64
	}

	// Step 1: Get user info
	userFuture := NewFuture()
	go func() {
		fmt.Println("Step 1: Getting user info...")
		time.Sleep(200 * time.Millisecond)
		user := User{ID: 123, Name: "John"}
		userFuture.SetResult(user)
	}()

	// Step 2: Get user orders
	ordersFuture := NewFuture()
	go func() {
		user, err := userFuture.Get()
		if err != nil {
			ordersFuture.SetError(err)
			return
		}

		fmt.Printf("Step 2: Getting orders for user %v...\n", user)
		time.Sleep(300 * time.Millisecond)
		orders := []Order{
			{ID: 1, UserID: 123, Items: []string{"Apple", "Banana"}, Total: 15.50},
			{ID: 2, UserID: 123, Items: []string{"Orange"}, Total: 8.00},
		}
		ordersFuture.SetResult(orders)
	}()

	// Step 3: Calculate order statistics
	statsFuture := NewFuture()
	go func() {
		orders, err := ordersFuture.Get()
		if err != nil {
			statsFuture.SetError(err)
			return
		}

		fmt.Println("Step 3: Calculating order statistics...")
		time.Sleep(150 * time.Millisecond)

		orderList := orders.([]Order)
		totalOrders := len(orderList)
		var totalAmount float64
		for _, order := range orderList {
			totalAmount += order.Total
		}

		stats := fmt.Sprintf("Order statistics: Total orders=%d, Total amount=%.2f", totalOrders, totalAmount)
		statsFuture.SetResult(stats)
	}()

	// Step 4: Generate report
	reportFuture := NewFuture()
	go func() {
		user, _ := userFuture.Get()
		stats, err := statsFuture.Get()
		if err != nil {
			reportFuture.SetError(err)
			return
		}

		fmt.Println("Step 4: Generating report...")
		time.Sleep(100 * time.Millisecond)

		report := fmt.Sprintf("User report: %v, %v", user, stats)
		reportFuture.SetResult(report)
	}()

	// Get final report
	fmt.Println("Waiting for report generation...")
	result, err := reportFuture.Get()
	if err != nil {
		fmt.Printf("Report generation error: %v\n", err)
	} else {
		fmt.Printf("Final report: %v\n", result)
	}

	// Demonstrate timeout handling
	fmt.Println("\nDemonstrating timeout handling:")
	slowFuture := NewFuture()
	go func() {
		time.Sleep(3 * time.Second)
		slowFuture.SetResult("Slow task completed")
	}()

	result, err = slowFuture.GetWithTimeout(1 * time.Second)
	if err != nil {
		fmt.Printf("Slow task timeout: %v\n", err)
	} else {
		fmt.Printf("Slow task result: %v\n", result)
	}
}
