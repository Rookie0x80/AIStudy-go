# Go Concurrency Programming Examples

[English](#english) | [中文](#chinese)

---

## English

This directory contains a comprehensive collection of Go concurrency programming examples, from basic concepts to advanced patterns, helping you deeply understand Go's concurrency features.

### 🚀 Quick Start

```bash
# View all available examples
go run .

# Run specific example (1-10)
go run . <example_number>
```

### 🧠 Core Concepts

#### Basic Concepts
- **Goroutine**: Go's lightweight threads, started with the `go` keyword, more lightweight and efficient than traditional threads
- **Channel**: Type-safe communication pipes for data transmission and synchronization between goroutines
- **Select**: Multiplexing statement that can monitor multiple channel operations simultaneously
- **Context**: Request context for controlling goroutine lifecycle, timeouts, and cancellation

#### Synchronization Primitives
- **Mutex**: Mutual exclusion lock ensuring only one goroutine accesses shared resources at a time
- **RWMutex**: Read-write lock allowing multiple reads or one write operation simultaneously
- **WaitGroup**: Wait group for waiting for a group of goroutines to complete
- **Once**: Ensures a function executes only once
- **Cond**: Condition variables for conditional waiting and notification between goroutines

#### Atomic Operations
- **Atomic Operations**: Thread-safe operations without locks, higher performance
- **CAS**: Compare-and-swap operations, foundation of lock-free programming
- **Atomic Values**: Can atomically store and load values of any type

#### Concurrency Patterns
- **Actor Model**: Message-passing based concurrency model where each Actor independently processes messages
- **CSP Pattern**: Communicating Sequential Processes, inter-process communication through channels
- **Future/Promise**: Asynchronous programming pattern for handling asynchronous operation results
- **Reactive Programming**: Reactive programming for handling data streams and event streams

### 📁 Project Structure

| File | Topic | Core Content |
|------|-------|--------------|
| `01_goroutines.go` | Goroutines Basics | Lightweight thread creation, closure traps, worker pool patterns |
| `02_channels.go` | Channel Communication | Type-safe communication, buffering, directionality, timeout handling |
| `03_select.go` | Select Multiplexing | Multi-channel monitoring, non-blocking operations, random selection |
| `04_context.go` | Context Management | Request lifecycle management, timeout cancellation, value passing |
| `05_sync.go` | Sync Package Primitives | Mutex, RWMutex, WaitGroup, Once, Cond |
| `06_atomic.go` | Atomic Operations | Lock-free programming, CAS operations, atomic value operations |
| `07_actor.go` | Actor Model | Message-based concurrency, Actor communication, supervision patterns |
| `08_csp.go` | CSP Pattern | Communicating Sequential Processes, pipeline processing, publish-subscribe |
| `09_future.go` | Future/Promise | Asynchronous result handling, composition operations, error handling |
| `10_reactive.go` | Reactive Programming | Reactive data streams, event handling, composition operations |

### 🎯 Learning Path

#### Basic Stage (1-4)
Master the core concepts of Go concurrency programming:
- **Goroutines**: Creation and management of lightweight threads
- **Channels**: Type-safe communication mechanisms
- **Select**: Multiplexing and timeout control
- **Context**: Request context and lifecycle management

#### Intermediate Stage (5-6)
Learn synchronization and atomic operations:
- **Sync Package**: Mutex, read-write locks, wait groups, and other synchronization primitives
- **Atomic Operations**: Lock-free programming and performance optimization

#### Advanced Stage (7-10)
Explore concurrency design patterns:
- **Actor Model**: Message-based concurrency architecture
- **CSP Pattern**: Communicating Sequential Processes pattern
- **Future/Promise**: Asynchronous programming patterns
- **Reactive Programming**: Reactive data stream processing

### 💡 Usage Recommendations

#### Running Examples
```bash
# View help information
go run .

# Run examples by number
go run . 1    # Goroutines basics
go run . 5    # Sync package primitives
go run . 10   # Reactive programming
```

#### Learning Suggestions
1. **Step by Step**: Learn in order of file numbers, build a solid foundation before learning advanced patterns
2. **Hands-on Practice**: Modify example parameters and observe different results
3. **Comparative Learning**: Focus on comparing correct examples with error examples
4. **Understand Principles**: Deeply understand the applicable scenarios and pros/cons of each pattern

### 🔍 Example Features

Each example file includes:
- ✅ **Basic Usage**: Practical application of core concepts
- ✅ **Advanced Techniques**: Advanced usage and best practices
- ✅ **Error Examples**: Common pitfalls and avoidance methods
- ✅ **Comprehensive Cases**: Typical scenarios in actual development
- ✅ **Detailed Comments**: Complete explanation of code logic

### ⚠️ Important Notes

- Time delays in examples are only for demonstration purposes
- Error examples intentionally show problems that should be avoided in actual development
- It's recommended to understand basic concepts before learning advanced patterns
- Different concurrency patterns have their own pros/cons, choose based on specific scenarios

### 🤝 Contributing

Welcome to submit Issues and Pull Requests to improve these examples!

---

## Chinese

本目录包含了Go语言并发编程的完整示例集合，从基础概念到高级模式，帮助您深入理解Go的并发特性。

### 🚀 快速开始

```bash
# 查看所有可用示例
go run .

# 运行特定示例（1-10）
go run . <示例编号>
```

### 🧠 核心概念

#### 基础概念
- **Goroutine**: Go的轻量级线程，使用`go`关键字启动，比传统线程更轻量、更高效
- **Channel**: 类型安全的通信管道，用于goroutine之间的数据传输和同步
- **Select**: 多路复用语句，可以同时监听多个channel的操作
- **Context**: 请求上下文，用于控制goroutine的生命周期、超时和取消

#### 同步原语
- **Mutex**: 互斥锁，确保同一时间只有一个goroutine访问共享资源
- **RWMutex**: 读写锁，允许多个读操作或一个写操作同时进行
- **WaitGroup**: 等待组，用于等待一组goroutine完成
- **Once**: 确保某个函数只执行一次
- **Cond**: 条件变量，用于goroutine间的条件等待和通知

#### 原子操作
- **原子操作**: 无需加锁的线程安全操作，性能更高
- **CAS**: 比较并交换操作，无锁编程的基础
- **原子值**: 可以原子地存储和加载任意类型的值

#### 并发模式
- **Actor模型**: 基于消息传递的并发模型，每个Actor独立处理消息
- **CSP模式**: 通信顺序进程，通过channel进行进程间通信
- **Future/Promise**: 异步编程模式，处理异步操作的结果
- **Reactive编程**: 响应式编程，处理数据流和事件流

### 📁 项目结构

| 文件 | 主题 | 核心内容 |
|------|------|----------|
| `01_goroutines.go` | Goroutines基础 | 轻量级线程创建、闭包陷阱、工作池模式 |
| `02_channels.go` | Channels通信 | 类型安全通信、缓冲、方向性、超时处理 |
| `03_select.go` | Select多路复用 | 多channel监听、非阻塞操作、随机选择 |
| `04_context.go` | Context上下文 | 请求生命周期管理、超时取消、值传递 |
| `05_sync.go` | Sync包原语 | Mutex、RWMutex、WaitGroup、Once、Cond |
| `06_atomic.go` | 原子操作 | 无锁编程、CAS操作、原子值操作 |
| `07_actor.go` | Actor模型 | 基于消息的并发、Actor通信、监督模式 |
| `08_csp.go` | CSP模式 | 通信顺序进程、管道处理、发布订阅 |
| `09_future.go` | Future/Promise | 异步结果处理、组合操作、错误处理 |
| `10_reactive.go` | Reactive编程 | 响应式数据流、事件处理、组合操作 |

### 🎯 学习路径

#### 基础阶段 (1-4)
掌握Go并发编程的核心概念：
- **Goroutines**: 轻量级线程的创建和管理
- **Channels**: 类型安全的通信机制
- **Select**: 多路复用和超时控制
- **Context**: 请求上下文和生命周期管理

#### 进阶阶段 (5-6)
学习同步和原子操作：
- **Sync包**: 互斥锁、读写锁、等待组等同步原语
- **原子操作**: 无锁编程和性能优化

#### 高级阶段 (7-10)
探索并发设计模式：
- **Actor模型**: 基于消息的并发架构
- **CSP模式**: 通信顺序进程模式
- **Future/Promise**: 异步编程模式
- **Reactive编程**: 响应式数据流处理

### 💡 使用建议

#### 运行示例
```bash
# 查看帮助信息
go run .

# 按编号运行示例
go run . 1    # Goroutines基础
go run . 5    # Sync包原语
go run . 10   # Reactive编程
```

#### 学习建议
1. **循序渐进**: 按照文件编号顺序学习，打好基础再学高级模式
2. **动手实践**: 修改示例参数，观察不同结果
3. **对比学习**: 重点关注正确示例与错误示例的对比
4. **理解原理**: 深入理解每个模式的适用场景和优缺点

### 🔍 示例特色

每个示例文件都包含：
- ✅ **基础用法**: 核心概念的实际应用
- ✅ **进阶技巧**: 高级用法和最佳实践
- ✅ **错误示例**: 常见陷阱和避免方法
- ✅ **综合案例**: 实际开发中的典型场景
- ✅ **详细注释**: 代码逻辑的完整说明

### ⚠️ 注意事项

- 示例中的时间延迟仅用于演示效果
- 错误示例故意展示问题，实际开发中应避免
- 建议在理解基础概念后再学习高级模式
- 不同并发模式各有优缺点，应根据具体场景选择

### 🤝 贡献

欢迎提交Issue和Pull Request来改进这些示例！

---

**Tip**: It's recommended to first run `go run .` to view all available examples, then gradually dive deeper following the learning path. 