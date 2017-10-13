// bit project main.go
package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct{ words []uint64 }

func main() {
	//	fmt.Println("Hello World!")
	var x IntSet
	//	x.Add(0)
	x.Add(1)
	x.Add(2)
	x.Add(3)
	//	x.Add(9)
	//	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println(x)
	fmt.Println(32 << (^uint(0) >> 63))
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	fmt.Println(word)
	fmt.Println(bit)
	fmt.Println(1 << bit)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte('}')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
