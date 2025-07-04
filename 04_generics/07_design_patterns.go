package main

import (
	"fmt"
	"sync"
)

// ==========================================
// Generic Factory Pattern
// ==========================================

// Factory interface for creating objects
type Factory[T any] interface {
	Create() T
}

// ConfigurableFactory allows parameterized object creation
type ConfigurableFactory[T any, Config any] interface {
	CreateWithConfig(Config) T
}

// UserFactory creates different types of users
type UserFactory[T any] struct {
	createFunc func() T
}

func NewUserFactory[T any](createFunc func() T) *UserFactory[T] {
	return &UserFactory[T]{createFunc: createFunc}
}

func (f *UserFactory[T]) Create() T {
	return f.createFunc()
}

// DatabaseFactory creates database connections with configuration
type DatabaseFactory[T any] struct {
	createFunc func(string) T
}

func NewDatabaseFactory[T any](createFunc func(string) T) *DatabaseFactory[T] {
	return &DatabaseFactory[T]{createFunc: createFunc}
}

func (f *DatabaseFactory[T]) CreateWithConfig(config string) T {
	return f.createFunc(config)
}

// ==========================================
// Generic Builder Pattern
// ==========================================

// Builder interface for step-by-step construction
type Builder[T any] interface {
	Build() T
}

// GenericBuilder provides a flexible building pattern
type GenericBuilder[T any] struct {
	buildFunc func() T
	steps     []func()
}

func NewGenericBuilder[T any](buildFunc func() T) *GenericBuilder[T] {
	return &GenericBuilder[T]{buildFunc: buildFunc}
}

func (b *GenericBuilder[T]) AddStep(step func()) *GenericBuilder[T] {
	b.steps = append(b.steps, step)
	return b
}

func (b *GenericBuilder[T]) Build() T {
	for _, step := range b.steps {
		step()
	}
	return b.buildFunc()
}

// ConfigBuilder builds configuration objects
type ConfigBuilder[T any] struct {
	config T
}

func NewConfigBuilder[T any](initial T) *ConfigBuilder[T] {
	return &ConfigBuilder[T]{config: initial}
}

func (cb *ConfigBuilder[T]) Set(modifier func(*T)) *ConfigBuilder[T] {
	modifier(&cb.config)
	return cb
}

func (cb *ConfigBuilder[T]) Build() T {
	return cb.config
}

// ==========================================
// Generic Decorator Pattern
// ==========================================

// Component interface that can be decorated
type Component[T any] interface {
	Execute() T
}

// BaseComponent provides basic functionality
type BaseComponent[T any] struct {
	value T
}

func NewBaseComponent[T any](value T) *BaseComponent[T] {
	return &BaseComponent[T]{value: value}
}

func (bc *BaseComponent[T]) Execute() T {
	return bc.value
}

// Decorator wraps a component with additional behavior
type Decorator[T any] struct {
	component Component[T]
	wrapper   func(T) T
}

func NewDecorator[T any](component Component[T], wrapper func(T) T) *Decorator[T] {
	return &Decorator[T]{
		component: component,
		wrapper:   wrapper,
	}
}

func (d *Decorator[T]) Execute() T {
	result := d.component.Execute()
	return d.wrapper(result)
}

// ==========================================
// Generic Observer Pattern
// ==========================================

// Observer interface for receiving notifications
type Observer[T any] interface {
	Update(T)
}

// Subject interface for managing observers
type Subject[T any] interface {
	Attach(Observer[T])
	Detach(Observer[T])
	Notify(T)
}

// GenericSubject implements the Subject interface
type GenericSubject[T any] struct {
	observers []Observer[T]
	mu        sync.RWMutex
}

func NewGenericSubject[T any]() *GenericSubject[T] {
	return &GenericSubject[T]{
		observers: make([]Observer[T], 0),
	}
}

func (s *GenericSubject[T]) Attach(observer Observer[T]) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.observers = append(s.observers, observer)
}

func (s *GenericSubject[T]) Detach(observer Observer[T]) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *GenericSubject[T]) Notify(data T) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, observer := range s.observers {
		observer.Update(data)
	}
}

// ConcreteObserver implements Observer
type ConcreteObserver[T any] struct {
	id       string
	callback func(string, T)
}

func NewConcreteObserver[T any](id string, callback func(string, T)) *ConcreteObserver[T] {
	return &ConcreteObserver[T]{id: id, callback: callback}
}

func (o *ConcreteObserver[T]) Update(data T) {
	o.callback(o.id, data)
}

// ==========================================
// Generic Strategy Pattern
// ==========================================

// Strategy interface for interchangeable algorithms
type Strategy[T, R any] interface {
	Execute(T) R
}

// Context uses strategies to perform operations
type Context[T, R any] struct {
	strategy Strategy[T, R]
}

func NewContext[T, R any](strategy Strategy[T, R]) *Context[T, R] {
	return &Context[T, R]{strategy: strategy}
}

func (c *Context[T, R]) SetStrategy(strategy Strategy[T, R]) {
	c.strategy = strategy
}

func (c *Context[T, R]) ExecuteStrategy(data T) R {
	return c.strategy.Execute(data)
}

// ConcreteStrategy implements a specific algorithm
type ConcreteStrategy[T, R any] struct {
	algorithm func(T) R
}

func NewConcreteStrategy[T, R any](algorithm func(T) R) *ConcreteStrategy[T, R] {
	return &ConcreteStrategy[T, R]{algorithm: algorithm}
}

func (s *ConcreteStrategy[T, R]) Execute(data T) R {
	return s.algorithm(data)
}

// ==========================================
// Generic Command Pattern
// ==========================================

// Command interface for encapsulating requests
type Command[T any] interface {
	Execute() T
	Undo() T
}

// GenericCommand implements Command with functions
type GenericCommand[T any] struct {
	executeFunc func() T
	undoFunc    func() T
}

func NewGenericCommand[T any](executeFunc, undoFunc func() T) *GenericCommand[T] {
	return &GenericCommand[T]{
		executeFunc: executeFunc,
		undoFunc:    undoFunc,
	}
}

func (c *GenericCommand[T]) Execute() T {
	return c.executeFunc()
}

func (c *GenericCommand[T]) Undo() T {
	return c.undoFunc()
}

// CommandInvoker manages and executes commands
type CommandInvoker[T any] struct {
	commands []Command[T]
	history  []T
}

func NewCommandInvoker[T any]() *CommandInvoker[T] {
	return &CommandInvoker[T]{
		commands: make([]Command[T], 0),
		history:  make([]T, 0),
	}
}

func (ci *CommandInvoker[T]) AddCommand(command Command[T]) {
	ci.commands = append(ci.commands, command)
}

func (ci *CommandInvoker[T]) ExecuteAll() []T {
	results := make([]T, 0, len(ci.commands))
	for _, cmd := range ci.commands {
		result := cmd.Execute()
		results = append(results, result)
		ci.history = append(ci.history, result)
	}
	return results
}

func (ci *CommandInvoker[T]) UndoLast() T {
	if len(ci.commands) == 0 {
		var zero T
		return zero
	}

	lastCmd := ci.commands[len(ci.commands)-1]
	ci.commands = ci.commands[:len(ci.commands)-1]
	return lastCmd.Undo()
}

// ==========================================
// Generic Repository Pattern
// ==========================================

// Repository interface for data access
type Repository[T any, ID comparable] interface {
	Save(T) ID
	FindByID(ID) (T, bool)
	FindAll() []T
	Delete(ID) bool
	Update(ID, T) bool
}

// MemoryRepository implements Repository with in-memory storage
type MemoryRepository[T any, ID comparable] struct {
	data   map[ID]T
	nextID func() ID
	mu     sync.RWMutex
}

func NewMemoryRepository[T any, ID comparable](nextIDFunc func() ID) *MemoryRepository[T, ID] {
	return &MemoryRepository[T, ID]{
		data:   make(map[ID]T),
		nextID: nextIDFunc,
	}
}

func (r *MemoryRepository[T, ID]) Save(entity T) ID {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := r.nextID()
	r.data[id] = entity
	return id
}

func (r *MemoryRepository[T, ID]) FindByID(id ID) (T, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	entity, exists := r.data[id]
	return entity, exists
}

func (r *MemoryRepository[T, ID]) FindAll() []T {
	r.mu.RLock()
	defer r.mu.RUnlock()

	entities := make([]T, 0, len(r.data))
	for _, entity := range r.data {
		entities = append(entities, entity)
	}
	return entities
}

func (r *MemoryRepository[T, ID]) Delete(id ID) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[id]; exists {
		delete(r.data, id)
		return true
	}
	return false
}

func (r *MemoryRepository[T, ID]) Update(id ID, entity T) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[id]; exists {
		r.data[id] = entity
		return true
	}
	return false
}

// ==========================================
// Generic Adapter Pattern
// ==========================================

// Target interface that client expects
type Target[T any] interface {
	Request() T
}

// Adaptee has incompatible interface
type Adaptee[T any] struct {
	data T
}

func NewAdaptee[T any](data T) *Adaptee[T] {
	return &Adaptee[T]{data: data}
}

func (a *Adaptee[T]) SpecificRequest() T {
	return a.data
}

// Adapter makes Adaptee compatible with Target
type Adapter[T any] struct {
	adaptee *Adaptee[T]
}

func NewAdapter[T any](adaptee *Adaptee[T]) *Adapter[T] {
	return &Adapter[T]{adaptee: adaptee}
}

func (a *Adapter[T]) Request() T {
	return a.adaptee.SpecificRequest()
}

// ==========================================
// Examples and Demo Types
// ==========================================

// AppUser example type (avoiding conflict with User in other files)
type AppUser struct {
	ID   int
	Name string
	Age  int
}

func (u AppUser) String() string {
	return fmt.Sprintf("AppUser{ID: %d, Name: %s, Age: %d}", u.ID, u.Name, u.Age)
}

// DatabaseConnection example type
type DatabaseConnection struct {
	ConnectionString string
	Connected        bool
}

func (db DatabaseConnection) String() string {
	return fmt.Sprintf("DB{%s, Connected: %t}", db.ConnectionString, db.Connected)
}

// Configuration example type
type Configuration struct {
	Host     string
	Port     int
	Database string
	SSL      bool
}

// ==========================================
// Main Example Function
// ==========================================

func runDesignPatternsExample() {
	fmt.Println("\n🔸 Generic Factory Pattern")

	// User factory
	userFactory := NewUserFactory(func() AppUser {
		return AppUser{ID: 1, Name: "John Doe", Age: 30}
	})
	user := userFactory.Create()
	fmt.Printf("Created user: %s\n", user.String())

	// Database factory with configuration
	dbFactory := NewDatabaseFactory(func(config string) DatabaseConnection {
		return DatabaseConnection{ConnectionString: config, Connected: true}
	})
	db := dbFactory.CreateWithConfig("localhost:5432/mydb")
	fmt.Printf("Created DB: %s\n", db.String())

	fmt.Println("\n🔸 Generic Builder Pattern")

	// Generic builder for strings
	var result string
	builder := NewGenericBuilder(func() string { return result }).
		AddStep(func() { result += "Hello" }).
		AddStep(func() { result += " " }).
		AddStep(func() { result += "World" }).
		AddStep(func() { result += "!" })

	builtString := builder.Build()
	fmt.Printf("Built string: %s\n", builtString)

	// Configuration builder
	configBuilder := NewConfigBuilder(Configuration{}).
		Set(func(c *Configuration) { c.Host = "localhost" }).
		Set(func(c *Configuration) { c.Port = 8080 }).
		Set(func(c *Configuration) { c.Database = "myapp" }).
		Set(func(c *Configuration) { c.SSL = true })

	config := configBuilder.Build()
	fmt.Printf("Built config: %+v\n", config)

	fmt.Println("\n🔸 Generic Decorator Pattern")

	// Base component with string
	base := NewBaseComponent("Hello")

	// Add decorators
	decorated := NewDecorator(base, func(s string) string { return s + " World" })
	doubleDecorated := NewDecorator(decorated, func(s string) string { return s + "!" })

	fmt.Printf("Original: %s\n", base.Execute())
	fmt.Printf("Decorated: %s\n", decorated.Execute())
	fmt.Printf("Double decorated: %s\n", doubleDecorated.Execute())

	fmt.Println("\n🔸 Generic Observer Pattern")

	// Create subject
	subject := NewGenericSubject[string]()

	// Create observers
	observer1 := NewConcreteObserver("Observer1", func(id, data string) {
		fmt.Printf("%s received: %s\n", id, data)
	})
	observer2 := NewConcreteObserver("Observer2", func(id, data string) {
		fmt.Printf("%s got notification: %s\n", id, data)
	})

	// Attach observers
	subject.Attach(observer1)
	subject.Attach(observer2)

	// Notify observers
	subject.Notify("Important event occurred!")

	// Detach one observer
	subject.Detach(observer1)
	subject.Notify("Second event")

	fmt.Println("\n🔸 Generic Strategy Pattern")

	// Different sorting strategies
	bubbleSort := NewConcreteStrategy(func(data []int) []int {
		result := make([]int, len(data))
		copy(result, data)
		BubbleSort(result)
		return result
	})

	quickSort := NewConcreteStrategy(func(data []int) []int {
		result := make([]int, len(data))
		copy(result, data)
		QuickSort(result)
		return result
	})

	context := NewContext[[]int, []int](bubbleSort)
	testData := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Printf("Original: %v\n", testData)
	fmt.Printf("Bubble sort: %v\n", context.ExecuteStrategy(testData))

	context.SetStrategy(quickSort)
	fmt.Printf("Quick sort: %v\n", context.ExecuteStrategy(testData))

	fmt.Println("\n🔸 Generic Command Pattern")

	// Create commands
	var counter int
	incrementCmd := NewGenericCommand(
		func() int { counter++; return counter },
		func() int { counter--; return counter },
	)
	doubleCmd := NewGenericCommand(
		func() int { counter *= 2; return counter },
		func() int { counter /= 2; return counter },
	)

	// Create invoker and add commands
	invoker := NewCommandInvoker[int]()
	invoker.AddCommand(incrementCmd)
	invoker.AddCommand(incrementCmd)
	invoker.AddCommand(doubleCmd)

	fmt.Printf("Initial counter: %d\n", counter)
	results := invoker.ExecuteAll()
	fmt.Printf("Execution results: %v\n", results)
	fmt.Printf("Final counter: %d\n", counter)

	undoResult := invoker.UndoLast()
	fmt.Printf("Undo result: %d\n", undoResult)
	fmt.Printf("Counter after undo: %d\n", counter)

	fmt.Println("\n🔸 Generic Repository Pattern")

	// Create user repository
	var idCounter int
	userRepo := NewMemoryRepository[AppUser, int](func() int {
		idCounter++
		return idCounter
	})

	// Save users
	id1 := userRepo.Save(AppUser{Name: "Alice", Age: 30})
	id2 := userRepo.Save(AppUser{Name: "Bob", Age: 25})
	id3 := userRepo.Save(AppUser{Name: "Charlie", Age: 35})

	fmt.Printf("Saved users with IDs: %d, %d, %d\n", id1, id2, id3)

	// Find user by ID
	if user, found := userRepo.FindByID(2); found {
		fmt.Printf("Found user: %s\n", user.String())
	}

	// Find all users
	allUsers := userRepo.FindAll()
	fmt.Printf("All users: %d total\n", len(allUsers))
	for _, u := range allUsers {
		fmt.Printf("  %s\n", u.String())
	}

	// Update user
	userRepo.Update(2, AppUser{Name: "Bob Updated", Age: 26})
	if updated, found := userRepo.FindByID(2); found {
		fmt.Printf("Updated user: %s\n", updated.String())
	}

	// Delete user
	userRepo.Delete(3)
	fmt.Printf("Users after deletion: %d total\n", len(userRepo.FindAll()))

	fmt.Println("\n🔸 Generic Adapter Pattern")

	// Create adaptee with incompatible interface
	adaptee := NewAdaptee("Legacy Data")

	// Create adapter
	adapter := NewAdapter(adaptee)

	// Use adapter through Target interface
	var target Target[string] = adapter
	fmt.Printf("Adapted result: %s\n", target.Request())

	fmt.Println("\n✅ Design patterns examples completed!")
}
