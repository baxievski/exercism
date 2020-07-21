// Package tree implements a solution for the exercism tree building challenge
package tree

import (
	"fmt"
	"sort"
)

// Node represents a single node in the tree of records
type Node struct {
	ID       int
	Children []*Node
}

// Record represents a forum post
type Record struct {
	ID     int
	Parent int
}

// Build creates a tree of nodes out of a list of records
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Sort(byID(records))
	nodes := map[int]*Node{}

	for i, record := range records {
		if i != record.ID {
			return nil, fmt.Errorf("node is our of order")
		}
		if record.ID < record.Parent {
			return nil, fmt.Errorf("record id %v is higher than its parents id %v", record.ID, record.Parent)
		}
		if record.ID == 0 {
			nodes[0] = &Node{ID: 0}
			continue
		}
		if record.ID == record.Parent {
			return nil, fmt.Errorf("record is it's own parent: %v", record)
		}
		nodes[record.ID] = &Node{ID: record.ID}
		parent := nodes[record.Parent]
		parent.Children = append(parent.Children, nodes[record.ID])
	}
	return nodes[0], nil
}

type byID []Record

func (records byID) Len() int {
	return len(records)
}

func (records byID) Less(i, j int) bool {
	return records[i].ID < records[j].ID
}

func (records byID) Swap(i, j int) {
	records[i], records[j] = records[j], records[i]
}
