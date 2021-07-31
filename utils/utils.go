package utils

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var records [][]string
var allRows []Row

func importThings(recordsMir [][]string, allRowsMir []Row) {
	records = recordsMir
	allRows = allRowsMir
}

var mu = &sync.Mutex{}

func curFuncName() string {
	fpcs := make([]uintptr, 1)

	// Skip 2 levels to get the caller
	n := runtime.Callers(2, fpcs)
	if n == 0 {
		fmt.Println("MSG: NO CALLER")
	}

	caller := runtime.FuncForPC(fpcs[0] - 1)
	if caller == nil {
		fmt.Println("MSG CALLER WAS NIL")
	}

	// Print the name of the function
	return caller.Name()
}
func Elapsed(start time.Time, name string) {
	ms := int(int64(time.Since(start) / time.Millisecond))
	var timeEnd string
	if ms < 1000 {
		timeEnd = rjust(strconv.Itoa(ms)+"ms", 10, " ")
	} else {
		timeEnd = rjust(fmt.Sprint(float32(ms)/1000)+"s", 10, " ")
	}
	fmt.Printf(">%s %s<\n", ljust(name, 40, " "), timeEnd)
}
func rjust(s string, n int, fill string) string {
	fillNum := n - len(s)
	if fillNum < 0 {
		fillNum = 0
	}
	return strings.Repeat(fill, fillNum) + s
}

func ljust(s string, n int, fill string) string {
	fillNum := n - len(s)
	if fillNum < 0 {
		fillNum = 0
	}
	return s + strings.Repeat(fill, fillNum)
}

func center(s string, n int, fill string) string {
	div := n / 2

	return strings.Repeat(fill, div) + s + strings.Repeat(fill, div)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

type Row struct {
	StateCode   string `json:"stateCode"`
	Sex         byte   `json:"sex"`
	YearOfBirth int    `json:"yearOfBirth"`
	Name        string `json:"name"`
	Number      int    `json:"number"`
	Id          int    `pg:",pk" json:"id"`
}

func writeCsvFile(filePath string, records [][]string) {
	f, err := os.Create("employee.csv")
	if err != nil {
		log.Fatal("Unable to create file"+filePath, err)
	}
	f.Close()

	csvWriter := csv.NewWriter(f)

	for _, r := range records {
		_ = csvWriter.Write(r)
	}
	csvWriter.Flush()
}

func NewRow(s []string, i int) Row {
	StateCode := &s[0]
	Sex := s[1][0]
	YearOfBirth, _ := strconv.Atoi(s[2])
	Name := &s[3]
	Number, _ := strconv.Atoi(s[4])
	ID := i
	return Row{*StateCode, Sex, int(YearOfBirth), *Name, int(Number), ID}
}

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

func CallMap() {
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
