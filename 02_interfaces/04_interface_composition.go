package main

import "fmt"

// --- Types for interface composition examples ---
type Reader interface {
	Read() string
}
type Writer interface {
	Write(s string)
}
type ReadWriter interface {
	Reader
	Writer
}
type File struct{ content string }

func (f *File) Read() string   { return f.content }
func (f *File) Write(s string) { f.content = s }

type Handler interface {
	Handle(req string) string
}
type BaseHandler struct{}

func (BaseHandler) Handle(req string) string {
	return "Base: " + req
}

type LoggingMiddleware struct {
	Handler
}

func (l LoggingMiddleware) Handle(req string) string {
	fmt.Println("Logging request:", req)
	return l.Handler.Handle(req)
}

type A interface{ Foo() }
type B interface{ Foo() }
type AB interface {
	A
	B
}

// type MyType struct{}
// func (MyType) Foo() {} // Only one Foo() method, but AB embeds both A and B

type Product interface {
	Use() string
}
type ProductA struct{}

func (ProductA) Use() string { return "ProductA used" }

type ProductB struct{}

func (ProductB) Use() string { return "ProductB used" }

type Factory interface {
	Create() Product
}
type FactoryA struct{}

func (FactoryA) Create() Product { return ProductA{} }

type FactoryB struct{}

func (FactoryB) Create() Product { return ProductB{} }

// Example 1: Interface embedding for ReadWriter
func interfaceEmbeddingReadWriter() {
	fmt.Println("\n--- Example 1: Interface embedding for ReadWriter ---")
	var rw ReadWriter = &File{}
	rw.Write("Hello, interface composition!")
	fmt.Println("ReadWriter reads:", rw.Read())
}

// Example 2: Middleware/Decorator pattern with interface composition
func interfaceMiddlewarePattern() {
	fmt.Println("\n--- Example 2: Middleware/Decorator pattern ---")
	h := LoggingMiddleware{BaseHandler{}}
	fmt.Println(h.Handle("request"))
}

// Example 3: Error - Conflicting embedded interfaces
func interfaceEmbeddingConflict() {
	fmt.Println("\n--- Example 3: Error - Conflicting embedded interfaces ---")
	fmt.Println("Uncomment code to see ambiguous method error if both A and B are implemented differently.")
}

// Example 4: Real-world - Factory pattern with interface composition
func interfaceFactoryPattern() {
	fmt.Println("\n--- Example 4: Factory pattern with interface composition ---")
	factories := []Factory{FactoryA{}, FactoryB{}}
	for _, f := range factories {
		p := f.Create()
		fmt.Println(p.Use())
	}
}

// InterfaceCompositionExamples runs all interface composition examples
func InterfaceCompositionExamples() {
	interfaceEmbeddingReadWriter()
	interfaceMiddlewarePattern()
	interfaceEmbeddingConflict()
	interfaceFactoryPattern()
}
