package utils

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	_ "github.com/go-pg/pg/v10"
)

var db *pg.DB

func saveToDB(l int, r int) {
	db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "123",
		Database: "name_usa",
	})
	defer db.Close()
	obs := []interface{}{}
	for i, row := range records[l:r] {
		ob := NewRow(row, i+l)
		obs = append(obs, &ob)

	}
	_, err := db.Model(obs...).Insert()
	if err != nil {
		panic(err)
	}
}
func writeFromCSVtoDB() {
	defer Elapsed(time.Now(), "writeFromCSVtoDB")
	fmt.Println("Writing ", curFuncName())
	max := 10
	for t := 0; t < max; t++ {
		fmt.Println("[Chunk] ", t)

		l := len(records) / max * t
		if l == 0 {
			l = 1
		}
		r := len(records) / max * (t + 1)
		if r > len(records) {
			r = len(records)
		}
		fmt.Println("Prechunk", l, r)
		saveToDB(l, r)
	}
}

func DBConnect() *pg.DB {
	db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "123",
		Database: "name_usa",
	})
	return db
}
func dbCreateTable() string {
	db.Exec("DROP TABLE ROWS;")
	var rows []Row
	exists, _ := db.Model(&rows).Where("Id = ?", 1).Exists()
	if !exists {
		fmt.Println("Create new table..")
		for _, model := range []interface{}{&Row{}} {
			err := db.Model(model).CreateTable(&orm.CreateTableOptions{
				Temp:          false,
				FKConstraints: true,
			})
			if err != nil {
				// panic(err)
			}
		}
	}
	writeFromCSVtoDB()
	db.Exec(`create index idx_name on rows(name)`)
	db.Exec(`create index idx_year_of_birth on rows(year_of_birth)`)
	db.Exec(`create index idx_id on rows(id)`)
	return "OK"
}
func DBGetAll() []Row {
	defer Elapsed(time.Now(), "dbGetAll")
	var rows []Row
	err := db.Model(&rows).Select()
	if err != nil {
		panic(err)
	}
	allRows = rows
	return rows
}

func dbFindByNameAndYearOfBirth(s string, i int) []Row {
	defer Elapsed(time.Now(), "dbFindByNameAndYearOfBirth")
	var rows []Row
	err := db.Model(&rows).Where("name = ? AND year_of_birth = ?", s, i).Select()
	if err != nil {
		panic(err)
	}
	return rows
}

func readNameCSV() string {
	defer Elapsed(time.Now(), "readNameCSV")
	records = readCsvFile("data/name.csv")
	fmt.Println("Length of records", len(records))
	return "OK"
}
func DBBuildTrie(rows []Row) Node {
	defer Elapsed(time.Now(), "dbBuildTrie")
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
	return Root
}
