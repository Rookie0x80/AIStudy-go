package main

import "fmt"

// --- Types for implicit implementation examples ---
type Storage interface {
	Save(data string) error
}
type FileStorage struct{}

func (FileStorage) Save(data string) error {
	fmt.Println("Saving to file:", data)
	return nil
}

type DBStorage struct{}

func (DBStorage) Save(data string) error {
	fmt.Println("Saving to database:", data)
	return nil
}

type Payment interface {
	Pay(amount float64) bool
}
type WechatPay struct{}

func (WechatPay) Pay(amount float64) bool {
	fmt.Printf("WechatPay: paid %.2f\n", amount)
	return true
}

type Alipay struct{}

func (Alipay) Pay(amount float64) bool {
	fmt.Printf("Alipay: paid %.2f\n", amount)
	return true
}

type Worker interface {
	Work()
}
type Person struct{ Name string }

func (p *Person) Work() {
	fmt.Printf("%s is working\n", p.Name)
}

type Clock interface {
	Now() int64
}
type RealClock struct{}

func (RealClock) Now() int64 { return 1234567890 }

type MockClock struct{ Value int64 }

func (m MockClock) Now() int64 { return m.Value }

type Service interface {
	Do()
}
type MyService struct{}

func (s *MyService) Do() { fmt.Println("MyService doing work") }

// Example 1: Multiple types implicitly implement the same interface
func implicitImplMultipleTypes() {
	fmt.Println("\n--- Example 1: Multiple types implement the same interface ---")
	storages := []Storage{FileStorage{}, DBStorage{}}
	for _, s := range storages {
		s.Save("important data")
	}
}

// Example 2: Polymorphism in real-world business logic
func implicitImplPolymorphism() {
	fmt.Println("\n--- Example 2: Polymorphism in business logic ---")
	checkout := func(p Payment, amount float64) {
		if !p.Pay(amount) {
			fmt.Println("Payment failed!")
		}
	}
	checkout(WechatPay{}, 100)
	checkout(Alipay{}, 200)
}

// Example 3: Method set and pointer vs value receiver
func implicitImplMethodSet() {
	fmt.Println("\n--- Example 3: Method set and pointer/value receiver ---")
	// var w Worker = Person{"Tom"} // Compile error: Person (value) does not implement Work() (pointer receiver)
	var w Worker = &Person{"Tom"}
	w.Work()
	fmt.Println("Uncomment code to see compile error for value receiver.")
}

// Example 4: Mocking in tests with interfaces
func implicitImplMockExample() {
	fmt.Println("\n--- Example 4: Using interfaces for mocking in tests ---")
	printTime := func(c Clock) {
		fmt.Println("Current time:", c.Now())
	}
	printTime(RealClock{})
	printTime(MockClock{Value: 42})
}

// Example 5: Error - assigning nil concrete value to interface
func implicitImplNilConcrete() {
	fmt.Println("\n--- Example 5: Pitfall - assigning nil concrete value to interface ---")
	var s *MyService = nil
	var svc Service = s
	fmt.Println("svc == nil?", svc == nil) // false! (interface holds type+value)
	// svc.Do() // Uncomment to see panic: method call on nil pointer
	fmt.Println("Uncomment code to see runtime panic.")
}

// ImplicitImplExamples runs all implicit implementation examples
func ImplicitImplExamples() {
	implicitImplMultipleTypes()
	implicitImplPolymorphism()
	implicitImplMethodSet()
	implicitImplMockExample()
	implicitImplNilConcrete()
}
