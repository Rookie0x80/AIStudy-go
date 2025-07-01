# Go Interfaces & Polymorphism Examples

[English](#english) | [中文](#chinese)

---

## English

This directory contains a comprehensive and practical collection of Go interface and polymorphism examples. Each topic includes real-world scenarios, best practices, and common pitfalls, helping you deeply understand Go's type system and interface features.

### 🚀 Quick Start

```bash
# View all available examples
go run .

# Run specific example (1-5)
go run . <example_number>
```

### 🧠 Core Concepts & Coverage
- **Interface Definition & Implementation**: Basic usage, decoupling, plugin pattern, interface zero value, nil pitfalls, compile-time errors.
- **Implicit Implementation & Polymorphism**: Multiple types, business polymorphism, pointer/value receiver, interface in testing/mocking, nil concrete pitfalls.
- **Type Assertion & Type Switch**: Safe/unsafe assertion, type switch for logging and business dispatch, anti-patterns.
- **Interface Composition**: Embedding, ReadWriter, middleware/decorator, factory pattern, method conflict errors.
- **Empty Interface & Any**: Generic containers, type switch, JSON decode, type safety pitfalls, reflection.

### 📁 Project Structure

| File | Topic | Core Content |
|------|-------|--------------|
| `01_interface_basic.go` | Interface Basics | Definition, implementation, decoupling, plugin, nil pitfalls, compile errors |
| `02_implicit_impl.go` | Implicit Implementation | Multiple types, business polymorphism, pointer/value receiver, mocking, nil pitfalls |
| `03_type_assertion.go` | Type Assertion | Safe/unsafe assertion, type switch, anti-patterns, business dispatch |
| `04_interface_composition.go` | Interface Composition | Embedding, ReadWriter, middleware, factory, conflict errors |
| `05_empty_interface.go` | Empty Interface & Any | Generic containers, type switch, JSON decode, type safety, reflection |

### 🎯 Learning Path

#### Basic Stage (1-2)
- **Interface Definition & Implementation**: Understand how interfaces are defined, implemented, and used for decoupling and extensibility.
- **Polymorphism**: Use interfaces for flexible, testable, and mockable code.

#### Intermediate Stage (3-4)
- **Type Assertion & Type Switch**: Extract and check concrete types, dynamic dispatch, and avoid common anti-patterns.
- **Interface Composition**: Build complex abstractions, middleware, and extensible factories.

#### Advanced Stage (5)
- **Empty Interface & Dynamic Typing**: Use `interface{}` for generic code, understand type safety, and leverage reflection.

### 💡 Usage Recommendations

#### Running Examples
```bash
# View help information
go run .

# Run examples by number
go run . 1    # Interface basics
go run . 3    # Type assertion
go run . 5    # Empty interface
```

#### Learning Suggestions
1. **Step by Step**: Learn in order of file numbers, build a solid foundation before learning advanced patterns
2. **Hands-on Practice**: Modify example parameters and observe different results
3. **Comparative Learning**: Focus on comparing correct examples with error examples
4. **Understand Principles**: Deeply understand the applicable scenarios and pros/cons of each pattern

### 🔍 Example Features

Each example file includes:
- ✅ **Basic Usage**: Practical application of core concepts
- ✅ **Advanced Techniques**: Real-world business scenarios and best practices
- ✅ **Error Examples**: Common pitfalls and avoidance methods
- ✅ **Comprehensive Cases**: Typical scenarios in actual development
- ✅ **Detailed Comments**: Complete explanation of code logic

### ✅ Covered Features & Scenarios
- Interface definition, implementation, and assignment
- Decoupling and plugin patterns
- Interface zero value and nil pitfalls
- Compile-time and runtime errors
- Multiple types and business polymorphism
- Pointer vs value receiver method sets
- Interface-based mocking for testing
- Type assertion (safe/unsafe), type switch, anti-patterns
- Interface composition, embedding, ReadWriter, middleware, factory
- Empty interface/any for generic containers, JSON decode, reflection

### 🚩 What's Not Covered (Advanced Topics)
- Recursive/nested interfaces
- Deep dive into interface dynamic type/value (reflect)
- Interface and nil: advanced edge cases
- Interface performance considerations
- Interface and generics (Go 1.18+ constraints)
- Interface comparability and panic cases
- Advanced reflection and dynamic proxy
- Interface serialization/deserialization caveats
- Standard library interface best practices (e.g., io.Reader, error)

If you need examples or explanations for these advanced topics, feel free to request!

### ⚠️ Important Notes

- Interface zero value is `nil`, using it without assignment causes runtime panic
- Error examples intentionally show problems that should be avoided in actual development
- It's recommended to understand basic concepts before learning advanced patterns
- Different interface patterns have their own pros/cons, choose based on specific scenarios

### 🤝 Contributing

Welcome to submit Issues and Pull Requests to improve these examples!

---

## Chinese

本目录包含了Go语言接口与多态的完整且实用的示例集合。每个主题都包含真实业务场景、最佳实践和常见陷阱，帮助你深入理解Go的类型系统和接口特性。

### 🧠 覆盖内容与特性
- **接口定义与实现**：基础用法、解耦、插件模式、零值与nil陷阱、编译时错误
- **隐式实现与多态**：多类型实现、业务多态、指针/值接收者、mock测试、nil指针陷阱
- **类型断言与type switch**：安全/不安全断言、type switch日志与业务分发、反例
- **接口组合**：嵌入、ReadWriter、中间件/装饰器、工厂模式、方法冲突错误
- **空接口/any**：通用容器、type switch、JSON解码、类型安全陷阱、反射

### 📁 项目结构

| 文件 | 主题 | 核心内容 |
|------|------|----------|
| `01_interface_basic.go` | 接口基础 | 定义、实现、解耦、插件、nil陷阱、编译错误 |
| `02_implicit_impl.go` | 隐式实现 | 多类型、业务多态、指针/值接收者、mock、nil陷阱 |
| `03_type_assertion.go` | 类型断言 | 安全/不安全断言、type switch、反例、业务分发 |
| `04_interface_composition.go` | 接口组合 | 嵌入、ReadWriter、中间件、工厂、冲突错误 |
| `05_empty_interface.go` | 空接口与any | 通用容器、type switch、JSON解码、类型安全、反射 |

### 🎯 学习路径
- **基础阶段**：接口定义与实现，多态与解耦
- **进阶阶段**：类型断言与type switch，接口组合与中间件
- **高级阶段**：空接口与动态类型、反射

### ✅ 已覆盖特性与场景
- 接口定义、实现、赋值
- 解耦与插件模式
- 接口零值与nil陷阱
- 编译时与运行时错误
- 多类型与业务多态
- 指针与值接收者方法集
- mock测试
- 类型断言（安全/不安全）、type switch、反例
- 接口组合、嵌入、ReadWriter、中间件、工厂
- 空接口/any、通用容器、JSON解码、反射

### 🚩 未覆盖的进阶话题
- 递归/嵌套接口
- 接口动态类型与值（reflect）
- 接口与nil的深层陷阱
- 性能开销
- 接口与泛型约束（Go 1.18+）
- 接口可比性与panic
- 高级反射与动态代理
- 接口序列化/反序列化注意事项
- 标准库接口最佳实践（如io.Reader、error）

如需这些进阶话题的示例或讲解，欢迎提出！

### ⚠️ 注意事项
- 接口零值是`nil`，未赋值直接使用会panic
- 错误示例故意展示问题，实际开发中应避免
- 建议先理解基础概念再学习高级模式
- 不同接口模式各有优缺点，应根据具体场景选择

---

**Tip**: It's recommended to first run `go run .` to view all available examples, then gradually dive deeper following the learning path. 