package main

import (
	"fmt"
	"sync"
	"time"
)

// ChannelsExamples runs all channels examples
func ChannelsExamples() {
	fmt.Println("=== Channels Communication Examples ===")

	// Example 1: Basic channel communication
	basicChannel()

	// Example 2: Buffered channel
	bufferedChannel()

	// Example 3: Directional channel
	directionalChannel()

	// Example 4: Channel close
	channelClose()

	// Example 5: Select statement
	selectExample()

	// Example 6: Timeout handling
	timeoutExample()

	// Example 7: Non-blocking channel operations
	nonBlockingChannel()

	// Example 8: Channel as signal
	channelAsSignal()

	// Example 9: Common mistakes examples
	channelMistakes()
}

// Example 1: Basic channel communication
func basicChannel() {
	fmt.Println("\n--- Example 1: Basic channel communication ---")

	// Create unbuffered channel
	ch := make(chan string)

	// Start sender goroutine
	go func() {
		fmt.Println("Sender: preparing to send message")
		ch <- "Hello from sender"
		fmt.Println("Sender: message sent")
	}()

	// Main goroutine receives message
	fmt.Println("Receiver: waiting for message...")
	message := <-ch
	fmt.Printf("Receiver: received message '%s'\n", message)
}

// Example 2: Buffered channel
func bufferedChannel() {
	fmt.Println("\n--- Example 2: Buffered channel ---")

	// Create channel with buffer size 2
	ch := make(chan int, 2)

	// Can send multiple values without blocking
	ch <- 1
	ch <- 2
	fmt.Println("Sent 2 values to buffered channel")

	// Receive values
	fmt.Printf("Received value: %d\n", <-ch)
	fmt.Printf("Received value: %d\n", <-ch)

	// Demonstrate blocking behavior of buffered channel
	go func() {
		fmt.Println("Sender: trying to send to full buffered channel")
		ch <- 3 // This will block because buffer is full
		fmt.Println("Sender: sent successfully")
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Receiver: received %d from buffered channel\n", <-ch)
	time.Sleep(100 * time.Millisecond)
}

// Example 3: Directional channel
func directionalChannel() {
	fmt.Println("\n--- Example 3: Directional channel ---")

	// Create bidirectional channel
	ch := make(chan int)

	// Start sender (write-only)
	go sender(ch)

	// Start receiver (read-only)
	go receiver(ch)

	time.Sleep(200 * time.Millisecond)
}

// sender write-only channel
func sender(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Sender: sending %d\n", i)
		ch <- i
		time.Sleep(50 * time.Millisecond)
	}
	close(ch)
}

// receiver read-only channel
func receiver(ch <-chan int) {
	for value := range ch {
		fmt.Printf("Receiver: received %d\n", value)
	}
	fmt.Println("Receiver: channel closed")
}

// Example 4: Channel close
func channelClose() {
	fmt.Println("\n--- Example 4: Channel close ---")

	ch := make(chan int)

	// Start sender
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
			time.Sleep(50 * time.Millisecond)
		}
		close(ch) // Close channel
		fmt.Println("Sender: channel closed")
	}()

	// Receiver uses range loop
	fmt.Println("Receiver: starting to receive...")
	for value := range ch {
		fmt.Printf("Receiver: received %d\n", value)
	}
	fmt.Println("Receiver: receiving completed")

	// Demonstrate behavior after channel is closed
	fmt.Println("\nDemonstrating behavior after channel is closed:")
	ch2 := make(chan int)
	close(ch2)

	// Reading from closed channel returns zero value immediately
	value, ok := <-ch2
	fmt.Printf("Reading from closed channel: value=%d, ok=%t\n", value, ok)

	// Sending to closed channel will panic
	// ch2 <- 1 // This will panic
}

// Example 5: Select statement
func selectExample() {
	fmt.Println("\n--- Example 5: Select statement ---")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start two senders
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "message from channel1"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "message from channel2"
	}()

	// Use select to listen to multiple channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from ch2: %s\n", msg2)
		case <-time.After(200 * time.Millisecond):
			fmt.Println("Timeout")
		}
	}
}

// Example 6: Timeout handling
func timeoutExample() {
	fmt.Println("\n--- Example 6: Timeout handling ---")

	ch := make(chan string)

	// Simulate slow operation
	go func() {
		time.Sleep(300 * time.Millisecond) // Longer than timeout
		ch <- "slow operation completed"
	}()

	// Set timeout
	select {
	case result := <-ch:
		fmt.Printf("Operation successful: %s\n", result)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Operation timeout")
	}

	// Fast operation example
	ch2 := make(chan string)
	go func() {
		time.Sleep(50 * time.Millisecond) // Shorter than timeout
		ch2 <- "fast operation completed"
	}()

	select {
	case result := <-ch2:
		fmt.Printf("Operation successful: %s\n", result)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Operation timeout")
	}
}

// Example 7: Non-blocking channel operations
func nonBlockingChannel() {
	fmt.Println("\n--- Example 7: Non-blocking channel operations ---")

	ch := make(chan int, 1)

	// Non-blocking send
	select {
	case ch <- 1:
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

	// Try to receive from empty channel again
	select {
	case value := <-ch:
		fmt.Printf("Receive successful: %d\n", value)
	default:
		fmt.Println("Receive failed, channel is empty")
	}
}

// Example 8: Channel as signal
func channelAsSignal() {
	fmt.Println("\n--- Example 8: Channel as signal ---")

	// Use channel as stop signal
	stop := make(chan struct{})

	// Start worker goroutine
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Received stop signal, exiting")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// Main goroutine waits for a while then sends stop signal
	time.Sleep(300 * time.Millisecond)
	close(stop) // Send stop signal

	time.Sleep(100 * time.Millisecond)

	// Use channel as completion signal
	done := make(chan struct{})

	go func() {
		fmt.Println("Executing task...")
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Task completed")
		close(done) // Send completion signal
	}()

	<-done // Wait for completion signal
	fmt.Println("Main goroutine received completion signal")
}

// Example 9: Common mistakes examples
func channelMistakes() {
	fmt.Println("\n--- Example 9: Common mistakes examples ---")

	fmt.Println("Mistake 1: Sending data to nil channel")
	var ch chan int // nil channel
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Caught panic: %v\n", r)
			}
		}()
		ch <- 1 // This will panic
	}()
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\nMistake 2: Receiving data from nil channel")
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Caught panic: %v\n", r)
			}
		}()
		<-ch // This will block, but won't panic
	}()
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\nMistake 3: Closing channel multiple times")
	ch2 := make(chan int)
	close(ch2)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Caught panic: %v\n", r)
			}
		}()
		close(ch2) // Closing multiple times will panic
	}()
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\nMistake 4: Forgetting to close channel causing goroutine leak")
	ch3 := make(chan int)
	go func() {
		for value := range ch3 {
			fmt.Printf("Processing value: %d\n", value)
		}
		fmt.Println("Goroutine exited")
	}()

	ch3 <- 1
	ch3 <- 2
	// Forgot to close(ch3), goroutine will wait forever

	fmt.Println("\nCorrect approach: Use sync.WaitGroup to ensure goroutine completion")
	var wg sync.WaitGroup
	ch4 := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range ch4 {
			fmt.Printf("Processing value: %d\n", value)
		}
		fmt.Println("Goroutine exited normally")
	}()

	ch4 <- 1
	ch4 <- 2
	close(ch4) // Correctly close

	wg.Wait() // Wait for goroutine completion
	fmt.Println("All goroutines completed")
}
