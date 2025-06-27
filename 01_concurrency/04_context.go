package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ContextExamples runs all context examples
func ContextExamples() {
	fmt.Println("=== Context Management Examples ===")

	// Example 1: Basic context
	basicContext()

	// Example 2: Timeout control
	timeoutContext()

	// Example 3: Cancel control
	cancelContext()

	// Example 4: Deadline
	deadlineContext()

	// Example 5: Value passing
	valueContext()

	// Example 6: Nested context
	nestedContext()

	// Example 7: Concurrent operation control
	concurrentContext()

	// Example 8: HTTP request context
	httpContext()

	// Example 9: Common mistakes examples
	contextMistakes()
}

// Example 1: Basic context
func basicContext() {
	fmt.Println("\n--- Example 1: Basic context ---")

	// Create root context
	ctx := context.Background()
	fmt.Printf("Root context: %v\n", ctx)

	// Check if context is cancelled
	select {
	case <-ctx.Done():
		fmt.Println("Context is cancelled")
	default:
		fmt.Println("Context is not cancelled")
	}

	// Get context error information
	err := ctx.Err()
	if err != nil {
		fmt.Printf("Context error: %v\n", err)
	} else {
		fmt.Println("Context has no error")
	}
}

// Example 2: Timeout control
func timeoutContext() {
	fmt.Println("\n--- Example 2: Timeout control ---")

	// Create context with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Starting long-running task...")

	// Simulate long-running task
	go func() {
		time.Sleep(10 * time.Second) // Longer than timeout
		fmt.Println("Task completed")
	}()

	// Wait for context cancellation
	select {
	case <-ctx.Done():
		fmt.Printf("Task cancelled: %v\n", ctx.Err())
	case <-time.After(6 * time.Second):
		fmt.Println("Wait timeout")
	}

	// Short task example
	fmt.Println("\nShort task example:")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel2()

	go func() {
		time.Sleep(1 * time.Second) // Shorter than timeout
		fmt.Println("Short task completed")
	}()

	select {
	case <-ctx2.Done():
		fmt.Printf("Short task cancelled: %v\n", ctx2.Err())
	case <-time.After(2 * time.Second):
		fmt.Println("Short task completed normally")
	}
}

// Example 3: Cancel control
func cancelContext() {
	fmt.Println("\n--- Example 3: Cancel control ---")

	// Create cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// Start worker goroutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Worker goroutine received cancel signal: %v\n", ctx.Err())
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Main goroutine waits for a while then cancels
	time.Sleep(2 * time.Second)
	fmt.Println("Main goroutine sends cancel signal")
	cancel()

	// Wait for worker goroutine to exit
	time.Sleep(100 * time.Millisecond)

	// Demonstrate multiple cancellations
	fmt.Println("\nDemonstrating multiple cancellations:")
	ctx2, cancel2 := context.WithCancel(context.Background())
	defer cancel2()

	cancel2() // First cancellation
	cancel2() // Second cancellation (safe)

	select {
	case <-ctx2.Done():
		fmt.Printf("Context cancelled: %v\n", ctx2.Err())
	default:
		fmt.Println("Context not cancelled")
	}
}

// Example 4: Deadline
func deadlineContext() {
	fmt.Println("\n--- Example 4: Deadline ---")

	// Set deadline to 5 seconds from now
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	fmt.Printf("Deadline: %v\n", deadline)

	// Start task
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Task cancelled due to deadline reached: %v\n", ctx.Err())
				return
			default:
				fmt.Printf("Task running, remaining time: %v\n", time.Until(deadline))
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Wait for context cancellation
	<-ctx.Done()
	fmt.Printf("Main goroutine received cancel signal: %v\n", ctx.Err())
}

// Example 5: Value passing
func valueContext() {
	fmt.Println("\n--- Example 5: Value passing ---")

	// Create context with values
	ctx := context.WithValue(context.Background(), "user_id", "12345")
	ctx = context.WithValue(ctx, "request_id", "req_67890")

	// Access values in goroutine
	go func(ctx context.Context) {
		userID := ctx.Value("user_id")
		requestID := ctx.Value("request_id")
		fmt.Printf("Getting values in goroutine: user_id=%v, request_id=%v\n", userID, requestID)

		// Access non-existent value
		nonExistent := ctx.Value("non_existent")
		fmt.Printf("Non-existent value: %v\n", nonExistent)
	}(ctx)

	time.Sleep(100 * time.Millisecond)

	// Demonstrate value passing chain
	fmt.Println("\nValue passing chain:")
	ctx1 := context.WithValue(context.Background(), "level", "1")
	ctx2 := context.WithValue(ctx1, "level", "2")
	ctx3 := context.WithValue(ctx2, "level", "3")

	fmt.Printf("ctx1 level: %v\n", ctx1.Value("level"))
	fmt.Printf("ctx2 level: %v\n", ctx2.Value("level"))
	fmt.Printf("ctx3 level: %v\n", ctx3.Value("level"))

	// Access parent context value
	fmt.Printf("ctx3 accessing parent context value: %v\n", ctx3.Value("level"))
}

// Example 6: Nested context
func nestedContext() {
	fmt.Println("\n--- Example 6: Nested context ---")

	// Create root context
	rootCtx := context.Background()

	// Add timeout
	timeoutCtx, cancel1 := context.WithTimeout(rootCtx, 10*time.Second)
	defer cancel1()

	// Add value
	valueCtx := context.WithValue(timeoutCtx, "operation", "database_query")

	// Add deadline
	deadlineCtx, cancel2 := context.WithDeadline(valueCtx, time.Now().Add(5*time.Second))
	defer cancel2()

	// Start worker goroutine
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Nested context cancelled: %v\n", ctx.Err())
				return
			default:
				operation := ctx.Value("operation")
				fmt.Printf("Executing operation: %v\n", operation)
				time.Sleep(1 * time.Second)
			}
		}
	}(deadlineCtx)

	// Wait for context cancellation
	<-deadlineCtx.Done()
	fmt.Printf("Nested context cancelled: %v\n", deadlineCtx.Err())
}

// Example 7: Concurrent operation control
func concurrentContext() {
	fmt.Println("\n--- Example 7: Concurrent operation control ---")

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	// Start multiple concurrent tasks
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Task %d cancelled: %v\n", id, ctx.Err())
					return
				default:
					fmt.Printf("Task %d working...\n", id)
					time.Sleep(time.Duration(id) * 500 * time.Millisecond)
				}
			}
		}(i)
	}

	// Wait for all tasks to complete or be cancelled
	wg.Wait()
	fmt.Println("All tasks completed")

	// Demonstrate resource cleanup
	fmt.Println("\nDemonstrating resource cleanup:")
	ctx2, cancel2 := context.WithCancel(context.Background())

	go func() {
		defer fmt.Println("Resources cleaned up")

		select {
		case <-ctx2.Done():
			fmt.Println("Received cancel signal, starting resource cleanup")
		}
	}()

	time.Sleep(100 * time.Millisecond)
	cancel2()
	time.Sleep(100 * time.Millisecond)
}

// Example 8: HTTP request context
func httpContext() {
	fmt.Println("\n--- Example 8: HTTP request context ---")

	// Simulate HTTP request context
	ctx := context.WithValue(context.Background(), "request_id", "http_req_123")

	// Add timeout
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// Simulate database query
	go func(ctx context.Context) {
		requestID := ctx.Value("request_id")
		fmt.Printf("Starting database query, request ID: %v\n", requestID)

		// Simulate database operation
		select {
		case <-time.After(3 * time.Second): // Simulate slow query
			fmt.Println("Database query completed")
		case <-ctx.Done():
			fmt.Printf("Database query cancelled: %v\n", ctx.Err())
		}
	}(ctx)

	// Simulate external API call
	go func(ctx context.Context) {
		requestID := ctx.Value("request_id")
		fmt.Printf("Starting external API call, request ID: %v\n", requestID)

		select {
		case <-time.After(1 * time.Second): // Simulate fast API call
			fmt.Println("External API call completed")
		case <-ctx.Done():
			fmt.Printf("External API call cancelled: %v\n", ctx.Err())
		}
	}(ctx)

	// Wait for context cancellation
	<-ctx.Done()
	fmt.Printf("HTTP request completed: %v\n", ctx.Err())
}

// Example 9: Common mistakes examples
func contextMistakes() {
	fmt.Println("\n--- Example 9: Common mistakes examples ---")

	fmt.Println("Mistake 1: Forgetting to call cancel function")
	_, _ = context.WithCancel(context.Background())
	// Forgetting to call cancel() may cause goroutine leaks
	// Intentionally not calling cancel here to demonstrate the mistake

	// Correct approach: use defer
	ctx2, cancel2 := context.WithCancel(context.Background())
	defer cancel2() // Ensure cancel is called

	fmt.Println("Mistake 2: Creating child context on already cancelled context")
	cancel2() // Cancel parent context

	// Create child context on already cancelled context
	childCtx, childCancel := context.WithTimeout(ctx2, 1*time.Second)
	defer childCancel()

	select {
	case <-childCtx.Done():
		fmt.Printf("Child context immediately cancelled: %v\n", childCtx.Err())
	default:
		fmt.Println("Child context not cancelled")
	}

	fmt.Println("\nMistake 3: Passing nil context")
	// Don't pass nil context
	// someFunction(nil) // Wrong

	// Correct approach: use context.Background()
	// someFunction(context.Background()) // Correct

	fmt.Println("\nMistake 4: Storing mutable reference types in context")
	type User struct {
		Name string
	}

	user := &User{Name: "Alice"}
	ctx3 := context.WithValue(context.Background(), "user", user)

	// Modifying user affects the value in context
	user.Name = "Bob"

	retrievedUser := ctx3.Value("user").(*User)
	fmt.Printf("User name retrieved from context: %s\n", retrievedUser.Name)

	fmt.Println("\nMistake 5: Forgetting to handle context cancellation in select")
	ctx5, cancel5 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel5()

	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second) // Longer than context timeout
		ch <- "data"
	}()

	// Wrong select - may block forever
	select {
	case data := <-ch:
		fmt.Printf("Received data: %s\n", data)
		// Forgot to handle context cancellation
	}

	fmt.Println("\nCorrect approach: Handle context cancellation in select")
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "data"
	}()

	select {
	case data := <-ch2:
		fmt.Printf("Received data: %s\n", data)
	case <-ctx5.Done():
		fmt.Printf("Operation cancelled: %v\n", ctx5.Err())
	}
}
