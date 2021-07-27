package main

import (
	"fmt"
	"strings"
	"time"
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

func findByName(t string) []int {
	u := &Root
	for _, char := range t {
		val, ok := u.to[char]
		if !ok {
			fmt.Println("!!Find")
			return nil
		}
		u = val
	}
	return u.ids
}

func findByNameAndYearOfBirth(t string, a int) []Row {
	defer elapsed(time.Now(), "findByNameAndYearOfBirth")
	ids := findByName(t)
	fmt.Println("Length of findByName", len(ids))

	re := []Row{}
	for _, id := range ids {
		func(YearOfBirth int, id int) {
			if YearOfBirth == a {
				re = append(re, allRows[id-1])
			}
		}(allRows[id-1].YearOfBirth, id)
	}

	return re
}

func buildTrie() string {
	strs := []string{}
	for _, r := range records[1:] {
		strs = append(strs, r[3])
	}
	start := time.Now()
	makeTrie(strs)
	fmt.Println(time.Since(start).String())

	// start = time.Now()
	// re := findByName("M")
	// fmt.Println(time.Since(start).String())

	start = time.Now()
	re := findByNameAndYearOfBirth("M", 2001)
	fmt.Println(time.Since(start).String())

	if false {
		fmt.Println(re)
	}
	return "OK"
}
func dbBuildTrie(rows []Row) string {
	defer elapsed(time.Now(), "dbBuildTrie")
	all := []struct {
		string
		int
	}{}
	for _, row := range rows {
		for j, _ := range row.Name {
			all = append(all, struct {
				string
				int
			}{row.Name[j:], row.Id})
		}
	}

	for _, s := range all {
		start := &Root
		for _, char := range s.string {
			start = AddNodeTrie(start, char, s.int)
		}
	}
	return "OK"
}
