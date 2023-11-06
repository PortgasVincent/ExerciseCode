package main

import "fmt"

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.M, obj.Head, obj.Tail)
	fmt.Println(obj.Get(1))
}

type Node struct {
	Key   int
	Value int
	Left  *Node
	Right *Node
}

type LRUCache struct {
	M      map[int]*Node
	Head   *Node
	Tail   *Node
	Length int
	Max    int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		M:      map[int]*Node{},
		Head:   nil,
		Tail:   nil,
		Length: 0,
		Max:    capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.M[key]
	if !exist {
		return -1
	}
	this.Reset(node)
	return node.Value
}

func (this *LRUCache) Reset(node *Node) {
	if this.Head == node {
		return
	}
	if this.Tail == node {
		this.Tail = node.Left
	} else {
		node.Right.Left = node.Left
	}

	node.Left.Right = node.Right
	node.Left = nil
	node.Right = this.Head
	this.Head.Left = node
	this.Head = node
	return
}

func (this *LRUCache) Put(key int, value int) {
	node, exist := this.M[key]
	if exist {
		node.Value = value
		this.Reset(node)
		return
	}

	new := &Node{
		Key:   key,
		Value: value,
	}
	if this.Length == this.Max {
		expired := this.Tail

		if this.Max == 1 {
			this.Head = new
			this.Tail = new
			delete(this.M, expired.Key)
			this.M[key] = new
			return
		}

		this.Tail = expired.Left
		this.Tail.Right = nil
		delete(this.M, expired.Key)
		this.Length--
	}
	if this.Length == 0 {
		this.Head = new
		this.Tail = new
		this.M[key] = new
		this.Length++
		return
	}
	new.Right = this.Head
	this.Head.Left = new
	this.Head = new
	this.M[key] = new
	this.Length++
	return
}
