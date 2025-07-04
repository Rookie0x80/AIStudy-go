package main

import (
	"fmt"
)

// ==========================================
// Generic Queue Implementation
// ==========================================

// Queue represents a generic first-in-first-out queue
type Queue[T any] struct {
	items []T
}

// NewQueue creates a new queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

// Enqueue adds an item to the back of the queue
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front item
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Front returns the front item without removing it
func (q *Queue[T]) Front() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	return q.items[0], true
}

// Size returns the number of items in the queue
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// IsEmpty checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// ==========================================
// Generic Set Implementation
// ==========================================

// Set represents a generic set
type Set[T comparable] struct {
	items map[T]struct{}
}

// NewSet creates a new set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]struct{})}
}

// Add adds an item to the set
func (s *Set[T]) Add(item T) {
	s.items[item] = struct{}{}
}

// Remove removes an item from the set
func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

// Contains checks if an item is in the set
func (s *Set[T]) Contains(item T) bool {
	_, exists := s.items[item]
	return exists
}

// Size returns the number of items in the set
func (s *Set[T]) Size() int {
	return len(s.items)
}

// ToSlice returns all items as a slice
func (s *Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s.items))
	for item := range s.items {
		result = append(result, item)
	}
	return result
}

// Union returns a new set with items from both sets
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.items {
		result.Add(item)
	}
	for item := range other.items {
		result.Add(item)
	}
	return result
}

// Intersection returns a new set with items common to both sets
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.items {
		if other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

// ==========================================
// Generic Binary Tree
// ==========================================

// TreeNode represents a node in a binary tree
type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

// BinaryTree represents a generic binary tree
type BinaryTree[T any] struct {
	Root *TreeNode[T]
	size int
}

// NewBinaryTree creates a new binary tree
func NewBinaryTree[T any]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

// Insert adds a value to the tree (for demonstration, adds left-to-right)
func (bt *BinaryTree[T]) Insert(value T) {
	newNode := &TreeNode[T]{Value: value}

	if bt.Root == nil {
		bt.Root = newNode
		bt.size++
		return
	}

	// Simple insertion: level-order insertion
	queue := []*TreeNode[T]{bt.Root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Left == nil {
			current.Left = newNode
			bt.size++
			return
		} else if current.Right == nil {
			current.Right = newNode
			bt.size++
			return
		} else {
			queue = append(queue, current.Left, current.Right)
		}
	}
}

// InOrder traverses the tree in-order
func (bt *BinaryTree[T]) InOrder() []T {
	var result []T
	bt.inOrderHelper(bt.Root, &result)
	return result
}

func (bt *BinaryTree[T]) inOrderHelper(node *TreeNode[T], result *[]T) {
	if node != nil {
		bt.inOrderHelper(node.Left, result)
		*result = append(*result, node.Value)
		bt.inOrderHelper(node.Right, result)
	}
}

// PreOrder traverses the tree pre-order
func (bt *BinaryTree[T]) PreOrder() []T {
	var result []T
	bt.preOrderHelper(bt.Root, &result)
	return result
}

func (bt *BinaryTree[T]) preOrderHelper(node *TreeNode[T], result *[]T) {
	if node != nil {
		*result = append(*result, node.Value)
		bt.preOrderHelper(node.Left, result)
		bt.preOrderHelper(node.Right, result)
	}
}

// Size returns the number of nodes in the tree
func (bt *BinaryTree[T]) Size() int {
	return bt.size
}

// ==========================================
// Generic Trie (Prefix Tree)
// ==========================================

// TrieNode represents a node in a trie
type TrieNode[T any] struct {
	Value    T
	IsEnd    bool
	Children map[rune]*TrieNode[T]
}

// Trie represents a generic trie
type Trie[T any] struct {
	Root *TrieNode[T]
}

// NewTrie creates a new trie
func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{
		Root: &TrieNode[T]{Children: make(map[rune]*TrieNode[T])},
	}
}

// Insert adds a word with associated value to the trie
func (t *Trie[T]) Insert(word string, value T) {
	current := t.Root

	for _, char := range word {
		if _, exists := current.Children[char]; !exists {
			current.Children[char] = &TrieNode[T]{
				Children: make(map[rune]*TrieNode[T]),
			}
		}
		current = current.Children[char]
	}

	current.IsEnd = true
	current.Value = value
}

// Search finds a word in the trie
func (t *Trie[T]) Search(word string) (T, bool) {
	current := t.Root

	for _, char := range word {
		if child, exists := current.Children[char]; exists {
			current = child
		} else {
			var zero T
			return zero, false
		}
	}

	if current.IsEnd {
		return current.Value, true
	}

	var zero T
	return zero, false
}

// StartsWith checks if any word starts with the given prefix
func (t *Trie[T]) StartsWith(prefix string) bool {
	current := t.Root

	for _, char := range prefix {
		if child, exists := current.Children[char]; exists {
			current = child
		} else {
			return false
		}
	}

	return true
}

// GetWordsWithPrefix returns all words that start with the given prefix
func (t *Trie[T]) GetWordsWithPrefix(prefix string) []string {
	var result []string
	current := t.Root

	// Navigate to the prefix
	for _, char := range prefix {
		if child, exists := current.Children[char]; exists {
			current = child
		} else {
			return result // No words with this prefix
		}
	}

	// Collect all words from this point
	t.collectWords(current, prefix, &result)
	return result
}

func (t *Trie[T]) collectWords(node *TrieNode[T], currentWord string, result *[]string) {
	if node.IsEnd {
		*result = append(*result, currentWord)
	}

	for char, child := range node.Children {
		t.collectWords(child, currentWord+string(char), result)
	}
}

// ==========================================
// Generic Graph Implementation
// ==========================================

// Graph represents a generic directed graph
type Graph[T comparable] struct {
	vertices map[T][]T
}

// NewGraph creates a new graph
func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{vertices: make(map[T][]T)}
}

// AddVertex adds a vertex to the graph
func (g *Graph[T]) AddVertex(vertex T) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = make([]T, 0)
	}
}

// AddEdge adds a directed edge from source to destination
func (g *Graph[T]) AddEdge(source, destination T) {
	g.AddVertex(source)
	g.AddVertex(destination)
	g.vertices[source] = append(g.vertices[source], destination)
}

// GetNeighbors returns the neighbors of a vertex
func (g *Graph[T]) GetNeighbors(vertex T) []T {
	if neighbors, exists := g.vertices[vertex]; exists {
		return neighbors
	}
	return []T{}
}

// BFS performs breadth-first search from a starting vertex
func (g *Graph[T]) BFS(start T) []T {
	visited := make(map[T]bool)
	queue := []T{start}
	result := []T{}

	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		for _, neighbor := range g.vertices[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

// DFS performs depth-first search from a starting vertex
func (g *Graph[T]) DFS(start T) []T {
	visited := make(map[T]bool)
	result := []T{}
	g.dfsHelper(start, visited, &result)
	return result
}

func (g *Graph[T]) dfsHelper(vertex T, visited map[T]bool, result *[]T) {
	visited[vertex] = true
	*result = append(*result, vertex)

	for _, neighbor := range g.vertices[vertex] {
		if !visited[neighbor] {
			g.dfsHelper(neighbor, visited, result)
		}
	}
}

// ==========================================
// Main Example Function
// ==========================================

func runGenericContainersExample() {
	fmt.Println("\n🔸 Generic Queue")

	// String queue
	stringQueue := NewQueue[string]()
	stringQueue.Enqueue("first")
	stringQueue.Enqueue("second")
	stringQueue.Enqueue("third")

	fmt.Printf("Queue size: %d\n", stringQueue.Size())

	if front, ok := stringQueue.Front(); ok {
		fmt.Printf("Front item: %s\n", front)
	}

	for !stringQueue.IsEmpty() {
		if item, ok := stringQueue.Dequeue(); ok {
			fmt.Printf("Dequeued: %s\n", item)
		}
	}

	fmt.Println("\n🔸 Generic Set")

	// Integer set operations
	set1 := NewSet[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)
	set1.Add(4)

	set2 := NewSet[int]()
	set2.Add(3)
	set2.Add(4)
	set2.Add(5)
	set2.Add(6)

	fmt.Printf("Set1: %v\n", set1.ToSlice())
	fmt.Printf("Set2: %v\n", set2.ToSlice())
	fmt.Printf("Set1 contains 3: %t\n", set1.Contains(3))
	fmt.Printf("Set1 contains 7: %t\n", set1.Contains(7))

	union := set1.Union(set2)
	fmt.Printf("Union: %v\n", union.ToSlice())

	intersection := set1.Intersection(set2)
	fmt.Printf("Intersection: %v\n", intersection.ToSlice())

	fmt.Println("\n🔸 Generic Binary Tree")

	// String binary tree
	tree := NewBinaryTree[string]()
	words := []string{"root", "left", "right", "leftleft", "leftright", "rightleft", "rightright"}

	for _, word := range words {
		tree.Insert(word)
	}

	fmt.Printf("Tree size: %d\n", tree.Size())
	fmt.Printf("In-order traversal: %v\n", tree.InOrder())
	fmt.Printf("Pre-order traversal: %v\n", tree.PreOrder())

	fmt.Println("\n🔸 Generic Trie (Prefix Tree)")

	// String-to-int mapping trie
	trie := NewTrie[int]()

	// Insert some words with associated values
	words_values := map[string]int{
		"car":     1,
		"card":    2,
		"care":    3,
		"careful": 4,
		"cars":    5,
		"cat":     6,
		"cats":    7,
	}

	for word, value := range words_values {
		trie.Insert(word, value)
	}

	// Search for words
	if value, found := trie.Search("car"); found {
		fmt.Printf("Found 'car' with value: %d\n", value)
	}

	if value, found := trie.Search("care"); found {
		fmt.Printf("Found 'care' with value: %d\n", value)
	}

	if _, found := trie.Search("careless"); !found {
		fmt.Println("'careless' not found")
	}

	// Check prefixes
	fmt.Printf("Words starting with 'car': %t\n", trie.StartsWith("car"))
	fmt.Printf("Words starting with 'dog': %t\n", trie.StartsWith("dog"))

	// Get all words with prefix
	carWords := trie.GetWordsWithPrefix("car")
	fmt.Printf("All words starting with 'car': %v\n", carWords)

	catWords := trie.GetWordsWithPrefix("cat")
	fmt.Printf("All words starting with 'cat': %v\n", catWords)

	fmt.Println("\n🔸 Generic Graph")

	// String graph
	graph := NewGraph[string]()

	// Add edges to create a simple graph
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "D")
	graph.AddEdge("D", "E")
	graph.AddEdge("B", "E")

	fmt.Printf("Neighbors of A: %v\n", graph.GetNeighbors("A"))
	fmt.Printf("Neighbors of B: %v\n", graph.GetNeighbors("B"))

	// Traverse the graph
	bfsResult := graph.BFS("A")
	fmt.Printf("BFS from A: %v\n", bfsResult)

	dfsResult := graph.DFS("A")
	fmt.Printf("DFS from A: %v\n", dfsResult)

	// Integer graph for numbers
	numGraph := NewGraph[int]()
	numGraph.AddEdge(1, 2)
	numGraph.AddEdge(1, 3)
	numGraph.AddEdge(2, 4)
	numGraph.AddEdge(3, 4)
	numGraph.AddEdge(4, 5)

	fmt.Printf("Number graph BFS from 1: %v\n", numGraph.BFS(1))
	fmt.Printf("Number graph DFS from 1: %v\n", numGraph.DFS(1))

	fmt.Println("\n✅ Generic containers examples completed!")
}
