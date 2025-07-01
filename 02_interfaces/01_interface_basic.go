package main

import (
	"errors"
	"fmt"
)

// --- Types for interface basic examples ---
type Notifier interface {
	Notify(msg string)
}
type Email struct {
	Address string
}

func (e Email) Notify(msg string) {
	fmt.Printf("Sending email to %s: %s\n", e.Address, msg)
}

type SMS struct{ Number string }

func (s SMS) Notify(msg string) {
	fmt.Printf("Sending SMS to %s: %s\n", s.Number, msg)
}

type Plugin interface {
	DoAction(data string) error
}
type LoggerPlugin struct{}

func (LoggerPlugin) DoAction(data string) error {
	fmt.Println("Logging:", data)
	return nil
}

type ValidatorPlugin struct{}

func (ValidatorPlugin) DoAction(data string) error {
	if data == "" {
		return errors.New("empty data")
	}
	fmt.Println("Validated:", data)
	return nil
}

// Example 1: Basic interface definition and implementation
func interfaceBasicDefinition() {
	fmt.Println("\n--- Example 1: Interface definition and implementation ---")
	var n Notifier = Email{"user@example.com"}
	n.Notify("Welcome!")
}

// Example 2: Real-world decoupling with interfaces
func interfaceDecouplingExample() {
	fmt.Println("\n--- Example 2: Decoupling with interfaces (real-world) ---")
	sendNotification := func(n Notifier, msg string) {
		n.Notify(msg)
	}
	sendNotification(SMS{"123456"}, "Your code is 1234")
	sendNotification(Email{"user@example.com"}, "Your code is 5678")
}

// Example 3: Error - Not implementing all interface methods
func interfaceNotImplementedError() {
	fmt.Println("\n--- Example 3: Not implementing all interface methods (compile error) ---")
	fmt.Println("Uncomment code to see compile error.")
	// type BadNotifier struct{}
	// var n Notifier = BadNotifier{} // Compile error: BadNotifier does not implement Notify
}

// Example 4: Interface zero value and nil pitfalls
func interfaceNilPitfall() {
	fmt.Println("\n--- Example 4: Interface zero value and nil pitfalls ---")
	var n Notifier
	fmt.Println("n == nil?", n == nil)
	// n.Notify("test") // Uncomment to see runtime panic: method call on nil interface
	fmt.Println("Uncomment code to see runtime panic.")
}

// Example 5: Real-world plugin pattern with interfaces
func interfacePluginPattern() {
	fmt.Println("\n--- Example 5: Plugin pattern with interfaces ---")
	plugins := []Plugin{LoggerPlugin{}, ValidatorPlugin{}}
	for _, p := range plugins {
		if err := p.DoAction("hello"); err != nil {
			fmt.Println("Plugin error:", err)
		}
	}
	// Error case
	for _, p := range plugins {
		if err := p.DoAction(""); err != nil {
			fmt.Println("Plugin error:", err)
		}
	}
}

// InterfaceBasicExamples runs all interface basic examples
func InterfaceBasicExamples() {
	interfaceBasicDefinition()
	interfaceDecouplingExample()
	interfaceNotImplementedError()
	interfaceNilPitfall()
	interfacePluginPattern()
}
