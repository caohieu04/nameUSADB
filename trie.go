package main

import (
	"strings"
)

type Node struct {
	value rune
	to    map[rune]*Node
	ids   []int
}

func NewNode(value rune) Node {
	node := Node{}
	node.value = value
	node.to = make(map[rune]*Node)
	node.ids = []int{}
	return node
}

var Root Node = NewNode('*')

func AddNodeTrie(ptr *Node, char rune, id int) *Node {
	val, ok := ptr.to[char]
	if !ok {
		node := NewNode(char)
		ptr.to[char] = &node
		val = &node
	}
	ptr = val
	val.ids = append(val.ids, id)
	return ptr
}

func makeTrie(strs []string) {
	// records := readCsvFile("name.csv")
	// fmt.Println(records[0], records[1])

	// strs := []string{"abc", "abx", "abcd"}
	all := []struct {
		string
		int
	}{}
	for id, str := range strs {
		for _, strField := range strings.Fields(str) {
			for j, _ := range strField {
				all = append(all, struct {
					string
					int
				}{strField[j:], id + 1})
			}
		}
	}

	for _, s := range all {
		start := &Root
		for _, char := range s.string {
			start = AddNodeTrie(start, char, s.int)
		}
	}
}
