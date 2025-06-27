package main

import (
	"fmt"
	"sync"
	"time"
)

// GoroutinesExamples runs all goroutines examples
func GoroutinesExamples() {
	fmt.Println("=== Basic Goroutines Examples ===")

	// Example 1: Basic goroutine
	basicGoroutine()

	// Example 2: Goroutine with parameters
	goroutineWithParams()

	// Example 3: Anonymous function goroutine
	anonymousGoroutine()

	// Example 4: Closure trap
	closureTrap()

	// Example 5: Wait for goroutines to complete
	waitForGoroutines()

	// Example 6: Concurrent counter
	concurrentCounter()

	// Example 7: Worker pool pattern
	workerPool()

	// Example 8: Fan-out fan-in pattern
	fanOutFanIn()

	// Example 9: Common mistakes examples
	commonMistakes()
}

// Example 1: Basic goroutine
func basicGoroutine() {
	fmt.Println("\n--- Example 1: Basic goroutine ---")

	fmt.Println("Main goroutine starts")

	// Start a simple goroutine
	go func() {
		fmt.Println("Child goroutine executing...")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Child goroutine completed")
	}()

	// Main goroutine continues execution
	fmt.Println("Main goroutine continues execution")
	time.Sleep(200 * time.Millisecond) // Wait for child goroutine to complete
	fmt.Println("Main goroutine ends")
}

// Example 2: Goroutine with parameters
func goroutineWithParams() {
	fmt.Println("\n--- Example 2: Goroutine with parameters ---")

	for i := 1; i <= 3; i++ {
		// Correct way: pass parameters to goroutine
		go func(id int) {
			fmt.Printf("Worker %d starts work\n", id)
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)
			fmt.Printf("Worker %d completes work\n", id)
		}(i)
	}

	time.Sleep(500 * time.Millisecond)
}

// Example 3: Anonymous function goroutine
func anonymousGoroutine() {
	fmt.Println("\n--- Example 3: Anonymous function goroutine ---")

	// Use anonymous function to create goroutine
	go func(name string) {
		fmt.Printf("Hello, %s! From goroutine\n", name)
	}("World")

	time.Sleep(100 * time.Millisecond)
}

// Example 4: Closure trap
func closureTrap() {
	fmt.Println("\n--- Example 4: Closure trap ---")

	fmt.Println("Wrong example - Closure trap:")
	for i := 1; i <= 3; i++ {
		// Wrong: all goroutines will capture the value of i when the loop ends
		go func() {
			fmt.Printf("Wrong closure: i = %d\n", i)
		}()
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("\nCorrect example - Pass parameters:")
	for i := 1; i <= 3; i++ {
		// Correct: pass i as parameter
		go func(id int) {
			fmt.Printf("Correct parameter passing: id = %d\n", id)
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
}

// Example 5: Wait for goroutines to complete
func waitForGoroutines() {
	fmt.Println("\n--- Example 5: Wait for goroutines to complete ---")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment wait counter

		go func(id int) {
			defer wg.Done() // Decrement counter when done
			fmt.Printf("Worker %d starts\n", id)
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)
			fmt.Printf("Worker %d completes\n", id)
		}(i)
	}

	fmt.Println("Waiting for all goroutines to complete...")
	wg.Wait() // Wait for all goroutines to complete
	fmt.Println("All goroutines completed")
}

// Example 6: Concurrent counter
func concurrentCounter() {
	fmt.Println("\n--- Example 6: Concurrent counter ---")

	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Start multiple goroutines to increment counter simultaneously
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d (Expected: 10000)\n", counter)
}

// Example 7: Worker pool pattern
func workerPool() {
	fmt.Println("\n--- Example 7: Worker pool pattern ---")

	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start worker pool
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, results)
	}

	// Send jobs
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	// Collect results
	for i := 1; i <= numJobs; i++ {
		result := <-results
		fmt.Printf("Received result: %d\n", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond) // Simulate work
		results <- job * 2
	}
}

// Example 8: Fan-out fan-in pattern
func fanOutFanIn() {
	fmt.Println("\n--- Example 8: Fan-out fan-in pattern ---")

	// Generate data
	input := generate(1, 2, 3, 4, 5)

	// Fan-out: distribute input to multiple workers
	c1 := square(input)
	c2 := square(input)

	// Fan-in: merge outputs from multiple channels
	output := merge(c1, c2)

	// Consume results
	for result := range output {
		fmt.Printf("Result: %d\n", result)
	}
}

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Example 9: Common mistakes examples
func commonMistakes() {
	fmt.Println("\n--- Example 9: Common mistakes examples ---")

	fmt.Println("Mistake 1: Not waiting for goroutines to complete")
	go func() {
		fmt.Println("This goroutine may not complete execution")
	}()
	// Main program ends immediately, goroutine may be terminated

	fmt.Println("Mistake 2: Starting too many goroutines in a loop")
	// This may cause resource exhaustion
	// for i := 0; i < 1000000; i++ {
	//     go func() { /* work */ }()
	// }

	fmt.Println("Mistake 3: Forgetting to close channels")
	ch := make(chan int)
	go func() {
		ch <- 1
		// Forgot to close(ch)
	}()

	// Correct approach: use defer to close channel
	go func() {
		defer close(ch)
		ch <- 1
	}()

	fmt.Println("Mistake 4: Panic in goroutine")
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Caught panic: %v\n", r)
			}
		}()
		panic("panic in goroutine")
	}()

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Mistake 5: Race condition")
	var shared int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			shared++ // Race condition!
		}()
	}

	wg.Wait()
	fmt.Printf("Shared variable value: %d (may not be 1000)\n", shared)

	// Correct approach: use mutex
	var safeShared int
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			safeShared++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("Safe shared variable value: %d\n", safeShared)
}
