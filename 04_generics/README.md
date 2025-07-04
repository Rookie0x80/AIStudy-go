# Go Generics Programming Examples

[English](#english) | [中文](#chinese)

---

## English

This directory contains a comprehensive collection of Go generics programming examples, from basic concepts to advanced patterns, helping you deeply understand Go's generic programming capabilities introduced in Go 1.18+.

### 🚀 Quick Start

```bash
# View all available examples
go run .

# Run specific example (1-8)
go run . <example_number>
```

### 🧬 Core Concepts

#### Basic Concepts
- **Type Parameters**: Generic type variables that allow functions and types to work with multiple types
- **Constraints**: Interfaces that specify what operations are allowed on type parameters
- **Type Inference**: Automatic deduction of type arguments from function arguments
- **Type Sets**: Collections of types that define what types satisfy a constraint
- **Instantiation**: The process of creating specific versions of generic functions or types

#### Constraint System
- **Built-in Constraints**: `any`, `comparable` and other predefined constraints
- **Custom Constraints**: User-defined interfaces for specific type requirements
- **Type Sets**: Using `~` for underlying types and `|` for unions
- **Approximation**: `~T` allows types with underlying type T

#### Generic Types & Methods
- **Generic Functions**: Functions with type parameters
- **Generic Structs**: Data structures that work with multiple types
- **Generic Interfaces**: Interfaces with type parameters
- **Method Sets**: How methods work with generic types and receivers

#### Advanced Features
- **Type Sets & Unions**: Complex constraint definitions with multiple types
- **Type Inference**: Advanced scenarios and limitations
- **Performance**: Compile-time vs runtime considerations
- **Interoperability**: Working with existing non-generic code

### 📁 Project Structure

| File | Topic | Core Content |
|------|-------|--------------|
| `01_basic_generics.go` | Basic Generics | Function generics, type parameters, type inference, basic syntax |
| `02_constraints.go` | Constraint System | Built-in constraints, custom constraints, type sets, comparable |
| `03_generic_types.go` | Generic Types | Generic structs, interfaces, method definitions, receivers |
| `04_advanced_constraints.go` | Advanced Constraints | Type sets, union types, approximation types, complex constraints |
| `05_generic_containers.go` | Generic Containers | Stack, queue, map, tree and other data structure implementations |
| `06_generic_algorithms.go` | Generic Algorithms | Sorting, searching, transformation, aggregation, functional programming |
| `07_design_patterns.go` | Design Patterns | Factory, builder, decorator patterns applied with generics |
| `08_best_practices.go` | Best Practices | Performance optimization, compile-time checks, common pitfall avoidance |

### 🎯 Learning Path

#### Basic Stage (1-3)
Master the fundamental concepts of Go generics:
- **Basic Generics**: Function type parameters and type inference
- **Constraints**: Understanding and creating constraints for type safety
- **Generic Types**: Defining and using generic structs and interfaces

#### Intermediate Stage (4-6)
Learn advanced features and practical applications:
- **Advanced Constraints**: Complex type sets and union types
- **Generic Containers**: Implementing common data structures
- **Generic Algorithms**: Creating reusable algorithmic patterns

#### Advanced Stage (7-8)
Explore design patterns and production considerations:
- **Design Patterns**: Applying classic patterns with generics
- **Best Practices**: Performance, maintainability, and team collaboration

### 💡 Usage Recommendations

#### Running Examples
```bash
# View help information
go run .

# Run examples by number
go run . 1    # Basic generics
go run . 4    # Advanced constraints
go run . 8    # Best practices
```

#### Learning Suggestions
1. **Prerequisites**: Ensure you're using Go 1.18+ with generics support
2. **Step by Step**: Learn in order of file numbers, building foundation before advanced patterns
3. **Hands-on Practice**: Modify type parameters and constraints to see how they affect compilation
4. **Compare Approaches**: Pay attention to generic vs non-generic implementations
5. **Performance Awareness**: Understand compile-time costs and runtime implications

### 🔍 Example Features

Each example file includes:
- ✅ **Basic Usage**: Practical application of core concepts
- ✅ **Progressive Complexity**: From simple to advanced usage patterns
- ✅ **Comparative Examples**: Generic vs non-generic implementations
- ✅ **Error Examples**: Common mistakes and how to avoid them
- ✅ **Real-world Scenarios**: Practical applications in actual development
- ✅ **Performance Notes**: Compilation and runtime considerations
- ✅ **Detailed Comments**: Complete explanation of code logic and design decisions

### ✅ Covered Features & Scenarios

#### Core Generic Features
- Function and type parameter definitions
- Built-in and custom constraint creation
- Type inference in various scenarios
- Generic struct and interface definitions
- Method sets and receiver type handling

#### Advanced Constraint Techniques
- Type sets with `~` approximation
- Union types with `|` operator
- Complex constraint compositions
- Constraint satisfaction checking
- Interface embedding in constraints

#### Practical Applications
- Generic data structures (containers, collections)
- Algorithmic patterns (sorting, searching, transforming)
- Design pattern implementations
- Type-safe API design
- Code reuse and abstraction patterns

### 🚩 What's Not Covered (Advanced Topics)

- Deep compiler internals and type checking algorithms
- Generic methods on interfaces (experimental features)
- Complex type inference edge cases
- Performance benchmarking and profiling
- Integration with reflection and runtime type information
- Advanced constraint solver behavior
- Generic code generation techniques
- Migration strategies for large codebases

If you need examples or explanations for these advanced topics, feel free to request!

### ⚠️ Important Notes

- **Go Version**: Requires Go 1.18 or later for generics support
- **Compilation**: Generic code may increase compilation time
- **Learning Curve**: Start with simple examples before attempting complex constraints
- **Best Practices**: Prefer simple, readable generics over complex type gymnastics
- **Performance**: Understand that generics are compile-time feature with minimal runtime overhead

### 🤝 Contributing

Welcome to submit Issues and Pull Requests to improve these examples!

---

## Chinese

本目录包含了Go语言泛型编程的完整示例集合，从基础概念到高级模式，帮助您深入理解Go 1.18+引入的泛型编程能力。

### 🚀 快速开始

```bash
# 查看所有可用示例
go run .

# 运行特定示例（1-8）
go run . <示例编号>
```

### 🧬 核心概念

#### 基础概念
- **类型参数**: 允许函数和类型使用多种类型的泛型类型变量
- **约束**: 指定类型参数允许哪些操作的接口
- **类型推导**: 从函数参数自动推导类型实参
- **类型集合**: 定义哪些类型满足约束的类型集合
- **实例化**: 创建泛型函数或类型的具体版本的过程

#### 约束系统
- **内置约束**: `any`、`comparable`和其他预定义约束
- **自定义约束**: 用户定义的特定类型需求接口
- **类型集合**: 使用`~`表示底层类型，`|`表示联合
- **近似类型**: `~T`允许底层类型为T的类型

#### 泛型类型与方法
- **泛型函数**: 带有类型参数的函数
- **泛型结构体**: 可以使用多种类型的数据结构
- **泛型接口**: 带有类型参数的接口
- **方法集**: 方法如何与泛型类型和接收器配合使用

#### 高级特性
- **类型集合与联合**: 复杂的多类型约束定义
- **类型推导**: 高级场景和限制
- **性能**: 编译时vs运行时考虑
- **互操作性**: 与现有非泛型代码的配合

### 📁 项目结构

| 文件 | 主题 | 核心内容 |
|------|------|----------|
| `01_basic_generics.go` | 泛型基础 | 函数泛型、类型参数、类型推导、基本语法 |
| `02_constraints.go` | 约束系统 | 内置约束、自定义约束、类型集合、comparable |
| `03_generic_types.go` | 泛型类型 | 泛型结构体、接口、方法定义、接收器 |
| `04_advanced_constraints.go` | 高级约束 | 类型集合、联合类型、近似类型、复杂约束 |
| `05_generic_containers.go` | 泛型容器 | 栈、队列、映射、树等数据结构实现 |
| `06_generic_algorithms.go` | 泛型算法 | 排序、搜索、变换、聚合、函数式编程 |
| `07_design_patterns.go` | 设计模式 | 工厂、建造者、装饰器模式的泛型应用 |
| `08_best_practices.go` | 最佳实践 | 性能优化、编译时检查、常见陷阱避免 |

### 🎯 学习路径

#### 基础阶段 (1-3)
掌握Go泛型的基础概念：
- **泛型基础**: 函数类型参数和类型推导
- **约束系统**: 理解和创建类型安全的约束
- **泛型类型**: 定义和使用泛型结构体和接口

#### 进阶阶段 (4-6)
学习高级特性和实际应用：
- **高级约束**: 复杂类型集合和联合类型
- **泛型容器**: 实现常用数据结构
- **泛型算法**: 创建可重用的算法模式

#### 高级阶段 (7-8)
探索设计模式和生产考虑：
- **设计模式**: 将经典模式与泛型结合
- **最佳实践**: 性能、可维护性和团队协作

### 💡 使用建议

#### 运行示例
```bash
# 查看帮助信息
go run .

# 按编号运行示例
go run . 1    # 泛型基础
go run . 4    # 高级约束
go run . 8    # 最佳实践
```

#### 学习建议
1. **前置条件**: 确保使用Go 1.18+版本以支持泛型
2. **循序渐进**: 按照文件编号顺序学习，打好基础再学高级模式
3. **动手实践**: 修改类型参数和约束，观察对编译的影响
4. **对比学习**: 重点关注泛型与非泛型实现的对比
5. **性能意识**: 理解编译时成本和运行时影响

### 🔍 示例特色

每个示例文件都包含：
- ✅ **基础用法**: 核心概念的实际应用
- ✅ **渐进复杂性**: 从简单到高级的使用模式
- ✅ **对比示例**: 泛型与非泛型实现对比
- ✅ **错误示例**: 常见错误和避免方法
- ✅ **真实场景**: 实际开发中的应用
- ✅ **性能说明**: 编译和运行时考虑
- ✅ **详细注释**: 代码逻辑和设计决策的完整说明

### ✅ 已覆盖特性与场景

#### 核心泛型特性
- 函数和类型参数定义
- 内置和自定义约束创建
- 各种场景下的类型推导
- 泛型结构体和接口定义
- 方法集和接收器类型处理

#### 高级约束技术
- 使用`~`近似的类型集合
- 使用`|`操作符的联合类型
- 复杂约束组合
- 约束满足性检查
- 约束中的接口嵌入

#### 实际应用
- 泛型数据结构（容器、集合）
- 算法模式（排序、搜索、变换）
- 设计模式实现
- 类型安全的API设计
- 代码重用和抽象模式

### 🚩 未覆盖的进阶话题

- 编译器内部和类型检查算法深度解析
- 接口上的泛型方法（实验性特性）
- 复杂类型推导边界情况
- 性能基准测试和分析
- 与反射和运行时类型信息的集成
- 高级约束求解器行为
- 泛型代码生成技术
- 大型代码库的迁移策略

如需这些进阶话题的示例或讲解，欢迎提出！

### ⚠️ 注意事项

- **Go版本**: 需要Go 1.18或更高版本支持泛型
- **编译**: 泛型代码可能增加编译时间
- **学习曲线**: 先从简单示例开始，再尝试复杂约束
- **最佳实践**: 优先使用简单、可读的泛型而非复杂的类型体操
- **性能**: 理解泛型是编译时特性，运行时开销极小

### 🤝 贡献

欢迎提交Issue和Pull Request来改进这些示例！

---

**Tip**: It's recommended to first run `go run .` to view all available examples, then gradually dive deeper following the learning path. 