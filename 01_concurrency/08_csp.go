package main

import (
	"fmt"
	"sync"
	"time"
)

// Define request and response structures
type Request struct {
	ID   int
	Data string
}

type Response struct {
	RequestID int
	Result    string
	Error     error
}

// CSPExamples runs all CSP pattern examples
func CSPExamples() {
	fmt.Println("=== CSP (Communicating Sequential Processes) Pattern Examples ===")

	// Example 1: Basic CSP communication
	basicCSPExample()

	// Example 2: Multi-process communication
	multiProcessCSPExample()

	// Example 3: Pipeline processing
	pipelineCSPExample()

	// Example 4: Fan-in Fan-out
	fanInFanOutCSPExample()

	// Example 5: Worker pool CSP
	workerPoolCSPExample()

	// Example 6: Publish Subscribe
	pubSubCSPExample()

	// Example 7: Request Response
	requestResponseCSPExample()

	// Example 8: Comprehensive example
	comprehensiveCSPExample()
}

// Example 1: Basic CSP communication
func basicCSPExample() {
	fmt.Println("\n--- Example 1: Basic CSP Communication ---")

	// Create communication channel
	ch := make(chan string)

	// Process 1: Sender
	go func() {
		defer close(ch)
		messages := []string{"Message1", "Message2", "Message3", "Message4"}
		for _, msg := range messages {
			fmt.Printf("Sender: sending %s\n", msg)
			ch <- msg
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("Sender: finished sending")
	}()

	// Process 2: Receiver
	go func() {
		for msg := range ch {
			fmt.Printf("Receiver: received %s\n", msg)
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Println("Receiver: finished receiving")
	}()

	time.Sleep(1 * time.Second)
}

// Example 2: Multi-process communication
func multiProcessCSPExample() {
	fmt.Println("\n--- Example 2: Multi-process Communication ---")

	// Create multiple channels for inter-process communication
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Process 1: Data generator
	go func() {
		defer close(ch1)
		for i := 1; i <= 5; i++ {
			fmt.Printf("Generator: generating data %d\n", i)
			ch1 <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Process 2: Data processor
	go func() {
		defer close(ch2)
		for data := range ch1 {
			processed := data * 2
			fmt.Printf("Processor: processing %d -> %d\n", data, processed)
			ch2 <- processed
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Process 3: Data validator
	go func() {
		defer close(ch3)
		for data := range ch2 {
			if data%2 == 0 {
				fmt.Printf("Validator: validation passed %d\n", data)
				ch3 <- data
			} else {
				fmt.Printf("Validator: validation failed %d\n", data)
			}
		}
	}()

	// Main process: Result collector
	go func() {
		for result := range ch3 {
			fmt.Printf("Collector: collecting result %d\n", result)
		}
		fmt.Println("Collector: finished collecting")
	}()

	time.Sleep(1 * time.Second)
}

// Example 3: Pipeline processing
func pipelineCSPExample() {
	fmt.Println("\n--- Example 3: Pipeline Processing ---")

	// Create pipeline stages
	stage1 := make(chan int, 10)
	stage2 := make(chan int, 10)
	stage3 := make(chan int, 10)

	// Stage 1: Data input
	go func() {
		defer close(stage1)
		for i := 1; i <= 10; i++ {
			fmt.Printf("Stage1: input %d\n", i)
			stage1 <- i
		}
	}()

	// Stage 2: Data filtering
	go func() {
		defer close(stage2)
		for data := range stage1 {
			if data%2 == 0 {
				fmt.Printf("Stage2: filter passed %d\n", data)
				stage2 <- data
			} else {
				fmt.Printf("Stage2: filter dropped %d\n", data)
			}
		}
	}()

	// Stage 3: Data transformation
	go func() {
		defer close(stage3)
		for data := range stage2 {
			transformed := data * data
			fmt.Printf("Stage3: transform %d -> %d\n", data, transformed)
			stage3 <- transformed
		}
	}()

	// Result output
	go func() {
		for result := range stage3 {
			fmt.Printf("Output: %d\n", result)
		}
		fmt.Println("Pipeline processing completed")
	}()

	time.Sleep(1 * time.Second)
}

// Example 4: Fan-in Fan-out
func fanInFanOutCSPExample() {
	fmt.Println("\n--- Example 4: Fan-in Fan-out ---")

	// Create input and output channels
	input := make(chan int, 10)
	output1 := make(chan int, 5)
	output2 := make(chan int, 5)
	output3 := make(chan int, 5)
	merged := make(chan int, 15)

	// Data generator
	go func() {
		defer close(input)
		for i := 1; i <= 15; i++ {
			fmt.Printf("Generator: generating %d\n", i)
			input <- i
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Fan-out: Distribute input to multiple workers
	go func() {
		defer close(output1)
		defer close(output2)
		defer close(output3)

		counter := 0
		for data := range input {
			switch counter % 3 {
			case 0:
				fmt.Printf("Dispatcher: sending to Worker1: %d\n", data)
				output1 <- data
			case 1:
				fmt.Printf("Dispatcher: sending to Worker2: %d\n", data)
				output2 <- data
			case 2:
				fmt.Printf("Dispatcher: sending to Worker3: %d\n", data)
				output3 <- data
			}
			counter++
		}
	}()

	// Worker 1
	go func() {
		for data := range output1 {
			processed := data * 10
			fmt.Printf("Worker1: processing %d -> %d\n", data, processed)
			merged <- processed
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Worker 2
	go func() {
		for data := range output2 {
			processed := data * 100
			fmt.Printf("Worker2: processing %d -> %d\n", data, processed)
			merged <- processed
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Worker 3
	go func() {
		for data := range output3 {
			processed := data * 1000
			fmt.Printf("Worker3: processing %d -> %d\n", data, processed)
			merged <- processed
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Fan-in: Merge results
	go func() {
		defer close(merged)
		var wg sync.WaitGroup

		// Wait for all workers to complete
		wg.Add(3)
		go func() {
			defer wg.Done()
			for range output1 {
			}
		}()
		go func() {
			defer wg.Done()
			for range output2 {
			}
		}()
		go func() {
			defer wg.Done()
			for range output3 {
			}
		}()

		wg.Wait()
	}()

	// Result collector
	go func() {
		for result := range merged {
			fmt.Printf("Collector: received result %d\n", result)
		}
		fmt.Println("Fan-in Fan-out processing completed")
	}()

	time.Sleep(2 * time.Second)
}

// Example 5: Worker pool CSP
func workerPoolCSPExample() {
	fmt.Println("\n--- Example 5: Worker Pool CSP ---")

	const numWorkers = 3
	const numJobs = 10

	// Create job and result channels
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	// Start worker pool
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Worker %d: starting to process job %d\n", workerID, job)
				time.Sleep(200 * time.Millisecond) // Simulate work
				result := fmt.Sprintf("Worker %d completed job %d", workerID, job)
				results <- result
				fmt.Printf("Worker %d: finished processing job %d\n", workerID, job)
			}
		}(i)
	}

	// Send jobs
	go func() {
		defer close(jobs)
		for i := 1; i <= numJobs; i++ {
			fmt.Printf("Scheduler: sending job %d\n", i)
			jobs <- i
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Collect results
	go func() {
		defer close(results)
		wg.Wait()
	}()

	// Process results
	go func() {
		for result := range results {
			fmt.Printf("Result processor: %s\n", result)
		}
		fmt.Println("Worker pool processing completed")
	}()

	time.Sleep(3 * time.Second)
}

// Example 6: Publish Subscribe
func pubSubCSPExample() {
	fmt.Println("\n--- Example 6: Publish Subscribe ---")

	// Create topic channels
	topic1 := make(chan string, 10)
	topic2 := make(chan string, 10)

	// Publisher
	go func() {
		messages := []struct {
			topic   chan string
			message string
		}{
			{topic1, "Topic1: Message1"},
			{topic2, "Topic2: Message1"},
			{topic1, "Topic1: Message2"},
			{topic2, "Topic2: Message2"},
			{topic1, "Topic1: Message3"},
		}

		for _, msg := range messages {
			fmt.Printf("Publisher: publishing to %v: %s\n", msg.topic, msg.message)
			msg.topic <- msg.message
			time.Sleep(100 * time.Millisecond)
		}

		close(topic1)
		close(topic2)
	}()

	// Subscriber 1: Subscribe to Topic1
	go func() {
		for msg := range topic1 {
			fmt.Printf("Subscriber1: received %s\n", msg)
		}
		fmt.Println("Subscriber1: finished subscribing")
	}()

	// Subscriber 2: Subscribe to Topic2
	go func() {
		for msg := range topic2 {
			fmt.Printf("Subscriber2: received %s\n", msg)
		}
		fmt.Println("Subscriber2: finished subscribing")
	}()

	// Subscriber 3: Subscribe to both Topics
	go func() {
		done := make(chan bool, 2)

		go func() {
			for msg := range topic1 {
				fmt.Printf("Subscriber3(Topic1): received %s\n", msg)
			}
			done <- true
		}()

		go func() {
			for msg := range topic2 {
				fmt.Printf("Subscriber3(Topic2): received %s\n", msg)
			}
			done <- true
		}()

		<-done
		<-done
		fmt.Println("Subscriber3: finished subscribing")
	}()

	time.Sleep(1 * time.Second)
}

// Example 7: Request Response
func requestResponseCSPExample() {
	fmt.Println("\n--- Example 7: Request Response ---")

	// Create request and response channels
	requests := make(chan Request, 10)
	responses := make(chan Response, 10)

	// Service processor
	go func() {
		defer close(responses)
		for req := range requests {
			fmt.Printf("Service processor: processing request %d: %s\n", req.ID, req.Data)

			// Simulate processing time
			time.Sleep(100 * time.Millisecond)

			// Generate response
			response := Response{
				RequestID: req.ID,
				Result:    fmt.Sprintf("Processing result: %s processed", req.Data),
				Error:     nil,
			}

			responses <- response
			fmt.Printf("Service processor: responding to request %d\n", req.ID)
		}
	}()

	// Client 1
	go func() {
		for i := 1; i <= 3; i++ {
			req := Request{ID: i, Data: fmt.Sprintf("Client1 data%d", i)}
			fmt.Printf("Client1: sending request %d\n", i)
			requests <- req
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Client 2
	go func() {
		for i := 4; i <= 6; i++ {
			req := Request{ID: i, Data: fmt.Sprintf("Client2 data%d", i)}
			fmt.Printf("Client2: sending request %d\n", i)
			requests <- req
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Response collector
	go func() {
		for response := range responses {
			fmt.Printf("Response collector: received response %d: %s\n", response.RequestID, response.Result)
		}
		fmt.Println("Request response processing completed")
	}()

	// Close request channel
	go func() {
		time.Sleep(1 * time.Second)
		close(requests)
	}()

	time.Sleep(2 * time.Second)
}

// Example 8: Comprehensive example
func comprehensiveCSPExample() {
	fmt.Println("\n--- Example 8: Comprehensive Example ---")

	// Simulate a simple order processing system
	type Order struct {
		ID       int
		Customer string
		Items    []string
		Total    float64
	}

	type OrderStatus struct {
		OrderID int
		Status  string
		Message string
	}

	// Create processing pipeline
	orders := make(chan Order, 10)
	validated := make(chan Order, 10)
	processed := make(chan Order, 10)
	status := make(chan OrderStatus, 10)

	// Order generator
	go func() {
		defer close(orders)
		orderData := []Order{
			{ID: 1, Customer: "Alice", Items: []string{"Apple", "Banana"}, Total: 15.50},
			{ID: 2, Customer: "Bob", Items: []string{"Orange"}, Total: 8.00},
			{ID: 3, Customer: "Charlie", Items: []string{"Apple", "Grape", "Pear"}, Total: 25.00},
			{ID: 4, Customer: "David", Items: []string{}, Total: 0.00},
		}

		for _, order := range orderData {
			fmt.Printf("Order generator: generating order %d\n", order.ID)
			orders <- order
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Order validator
	go func() {
		defer close(validated)
		for order := range orders {
			fmt.Printf("Validator: validating order %d\n", order.ID)

			if len(order.Items) == 0 {
				status <- OrderStatus{OrderID: order.ID, Status: "REJECTED", Message: "Order is empty"}
				continue
			}

			if order.Total <= 0 {
				status <- OrderStatus{OrderID: order.ID, Status: "REJECTED", Message: "Invalid amount"}
				continue
			}

			validated <- order
			status <- OrderStatus{OrderID: order.ID, Status: "VALIDATED", Message: "Validation passed"}
		}
	}()

	// Order processor
	go func() {
		defer close(processed)
		for order := range validated {
			fmt.Printf("Processor: processing order %d\n", order.ID)
			time.Sleep(200 * time.Millisecond) // Simulate processing time
			processed <- order
			status <- OrderStatus{OrderID: order.ID, Status: "PROCESSED", Message: "Processing completed"}
		}
	}()

	// Order completer
	go func() {
		for order := range processed {
			fmt.Printf("Completer: completing order %d (Customer: %s, Total: %.2f)\n",
				order.ID, order.Customer, order.Total)
			status <- OrderStatus{OrderID: order.ID, Status: "COMPLETED", Message: "Order completed"}
		}
	}()

	// Status monitor
	go func() {
		for orderStatus := range status {
			fmt.Printf("Monitor: Order %d - %s: %s\n",
				orderStatus.OrderID, orderStatus.Status, orderStatus.Message)
		}
		fmt.Println("Order processing system completed")
	}()

	time.Sleep(2 * time.Second)
}
