package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

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

type Row struct {
	StateCode   string
	Sex         byte
	YearOfBirth int
	Name        string
	Number      int
}

func NewRow(s []string) Row {
	StateCode := &s[0]
	Sex := s[1][0]
	YearOfBirth, _ := strconv.ParseInt(s[2], 0, 32)
	Name := &s[3]
	Number, _ := strconv.ParseInt(s[4], 0, 32)
	return Row{*StateCode, Sex, int(YearOfBirth), *Name, int(Number)}
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

func findByNameAndYearOfBirth(t string, a int, records [][]string) []int {
	ids := findByName(t)
	fmt.Println("Length of findByName", len(ids))
	re := []int{}
	for _, id := range ids {
		s := records[id]
		YearOfBirth, _ := strconv.Atoi(s[2])
		// fmt.Println(YearOfBirth, s[2])
		if YearOfBirth == a {
			re = append(re, id)
		}
	}
	return re
}
func main() {

	records := readCsvFile("name.csv")
	strs := []string{}
	fmt.Println("Length of records", len(records))
	for _, r := range records[1:] {
		strs = append(strs, r[3])
	}
	start := time.Now()

	makeTrie(strs)
	fmt.Println(time.Since(start).String())

	start = time.Now()
	re := findByName("M")
	fmt.Println(time.Since(start).String())

	start = time.Now()
	re = findByNameAndYearOfBirth("M", 2001, records)
	fmt.Println(time.Since(start).String())

	if false {
		fmt.Println(re)
	}
}
