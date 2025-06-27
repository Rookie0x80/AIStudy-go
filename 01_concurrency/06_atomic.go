package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// AtomicExamples runs all atomic operation examples
func AtomicExamples() {
	fmt.Println("=== Atomic Operations Examples ===")

	// Example 1: Atomic integer operations
	atomicIntExample()

	// Example 2: Atomic pointer operations
	atomicPointerExample()

	// Example 3: Atomic value operations
	atomicValueExample()

	// Example 4: Atomic compare and swap
	atomicCASExample()

	// Example 5: Atomic load and store
	atomicLoadStoreExample()

	// Example 6: Atomic addition operations
	atomicAddExample()

	// Example 7: Atomic bitwise operations
	atomicBitwiseExample()

	// Example 8: Comprehensive example
	comprehensiveAtomicExample()
}

// Example 1: Atomic integer operations
func atomicIntExample() {
	fmt.Println("\n--- Example 1: Atomic integer operations ---")

	var counter int64
	var wg sync.WaitGroup

	// Start multiple goroutines to increment counter simultaneously
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				// Atomic increment
				atomic.AddInt64(&counter, 1)
			}
			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d (Expected: 10000)\n", atomic.LoadInt64(&counter))

	// Demonstrate other atomic operations
	fmt.Println("\nDemonstrating other atomic operations:")

	// Atomic decrement
	atomic.AddInt64(&counter, -100)
	fmt.Printf("After decrementing 100: %d\n", atomic.LoadInt64(&counter))

	// Atomic swap
	oldValue := atomic.SwapInt64(&counter, 5000)
	fmt.Printf("Before swap: %d, After swap: %d\n", oldValue, atomic.LoadInt64(&counter))

	// Atomic compare and swap
	swapped := atomic.CompareAndSwapInt64(&counter, 5000, 6000)
	fmt.Printf("CAS operation: %t, Current value: %d\n", swapped, atomic.LoadInt64(&counter))
}

// Example 2: Atomic pointer operations
func atomicPointerExample() {
	fmt.Println("\n--- Example 2: Atomic pointer operations ---")

	type Data struct {
		ID   int
		Name string
	}

	var ptr atomic.Value // Store pointer
	var wg sync.WaitGroup

	// Initialize pointer
	initialData := &Data{ID: 0, Name: "Initial data"}
	ptr.Store(initialData)

	// Start multiple goroutines to update pointer
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Create new data
			newData := &Data{ID: id, Name: fmt.Sprintf("Data%d", id)}

			// Atomically store new pointer
			ptr.Store(newData)
			fmt.Printf("Goroutine %d updated pointer\n", id)

			time.Sleep(50 * time.Millisecond)
		}(i)
	}

	// Start reader goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if data, ok := ptr.Load().(*Data); ok {
				fmt.Printf("Read: ID=%d, Name=%s\n", data.ID, data.Name)
			}
			time.Sleep(30 * time.Millisecond)
		}
	}()

	wg.Wait()

	// Show final pointer value
	if finalData, ok := ptr.Load().(*Data); ok {
		fmt.Printf("Final pointer: ID=%d, Name=%s\n", finalData.ID, finalData.Name)
	}
}

// Example 3: Atomic value operations
func atomicValueExample() {
	fmt.Println("\n--- Example 3: Atomic value operations ---")

	var value atomic.Value
	var wg sync.WaitGroup

	// Store different types of values
	value.Store("String value")
	fmt.Printf("Stored string: %v\n", value.Load())

	value.Store(42)
	fmt.Printf("Stored integer: %v\n", value.Load())

	value.Store([]int{1, 2, 3})
	fmt.Printf("Stored slice: %v\n", value.Load())

	// Concurrent store and read
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Store new value
			newValue := fmt.Sprintf("Goroutine%d's value", id)
			value.Store(newValue)
			fmt.Printf("Goroutine %d stored: %s\n", id, newValue)

			// Read current value
			currentValue := value.Load()
			fmt.Printf("Goroutine %d read: %v\n", id, currentValue)
		}(i)
	}

	wg.Wait()
}

// Example 4: Atomic compare and swap
func atomicCASExample() {
	fmt.Println("\n--- Example 4: Atomic compare and swap ---")

	var value int64 = 100
	var wg sync.WaitGroup

	// Start multiple goroutines to attempt CAS operations
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Try to change 100 to 200
			swapped := atomic.CompareAndSwapInt64(&value, 100, 200)
			if swapped {
				fmt.Printf("Goroutine %d: CAS successful, value changed from 100 to 200\n", id)
			} else {
				fmt.Printf("Goroutine %d: CAS failed, current value: %d\n", id, atomic.LoadInt64(&value))
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final value: %d\n", atomic.LoadInt64(&value))

	// Demonstrate CAS for pointers
	fmt.Println("\nDemonstrating CAS for pointers:")
	type Node struct {
		Value int
		Next  *Node
	}

	var head atomic.Value
	initialNode := &Node{Value: 1, Next: nil}
	head.Store(initialNode)

	// Try to update head node
	newNode := &Node{Value: 2, Next: initialNode}
	swapped := head.CompareAndSwap(initialNode, newNode)
	fmt.Printf("Head node CAS: %t\n", swapped)
}

// Example 5: Atomic load and store
func atomicLoadStoreExample() {
	fmt.Println("\n--- Example 5: Atomic load and store ---")

	var flag int32
	var wg sync.WaitGroup

	// Start multiple goroutines to check flag
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Atomically load flag
			currentFlag := atomic.LoadInt32(&flag)
			fmt.Printf("Goroutine %d read flag: %d\n", id, currentFlag)

			// If flag is 0, set it to 1
			if currentFlag == 0 {
				atomic.StoreInt32(&flag, 1)
				fmt.Printf("Goroutine %d set flag to 1\n", id)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final flag value: %d\n", atomic.LoadInt32(&flag))

	// Demonstrate atomic swap
	fmt.Println("\nDemonstrating atomic swap:")
	var value int64 = 100
	oldValue := atomic.SwapInt64(&value, 200)
	fmt.Printf("Swap: old value=%d, new value=%d\n", oldValue, atomic.LoadInt64(&value))
}

// Example 6: Atomic addition operations
func atomicAddExample() {
	fmt.Println("\n--- Example 6: Atomic addition operations ---")

	var counter int64
	var wg sync.WaitGroup

	// Start multiple goroutines for addition operations
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Atomic addition
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, int64(id+1))
			}
			fmt.Printf("Goroutine %d completed addition operations\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", atomic.LoadInt64(&counter))

	// Demonstrate addition with negative numbers and zero
	fmt.Println("\nDemonstrating addition with negative numbers and zero:")
	atomic.AddInt64(&counter, -50)
	fmt.Printf("After subtracting 50: %d\n", atomic.LoadInt64(&counter))

	atomic.AddInt64(&counter, 0)
	fmt.Printf("After adding 0: %d\n", atomic.LoadInt64(&counter))
}

// Example 7: Atomic bitwise operations
func atomicBitwiseExample() {
	fmt.Println("\n--- Example 7: Atomic bitwise operations ---")

	var flags uint32
	var wg sync.WaitGroup

	// Set different flag bits
	flag1 := uint32(1 << 0) // Bit 0
	flag2 := uint32(1 << 1) // Bit 1
	flag3 := uint32(1 << 2) // Bit 2

	// Start multiple goroutines to set flags
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Use CAS operation to set flag 1
		for {
			old := atomic.LoadUint32(&flags)
			if atomic.CompareAndSwapUint32(&flags, old, old|flag1) {
				break
			}
		}
		fmt.Println("Set flag 1")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Use CAS operation to set flag 2
		for {
			old := atomic.LoadUint32(&flags)
			if atomic.CompareAndSwapUint32(&flags, old, old|flag2) {
				break
			}
		}
		fmt.Println("Set flag 2")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Use CAS operation to set flag 3
		for {
			old := atomic.LoadUint32(&flags)
			if atomic.CompareAndSwapUint32(&flags, old, old|flag3) {
				break
			}
		}
		fmt.Println("Set flag 3")
	}()

	wg.Wait()
	fmt.Printf("Final flag value: %b (binary)\n", atomic.LoadUint32(&flags))

	// Clear flag
	for {
		old := atomic.LoadUint32(&flags)
		if atomic.CompareAndSwapUint32(&flags, old, old&^flag2) {
			break
		}
	}
	fmt.Printf("After clearing flag 2: %b\n", atomic.LoadUint32(&flags))

	// Check if flags are set
	if atomic.LoadUint32(&flags)&flag1 != 0 {
		fmt.Println("Flag 1 is set")
	}
	if atomic.LoadUint32(&flags)&flag2 != 0 {
		fmt.Println("Flag 2 is set")
	} else {
		fmt.Println("Flag 2 is not set")
	}
}

// Example 8: Comprehensive example
func comprehensiveAtomicExample() {
	fmt.Println("\n--- Example 8: Comprehensive example ---")

	// Simulate an atomic counter system
	type AtomicCounter struct {
		value    int64
		requests int64
		errors   int64
	}

	counter := &AtomicCounter{}
	var wg sync.WaitGroup

	// Start multiple goroutines for counting operations
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < 100; j++ {
				// Increment request count
				atomic.AddInt64(&counter.requests, 1)

				// Simulate success or failure
				if j%10 == 0 {
					// Simulate error
					atomic.AddInt64(&counter.errors, 1)
				} else {
					// Successful operation, increment counter
					atomic.AddInt64(&counter.value, 1)
				}
			}
		}(i)
	}

	wg.Wait()

	// Show final statistics
	fmt.Printf("Final statistics:\n")
	fmt.Printf("  Total requests: %d\n", atomic.LoadInt64(&counter.requests))
	fmt.Printf("  Successful operations: %d\n", atomic.LoadInt64(&counter.value))
	fmt.Printf("  Error count: %d\n", atomic.LoadInt64(&counter.errors))

	// Demonstrate complex usage of atomic values
	fmt.Println("\nDemonstrating complex usage of atomic values:")

	type Config struct {
		MaxConnections int
		Timeout        time.Duration
		Enabled        bool
	}

	var config atomic.Value
	config.Store(Config{
		MaxConnections: 100,
		Timeout:        30 * time.Second,
		Enabled:        true,
	})

	// Concurrently read configuration
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			if cfg, ok := config.Load().(Config); ok {
				fmt.Printf("Goroutine %d read config: MaxConn=%d, Timeout=%v, Enabled=%t\n",
					id, cfg.MaxConnections, cfg.Timeout, cfg.Enabled)
			}
		}(i)
	}

	// Update configuration
	wg.Add(1)
	go func() {
		defer wg.Done()

		newConfig := Config{
			MaxConnections: 200,
			Timeout:        60 * time.Second,
			Enabled:        false,
		}
		config.Store(newConfig)
		fmt.Println("Configuration updated")
	}()

	wg.Wait()

	// Show final configuration
	if finalConfig, ok := config.Load().(Config); ok {
		fmt.Printf("Final configuration: MaxConn=%d, Timeout=%v, Enabled=%t\n",
			finalConfig.MaxConnections, finalConfig.Timeout, finalConfig.Enabled)
	}
}
