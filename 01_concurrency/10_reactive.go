package main

import (
	"fmt"
	"sync"
	"time"
)

// Reactive programming examples

// Observable observable object
type Observable struct {
	observers []chan interface{}
	mu        sync.RWMutex
	closed    bool
}

// NewObservable creates a new Observable
func NewObservable() *Observable {
	return &Observable{
		observers: make([]chan interface{}, 0),
	}
}

// Subscribe subscribes to Observable
func (o *Observable) Subscribe() chan interface{} {
	o.mu.Lock()
	defer o.mu.Unlock()

	if o.closed {
		return nil
	}

	ch := make(chan interface{}, 10)
	o.observers = append(o.observers, ch)
	return ch
}

// Emit sends data to all observers
func (o *Observable) Emit(data interface{}) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	if o.closed {
		return
	}

	for _, observer := range o.observers {
		select {
		case observer <- data:
		default:
			// Skip if channel is full
		}
	}
}

// Close closes Observable
func (o *Observable) Close() {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.closed = true
	for _, observer := range o.observers {
		close(observer)
	}
	o.observers = nil
}

// Subject topic, both Observable and Observer
type Subject struct {
	*Observable
	observers []chan interface{}
	mu        sync.RWMutex
}

// NewSubject creates a new Subject
func NewSubject() *Subject {
	return &Subject{
		Observable: NewObservable(),
		observers:  make([]chan interface{}, 0),
	}
}

// Next sends next value
func (s *Subject) Next(value interface{}) {
	s.Emit(value)
}

// Error sends error
func (s *Subject) Error(err error) {
	s.Emit(fmt.Sprintf("ERROR: %v", err))
}

// Complete completes
func (s *Subject) Complete() {
	s.Emit("COMPLETE")
	s.Close()
}

// ReactiveExamples runs all Reactive programming examples
func ReactiveExamples() {
	fmt.Println("=== Reactive Programming Examples ===")

	// Example 1: Basic Observable
	basicObservableExample()

	// Example 2: Subject usage
	subjectExample()

	// Example 3: Data stream processing
	dataStreamExample()

	// Example 4: Event handling
	eventHandlingExample()

	// Example 5: Asynchronous data stream
	asyncDataStreamExample()

	// Example 6: Error handling
	reactiveErrorHandlingExample()

	// Example 7: Compose operations
	composeReactiveExample()

	// Example 8: Comprehensive example
	comprehensiveReactiveExample()
}

// Example 1: Basic Observable
func basicObservableExample() {
	fmt.Println("\n--- Example 1: Basic Observable ---")

	observable := NewObservable()

	// Subscriber 1
	ch1 := observable.Subscribe()
	go func() {
		for data := range ch1 {
			fmt.Printf("Subscriber1: received %v\n", data)
		}
		fmt.Println("Subscriber1: completed")
	}()

	// Subscriber 2
	ch2 := observable.Subscribe()
	go func() {
		for data := range ch2 {
			fmt.Printf("Subscriber2: received %v\n", data)
		}
		fmt.Println("Subscriber2: completed")
	}()

	// Send data
	fmt.Println("Sending data...")
	observable.Emit("Message1")
	time.Sleep(100 * time.Millisecond)
	observable.Emit("Message2")
	time.Sleep(100 * time.Millisecond)
	observable.Emit("Message3")

	// Close Observable
	time.Sleep(200 * time.Millisecond)
	observable.Close()
}

// Example 2: Subject usage
func subjectExample() {
	fmt.Println("\n--- Example 2: Subject Usage ---")

	subject := NewSubject()

	// Subscriber
	ch := subject.Subscribe()
	go func() {
		for data := range ch {
			fmt.Printf("Subject subscriber: received %v\n", data)
		}
		fmt.Println("Subject subscriber: completed")
	}()

	// Send data
	fmt.Println("Subject sending data...")
	subject.Next("Data1")
	time.Sleep(50 * time.Millisecond)
	subject.Next("Data2")
	time.Sleep(50 * time.Millisecond)
	subject.Next("Data3")

	// Send error
	subject.Error(fmt.Errorf("Simulated error"))

	// Complete
	subject.Complete()
}

// Example 3: Data stream processing
func dataStreamExample() {
	fmt.Println("\n--- Example 3: Data Stream Processing ---")

	// Create data source
	dataSource := NewObservable()

	// Data processor 1: Filter even numbers
	evenFilter := NewObservable()
	go func() {
		ch := dataSource.Subscribe()
		for data := range ch {
			if num, ok := data.(int); ok && num%2 == 0 {
				evenFilter.Emit(num)
			}
		}
		evenFilter.Close()
	}()

	// Data processor 2: Square operation
	squareProcessor := NewObservable()
	go func() {
		ch := evenFilter.Subscribe()
		for data := range ch {
			if num, ok := data.(int); ok {
				squareProcessor.Emit(num * num)
			}
		}
		squareProcessor.Close()
	}()

	// Result collector
	go func() {
		ch := squareProcessor.Subscribe()
		for data := range ch {
			fmt.Printf("Result: %v\n", data)
		}
		fmt.Println("Data stream processing completed")
	}()

	// Send data
	fmt.Println("Sending data stream...")
	for i := 1; i <= 10; i++ {
		dataSource.Emit(i)
		time.Sleep(50 * time.Millisecond)
	}

	// Close data source
	dataSource.Close()
	time.Sleep(200 * time.Millisecond)
}

// Example 4: Event handling
func eventHandlingExample() {
	fmt.Println("\n--- Example 4: Event Handling ---")

	// Create event bus
	eventBus := NewSubject()

	// Event handler 1: User events
	userEventHandler := eventBus.Subscribe()
	go func() {
		for event := range userEventHandler {
			if eventStr, ok := event.(string); ok {
				if eventStr == "user_login" {
					fmt.Println("User event handler: User login")
				} else if eventStr == "user_logout" {
					fmt.Println("User event handler: User logout")
				}
			}
		}
	}()

	// Event handler 2: System events
	systemEventHandler := eventBus.Subscribe()
	go func() {
		for event := range systemEventHandler {
			if eventStr, ok := event.(string); ok {
				if eventStr == "system_start" {
					fmt.Println("System event handler: System startup")
				} else if eventStr == "system_shutdown" {
					fmt.Println("System event handler: System shutdown")
				}
			}
		}
	}()

	// Event handler 3: Log recording
	logEventHandler := eventBus.Subscribe()
	go func() {
		for event := range logEventHandler {
			fmt.Printf("Log recorder: Recording event %v\n", event)
		}
	}()

	// Send events
	fmt.Println("Sending events...")
	eventBus.Next("user_login")
	time.Sleep(100 * time.Millisecond)
	eventBus.Next("system_start")
	time.Sleep(100 * time.Millisecond)
	eventBus.Next("user_logout")
	time.Sleep(100 * time.Millisecond)
	eventBus.Next("system_shutdown")

	// Complete
	eventBus.Complete()
}

// Example 5: Asynchronous data stream
func asyncDataStreamExample() {
	fmt.Println("\n--- Example 5: Asynchronous Data Stream ---")

	// Create asynchronous data source
	asyncSource := NewObservable()

	// Asynchronous data generator
	go func() {
		defer asyncSource.Close()
		for i := 1; i <= 5; i++ {
			fmt.Printf("Async generator: generating data %d\n", i)
			asyncSource.Emit(i)
			time.Sleep(200 * time.Millisecond) // Simulate async operation
		}
	}()

	// Async processor 1
	processor1 := NewObservable()
	go func() {
		ch := asyncSource.Subscribe()
		for data := range ch {
			if num, ok := data.(int); ok {
				// Simulate async processing
				time.Sleep(100 * time.Millisecond)
				processed := num * 2
				fmt.Printf("Async processor1: %d -> %d\n", num, processed)
				processor1.Emit(processed)
			}
		}
		processor1.Close()
	}()

	// Async processor 2
	processor2 := NewObservable()
	go func() {
		ch := asyncSource.Subscribe()
		for data := range ch {
			if num, ok := data.(int); ok {
				// Simulate async processing
				time.Sleep(150 * time.Millisecond)
				processed := num * num
				fmt.Printf("Async processor2: %d -> %d\n", num, processed)
				processor2.Emit(processed)
			}
		}
		processor2.Close()
	}()

	// Result merger
	mergedResults := NewObservable()
	go func() {
		var wg sync.WaitGroup
		wg.Add(2)

		// Merge processor1 results
		go func() {
			defer wg.Done()
			ch := processor1.Subscribe()
			for data := range ch {
				mergedResults.Emit(fmt.Sprintf("P1: %v", data))
			}
		}()

		// Merge processor2 results
		go func() {
			defer wg.Done()
			ch := processor2.Subscribe()
			for data := range ch {
				mergedResults.Emit(fmt.Sprintf("P2: %v", data))
			}
		}()

		wg.Wait()
		mergedResults.Close()
	}()

	// Result collector
	go func() {
		ch := mergedResults.Subscribe()
		for data := range ch {
			fmt.Printf("Result collector: %v\n", data)
		}
		fmt.Println("Asynchronous data stream processing completed")
	}()

	time.Sleep(2 * time.Second)
}

// Example 6: Error handling
func reactiveErrorHandlingExample() {
	fmt.Println("\n--- Example 6: Error Handling ---")

	// Create Observable that may error
	errorProneSource := NewObservable()

	// Error handler
	errorHandler := NewObservable()
	go func() {
		ch := errorProneSource.Subscribe()
		for data := range ch {
			if errStr, ok := data.(string); ok && len(errStr) > 5 && errStr[:5] == "ERROR" {
				fmt.Printf("Error handler: Processing error %s\n", errStr)
				errorHandler.Emit(fmt.Sprintf("Processed: %s", errStr))
			} else {
				fmt.Printf("Error handler: Normal data %v\n", data)
			}
		}
		errorHandler.Close()
	}()

	// Recovery handler
	recoveryHandler := NewObservable()
	go func() {
		ch := errorHandler.Subscribe()
		for data := range ch {
			fmt.Printf("Recovery handler: %v\n", data)
			recoveryHandler.Emit("Recovery successful")
		}
		recoveryHandler.Close()
	}()

	// Send normal data and error data
	go func() {
		defer errorProneSource.Close()

		errorProneSource.Emit("Normal data1")
		time.Sleep(100 * time.Millisecond)
		errorProneSource.Emit("ERROR: Simulated error1")
		time.Sleep(100 * time.Millisecond)
		errorProneSource.Emit("Normal data2")
		time.Sleep(100 * time.Millisecond)
		errorProneSource.Emit("ERROR: Simulated error2")
	}()

	// Final result collector
	go func() {
		ch := recoveryHandler.Subscribe()
		for data := range ch {
			fmt.Printf("Final result: %v\n", data)
		}
		fmt.Println("Error handling completed")
	}()

	time.Sleep(1 * time.Second)
}

// Example 7: Compose operations
func composeReactiveExample() {
	fmt.Println("\n--- Example 7: Compose Operations ---")

	// Create multiple data sources
	source1 := NewObservable()
	source2 := NewObservable()
	source3 := NewObservable()

	// Compose processor
	composed := NewObservable()
	go func() {
		var wg sync.WaitGroup
		wg.Add(3)

		// Process source1
		go func() {
			defer wg.Done()
			ch := source1.Subscribe()
			for data := range ch {
				composed.Emit(fmt.Sprintf("S1: %v", data))
			}
		}()

		// Process source2
		go func() {
			defer wg.Done()
			ch := source2.Subscribe()
			for data := range ch {
				composed.Emit(fmt.Sprintf("S2: %v", data))
			}
		}()

		// Process source3
		go func() {
			defer wg.Done()
			ch := source3.Subscribe()
			for data := range ch {
				composed.Emit(fmt.Sprintf("S3: %v", data))
			}
		}()

		wg.Wait()
		composed.Close()
	}()

	// Result collector
	go func() {
		ch := composed.Subscribe()
		for data := range ch {
			fmt.Printf("Compose result: %v\n", data)
		}
		fmt.Println("Compose operations completed")
	}()

	// Send data to different sources
	go func() {
		for i := 1; i <= 3; i++ {
			source1.Emit(fmt.Sprintf("Data%d", i))
			time.Sleep(50 * time.Millisecond)
		}
		source1.Close()
	}()

	go func() {
		for i := 4; i <= 6; i++ {
			source2.Emit(fmt.Sprintf("Data%d", i))
			time.Sleep(50 * time.Millisecond)
		}
		source2.Close()
	}()

	go func() {
		for i := 7; i <= 9; i++ {
			source3.Emit(fmt.Sprintf("Data%d", i))
			time.Sleep(50 * time.Millisecond)
		}
		source3.Close()
	}()

	time.Sleep(1 * time.Second)
}

// Example 8: Comprehensive example
func comprehensiveReactiveExample() {
	fmt.Println("\n--- Example 8: Comprehensive Example ---")

	// Simulate a reactive user interface system
	type UIEvent struct {
		Type string
		Data interface{}
	}

	// Create UI event bus
	uiEventBus := NewSubject()

	// User input processor
	inputProcessor := NewObservable()
	go func() {
		ch := uiEventBus.Subscribe()
		for event := range ch {
			if uiEvent, ok := event.(UIEvent); ok {
				if uiEvent.Type == "input" {
					fmt.Printf("Input processor: Processing input %v\n", uiEvent.Data)
					inputProcessor.Emit(uiEvent.Data)
				}
			}
		}
		inputProcessor.Close()
	}()

	// Data validator
	validator := NewObservable()
	go func() {
		ch := inputProcessor.Subscribe()
		for data := range ch {
			if str, ok := data.(string); ok {
				if len(str) > 0 {
					fmt.Printf("Validator: Validation passed %s\n", str)
					validator.Emit(data)
				} else {
					fmt.Printf("Validator: Validation failed %s\n", str)
				}
			}
		}
		validator.Close()
	}()

	// Data processor
	dataProcessor := NewObservable()
	go func() {
		ch := validator.Subscribe()
		for data := range ch {
			if str, ok := data.(string); ok {
				processed := fmt.Sprintf("Processed: %s", str)
				fmt.Printf("Data processor: %s -> %s\n", str, processed)
				dataProcessor.Emit(processed)
			}
		}
		dataProcessor.Close()
	}()

	// UI updater
	uiUpdater := NewObservable()
	go func() {
		ch := dataProcessor.Subscribe()
		for data := range ch {
			fmt.Printf("UI updater: Updating interface %v\n", data)
			uiUpdater.Emit(fmt.Sprintf("Interface updated: %v", data))
		}
		uiUpdater.Close()
	}()

	// Log recorder
	go func() {
		ch := uiUpdater.Subscribe()
		for data := range ch {
			fmt.Printf("Log recorder: %v\n", data)
		}
		fmt.Println("UI system processing completed")
	}()

	// Simulate user input
	fmt.Println("Simulating user input...")
	inputs := []string{"Hello", "", "World", "Reactive", "Programming"}
	for _, input := range inputs {
		uiEventBus.Next(UIEvent{Type: "input", Data: input})
		time.Sleep(200 * time.Millisecond)
	}

	// Complete
	uiEventBus.Complete()
	time.Sleep(500 * time.Millisecond)
}
