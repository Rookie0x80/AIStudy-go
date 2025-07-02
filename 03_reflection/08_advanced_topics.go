package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
	"unsafe"
)

// AdvancedTopicsExamples demonstrates advanced reflection topics
func AdvancedTopicsExamples() {
	fmt.Println("\n=== Advanced Topics Examples ===")

	// Example 1: Performance optimization techniques
	performanceOptimization()

	// Example 2: Reflection caching strategies
	reflectionCaching()

	// Example 3: Memory and safety considerations
	memorySafetyConsiderations()

	// Example 4: Concurrent reflection operations
	concurrentReflection()

	// Example 5: Debugging reflection code
	debuggingReflection()

	// Example 6: Advanced type manipulation
	advancedTypeManipulation()
}

// Example 1: Performance optimization techniques
func performanceOptimization() {
	fmt.Println("\n--- Example 1: Performance Optimization ---")

	// Benchmark different approaches
	fmt.Printf("Benchmarking different field access methods:\n")

	people := make([]Person, 1000)
	for i := range people {
		people[i] = Person{Name: fmt.Sprintf("Person%d", i), Age: 20 + i%50}
	}

	// Method 1: Direct access (baseline)
	start := time.Now()
	total1 := 0
	for _, person := range people {
		total1 += person.Age
	}
	direct := time.Since(start)

	// Method 2: Reflection without caching
	start = time.Now()
	total2 := 0
	for _, person := range people {
		v := reflect.ValueOf(person)
		ageField := v.FieldByName("Age")
		total2 += int(ageField.Int())
	}
	reflection := time.Since(start)

	// Method 3: Reflection with field index caching
	start = time.Now()
	total3 := 0
	personType := reflect.TypeOf(Person{})
	ageIndex := findFieldIndex(personType, "Age")
	for _, person := range people {
		v := reflect.ValueOf(person)
		ageField := v.Field(ageIndex)
		total3 += int(ageField.Int())
	}
	cached := time.Since(start)

	fmt.Printf("  Direct access: %v (total: %d)\n", direct, total1)
	fmt.Printf("  Reflection (no cache): %v (total: %d, %.1fx slower)\n",
		reflection, total2, float64(reflection)/float64(direct))
	fmt.Printf("  Reflection (cached): %v (total: %d, %.1fx slower)\n",
		cached, total3, float64(cached)/float64(direct))

	// Demonstrate compilation optimization
	demonstrateCompilerOptimization()
}

// Example 2: Reflection caching strategies
func reflectionCaching() {
	fmt.Println("\n--- Example 2: Reflection Caching ---")

	cache := NewReflectionCache()

	// Cache type information
	personType := reflect.TypeOf(Person{})
	cache.CacheType("Person", personType)

	// Cache field information
	cache.CacheFields("Person", personType)

	// Cache method information
	calcType := reflect.TypeOf(SimpleCalculator{})
	cache.CacheMethods("Calculator", calcType)

	// Use cached information
	fmt.Printf("Using cached reflection data:\n")

	person := Person{Name: "Alice", Age: 30}

	// Fast field access using cache
	if fieldInfo, exists := cache.GetField("Person", "Name"); exists {
		v := reflect.ValueOf(person)
		nameField := v.Field(fieldInfo.Index)
		fmt.Printf("  Name (cached): %s\n", nameField.String())
	}

	// Fast method call using cache
	calc := SimpleCalculator{Name: "TestCalc"}
	if methodInfo, exists := cache.GetMethod("Calculator", "Add"); exists {
		v := reflect.ValueOf(calc)
		method := v.Method(methodInfo.Index)
		args := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
		results := method.Call(args)
		fmt.Printf("  Add result (cached): %v\n", results[0].Interface())
	}
}

// Example 3: Memory and safety considerations
func memorySafetyConsiderations() {
	fmt.Println("\n--- Example 3: Memory & Safety Considerations ---")

	// Demonstrate memory layout understanding
	fmt.Printf("Memory layout analysis:\n")
	analyzeMemoryLayout()

	// Show unsafe reflection operations
	fmt.Printf("\nUnsafe operations (demonstration only):\n")
	demonstrateUnsafeOperations()

	// Security considerations
	fmt.Printf("\nSecurity considerations:\n")
	demonstrateSecurityConsiderations()
}

// Example 4: Concurrent reflection operations
func concurrentReflection() {
	fmt.Println("\n--- Example 4: Concurrent Reflection ---")

	// Thread-safe reflection cache
	safeCache := NewThreadSafeCache()

	var wg sync.WaitGroup
	numGoroutines := 10

	// Concurrent cache access
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Each goroutine performs reflection operations
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("type_%d_%d", id, j)
				safeCache.SetType(key, reflect.TypeOf(Person{}))

				if cached := safeCache.GetType(key); cached != nil {
					// Use cached type
					_ = cached.Name()
				}
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Concurrent reflection operations completed\n")
	fmt.Printf("Cache size: %d entries\n", safeCache.Size())

	// Demonstrate race condition without synchronization
	demonstrateReflectionRaceCondition()
}

// Example 5: Debugging reflection code
func debuggingReflection() {
	fmt.Println("\n--- Example 5: Debugging Reflection ---")

	// Debug reflection operations
	debugger := NewReflectionDebugger()

	person := Person{Name: "Alice", Age: 30}

	// Trace reflection operations
	fmt.Printf("Debugging reflection operations:\n")
	debugger.TraceFieldAccess(person, "Name")
	debugger.TraceFieldAccess(person, "Age")
	debugger.TraceFieldAccess(person, "NonExistent")

	// Debug type information
	fmt.Printf("\nDebugging type information:\n")
	debugger.DebugType(reflect.TypeOf(person))

	// Debug value information
	fmt.Printf("\nDebugging value information:\n")
	debugger.DebugValue(reflect.ValueOf(person))
}

// Example 6: Advanced type manipulation
func advancedTypeManipulation() {
	fmt.Println("\n--- Example 6: Advanced Type Manipulation ---")

	// Dynamic type creation
	fmt.Printf("Dynamic type creation:\n")
	demonstrateDynamicTypeCreation()

	// Type composition
	fmt.Printf("\nType composition:\n")
	demonstrateTypeComposition()

	// Interface synthesis
	fmt.Printf("\nInterface synthesis:\n")
	demonstrateInterfaceSynthesis()
}

// Helper implementations

func findFieldIndex(t reflect.Type, fieldName string) int {
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name == fieldName {
			return i
		}
	}
	return -1
}

func demonstrateCompilerOptimization() {
	fmt.Printf("\nCompiler optimization demonstration:\n")

	// Hot path that might benefit from optimization
	people := []Person{{Name: "Test", Age: 25}}

	// Reflection operations that could be optimized
	t := reflect.TypeOf(Person{})
	nameFieldIndex := findFieldIndex(t, "Name")
	ageFieldIndex := findFieldIndex(t, "Age")

	// Pre-compiled field accessors
	nameAccessor := func(p Person) string {
		v := reflect.ValueOf(p)
		return v.Field(nameFieldIndex).String()
	}

	ageAccessor := func(p Person) int {
		v := reflect.ValueOf(p)
		return int(v.Field(ageFieldIndex).Int())
	}

	for _, person := range people {
		name := nameAccessor(person)
		age := ageAccessor(person)
		fmt.Printf("  Optimized access: %s, %d\n", name, age)
	}
}

// Reflection Cache Implementation
type ReflectionCache struct {
	types   map[string]reflect.Type
	fields  map[string]map[string]FieldInfo
	methods map[string]map[string]MethodInfo
}

type FieldInfo struct {
	Index int
	Type  reflect.Type
	Tag   reflect.StructTag
}

type MethodInfo struct {
	Index int
	Type  reflect.Type
}

func NewReflectionCache() *ReflectionCache {
	return &ReflectionCache{
		types:   make(map[string]reflect.Type),
		fields:  make(map[string]map[string]FieldInfo),
		methods: make(map[string]map[string]MethodInfo),
	}
}

func (rc *ReflectionCache) CacheType(name string, t reflect.Type) {
	rc.types[name] = t
}

func (rc *ReflectionCache) CacheFields(typeName string, t reflect.Type) {
	fields := make(map[string]FieldInfo)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fields[field.Name] = FieldInfo{
			Index: i,
			Type:  field.Type,
			Tag:   field.Tag,
		}
	}

	rc.fields[typeName] = fields
}

func (rc *ReflectionCache) CacheMethods(typeName string, t reflect.Type) {
	methods := make(map[string]MethodInfo)

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		methods[method.Name] = MethodInfo{
			Index: i,
			Type:  method.Type,
		}
	}

	rc.methods[typeName] = methods
}

func (rc *ReflectionCache) GetField(typeName, fieldName string) (FieldInfo, bool) {
	if fields, exists := rc.fields[typeName]; exists {
		if field, exists := fields[fieldName]; exists {
			return field, true
		}
	}
	return FieldInfo{}, false
}

func (rc *ReflectionCache) GetMethod(typeName, methodName string) (MethodInfo, bool) {
	if methods, exists := rc.methods[typeName]; exists {
		if method, exists := methods[methodName]; exists {
			return method, true
		}
	}
	return MethodInfo{}, false
}

// Memory layout analysis
func analyzeMemoryLayout() {
	person := Person{Name: "Alice", Age: 30}

	fmt.Printf("  Person struct size: %d bytes\n", unsafe.Sizeof(person))
	fmt.Printf("  Name field offset: %d bytes\n", unsafe.Offsetof(person.Name))
	fmt.Printf("  Age field offset: %d bytes\n", unsafe.Offsetof(person.Age))

	// Show reflection overhead
	v := reflect.ValueOf(person)
	fmt.Printf("  reflect.Value size: %d bytes\n", unsafe.Sizeof(v))

	t := reflect.TypeOf(person)
	fmt.Printf("  reflect.Type size: %d bytes\n", unsafe.Sizeof(t))
}

func demonstrateUnsafeOperations() {
	// WARNING: These are dangerous operations for demonstration only
	person := Person{Name: "Alice", Age: 30}

	// Get pointer to struct
	ptr := unsafe.Pointer(&person)

	// Calculate offset to Age field (assuming it's at offset 16 on 64-bit)
	// This is extremely unsafe and platform-dependent
	ageOffset := unsafe.Offsetof(person.Age)
	agePtr := (*int)(unsafe.Pointer(uintptr(ptr) + ageOffset))

	fmt.Printf("  Original age: %d\n", person.Age)
	fmt.Printf("  Age via unsafe pointer: %d\n", *agePtr)

	// Modify via unsafe pointer (DON'T DO THIS IN REAL CODE)
	*agePtr = 99
	fmt.Printf("  Modified age: %d\n", person.Age)
}

func demonstrateSecurityConsiderations() {
	// Show potential security issues with reflection

	// 1. Access to unexported fields
	type sensitiveData struct {
		Public  string
		private string
	}

	data := sensitiveData{Public: "safe", private: "secret"}
	v := reflect.ValueOf(&data).Elem()

	// Can access unexported fields through reflection
	privateField := v.FieldByName("private")
	if privateField.IsValid() {
		// This accesses private data!
		fmt.Printf("  Accessed private field: %s\n", privateField.String())
	}

	// 2. Type confusion attacks
	var iface interface{} = "string"
	v2 := reflect.ValueOf(&iface).Elem()

	// Changing interface type (dangerous!)
	fmt.Printf("  Original type: %v\n", v2.Type())
	// Note: Actual type confusion would require more complex unsafe operations
}

// Thread-safe cache
type ThreadSafeCache struct {
	mu    sync.RWMutex
	types map[string]reflect.Type
}

func NewThreadSafeCache() *ThreadSafeCache {
	return &ThreadSafeCache{
		types: make(map[string]reflect.Type),
	}
}

func (tsc *ThreadSafeCache) SetType(key string, t reflect.Type) {
	tsc.mu.Lock()
	defer tsc.mu.Unlock()
	tsc.types[key] = t
}

func (tsc *ThreadSafeCache) GetType(key string) reflect.Type {
	tsc.mu.RLock()
	defer tsc.mu.RUnlock()
	return tsc.types[key]
}

func (tsc *ThreadSafeCache) Size() int {
	tsc.mu.RLock()
	defer tsc.mu.RUnlock()
	return len(tsc.types)
}

func demonstrateReflectionRaceCondition() {
	fmt.Printf("\nDemonstrating race condition (without synchronization):\n")

	// Unsafe cache (for demonstration)
	unsafeCache := make(map[string]reflect.Type)

	var wg sync.WaitGroup

	// This would cause race conditions in real code
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("race_%d", id)

			// Concurrent map access without synchronization
			// In real code, this would panic with "concurrent map writes"
			// We'll just read here to avoid the panic
			_ = unsafeCache[key]
		}(i)
	}

	wg.Wait()
	fmt.Printf("  Race condition demonstration completed (safely)\n")
}

// Reflection Debugger
type ReflectionDebugger struct{}

func NewReflectionDebugger() *ReflectionDebugger {
	return &ReflectionDebugger{}
}

func (rd *ReflectionDebugger) TraceFieldAccess(obj interface{}, fieldName string) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	fmt.Printf("  Tracing field access: %s.%s\n", t.Name(), fieldName)

	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		fmt.Printf("    ERROR: Field not found\n")
		return
	}

	fmt.Printf("    Field type: %v\n", field.Type())
	fmt.Printf("    Field kind: %v\n", field.Kind())
	fmt.Printf("    Can interface: %v\n", field.CanInterface())
	fmt.Printf("    Can addr: %v\n", field.CanAddr())
	fmt.Printf("    Can set: %v\n", field.CanSet())

	if field.CanInterface() {
		fmt.Printf("    Value: %v\n", field.Interface())
	}
}

func (rd *ReflectionDebugger) DebugType(t reflect.Type) {
	fmt.Printf("  Type debug: %s\n", t.Name())
	fmt.Printf("    Kind: %v\n", t.Kind())
	fmt.Printf("    Size: %d bytes\n", t.Size())
	fmt.Printf("    Align: %d bytes\n", t.Align())
	fmt.Printf("    Package: %s\n", t.PkgPath())
	fmt.Printf("    Comparable: %v\n", t.Comparable())

	if t.Kind() == reflect.Struct {
		fmt.Printf("    Fields: %d\n", t.NumField())
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fmt.Printf("      %d: %s (%v)\n", i, field.Name, field.Type)
		}
	}

	fmt.Printf("    Methods: %d\n", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("      %d: %s%s\n", i, method.Name, formatFunctionSignature(method.Type))
	}
}

func (rd *ReflectionDebugger) DebugValue(v reflect.Value) {
	t := v.Type()

	fmt.Printf("  Value debug: %v\n", t)
	fmt.Printf("    Kind: %v\n", v.Kind())
	fmt.Printf("    Type: %v\n", v.Type())
	fmt.Printf("    Is valid: %v\n", v.IsValid())

	if v.IsValid() {
		fmt.Printf("    Is zero: %v\n", v.IsZero())
		fmt.Printf("    Can interface: %v\n", v.CanInterface())
		fmt.Printf("    Can addr: %v\n", v.CanAddr())
		fmt.Printf("    Can set: %v\n", v.CanSet())

		switch v.Kind() {
		case reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func, reflect.Slice, reflect.Map:
			fmt.Printf("    Is nil: %v\n", v.IsNil())
		}
	}
}

func demonstrateDynamicTypeCreation() {
	// Create slice type dynamically
	intType := reflect.TypeOf(int(0))
	sliceType := reflect.SliceOf(intType)

	fmt.Printf("  Created slice type: %v\n", sliceType)

	// Create instance of dynamic type
	slice := reflect.MakeSlice(sliceType, 0, 5)
	slice = reflect.Append(slice, reflect.ValueOf(1))
	slice = reflect.Append(slice, reflect.ValueOf(2))

	fmt.Printf("  Dynamic slice: %v\n", slice.Interface())

	// Create map type dynamically
	mapType := reflect.MapOf(reflect.TypeOf(""), intType)
	m := reflect.MakeMap(mapType)
	m.SetMapIndex(reflect.ValueOf("key"), reflect.ValueOf(42))

	fmt.Printf("  Dynamic map: %v\n", m.Interface())
}

func demonstrateTypeComposition() {
	// Analyze composed types
	rwType := reflect.TypeOf((*ReadWriter)(nil)).Elem()

	fmt.Printf("  Analyzing composed interface: %s\n", rwType.Name())
	fmt.Printf("  Methods: %d\n", rwType.NumMethod())

	for i := 0; i < rwType.NumMethod(); i++ {
		method := rwType.Method(i)
		fmt.Printf("    %d: %s%s\n", i, method.Name, formatFunctionSignature(method.Type))
	}

	// Analyze component interfaces
	readerType := reflect.TypeOf((*Reader)(nil)).Elem()
	writerType := reflect.TypeOf((*Writer)(nil)).Elem()

	fmt.Printf("  Component Reader interface:\n")
	for i := 0; i < readerType.NumMethod(); i++ {
		method := readerType.Method(i)
		fmt.Printf("    %s%s\n", method.Name, formatFunctionSignature(method.Type))
	}

	fmt.Printf("  Component Writer interface:\n")
	for i := 0; i < writerType.NumMethod(); i++ {
		method := writerType.Method(i)
		fmt.Printf("    %s%s\n", method.Name, formatFunctionSignature(method.Type))
	}
}

func demonstrateInterfaceSynthesis() {
	// Check interface satisfaction dynamically
	calc := SimpleCalculator{}
	calcType := reflect.TypeOf(calc)

	// Check if type satisfies Calculator interface
	calculatorType := reflect.TypeOf((*Calculator)(nil)).Elem()

	fmt.Printf("  Type %s satisfies Calculator: %v\n",
		calcType.Name(), calcType.Implements(calculatorType))

	// Show required methods
	fmt.Printf("  Calculator interface requires:\n")
	for i := 0; i < calculatorType.NumMethod(); i++ {
		method := calculatorType.Method(i)
		fmt.Printf("    %s%s\n", method.Name, formatFunctionSignature(method.Type))
	}
}
