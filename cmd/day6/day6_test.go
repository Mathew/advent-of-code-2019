package main

import (
	"github.com/mathew/advent-of-code-2019/internal/pkg/asserts"
	"testing"
)

var orbits = []string{
	"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L",
}

func TestNode_GetDepth(t *testing.T) {
	tree := NewTree(orbits)

	asserts.Equals(t, 3, tree.nodes["D"].GetDepth())
}

func TestTree_GetNodesWithParentsCount(t *testing.T) {
	tree := NewTree(orbits)

	asserts.Equals(t, 42, tree.GetNodesWithParentsCount())
}

func TestTree_GetNodesWithParentsCountOutOfOrder(t *testing.T) {
	tree := NewTree([]string{
		"B)C", "COM)B", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L",
	})
	asserts.Equals(t, 42, tree.GetNodesWithParentsCount())
}