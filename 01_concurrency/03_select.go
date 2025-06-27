package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SelectExamples runs all select examples
func SelectExamples() {
	fmt.Println("=== Select Multiplexing Examples ===")

	// Example 1: Basic select
	basicSelect()

	// Example 2: Multi-channel listening
	multiChannelSelect()

	// Example 3: Timeout control
	timeoutSelect()

	// Example 4: Non-blocking operations
	nonBlockingSelect()

	// Example 5: Random selection
	randomSelect()

	// Example 6: Loop select
	loopSelect()

	// Example 7: Select with default
	selectWithDefault()

	// Example 8: Dynamic channel selection
	dynamicChannelSelect()

	// Example 9: Common mistakes examples
	selectMistakes()
}

// Example 1: Basic select
func basicSelect() {
	fmt.Println("\n--- Example 1: Basic select ---")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start two goroutines to send data
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from channel1"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "from channel2"
	}()

	// Use select to listen to two channels
	select {
	case msg1 := <-ch1:
		fmt.Printf("Received: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("Received: %s\n", msg2)
	}
}

// Example 2: Multi-channel listening
func multiChannelSelect() {
	fmt.Println("\n--- Example 2: Multi-channel listening ---")

	ch1 := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan bool)

	// Start multiple senders
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(100 * time.Millisecond)
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		time.Sleep(150 * time.Millisecond)
		ch2 <- "hello"
		time.Sleep(100 * time.Millisecond)
		ch2 <- "world"
		close(ch2)
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch3 <- true
		close(ch3)
	}()

	// Listen to multiple channels until all are closed
	ch1Closed, ch2Closed, ch3Closed := false, false, false

	for !ch1Closed || !ch2Closed || !ch3Closed {
		select {
		case value, ok := <-ch1:
			if !ok {
				ch1Closed = true
				fmt.Println("ch1 closed")
			} else {
				fmt.Printf("Received from ch1: %d\n", value)
			}
		case value, ok := <-ch2:
			if !ok {
				ch2Closed = true
				fmt.Println("ch2 closed")
			} else {
				fmt.Printf("Received from ch2: %s\n", value)
			}
		case value, ok := <-ch3:
			if !ok {
				ch3Closed = true
				fmt.Println("ch3 closed")
			} else {
				fmt.Printf("Received from ch3: %t\n", value)
			}
		}
	}
}

// Example 3: Timeout control
func timeoutSelect() {
	fmt.Println("\n--- Example 3: Timeout control ---")

	// Simulate slow operation
	slowOperation := make(chan string)

	go func() {
		// Random delay 0-300ms
		delay := time.Duration(rand.Intn(300)) * time.Millisecond
		time.Sleep(delay)
		slowOperation <- "operation completed"
	}()

	// Set timeout
	timeout := time.After(200 * time.Millisecond)

	select {
	case result := <-slowOperation:
		fmt.Printf("Operation successful: %s\n", result)
	case <-timeout:
		fmt.Println("Operation timeout")
	}

	// Use context-style timeout
	fmt.Println("\nUsing time.After timeout pattern:")
	done := make(chan struct{})

	go func() {
		time.Sleep(150 * time.Millisecond)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Operation completed before timeout")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Operation timeout")
	}
}

// Example 4: Non-blocking operations
func nonBlockingSelect() {
	fmt.Println("\n--- Example 4: Non-blocking operations ---")

	ch := make(chan int, 1)

	// Non-blocking send
	select {
	case ch <- 1:
		fmt.Println("Send successful")
	default:
		fmt.Println("Send failed, channel is full")
	}

	// Try to send again
	select {
	case ch <- 2:
		fmt.Println("Send successful")
	default:
		fmt.Println("Send failed, channel is full")
	}

	// Non-blocking receive
	select {
	case value := <-ch:
		fmt.Printf("Receive successful: %d\n", value)
	default:
		fmt.Println("Receive failed, channel is empty")
	}

	// Receive again
	select {
	case value := <-ch:
		fmt.Printf("Receive successful: %d\n", value)
	default:
		fmt.Println("Receive failed, channel is empty")
	}
}

// Example 5: Random selection
func randomSelect() {
	fmt.Println("\n--- Example 5: Random selection ---")

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	// Start multiple senders, send simultaneously
	go func() {
		ch1 <- "A"
	}()

	go func() {
		ch2 <- "B"
	}()

	go func() {
		ch3 <- "C"
	}()

	// select will randomly choose an available case
	select {
	case msg := <-ch1:
		fmt.Printf("Randomly selected: %s\n", msg)
	case msg := <-ch2:
		fmt.Printf("Randomly selected: %s\n", msg)
	case msg := <-ch3:
		fmt.Printf("Randomly selected: %s\n", msg)
	}

	// Demonstrate multiple random selections
	fmt.Println("\nMultiple random selections:")
	for i := 0; i < 5; i++ {
		ch := make(chan string)
		go func(id int) {
			ch <- fmt.Sprintf("message%d", id)
		}(i)

		select {
		case msg := <-ch:
			fmt.Printf("Selection %d: %s\n", i+1, msg)
		}
	}
}

// Example 6: Loop select
func loopSelect() {
	fmt.Println("\n--- Example 6: Loop select ---")

	ch := make(chan int)
	done := make(chan struct{})

	// Start sender
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(100 * time.Millisecond)
			ch <- i
		}
		close(done)
	}()

	// Loop select until completion signal is received
	for {
		select {
		case value := <-ch:
			fmt.Printf("Received value: %d\n", value)
		case <-done:
			fmt.Println("Received completion signal, exiting loop")
			return
		}
	}
}

// Example 7: Select with default
func selectWithDefault() {
	fmt.Println("\n--- Example 7: Select with default ---")

	ch := make(chan int)

	// Select without default will block
	fmt.Println("Select without default:")
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- 1
	}()

	select {
	case value := <-ch:
		fmt.Printf("Received value: %d\n", value)
	}

	// Select with default won't block
	fmt.Println("\nSelect with default:")
	select {
	case value := <-ch:
		fmt.Printf("Received value: %d\n", value)
	default:
		fmt.Println("No data to read, executing default")
	}

	// Use default to implement non-blocking operations
	fmt.Println("\nUsing default to implement non-blocking operations:")
	ch2 := make(chan int, 1)

	// Non-blocking send
	select {
	case ch2 <- 1:
		fmt.Println("Send successful")
	default:
		fmt.Println("Send failed")
	}

	// Non-blocking receive
	select {
	case value := <-ch2:
		fmt.Printf("Receive successful: %d\n", value)
	default:
		fmt.Println("Receive failed")
	}
}

// Example 8: Dynamic channel selection
func dynamicChannelSelect() {
	fmt.Println("\n--- Example 8: Dynamic channel selection ---")

	// Create multiple channels
	channels := make([]chan int, 3)
	for i := range channels {
		channels[i] = make(chan int)
	}

	// Start senders
	for i, ch := range channels {
		go func(id int, c chan int) {
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			c <- id + 1
		}(i, ch)
	}

	// Dynamically build select cases
	fmt.Println("Dynamically select first available channel:")
	select {
	case value := <-channels[0]:
		fmt.Printf("Received from channels[0]: %d\n", value)
	case value := <-channels[1]:
		fmt.Printf("Received from channels[1]: %d\n", value)
	case value := <-channels[2]:
		fmt.Printf("Received from channels[2]: %d\n", value)
	}

	// Use reflection to implement dynamic select (advanced usage)
	fmt.Println("\nUsing reflection to implement dynamic select:")
	dynamicSelectWithReflection(channels)
}

// Use reflection to implement dynamic select (simplified version)
func dynamicSelectWithReflection(channels []chan int) {
	// Simplified implementation here, actual usage requires reflect package
	fmt.Println("Note: Complete reflection implementation requires reflect package")
	fmt.Println("This shows the concept, actual projects can use reflect.Select")

	// Simulate dynamic selection
	for i, ch := range channels {
		select {
		case value := <-ch:
			fmt.Printf("Dynamic selection received from channels[%d]: %d\n", i, value)
			return
		default:
			continue
		}
	}
}

// Example 9: Common mistakes examples
func selectMistakes() {
	fmt.Println("\n--- Example 9: Common mistakes examples ---")

	fmt.Println("Mistake 1: Using nil channel in select")
	var ch chan int // nil channel

	select {
	case value := <-ch:
		fmt.Printf("Received: %d\n", value)
	default:
		fmt.Println("Nil channel will be ignored")
	}

	fmt.Println("\nMistake 2: Forgetting to handle channel closure")
	ch2 := make(chan int)
	go func() {
		ch2 <- 1
		close(ch2)
	}()

	select {
	case value := <-ch2:
		fmt.Printf("Received: %d\n", value)
	}

	// Should check if channel is closed
	select {
	case value, ok := <-ch2:
		if !ok {
			fmt.Println("Channel is closed")
		} else {
			fmt.Printf("Received: %d\n", value)
		}
	default:
		fmt.Println("Channel is empty")
	}

	fmt.Println("\nMistake 3: Reusing the same channel multiple times in select")
	ch3 := make(chan int)
	go func() {
		ch3 <- 1
		ch3 <- 2
	}()

	// This is legal but may not be expected behavior
	select {
	case value := <-ch3:
		fmt.Printf("First received: %d\n", value)
	case value := <-ch3:
		fmt.Printf("Second received: %d\n", value)
	}

	fmt.Println("\nMistake 4: Forgetting to handle timeout in select")
	slowCh := make(chan string)
	go func() {
		time.Sleep(300 * time.Millisecond)
		slowCh <- "slow operation completed"
	}()

	// No timeout handling, may block forever
	select {
	case result := <-slowCh:
		fmt.Printf("Operation completed: %s\n", result)
		// Forgot to add timeout case
	}

	fmt.Println("\nCorrect approach: Add timeout handling")
	select {
	case result := <-slowCh:
		fmt.Printf("Operation completed: %s\n", result)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Operation timeout")
	}
}
