package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

// Basic types for reflection examples
type Person struct {
	Name    string  `json:"name" validate:"required"`
	Age     int     `json:"age" validate:"min=0,max=150"`
	Email   string  `json:"email" validate:"email"`
	Address Address `json:"address"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country" validate:"required"`
}

// Embedded struct example
type Employee struct {
	Person
	ID       int     `json:"id" validate:"required"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary" validate:"min=0"`
}

// Interface types for reflection
type Writer interface {
	Write(data []byte) (int, error)
}

type Reader interface {
	Read(data []byte) (int, error)
}

type ReadWriter interface {
	Reader
	Writer
}

// Calculator interface for method reflection
type Calculator interface {
	Add(a, b int) int
	Subtract(a, b int) int
	Multiply(a, b int) int
	Divide(a, b int) (int, error)
}

// Simple calculator implementation
type SimpleCalculator struct {
	Name string
}

func (c SimpleCalculator) Add(a, b int) int {
	return a + b
}

func (c SimpleCalculator) Subtract(a, b int) int {
	return a - b
}

func (c SimpleCalculator) Multiply(a, b int) int {
	return a * b
}

func (c SimpleCalculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Handler function type for function reflection
type Handler func(ctx context.Context, data string) error

// Event system types
type Event struct {
	Type      string                 `json:"type"`
	Timestamp time.Time              `json:"timestamp"`
	Data      map[string]interface{} `json:"data"`
}

type EventHandler interface {
	Handle(event Event) error
}

type LogEventHandler struct {
	Name string
}

func (h LogEventHandler) Handle(event Event) error {
	fmt.Printf("[%s] Event: %s at %v\n", h.Name, event.Type, event.Timestamp)
	return nil
}

// Generic container for value operations
type Container struct {
	Items []interface{}     `json:"items"`
	Meta  map[string]string `json:"meta"`
}

// Database model with tags
type User struct {
	ID        int       `db:"id" json:"id"`
	Username  string    `db:"username" json:"username" validate:"required,min=3"`
	Password  string    `db:"password" json:"-" validate:"required,min=8"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// Configuration struct for pattern examples
type Config struct {
	Database DatabaseConfig `json:"database" mapstructure:"database"`
	Server   ServerConfig   `json:"server" mapstructure:"server"`
	Redis    RedisConfig    `json:"redis" mapstructure:"redis"`
}

type DatabaseConfig struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"database" mapstructure:"database"`
}

type ServerConfig struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
}

type RedisConfig struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	Password string `json:"password" mapstructure:"password"`
}

// Service interfaces for dependency injection examples
type UserService interface {
	GetUser(id int) (*User, error)
	CreateUser(user *User) error
}

type EmailService interface {
	SendEmail(to, subject, body string) error
}

type UserServiceImpl struct {
	db    DatabaseService
	email EmailService
}

func (s *UserServiceImpl) GetUser(id int) (*User, error) {
	return &User{ID: id, Username: "test"}, nil
}

func (s *UserServiceImpl) CreateUser(user *User) error {
	return nil
}

type DatabaseService interface {
	Query(sql string, args ...interface{}) ([]map[string]interface{}, error)
	Execute(sql string, args ...interface{}) error
}

// Validator interface and implementation
type Validator interface {
	Validate(value interface{}) error
}

type FieldValidator struct {
	Rules map[string]string
}

func (v *FieldValidator) Validate(value interface{}) error {
	return nil // Implementation would validate based on struct tags
}

// Mapper interface for object mapping patterns
type Mapper interface {
	Map(src, dst interface{}) error
}

type ReflectionMapper struct{}

func (m *ReflectionMapper) Map(src, dst interface{}) error {
	return nil // Implementation would use reflection to map fields
}

// Factory interface for dynamic creation
type Factory interface {
	Create(typeName string) (interface{}, error)
}

type ReflectionFactory struct {
	types map[string]reflect.Type
}

func (f *ReflectionFactory) Create(typeName string) (interface{}, error) {
	if t, exists := f.types[typeName]; exists {
		return reflect.New(t).Interface(), nil
	}
	return nil, fmt.Errorf("unknown type: %s", typeName)
}

// Performance testing types
type BenchmarkStruct struct {
	Field1 string
	Field2 int
	Field3 bool
	Field4 float64
}

// Error types for common mistakes examples
type CustomError struct {
	Message string
	Code    int
}

func (e CustomError) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}
