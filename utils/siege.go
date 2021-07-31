package utils

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

var in chan string
var out chan string
var quit chan bool

func SiegeMake(limNum int, limSec int) {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	const worker = 12
	in = make(chan string, 2*worker)
	out = make(chan string, 2*worker)
	quit = make(chan bool, 2)
	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range in {
				out <- getRespBody(url)
			}
		}()
	}

	rand.Seed(time.Now().UnixNano())
	go SiegeSendAll(limNum, limSec, wg)
	go SiegeReceiveAll(limNum, limSec, wg)
	wg.Wait()
	close(in)
	close(out)
}

func SiegeRand() string {
	// defer Elapsed(time.Now(), "SiegeRand")
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000/names", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("name", randName())
	q.Add("yearOfBirth", randYearOfBirth())
	req.URL.RawQuery = q.Encode()

	// resp, err = http.Get(req.URL.String())
	url := req.URL.String()
	return url
}
func SiegeSendAll(limNum int, limSec int, wg2 *sync.WaitGroup) {
	defer Elapsed(time.Now(), "SiegeSendAll")
	defer wg2.Done()
	timeStart := time.Now()
	var timeEnd time.Duration
	cnt := 0
	for {
		cnt++
		timeEnd = time.Since(timeStart)
		select {
		case <-quit:
			fmt.Println("Send ", cnt, "requests")
			return
		default:
			url := SiegeRand()
			in <- url
		}
		if limNum != 0 && cnt >= limNum {
			fmt.Println("Send ", cnt, "requests")
			return
		}
		if limSec != 0 && timeEnd.Seconds() >= float64(limSec) {
			fmt.Println("Send ", cnt, "requests")
			return
		}
	}
}

func SiegeReceiveAll(limNum int, limSec int, wg2 *sync.WaitGroup) {
	defer Elapsed(time.Now(), "SiegeReceiveAll")
	defer wg2.Done()
	timeStart := time.Now()
	var timeEnd time.Duration
	cnt := 0
	for {
		timeEnd = time.Since(timeStart)
		if body, ok := <-out; ok {
			cnt += 1
			if false {
				fmt.Println(body)
			}
		}
		if limNum != 0 && cnt >= limNum {
			fmt.Println("Received ", cnt, "requests")
			quit <- true
			return
		}
		if limSec != 0 && timeEnd.Seconds() >= float64(limSec) {
			fmt.Println("Received ", cnt, "requests")
			quit <- true
			return
		}
	}
}
