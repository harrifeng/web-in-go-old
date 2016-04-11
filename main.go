package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"createdon"`
}

var noteStore = make(map[string]Note)

var id int = 0

func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	// var note Note

	bufferSize := 100

	p := make([]byte, bufferSize)
	len, err2 := r.Body.Read(p)
	fmt.Println(len)
	if err2 != nil {
		fmt.Println("error is ", err2.Error())
	}

	for i := 0; i < len; i++ {
		fmt.Printf("%c", p[i])
	}

	// err := json.NewDecoder(r.Body).Decode(&note)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// note.CreatedOn = time.Now()
	// id++
	// k := strconv.Itoa(id)
	// noteStore[k] = note
	//
	// j, err := json.Marshal(note)
	// if err != nil {
	// 	panic(err)
	// }
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// w.Write(j)
}

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note

	for _, v := range noteStore {
		notes = append(notes, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()

	fmt.Println()
	os.Exit(0)
}
