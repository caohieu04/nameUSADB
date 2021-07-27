package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

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
func elapsed(start time.Time, name string) {
	timeEnd := time.Since(start)
	log.Printf("%s took %s", name, timeEnd)
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
	StateCode   string
	Sex         byte
	YearOfBirth int
	Name        string
	Number      int
	Id          int `pg:",pk"`
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
