package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"
)

type stubMapping map[string]interface{}

var StubStorage = stubMapping{}

func Call(funcName string, params ...interface{}) (result interface{}, err error) {
	f := reflect.ValueOf(StubStorage[funcName])
	if len(params) != f.Type().NumIn() {
		err = errors.New("the number of params is out of index")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	fmt.Println("==Calling ", funcName)
	var res []reflect.Value = f.Call(in)
	result = res[0].Interface()
	fmt.Println("====================")
	return
}

var records [][]string

func readNameCSV() string {
	defer elapsed(time.Now(), "readNameCSV")
	records = readCsvFile("name.csv")
	fmt.Println("Length of records", len(records))
	return "OK"
}

var allRows []Row

func main() {
	if false {
		StubStorage = map[string]interface{}{
			"readNameCSV":   readNameCSV,
			"buildTrie":     buildTrie,
			"dbCreateTable": dbCreateTable,
		}

		var line string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line = scanner.Text()
			res, _ := Call(line)
			if false {
				fmt.Println(res)
			}
		}
	}
	// readNameCSV()
	dbConnect()
	// dbCreateTable()
	allRows = dbGetAll()
	dbBuildTrie(allRows)
	q1 := findByNameAndYearOfBirth("nn", 2001)
	fmt.Println(q1[0])
	q2 := dbFindByNameAndYearOfBirth("Anna", 2001)
	fmt.Println(q2[0])
	defer db.Close()
}
