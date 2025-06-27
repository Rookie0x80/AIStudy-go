package main

import (
	"fmt"
	"sync"
	"time"
)

// SyncExamples runs all sync package examples
func SyncExamples() {
	fmt.Println("=== Sync Package Synchronization Primitives Examples ===")

	// Example 1: Mutex mutual exclusion lock
	mutexExample()

	// Example 2: RWMutex read-write lock
	rwMutexExample()

	// Example 3: WaitGroup wait group
	waitGroupExample()

	// Example 4: Once single execution
	onceExample()

	// Example 5: Cond condition variable
	condExample()

	// Example 6: Pool object pool
	poolExample()

	// Example 7: Map concurrent safe mapping
	mapExample()

	// Example 8: Comprehensive example
	comprehensiveSyncExample()
}

// Example 1: Mutex mutual exclusion lock
func mutexExample() {
	fmt.Println("\n--- Example 1: Mutex mutual exclusion lock ---")

	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Start multiple goroutines to access shared resource simultaneously
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock() // Acquire lock
				counter++
				mu.Unlock() // Release lock
			}
			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d (Expected: 5000)\n", counter)

	// Demonstrate nested lock usage
	fmt.Println("\nDemonstrating nested lock usage:")
	var nestedMu sync.Mutex
	nestedMu.Lock()
	fmt.Println("First lock acquired")
	nestedMu.Lock()                     // This will cause deadlock!
	fmt.Println("Second lock acquired") // Will never execute here
}

// Example 2: RWMutex read-write lock
func rwMutexExample() {
	fmt.Println("\n--- Example 2: RWMutex read-write lock ---")

	var data map[string]int
	var rwMu sync.RWMutex
	var wg sync.WaitGroup

	// Initialize data
	data = make(map[string]int)
	data["counter"] = 0

	// Start multiple reader goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				rwMu.RLock() // Read lock
				value := data["counter"]
				fmt.Printf("Reader %d: read value %d\n", id, value)
				rwMu.RUnlock() // Release read lock
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	// Start writer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			rwMu.Lock() // Write lock
			data["counter"]++
			fmt.Printf("Writer: wrote value %d\n", data["counter"])
			rwMu.Unlock() // Release write lock
			time.Sleep(200 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Printf("Final data: %v\n", data)
}

// Example 3: WaitGroup wait group
func waitGroupExample() {
	fmt.Println("\n--- Example 3: WaitGroup wait group ---")

	var wg sync.WaitGroup

	// Start multiple goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment wait counter
		go func(id int) {
			defer wg.Done() // Decrement counter when done
			fmt.Printf("Worker %d starting work\n", id)
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)
			fmt.Printf("Worker %d completed work\n", id)
		}(i)
	}

	fmt.Println("Waiting for all goroutines to complete...")
	wg.Wait() // Wait for all goroutines to complete
	fmt.Println("All goroutines completed")

	// Demonstrate incorrect usage
	fmt.Println("\nDemonstrating incorrect usage:")
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		fmt.Println("This goroutine will complete normally")
	}()

	// Wrong: calling Done() outside goroutine
	// wg2.Done() // This will cause panic

	wg2.Wait()
}

// Example 4: Once single execution
func onceExample() {
	fmt.Println("\n--- Example 4: Once single execution ---")

	var once sync.Once
	var wg sync.WaitGroup

	// Start multiple goroutines, but initialization function will only execute once
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			once.Do(func() {
				fmt.Printf("Initialization function called (by goroutine %d)\n", id)
				time.Sleep(100 * time.Millisecond) // Simulate initialization work
			})
			fmt.Printf("Goroutine %d continues execution\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed")

	// Demonstrate Once's non-repeatability
	fmt.Println("\nDemonstrating Once's non-repeatability:")
	var once2 sync.Once
	once2.Do(func() {
		fmt.Println("First call")
	})
	once2.Do(func() {
		fmt.Println("Second call") // Won't execute
	})
}

// Example 5: Cond condition variable
func condExample() {
	fmt.Println("\n--- Example 5: Cond condition variable ---")

	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	var ready bool
	var wg sync.WaitGroup

	// Waiter goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		for !ready {
			fmt.Println("Waiter: waiting for condition to be met...")
			cond.Wait() // Wait for condition to be met
		}
		fmt.Println("Waiter: condition met, continuing execution")
		mu.Unlock()
	}()

	// Notifier goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second) // Simulate work
		mu.Lock()
		ready = true
		fmt.Println("Notifier: setting condition to true")
		cond.Signal() // Notify one waiter
		mu.Unlock()
	}()

	wg.Wait()

	// Demonstrate Broadcast
	fmt.Println("\nDemonstrating Broadcast:")
	var mu2 sync.Mutex
	cond2 := sync.NewCond(&mu2)
	var ready2 bool

	// Start multiple waiters
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu2.Lock()
			for !ready2 {
				fmt.Printf("Waiter %d: waiting for condition to be met...\n", id)
				cond2.Wait()
			}
			fmt.Printf("Waiter %d: condition met, continuing execution\n", id)
			mu2.Unlock()
		}(i)
	}

	// Notifier
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		mu2.Lock()
		ready2 = true
		fmt.Println("Notifier: broadcasting condition met")
		cond2.Broadcast() // Notify all waiters
		mu2.Unlock()
	}()

	wg.Wait()
}

// Example 6: Pool object pool
func poolExample() {
	fmt.Println("\n--- Example 6: Pool object pool ---")

	var pool sync.Pool

	// Set New function
	pool.New = func() interface{} {
		fmt.Println("Creating new object")
		return &struct {
			ID   int
			Data string
		}{}
	}

	// Get object from pool
	obj1 := pool.Get().(*struct {
		ID   int
		Data string
	})
	obj1.ID = 1
	obj1.Data = "Data1"
	fmt.Printf("Got object: %+v\n", obj1)

	// Put object back to pool
	pool.Put(obj1)
	fmt.Println("Object put back to pool")

	// Get object again (might be the same object)
	obj2 := pool.Get().(*struct {
		ID   int
		Data string
	})
	fmt.Printf("Got object again: %+v\n", obj2)

	// Demonstrate concurrent usage
	fmt.Println("\nDemonstrating concurrent usage:")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			obj := pool.Get().(*struct {
				ID   int
				Data string
			})
			obj.ID = id
			obj.Data = fmt.Sprintf("Data%d", id)
			fmt.Printf("Goroutine %d: got object %+v\n", id, obj)
			time.Sleep(100 * time.Millisecond)
			pool.Put(obj)
			fmt.Printf("Goroutine %d: put object back\n", id)
		}(i)
	}
	wg.Wait()
}

// Example 7: Map concurrent safe mapping
func mapExample() {
	fmt.Println("\n--- Example 7: Map concurrent safe mapping ---")

	var syncMap sync.Map
	var wg sync.WaitGroup

	// Concurrent writes
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			value := fmt.Sprintf("value%d", id)
			syncMap.Store(key, value)
			fmt.Printf("Write: %s = %s\n", key, value)
		}(i)
	}

	// Concurrent reads
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			if value, ok := syncMap.Load(key); ok {
				fmt.Printf("Read: %s = %s\n", key, value)
			} else {
				fmt.Printf("Read: %s does not exist\n", key)
			}
		}(i)
	}

	wg.Wait()

	// Demonstrate other operations
	fmt.Println("\nDemonstrating other operations:")

	// LoadOrStore
	value, loaded := syncMap.LoadOrStore("key10", "value10")
	fmt.Printf("LoadOrStore: value=%v, loaded=%t\n", value, loaded)

	// Delete
	syncMap.Delete("key0")
	fmt.Println("Deleted key0")

	// Range iteration
	fmt.Println("Iterating all key-value pairs:")
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v = %v\n", key, value)
		return true // Continue iteration
	})
}

// Example 8: Comprehensive example
func comprehensiveSyncExample() {
	fmt.Println("\n--- Example 8: Comprehensive example ---")

	// Simulate a simple cache system
	type Cache struct {
		data  map[string]interface{}
		mu    sync.RWMutex
		stats struct {
			hits   int64
			misses int64
		}
		statsMu sync.RWMutex
	}

	cache := &Cache{
		data: make(map[string]interface{}),
	}

	// Start multiple goroutines for read-write operations
	var wg sync.WaitGroup

	// Write goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				key := fmt.Sprintf("key%d_%d", id, j)
				value := fmt.Sprintf("value%d_%d", id, j)

				cache.mu.Lock()
				cache.data[key] = value
				cache.mu.Unlock()

				fmt.Printf("Write: %s = %s\n", key, value)
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	// Read goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				key := fmt.Sprintf("key%d_%d", id%3, j)

				cache.mu.RLock()
				value, exists := cache.data[key]
				cache.mu.RUnlock()

				if exists {
					cache.statsMu.Lock()
					cache.stats.hits++
					cache.statsMu.Unlock()
					fmt.Printf("Hit: %s = %v\n", key, value)
				} else {
					cache.statsMu.Lock()
					cache.stats.misses++
					cache.statsMu.Unlock()
					fmt.Printf("Miss: %s\n", key)
				}

				time.Sleep(30 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()

	// Show statistics
	cache.statsMu.RLock()
	fmt.Printf("\nCache statistics: hits=%d, misses=%d\n",
		cache.stats.hits, cache.stats.misses)
	cache.statsMu.RUnlock()

	// Show final data
	fmt.Println("Final cache data:")
	cache.mu.RLock()
	for key, value := range cache.data {
		fmt.Printf("  %s = %v\n", key, value)
	}
	cache.mu.RUnlock()
}
