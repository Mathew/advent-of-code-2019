package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/files"
	"github.com/mathew/advent-of-code-2019/internal/pkg/structures"
	"log"
	"strings"
)

type Tree struct {
	nodes map[string]*Node
}

func NewTree(links []string) Tree {
	t := Tree{
		nodes: map[string]*Node{},
	}

	for _, l := range links {
		ss := strings.Split(l, ")")
		t.Insert(ss[0], ss[1])
	}

	return t
}

func (t *Tree) Insert(parent, child string) {
	p, ok := t.nodes[parent]
	if !ok {
		t.nodes[parent] = NewNode(parent, nil)
		p = t.nodes[parent]
	}

	c, ok := t.nodes[child]
	if !ok {
		t.nodes[child] = NewNode(child, p)
		c = t.nodes[child]
	} else {
		c.Parent = p
	}

	t.nodes[parent].AddChild(c)
}

func (t Tree) GetNodesWithParentsCount() int {
	c := 0
	for _, n := range t.nodes {
		c += n.GetDepth()
	}

	return c
}

func (t Tree) GetNodesToRootFrom(value string) []*Node {
	if n, ok := t.nodes[value]; ok {
		return n.GetNodesToRoot()
	}

	return []*Node{}
}

type Node struct {
	value    string
	Parent   *Node
	Children []*Node
}

func (t Tree) GetDistanceBetweenNodes(value, value2 string) int {
	ns := t.GetNodesToRootFrom(value)
	ns2 := t.GetNodesToRootFrom(value2)

	var vals []string
	var vals2 []string

	for _, n := range ns {
		vals = append(vals, n.value)
	}

	for _, n := range ns2 {
		vals2 = append(vals2, n.value)
	}

	intr, ok := structures.GetFirstIntersection(vals, vals2)
	if !ok {
		log.Fatalf("No intersection found for %v and %v", vals, vals2)
	}

	return len(t.nodes[value].GetNodesToParent(intr)) + len(t.nodes[value2].GetNodesToParent(intr))
}

func NewNode(value string, parent *Node) *Node {
	root := Node{
		value,
		parent,
		[]*Node{},
	}

	return &root
}

func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

func (n *Node) GetDepth() int {
	return len(n.GetNodesToRoot())
}

func (n *Node) GetNodesToRoot() []*Node {
	p := n.Parent
	var ns []*Node

	for p != nil {
		ns = append(ns, p)
		p = p.Parent
	}

	return ns
}

func (n *Node) GetNodesToParent(val string) []*Node {
	p := n.Parent
	var ns []*Node

	for true {
		if p.Parent == nil {
			return []*Node{}
		}

		ns = append(ns, p)

		if p.Parent.value == val {
			return ns
		}

		p = p.Parent
	}

	return ns
}

func main() {
	orbits := files.Load("cmd/day6/input.txt", "\n")

	t := NewTree(orbits)
	log.Printf("Total orbits: %v", t.GetNodesWithParentsCount())

	// Part 2
	log.Printf("Distance between nodes: %v", t.GetDistanceBetweenNodes("YOU", "SAN"))
}
