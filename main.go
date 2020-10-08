package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Record , comment
type Record struct {
	Id      string `json:"Id"`
	Name    string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Records, comment
var Records []Record

func createNewRecord(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var record Record
	json.Unmarshal(reqBody, &record)
	Records = append(Records, record)
	json.NewEncoder(w).Encode(record)
	// fmt.Fprintf(w, "%+v", string(reqBody))
}

func deleteRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, record := range Records {
		fmt.Println(index)
		if record.Id == id {
			Records = append(Records[:index], Records[index+1:]...)
		}
	}
}

func returnSingleRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key+"\n")

	for _, record := range Records {
		if record.Id == key {
			json.NewEncoder(w).Encode(record)
		}
	}
}

func returnAllRecords(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: return all")
	json.NewEncoder(w).Encode(Records)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
	fmt.Println("Endpoint Hit: homepage")
}

func handleRequests() {

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homePage)
	r.HandleFunc("/all", returnAllRecords)
	r.HandleFunc("/record", createNewRecord).Methods("POST")
	r.HandleFunc("/record/{id}", deleteRecord).Methods("DELETE")
	r.HandleFunc("/record/{id}", returnSingleRecord)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	fmt.Println("API INIT")
	Records = []Record{
		Record{Id: "1", Name: "This", Desc: "This DESC", Content: "This Content"},
		Record{Id: "2", Name: "This", Desc: "This DESC again", Content: "This Content again..."},
		Record{Id: "3", Name: "This", Desc: "This DESC again", Content: "This Content again..."},
		Record{Id: "4", Name: "This", Desc: "This DESC again", Content: "This Content again..."},
		Record{Id: "5", Name: "This", Desc: "This DESC again", Content: "This Content again..."},
	}
	handleRequests()
}
