package main

import (
	"fmt"
	"sync"
	"time"
)

// Actor model examples

// Message defines message interface
type Message interface {
	Type() string
}

// StringMessage string message
type StringMessage struct {
	Content string
}

func (m StringMessage) Type() string {
	return "string"
}

// NumberMessage number message
type NumberMessage struct {
	Value int
}

func (m NumberMessage) Type() string {
	return "number"
}

// StopMessage stop message
type StopMessage struct{}

func (m StopMessage) Type() string {
	return "stop"
}

// PingMessage Ping message
type PingMessage struct {
	Count int
}

func (m PingMessage) Type() string {
	return "ping"
}

// PongMessage Pong message
type PongMessage struct {
	Count int
}

func (m PongMessage) Type() string {
	return "pong"
}

// WorkMessage work message
type WorkMessage struct {
	ID   int
	Task string
}

func (m WorkMessage) Type() string {
	return "work"
}

// ResultMessage result message
type ResultMessage struct {
	ID     int
	Result string
}

func (m ResultMessage) Type() string {
	return "result"
}

// ErrorMessage error message
type ErrorMessage struct {
	ActorID string
	Error   string
}

func (m ErrorMessage) Type() string {
	return "error"
}

// SetMessage set message
type SetMessage struct {
	Key   string
	Value interface{}
}

func (m SetMessage) Type() string {
	return "set"
}

// GetMessage get message
type GetMessage struct {
	Key      string
	Response chan interface{}
}

func (m GetMessage) Type() string {
	return "get"
}

// RouteMessage route message
type RouteMessage struct {
	Target string
	Data   interface{}
}

func (m RouteMessage) Type() string {
	return "route"
}

// PipelineMessage pipeline message
type PipelineMessage struct {
	Data  string
	Stage int
}

func (m PipelineMessage) Type() string {
	return "pipeline"
}

// BankMessage bank message
type BankMessage struct {
	Operation string
	AccountID string
	Amount    int
	Response  chan interface{}
}

func (m BankMessage) Type() string {
	return "bank"
}

// Actor Actor struct
type Actor struct {
	ID       string
	mailbox  chan Message
	handlers map[string]func(Message)
	wg       sync.WaitGroup
	stop     chan struct{}
}

// NewActor create new Actor
func NewActor(id string) *Actor {
	return &Actor{
		ID:       id,
		mailbox:  make(chan Message, 100),
		handlers: make(map[string]func(Message)),
		stop:     make(chan struct{}),
	}
}

// RegisterHandler register message handler
func (a *Actor) RegisterHandler(msgType string, handler func(Message)) {
	a.handlers[msgType] = handler
}

// Send send message to Actor
func (a *Actor) Send(msg Message) {
	a.mailbox <- msg
}

// Start start Actor
func (a *Actor) Start() {
	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		for {
			select {
			case msg := <-a.mailbox:
				if handler, exists := a.handlers[msg.Type()]; exists {
					handler(msg)
				} else {
					fmt.Printf("Actor %s: unknown message type %s\n", a.ID, msg.Type())
				}
			case <-a.stop:
				fmt.Printf("Actor %s: stopping\n", a.ID)
				return
			}
		}
	}()
}

// Stop stop Actor
func (a *Actor) Stop() {
	close(a.stop)
	a.wg.Wait()
}

// ActorExamples runs all Actor model examples
func ActorExamples() {
	fmt.Println("=== Actor Model Examples ===")

	// Example 1: Basic Actor
	basicActorExample()

	// Example 2: Actor communication
	actorCommunicationExample()

	// Example 3: Actor pool
	actorPoolExample()

	// Example 4: Supervisor Actor
	supervisorActorExample()

	// Example 5: Stateful Actor
	statefulActorExample()

	// Example 6: Router Actor
	routerActorExample()

	// Example 7: Pipeline Actor
	pipelineActorExample()

	// Example 8: Comprehensive example
	comprehensiveActorExample()
}

// Example 1: Basic Actor
func basicActorExample() {
	fmt.Println("\n--- Example 1: Basic Actor ---")

	// Create Actor
	actor := NewActor("printer")

	// Register message handlers
	actor.RegisterHandler("string", func(msg Message) {
		if strMsg, ok := msg.(StringMessage); ok {
			fmt.Printf("Actor %s: received string message '%s'\n", actor.ID, strMsg.Content)
		}
	})

	actor.RegisterHandler("number", func(msg Message) {
		if numMsg, ok := msg.(NumberMessage); ok {
			fmt.Printf("Actor %s: received number message %d\n", actor.ID, numMsg.Value)
		}
	})

	actor.RegisterHandler("stop", func(msg Message) {
		fmt.Printf("Actor %s: received stop message\n", actor.ID)
		actor.Stop()
	})

	// Start Actor
	actor.Start()

	// Send messages
	actor.Send(StringMessage{Content: "Hello, Actor!"})
	actor.Send(NumberMessage{Value: 42})
	actor.Send(StringMessage{Content: "Another message"})

	// Wait for a while then stop
	time.Sleep(100 * time.Millisecond)
	actor.Send(StopMessage{})
	actor.Stop()
}

// Example 2: Actor communication
func actorCommunicationExample() {
	fmt.Println("\n--- Example 2: Actor communication ---")

	// Create two Actors
	pingActor := NewActor("ping")
	pongActor := NewActor("pong")

	// Register Ping Actor handler
	pingActor.RegisterHandler("pong", func(msg Message) {
		if pongMsg, ok := msg.(PongMessage); ok {
			fmt.Printf("Ping: received Pong, count=%d\n", pongMsg.Count)
			if pongMsg.Count < 5 {
				time.Sleep(100 * time.Millisecond)
				pongActor.Send(PingMessage{Count: pongMsg.Count + 1})
			} else {
				fmt.Println("Ping: game over")
				pongActor.Send(StopMessage{})
				pingActor.Send(StopMessage{})
			}
		}
	})

	// Register Pong Actor handler
	pongActor.RegisterHandler("ping", func(msg Message) {
		if pingMsg, ok := msg.(PingMessage); ok {
			fmt.Printf("Pong: received Ping, count=%d\n", pingMsg.Count)
			time.Sleep(100 * time.Millisecond)
			pingActor.Send(PongMessage{Count: pingMsg.Count})
		}
	})

	// Start Actors
	pingActor.Start()
	pongActor.Start()

	// Start game
	pingActor.Send(PingMessage{Count: 1})

	// Wait for game to end
	time.Sleep(1 * time.Second)
	pingActor.Stop()
	pongActor.Stop()
}

// Example 3: Actor pool
func actorPoolExample() {
	fmt.Println("\n--- Example 3: Actor pool ---")

	// Create Actor pool
	const poolSize = 3
	actors := make([]*Actor, poolSize)

	// Create result collector Actor
	resultCollector := NewActor("collector")
	resultCollector.RegisterHandler("result", func(msg Message) {
		if resultMsg, ok := msg.(ResultMessage); ok {
			fmt.Printf("Collector: received result ID=%d, Result=%s\n", resultMsg.ID, resultMsg.Result)
		}
	})
	resultCollector.Start()

	// Create worker Actors
	for i := 0; i < poolSize; i++ {
		actor := NewActor(fmt.Sprintf("worker-%d", i))
		actor.RegisterHandler("work", func(msg Message) {
			if workMsg, ok := msg.(WorkMessage); ok {
				fmt.Printf("Worker %s: processing task ID=%d, Task=%s\n", actor.ID, workMsg.ID, workMsg.Task)
				time.Sleep(200 * time.Millisecond) // Simulate work
				result := fmt.Sprintf("Completed task %d", workMsg.ID)
				resultCollector.Send(ResultMessage{ID: workMsg.ID, Result: result})
			}
		})
		actor.Start()
		actors[i] = actor
	}

	// Send tasks
	for i := 1; i <= 10; i++ {
		workerIndex := (i - 1) % poolSize
		actors[workerIndex].Send(WorkMessage{ID: i, Task: fmt.Sprintf("Task%d", i)})
	}

	// Wait for all tasks to complete
	time.Sleep(2 * time.Second)

	// Stop all Actors
	for _, actor := range actors {
		actor.Send(StopMessage{})
		actor.Stop()
	}
	resultCollector.Send(StopMessage{})
	resultCollector.Stop()
}

// Example 4: Supervisor Actor
func supervisorActorExample() {
	fmt.Println("\n--- Example 4: Supervisor Actor ---")

	// Create supervisor Actor
	supervisor := NewActor("supervisor")
	supervisor.RegisterHandler("error", func(msg Message) {
		if errMsg, ok := msg.(ErrorMessage); ok {
			fmt.Printf("Supervisor: received error report Actor=%s, Error=%s\n", errMsg.ActorID, errMsg.Error)
			fmt.Printf("Supervisor: restarting Actor %s\n", errMsg.ActorID)
		}
	})

	// Create worker Actor
	worker := NewActor("worker")
	worker.RegisterHandler("work", func(msg Message) {
		if workMsg, ok := msg.(WorkMessage); ok {
			fmt.Printf("Worker: processing task %d\n", workMsg.ID)
			// Simulate occasional errors
			if workMsg.ID%3 == 0 {
				fmt.Printf("Worker: task %d failed\n", workMsg.ID)
				supervisor.Send(ErrorMessage{ActorID: worker.ID, Error: "Task processing failed"})
			}
		}
	})

	supervisor.Start()
	worker.Start()

	// Send tasks
	for i := 1; i <= 6; i++ {
		worker.Send(WorkMessage{ID: i, Task: fmt.Sprintf("Task%d", i)})
		time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(500 * time.Millisecond)
	supervisor.Send(StopMessage{})
	worker.Send(StopMessage{})
	supervisor.Stop()
	worker.Stop()
}

// Example 5: Stateful Actor
func statefulActorExample() {
	fmt.Println("\n--- Example 5: Stateful Actor ---")

	// Define stateful Actor
	type StatefulActor struct {
		*Actor
		state map[string]interface{}
		mu    sync.RWMutex
	}

	// Create stateful Actor
	stateActor := &StatefulActor{
		Actor: NewActor("stateful"),
		state: make(map[string]interface{}),
	}

	// Register handlers
	stateActor.RegisterHandler("set", func(msg Message) {
		if setMsg, ok := msg.(SetMessage); ok {
			stateActor.mu.Lock()
			stateActor.state[setMsg.Key] = setMsg.Value
			stateActor.mu.Unlock()
			fmt.Printf("StatefulActor: set %s = %v\n", setMsg.Key, setMsg.Value)
		}
	})

	stateActor.RegisterHandler("get", func(msg Message) {
		if getMsg, ok := msg.(GetMessage); ok {
			stateActor.mu.RLock()
			value, exists := stateActor.state[getMsg.Key]
			stateActor.mu.RUnlock()
			if exists {
				getMsg.Response <- value
			} else {
				getMsg.Response <- nil
			}
		}
	})

	stateActor.Start()

	// Set state
	stateActor.Send(SetMessage{Key: "name", Value: "Alice"})
	stateActor.Send(SetMessage{Key: "age", Value: 30})
	stateActor.Send(SetMessage{Key: "city", Value: "Beijing"})

	// Get state
	responseChan := make(chan interface{}, 1)
	stateActor.Send(GetMessage{Key: "name", Response: responseChan})
	value := <-responseChan
	fmt.Printf("Get state: name = %v\n", value)

	stateActor.Send(GetMessage{Key: "age", Response: responseChan})
	value = <-responseChan
	fmt.Printf("Get state: age = %v\n", value)

	time.Sleep(100 * time.Millisecond)
	stateActor.Send(StopMessage{})
	stateActor.Stop()
}

// Example 6: Router Actor
func routerActorExample() {
	fmt.Println("\n--- Example 6: Router Actor ---")

	// Create router Actor
	router := NewActor("router")
	targets := make(map[string]*Actor)

	// Create target Actors
	for i := 1; i <= 3; i++ {
		targetID := fmt.Sprintf("target-%d", i)
		target := NewActor(targetID)
		target.RegisterHandler("data", func(msg Message) {
			if dataMsg, ok := msg.(StringMessage); ok {
				fmt.Printf("%s: received data '%s'\n", target.ID, dataMsg.Content)
			}
		})
		target.Start()
		targets[targetID] = target
	}

	// Register router handler
	router.RegisterHandler("route", func(msg Message) {
		if routeMsg, ok := msg.(RouteMessage); ok {
			if target, exists := targets[routeMsg.Target]; exists {
				fmt.Printf("Router: routing message to %s\n", routeMsg.Target)
				target.Send(StringMessage{Content: fmt.Sprintf("%v", routeMsg.Data)})
			} else {
				fmt.Printf("Router: target %s does not exist\n", routeMsg.Target)
			}
		}
	})

	router.Start()

	// Send route messages
	router.Send(RouteMessage{Target: "target-1", Data: "Data1"})
	router.Send(RouteMessage{Target: "target-2", Data: "Data2"})
	router.Send(RouteMessage{Target: "target-3", Data: "Data3"})
	router.Send(RouteMessage{Target: "target-1", Data: "More data"})

	time.Sleep(200 * time.Millisecond)

	// Stop all Actors
	router.Send(StopMessage{})
	router.Stop()
	for _, target := range targets {
		target.Send(StopMessage{})
		target.Stop()
	}
}

// Example 7: Pipeline Actor
func pipelineActorExample() {
	fmt.Println("\n--- Example 7: Pipeline Actor ---")

	// Create pipeline Actors
	stage1 := NewActor("stage1")
	stage2 := NewActor("stage2")
	stage3 := NewActor("stage3")

	// Register handlers
	stage1.RegisterHandler("pipeline", func(msg Message) {
		if pipeMsg, ok := msg.(PipelineMessage); ok {
			processed := fmt.Sprintf("Stage1 processed: %s", pipeMsg.Data)
			fmt.Printf("Stage1: %s\n", processed)
			stage2.Send(PipelineMessage{Data: processed, Stage: 1})
		}
	})

	stage2.RegisterHandler("pipeline", func(msg Message) {
		if pipeMsg, ok := msg.(PipelineMessage); ok {
			processed := fmt.Sprintf("Stage2 processed: %s", pipeMsg.Data)
			fmt.Printf("Stage2: %s\n", processed)
			stage3.Send(PipelineMessage{Data: processed, Stage: 2})
		}
	})

	stage3.RegisterHandler("pipeline", func(msg Message) {
		if pipeMsg, ok := msg.(PipelineMessage); ok {
			processed := fmt.Sprintf("Stage3 processed: %s", pipeMsg.Data)
			fmt.Printf("Stage3: %s (final result)\n", processed)
		}
	})

	// Start pipeline
	stage1.Start()
	stage2.Start()
	stage3.Start()

	// Send data to pipeline
	data := []string{"Data1", "Data2", "Data3", "Data4"}
	for _, item := range data {
		stage1.Send(PipelineMessage{Data: item, Stage: 0})
		time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(500 * time.Millisecond)

	// Stop pipeline
	stage1.Send(StopMessage{})
	stage2.Send(StopMessage{})
	stage3.Send(StopMessage{})
	stage1.Stop()
	stage2.Stop()
	stage3.Stop()
}

// Example 8: Comprehensive example
func comprehensiveActorExample() {
	fmt.Println("\n--- Example 8: Comprehensive example ---")

	// Simulate a simple banking system
	type Account struct {
		ID      string
		Balance int
	}

	// Create bank Actor
	bank := NewActor("bank")
	accounts := make(map[string]*Account)
	var mu sync.RWMutex

	bank.RegisterHandler("bank", func(msg Message) {
		if bankMsg, ok := msg.(BankMessage); ok {
			mu.Lock()
			defer mu.Unlock()

			switch bankMsg.Operation {
			case "create":
				accounts[bankMsg.AccountID] = &Account{ID: bankMsg.AccountID, Balance: 0}
				fmt.Printf("Bank: created account %s\n", bankMsg.AccountID)
				bankMsg.Response <- "Account created successfully"

			case "deposit":
				if account, exists := accounts[bankMsg.AccountID]; exists {
					account.Balance += bankMsg.Amount
					fmt.Printf("Bank: account %s deposited %d, balance %d\n", bankMsg.AccountID, bankMsg.Amount, account.Balance)
					bankMsg.Response <- account.Balance
				} else {
					bankMsg.Response <- "Account does not exist"
				}

			case "withdraw":
				if account, exists := accounts[bankMsg.AccountID]; exists {
					if account.Balance >= bankMsg.Amount {
						account.Balance -= bankMsg.Amount
						fmt.Printf("Bank: account %s withdrew %d, balance %d\n", bankMsg.AccountID, bankMsg.Amount, account.Balance)
						bankMsg.Response <- account.Balance
					} else {
						bankMsg.Response <- "Insufficient balance"
					}
				} else {
					bankMsg.Response <- "Account does not exist"
				}

			case "balance":
				if account, exists := accounts[bankMsg.AccountID]; exists {
					fmt.Printf("Bank: account %s balance inquiry %d\n", bankMsg.AccountID, account.Balance)
					bankMsg.Response <- account.Balance
				} else {
					bankMsg.Response <- "Account does not exist"
				}
			}
		}
	})

	bank.Start()

	// Create client Actors
	client1 := NewActor("client1")
	client2 := NewActor("client2")

	// Client handlers
	clientHandler := func(clientID string) func(Message) {
		return func(msg Message) {
			if bankMsg, ok := msg.(BankMessage); ok {
				response := <-bankMsg.Response
				fmt.Printf("%s: received response %v\n", clientID, response)
			}
		}
	}

	client1.RegisterHandler("bank", clientHandler("Client1"))
	client2.RegisterHandler("bank", clientHandler("Client2"))

	client1.Start()
	client2.Start()

	// Simulate banking operations
	responseChan := make(chan interface{}, 1)

	// Client1 creates account and deposits
	client1.Send(BankMessage{Operation: "create", AccountID: "ACC001", Response: responseChan})
	time.Sleep(50 * time.Millisecond)
	client1.Send(BankMessage{Operation: "deposit", AccountID: "ACC001", Amount: 1000, Response: responseChan})
	time.Sleep(50 * time.Millisecond)

	// Client2 creates account and deposits
	client2.Send(BankMessage{Operation: "create", AccountID: "ACC002", Response: responseChan})
	time.Sleep(50 * time.Millisecond)
	client2.Send(BankMessage{Operation: "deposit", AccountID: "ACC002", Amount: 500, Response: responseChan})
	time.Sleep(50 * time.Millisecond)

	// Query balances
	client1.Send(BankMessage{Operation: "balance", AccountID: "ACC001", Response: responseChan})
	client2.Send(BankMessage{Operation: "balance", AccountID: "ACC002", Response: responseChan})
	time.Sleep(50 * time.Millisecond)

	// Withdraw
	client1.Send(BankMessage{Operation: "withdraw", AccountID: "ACC001", Amount: 300, Response: responseChan})
	client2.Send(BankMessage{Operation: "withdraw", AccountID: "ACC002", Amount: 200, Response: responseChan})
	time.Sleep(50 * time.Millisecond)

	time.Sleep(200 * time.Millisecond)

	// Stop all Actors
	bank.Send(StopMessage{})
	client1.Send(StopMessage{})
	client2.Send(StopMessage{})
	bank.Stop()
	client1.Stop()
	client2.Stop()
}
