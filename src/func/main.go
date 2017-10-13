// func project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	// prereqs记录了每个课程的前置课程
	var prereqs = map[string][]string{
		"algorithms":            {"data structures"},
		"calculus":              {"linear algebra"},
		"compilers":             {"data structures", "formal languages", "computer organization"},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"operating systems":     {"data structures", "computer organization"},
		"networks":              {"operating systems"},
		"programming languages": {"data structures", "computer organization"}}
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

}
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			//fmt.Println(item)
			//fmt.Println(seen[item])
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
				//fmt.Println(order)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
