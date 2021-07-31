package utils

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var vowels = []rune("bcdfghjklmnpqrstvwxyz")
var consonants = []rune("aeiou")

func randName() string {
	b := make([]rune, 2)
	b[0] = consonants[rand.Intn(len(consonants))]
	b[1] = vowels[rand.Intn(len(vowels))]
	return string(b)
}
func randYearOfBirth() string {
	return strconv.Itoa(rand.Intn(2005-1950) + 1950)
}

// var outNames chan string
// func receiveResp() {
// 	resp := <-out
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	if true {
// 		fmt.Println(body)
// 	}
// 	defer resp.Body.Close()
// 	outNames <- string(body)
// }

func getNames(w http.ResponseWriter, r *http.Request) {
	// defer Elapsed(time.Now(), "/names")
	w.Header().Set("Content-Type", "application/json")

	nameSub := r.URL.Query().Get("name")
	yearOfBirth, err := strconv.Atoi(r.URL.Query().Get("yearOfBirth"))
	if err != nil {
		fmt.Println(err)
	}
	result := findByNameAndYearOfBirth(nameSub, yearOfBirth)
	var resultMaxLen int
	if len(result) > 3 {
		resultMaxLen = 3
	} else {
		resultMaxLen = len(result)
	}
	json.NewEncoder(w).Encode(result[:resultMaxLen])
}

func getRespBody(url string) string {
	// defer Elapsed(time.Now(), "  getRespBody")
	// fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching: %v", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	// <-in
	return string(body)
}

func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}
func MakeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/names", getNames).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
