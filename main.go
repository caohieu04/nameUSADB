package main

import (
	"fmt"
	"nameUSADB/utils"
	"sync"
	"time"
)

func main() {

	db := utils.DBConnect()
	allRows := utils.DBGetAll()
	Root := utils.DBBuildTrie(allRows)
	if false {
		fmt.Println(Root)
	}
	defer db.Close()

	wg := new(sync.WaitGroup)
	go func() {
		defer wg.Done()
		wg.Add(1)
		utils.MakeRouter()
	}()
	time.Sleep(time.Millisecond)

	go func() {
		defer wg.Done()
		wg.Add(1)
		utils.SiegeMake(0, 1)
	}()
	wg.Wait()
}
